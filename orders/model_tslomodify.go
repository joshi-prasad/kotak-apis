/*
 * Orders-API
 *
 * Order Placement API's
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package orders

// TSLO Order parmaeter for Product TSLO
type Tslomodify struct {
	// Order Indicator to modify Order
	OrderIndicator int64 `json:"orderIndicator,omitempty"`
	// Spread of the order
	Spread float32 `json:"spread,omitempty"`
	// Triling price of TSLO Order.
	TrailingPrice float32 `json:"trailingPrice,omitempty"`
}