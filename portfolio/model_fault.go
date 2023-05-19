/*
 * Portfolio - API
 *
 * This is a sample server for Kotak Trade API - Portfolio
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package portfolio

// Fault Details
type Fault struct {
	// error code
	Code int32 `json:"code,omitempty"`
	// error message
	Message string `json:"message,omitempty"`
	// error description
	Description string `json:"description,omitempty"`
}
