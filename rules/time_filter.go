// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package rules

import "github.com/cuttle-ai/octopus/interpreter"

/*
 * This file contains the rule defnition for identifying the filter with time filter applied to the collection
 */

//TimeFilter will mark all the filter associated in the query with <operator> <time>
var TimeFilter = interpreter.Rule{
	Name: "Filter with date/time",
	Description: "This rule will find the date/time filters in the query. " + "It will assign a filter if found in the template <time>." +
		"Filter will be applied to the default date field of a table. If the default date field is missing, then filter is applied to the available field with date type." +
		"Table is decided based on the parent of the select/group by field. If multiple tables are avaiable in case of a join query, the table with default date/having date fields is selected." +
		"Improvement for this rule is kept for future scope.",
	Template: []interpreter.Type{interpreter.Time},
	Resolve: func(qu interpreter.Query, toks []interpreter.FastToken, index int) (interpreter.Query, error) {
		/*
		 * If the operator, time in the given index is not resolved we will then proceed further
		 * We will select the tables and in the process will select the coumn to apply the filter
		 * Then we will mark nodes as resolved
		 */
		if index >= len(toks) || len(toks[index].Times) == 0 {
			//we don't have enough the tokens for the given index
			return qu, nil
		}
		if toks[index].Times[0].IsResolved() {
			//the operator or time is already resolved
			return qu, nil
		}

		//selecting the table
		//then will iterate through the tables and check for the conditions
		var selectedField *interpreter.ColumnNode
		for _, t := range qu.Tables {
			//first check for the default date field
			//then check for the date type fields
			if t.DefaultDateField != nil {
				selectedField = t.DefaultDateField.Copy().(*interpreter.ColumnNode)
				continue
			}
			for _, f := range t.Children {
				if f.DataType == interpreter.DataTypeDate {
					selectedField = f.Copy().(*interpreter.ColumnNode)
				}
			}
		}

		//if a field is not found, then we will skip
		if selectedField == nil {
			return qu, nil
		}

		//marking the nodes as resolved
		selectedField.SetResolved(true)

		operator := interpreter.OperatorNode{
			UID:       "Operator-" + toks[index].Times[0].UID,
			Resolved:  true,
			Word:      []rune("is"),
			Column:    selectedField,
			Time:      &toks[index].Times[0],
			Operation: interpreter.EqOperator,
		}
		if toks[index].Times[0].Value.Type == "interval" && toks[index].Times[0].Value.From != nil {
			operator.Operation = interpreter.GreaterOperator
			operator.Word = []rune("since")
		} else if toks[index].Times[0].Value.Type == "interval" && toks[index].Times[0].Value.To != nil {
			operator.Operation = interpreter.LessOperator
			operator.Word = []rune("till")
		}
		toks[index].Times[0].SetResolved(true)
		if len(qu.Filters) == 0 {
			qu.Filters = []interpreter.OperatorNode{}
		}
		qu.Filters = append(qu.Filters, operator)

		return qu, nil
	},
}
