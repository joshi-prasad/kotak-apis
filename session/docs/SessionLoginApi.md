# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/session/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SessionLoginNicknamePost**](SessionLoginApi.md#SessionLoginNicknamePost) | **Post** /session/login/nickname | Login using nickname
[**SessionLoginUseridPost**](SessionLoginApi.md#SessionLoginUseridPost) | **Post** /session/login/userid | Login using Userid
[**SessionLogoutDelete**](SessionLoginApi.md#SessionLogoutDelete) | **Delete** /session/logout | Invalidate Session Token

# **SessionLoginNicknamePost**
> InlineResponse2001 SessionLoginNicknamePost(ctx, consumerKey, ip, appId, optional)
Login using nickname

Authenticate nickname serid and password to gnerrated one time token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 
 **optional** | ***SessionLoginApiSessionLoginNicknamePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SessionLoginApiSessionLoginNicknamePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of LoginNicknameBody**](LoginNicknameBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionLoginUseridPost**
> InlineResponse2001 SessionLoginUseridPost(ctx, consumerKey, ip, appId, optional)
Login using Userid

Authenticate userid and password to gnerrated one time token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 
 **optional** | ***SessionLoginApiSessionLoginUseridPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SessionLoginApiSessionLoginUseridPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of LoginUseridBody**](LoginUseridBody.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SessionLogoutDelete**
> InlineResponse2004 SessionLogoutDelete(ctx, sessionToken, consumerKey, ip, appId, userid)
Invalidate Session Token

API to invalidate final session for user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **sessionToken** | **string**|  | 
  **consumerKey** | **string**|  | 
  **ip** | **string**|  | 
  **appId** | **string**|  | 
  **userid** | **string**|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

