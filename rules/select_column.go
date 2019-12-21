// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package rules

import "github.com/cuttle-ai/octopus/interpreter"

/*
 * This file contains the rule defnition for identifying the columns in the query as select fields
 */

//SelectColumn will mark all the fields in the query as select columns
var SelectColumn = interpreter.Rule{
	Name:        "Select Columns",
	Description: "This rule will mark all the remaining columns in the query as columns to be selected in a SQL query",
	Template:    []interpreter.Type{interpreter.Column},
	Resolve: func(qu interpreter.Query, toks []interpreter.FastToken, index int) (interpreter.Query, error) {
		/*
		 * If the column in the given index is not resolved we will add it to the query as select column and mark it as resolved
		 */
		if index >= len(toks) || len(toks[index].Columns) == 0 {
			//we don't have enough the tokens for the given index
			return qu, nil
		}
		if toks[index].Columns[0].IsResolved() {
			//the column is already resolved
			return qu, nil
		}
		toks[index].Columns[0].SetResolved(true)
		if len(qu.Select) == 0 {
			qu.Select = []interpreter.ColumnNode{}
		}
		qu.Select = append(qu.Select, toks[index].Columns[0])
		qu.Tables[toks[index].Columns[0].PUID] = *((toks[index].Columns[0].PN.Copy()).(*interpreter.TableNode))

		return qu, nil
	},
}
