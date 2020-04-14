// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cuttle-ai/octopus/interpreter"
)

/*
 * This file contains the tests and test utilities for interpreter
 */

func TestInterpret(t *testing.T) {
	loadTestDICT()
	time.Sleep(time.Second)
	toks, err := interpreter.Tokenize(testUser, []rune("show me the brands of with Swift cars"))
	if err != nil {
		t.Fatal("error while tokenizing the sentence", err)
	}
	if len(toks) == 0 {
		t.Error("Expected to have more than 1 tokens found. Go none")
	}

	qu, err := interpreter.Interpret(toks)

	fmt.Println(qu)
}
