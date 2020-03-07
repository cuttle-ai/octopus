// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import "encoding/json"

/*
 * This file contains the defnition of knowledgebase type node
 */

//KnowledgeBaseNode is the node storing the knownledgebase in the system.
//It doesn't have any parents. But Knowledgebases are connected by universe
type KnowledgeBaseNode struct {
	//UID is the unique id of the knowledgebase node
	UID string
	//Word is the word with which the table node has to be matched
	Word []rune
	//Name is the name of the node
	Name string
	//Children are the children node of the knowledgebase node. Ie tables
	Children []TableNode
}

type knowledgebaseNode struct {
	UID      string      `json:"uid,omitempty"`
	Word     string      `json:"word,omitempty"`
	Name     string      `json:"name,omitempty"`
	Children []TableNode `json:"children,omitempty"`
	Type     string      `json:"type,omitempty"`
}

//Copy will return a copy of the node
func (k *KnowledgeBaseNode) Copy() Node {
	return &KnowledgeBaseNode{
		UID:      k.UID,
		Word:     k.Word,
		Name:     k.Name,
		Children: k.Children,
	}
}

//ID returns the unique id of the node
func (k *KnowledgeBaseNode) ID() string {
	return k.UID
}

//Type returns Table Type
func (k *KnowledgeBaseNode) Type() Type {
	return KnowledgeBase
}

//TokenWord returns the word property of the node
func (k *KnowledgeBaseNode) TokenWord() []rune {
	return k.Word
}

//PID returns the PUID if the node
func (k *KnowledgeBaseNode) PID() string {
	return ""
}

//Parent returns the PN of the node
func (k *KnowledgeBaseNode) Parent() Node {
	return nil
}

//MarshalJSON encodes the node into a serializable json
func (k *KnowledgeBaseNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&knowledgebaseNode{
		k.UID, string(k.Word), k.Name, k.Children, "KnowledgeBase",
	})
}

//UnmarshalJSON decodes the node from a json
func (k *KnowledgeBaseNode) UnmarshalJSON(data []byte) error {
	m := &knowledgebaseNode{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	k.UID = m.UID
	k.Word = []rune(m.Word)
	k.Name = m.Name
	k.Children = m.Children
	return nil
}

//IsResolved will return true if the node is resolved
func (k *KnowledgeBaseNode) IsResolved() bool {
	return false
}

//SetResolved will set the resolved state of the node
func (k *KnowledgeBaseNode) SetResolved(state bool) {
}
