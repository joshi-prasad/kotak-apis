/*
 * Watchlist-API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package watchlist

type InlineResponse2002Success struct {
	// Market Exchange of Security.
	Exchange string `json:"exchange,omitempty"`
	// Expiry Date of Token.
	ExpiryDate string `json:"expiryDate,omitempty"`
	// Unique Security Token.
	InstrumentToken float64 `json:"instrumentToken,omitempty"`
	// Instrument type of security.
	It string `json:"it,omitempty"`
	// Option type of security.
	OptionType string `json:"optionType,omitempty"`
	// Display Name of security.
	SecurityName string `json:"securityName,omitempty"`
	// Strike price of security.
	StrikePrice float64 `json:"strikePrice,omitempty"`
	// Sybol of Security
	Symbol string `json:"symbol,omitempty"`
	// LTP of security when added to watchlist.
	TokenAddedLtp float64 `json:"tokenAddedLtp,omitempty"`
	// time when security added to watchlist.
	TokenAddedTime string `json:"tokenAddedTime,omitempty"`
}
