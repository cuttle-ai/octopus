// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package dict has the implementation of the dictionary api for the server
package dict

import (
	"context"
	"net/http"
	"strings"

	"github.com/cuttle-ai/octopus/interpreter"
	"github.com/cuttle-ai/octopus/lsp/routes"
	"github.com/cuttle-ai/octopus/lsp/routes/response"
)

var testCollection = &interpreter.TableNode{UID: "automobile-sales", Word: []rune("Automobile sales")}

var testColumn = &interpreter.ColumnNode{UID: "car", PUID: "automobile-sales", PN: testCollection, Word: []rune("car"), DataType: interpreter.DataTypeString}

var testColumn1 = &interpreter.ColumnNode{UID: "brand", PUID: "automobile-sales", PN: testCollection, Word: []rune("brand"), Dimension: true, DataType: interpreter.DataTypeString}

var testColumn2 = &interpreter.ColumnNode{UID: "year", PUID: "automobile-sales", PN: testCollection, Word: []rune("year"), Dimension: true, DataType: interpreter.DataTypeDate}

var testColumn3 = &interpreter.ColumnNode{UID: "sales", PUID: "automobile-sales", PN: testCollection, Word: []rune("sales"), Measure: true, DataType: interpreter.DataTypeInt}

var testValue = &interpreter.ValueNode{UID: "swift", Word: []rune("Swift"), PN: testColumn, PUID: "car"}

var testOperator = &interpreter.OperatorNode{UID: "equal-is", Word: []rune("is"), Operation: interpreter.EqOperator}

var notEqOperator = &interpreter.OperatorNode{UID: "not-equal", Word: []rune("not"), Operation: interpreter.NotEqOperator}

var greaterThanOperator = &interpreter.OperatorNode{UID: "greater-than", Word: []rune(">="), Operation: interpreter.LessOperator}

var lessThanOperator = &interpreter.OperatorNode{UID: "less-than", Word: []rune("<="), Operation: interpreter.GreaterOperator}

func init() {
	testCollection.DefaultDateField = testColumn2
	testCollection.DefaultDateFieldUID = testColumn2.UID
	testCollection.Children = []interpreter.ColumnNode{*testColumn, *testColumn1, *testColumn2, *testColumn3}
}

var testTokens = map[string]interpreter.Token{
	"year": interpreter.Token{
		Word:  []rune("year"),
		Nodes: []interpreter.Node{testColumn2},
	},
	"sales": interpreter.Token{
		Word:  []rune("sales"),
		Nodes: []interpreter.Node{testColumn3},
	},
	"brand": interpreter.Token{
		Word:  []rune("brand"),
		Nodes: []interpreter.Node{testColumn1},
	},
	"car": interpreter.Token{
		Word:  []rune("car"),
		Nodes: []interpreter.Node{testColumn},
	},
	"Swift": interpreter.Token{
		Word:  []rune("Swift"),
		Nodes: []interpreter.Node{testValue},
	},
	"is": interpreter.Token{
		Word:  []rune("is"),
		Nodes: []interpreter.Node{testOperator},
	},
	"since": interpreter.Token{
		Word:  []rune("since"),
		Nodes: []interpreter.Node{greaterThanOperator},
	},
	"not": interpreter.Token{
		Word:  []rune("not"),
		Nodes: []interpreter.Node{notEqOperator},
	},
	"before": interpreter.Token{
		Word:  []rune("before"),
		Nodes: []interpreter.Node{lessThanOperator},
	},
	"<": interpreter.Token{
		Word:  []rune("<"),
		Nodes: []interpreter.Node{lessThanOperator},
	},
	">": interpreter.Token{
		Word:  []rune(">"),
		Nodes: []interpreter.Node{greaterThanOperator},
	},
}

var testDICT = interpreter.DICT{Map: testTokens}

//TestUser for the lsp
var TestUser = "testuser"

func loadTestDICT() {
	tTok := map[string]interpreter.Token{}
	for k, v := range testDICT.Map {
		tTok[strings.ToLower(k)] = v
	}
	testDICT.Map = tTok
	req := interpreter.DICTRequest{
		ID:   TestUser,
		Type: interpreter.DICTAdd,
		DICT: testDICT,
	}
	interpreter.SendDICTToChannel(interpreter.DICTInputChannel, req)
}

func init() {
	loadTestDICT()
}

//GetDict will return the dictionary being used
func GetDict(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := interpreter.DICTRequest{ID: TestUser, Type: interpreter.DICTGet, Out: make(chan interpreter.DICTRequest)}
	go interpreter.SendDICTToChannel(interpreter.DICTInputChannel, req)
	res := <-req.Out
	response.Write(w, res.DICT)
}

func init() {
	routes.AddRoutes(
		routes.Route{
			Version:     "v1",
			Pattern:     "/dict",
			HandlerFunc: GetDict,
		},
	)
}
