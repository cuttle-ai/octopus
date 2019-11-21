// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package tokenizer

/*
 * This file contains the defnition of table type node
 */

//TableNode is the node storing the information about a table.
//It can be the child of a KnowledgeBase and can have Column as children
type TableNode struct {
	//UID is the unique id of the table node
	UID string
	//Word is the word with which the table node has to be matched
	Word []rune
	//PUID is the UID of table's parent node
	PUID string
	//PN is the parent node of the table. It will be a KnowledgeBase
	PN Node
	//Name is the name of the node
	Name string
	//Children are the children node of the table node
	Children []ColumnNode
	//Resolved indicates that the node is resolved
	Resolved bool
}

//ID returns the unique id of the node
func (t *TableNode) ID() string {
	return t.UID
}

//Type returns Table Type
func (t *TableNode) Type() Type {
	return Table
}

//TokenWord returns the word property of the node
func (t *TableNode) TokenWord() []rune {
	return t.Word
}

//PID returns the PUID if the node
func (t *TableNode) PID() string {
	return t.PUID
}

//Parent returns the PN of the node
func (t *TableNode) Parent() Node {
	return t.PN
}

//Encode encodes the node into a serializable form
func (t *TableNode) Encode() []byte {
	return nil
}

//Decode decodes the node from the serialized data
func (t *TableNode) Decode(enc []byte) bool {
	return false
}

//Resolve will try resolve the node with the given tokens
func (t *TableNode) Resolve(tokens []FastToken, pos int) bool {
	return false
}
