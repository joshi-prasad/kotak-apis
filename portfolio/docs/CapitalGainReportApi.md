# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioCapitalgainDetailsFyQuarterGet**](CapitalGainReportApi.md#PortfolioCapitalgainDetailsFyQuarterGet) | **Get** /portfolio/capitalgain/details/{fy}/{quarter} | Get Capital Gain Report by Details
[**PortfolioCapitalgainSummaryFyQuarterGet**](CapitalGainReportApi.md#PortfolioCapitalgainSummaryFyQuarterGet) | **Get** /portfolio/capitalgain/summary/{fy}/{quarter} | Get Capital Gain Report by summary

# **PortfolioCapitalgainDetailsFyQuarterGet**
> InlineResponse2003 PortfolioCapitalgainDetailsFyQuarterGet(ctx, consumerKey, sessionToken, fy, quarter, optional)
Get Capital Gain Report by Details

returns Capital Gain Report by Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **quarter** | **float64**| Quarter of the year- 0-ALL, 1-Q1, 2-Q2, 3-Q3, 4-Q4 | 
 **optional** | ***CapitalGainReportApiPortfolioCapitalgainDetailsFyQuarterGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CapitalGainReportApiPortfolioCapitalgainDetailsFyQuarterGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioCapitalgainSummaryFyQuarterGet**
> InlineResponse2002 PortfolioCapitalgainSummaryFyQuarterGet(ctx, consumerKey, sessionToken, fy, quarter, optional)
Get Capital Gain Report by summary

returns Capital Gain Report by summary

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **quarter** | **float64**| Quarter of the year- 0-ALL, 1-Q1, 2-Q2, 3-Q3, 4-Q4 | 
 **optional** | ***CapitalGainReportApiPortfolioCapitalgainSummaryFyQuarterGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CapitalGainReportApiPortfolioCapitalgainSummaryFyQuarterGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

