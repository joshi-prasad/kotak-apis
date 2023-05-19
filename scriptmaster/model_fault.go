/*
 * ScripMaster-API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package scriptmaster

// Fault Details
type Fault struct {
	// error code
	Code int32 `json:"code,omitempty"`
	// error message
	Message string `json:"message,omitempty"`
	// error description
	Description string `json:"description,omitempty"`
}
