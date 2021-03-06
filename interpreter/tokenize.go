// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	goahocorasick "github.com/anknown/ahocorasick"
	"github.com/cuttle-ai/octopus/datetime"
)

/*
 * This file contains the utilities for tokenizing a sentence.
 */

//Tokenize will tokenize a given sentence according to the tokenizer of the given id
func Tokenize(id string, sentence []rune) ([]FastToken, error) {
	/*
	 * We will initiate the datetime service
	 * Then we will prepcache the dictionary for the id
	 * We will make a request to the tokenizer to get the sentence tokenized
	 * Then we will build the unknowns
	 * Then we will adjust the date nodes
	 * Then we will adjust the postions
	 * Then we will do a fast token for all the tokens and return the same
	 */
	//start checking for the dates
	ch, err := StartCheckingForDates(sentence)
	if err != nil {
		//error while checking for the dates
		return nil, err
	}

	//precaching
	dictReq := DICTRequest{ID: id, Type: DICTPreCache, Out: make(chan DICTRequest)}
	SendDICTToChannel(DICTInputChannel, dictReq)
	<-dictReq.Out

	//requesting the tokenizer to tokenize the given sentence
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

	//adjusting the date fields
	//adding the date fields before adding unknowns is important as the
	//positions still refers to the original position in the sentence
	res.Matches = BuildTimeNodes(res.Matches, ch)

	//building the unknowns
	res.Matches = BuildUnknowns(sentence, res.Matches)

	//adjusting the positions of the tokens according to the position in the token list
	res.Matches = AdjustPositions(res.Matches)

	//fast tokenizing the tokens
	result := []FastToken{}
	for _, v := range res.Matches {
		result = append(result, v.FastToken())
	}
	return result, nil
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
				if len(word) == 0 {
					continue
				}
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

//BuildUnknowns build unknown nodes.
//It return the tokens with unidentified words in the sentence transfomed to tokens with unknowns
func BuildUnknowns(sentence []rune, toks []Token) []Token {
	/*
	 * We will iterate through the sentence
	 * Till the positions of each token is reached we will itearte through the sentence and build each word
	 * Upon reaching a position, it will add the word to the words list
	 * Once all the unknown words are recorded, we will split each word based on space
	 * The resultant words are consolidated as tokens
	 */
	result := []Token{}
	words := [][]rune{}
	tok := 0
	word := []rune{}
	tokenMap := map[string]Token{}

	//iterating through the sentence
	for p := 0; p < len(sentence); p++ {

		//checking whether the rune is outside the token
		if (len(toks) > tok && p < toks[tok].Pos) || len(toks) <= tok {
			word = append(word, sentence[p])
		} else if len(toks) > tok && toks[tok].Pos == p {

			//need to append only if there exist a word
			if len(word) > 0 {
				words = append(words, word)
				word = []rune{}
			}
			tokenMap[string(toks[tok].Word)] = toks[tok]
			words = append(words, toks[tok].Word)
			p = toks[tok].Pos + len(toks[tok].Word)
			tok++
		}
	}

	//if any pending word, copy it to the words list
	if len(word) > 0 {
		words = append(words, word)
	}

	//split the words with space
	for i, w := range words {
		// ws := strings.Split(string(w), " ")
		// for i, wr := range ws {
		//no need to take empty strings
		if len(w) == 0 {
			continue
		}
		tok := Token{
			Word: []rune(w),
		}
		oTok, ok := tokenMap[string(w)]
		if ok {
			tok.Nodes = oTok.Nodes
		} else {
			tok.Nodes = []Node{&UnknownNode{UID: fmt.Sprint("U", i), Word: []rune(w)}}
		}
		result = append(result, tok)
		// }
	}

	return result
}

//AdjustPositions will adjust the positions of the tokens corresponding to the
//position of the token in the token list rather than the sentence
func AdjustPositions(toks []Token) []Token {
	/*
	 * We will iterate through the tokens and put the position as the index of the token
	 */
	result := []Token{}

	for k, tok := range toks {
		result = append(result, Token{Pos: k, Word: tok.Word, Nodes: tok.Nodes})
	}

	return result
}

//StartCheckingForDates will initaite the service which starts checking for the date/time presence in the query
func StartCheckingForDates(sentence []rune) (chan datetime.Results, error) {
	ser, err := datetime.DefaultService()
	if err != nil {
		return nil, errors.New("Error while getting the date service" + err.Error())
	}
	return ser.Query(sentence), nil
}

type timeResults []datetime.Response

func (r timeResults) Len() int           { return len(r) }
func (r timeResults) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r timeResults) Less(i, j int) bool { return r[i].Start < r[j].Start }

//BuildTimeNodes will insert time nodes if a valid response from date service is received for its place
//This function won't replace any existing tokens. If conflict between existing node and time node come,
// the time node will be skipped with priority given to the existing node.
func BuildTimeNodes(toks []Token, ch chan datetime.Results) []Token {
	/*
	 * Then we will wait for the channel to dump results
	 * We will check whether the results are valid or not
	 * Then we will process it
	 */
	//wait for the channel to dump response
	result := &datetime.Results{}
	select {
	case res := <-ch:
		result = &res
	case <-time.After(3 * time.Second):
		break
	}

	//checking whether the results are valid or not
	if !result.IsValid() {
		return toks
	}

	//processing the result
	//first we will sort the valid results according to the start and end position
	//then we will insert the tokens for time nodes in between the tokens
	validResults := []datetime.Response{}
	for _, resp := range result.Res {
		if resp.IsValid() {
			validResults = append(validResults, resp)
		}
	}
	sort.Sort(timeResults(validResults))
	i, j := 0, 0
	for i < len(validResults) {
		result := validResults[i]
		startIndex := result.Start
		endIndex := result.End
		//if the end index is < the token's position, then we can
		//insert the result as a time node to the tokens
		if endIndex < toks[j].Pos {
			tok := Token{
				Pos:  startIndex,
				Word: []rune(result.Body),
				Nodes: []Node{
					&TimeNode{
						UID:   "T" + strconv.Itoa(i),
						Word:  []rune(result.Value.Value),
						Value: result.Value,
					},
				},
			}
			toks = append(toks[:j], append([]Token{tok}, toks[j:]...)...)
			i++
		} else if startIndex <= toks[j].Pos {
			//if the time result has intersected a known node
			//then priority is for the known know node, so we skip and move ahead
			i++
			j++
		} else if startIndex > toks[j].Pos {
			//we haven't reached the relevant position in the token
			j++
		}

		if j >= len(toks) {
			//we have exhausted the token list now we will append
			//the remianing time results to the tokens list
			timeNodes := []Token{}
			for i < len(validResults) {

				result = validResults[i]
				startIndex = result.Start
				timeNodes = append(timeNodes, Token{
					Pos:  startIndex,
					Word: []rune(result.Body),
					Nodes: []Node{
						&TimeNode{
							UID:   "T" + strconv.Itoa(i),
							Word:  []rune(result.Body),
							Value: result.Value,
						},
					},
				})
				i++
			}
			toks = append(toks, timeNodes...)
		}
	}

	return toks
}
