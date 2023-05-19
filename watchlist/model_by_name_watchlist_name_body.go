/*
 * Watchlist-API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package watchlist

type ByNameWatchlistNameBody struct {
	// Unique token of security
	InstumentToken int32 `json:"instumentToken,omitempty"`
	// Blank in case not available else as recieved in Get Watchlist details
	TokenAddedTime string `json:"tokenAddedTime,omitempty"`
	// 0 if not available else as recieved in get Watchlist Details
	TokenAddedLtp float64 `json:"tokenAddedLtp,omitempty"`
}
