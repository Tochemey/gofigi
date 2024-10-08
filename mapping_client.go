/*
 * MIT License
 *
 * Copyright (c) 2024 tochemey
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
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/carlmjohnson/requests"
	"golang.org/x/time/rate"

	"github.com/tochemey/gofigi/model"
)

// MappingClient implements the Mapping API
// reference: https://www.openfigi.com/api#post-v3-mapping
type MappingClient struct {
	client
}

// NewDefaultMappingClient creates a new instance of MappingClient using the http.DefaultClient
func NewDefaultMappingClient(apiKey string) *MappingClient {
	// create the instance of the MappingClient
	return &MappingClient{
		client: client{
			httpClient: http.DefaultClient,
			// let us set the rate limits according the API doc for 25 per 7 seconds
			// reference: https://www.openfigi.com/api#rate-limit
			limiter: rate.NewLimiter(rate.Every(7*time.Second), 25),
			apiKey:  apiKey,
			baseURL: baseURL,
		},
	}
}

// NewMappingClient creates a new instance of MappingClient with the provided  http.Client
func NewMappingClient(apiKey string, httpClient *http.Client) *MappingClient {
	// create the instance of the MappingClient
	return &MappingClient{
		client: client{
			httpClient: httpClient,
			limiter:    rate.NewLimiter(rate.Every(7*time.Second/25), 25),
			apiKey:     apiKey,
			baseURL:    baseURL,
		},
	}
}

// Mappings maps third party identifiers to FIGIs.
// reference: https://www.openfigi.com/api#post-v3-mapping
func (c MappingClient) Mappings(ctx context.Context, mappingRequests []*model.MappingRequest) ([]*model.MappingResponse, error) {
	// per the rate limit the maximum jobs for mapping is 100
	// https://www.openfigi.com/api#rate-limit
	const jobsSizePerRequest = 100
	// make a copy of the request before using it
	mappingReqs := mappingRequests
	// create an instance of the search response
	var responses []*model.MappingResponse
	// let us chunk the data to send
	batches := chunkMappingRequestsSlice(mappingReqs, jobsSizePerRequest)
	// let us send each batchRequests
	for _, batchRequests := range batches {
		// create an instance of the responses for that batchRequests
		var batchResponses []*model.MappingResponse
		// build the http request
		requestBuilder := requests.
			URL(c.baseURL).
			Client(c.httpClient).
			Path(mappingResourcePath).
			Method(http.MethodPost).
			Header("X-OPENFIGI-APIKEY", c.apiKey).
			Header("Content-Type", "application/json; charset=utf-8").
			Header("Accept", "application/json; charset=utf-8").
			BodyJSON(batchRequests).
			Handle(func(response *http.Response) error {
				data, err := io.ReadAll(response.Body)
				if err != nil {
					return err
				}
				if err = json.Unmarshal(data, &batchResponses); err != nil {
					return fmt.Errorf("unmarshalling response body failed: %w", err)
				}
				return nil
			}).
			AddValidator(checkStatus(http.StatusOK, http.StatusTooManyRequests))

		// let us hook in the rate limit and handle the error
		err := c.limiter.Wait(ctx) // This is a blocking call. Honors the rate limit
		if err != nil {
			return nil, fmt.Errorf("rate limiter reached: %w", err)
		}

		// execute the request and  handle the error
		if err = requestBuilder.Fetch(ctx); err != nil {
			return nil, err
		}
		// add the batchRequests response to the overall list
		responses = append(responses, batchResponses...)
	}
	return responses, nil
}

// chunkMappingRequestsSlice helps chunk the size of mapping requests into the required chunkSize per
// each mapping api call.
// https://www.openfigi.com/api#rate-limit
func chunkMappingRequestsSlice(slice []*model.MappingRequest, chunkSize int) [][]*model.MappingRequest {
	var chunks [][]*model.MappingRequest
	for {
		if len(slice) == 0 {
			break
		}

		// necessary check to avoid slicing beyond
		// slice capacity
		if len(slice) < chunkSize {
			chunkSize = len(slice)
		}

		chunks = append(chunks, slice[0:chunkSize])
		slice = slice[chunkSize:]
	}

	return chunks
}
