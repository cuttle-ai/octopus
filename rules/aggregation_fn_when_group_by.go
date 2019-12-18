// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package rules

import "github.com/cuttle-ai/octopus/interpreter"

/*
 * This file contains the rule defnition for adjusting the query so that the query will have atleast one select field with borrowing from group by column
 */

//AggregationFnWhenGroupBy will check whether the query has atleast one group by. If yes, will add aggresgation fns if not available for the select fields
var AggregationFnWhenGroupBy = interpreter.Rule{
	Name:        "Aggregation funtion when group by",
	Description: "This will check whether the query has atleast one group by. If yes, will add aggresgation fns if not available for the select fields",
	Template:    []interpreter.Type{interpreter.GroupBy},
	Resolve: func(qu interpreter.Query, toks []interpreter.FastToken, index int) (interpreter.Query, error) {
		/*
		 * If the query has a group by field, we will add aggregation fn for the select fields which doesn't have one
		 */
		if len(qu.Select) == 0 || len(qu.GroupBy) == 0 {
			return qu, nil
		}
		for i := len(qu.Select) - 1; i >= 0; i-- {
			if len(qu.Select[i].AggregationFn) != 0 {
				continue
			}
			if qu.Select[i].DataType == interpreter.DataTypeFloat || qu.Select[i].DataType == interpreter.DataTypeInt {
				qu.Select[i].AggregationFn = interpreter.AggregationFnSum
			} else {
				qu.Select[i].AggregationFn = interpreter.AggregationFnCount
			}
		}

		return qu, nil
	},
}
