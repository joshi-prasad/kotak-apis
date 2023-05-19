/*
 * Portfolio - API
 *
 * This is a sample server for Kotak Trade API - Portfolio
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package portfolio

type InlineResponse2003SuccessDarivatives struct {
	Futures        float64                                              `json:"futures,omitempty"`
	Options        float64                                              `json:"options,omitempty"`
	Total          float64                                              `json:"total,omitempty"`
	FuturesDetails []InlineResponse2003SuccessDarivativesFuturesDetails `json:"futuresDetails,omitempty"`
	OptionsDetails []InlineResponse2003SuccessDarivativesFuturesDetails `json:"optionsDetails,omitempty"`
}
