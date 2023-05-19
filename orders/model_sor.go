/*
 * Orders-API
 *
 * Order Placement API's
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package orders

// SOR Order Details Repsonse.
type Sor struct {
	// Price of the order at which order is placed
	Price float32 `json:"price,omitempty"`
	// Qunatity of the order placed in exchange.
	Quantity float32 `json:"quantity,omitempty"`
	// Order confirmation mesage
	Mesage string `json:"mesage,omitempty"`
	// Order ID of the order to be modified
	OrderId int64 `json:"orderId,omitempty"`
}