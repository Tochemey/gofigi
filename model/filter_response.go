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

package model

// FilterResponse is used to handle the SearchRequest api call
// ref: https://www.openfigi.com/api#post-v3-filter
type FilterResponse struct {
	// Data contains the actual search response
	Data  []*Data `json:"data,omitempty"`  // Data is present when Figi(s) are found for the associated
	Error string  `json:"error,omitempty"` // Error when there was an error when processing the request.
	Next  string  `json:"next,omitempty"`  // Next represents the next page token data to fetch.
	Total int     `json:"total"`           // Total represents how many figis match filter options
}
