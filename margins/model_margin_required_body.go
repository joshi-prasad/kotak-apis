/*
 * Margins-API
 *
 * Margin API's
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package margins

type MarginRequiredBody struct {
	// Transaction Type - BUY or SELL
	TransactionType string `json:"transactionType,omitempty"`
	OrderInfo []OrderInfo `json:"orderInfo,omitempty"`
}
