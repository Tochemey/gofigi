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

const (
	IDTypeTicker             = "TICKER"                         // IDTypeTicker represents the Ticker ID Type
	IDTypeISIN               = "ID_ISIN"                        // IDTypeISIN represents the ISIN ID Type
	IDTypeSedol              = "ID_SEDOL"                       // IDTypeSedol represents the SEDOL ID Type
	IDTypeCusip              = "ID_CUSIP"                       // IDTypeCusip represents the CUSIP ID Type
	IDTypeCompositeFIGI      = "COMPOSITE_ID_BB_GLOBAL"         // IDTypeCompositeFIGI  represents the Composite Financial Instrument Global Identifier Type
	IDTypeShareClassFIGI     = "ID_BB_GLOBAL_SHARE_CLASS_LEVEL" // IDTypeShareClassFIGI represents the Share Class Financial Instrument Global Identifier Type
	IDTypeFIGI               = "ID_BB_GLOBAL"                   // IDTypeFIGI represents the Figi Type
	IDTypeLegacyBBUnique     = "ID_BB_UNIQUE"                   // IDTypeLegacyBBUnique represents the Unique Bloomberg Identifier - A legacy, internal Bloomberg identifier
	IDTypeCommon             = "ID_COMMON"                      // IDTypeCommon is the Common Code - A nine digit identification number.
	IDTypeWKN                = "ID_WERTPAPIER"                  // IDTypeWKN is the Wertpapierkennnummer/WKN - German securities identification code.
	IDTypeCINS               = "ID_CINS"                        // IDTypeCINS is the CINS - CUSIP International Numbering System
	IDTypeBB                 = "ID_BB"                          // IDTypeBB is the legacy Bloomberg identifier.
	IDTypeBB8                = "ID_BB_8_CHR"                    // IDTypeBB8 is the legacy Bloomberg identifier (8 characters only).
	IDTypeFINRAID            = "ID_TRACE"                       // IDTypeFINRAID is the Trace eligible bond identifier issued by FINRA
	IDTypeItalyNumber        = "ID_ITALY"                       // IDTypeItalyNumber is the Italian Identifier Number - The Italian Identification number consisting of five or six digits.
	IDTypeExchangeSymbol     = "ID_EXCH_SYMBOL"                 // IDTypeExchangeSymbol is the Local Exchange Security Symbol - Local exchange security symbol
	IDTypeFullExchangeSymbol = "ID_FULL_EXCHANGE_SYMBOL"        // IDTypeFullExchangeSymbol is the Full Exchange Symbol - Contains the exchange symbol for futures, options, indices inclusive of base symbol and other security elements.
	IDTypeSecurityNumberDesc = "ID_BB_SEC_NUM_DES"              // IDTypeSecurityNumberDesc is the Security ID Number Description - Descriptor for a financial instrument. Similar to the ticker field, but will provide additional metadata data.
	IDTypeBaseTicker         = "BASE_TICKER"                    // IDTypeBaseTicker is the  indistinct identifier which may be linked to multiple instruments. May need to be combined with other values to identify a unique instrument.
	IDTypeCusip8             = "ID_CUSIP_8_CHR"                 // IDTypeCusip8 is the CUSIP (8 Characters Only) - Committee on Uniform Securities Identification Procedures.
	IDTypeOCCSymbol          = "OCC_SYMBOL"                     // IDTypeOCCSymbol is the OCC Symbol - A twenty-one character option symbol standardized by the Options Clearing Corporation (OCC) to identify a U.S. option.
	IDTypeFutureOptionID     = "UNIQUE_ID_FUT_OPT"              // IDTypeFutureOptionID is the Unique Identifier for Future Option - Bloomberg unique ticker with logic for index, currency, single stock futures, commodities and commodity options.
	IDTypeOPRASymbol         = "OPRA_SYMBOL"                    // IDTypeOPRASymbol is the OPRA Symbol - Option symbol standardized by the Options Price Reporting Authority (OPRA) to identify a U.S. option
	IDTypeTradingSysID       = "TRADING_SYSTEM_IDENTIFIER"      // IDTypeTradingSysID is the Trading System Identifier - Unique identifier for the instrument as used on the source trading system
	IDTypeShortCode          = "ID_SHORT_CODE"                  // IDTypeShortCode is the  exchange venue specific code to identify fixed income instruments primarily traded in Asia.
	IDTypeVendorIndex        = "VENDOR_INDEX_CODE"              // IDTypeVendorIndex is the Index code assigned by the index provider for the purpose of identifying the security.
)
