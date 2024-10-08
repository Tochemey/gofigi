/*
 * MIT License
 *
 * Copyright (c) 2024 tochemey-lab
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

package gofigi

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/carlmjohnson/requests"
	"golang.org/x/time/rate"
)

const (
	searchResourcePath  = "/v3/search"
	mappingResourcePath = "/v3/mapping"
	filterResourcePath  = "/v3/filter"
	baseURL             = "https://api.openfigi.com"
)

// client is an implementation of the v3  API
type client struct {
	httpClient *http.Client  // nolint
	limiter    *rate.Limiter // nolint
	apiKey     string        // the v3 api key
	baseURL    string        // the v3 api endpoint
}

// checkStatus validates the response has an acceptable status code.
func checkStatus(acceptStatuses ...int) requests.ResponseHandler {
	// return the response handler
	return func(res *http.Response) error {
		for _, code := range acceptStatuses {
			if res.StatusCode == code {
				return nil
			}
		}
		// cast error into error response
		se := (*requests.ResponseError)(res)
		// read the error response body
		var buf strings.Builder
		_, copyErr := io.Copy(&buf, se.Body)
		if copyErr == nil {
			return fmt.Errorf("response error for %s: unexpected status: %d and mesg: %s",
				se.Request.URL.Redacted(), res.StatusCode, buf.String())
		}
		return copyErr
	}
}
