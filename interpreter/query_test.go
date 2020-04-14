// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter_test

import (
	"strings"
	"testing"

	"github.com/cuttle-ai/octopus/interpreter"
)

/*
 * This file contains the tests for the functions in the query file
 */

type query struct {
	q interpreter.Query
	s string
	n string
}

var testTable = interpreter.TableNode{Name: "stores"}

var testColumnCity = interpreter.ColumnNode{Name: "city", DataType: interpreter.DataTypeString}

var testUnknownCity = interpreter.UnknownNode{Word: []rune("Delhi")}

var testQueries = []query{
	{
		q: interpreter.Query{
			Select: []interpreter.ColumnNode{
				testColumnCity,
			},
			Tables: map[string]interpreter.TableNode{
				"1": testTable,
			},
		},
		s: `SELECT "city" AS "city" FROM "stores"`,
		n: "normal select",
	},
	{
		q: interpreter.Query{
			Select: []interpreter.ColumnNode{
				testColumnCity,
			},
			Tables: map[string]interpreter.TableNode{
				"1": testTable,
			},
			Filters: []interpreter.OperatorNode{
				{
					Operation: interpreter.EqOperator,
					Column:    &testColumnCity,
					Unknown:   &testUnknownCity,
				},
			},
		},
		s: `SELECT "city" AS "city" FROM "stores" WHERE "city" = ?`,
		n: "normal select with filters",
	},
}

func TestToSQL(t *testing.T) {
	for _, v := range testQueries {
		t.Run(v.n, func(t *testing.T) {
			s, err := v.q.ToSQL()
			if err != nil {
				t.Error("error while running the test", err)
				return
			}
			if strings.Compare(s.Query, v.s) != 0 {
				t.Error("Expected query", "`"+v.s+"`", "got", "`"+s.Query+"`")
			}
		})
	}
}
