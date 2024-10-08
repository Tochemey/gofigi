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

// Data represents a figi data
type Data struct {
	Figi                string `json:"figi"`                          // Figi the 12-character alphanumerical code of the instrument
	Name                string `json:"name,omitempty"`                // Name represents the instrument name
	Ticker              string `json:"ticker,omitempty"`              // Ticker represents the instrument ticker symbol
	ExchCode            string `json:"exchCode,omitempty"`            // ExchCode the exchange code
	CompositeFIGI       string `json:"compositeFIGI,omitempty"`       // CompositeFIGI represents the instrument composite Figi
	SecurityType        string `json:"securityType,omitempty"`        // SecurityType represents the security type of the instrument
	MarketSector        string `json:"marketSector,omitempty"`        // MarketSector represents the market sector data of the instrument
	ShareClassFIGI      string `json:"shareClassFIGI,omitempty"`      // ShareClassFIGI represents the share class Figi
	SecurityType2       string `json:"securityType2,omitempty"`       // SecurityType2 represents the security type 2 information
	SecurityDescription string `json:"securityDescription,omitempty"` // SecurityDescription the security description
}
