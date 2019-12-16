// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

//Package interpreter has the utilities and defnition of the interpreter implementation of octopus
package interpreter

/*
 * This file contains the defnition of rule templates to be used for the interpreter
 */

//Rule represents a rule with template and resolver to resolve the parsed tokens
type Rule struct {
	//Name of the rule for debugging purposes
	Name string
	//Description of the rule.
	Description string
	//Disabled indicates wether the riules is emnabled
	Disbled bool
	//Template is the template of the of the rule
	Template []Type
	//Resolve function will try to run the resolution for the rule.
	//Query argument is the query to which the resolved tokens has to be attached
	//The int argument gives the index of the fasttokoen we are referring to
	//Resolve function should not mutate the state of the rule
	Resolve func(Query, []FastToken, int) (Query, error)
	//Matches are the indices of the tokens in the list of tokens to which the rule template has found match
	Matches []int
	//Pattern is the kmp buildup of the pattern
	Pattern *KMP
}

//RuleGroup stores the list of rules to be executed together with priority
type RuleGroup struct {
	//Rules in the group
	Rules []Rule
	//Tag identifier for the group
	Tag string
}

var rules = []*RuleGroup{}

//AddRule will add a given rule with the given priority and priority in the group
//It will also build the kmp pattern for the rule
func AddRule(rule Rule, priority, groupPriority int, tag string) {
	/*
	 * If the group rules doesn't exist add a new one
	 * If a different group exist replace it
	 * If the group doesn't have space to add a new rule. Increase the size of the rules
	 * Initialize the rule pattern
	 * Add the rule
	 */
	//Adding new group if required
	if len(rules) > priority && rules[priority] == nil {
		rules[priority] = &RuleGroup{Rules: []Rule{}, Tag: tag}
	} else if len(rules) <= priority {
		rules = append(rules, make([]*RuleGroup, priority-len(rules)+1)...)
		rules[priority] = &RuleGroup{Rules: []Rule{}, Tag: tag}
	} else if len(rules) > priority && rules[priority] != nil && rules[priority].Tag != tag {
		//replacing an existing group
		rules[priority] = &RuleGroup{Rules: []Rule{}, Tag: tag}
	}
	groupRules := rules[priority].Rules

	//add space for rules to group if necessary
	if len(groupRules) <= groupPriority {
		groupRules = append(groupRules, make([]Rule, groupPriority-len(groupRules)+1)...)
	}

	//initialize the rule pattern
	rule.Pattern = NewKMP(rule.Template)

	//adding thr rule
	groupRules[groupPriority] = rule
	rules[priority].Rules = groupRules
}

//MatchRules will try to match the tokens passed with rules existing in the interpreter
//If no match is found, will return nil
func MatchRules(tokens []FastToken) []Rule {
	/*
	 * We will first build a pattern for the given tokens
	 * Then will match with the existing rule patterns using kmp
	 */
	result := []Rule{}
	//building the pattern
	tokPattern := BuildPattern(tokens)

	//trying to find matches for the rules with the token pattern
	for _, gr := range rules {
		for _, r := range gr.Rules {
			//if the disabled skip the rule
			if r.Disbled {
				continue
			}
			pos := r.Pattern.Matches(tokPattern)
			if len(pos) > 0 {
				newRule := Rule{Name: r.Name, Matches: pos, Resolve: r.Resolve}
				copy(newRule.Template, r.Template)
				result = append(result, newRule)
			}
		}
	}
	return result
}

//SetRuleDisableState will set the disable state of a rule
func SetRuleDisableState(pos, groupPos int, state bool) {
	for i := 0; i < len(rules); i++ {
		for j := 0; j < len(rules[i].Rules); j++ {
			if i == pos && groupPos == j {
				rules[i].Rules[j].Disbled = state
			}
		}
	}
}

//GetRules return the rules used in the interpreter
func GetRules() []*RuleGroup {
	return rules
}

//BuildPattern will build pattern for the given tokens
func BuildPattern(tokens []FastToken) []Type {
	/*
	 * We will iterate through the tokens and retrive the types
	 */
	result := []Type{}
	for _, v := range tokens {
		//We will only take the first node in the token in the following order
		// Operator
		// Value
		// Column
		// Table
		// Unknown
		if len(v.Operators) > 0 {
			result = append(result, Operator)
		} else if len(v.Values) > 0 {
			result = append(result, Value)
		} else if len(v.Columns) > 0 {
			result = append(result, Column)
		} else if len(v.Tables) > 0 {
			result = append(result, Table)
		} else if len(v.Unknowns) > 0 {
			result = append(result, Unknown)
		}
	}
	return result
}
