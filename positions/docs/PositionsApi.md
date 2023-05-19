# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/positions/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PositionsGet**](PositionsApi.md#PositionsGet) | **Get** /positions | Get&#x27;s raw position from Trading Engine.
[**PositionsOpenGet**](PositionsApi.md#PositionsOpenGet) | **Get** /positions/open | Get&#x27;s Open position.
[**PositionsStocksGet**](PositionsApi.md#PositionsStocksGet) | **Get** /positions/stocks | Get&#x27;s Sell from Existing stocks.
[**PositionsTodaysGet**](PositionsApi.md#PositionsTodaysGet) | **Get** /positions/todays | Get&#x27;s Todays position.

# **PositionsGet**
> []Positions PositionsGet(ctx, consumerKey, sessionToken)
Get's raw position from Trading Engine.

Snapshot of Position data for a client available in System.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**[]Positions**](positions.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PositionsOpenGet**
> []Open PositionsOpenGet(ctx, consumerKey, sessionToken)
Get's Open position.

Gets Open Position of a client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**[]Open**](open.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PositionsStocksGet**
> []Stocks PositionsStocksGet(ctx, consumerKey, sessionToken)
Get's Sell from Existing stocks.

Gets Stocks position eligible for sell from existing for a client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**[]Stocks**](stocks.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PositionsTodaysGet**
> []Todays PositionsTodaysGet(ctx, consumerKey, sessionToken)
Get's Todays position.

Gets Today's Position of a client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**[]Todays**](todays.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

