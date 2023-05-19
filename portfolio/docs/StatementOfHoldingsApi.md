# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioHoldingsAllDateGet**](StatementOfHoldingsApi.md#PortfolioHoldingsAllDateGet) | **Get** /portfolio/holdings/all/{date} | Get user holdings for a date
[**PortfolioHoldingsAllGet**](StatementOfHoldingsApi.md#PortfolioHoldingsAllGet) | **Get** /portfolio/holdings/all | Get user holdings
[**PortfolioHoldingsDateGet**](StatementOfHoldingsApi.md#PortfolioHoldingsDateGet) | **Get** /portfolio/holdings/{date} | Get user holdings for a date

# **PortfolioHoldingsAllDateGet**
> []Holdings PortfolioHoldingsAllDateGet(ctx, consumerKey, sessionToken, date, optional)
Get user holdings for a date

Returns the full DP holdings of the user on a specific date inluding of Sold holdings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **date** | **string**| Date of specific holdings to fetch Format - YYYY-MM-DD | 
 **optional** | ***StatementOfHoldingsApiPortfolioHoldingsAllDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfHoldingsApiPortfolioHoldingsAllDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]Holdings**](holdings.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioHoldingsAllGet**
> []Holdings PortfolioHoldingsAllGet(ctx, consumerKey, sessionToken, optional)
Get user holdings

Returns the full DP holdings of the user including Sold Holdings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
 **optional** | ***StatementOfHoldingsApiPortfolioHoldingsAllGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfHoldingsApiPortfolioHoldingsAllGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]Holdings**](holdings.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioHoldingsDateGet**
> []Holdings PortfolioHoldingsDateGet(ctx, consumerKey, sessionToken, date, optional)
Get user holdings for a date

Returns the full DP holdings of the user on a specific date exlusive of Sold holdings

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **date** | **string**| Date of specific holdings to fetch Format - YYYY-MM-DD | 
 **optional** | ***StatementOfHoldingsApiPortfolioHoldingsDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfHoldingsApiPortfolioHoldingsDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]Holdings**](holdings.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

