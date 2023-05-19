# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/watchlist/2.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatchlistsByIDWatchlistIdDelete**](WatchlistsByIDApi.md#WatchlistsByIDWatchlistIdDelete) | **Delete** /watchlists/byID/{watchlistId} | Delete Watchlist by Watchlist Id
[**WatchlistsByIDWatchlistIdGet**](WatchlistsByIDApi.md#WatchlistsByIDWatchlistIdGet) | **Get** /watchlists/byID/{watchlistId} | Get Watchlist by Watchlist Id
[**WatchlistsByIDWatchlistIdNewWatchlistNamePut**](WatchlistsByIDApi.md#WatchlistsByIDWatchlistIdNewWatchlistNamePut) | **Put** /watchlists/byID/{watchlistId}/{newWatchlistName} | Rename watchlist by Id

# **WatchlistsByIDWatchlistIdDelete**
> InlineResponse2003 WatchlistsByIDWatchlistIdDelete(ctx, userid, consumerKey, watchlistId, sessionToken)
Delete Watchlist by Watchlist Id

API to delete watchlist by watchlist Id for a user id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistId** | **int32**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchlistsByIDWatchlistIdGet**
> InlineResponse2002 WatchlistsByIDWatchlistIdGet(ctx, userid, consumerKey, watchlistId, sessionToken)
Get Watchlist by Watchlist Id

API To get Watchlist by watchlist ID for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistId** | **int32**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchlistsByIDWatchlistIdNewWatchlistNamePut**
> InlineResponse2004 WatchlistsByIDWatchlistIdNewWatchlistNamePut(ctx, userid, consumerKey, watchlistId, newWatchlistName, sessionToken)
Rename watchlist by Id

API to rename Watchlist by Name for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistId** | **int32**|  | 
  **newWatchlistName** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

