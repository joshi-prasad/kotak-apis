# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/margin/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarginCashbalanceHistoryGet**](MarginBasecapitalHistoryApi.md#MarginCashbalanceHistoryGet) | **Get** /margin/cashbalance-history | Base Capital History

# **MarginCashbalanceHistoryGet**
> []MarginDet MarginCashbalanceHistoryGet(ctx, consumerKey, sessionToken)
Base Capital History

Gives Transaction History of Basecapital i.e. Availalbe Cash Balance.

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

