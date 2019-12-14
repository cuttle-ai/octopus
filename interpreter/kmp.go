// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

/*
 * This file contains the implmentation of variant of kmp pattern matching for mathcing
 * the rules with the tokens
 */

//KMP holds the infomartaion on a pattern to find match with a given target array
type KMP struct {
	//Pos holds the position information about the pattern
	Pos []int
	//KMP Pattern
	Pattern []Type
}

//NewKMP initializes a kmp pattern and does the precomputation and returns the pointer to it
func NewKMP(pattern []Type) *KMP {
	result := &KMP{Pos: make([]int, len(pattern)), Pattern: pattern}
	j, i := 0, 1
	result.Pos[0] = 0

	for i < len(pattern) {
		if pattern[i] == pattern[j] {
			j++
			result.Pos[i] = j
			i++
			continue
		}
		if j != 0 {
			j = result.Pos[j-1]
			continue
		}
		i++
	}
	return result
}

//Matches finds if the given target has any substring match with the kmp pattern
func (k KMP) Matches(target []Type) []int {
	i, j := 0, 0
	n, m := len(target), len(k.Pattern)
	var ret []int

	//got zero target or want, just return empty result
	if m > n || m == 0 || n == 0 {
		return ret
	}

	for i < n {
		if k.Pattern[j] == target[i] {
			i++
			j++
		}
		if j == m {
			ret = append(ret, i-j)
			j = k.Pos[j-1]
		} else if i < n && k.Pattern[j] != target[i] {
			if j != 0 {
				j = k.Pos[j-1]
			} else {
				i++
			}
		}
	}

	return ret
}
