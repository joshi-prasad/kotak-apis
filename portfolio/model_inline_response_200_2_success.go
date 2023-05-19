/*
 * Portfolio - API
 *
 * This is a sample server for Kotak Trade API - Portfolio
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package portfolio

type InlineResponse2002Success struct {
	Debentures  *InlineResponse2002SuccessDebentures  `json:"debentures,omitempty"`
	ETF         *InlineResponse2002SuccessDebentures  `json:"ETF,omitempty"`
	Equity      *InlineResponse2002SuccessDebentures  `json:"equity,omitempty"`
	Darivatives *InlineResponse2002SuccessDarivatives `json:"darivatives,omitempty"`
	MutualFund  *InlineResponse2002SuccessMutualFund  `json:"mutualFund,omitempty"`
}
