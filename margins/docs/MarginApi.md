# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/margin/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginGet**](MarginApi.md#MarginGet) | **Get** /margin | Complete Margin
[**MarginMiniGet**](MarginApi.md#MarginMiniGet) | **Get** /margin-mini | Margin Minified Data

# **MarginGet**
> []MarginDet MarginGet(ctx, consumerKey, sessionToken)
Complete Margin

Gives complete Margin Details of a Client from RMS

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

# **MarginMiniGet**
> []MarginDet MarginMiniGet(ctx, consumerKey, sessionToken)
Margin Minified Data

Gives Mini version of Margin Details of a Client from RMS

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

