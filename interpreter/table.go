// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import "encoding/json"

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
	//DefaultDateFieldUID is the uid of the default date field in the table
	DefaultDateFieldUID string
	//DefaultDateField is the default column to be selected as date in the table
	DefaultDateField *ColumnNode
}

type tableNode struct {
	UID      string       `json:"uid,omitempty"`
	Word     string       `json:"word,omitempty"`
	PUID     string       `json:"puid,omitempty"`
	Name     string       `json:"name,omitempty"`
	Children []ColumnNode `json:"children,omitempty"`
	Resolved bool         `json:"resolved,omitempty"`
	Type     string       `json:"type,omitempty"`
}

//Copy will return a copy of the node
func (t *TableNode) Copy() Node {
	return &TableNode{
		UID:      t.UID,
		Word:     t.Word,
		PN:       t.PN,
		PUID:     t.PUID,
		Name:     t.Name,
		Children: t.Children,
		Resolved: t.Resolved,
	}
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

//MarshalJSON encodes the node into a serializable json
func (t *TableNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&tableNode{
		t.UID, string(t.Word), t.PUID, t.Name, t.Children, t.Resolved, "Table",
	})
}

//UnmarshalJSON decodes the node from a json
func (t *TableNode) UnmarshalJSON(data []byte) error {
	m := &tableNode{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	t.UID = m.UID
	t.Word = []rune(m.Word)
	t.PUID = m.PUID
	t.Name = m.Name
	t.Children = m.Children
	t.Resolved = m.Resolved
	return nil
}

//IsResolved will return true if the node is resolved
func (t *TableNode) IsResolved() bool {
	return t.Resolved
}

//SetResolved will set the resolved state of the node
func (t *TableNode) SetResolved(state bool) {
	t.Resolved = state
}
