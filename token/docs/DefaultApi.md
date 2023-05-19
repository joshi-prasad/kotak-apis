# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/oauth/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RevokePost**](DefaultApi.md#RevokePost) | **Post** /revoke | 
[**TokenPost**](DefaultApi.md#TokenPost) | **Post** /token | 

# **RevokePost**
> RevokePost(ctx, token, tokenTypeHint, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **token** | **string**|  | 
  **tokenTypeHint** | **string**|  | 
  **authorization** | **string**| Basic Base64(consumer-key:consumer-secret) | 

### Return type

 (empty response body)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokenPost**
> TokenPost(ctx, grantType, username, password, refreshToken, authorization)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **grantType** | **string**|  | 
  **username** | **string**|  | 
  **password** | **string**|  | 
  **refreshToken** | **string**|  | 
  **authorization** | **string**| Basic Base64(consumer-key:consumer-secret) | 

### Return type

 (empty response body)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

