// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

/*
 * This file contains the defnition of query of octopus
 */

//Query has the interpreted query info
type Query struct {
	//Tables has the map of tables whose data is being accessed by the query
	Tables map[string]TableNode
	//Select has the list of columns to be selected from the data
	Select []ColumnNode `json:"select,omitempty"`
	//GroupBy has the list of columns to be used for group the data
	GroupBy []ColumnNode `json:"group_by,omitempty"`
	//Filters has the list of filters applied in the query
	Filters []OperatorNode `json:"filters,omitempty"`
	//Result has the result of the query
	Result interface{}
}

//SQLQuery stores a sql query to be executed
type SQLQuery struct {
	//Query is the query string with arguments
	Query string
	//Args has the arguments to be passed to the query string
	Args []interface{}
}

//ToSQL converts the the query to a sql query
func (q Query) ToSQL() (*SQLQuery, error) {
	/*
	 * First we will add the select fields
	 * Then we will add the filters
	 * Then we will ad the group by
	 */
	return nil, nil
}
