// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package rules

import "github.com/cuttle-ai/octopus/interpreter"

/*
 * This file contains the rule defnition for identifying the filter with known values assigned to a column
 */

//ValueFilter will mark all the filter associated in the query with <field> <operator> <value>
var ValueFilter = interpreter.Rule{
	Name:        "Filter with value",
	Description: "This rule will find the filters in the query. It will assign a filter if found in the template <field> <operator> <value>",
	Template:    []interpreter.Type{interpreter.Column, interpreter.Operator, interpreter.Value},
	Resolve: func(qu interpreter.Query, toks []interpreter.FastToken, index int) (interpreter.Query, error) {
		/*
		 * If the column, operator, value in the given index is not resolved we will add it to the query as filters and mark them as resolved
		 */
		if index+2 >= len(toks) || len(toks[index].Columns) == 0 || len(toks[index+1].Operators) == 0 || len(toks[index+2].Values) == 0 {
			//we don't have enough the tokens for the given index
			return qu, nil
		}
		if toks[index].Columns[0].IsResolved() || toks[index+1].Operators[0].IsResolved() || toks[index+2].Values[0].IsResolved() {
			//the column or operator or value is already resolved
			return qu, nil
		}
		if toks[index].Columns[0].UID != toks[index+2].Values[0].PUID {
			//Parent of the value is not the column
			return qu, nil
		}
		toks[index].Columns[0].SetResolved(true)
		toks[index+1].Operators[0].SetResolved(true)
		toks[index+2].Values[0].SetResolved(true)
		toks[index+1].Operators[0].Column = &toks[index].Columns[0]
		toks[index+1].Operators[0].Value = &toks[index+2].Values[0]
		if len(qu.Filters) == 0 {
			qu.Filters = []interpreter.OperatorNode{}
		}
		qu.Filters = append(qu.Filters, toks[index+1].Operators[0])
		qu.Tables[toks[index].Columns[0].PUID] = *((toks[index].Columns[0].PN.Copy()).(*interpreter.TableNode))

		return qu, nil
	},
}
