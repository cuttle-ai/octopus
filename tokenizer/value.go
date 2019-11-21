// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package tokenizer

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

//Encode encodes the node into a serializable form
func (v *ValueNode) Encode() []byte {
	return nil
}

//Decode decodes the node from the serialized data
func (v *ValueNode) Decode(enc []byte) bool {
	return false
}

//Resolve will try resolve the node with the given tokens
func (v *ValueNode) Resolve(tokens []FastToken, pos int) bool {
	return false
}
