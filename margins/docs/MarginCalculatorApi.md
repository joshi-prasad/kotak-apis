# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/margin/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginRequiredPost**](MarginCalculatorApi.md#MarginRequiredPost) | **Post** /margin/required | Get Margin Required for an order by amount or quantity.

# **MarginRequiredPost**
> []MarginDet MarginRequiredPost(ctx, consumerKey, sessionToken, optional)
Get Margin Required for an order by amount or quantity.

Returns margin required for Equity, Super Multiple & MTF Order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
 **optional** | ***MarginCalculatorApiMarginRequiredPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MarginCalculatorApiMarginRequiredPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MarginRequiredBody**](MarginRequiredBody.md)|  | 

### Return type

[**[]MarginDet**](marginDet.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

