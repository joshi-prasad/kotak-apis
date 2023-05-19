# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/margin/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCommoditiesExchangeGet**](MarginCommoditiesApi.md#MarginCommoditiesExchangeGet) | **Get** /margin/commodities/{exchange} | Exchange wise Commodities Margin
[**MarginCommoditiesGet**](MarginCommoditiesApi.md#MarginCommoditiesGet) | **Get** /margin/commodities | Commodities Margin

# **MarginCommoditiesExchangeGet**
> []MarginDet MarginCommoditiesExchangeGet(ctx, consumerKey, sessionToken, exchange)
Exchange wise Commodities Margin

Gives Exchagne wise Commodities Margin Details of a Client from RMS

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **exchange** | **string**| Possible Values  : MCX | 

### Return type

[**[]MarginDet**](marginDet.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MarginCommoditiesGet**
> []MarginDet MarginCommoditiesGet(ctx, consumerKey, sessionToken)
Commodities Margin

Gives Commodities Margin Details of a Client from RMS

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

