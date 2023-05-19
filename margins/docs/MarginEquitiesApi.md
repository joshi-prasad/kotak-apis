# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/margin/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginEquitiesExchangeGet**](MarginEquitiesApi.md#MarginEquitiesExchangeGet) | **Get** /margin/equities/{exchange} | Exchange wise Equities Margin
[**MarginEquitiesGet**](MarginEquitiesApi.md#MarginEquitiesGet) | **Get** /margin/equities | Equity Margin

# **MarginEquitiesExchangeGet**
> []MarginDet MarginEquitiesExchangeGet(ctx, consumerKey, sessionToken, exchange)
Exchange wise Equities Margin

Gives Exchagne wise Equities Margin Details of a Client from RMS

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

# **MarginEquitiesGet**
> []MarginDet MarginEquitiesGet(ctx, consumerKey, sessionToken)
Equity Margin

Gives complete Equity Margin Details of a Client from RMS

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

