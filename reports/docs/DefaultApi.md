# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/reports/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrdersGet**](DefaultApi.md#OrdersGet) | **Get** /orders | Get order report
[**OrdersOrderIdGet**](DefaultApi.md#OrdersOrderIdGet) | **Get** /orders/{orderId} | Get order report by orderId
[**OrdersOrderIdIsFNOGet**](DefaultApi.md#OrdersOrderIdIsFNOGet) | **Get** /orders/{orderId}/{isFNO} | Get order report by orderId
[**TradesGet**](DefaultApi.md#TradesGet) | **Get** /trades | Get trade report
[**TradesOrderIdGet**](DefaultApi.md#TradesOrderIdGet) | **Get** /trades/{orderId} | Get trade report by orderId
[**TradesOrderIdIsFNOGet**](DefaultApi.md#TradesOrderIdIsFNOGet) | **Get** /trades/{orderId}/{isFNO} | Get trade report by orderId

# **OrdersGet**
> []Orders OrdersGet(ctx, consumerKey, sessionToken)
Get order report

Returns the full order report for a client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**[]Orders**](Orders.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrdersOrderIdGet**
> InlineResponse200 OrdersOrderIdGet(ctx, orderId, consumerKey, sessionToken)
Get order report by orderId

Returns the specific order report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**|  | 
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

# **OrdersOrderIdIsFNOGet**
> InlineResponse200 OrdersOrderIdIsFNOGet(ctx, orderId, consumerKey, sessionToken, isFNO)
Get order report by orderId

Returns the specific order report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**|  | 
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **isFNO** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradesGet**
> InlineResponse2001 TradesGet(ctx, optional)
Get trade report

Returns the full trade report

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiTradesGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiTradesGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consumerKey** | **optional.String**| Unique ID for your application | 
 **sessionToken** | **optional.String**| Session ID for your application | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradesOrderIdGet**
> InlineResponse200 TradesOrderIdGet(ctx, orderId, consumerKey, sessionToken)
Get trade report by orderId

Returns the trade report for a orderId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**|  | 
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TradesOrderIdIsFNOGet**
> InlineResponse200 TradesOrderIdIsFNOGet(ctx, orderId, consumerKey, sessionToken, isFNO)
Get trade report by orderId

Returns the trade report for a orderId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**|  | 
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **isFNO** | **string**| Session ID for your application | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

