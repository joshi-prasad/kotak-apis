/*
 * Quotes - KSEC
 *
 * This is a sample server for Kotak Trade API - Quote
 *
 * API version: v1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package quotes_ksec

type InlineResponse200 struct {
	// Name of the instrument
	InstrumentName string `json:"instrumentName,omitempty"`
	InstrumentToken int32 `json:"instrumentToken,omitempty"`
	// Last time in epoch format when this data was updated
	LastUpdatedTime int64 `json:"lastUpdatedTime,omitempty"`
	// Last time in epoch format when this scrip was traded
	LastTradedTime int64 `json:"lastTradedTime,omitempty"`
	// Last traded price
	LastPrice float32 `json:"lastPrice,omitempty"`
	// Last traded quantity
	LastTradedQuantity int32 `json:"lastTradedQuantity,omitempty"`
	// Total bid quantity
	TotalBuyQuantity int64 `json:"totalBuyQuantity,omitempty"`
	// Total ask quantity
	TotalSellQuantity int64 `json:"totalSellQuantity,omitempty"`
	// Average traded price for the day
	AverageTradedPrice float32 `json:"averageTradedPrice,omitempty"`
	// Open interest in case of derivateive contracts
	OpenInterest int64 `json:"openInterest,omitempty"`
	// Absolute change in price from previous day close
	NetChange float32 `json:"netChange,omitempty"`
	// Lower value of daily price range
	DprLow float32 `json:"dprLow,omitempty"`
	// Higher value of daily price range
	DprHigh float32 `json:"dprHigh,omitempty"`
	Ohlc *InlineResponse200Ohlc `json:"ohlc,omitempty"`
	Depth *InlineResponse200Depth `json:"depth,omitempty"`
}
