/*
 * Watchlist-API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package watchlist

type WatchlistsbyIdSuccess struct {
	// Unique Watchlist Id
	WatchlistId float64 `json:"watchlistId,omitempty"`
	// Name of the Wathchlist
	WatchlistName string `json:"watchlistName,omitempty"`
}
