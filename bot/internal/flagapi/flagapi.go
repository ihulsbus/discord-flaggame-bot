/*
Copyright (c) 2023 Ian Hulsbus

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/

package flagapi

import (
	"encoding/json"
	m "flaggame/internal/models"
	"net/http"
)

type FlagApi struct {
	httpClient *http.Client
	baseURL    string
}

func FlagApiConstructor(baseURL string, client *http.Client) *FlagApi {

	// Return lib with provided http client if it is passed on init
	if client != nil {
		return &FlagApi{
			baseURL:    baseURL,
			httpClient: client,
		}
	}

	// Return lib with own http client (default)
	return &FlagApi{
		baseURL:    baseURL,
		httpClient: httpClient(),
	}
}

func (f FlagApi) GetRandomFlag() (*m.RandomCountryResponse, error) {
	var response m.RandomCountryResponse

	rawResponse, err := f.get("getRandomCountry")
	if err != nil {
		return nil, err
	}

	json.Unmarshal(rawResponse, &response)

	return &response, nil
}
