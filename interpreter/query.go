// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

/*
 * This file contains the defnition of query of octopus
 */

//Query has the interpreted query info
type Query struct {
	//Select has the list of columns to be selected from the data
	Select []ColumnNode
}
