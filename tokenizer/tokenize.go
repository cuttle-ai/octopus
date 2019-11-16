// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

//Package tokenizer has the functions and structs required for finding the token matches in a sentence
package tokenizer

import "fmt"

import goahocorasick "github.com/anknown/ahocorasick"

import "errors"

import "strings"

/*
 * This file contains the utilities for tokenizing a sentence.
 */

//Tokenize will tokenize a given sentence according to the tokenizer of the given id
func Tokenize(id string, sentence []rune) ([]Token, error) {
	/*
	 * We will make a request to the tokenizer to get the sentence tokenized
	 */
	req := Request{
		ID:       id,
		Type:     TokenizerGet,
		Sentence: sentence,
		Out:      make(chan Request),
	}
	go SendTokenizerToChannel(TokenizerInputChannel, req)
	res := <-req.Out
	if !res.Valid {
		return nil, errors.New("couldn't tokenize the given sentence for the id " + id)
	}
	return res.Matches, nil
}

//RequestType is the type of the request for the tokenizer
type RequestType uint

const (
	//TokenizerAdd adds a tokenizer for the given id
	TokenizerAdd RequestType = 1
	//TokenizerGet returns the tokenizer of a given id
	TokenizerGet RequestType = 2
	//TokenizerRemove the tokenizer from the cache
	TokenizerRemove RequestType = 3
)

//Tokenizer has the tokens map and machine for storing the state of the tokens trie
type Tokenizer struct {
	//Machine has the machine storing the trie
	Machine *goahocorasick.Machine
	//map has the tokens mapped to their word
	Map map[string]Token
}

//Request can be used to make a request to tokenizer cache
type Request struct {
	//ID to which the tokenizer belong to
	ID string
	//Type is the type of the tokenizer request. It can have Add, Get, Remove
	Type RequestType
	//Tokenizer is the tokenizer under watch
	Tokenizer Tokenizer
	//Sentence is the sentence to be tokenized
	Sentence []rune
	//Valid indicates whethe the result is valid or not
	Valid bool
	//matches returns the matched tokens
	Matches []Token
	//Out channel for sending response to the requester
	Out chan Request
}

//TokenizerInputChannel is the input channel to communicate with the cache
var TokenizerInputChannel chan Request

func init() {
	TokenizerInputChannel = make(chan Request)
	go Cache(TokenizerInputChannel)
}

//SendTokenizerToChannel sends a dict request to the channel. This function is to be used with go routines so that
//tokenizer isn't blocked by the requests
func SendTokenizerToChannel(ch chan Request, req Request) {
	ch <- req
}

//Cache is the cache for providing the tokenizer to the platform on demand
func Cache(in chan Request) {
	/*
	 * We will go into an infinte loop
	 * Will wait for the requests to come through the channel
	 * Based on the type of the request we will remove them from memory
	 */
	dict := make(map[string]Tokenizer)
	for {
		req := <-in
		switch req.Type {
		case TokenizerAdd:
			m := new(goahocorasick.Machine)
			d := [][]rune{}
			for word := range req.Tokenizer.Map {
				d = append(d, []rune(strings.ToLower(word)))
			}
			if err := m.Build(d); err != nil {
				fmt.Println(err)
				break
			}
			dict[req.ID] = Tokenizer{Map: req.Tokenizer.Map, Machine: m}
			break
		case TokenizerGet:
			t, mOk := dict[req.ID]
			if !mOk {
				req.Valid = false
				go SendTokenizerToChannel(req.Out, req)
				break
			}
			terms := t.Machine.MultiPatternSearch([]rune(strings.ToLower(string(req.Sentence))), false)
			result := []Token{}
			for _, term := range terms {
				tok, ok := t.Map[string(term.Word)]
				if ok {
					result = append(result, Token{Pos: term.Pos, Word: term.Word, Nodes: tok.Nodes})
				}
			}
			req.Matches = result
			req.Valid = true
			go SendTokenizerToChannel(req.Out, req)
			break
		case TokenizerRemove:
			delete(dict, req.ID)
		}
	}
}
