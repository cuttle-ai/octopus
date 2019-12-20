// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import "encoding/json"

/*
 * This file contains the defnition value type node
 */

//ValueNode is the node storing the information about a value in a column.
//It can be the child of a Column Node and cannot have any children
type ValueNode struct {
	//UID is the unique id of the value node
	UID string
	//Word is the word with which the value node has to be matched
	Word []rune
	//PN is the parent node of the value. It will be a column node
	PN *ColumnNode
	//PUID is the UID of value's parent node
	PUID string
	//Name is the name of the node
	Name string
	//Resolved indicates that the node is resolved
	Resolved bool
}

type valueNode struct {
	UID      string `json:"uid,omitempty"`
	Word     string `json:"word,omitempty"`
	PUID     string `json:"puid,omitempty"`
	Name     string `json:"name,omitempty"`
	Resolved bool   `json:"resolved,omitempty"`
	Type     string `json:"type,omitempty"`
}

//Copy will return a copy of the node
func (v *ValueNode) Copy() Node {
	return &ValueNode{
		UID:      v.UID,
		Word:     v.Word,
		PN:       v.PN,
		PUID:     v.PUID,
		Name:     v.Name,
		Resolved: v.Resolved,
	}
}

//ID returns the unique id of the node
func (v *ValueNode) ID() string {
	return v.UID
}

//Type returns Value Type
func (v *ValueNode) Type() Type {
	return Value
}

//TokenWord returns the word property of the node
func (v *ValueNode) TokenWord() []rune {
	return v.Word
}

//PID returns the PUID if the node
func (v *ValueNode) PID() string {
	return v.PUID
}

//Parent returns the PN of the node
func (v *ValueNode) Parent() Node {
	return v.PN
}

//MarshalJSON encodes the node into a serializable json
func (v *ValueNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&valueNode{
		v.UID, string(v.Word), v.PUID, v.Name, v.Resolved, "Value",
	})
}

//UnmarshalJSON decodes the node from a json
func (v *ValueNode) UnmarshalJSON(data []byte) error {
	m := &valueNode{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	v.UID = m.UID
	v.Word = []rune(m.Word)
	v.PUID = m.PUID
	v.Name = m.Name
	v.Resolved = m.Resolved
	return nil
}

//IsResolved will return true if the node is resolved
func (v *ValueNode) IsResolved() bool {
	return v.Resolved
}

//SetResolved will set the resolved state of the node
func (v *ValueNode) SetResolved(state bool) {
	v.Resolved = state
}
