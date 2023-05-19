# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/margin/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginSecurityCashGet**](MarginSecurityApi.md#MarginSecurityCashGet) | **Get** /margin/security/cash | Security Margin Cash
[**MarginSecurityMtfGet**](MarginSecurityApi.md#MarginSecurityMtfGet) | **Get** /margin/security/mtf | Security Margin MTF

# **MarginSecurityCashGet**
> []MarginDet MarginSecurityCashGet(ctx, consumerKey, sessionToken)
Security Margin Cash

Gives Security Margin for Cash Segment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**[]MarginDet**](marginDet.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginSecurityMtfGet**
> []MarginDet MarginSecurityMtfGet(ctx, consumerKey, sessionToken)
Security Margin MTF

Gives Security Margin for MTF.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 

### Return type

[**[]MarginDet**](marginDet.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

