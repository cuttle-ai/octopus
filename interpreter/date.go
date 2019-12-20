// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import "encoding/json"

import "time"

/*
 * This file contains the defnition of time type node
 */

//TimeNode is the node storing the information about a time.
type TimeNode struct {
	//UID is the unique id of the time node
	UID string
	//Word is the word with which the time node has to be matched
	Word []rune
	//PUID is the UID of time's parent node
	PUID string
	//PN is the parent node of the time. It will be a KnowledgeBase
	PN Node
	//Time is the time represented by the node
	Time time.Time
	//Gran is the granularity of the time
	Gran string
	//Resolved indicates that the node is resolved
	Resolved bool
}

type timeNode struct {
	UID      string    `json:"uid,omitempty"`
	Word     string    `json:"word,omitempty"`
	PUID     string    `json:"puid,omitempty"`
	Resolved bool      `json:"resolved,omitempty"`
	Time     time.Time `json:"time,omitempty"`
	Gran     string    `json:"gran,omitempty"`
	Type     string    `json:"type,omitempty"`
}

//Copy will return a copy of the node
func (u *TimeNode) Copy() Node {
	return &TimeNode{
		UID:      u.UID,
		Word:     u.Word,
		PN:       u.PN,
		PUID:     u.PUID,
		Resolved: u.Resolved,
	}
}

//ID returns the unique id of the node
func (u *TimeNode) ID() string {
	return u.UID
}

//Type returns Unknown Type
func (u *TimeNode) Type() Type {
	return Unknown
}

//TokenWord returns the word property of the node
func (u *TimeNode) TokenWord() []rune {
	return u.Word
}

//PID returns the PUID if the node
func (u *TimeNode) PID() string {
	return u.PUID
}

//Parent returns the PN of the node
func (u *TimeNode) Parent() Node {
	return u.PN
}

//MarshalJSON encodes the node into a serializable json
func (u *TimeNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&timeNode{
		u.UID, string(u.Word), u.PUID, u.Resolved, u.Time, u.Gran, "Date",
	})
}

//UnmarshalJSON decodes the node from a json
func (u *TimeNode) UnmarshalJSON(data []byte) error {
	m := &timeNode{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	u.UID = m.UID
	u.Word = []rune(m.Word)
	u.PUID = m.PUID
	u.Resolved = m.Resolved
	u.Time = m.Time
	u.Gran = m.Gran
	return nil
}

//IsResolved will return true if the node is resolved
func (u *TimeNode) IsResolved() bool {
	return u.Resolved
}

//SetResolved will set the resolved state of the node
func (u *TimeNode) SetResolved(state bool) {
	u.Resolved = state
}
