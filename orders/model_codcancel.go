/*
 * Orders-API
 *
 * Order Placement API's
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package orders

// Call of the day Order parmaeter for Product COD
type Codcancel struct {
	// Order Number to modify
	OrderNo string `json:"orderNo,omitempty"`
}
