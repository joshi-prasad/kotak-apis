# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioAsOnDateGet**](AsOnDateApi.md#PortfolioAsOnDateGet) | **Get** /portfolio/asOnDate | Get As On Date

# **PortfolioAsOnDateGet**
> InlineResponse200 PortfolioAsOnDateGet(ctx, consumerKey, sessionToken, optional)
Get As On Date

Returns the date As on which data is available in system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
 **optional** | ***AsOnDateApiPortfolioAsOnDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsOnDateApiPortfolioAsOnDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

