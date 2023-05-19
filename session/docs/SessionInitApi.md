# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/session/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionInitGet**](SessionInitApi.md#SessionInitGet) | **Get** /session/init | Initialise Session
[**SessionInitNicknameNicknameGet**](SessionInitApi.md#SessionInitNicknameNicknameGet) | **Get** /session/init/nickname/{nickname} | Initialise Session by nickname
[**SessionInitUseridUseridGet**](SessionInitApi.md#SessionInitUseridUseridGet) | **Get** /session/init/userid/{userid} | Initialise Session by UserId

# **SessionInitGet**
> InlineResponse200 SessionInitGet(ctx, userid, consumerKey, ip, appId)
Initialise Session

API to initiate trading session for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionInitNicknameNicknameGet**
> InlineResponse200 SessionInitNicknameNicknameGet(ctx, nickname, consumerKey, ip, appId)
Initialise Session by nickname

API to initiate trading session for a nickname

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **nickname** | **string**|  | 
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionInitUseridUseridGet**
> InlineResponse200 SessionInitUseridUseridGet(ctx, userid, consumerKey, ip, appId)
Initialise Session by UserId

API to initiate trading session for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userid** | **string**|  | 
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

