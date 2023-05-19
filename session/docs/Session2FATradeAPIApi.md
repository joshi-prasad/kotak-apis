# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/session/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Session2FAOneTimeTokenPost**](Session2FATradeAPIApi.md#Session2FAOneTimeTokenPost) | **Post** /session/2FA/oneTimeToken | Generate Final Session Token using One Time Token for Trade API subcribed clients

# **Session2FAOneTimeTokenPost**
> InlineResponse2003 Session2FAOneTimeTokenPost(ctx, oneTimeToken, consumerKey, ip, appId, optional)
Generate Final Session Token using One Time Token for Trade API subcribed clients

API to generate final session token for user based on One time token Generated using Login API

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oneTimeToken** | **string**|  | 
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 
 **optional** | ***Session2FATradeAPIApiSession2FAOneTimeTokenPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Session2FATradeAPIApiSession2FAOneTimeTokenPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of Model2FaOneTimeTokenBody**](Model2FaOneTimeTokenBody.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

