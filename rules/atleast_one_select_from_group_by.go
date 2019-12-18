// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package rules

import (
	"github.com/cuttle-ai/octopus/interpreter"
)

/*
 * This file contains the rule defnition for adjusting the query so that the query will have atleast one select field with borrowing from group by column
 */

//AtleastOneColumnFromGroupBy will check whether the query has atleast one select if none available and if atleast one group by available, will select it from that
var AtleastOneColumnFromGroupBy = interpreter.Rule{
	Name:        "Atleast one select with borrow from group by",
	Description: "This will check whether the query has atleast one select if none available and if atleast one group by available, will select it from that",
	Template:    []interpreter.Type{interpreter.Column},
	Resolve: func(qu interpreter.Query, toks []interpreter.FastToken, index int) (interpreter.Query, error) {
		/*
		 * If the query doesn't have any select fields and since has a group by field, we will add it to the select
		 */
		if len(qu.Select) != 0 || len(qu.GroupBy) == 0 {
			return qu, nil
		}
		qu.Select = []interpreter.ColumnNode{qu.GroupBy[0]}
		qu.GroupBy = qu.GroupBy[1:]

		return qu, nil
	},
}
