// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

/*
 * This file contains the defnition of operator type node
 */

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
	Column ColumnNode
	//Unknown is the value to be applied to the column node with the operator
	Unknown UnknownNode
	//Value is the value to be applied to the column node with the operator
	Value ValueNode
}

//Copy will return a copy of the node
func (o *OperatorNode) Copy() Node {
	return &OperatorNode{
		UID:      o.UID,
		Word:     o.Word,
		PN:       o.PN,
		PUID:     o.PUID,
		Resolved: o.Resolved,
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

//Encode encodes the node into a serializable form
func (o *OperatorNode) Encode() []byte {
	return nil
}

//Decode decodes the node from the serialized data
func (o *OperatorNode) Decode(enc []byte) bool {
	return false
}

//IsResolved will return true if the node is resolved
func (o *OperatorNode) IsResolved() bool {
	return o.Resolved
}


//SetResolved will set the resolved state of the node
func (o *OperatorNode) SetResolved(state bool) {
	o.Resolved = state
}