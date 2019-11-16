// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package tokenizer

/*
 * This file contains the defnition of tokens in the system
 */

//Token is the parsed word and possible nodes a word associated to
type Token struct {
	//Pos is the position of the token in the root sentence
	Pos int
	//Word is the match word corresponding to the token
	Word []rune
	//Node has the list of nodes applicable to a token
	Nodes []Node
}
