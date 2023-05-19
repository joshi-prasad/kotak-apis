# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/orders/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderSorOrderIdDelete**](SmartOrderRoutingApi.md#OrderSorOrderIdDelete) | **Delete** /order/sor/{orderId} | Cancel an order
[**OrderSorPost**](SmartOrderRoutingApi.md#OrderSorPost) | **Post** /order/sor | Place a New order
[**OrderSorPut**](SmartOrderRoutingApi.md#OrderSorPut) | **Put** /order/sor | Modify an existing order

# **OrderSorOrderIdDelete**
> InlineResponse2001 OrderSorOrderIdDelete(ctx, consumerKey, sessionToken, orderId)
Cancel an order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **orderId** | **int64**| Order ID to cancel. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSorPost**
> InlineResponse2001 OrderSorPost(ctx, body, consumerKey, sessionToken)
Place a New order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderSorBody1**](OrderSorBody1.md)|  | 
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID Generated with successful login. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OrderSorPut**
> InlineResponse2001 OrderSorPut(ctx, body, consumerKey, sessionToken)
Modify an existing order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderSorBody**](OrderSorBody.md)|  | 
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

