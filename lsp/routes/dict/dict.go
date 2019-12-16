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

var testColumn = &interpreter.ColumnNode{UID: "cars", Word: []rune("cars")}

var testColumn1 = &interpreter.ColumnNode{UID: "brands", Word: []rune("brand")}

var testValue = &interpreter.ValueNode{UID: "swift", Word: []rune("Swift"), PN: testColumn, PUID: "cars"}

var testTokens = map[string]interpreter.Token{
	"cars": interpreter.Token{
		Word:  []rune("cars"),
		Nodes: []interpreter.Node{testColumn},
	},
	"Swift": interpreter.Token{
		Word:  []rune("Swift"),
		Nodes: []interpreter.Node{testValue},
	},
	"brands": interpreter.Token{
		Word:  []rune("brands"),
		Nodes: []interpreter.Node{testColumn1},
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
