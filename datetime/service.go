// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

//Package datetime has the utilities required for interpreting date/time from a nlp query
package datetime

import "time"

//Value is the value struct holding the time
type Value struct {
	//String value of the time
	Value string `json:"value,omitempty"`
	//Grain of the time to like hour, day etc described by the text
	Grain string `json:"grain,omitempty"`
	//Time stores the time value
	Time *time.Time `json:"-"`
	//Error will be set once isvalid is run if parsing failed
	Error error `json:"-"`
}

//Response is the response of the package for external use
type Response struct {
	//Start index of the words denoting date-time in the given query
	Start int `json:"start,omitempty"`
	//End index of the words denoting date-time in the given query
	End int `json:"end,omitempty"`
	//Dim is the dimension of the response. We are intrested only if the Dim is time
	Dim string `json:"dim,omitempty"`
	//Value stores the value of the response.
	Value Value `json:"value,omitempty"`
}

//Results has list of responses from the service
type Results struct {
	//Res has the list of response
	Res []Response
}

//Service interface produces the querying service
type Service interface {
	Query(query []rune) chan Results
}

//IsValid checks whether the value is valid or not
func (v *Value) IsValid() bool {
	/*
	 * If an error exists or parsing failed we will return false
	 * If time exists or parsing was successful we will return true
	 */
	if v.Error != nil {
		return false
	}
	if v.Time != nil {
		return true
	}
	t, err := time.Parse(time.RFC3339, v.Value)
	if err != nil {
		v.Error = err
		return false
	}
	v.Time = &t
	return true
}

//IsValid denotes whether the response valid or not. It is valid only if the underlying value is valid and dimension is time
func (r *Response) IsValid() bool {
	return r.Value.IsValid() && r.Dim == "time"
}

//IsValid will check and update the state of the results
func (r *Results) IsValid() bool {
	one := false
	for i := 0; i < len(r.Res); i++ {
		if (&(r.Res[i])).IsValid() {
			one = true
		}
	}
	return one
}

//DefaultService will return a service to provide the interpretation of datetime
func DefaultService() (Service, error) {
	return NewDuckling()
}
