# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioMonthlypnlFyFromDateToDateGet**](MonthlyPnlForPortfolioApi.md#PortfolioMonthlypnlFyFromDateToDateGet) | **Get** /portfolio/monthlypnl/{fy}/{fromDate}/{toDate} | Monthly for portfolio

# **PortfolioMonthlypnlFyFromDateToDateGet**
> []InlineResponse2007 PortfolioMonthlypnlFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, optional)
Monthly for portfolio

Monthly for portfolio

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***MonthlyPnlForPortfolioApiPortfolioMonthlypnlFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MonthlyPnlForPortfolioApiPortfolioMonthlypnlFyFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]InlineResponse2007**](inline_response_200_7.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

