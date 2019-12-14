// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import (
	"testing"

	"github.com/cuttle-ai/octopus/testutils"
)

/*
 * This file contains the tests for kmp
 */

type newKmpTest struct {
	testutils.Test
	Input    []Type
	Expected []int
}

var newKmpTestcases = []newKmpTest{
	{
		Test:     testutils.Test{Name: "Simple test", Description: "Basic working of new kmp"},
		Input:    []Type{Column, Column},
		Expected: []int{0, 1},
	},
}

func TestNewKMP(t *testing.T) {
	for _, v := range newKmpTestcases {
		t.Run(v.Name, func(t *testing.T) {
			p := NewKMP(v.Input)
			if len(p.Pos) != len(v.Expected) {
				t.Error("Expected output length differs. Expected", len(v.Expected), "got", len(p.Pos))
				t.Error("Expected", v.Expected, "got", p.Pos)
				return
			}
			for i := range p.Pos {
				if p.Pos[i] != v.Expected[i] {
					t.Error("Expcted result varies at the index", i)
					t.Error("Expected", v.Expected, "got", p.Pos)
					return
				}
			}
		})
	}
}

type kmpMatchesTest struct {
	testutils.Test
	Pattern  []Type
	Input    []Type
	Expected []int
}

var kmpMatchesTestcases = []kmpMatchesTest{
	{
		Test:     testutils.Test{Name: "Simple test", Description: "Basic working of the kmp"},
		Pattern:  []Type{Column, Column},
		Input:    []Type{Column, Column, Value},
		Expected: []int{0},
	},
	{
		Test:     testutils.Test{Name: "Multiple case test", Description: "multiple match"},
		Pattern:  []Type{Column, Column},
		Input:    []Type{Column, Column, Value, Column, Column, Column},
		Expected: []int{0, 3, 4},
	},
}

func TestKMPMatches(t *testing.T) {
	for _, v := range kmpMatchesTestcases {
		t.Run(v.Name, func(t *testing.T) {
			p := NewKMP(v.Pattern)
			res := p.Matches(v.Input)
			if len(res) != len(v.Expected) {
				t.Error("Expected output length differs. Expected", len(v.Expected), "got", len(res))
				t.Error("Expected", v.Expected, "got", res)
				return
			}
			for i := range res {
				if res[i] != v.Expected[i] {
					t.Error("Expcted result varies at the index", i)
					t.Error("Expected", v.Expected, "got", res)
					return
				}
			}
		})
	}
}
