// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import (
	"strconv"
)

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

//FastToken is used to store the token with nodes converted into their concrete type so
//processing become becomes easy
type FastToken struct {
	//Pos is the position of the token in the root sentence
	Pos int
	//Word is the match word corresponding to the token
	Word []rune
	//Tables is the list of table nodes in the token
	Tables []TableNode
	//Columns is the list of column nodes in the token
	Columns []ColumnNode
	//Values is the list of value nodes in the token
	Values []ValueNode
	//Operators is the list of operators nodes in the token
	Operators []OperatorNode
	//Unknowns is the list of unknows nodes in the token
	Unknowns []UnknownNode
	//Times is the list of time nodes in the token
	Times []TimeNode
}

//FastToken returns the converted fast token of the token
func (t Token) FastToken() FastToken {
	result := FastToken{Pos: t.Pos, Word: t.Word}

	for _, n := range t.Nodes {
		switch n.Type() {
		case Table:
			tab, ok := n.(*TableNode)
			if ok {
				if result.Tables == nil {
					result.Tables = []TableNode{}
				}
				result.Tables = append(result.Tables, *tab)
			}
		case Column:
			col, ok := n.(*ColumnNode)
			if ok {
				if result.Columns == nil {
					result.Columns = []ColumnNode{}
				}
				result.Columns = append(result.Columns, *col)
			}
		case Value:
			val, ok := n.(*ValueNode)
			if ok {
				if result.Values == nil {
					result.Values = []ValueNode{}
				}
				result.Values = append(result.Values, *val)
			}
		case Operator:
			op, ok := n.(*OperatorNode)
			if ok {
				if result.Operators == nil {
					result.Operators = []OperatorNode{}
				}
				result.Operators = append(result.Operators, *op)
			}
		case Unknown:
			un, ok := n.(*UnknownNode)
			if ok {
				if result.Unknowns == nil {
					result.Unknowns = []UnknownNode{}
				}
				result.Unknowns = append(result.Unknowns, *un)
			}
		case Time:
			tn, ok := n.(*TimeNode)
			if ok {
				if result.Times == nil {
					result.Times = []TimeNode{}
				}
				result.Times = append(result.Times, *tn)
			}
		}
	}

	return result
}

//Copy makes a deep copy of the token
func (t Token) Copy() Token {
	res := Token{Pos: t.Pos, Word: t.Word, Nodes: []Node{}}
	for _, v := range t.Nodes {
		res.Nodes = append(res.Nodes, v.Copy())
	}
	return res
}

//String is the stringer implementation of the fast token
func (f FastToken) String() string {
	return string(f.Word) + "-" + strconv.Itoa(f.Pos)
}
