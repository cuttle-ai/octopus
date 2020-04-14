// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter_test

import (
	"strings"

	"github.com/cuttle-ai/octopus/interpreter"
)

/*
 * This file contains the test and test utilities for testing dictionary
 */

var testColumn = &interpreter.ColumnNode{UID: "cars", Word: []rune("cars")}

var testColumn1 = &interpreter.ColumnNode{UID: "brands", Word: []rune("brand")}

var testValue = &interpreter.ValueNode{UID: "swift", Word: []rune("Swift"), PN: testColumn, PUID: "cars"}

var testTokens = map[string]interpreter.Token{
	"cars": interpreter.Token{
		Word:  []rune("cars"),
		Nodes: []interpreter.Node{testColumn},
	},
	"Swift": interpreter.Token{
		Word:  []rune("Swift"),
		Nodes: []interpreter.Node{testValue},
	},
	"brands": interpreter.Token{
		Word:  []rune("brands"),
		Nodes: []interpreter.Node{testColumn1},
	},
}

var testDICT = interpreter.DICT{Map: testTokens}

var testUser = "testuser"

func loadTestDICT() {
	tTok := map[string]interpreter.Token{}
	for k, v := range testDICT.Map {
		tTok[strings.ToLower(k)] = v
	}
	testDICT.Map = tTok
	req := interpreter.DICTRequest{
		ID:   testUser,
		Type: interpreter.DICTAdd,
		DICT: testDICT,
	}
	interpreter.SendDICTToChannel(interpreter.DICTInputChannel, req)
}
