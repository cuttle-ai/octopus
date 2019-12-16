// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package interpreter has the implementation of the interpreter api for the server
package interpreter

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cuttle-ai/octopus/interpreter"
	"github.com/cuttle-ai/octopus/lsp/routes"
	"github.com/cuttle-ai/octopus/lsp/routes/dict"
	"github.com/cuttle-ai/octopus/lsp/routes/response"
)

//Query struct as input for interperter
type Query struct {
	//NL is the natural language query
	NL string `json:"nl,omitempty"`
}

//Interpret will interpret a given natural language query
func Interpret(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	rq := &Query{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(rq)
	if err != nil {
		//error while decoding the request param
		response.WriteError(w, response.Error{Err: err.Error()}, http.StatusBadRequest)
		return
	}
	toks, err := interpreter.Tokenize(dict.TestUser, []rune(rq.NL))
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
