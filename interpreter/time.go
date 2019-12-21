// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import (
	"encoding/json"
	"time"
)

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
func (t *TimeNode) Copy() Node {
	return &TimeNode{
		UID:      t.UID,
		Word:     t.Word,
		PN:       t.PN,
		PUID:     t.PUID,
		Resolved: t.Resolved,
		Time:     t.Time,
		Gran:     t.Gran,
	}
}

//ID returns the unique id of the node
func (t *TimeNode) ID() string {
	return t.UID
}

//Type returns Unknown Type
func (t *TimeNode) Type() Type {
	return Unknown
}

//TokenWord returns the word property of the node
func (t *TimeNode) TokenWord() []rune {
	return t.Word
}

//PID returns the PUID if the node
func (t *TimeNode) PID() string {
	return t.PUID
}

//Parent returns the PN of the node
func (t *TimeNode) Parent() Node {
	return t.PN
}

//MarshalJSON encodes the node into a serializable json
func (t *TimeNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&timeNode{
		t.UID, string(t.Word), t.PUID, t.Resolved, t.Time, t.Gran, "Date",
	})
}

//UnmarshalJSON decodes the node from a json
func (t *TimeNode) UnmarshalJSON(data []byte) error {
	m := &timeNode{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	t.UID = m.UID
	t.Word = []rune(m.Word)
	t.PUID = m.PUID
	t.Resolved = m.Resolved
	t.Time = m.Time
	t.Gran = m.Gran
	return nil
}

//IsResolved will return true if the node is resolved
func (t *TimeNode) IsResolved() bool {
	return t.Resolved
}

//SetResolved will set the resolved state of the node
func (t *TimeNode) SetResolved(state bool) {
	t.Resolved = state
}
