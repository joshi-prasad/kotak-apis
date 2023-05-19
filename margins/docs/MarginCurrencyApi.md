# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/margin/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCurrencyExchangeGet**](MarginCurrencyApi.md#MarginCurrencyExchangeGet) | **Get** /margin/currency/{exchange} | Exchange wise Currency Margin
[**MarginCurrencyGet**](MarginCurrencyApi.md#MarginCurrencyGet) | **Get** /margin/currency | Currency Margin

# **MarginCurrencyExchangeGet**
> []MarginDet MarginCurrencyExchangeGet(ctx, consumerKey, sessionToken, exchange)
Exchange wise Currency Margin

Gives Exchagne wise Currency Margin Details of a Client from RMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **exchange** | **string**| Possible Values  : NSE, MCX | 

### Return type

[**[]MarginDet**](marginDet.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginCurrencyGet**
> []MarginDet MarginCurrencyGet(ctx, consumerKey, sessionToken)
Currency Margin

Gives Currency Margin Details of a Client from RMS

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

