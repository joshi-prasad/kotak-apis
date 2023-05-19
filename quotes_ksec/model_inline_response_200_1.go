/*
 * Quotes - KSEC
 *
 * This is a sample server for Kotak Trade API - Quote
 *
 * API version: v1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package quotes_ksec

type InlineResponse2001 struct {
	// Name of the instrument
	InstrumentName string `json:"instrumentName,omitempty"`
	InstrumentToken int32 `json:"instrumentToken,omitempty"`
	Open float32 `json:"open,omitempty"`
	High float32 `json:"high,omitempty"`
	Low float32 `json:"low,omitempty"`
	Close float32 `json:"close,omitempty"`
}
