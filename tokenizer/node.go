// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package tokenizer

/*
 * This file contains the defntion of node interface
 */

//Type is type of node
type Type uint

const (
	//KnowledgeBase is the collection of diffent tables etcs
	KnowledgeBase Type = 1
	//Table is the table to which referring data belongs to
	Table Type = 2
	//Column is the column that is being referred in a table
	Column Type = 3
	//Value is the value that is present in the table's specific column
	Value Type = 4
	//Operation is the operation to be applied when doing a filter
	Operation Type = 5
	//GroupBy is the based on which the values of columns should be grouped
	GroupBy Type = 6
	//AggregationFn is the aggregation function to be used for a column in a query
	AggregationFn Type = 7
	//Unknown is a node whose purpose has still not been resolved
	Unknown Type = 8
	//Ignore is node that has to be ignored without going for further processing
	Ignore Type = 9
	//Context node if found indicates that there is an context to the query and certain values can be inferrerd from that
	Context Type = 10
)

//Node is the interface to be implemented for considering it as a basic building block in octopus
type Node interface {
	//ID is the unique identifier of the node
	ID() string
	//Type is the type of the node
	Type() Type
	//TokenWord returns the word to be matched with the token
	TokenWord() []rune
	//PID returns the id of the parent associated with the node
	PID() string
	//Parent returns the parent node
	Parent() Node
	//Encode will encode a node to binary string which can be stored
	Encode() []byte
	//Decode will decode a binary string to a node. Returns true if sucessfully decoded
	Decode([]byte) bool
	//Resolve returns true if the node is resolved
	Resolve([]Token) bool
}
