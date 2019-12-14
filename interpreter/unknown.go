// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

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
	//PUID is the UID of operator's parent node
	PUID string
	//PN is the parent node of the operator. It will be a KnowledgeBase
	PN Node
	//Resolved indicates that the node is resolved
	Resolved bool
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

//Encode encodes the node into a serializable form
func (u *UnknownNode) Encode() []byte {
	return nil
}

//Decode decodes the node from the serialized data
func (u *UnknownNode) Decode(enc []byte) bool {
	return false
}

//IsResolved will return true if the node is resolved
func (u *UnknownNode) IsResolved() bool {
	return u.Resolved
}

//SetResolved will set the resolved state of the node
func (u *UnknownNode) SetResolved(state bool) {
	u.Resolved = state
}
