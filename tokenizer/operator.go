// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package tokenizer

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

//Resolve will try resolve the node with the given tokens
func (o *OperatorNode) Resolve(tokens []FastToken, pos int) bool {
	/*
	 * 1. If there exists a value node after the operator.
	 * 2. If there exists a column node before the operator node and
	 *    there exists a unknow node after the operator node.
	 */
	//if already resolved, we return
	if o.Resolved {
		return true
	}

	//check if the position is 0 or end of the tokens. then we can't go forward
	if pos == 0 || len(tokens)-1 == pos {
		o.Resolved = false
		return false
	}

	//checking if there exists a value node after the operator node
	var valueNodeC = len(tokens[pos+1].Values) > 0
	if valueNodeC {
		o.Resolved = true
		tokens[pos+1].Values[0].Resolved = true
		return true
	}

	//get the token before the operator and try to find if there exists a column node in that
	var columnC = len(tokens[pos-1].Columns) == 0 || tokens[pos-1].Columns[0].Resolved
	var unknownC = len(tokens[pos+1].Unknowns) == 0 || tokens[pos+1].Unknowns[0].Resolved
	if columnC || unknownC {
		o.Resolved = false
		return false
	}

	o.Resolved = true
	tokens[pos-1].Columns[0].Resolved = true
	tokens[pos+1].Unknowns[0].Resolved = true
	return false
}
