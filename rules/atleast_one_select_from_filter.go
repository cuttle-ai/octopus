// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package rules

import "github.com/cuttle-ai/octopus/interpreter"

/*
 * This file contains the rule defnition for adjusting the query so that the query will have atleast one select field with borrowing from filter's column
 */

//AtleastOneColumnFromFilter will check whether the query has atleast one select if none available and if atleast one filter available, will select its column
var AtleastOneColumnFromFilter = interpreter.Rule{
	Name:        "Atleast one select with borrow from operator",
	Description: "This will check whether the query has atleast one select if none available and if atleast one filter available, will select its column",
	Template:    []interpreter.Type{interpreter.Operator},
	Resolve: func(qu interpreter.Query, toks []interpreter.FastToken, index int) (interpreter.Query, error) {
		/*
		 * If the query doesn't have any select fields and since has a filter, we will add its column to the select
		 */
		if len(qu.Select) != 0 || len(qu.Filters) == 0 || !qu.Filters[0].IsResolved() {
			return qu, nil
		}
		qu.Select = []interpreter.ColumnNode{*qu.Filters[0].Column}

		return qu, nil
	},
}

//AtleastOneColumnFromValue will check whether the query has atleast one select if none available and if atleast one filter available, will select its column
var AtleastOneColumnFromValue = interpreter.Rule{
	Name:        "Atleast one select with borrow from value",
	Description: "This will check whether the query has atleast one select if none available and if atleast one filter available, will select its column",
	Template:    []interpreter.Type{interpreter.Value},
	Resolve: func(qu interpreter.Query, toks []interpreter.FastToken, index int) (interpreter.Query, error) {
		/*
		 * If the query doesn't have any select fields and since has a filter, we will add its column to the select
		 */
		if len(qu.Select) != 0 || len(qu.Filters) == 0 || !qu.Filters[0].IsResolved() {
			return qu, nil
		}
		qu.Select = []interpreter.ColumnNode{*qu.Filters[0].Column}

		return qu, nil
	},
}
