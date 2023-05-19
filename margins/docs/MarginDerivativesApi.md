# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/margin/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginDerivativesExchangeGet**](MarginDerivativesApi.md#MarginDerivativesExchangeGet) | **Get** /margin/derivatives/{exchange} | Exchange wise Derivatives Margin
[**MarginDerivativesGet**](MarginDerivativesApi.md#MarginDerivativesGet) | **Get** /margin/derivatives | Derivatives Margin

# **MarginDerivativesExchangeGet**
> []MarginDet MarginDerivativesExchangeGet(ctx, consumerKey, sessionToken, exchange)
Exchange wise Derivatives Margin

Gives Exchagne wise Derivatives Margin Details of a Client from RMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **exchange** | **string**| Possible Values  : NSE, BSE | 

### Return type

[**[]MarginDet**](marginDet.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginDerivativesGet**
> []MarginDet MarginDerivativesGet(ctx, consumerKey, sessionToken)
Derivatives Margin

Gives Derivatives Margin Details of a Client from RMS

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

