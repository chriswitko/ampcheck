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
    "net/http"
    "testing"
)

var testURLs = []string{
    "https://ymotongpoo.github.io/demos/amp/amp.html",
}

func TestValidate(t *testing.T) {
    for _, u := range testURLs {
        src, err := http.Get(u)
        if err != nil {
            t.Errorf("%v", err)
        }
        resp, err := Validate(src.Body)
        if err != nil {
            t.Errorf("%v", err)
        }
        if resp == nil {
            t.Errorf("%v", err)
        }
    }
}

func TestValidateURL(t *testing.T) {
    for _, u := range testURLs {
        resp, err := ValidateURL(u)
        if err != nil {
            t.Errorf("%v", err)
        }
        if resp == nil {
            t.Errorf("nil returned.")
        }
    }
}