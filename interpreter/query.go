// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/*
 * This file contains the defnition of query of octopus
 */

//Query has the interpreted query info
type Query struct {
	//Tables has the map of tables whose data is being accessed by the query
	Tables map[string]TableNode `json:"tables,omitempty"`
	//Select has the list of columns to be selected from the data
	Select []ColumnNode `json:"select,omitempty"`
	//GroupBy has the list of columns to be used for group the data
	GroupBy []ColumnNode `json:"group_by,omitempty"`
	//Filters has the list of filters applied in the query
	Filters []OperatorNode `json:"filters,omitempty"`
	//Result has the result of the query
	Result interface{} `json:"result,omitempty"`
}

//SQLQuery stores a sql query to be executed
type SQLQuery struct {
	//Query is the query string with arguments
	Query string
	//Args has the arguments to be passed to the query string
	Args []interface{}
}

//ToSQL converts the the query to a sql query
func (q Query) ToSQL() (*SQLQuery, error) {
	/*
	 * We will add check for zero table
	 * If the no of tables is one we will choose the single table query mode
	 */
	if len(q.Tables) == 0 {
		return nil, errors.New("couldn't find any tables")
	}
	if len(q.Tables) == 1 {
		return q.ToSingleTableSQL()
	}
	return nil, errors.New("couldn't convert the query to sql as no of tables not suppoerted")
}

//ToSingleTableSQL will convert the query to sql if the query has only one table
func (q Query) ToSingleTableSQL() (*SQLQuery, error) {
	/*
	 * We will add a table check
	 * Then we will get the table
	 * Then we will iterate through the select fields
	 * If group by fields are there, will add them to be selected
	 * Then we will add the filters
	 * Then we will add the group by if any
	 */
	//adding the table number check
	if len(q.Tables) != 1 {
		return nil, fmt.Errorf("couldn't find any table. Length of the tables to which the query belong to is %d", len(q.Tables))
	}

	//getting the table
	var tableNode TableNode
	for _, v := range q.Tables {
		tableNode = v
	}

	var queryB strings.Builder
	queryB.WriteString("SELECT ")
	hasGroupBy := false
	for _, v := range q.GroupBy {
		if len(v.Name) > 0 {
			hasGroupBy = true
		}
	}
	//iterating through the fields to be selected
	count := 0
	for _, v := range q.Select {
		if len(v.Name) == 0 {
			continue
		}
		addColumnString(count, v, &queryB, hasGroupBy)
		count++
	}

	//if group fields are there, then add then for selection
	for _, v := range q.GroupBy {
		if len(v.Name) == 0 {
			continue
		}
		addColumnString(count, v, &queryB, false)
		count++
	}

	//adding the filters
	queryB.WriteString(" FROM \"" + tableNode.Name + "\"")
	count = 0
	values := []interface{}{}
	done := false
	index := 0
	for _, v := range q.Filters {
		if v.Column == nil || len(v.Column.Name) == 0 || (v.Value == nil && v.Unknown == nil && v.Time == nil) {
			continue
		}
		if v.Unknown == nil && v.Time == nil && len(v.Value.Name) == 0 {
			continue
		}
		if v.Value == nil && v.Time == nil && len(v.Unknown.Word) == 0 {
			continue
		}
		if v.Value == nil && v.Unknown == nil && (!v.Time.Value.IsValid() || !v.Time.Value.From.IsValid()) {
			continue
		}
		if (v.Column.DataType == DataTypeInt || v.Column.DataType == DataTypeFloat || v.Column.DataType == DataTypeDate) &&
			(v.Operation != EqOperator && v.Operation != NotEqOperator && v.Operation != GreaterOperator && v.Operation != LessOperator) {
			continue
		}
		var convertedVal interface{}
		if v.Value != nil {
			vl, ok := getValue(v.Column.DataType, v.Value.Name)
			if !ok {
				continue
			}
			convertedVal = vl
		} else if v.Unknown != nil {
			vl, ok := getValue(v.Column.DataType, string(v.Unknown.Word))
			if !ok {
				continue
			}
			convertedVal = vl
		} else if v.Time != nil {
			if v.Column.DataType != DataTypeDate {
				continue
			}
			convertedVal = v.Time.Value.From.Time
		}

		if len(q.Filters) > 0 && !done {
			queryB.WriteString(" WHERE ")
			done = true
		}
		if count != 0 {
			queryB.WriteString(" AND ")
		}
		count++
		columnName := "\"" + v.Column.Name + "\""
		value := ""
		if v.Column.DataType == DataTypeString {
			value += "'"
		}
		if (v.Operation == LikeOperator || v.Operation == ContainsOperator) && v.Column.DataType == DataTypeString {
			value += "%"
		}
		if v.Column.DataType == DataTypeString {
			value += convertedVal.(string)
		}
		if (v.Operation == LikeOperator || v.Operation == ContainsOperator) && v.Column.DataType == DataTypeString {
			value += "%"
		}
		if v.Column.DataType == DataTypeString {
			value += "'"
		}
		var finalValue interface{} = value
		if v.Column.DataType != DataTypeString {
			finalValue = convertedVal
		}
		index++
		queryB.WriteString(columnName + " " + v.Operation + " $" + strconv.Itoa(index))
		values = append(values, finalValue)
	}

	//add the group by if required
	result := &SQLQuery{Args: values}
	if !hasGroupBy {
		result.Query = queryB.String()
		return result, nil
	}
	count = 0
	queryB.WriteString(" GROUP BY ")
	for _, v := range q.GroupBy {
		if len(v.Name) == 0 {
			continue
		}
		if count != 0 {
			queryB.WriteString(", ")
		}
		count++
		queryB.WriteString("\"" + v.Name + "\"")
	}
	result.Query = queryB.String()

	return result, nil
}

func addColumnString(i int, v ColumnNode, qS *strings.Builder, enforceGroupBy bool) {
	if i != 0 {
		qS.WriteString(", ")
	}
	columnName := "\"" + v.Name + "\""
	if enforceGroupBy && len(v.AggregationFn) != 0 {
		columnName = v.AggregationFn + "(" + columnName + ") AS " + columnName
	} else if enforceGroupBy && len(v.AggregationFn) == 0 {
		columnName = DefaultAggregationFn + "(" + columnName + ") AS " + columnName
	}
	qS.WriteString(columnName + " ")
}

func getValue(dataType string, value string) (interface{}, bool) {
	if dataType == DataTypeDate || dataType == DataTypeString {
		return value, true
	}
	if dataType == DataTypeInt {
		val, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return nil, false
		}
		return val, true
	}
	if dataType == DataTypeFloat {
		val, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, false
		}
		return val, true
	}
	return value, false
}
