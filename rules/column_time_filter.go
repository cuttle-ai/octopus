// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package rules

import "github.com/cuttle-ai/octopus/interpreter"

/*
 * This file contains the rule defnition for identifying the filter with time assigned to a column having date datatype
 */

//ColumnTimeFilter will mark all the filter associated in the query with <field> <operator> <time>
var ColumnTimeFilter = interpreter.Rule{
	Name:        "Filter with column and time",
	Description: "This rule will find the filters in the query. It will assign a filter if found in the template <field> <operator> <time> and field has data type date",
	Template:    []interpreter.Type{interpreter.Column, interpreter.Operator, interpreter.Time},
	Resolve: func(qu interpreter.Query, toks []interpreter.FastToken, index int) (interpreter.Query, error) {
		/*
		 * If the column, operator, time in the given index is not resolved we will add it to the query as filters and mark them as resolved
		 */
		if index+2 >= len(toks) || len(toks[index].Columns) == 0 || len(toks[index+1].Operators) == 0 || len(toks[index+2].Times) == 0 {
			//we don't have enough the tokens for the given index
			return qu, nil
		}
		if toks[index].Columns[0].IsResolved() || toks[index+1].Operators[0].IsResolved() || toks[index+2].Times[0].IsResolved() {
			//the column or operator or unknown is already resolved
			return qu, nil
		}
		//if the data type of the column is not date, we will skip
		if toks[index].Columns[0].DataType != interpreter.DataTypeDate {
			return qu, nil
		}
		toks[index].Columns[0].SetResolved(true)
		toks[index+1].Operators[0].SetResolved(true)
		toks[index+2].Times[0].SetResolved(true)
		toks[index+1].Operators[0].Column = &toks[index].Columns[0]
		toks[index+1].Operators[0].Time = &toks[index+2].Times[0]
		if len(qu.Filters) == 0 {
			qu.Filters = []interpreter.OperatorNode{}
		}
		qu.Filters = append(qu.Filters, toks[index+1].Operators[0])
		qu.Tables[toks[index].Columns[0].PUID] = *((toks[index].Columns[0].PN.Copy()).(*interpreter.TableNode))

		return qu, nil
	},
}
