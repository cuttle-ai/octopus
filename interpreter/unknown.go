// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import "encoding/json"

/*
 * This file contains the defnition of unknown type node
 */

//UnknownNode is the node storing the information about a unknown.
//Unkown tokens are tokens that are not identified by the system but can have potential information
//for resolving the query
type UnknownNode struct {
	//UID is the unique id of the unkonown node
	UID string
	//Word is the word with which the unknown node has to be matched
	Word []rune
	//PUID is the UID of unknown's parent node
	PUID string
	//PN is the parent node of the unknown. It will be a KnowledgeBase
	PN Node
	//Resolved indicates that the node is resolved
	Resolved bool
}

type unknownNode struct {
	UID      string `json:"uid,omitempty"`
	Word     string `json:"word,omitempty"`
	PUID     string `json:"puid,omitempty"`
	Resolved bool   `json:"resolved,omitempty"`
	Type     string `json:"type,omitempty"`
}

//Copy will return a copy of the node
func (u *UnknownNode) Copy() Node {
	return &UnknownNode{
		UID:      u.UID,
		Word:     u.Word,
		PN:       u.PN,
		PUID:     u.PUID,
		Resolved: u.Resolved,
	}
}

//ID returns the unique id of the node
func (u *UnknownNode) ID() string {
	return u.UID
}

//Type returns Unknown Type
func (u *UnknownNode) Type() Type {
	return Unknown
}

//TokenWord returns the word property of the node
func (u *UnknownNode) TokenWord() []rune {
	return u.Word
}

//PID returns the PUID if the node
func (u *UnknownNode) PID() string {
	return u.PUID
}

//Parent returns the PN of the node
func (u *UnknownNode) Parent() Node {
	return u.PN
}

//MarshalJSON encodes the node into a serializable json
func (u *UnknownNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&unknownNode{
		u.UID, string(u.Word), u.PUID, u.Resolved, "Unknown",
	})
}

//UnmarshalJSON decodes the node from a json
func (u *UnknownNode) UnmarshalJSON(data []byte) error {
	m := &unknownNode{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	u.UID = m.UID
	u.Word = []rune(m.Word)
	u.PUID = m.PUID
	u.Resolved = m.Resolved
	return nil
}

//IsResolved will return true if the node is resolved
func (u *UnknownNode) IsResolved() bool {
	return u.Resolved
}

//SetResolved will set the resolved state of the node
func (u *UnknownNode) SetResolved(state bool) {
	u.Resolved = state
}
