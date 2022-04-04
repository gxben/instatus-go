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

// Package instatus is the main package entrypoint for using `instatus-go`. To use it
// in your application, you must initialize a new instatus.Client instance using the
// instatus.NewClient func.
package instatus

import (
	"errors"
	"github.com/gxben/instatus-go/types"
	"net/http"
)

// Client is the main client to use to interact with the Instatus API.
type Client struct {
	options Options
}

// NewClient creates a new Client instance.
func NewClient(opts ...OverrideOptionsFunc) Client {
	options := Options{
		httpClient:  http.Client{},
		accessToken: "",
		userAgent:   "auguwu/instatus-go",
	}

	for _, opt := range opts {
		options = opt(options)
	}

	// TODO: should this be as a (*Client, error) return signature?
	if options.accessToken == "" {
		panic(errors.New("missing accessToken property in NewClient"))
	}

	return Client{
		options: options,
	}
}

func (c Client) User() (*types.User, error) {
	var data *types.User
	if err := c.PerformRequest(
		"/user",
		MethodGET,
		data,
		nil,
	); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}
