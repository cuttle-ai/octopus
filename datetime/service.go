// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

//Package datetime has the utilities required for interpreting date/time from a nlp query
package datetime

import "time"

//TimeValue is the Value of the time stored
type TimeValue struct {
	//String value of the time
	Value string `json:"value,omitempty"`
	//Gran of the time to like hour, day etc described by the text
	Gran string `json:"gran,omitempty"`
	//Time stores the time value
	Time *time.Time `json:"-"`
	//Error will be set once isvalid is run if parsing failed
	Error error `json:"-"`
}

//Value is the value struct holding the time
type Value struct {
	//To indicates the to interval
	To *TimeValue `json:"to,omitempty"`
	//From indicates the from interval
	From *TimeValue `json:"from,omitempty"`
	//Type of value
	Type string `json:"type,omitempty"`
	//String value of the time
	Value string `json:"value,omitempty"`
	//Gran of the time to like hour, day etc described by the text
	Gran string `json:"gran,omitempty"`
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
	//Body is the string from which date is referred
	Body string `json:"body,omiempty"`
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

//IsValid checks whether the time value is valid or not
func (tv *TimeValue) IsValid() bool {
	/*
	 * If an error exists or parsing failed we will return false
	 * If time exists or parsing was successful we will return true
	 */
	if tv.Error != nil {
		return false
	}
	if tv.Time != nil {
		return true
	}
	t, err := time.Parse(time.RFC3339, tv.Value)
	if err != nil {
		tv.Error = err
		return false
	}
	tv.Time = &t
	return true
}

//IsValid checks whether the value is valid or not
func (v *Value) IsValid() bool {
	/*
	 * if type is value then will try to do the normal flow
	 * Else if it is interval we will parse from and to if available
	 */
	if v.Error != nil {
		return false
	}
	if v.Time != nil {
		return true
	}
	if v.Type == "value" {
		t, err := time.Parse(time.RFC3339, v.Value)
		if err != nil {
			v.Error = err
			return false
		}
		v.Time = &t
		return true
	}
	if v.Type == "interval" {
		if v.From == nil && v.To == nil {
			return false
		}
		resT, resF := false, false
		if v.From != nil {
			resF = v.From.IsValid()
		}
		if v.To != nil {
			resT = v.To.IsValid()
		}
		return resT || resF
	}
	return false
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
