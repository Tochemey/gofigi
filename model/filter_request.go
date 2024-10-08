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

// FilterRequest is used to fetch stocks' information
// across all markets using the v3 api
// ref: https://www.openfigi.com/api#post-v3-filter
type FilterRequest struct {
	// Query represents any keyword to include in the search
	Query string `json:"query,omitempty"`
	// SecurityType security type of the desired instrument(s)
	// Accepted values: https://www.openfigi.com/api/enumValues/v3/securityType
	SecurityType string `json:"securityType,omitempty"`
	// MarketSecDes is the  Market sector description of the desired instrument(s).
	MarketSecDes string `json:"marketSecDes,omitempty"`
	// IncludeUnlistedEquities Set to true to include equity instruments that are not listed on an exchange.
	IncludeUnlistedEquities bool `json:"includeUnlistedEquities,omitempty"`
	// ExchCode Exchange code of the desired instrument(s)
	// Accepted values: https://www.openfigi.com/api/enumValues/v3/exchCode
	ExchCode string `json:"exchCode,omitempty"`
	// Start The number of results returned for any given request is fixed.
	// When more results are accessible, the response will contain a next property whose value should be sent in succeeding requests as the value of the start property.
	// This will notify the API to return the next "page" of results.
	Start string `json:"start,omitempty"`
	// MicCode is the ISO market identification code(MIC) of the desired instrument(s) (cannot use in conjunction with exchCode).
	// Accepted values: https://www.openfigi.com/api/enumValues/v3/micCode
	MicCode string `json:"micCode,omitempty"`
	// Currency is the currency associated to the desired instrument(s)
	// Accepted values: https://www.openfigi.com/api/enumValues/v3/currency
	Currency string `json:"currency,omitempty"`
	// SecurityType2 is An alternative security type of the desired instrument(s). securityType2 is typically less specific than securityType
	// Use MarketSecDes if SecurityType2 is not available
	SecurityType2 string `json:"securityType2,omitempty"`
	// OptionType helps filter instruments based on option type
	OptionType string `json:"optionType,omitempty"`
	// Strike helps find instruments whose strike price falls in an interval.
	// Kindly read the v3 API doc to understand the different type of values to set.
	Strike []any `json:"strike,omitempty"`
	// ContractSize helps find instruments whose contract size falls in an interval.
	// Kindly read the v3 API doc to understand the different type of values to set.
	ContractSize []any `json:"contractSize,omitempty"`
	// Coupon helps  find instruments whose coupon falls in an interval.
	// Kindly read the v3 API doc to understand the different type of values to set.
	Coupon []any `json:"coupon,omitempty"`
	// Expiration helps find instruments whose expiration date falls in an interval.
	// Date strings must be of the form YYYY-MM-DD
	Expiration []string `json:"expiration,omitempty"`
	// Maturity helps find instruments whose maturity date falls in an interval.
	// This is required when SecurityType2 is set to Pool
	// Date strings must be of the form YYYY-MM-DD
	Maturity []string `json:"maturity,omitempty"`
	// StateCode is the state code of the instrument
	// Accepted values: https://www.openfigi.com/api/enumValues/v3/stateCode
	StateCode string `json:"stateCode,omitempty"`
}
