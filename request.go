// 🪁 instatus-go: Lightweight and speedy Go client for Instatus
// Copyright (c) 2022 Noel <cutie@floofy.dev>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package instatus

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// httpMethod is a string type to represent an Instatus request method.
type httpMethod string

var (
	// MethodGET is the httpMethod to use to create GET requests.
	MethodGET httpMethod = "GET"

	// MethodPOST is the httpMethod to use to create POST requests.
	MethodPOST httpMethod = "POST"

	// MethodPUT is the httpMethod to use to create PUT requests.
	MethodPUT httpMethod = "PUT"

	// MethodDELETE is the httpMethod to use to create DELETE requests.
	MethodDELETE httpMethod = "DELETE"
)

// String stringifies the httpMethod to its correspondent value.
func (s httpMethod) String() string {
	switch s {
	case MethodGET:
		return "GET"

	case MethodPUT:
		return "PUT"

	case MethodPOST:
		return "POST"

	case MethodDELETE:
		return "DELETE"

	default:
		return "?"
	}
}

// PerformRequest creates a request to the Instatus API and returns the data
// from the `data` parameter.
func (c Client) PerformRequest(
	endpoint string,
	method httpMethod,
	data interface{},
	body io.Reader,
) error {
	req, err := http.NewRequest(method.String(), "https://api.instatus.com/v1"+endpoint, body)
	if err != nil {
		return err
	}

	// add the `Authorization` header to the request
	req.Header.Add("Authorization", "Bearer "+c.options.accessToken)

	// Add the `User-Agent` header to the request
	req.Header.Add("User-Agent", c.options.userAgent)

	// Set the content type to `application/json`
	req.Header.Add("Content-Type", "application/json")

	if method == MethodGET && body != nil {
		return errors.New("cannot send payload in GET requests")
	}

	if body != nil {
		req.Body = io.NopCloser(body)
	}

	res, err := c.options.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Before we do anything, let's check if the status code is NOT 200
	if res.StatusCode != 200 {
		var errorData map[string]interface{}
		if err := json.Unmarshal(b, &errorData); err != nil {
			return err
		} else {
			errorOwo := errorData["error"].(map[string]interface{})
			return fmt.Errorf("%s [%d %s]: %s",
				errorOwo["code"].(string),
				res.StatusCode,
				res.Status,
				errorOwo["message"].(string),
			)
		}
	}

	// Try to convert it to JSON
	if err := json.Unmarshal(b, data); err != nil {
		return err
	} else {
		return nil
	}
}
