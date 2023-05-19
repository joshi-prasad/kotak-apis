# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/orders/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderSupermultipleOrderIdDelete**](SuperMultipleApi.md#OrderSupermultipleOrderIdDelete) | **Delete** /order/supermultiple/{orderId} | Cancel an order
[**OrderSupermultiplePost**](SuperMultipleApi.md#OrderSupermultiplePost) | **Post** /order/supermultiple | Place a New order
[**OrderSupermultiplePut**](SuperMultipleApi.md#OrderSupermultiplePut) | **Put** /order/supermultiple | Modify an existing order

# **OrderSupermultipleOrderIdDelete**
> InlineResponse200 OrderSupermultipleOrderIdDelete(ctx, consumerKey, sessionToken, orderId)
Cancel an order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **orderId** | **int64**| Order ID to cancel. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSupermultiplePost**
> InlineResponse200 OrderSupermultiplePost(ctx, body, consumerKey, sessionToken)
Place a New order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderSupermultipleBody1**](OrderSupermultipleBody1.md)|  | 
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID Generated with successful login. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSupermultiplePut**
> InlineResponse200 OrderSupermultiplePut(ctx, body, consumerKey, sessionToken)
Modify an existing order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderSupermultipleBody**](OrderSupermultipleBody.md)|  | 
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

