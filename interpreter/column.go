// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

/*
 * This file contains the defnition of column type node
 */

//ColumnNode is the node storing the information about a column.
//It can be the child of a Table Node and can have value as children
type ColumnNode struct {
	//UID is the unique id of the column node
	UID string
	//Word is the word with which the column node has to be matched
	Word []rune
	//PN is the parent node of the column. It will be a table node
	PN *TableNode
	//PUID is the UID of column's parent node
	PUID string
	//Name is the name of the node
	Name string
	//Children are the children node of the column node
	Children []ValueNode
	//Resolved indicates that the node is resolved
	Resolved bool
}

//Copy will return a copy of the node
func (c *ColumnNode) Copy() Node {
	return &ColumnNode{
		UID:      c.UID,
		Word:     c.Word,
		PN:       c.PN,
		PUID:     c.PUID,
		Name:     c.Name,
		Children: c.Children,
		Resolved: c.Resolved,
	}
}

//ID returns the unique id of the node
func (c *ColumnNode) ID() string {
	return c.UID
}

//Type returns Column Type
func (c *ColumnNode) Type() Type {
	return Column
}

//TokenWord returns the word property of the node
func (c *ColumnNode) TokenWord() []rune {
	return c.Word
}

//PID returns the PUID if the node
func (c *ColumnNode) PID() string {
	return c.PUID
}

//Parent returns the PN of the node
func (c *ColumnNode) Parent() Node {
	return c.PN
}

//Encode encodes the node into a serializable form
func (c *ColumnNode) Encode() []byte {
	return nil
}

//Decode decodes the node from the serialized data
func (c *ColumnNode) Decode(enc []byte) bool {
	return false
}

//IsResolved will return true if the node is resolved
func (c *ColumnNode) IsResolved() bool {
	return c.Resolved
}

//SetResolved will set the resolved state of the node
func (c *ColumnNode) SetResolved(state bool) {
	c.Resolved = state
}
