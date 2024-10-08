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
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/carlmjohnson/requests"
	"golang.org/x/time/rate"

	"github.com/tochemey-lab/gofigi/model"
)

// FilterClient implements the Filter API
// reference: https://www.openfigi.com/api#post-v3-filter
type FilterClient struct {
	client
}

// NewDefaultFilterClient creates a new instance of FilterClient using the http.DefaultClient
func NewDefaultFilterClient(apiKey string) *FilterClient {
	// create the instance of the SearchClient
	return &FilterClient{
		client: client{
			httpClient: http.DefaultClient,
			// let us set the rate limits according the API doc for 20 requests per minute
			// reference: https://www.openfigi.com/api#rate-limit
			limiter: rate.NewLimiter(rate.Every(1*time.Minute/20), 20),
			apiKey:  apiKey,
			baseURL: baseURL,
		},
	}
}

// NewFilterClient creates a new instance of SearchClient with the given http.Client
func NewFilterClient(apiKey string, httpClient *http.Client) *FilterClient {
	// create an instance of SearchClient
	return &FilterClient{
		client: client{
			httpClient: httpClient,
			// let us set the rate limits according the API doc for 20 requests per minute
			// reference: https://www.openfigi.com/api#rate-limit
			limiter: rate.NewLimiter(rate.Every(1*time.Minute), 20),
			apiKey:  apiKey,
			baseURL: baseURL,
		},
	}
}

// Filter fetches the list of stocks' information using the polygon API across all exchanges
// reference: https://www.openfigi.com/api#post-v3-filter
func (c FilterClient) Filter(ctx context.Context, request *model.FilterRequest) (*model.FilterResponse, error) {
	// make a copy of the request before using it
	requestCopy := request
	// create an instance of the search response
	response := new(model.FilterResponse)
	// build the http request
	requestBuilder := requests.
		URL(c.baseURL).
		Client(c.httpClient).
		Path(filterResourcePath).
		Method(http.MethodPost).
		Header("X-OPENFIGI-APIKEY", c.apiKey).
		Header("Content-Type", "application/json; charset=utf-8").
		Header("Accept", "application/json; charset=utf-8").
		BodyJSON(requestCopy).
		ToJSON(response).
		CheckStatus(http.StatusOK, http.StatusTooManyRequests)

	// This is a blocking call. Honors the rate limit
	err := c.limiter.Wait(ctx)
	if err != nil {
		return nil, fmt.Errorf("rate limiter reached: %w", err)
	}
	// execute the request and  handle the error
	if err = requestBuilder.Fetch(ctx); err != nil {
		return nil, err
	}
	// return the response
	return response, nil
}
