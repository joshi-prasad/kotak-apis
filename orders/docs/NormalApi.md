# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/orders/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrderNormalOrderIdDelete**](NormalApi.md#OrderNormalOrderIdDelete) | **Delete** /order/normal/{orderId} | Cancel a Normal order
[**OrderNormalPost**](NormalApi.md#OrderNormalPost) | **Post** /order/normal | Place a New order
[**OrderNormalPut**](NormalApi.md#OrderNormalPut) | **Put** /order/normal | Modify an existing order

# **OrderNormalOrderIdDelete**
> InlineResponse200 OrderNormalOrderIdDelete(ctx, consumerKey, sessionToken, orderId)
Cancel a Normal order

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

# **OrderNormalPost**
> InlineResponse200 OrderNormalPost(ctx, body, consumerKey, sessionToken)
Place a New order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderNormalBody1**](OrderNormalBody1.md)|  | 
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

# **OrderNormalPut**
> InlineResponse200 OrderNormalPut(ctx, body, consumerKey, sessionToken)
Modify an existing order

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**OrderNormalBody**](OrderNormalBody.md)|  | 
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

