# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioPnlCommodityFyFromDateToDateGet**](RealisedProfitLossSegmentwiseApi.md#PortfolioPnlCommodityFyFromDateToDateGet) | **Get** /portfolio/pnl/commodity/{fy}/{fromDate}/{toDate} | Get Gain Loss Statement for Cash
[**PortfolioPnlCurrencyFyFromDateToDateGet**](RealisedProfitLossSegmentwiseApi.md#PortfolioPnlCurrencyFyFromDateToDateGet) | **Get** /portfolio/pnl/currency/{fy}/{fromDate}/{toDate} | Get Gain Loss Statement for Cash
[**PortfolioPnlDerivativeFyFromDateToDateGet**](RealisedProfitLossSegmentwiseApi.md#PortfolioPnlDerivativeFyFromDateToDateGet) | **Get** /portfolio/pnl/derivative/{fy}/{fromDate}/{toDate} | Get Gain Loss Statement for Cash
[**PortfolioPnlEquityFyFromDateToDateGet**](RealisedProfitLossSegmentwiseApi.md#PortfolioPnlEquityFyFromDateToDateGet) | **Get** /portfolio/pnl/equity/{fy}/{fromDate}/{toDate} | Get Gain Loss Statement for Cash
[**PortfolioPnlMutualfundFyFromDateToDateGet**](RealisedProfitLossSegmentwiseApi.md#PortfolioPnlMutualfundFyFromDateToDateGet) | **Get** /portfolio/pnl/mutualfund/{fy}/{fromDate}/{toDate} | Get Gain Loss Statement for Cash

# **PortfolioPnlCommodityFyFromDateToDateGet**
> []PnlCommodity PortfolioPnlCommodityFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, optional)
Get Gain Loss Statement for Cash

Returns Gain Loss Statement for Cash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***RealisedProfitLossSegmentwiseApiPortfolioPnlCommodityFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossSegmentwiseApiPortfolioPnlCommodityFyFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]PnlCommodity**](pnl-commodity.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioPnlCurrencyFyFromDateToDateGet**
> []PnlCurrency PortfolioPnlCurrencyFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, optional)
Get Gain Loss Statement for Cash

Returns Gain Loss Statement for Cash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***RealisedProfitLossSegmentwiseApiPortfolioPnlCurrencyFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossSegmentwiseApiPortfolioPnlCurrencyFyFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]PnlCurrency**](pnl-currency.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioPnlDerivativeFyFromDateToDateGet**
> []PnlDerivative PortfolioPnlDerivativeFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, optional)
Get Gain Loss Statement for Cash

Returns Gain Loss Statement for Cash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***RealisedProfitLossSegmentwiseApiPortfolioPnlDerivativeFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossSegmentwiseApiPortfolioPnlDerivativeFyFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]PnlDerivative**](pnl-derivative.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioPnlEquityFyFromDateToDateGet**
> []PnlEquity PortfolioPnlEquityFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, optional)
Get Gain Loss Statement for Cash

Returns Gain Loss Statement for Cash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***RealisedProfitLossSegmentwiseApiPortfolioPnlEquityFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossSegmentwiseApiPortfolioPnlEquityFyFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]PnlEquity**](pnl-equity.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioPnlMutualfundFyFromDateToDateGet**
> []PnlMutualfund PortfolioPnlMutualfundFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, optional)
Get Gain Loss Statement for Cash

Returns Gain Loss Statement for Cash

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***RealisedProfitLossSegmentwiseApiPortfolioPnlMutualfundFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossSegmentwiseApiPortfolioPnlMutualfundFyFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]PnlMutualfund**](pnl-mutualfund.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

