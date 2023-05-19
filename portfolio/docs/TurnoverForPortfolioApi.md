# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioTurnoverFyFromDateToDateGet**](TurnoverForPortfolioApi.md#PortfolioTurnoverFyFromDateToDateGet) | **Get** /portfolio/turnover/{fy}/{fromDate}/{toDate} | turnover for portfolio

# **PortfolioTurnoverFyFromDateToDateGet**
> []InlineResponse2006 PortfolioTurnoverFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, optional)
turnover for portfolio

turnover for portfolio

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***TurnoverForPortfolioApiPortfolioTurnoverFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TurnoverForPortfolioApiPortfolioTurnoverFyFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]InlineResponse2006**](inline_response_200_6.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

