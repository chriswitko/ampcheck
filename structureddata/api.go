//  Copyright 2016 Yoshi Yamaguchi
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package structureddata

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// EndPoint is unofficial API Endpoint to call Structured Data Testing Tool.
//   $ curl --data "url=https://ymotongpoo.github.io/demos/amp/amp.html" \
//  > -X POST https://structured-data-testing-tool.developers.google.com/sdtt/web
const EndPoint = "https://structured-data-testing-tool.developers.google.com/sdtt/web/validate"

// Client is used for API call.
var Client = http.DefaultClient

// Response is response data from API.
type Response struct {
	TripleGroups []struct {
		ErrorsByOwner struct {
			AMP int `json:"AMP"`
		} `json:"errorsByOwner"`
		WarningsByOwner struct {
			AMP int `json:"AMP"`
		} `json:"warningsByOwner"`
		Errors []*ErrorInfo `json:"errors"`
	} `json:"tripleGroups"`
}

// ErrorInfo has error related information in Response.
type ErrorInfo struct {
	OwnerSet struct {
		AMP bool `json:"AMP"`
	} `json:"ownerSet"`
	Args            []string `json:"args"`
	OwnerToSeverity struct {
		AMP string `json:"AMP"`
	} `json:"ownerToSeverity"`
}

// NumErrors returns total number of errors and warnings for AMP respectively.
func (r *Response) NumErrors() (int, int) {
	errors, warnings := 0, 0
	for _, t := range r.TripleGroups {
		errors += t.ErrorsByOwner.AMP
		warnings += t.WarningsByOwner.AMP
	}
	return errors, warnings
}

// ValidateURL post urlStr to API and get result.
func ValidateURL(urlStr string) (*Response, error) {
	req, err := http.NewRequest("POST", EndPoint, nil)
	if err != nil {
		return nil, err
	}
	v := url.Values{}
	v.Add("url", urlStr)
	req.Form = v
	return callAPI(req)
}

// Validate post AMP HTML data in r and get result.
func Validate(r io.Reader) (*Response, error) {
	req, err := http.NewRequest("POST", EndPoint, r)
	if err != nil {
		return nil, err
	}
	return callAPI(req)
}

func callAPI(req *http.Request) (*Response, error) {
	resp, err := Client.Do(req)
	if err != nil {
		return nil, err
	}

	// Response from the API starts from unnecessary byte array ")]}'",
	// so throw that part away.
	ignore := make([]byte, 4)
	_, err = resp.Body.Read(ignore)
	if err != nil {
		return nil, err
	}

	var apiResp Response
	err = json.NewDecoder(resp.Body).Decode(&apiResp)
	if err != nil {
		return nil, err
	}
	return &apiResp, nil
}
