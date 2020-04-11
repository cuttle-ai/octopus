// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package interpreter

import (
	"sync"
	"time"
)

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

//DICTClearCheckInterval is the interval after which the dict removal check has to run
const DICTClearCheckInterval = time.Minute * 20

//DICTExpiry is the expiry time after which the dictionary expiries without any active usage
const DICTExpiry = time.Hour * 4

//DICT holds the dictionary in the system
type DICT struct {
	//LastUsed indicates when the token was used last
	LastUsed time.Time
	//Map has tokens mapped to their word
	Map map[string]Token
}

//Copy returns the deep copy of the dict
func (d DICT) Copy() DICT {
	res := DICT{LastUsed: d.LastUsed, Map: map[string]Token{}}
	for k, v := range d.Map {
		res.Map[k] = v.Copy()
	}
	return res
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

//DICTAggregator is the aggregator to get the dict from a service or database
type DICTAggregator interface {
	Get(ID string) (DICT, error)
}

//DICTInputChannel is the input channel to communicate with the cache
var DICTInputChannel chan DICTRequest

//defaultAggregator to be used in the dictionary
var defaultAggregator aggregator

type aggregator struct {
	agg DICTAggregator
	m   sync.Mutex
}

//SetDefaultDICTAggregator sets the default aggregator as the passed param
func SetDefaultDICTAggregator(agg DICTAggregator) {
	defaultAggregator.m.Lock()
	defaultAggregator.agg = agg
	defaultAggregator.m.Unlock()
}

func getDICT(ID string) (DICT, bool) {
	defaultAggregator.m.Lock()
	if defaultAggregator.agg == nil {
		return DICT{}, false
	}
	d, err := defaultAggregator.agg.Get(ID)
	if err != nil {
		return DICT{}, false
	}
	defaultAggregator.m.Unlock()
	return d, true
}

func init() {
	DICTInputChannel = make(chan DICTRequest)
	defaultAggregator = aggregator{}
	go Dictionary(DICTInputChannel)
	go cacheClearCheck(DICTInputChannel)
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
			if !req.Valid {
				req.DICT, req.Valid = getDICT(req.ID)
			}
			//we will only give the copies of the dict item to avoid mutation
			req.DICT = req.DICT.Copy()
			req.DICT.LastUsed = time.Now()
			dict[req.ID] = req.DICT
			go SendDICTToChannel(req.Out, req)
			break
		case DICTRemove:
			//we will iterate over the cache and check the last usage
			t := time.Now()
			for k, v := range dict {
				if v.LastUsed.Add(DICTExpiry).After(t) {
					continue
				}
				delete(dict, k)
				go SendTokenizerToChannel(
					TokenizerInputChannel,
					Request{
						ID:   k,
						Type: TokenizerRemove,
					})
			}
			break
		}
	}
}

func cacheClearCheck(in chan DICTRequest) {
	for {
		time.Sleep(DICTClearCheckInterval)
		go SendDICTToChannel(in, DICTRequest{Type: DICTRemove})
	}
}
