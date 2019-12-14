// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import "strings"

/*
 * This file contains the test and test utilities for testing dictionary
 */

var testColumn = &ColumnNode{UID: "cars", Word: []rune("cars")}

var testColumn1 = &ColumnNode{UID: "brands", Word: []rune("brand")}

var testValue = &ValueNode{UID: "swift", Word: []rune("Swift"), PN: testColumn, PUID: "cars"}

var testTokens = map[string]Token{
	"cars": Token{
		Word:  []rune("cars"),
		Nodes: []Node{testColumn},
	},
	"Swift": Token{
		Word:  []rune("Swift"),
		Nodes: []Node{testValue},
	},
	"brands": Token{
		Word:  []rune("brands"),
		Nodes: []Node{testColumn1},
	},
}

var testDICT = DICT{Map: testTokens}

var testUser = "testuser"

func loadTestDICT() {
	tTok := map[string]Token{}
	for k, v := range testDICT.Map {
		tTok[strings.ToLower(k)] = v
	}
	testDICT.Map = tTok
	req := DICTRequest{
		ID:   testUser,
		Type: DICTAdd,
		DICT: testDICT,
	}
	SendDICTToChannel(DICTInputChannel, req)
}
