/*
 * Session-API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package session

type InlineResponse2003SuccessApiTokenRedirect struct {
	// Redirect host where client login is exected to be done..
	Host string `json:"host,omitempty"`
	// Redirect Host Port.
	Port float64 `json:"port,omitempty"`
}
