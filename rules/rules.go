// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

//Package rules has the list of defaults rules for the octopus interpreter
package rules

import "github.com/cuttle-ai/octopus/interpreter"

/*
 * This file contains the list of default rules to be loaded to be added
 */

//DefaultRulesTag is the tag used for the deafult rules group
const DefaultRulesTag = "DEFAULT_RULES"

//LoadDefaultRules will load the default rules to the interpreter rule engine
func LoadDefaultRules() {
	interpreter.AddRule(UnknownFilter, 0, 0, DefaultRulesTag)
	interpreter.AddRule(GroupByColumn, 0, 1, DefaultRulesTag)
	interpreter.AddRule(SelectColumn, 0, 2, DefaultRulesTag)
}
