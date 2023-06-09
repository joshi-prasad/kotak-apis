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
type Codplace struct {
	// Square off flag for COD order, 1 to auto square off order. 0 - for no auto Square off by system.
	SquareOffFlag int32 `json:"squareOffFlag,omitempty"`
}
