# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/watchlist/2.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatchlistsByNameWatchlistNameDelete**](WatchlistsByNameApi.md#WatchlistsByNameWatchlistNameDelete) | **Delete** /watchlists/byName/{watchlistName} | Delete watchlist by NAme
[**WatchlistsByNameWatchlistNameGet**](WatchlistsByNameApi.md#WatchlistsByNameWatchlistNameGet) | **Get** /watchlists/byName/{watchlistName} | Get Watchlist by NAme
[**WatchlistsByNameWatchlistNameNewWatchlistNamePut**](WatchlistsByNameApi.md#WatchlistsByNameWatchlistNameNewWatchlistNamePut) | **Put** /watchlists/byName/{watchlistName}/{newWatchlistName} | Rename watchlist by watchlist name.
[**WatchlistsByNameWatchlistNamePost**](WatchlistsByNameApi.md#WatchlistsByNameWatchlistNamePost) | **Post** /watchlists/byName/{watchlistName} | Create Watchlist

# **WatchlistsByNameWatchlistNameDelete**
> InlineResponse2006 WatchlistsByNameWatchlistNameDelete(ctx, userid, consumerKey, watchlistName, sessionToken)
Delete watchlist by NAme

API to delete watchlist by watchlist Name for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistName** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchlistsByNameWatchlistNameGet**
> InlineResponse2002 WatchlistsByNameWatchlistNameGet(ctx, userid, consumerKey, watchlistName, sessionToken)
Get Watchlist by NAme

API To get watchlist by Name for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistName** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchlistsByNameWatchlistNameNewWatchlistNamePut**
> InlineResponse2007 WatchlistsByNameWatchlistNameNewWatchlistNamePut(ctx, userid, consumerKey, watchlistName, newWatchlistName, sessionToken)
Rename watchlist by watchlist name.

API To rename watchlist by Watchlist Name for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistName** | **string**|  | 
  **newWatchlistName** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchlistsByNameWatchlistNamePost**
> InlineResponse2005 WatchlistsByNameWatchlistNamePost(ctx, userid, consumerKey, watchlistName, sessionToken)
Create Watchlist

API To create new watchlist for UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistName** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

