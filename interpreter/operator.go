// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import "encoding/json"

/*
 * This file contains the defnition of operator type node
 */

//EqOperator indicates equal to operator
const EqOperator = "="

//NotEqOperator indicates not equal to operator
const NotEqOperator = "!="

//GreaterOperator indicates greater than operator
const GreaterOperator = ">="

//LessOperator indicates less than or equal to operator
const LessOperator = "<="

//ContainsOperator indicates contains operator
const ContainsOperator = "has"

//LikeOperator indicates like operator
const LikeOperator = "%"

//OperatorNode is the node storing the information about a operator.
//Filters are set based on this node. It depends upon a column and value/unknown
type OperatorNode struct {
	//UID is the unique id of the operator node
	UID string
	//Word is the word with which the operator node has to be matched
	Word []rune
	//PUID is the UID of operator's parent node
	PUID string
	//PN is the parent node of the operator. It will be a KnowledgeBase
	PN Node
	//Resolved indicates that the node is resolved
	Resolved bool
	//Column is the column with which the operator is applied
	Column *ColumnNode
	//Unknown is the value to be applied to the column node with the operator
	Unknown *UnknownNode
	//Value is the value to be applied to the column node with the operator
	Value *ValueNode
	//Time is the time node to be applied to the column node with the operator
	Time *TimeNode
	//Operation is the operation applied by the node
	Operation string
}

type operatorNode struct {
	UID       string       `json:"uid,omitempty"`
	Word      string       `json:"word,omitempty"`
	PUID      string       `json:"puid,omitempty"`
	Column    *ColumnNode  `json:"column,omitempty"`
	Unknown   *UnknownNode `json:"unknown,omitempty"`
	Value     *ValueNode   `json:"value,omitempty"`
	Resolved  bool         `json:"resolved,omitempty"`
	Type      string       `json:"type,omitempty"`
	Operation string       `json:"operation,omitempty"`
}

//Copy will return a copy of the node
func (o *OperatorNode) Copy() Node {
	return &OperatorNode{
		UID:       o.UID,
		Word:      o.Word,
		PN:        o.PN,
		PUID:      o.PUID,
		Resolved:  o.Resolved,
		Column:    o.Column,
		Value:     o.Value,
		Unknown:   o.Unknown,
		Operation: o.Operation,
	}
}

//ID returns the unique id of the node
func (o *OperatorNode) ID() string {
	return o.UID
}

//Type returns Operator Type
func (o *OperatorNode) Type() Type {
	return Operator
}

//TokenWord returns the word property of the node
func (o *OperatorNode) TokenWord() []rune {
	return o.Word
}

//PID returns the PUID if the node
func (o *OperatorNode) PID() string {
	return o.PUID
}

//Parent returns the PN of the node
func (o *OperatorNode) Parent() Node {
	return o.PN
}

//MarshalJSON encodes the node into a serializable json
func (o *OperatorNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&operatorNode{
		o.UID, string(o.Word), o.PUID, o.Column, o.Unknown, o.Value, o.Resolved, "Operator", o.Operation,
	})
}

//UnmarshalJSON decodes the node from a json
func (o *OperatorNode) UnmarshalJSON(data []byte) error {
	m := &operatorNode{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	o.UID = m.UID
	o.Word = []rune(m.Word)
	o.PUID = m.PUID
	o.Column = m.Column
	o.Unknown = m.Unknown
	o.Value = m.Value
	o.Resolved = m.Resolved
	o.Operation = m.Operation
	return nil
}

//IsResolved will return true if the node is resolved
func (o *OperatorNode) IsResolved() bool {
	return o.Resolved
}

//SetResolved will set the resolved state of the node
func (o *OperatorNode) SetResolved(state bool) {
	o.Resolved = state
}
