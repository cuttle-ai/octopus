// Copyright 2019 Cuttle.ai. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

//Package response handles utilities for writing error and normal responses to the response writer
package response

import (
	"encoding/json"
	"net/http"

	"github.com/cuttle-ai/octopus/lsp/log"
)

/*
 * This file contains the response templates
 */

//Error is the datastructure for writing error response
type Error struct {
	//Err is the error happened in string format
	Err string `json:"error"`
}

//WriteError will write to the error response to the response writer
func WriteError(res http.ResponseWriter, err Error, code int) {
	/*
	 * Will use json encoder to write response
	 */
	res.WriteHeader(code)
	en := json.NewEncoder(res)
	er := en.Encode(err)
	if er != nil {
		//Error while writing the response
		log.Error("Error while writing the error response", er)
	}
}

//Write will write the response to the response writer
//payload is any json serializable object
func Write(res http.ResponseWriter, payload interface{}) {
	/*
	 * Will use json encoder to write response
	 */
	en := json.NewEncoder(res)
	er := en.Encode(payload)
	if er != nil {
		//Error while writing the response
		log.Error("Error while writing the response", er)
	}
}
