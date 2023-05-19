# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/watchlist/2.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatchdetailsByNameWatchlistNameInstrumentTokenDelete**](WatchdetailsByNameApi.md#WatchdetailsByNameWatchlistNameInstrumentTokenDelete) | **Delete** /watchdetails/byName/{watchlistName}/{instrumentToken} | Delete Token by watchlist Name
[**WatchdetailsByNameWatchlistNameInstrumentTokenPost**](WatchdetailsByNameApi.md#WatchdetailsByNameWatchlistNameInstrumentTokenPost) | **Post** /watchdetails/byName/{watchlistName}/{instrumentToken} | Add token by watchlist Name
[**WatchdetailsByNameWatchlistNamePut**](WatchdetailsByNameApi.md#WatchdetailsByNameWatchlistNamePut) | **Put** /watchdetails/byName/{watchlistName} | Synchronise Watchlist by Watchlist Name

# **WatchdetailsByNameWatchlistNameInstrumentTokenDelete**
> InlineResponse2001 WatchdetailsByNameWatchlistNameInstrumentTokenDelete(ctx, userid, consumerKey, watchlistName, instrumentToken, sessionToken)
Delete Token by watchlist Name

API to Delete token by Watchlist Name for a userId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistName** | **string**|  | 
  **instrumentToken** | **int32**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchdetailsByNameWatchlistNameInstrumentTokenPost**
> InlineResponse2008 WatchdetailsByNameWatchlistNameInstrumentTokenPost(ctx, userid, consumerKey, watchlistName, instrumentToken, sessionToken)
Add token by watchlist Name

API to add token to watchlist by Watchlist Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistName** | **string**|  | 
  **instrumentToken** | **int32**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchdetailsByNameWatchlistNamePut**
> InlineResponse2001 WatchdetailsByNameWatchlistNamePut(ctx, userid, consumerKey, watchlistName, sessionToken, optional)
Synchronise Watchlist by Watchlist Name

API to save watchlist in same order seqence tokens arrya passed in request json payload by Watchlist Name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistName** | **string**|  | 
  **sessionToken** | **string**|  | 
 **optional** | ***WatchdetailsByNameApiWatchdetailsByNameWatchlistNamePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WatchdetailsByNameApiWatchdetailsByNameWatchlistNamePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of []ByNameWatchlistNameBody**](byName_watchlistName_body.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

