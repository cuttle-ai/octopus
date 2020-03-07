// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

//Package interpreter has the utilities and defnition of the interpreter implementation of octopus
package interpreter

import "fmt"

/*
 * This file contains the defnition of interpreter of octopus
 */

//Interpret the given list of token to meaningful query
func Interpret(toks []FastToken) (*Query, error) {
	/*
	 * Will run the tokens through the rule match to get the rules to be run
	 * Then will run the rules on the tokens
	 */
	//running through rules for finding matches
	rules := MatchRules(toks)

	//iterating through the rules to resolve them
	q := &Query{Tables: map[string]TableNode{}}
	for _, rule := range rules {
		for _, i := range rule.Matches {
			qu, err := rule.Resolve(*q, toks, i)
			if err != nil {
				fmt.Println("Couldn't apply the rule", rule.Name, "to the query at index", i, err)
				continue
			}
			*q = qu
		}
	}
	return q, nil
}
