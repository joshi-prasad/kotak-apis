# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/session/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Session2FAAccesscodeGet**](Session2FAApi.md#Session2FAAccesscodeGet) | **Get** /session/2FA/accesscode | Generate Access Code
[**Session2FAAccesscodePost**](Session2FAApi.md#Session2FAAccesscodePost) | **Post** /session/2FA/accesscode | Generate final Session Token
[**Session2FABiometricPost**](Session2FAApi.md#Session2FABiometricPost) | **Post** /session/2FA/biometric | Generate final Session Token
[**Session2FAMpinPost**](Session2FAApi.md#Session2FAMpinPost) | **Post** /session/2FA/mpin | Generate final Session Token

# **Session2FAAccesscodeGet**
> InlineResponse2001 Session2FAAccesscodeGet(ctx, oneTimeToken, consumerKey, ip, appId, userid)
Generate Access Code

API To generate access code for a UserId

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oneTimeToken** | **string**|  | 
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 
  **userid** | **string**|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Session2FAAccesscodePost**
> InlineResponse2002 Session2FAAccesscodePost(ctx, oneTimeToken, consumerKey, ip, appId, optional)
Generate final Session Token

API to generate final session for user based on one time token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oneTimeToken** | **string**|  | 
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 
 **optional** | ***Session2FAApiSession2FAAccesscodePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Session2FAApiSession2FAAccesscodePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of Model2FaAccesscodeBody**](Model2FaAccesscodeBody.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Session2FABiometricPost**
> InlineResponse2002 Session2FABiometricPost(ctx, oneTimeToken, consumerKey, ip, appId, optional)
Generate final Session Token

API to generate final session for user based on biometric

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oneTimeToken** | **string**|  | 
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 
 **optional** | ***Session2FAApiSession2FABiometricPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Session2FAApiSession2FABiometricPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of Model2FaBiometricBody**](Model2FaBiometricBody.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Session2FAMpinPost**
> InlineResponse2002 Session2FAMpinPost(ctx, oneTimeToken, consumerKey, ip, appId, optional)
Generate final Session Token

API to generate final session for user based on mpin

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **oneTimeToken** | **string**|  | 
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 
 **optional** | ***Session2FAApiSession2FAMpinPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a Session2FAApiSession2FAMpinPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **body** | [**optional.Interface of Model2FaMpinBody**](Model2FaMpinBody.md)|  | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

