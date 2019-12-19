// Copyright 2019 Melvin Davis<hi@melvindavis.me>. All rights reserved.
// Use of this source code is governed by a Melvin Davis<hi@melvindavis.me>
// license that can be found in the LICENSE file.

package datetime

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

/*
 * This file contains the implementation of the duckling as datetime service
 */

//Duckling is a datetime service
type Duckling struct {
	url      string
	parseAPI string
}

//NewDuckling returns a duckling service instance.
//If it couldn't be initialized due to missing env variables - DUCKLING_SERVER holding host address
//will return an error
func NewDuckling() (*Duckling, error) {

	url := "http://duckling.cuttle.ai:8000" //os.Getenv("DUCKLING_SERVER")
	if len(url) == 0 {
		return nil, errors.New("DUCKLING_SERVER host address is missing in environment variables. Add host address where duckling is running to use Duckling as datetime service")
	}
	return &Duckling{url: url, parseAPI: "/parse"}, nil
}

//Query returns a channel of Response. Which will return the reponse in an concurrent environment
func (d *Duckling) Query(query []rune) chan Results {
	ch := make(chan Results)
	go d.hitAPI(ch, string(query))
	return ch
}

func (d *Duckling) hitAPI(resp chan Results, query string) {
	/*
	 * We will hit the api
	 * Then we will decode the response
	 */
	//hitting the api
	result := Results{}
	formData := url.Values{
		"text": {query},
	}
	client := http.Client{Timeout: time.Second * 3}
	req, err := http.NewRequest("POST", d.url+d.parseAPI, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Println("Error while creating the request for duckling api", d.url+d.parseAPI, err)
		resp <- result
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error while htting the duckling api", d.url+d.parseAPI, err)
		resp <- result
		return
	}

	//parsing the response
	dec := json.NewDecoder(res.Body)
	dm := []Response{}
	er := dec.Decode(&dm)
	if er != nil {
		log.Println("Error while parsing the response from the duckling api", d.url+d.parseAPI, er)
		resp <- result
		return
	}

	result.Res = dm
	resp <- result
}
