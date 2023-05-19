/*
 * Quotes - KSEC
 *
 * This is a sample server for Kotak Trade API - Quote
 *
 * API version: v1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package quotes_ksec

type InlineResponse2002 struct {
	// Name of the instrument
	InstrumentName  string  `json:"instrumentName,omitempty"`
	InstrumentToken int32   `json:"instrumentToken,omitempty"`
	LastPrice       float32 `json:"lastPrice,omitempty"`
}
