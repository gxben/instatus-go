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

import "net/http"

// OverrideOptionsFunc is a type to override the Options struct.
type OverrideOptionsFunc func(o Options) Options

// Options represents the instatus.Client options to use to customize
// the HTTP client.
type Options struct {
	userAgent   string
	accessToken string
	httpClient  http.Client
}

// WithUserAgent overrides the default user agent when requests
// are sent out.
func WithUserAgent(agent string) OverrideOptionsFunc {
	return func(o Options) Options {
		o.userAgent = agent
		return o
	}
}

// WithToken overrides the access token, if no token is provided, then
// the client will panic.
func WithToken(token string) OverrideOptionsFunc {
	return func(o Options) Options {
		o.accessToken = token
		return o
	}
}

// WithHttpClient is the net/http client to use if you wish to extend it.
func WithHttpClient(client http.Client) OverrideOptionsFunc {
	return func(o Options) Options {
		o.httpClient = client
		return o
	}
}
