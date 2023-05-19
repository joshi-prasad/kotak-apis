/*
 * Portfolio - API
 *
 * This is a sample server for Kotak Trade API - Portfolio
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package portfolio

type InlineResponse2007 struct {
	// Month of year i.e 1 - Jan, 2- Feb like wise
	Month int32 `json:"month,omitempty"`
	// turnover for mention month
	OptTurnOver float64 `json:"optTurnOver,omitempty"`
	// year of month
	Year float64 `json:"year,omitempty"`
}