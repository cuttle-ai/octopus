// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import "encoding/json"

/*
 * This file contains the defnition of column type node
 */

const (
	//AggregationFnCount for count aggregation funtion
	AggregationFnCount = "COUNT"
	//AggregationFnSum for sum aggregation funtion
	AggregationFnSum = "SUM"
	//AggregationFnAvg for average aggregation funtion
	AggregationFnAvg = "AVG"
)

const (
	//DataTypeInt denotes integer data type
	DataTypeInt = "INT"
	//DataTypeFloat denotes float data type
	DataTypeFloat = "FLOAT"
	//DataTypeString denotes string data type
	DataTypeString = "STRING"
	//DataTypeDate denotes date data type
	DataTypeDate = "DATE"
)

//ColumnNode is the node storing the information about a column.
//It can be the child of a Table Node and can have value as children
type ColumnNode struct {
	//UID is the unique id of the column node
	UID string
	//Word is the word with which the column node has to be matched
	Word []rune
	//PN is the parent node of the column. It will be a table node
	PN *TableNode
	//PUID is the UID of column's parent node
	PUID string
	//Name is the name of the node
	Name string
	//Children are the children node of the column node
	Children []ValueNode
	//Resolved indicates that the node is resolved
	Resolved bool
	//Dimension indicates that the column can be used as dimension
	Dimension bool
	//Measure indicates that the column can be used as measure
	Measure bool
	//AggregationFn is preferred aggregation with the column node if used as a measure across a dimension
	AggregationFn string
	//DataType of the column
	DataType string
}

type columnNode struct {
	UID           string      `json:"uid,omitempty"`
	Word          string      `json:"word,omitempty"`
	PUID          string      `json:"puid,omitempty"`
	Name          string      `json:"name,omitempty"`
	Children      []ValueNode `json:"children,omitempty"`
	Resolved      bool        `json:"resolved,omitempty"`
	Type          string      `json:"type,omitempty"`
	Dimension     bool        `json:"dimension,omitempty"`
	Measure       bool        `json:"measure,omitempty"`
	AggregationFn string      `json:"aggregation_fn,omitempty"`
	DataType      string      `json:"data_type,omitempty"`
}

//Copy will return a copy of the node
func (c *ColumnNode) Copy() Node {
	return &ColumnNode{
		UID:           c.UID,
		Word:          c.Word,
		PN:            c.PN,
		PUID:          c.PUID,
		Name:          c.Name,
		Children:      c.Children,
		Resolved:      c.Resolved,
		Dimension:     c.Dimension,
		Measure:       c.Measure,
		AggregationFn: c.AggregationFn,
		DataType:      c.DataType,
	}
}

//ID returns the unique id of the node
func (c *ColumnNode) ID() string {
	return c.UID
}

//Type returns Column Type
func (c *ColumnNode) Type() Type {
	return Column
}

//TokenWord returns the word property of the node
func (c *ColumnNode) TokenWord() []rune {
	return c.Word
}

//PID returns the PUID if the node
func (c *ColumnNode) PID() string {
	return c.PUID
}

//Parent returns the PN of the node
func (c *ColumnNode) Parent() Node {
	return c.PN
}

//MarshalJSON encodes the node into a serializable json
func (c *ColumnNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(&columnNode{
		c.UID, string(c.Word), c.PUID, c.Name, c.Children, c.Resolved, "Column", c.Dimension, c.Measure, c.AggregationFn, c.DataType,
	})
}

//UnmarshalJSON decodes the node from a json
func (c *ColumnNode) UnmarshalJSON(data []byte) error {
	m := &columnNode{}
	err := json.Unmarshal(data, m)
	if err != nil {
		return err
	}
	c.UID = m.UID
	c.Word = []rune(m.Word)
	c.PUID = m.PUID
	c.Name = m.Name
	c.Children = m.Children
	c.Resolved = m.Resolved
	c.Dimension = m.Dimension
	c.Measure = m.Measure
	c.AggregationFn = m.AggregationFn
	c.DataType = m.DataType
	return nil
}

//IsResolved will return true if the node is resolved
func (c *ColumnNode) IsResolved() bool {
	return c.Resolved
}

//SetResolved will set the resolved state of the node
func (c *ColumnNode) SetResolved(state bool) {
	c.Resolved = state
}
