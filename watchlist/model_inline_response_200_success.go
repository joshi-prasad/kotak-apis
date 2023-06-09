/*
 * Watchlist-API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package watchlist

type InlineResponse200Success struct {
	// Unique Security Token.
	InstrumentCount float64 `json:"instrumentCount,omitempty"`
	// Last Access Flag
	LastAccess float64 `json:"lastAccess,omitempty"`
	// Unique Id of watchlist.
	WatchlistId float64 `json:"watchlistId,omitempty"`
	// Name of the watchlist.
	WatchlistName string `json:"watchlistName,omitempty"`
}
