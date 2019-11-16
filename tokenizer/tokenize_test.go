// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package tokenizer

import (
	"testing"
	"time"
)

/*
 * This file contains the tests and test utilities for tokenizer
 */

func TestTokenize(t *testing.T) {
	loadTestDICT()
	time.Sleep(time.Second)
	toks, err := Tokenize(testUser, []rune("show me the brands of with Swift cars"))
	if err != nil {
		t.Fatal("error while tokenizing the sentence", err)
	}
	if len(toks) == 0 {
		t.Error("Expected to have more than 1 tokens found. Go none")
	}
}
