# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/watchlist/2.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WatchlistsByIDPut**](WatchlistsApi.md#WatchlistsByIDPut) | **Put** /watchlists/byID | Rearrange watchlist by Id
[**WatchlistsGet**](WatchlistsApi.md#WatchlistsGet) | **Get** /watchlists | List Watchlist

# **WatchlistsByIDPut**
> InlineResponse2001 WatchlistsByIDPut(ctx, userid, consumerKey, sessionToken, optional)
Rearrange watchlist by Id

API to Rearrange Watchlist by Name for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
 **optional** | ***WatchlistsApiWatchlistsByIDPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WatchlistsApiWatchlistsByIDPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of WatchlistsByIdBody**](WatchlistsByIdBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WatchlistsGet**
> InlineResponse200 WatchlistsGet(ctx, userid, consumerKey, sessionToken)
List Watchlist

API to List watchlist for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

