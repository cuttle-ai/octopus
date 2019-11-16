// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package tokenizer

import "time"

/*
 * This file contains the defnition of dictionary in the platform for tokenizing
 */

//DICTRequestType is the type of the request for the dictionary
type DICTRequestType uint

const (
	//DICTAdd adds a dictionary for the given id
	DICTAdd DICTRequestType = 1
	//DICTGet returns the dictionary of a given id
	DICTGet DICTRequestType = 2
	//DICTRemove the dictionary from the cache
	DICTRemove DICTRequestType = 3
)

//DICT holds the dictionary in the system
type DICT struct {
	//LastUsed indicates when the token was used last
	LastUsed time.Time
	//Map has tokens mapped to their word
	Map map[string]Token
}

//DICTRequest can be used to make a request to dictionary cache
type DICTRequest struct {
	//ID to which the dictionary belong to
	ID string
	//Type is the type of the dictionary request. It can have Add, Get, Remove
	Type DICTRequestType
	//DICT is the dictionary under watch
	DICT DICT
	//Valid indicates that the dict is valid. During get requests, if valid is false then cache couldn't find the dict
	Valid bool
	//Out channel for sending response to the requester
	Out chan DICTRequest
}

//DICTInputChannel is the input channel to communicate with the cache
var DICTInputChannel chan DICTRequest

func init() {
	DICTInputChannel = make(chan DICTRequest)
	go Dictionary(DICTInputChannel)
}

//SendDICTToChannel sends a dict request to the channel. This function is to be used with go routines so that
//dictionary isn't blocked by the requests
func SendDICTToChannel(ch chan DICTRequest, req DICTRequest) {
	ch <- req
}

//Dictionary is the cache for providing the dictionary to the platform on demand
func Dictionary(in chan DICTRequest) {
	/*
	 * We will go into an infinte loop
	 * Will wait for the requests to come through the channel
	 * Based on the type of the request we will remove them from memory
	 */
	dict := make(map[string]DICT)
	for {
		req := <-in
		switch req.Type {
		case DICTAdd:
			req.DICT.LastUsed = time.Now()
			dict[req.ID] = req.DICT
			go SendTokenizerToChannel(
				TokenizerInputChannel,
				Request{
					ID:        req.ID,
					Type:      TokenizerAdd,
					Out:       make(chan Request),
					Tokenizer: Tokenizer{Map: req.DICT.Map},
				})
			break
		case DICTGet:
			req.DICT, req.Valid = dict[req.ID]
			req.DICT.LastUsed = time.Now()
			dict[req.ID] = req.DICT
			go SendDICTToChannel(req.Out, req)
			break
		case DICTRemove:
			delete(dict, req.ID)
			go SendTokenizerToChannel(
				TokenizerInputChannel,
				Request{
					ID:   req.ID,
					Type: TokenizerRemove,
				})
			break
		}
	}
}
