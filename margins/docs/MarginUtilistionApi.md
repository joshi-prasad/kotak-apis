# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/margin/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginUtilisationCashGet**](MarginUtilistionApi.md#MarginUtilisationCashGet) | **Get** /margin/utilisation/cash | Margin Utilisaction for Cash
[**MarginUtilisationMtfGet**](MarginUtilistionApi.md#MarginUtilisationMtfGet) | **Get** /margin/utilisation/mtf | Margin Utilisaction for MTF

# **MarginUtilisationCashGet**
> []MarginDet MarginUtilisationCashGet(ctx, consumerKey, sessionToken)
Margin Utilisaction for Cash

Gives Margin Utilisation for Cash Segment

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

# **MarginUtilisationMtfGet**
> []MarginDet MarginUtilisationMtfGet(ctx, consumerKey, sessionToken)
Margin Utilisaction for MTF

Gives Margin Utilisation for MTF

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

