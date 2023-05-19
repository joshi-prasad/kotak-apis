# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioPnlCommodityFyFromDateToDateInstrumentTokenGet**](RealisedProfitLossScripwiseApi.md#PortfolioPnlCommodityFyFromDateToDateInstrumentTokenGet) | **Get** /portfolio/pnl/commodity/{fy}/{fromDate}/{toDate}/{instrumentToken} | Get Gain Loss Statement for Cash
[**PortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGet**](RealisedProfitLossScripwiseApi.md#PortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGet) | **Get** /portfolio/pnl/currency/{fy}/{fromDate}/{toDate}/{instrumentToken} | Get Gain Loss Statement for Cash
[**PortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGet**](RealisedProfitLossScripwiseApi.md#PortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGet) | **Get** /portfolio/pnl/derivative/{fy}/{fromDate}/{toDate}/{instrumentToken} | Get Gain Loss Statement for Cash
[**PortfolioPnlEquityFyFromDateToDateInstrumentTokenGet**](RealisedProfitLossScripwiseApi.md#PortfolioPnlEquityFyFromDateToDateInstrumentTokenGet) | **Get** /portfolio/pnl/equity/{fy}/{fromDate}/{toDate}/{instrumentToken} | Get Gain Loss Statement for Cash
[**PortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGet**](RealisedProfitLossScripwiseApi.md#PortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGet) | **Get** /portfolio/pnl/mutualfund/{fy}/{fromDate}/{toDate}/{instrumentToken} | Get Gain Loss Statement for Cash

# **PortfolioPnlCommodityFyFromDateToDateInstrumentTokenGet**
> []PnlCommodity PortfolioPnlCommodityFyFromDateToDateInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, instrumentToken, optional)
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
  **instrumentToken** | **string**| Unique token of Security | 
 **optional** | ***RealisedProfitLossScripwiseApiPortfolioPnlCommodityFyFromDateToDateInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossScripwiseApiPortfolioPnlCommodityFyFromDateToDateInstrumentTokenGetOpts struct
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

# **PortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGet**
> []PnlCurrency PortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, instrumentToken, optional)
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
  **instrumentToken** | **string**| Unique token of Security | 
 **optional** | ***RealisedProfitLossScripwiseApiPortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossScripwiseApiPortfolioPnlCurrencyFyFromDateToDateInstrumentTokenGetOpts struct
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

# **PortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGet**
> []PnlDerivative PortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, instrumentToken, optional)
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
  **instrumentToken** | **string**| Unique token of Security | 
 **optional** | ***RealisedProfitLossScripwiseApiPortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossScripwiseApiPortfolioPnlDerivativeFyFromDateToDateInstrumentTokenGetOpts struct
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

# **PortfolioPnlEquityFyFromDateToDateInstrumentTokenGet**
> []PnlEquity PortfolioPnlEquityFyFromDateToDateInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, instrumentToken, optional)
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
  **instrumentToken** | **string**| Unique token of Security | 
 **optional** | ***RealisedProfitLossScripwiseApiPortfolioPnlEquityFyFromDateToDateInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossScripwiseApiPortfolioPnlEquityFyFromDateToDateInstrumentTokenGetOpts struct
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

# **PortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGet**
> []PnlMutualfund PortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, instrumentToken, optional)
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
  **instrumentToken** | **string**| Unique token of Security | 
 **optional** | ***RealisedProfitLossScripwiseApiPortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossScripwiseApiPortfolioPnlMutualfundFyFromDateToDateInstrumentTokenGetOpts struct
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

