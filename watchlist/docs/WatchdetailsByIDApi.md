# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/watchlist/2.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatchdetailsByIDWatchlistIdInstrumentTokenDelete**](WatchdetailsByIDApi.md#WatchdetailsByIDWatchlistIdInstrumentTokenDelete) | **Delete** /watchdetails/byID/{watchlistId}/{instrumentToken} | Delete token by Watchist Id
[**WatchdetailsByIDWatchlistIdInstrumentTokenPost**](WatchdetailsByIDApi.md#WatchdetailsByIDWatchlistIdInstrumentTokenPost) | **Post** /watchdetails/byID/{watchlistId}/{instrumentToken} | Add Token to watchlist by watchlist id
[**WatchdetailsByIDWatchlistIdPut**](WatchdetailsByIDApi.md#WatchdetailsByIDWatchlistIdPut) | **Put** /watchdetails/byID/{watchlistId} | Synchronise Watchlist By watchlist Id

# **WatchdetailsByIDWatchlistIdInstrumentTokenDelete**
> InlineResponse2001 WatchdetailsByIDWatchlistIdInstrumentTokenDelete(ctx, userid, consumerKey, watchlistId, instrumentToken, sessionToken)
Delete token by Watchist Id

API to delete token from watchlist by watchlist Id for a UsreId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistId** | **int32**|  | 
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

# **WatchdetailsByIDWatchlistIdInstrumentTokenPost**
> InlineResponse2008 WatchdetailsByIDWatchlistIdInstrumentTokenPost(ctx, userid, consumerKey, watchlistId, instrumentToken, sessionToken)
Add Token to watchlist by watchlist id

API to add token token to a watchlist by Watchlist Id

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistId** | **int32**|  | 
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

# **WatchdetailsByIDWatchlistIdPut**
> InlineResponse2001 WatchdetailsByIDWatchlistIdPut(ctx, userid, consumerKey, watchlistId, sessionToken, optional)
Synchronise Watchlist By watchlist Id

API to save watchlist in same order seqence tokens arrya passed in request json payload by Watchlist ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **watchlistId** | **int32**|  | 
  **sessionToken** | **string**|  | 
 **optional** | ***WatchdetailsByIDApiWatchdetailsByIDWatchlistIdPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WatchdetailsByIDApiWatchdetailsByIDWatchlistIdPutOpts struct
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

