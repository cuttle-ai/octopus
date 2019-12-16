// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package rules has the implmentation of the rules related api
package rules

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/cuttle-ai/octopus/interpreter"
	"github.com/cuttle-ai/octopus/lsp/routes"
	"github.com/cuttle-ai/octopus/lsp/routes/response"

	defaultRules "github.com/cuttle-ai/octopus/rules"
)

//RuleDisableStateDO is the data layer object to communicate with the api
type RuleDisableStateDO struct {
	//Position is the position of the group to which the rule belongs to in the rules array
	Position int `json:"position,omitempty"`
	//GroupPosition is the poistion of the group in its group
	GroupPosition int `json:"group_position,omitempty"`
	//DisabledState is the new disabled state of the rule
	DisabledState bool `json:"disabled_state,omitempty"`
}

//Success is the success message response
type Success struct {
	//Message of the response
	Message string
}

func init() {
	defaultRules.LoadDefaultRules()
}

//GetRules will return the list of rules in the interpreter
func GetRules(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	response.Write(w, interpreter.GetRules())
}

//SetRuleDisableState will set the disable state of a rule at the given position and group position
func SetRuleDisableState(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	rq := &RuleDisableStateDO{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(rq)
	if err != nil {
		//error while decoding the request param
		response.WriteError(w, response.Error{Err: err.Error()}, http.StatusBadRequest)
		return
	}
	interpreter.SetRuleDisableState(rq.Position, rq.GroupPosition, rq.DisabledState)
	response.Write(w, Success{"Sucessfull"})
}

func init() {
	routes.AddRoutes(
		routes.Route{
			Version:     "v1",
			Pattern:     "/rules",
			HandlerFunc: GetRules,
		},
		routes.Route{
			Version:     "v1",
			Pattern:     "/rules/state",
			HandlerFunc: SetRuleDisableState,
		},
	)
}
