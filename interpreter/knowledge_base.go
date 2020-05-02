// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import "encoding/json"

/*
 * This file contains the definition of knowledge base type node
 */

//KnowledgeBaseType indicates the type of the knowledgebase
type KnowledgeBaseType uint

const (
	//SystemKB stands for the knowledgebase of the system
	SystemKB KnowledgeBaseType = 1
	//UserKB stands for the knowledgebase of the user
	UserKB KnowledgeBaseType = 2
)

//KnowledgeBaseNode can store info about any node. It doesn't have a parent node
type KnowledgeBaseNode struct {
	//UID is the unique id of the node
	UID string
	//Word is the word with which the knowledge base node has to be matched
	Word []rune
	//Name is the name of the node
	Name string
	//Children are the children node of the node
	Children []Node
	//Resolved indicates that the node is resolved
	Resolved bool
	//Description of the node
	Description string
	//KBType indicates the type of the knowledgebase
	KBType KnowledgeBaseType
}

type knowledgebaseNode struct {
	UID         string            `json:"uid,omitempty"`
	Word        string            `json:"word,omitempty"`
	Name        string            `json:"name,omitempty"`
	Children    []Node            `json:"children,omitempty"`
	Resolved    bool              `json:"resolved,omitempty"`
	Type        string            `json:"type,omitempty"`
	Description string            `json:"description"`
	KBType      KnowledgeBaseType `json:"kb_type"`
}

//Copy will return a copy of the node
func (k *KnowledgeBaseNode) Copy() Node {
	return &KnowledgeBaseNode{
		UID:         k.UID,
		Word:        k.Word,
		Name:        k.Name,
		Children:    k.Children,
		Resolved:    k.Resolved,
		Description: k.Description,
		KBType:      k.KBType,
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
		k.UID, string(k.Word), k.Name, k.Children, k.Resolved, "KnowledgeBase", k.Description, k.KBType,
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
	k.Resolved = m.Resolved
	k.Description = m.Description
	k.KBType = m.KBType
	return nil
}

//IsResolved will return true if the node is resolved
func (k *KnowledgeBaseNode) IsResolved() bool {
	return k.Resolved
}

//SetResolved will set the resolved state of the node
func (k *KnowledgeBaseNode) SetResolved(state bool) {
	k.Resolved = state
}
