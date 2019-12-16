// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package interpreter has the implementation of the interpreter api for the server
package interpreter

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cuttle-ai/octopus/interpreter"
	"github.com/cuttle-ai/octopus/lsp/routes"
	"github.com/cuttle-ai/octopus/lsp/routes/response"
)

//Query struct as input for interperter
type Query struct {
	//NL is the natural language query
	NL string `json:"nl,omitempty"`
}

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

var testUser = "testuser"

func loadTestDICT() {
	tTok := map[string]interpreter.Token{}
	for k, v := range testDICT.Map {
		tTok[strings.ToLower(k)] = v
	}
	testDICT.Map = tTok
	req := interpreter.DICTRequest{
		ID:   testUser,
		Type: interpreter.DICTAdd,
		DICT: testDICT,
	}
	interpreter.SendDICTToChannel(interpreter.DICTInputChannel, req)
}

func init() {
	loadTestDICT()
}

//Interpret will interpret a given natural language query
func Interpret(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("query")
	rq := &Query{}
	dec := json.NewDecoder(strings.NewReader(query))
	err := dec.Decode(rq)
	if err != nil {
		//error while decoding the request param
		response.WriteError(w, response.Error{Err: err.Error()}, http.StatusBadRequest)
		return
	}
	toks, err := interpreter.Tokenize(testUser, []rune(rq.NL))
	if err != nil {
		//error while tokenizing the user query
		response.WriteError(w, response.Error{Err: err.Error()}, http.StatusBadRequest)
		return
	}
	ins, err := interpreter.Interpret(toks)
	if err != nil {
		//error while interpreting the user query
		response.WriteError(w, response.Error{Err: err.Error()}, http.StatusBadRequest)
		return
	}
	response.Write(w, ins)
}

func init() {
	routes.AddRoutes(
		routes.Route{
			Version:     "v1",
			Pattern:     "/interpret",
			HandlerFunc: Interpret,
		},
	)
}
