# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioTransactionsScripwiseCommoditySsymbolFyFromDateToDateSsymbolGet**](StatementOfTransactionsScripwiseForSSymbolApi.md#PortfolioTransactionsScripwiseCommoditySsymbolFyFromDateToDateSsymbolGet) | **Get** /portfolio/transactions-scripwise/commodity/ssymbol/{fy}/{fromDate}/{toDate}/{ssymbol} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseCommoditySsymbolFySsymbolGet**](StatementOfTransactionsScripwiseForSSymbolApi.md#PortfolioTransactionsScripwiseCommoditySsymbolFySsymbolGet) | **Get** /portfolio/transactions-scripwise/commodity/ssymbol/{fy}/{ssymbol} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseCurrencySsymbolFyFromDateToDateSsymbolGet**](StatementOfTransactionsScripwiseForSSymbolApi.md#PortfolioTransactionsScripwiseCurrencySsymbolFyFromDateToDateSsymbolGet) | **Get** /portfolio/transactions-scripwise/currency/ssymbol/{fy}/{fromDate}/{toDate}/{ssymbol} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseCurrencySsymbolFySsymbolGet**](StatementOfTransactionsScripwiseForSSymbolApi.md#PortfolioTransactionsScripwiseCurrencySsymbolFySsymbolGet) | **Get** /portfolio/transactions-scripwise/currency/ssymbol/{fy}/{ssymbol} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseDerivativeSsymbolFyFromDateToDateSsymbolGet**](StatementOfTransactionsScripwiseForSSymbolApi.md#PortfolioTransactionsScripwiseDerivativeSsymbolFyFromDateToDateSsymbolGet) | **Get** /portfolio/transactions-scripwise/derivative/ssymbol/{fy}/{fromDate}/{toDate}/{ssymbol} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseDerivativeSsymbolFySsymbolGet**](StatementOfTransactionsScripwiseForSSymbolApi.md#PortfolioTransactionsScripwiseDerivativeSsymbolFySsymbolGet) | **Get** /portfolio/transactions-scripwise/derivative/ssymbol/{fy}/{ssymbol} | Get Scripwie Transaction Statement

# **PortfolioTransactionsScripwiseCommoditySsymbolFyFromDateToDateSsymbolGet**
> []TransactionsScripwiseSsymbol PortfolioTransactionsScripwiseCommoditySsymbolFyFromDateToDateSsymbolGet(ctx, consumerKey, sessionToken, fy, ssymbol, fromDate, toDate, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **ssymbol** | **string**| unique token of ssymbol | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseCommoditySsymbolFyFromDateToDateSsymbolGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseCommoditySsymbolFyFromDateToDateSsymbolGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwiseSsymbol**](transactions-scripwise-ssymbol.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseCommoditySsymbolFySsymbolGet**
> []TransactionsScripwiseSsymbol PortfolioTransactionsScripwiseCommoditySsymbolFySsymbolGet(ctx, consumerKey, sessionToken, fy, ssymbol, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **ssymbol** | **string**| unique token of ssymbol | 
 **optional** | ***StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseCommoditySsymbolFySsymbolGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseCommoditySsymbolFySsymbolGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwiseSsymbol**](transactions-scripwise-ssymbol.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseCurrencySsymbolFyFromDateToDateSsymbolGet**
> []TransactionsScripwiseSsymbol PortfolioTransactionsScripwiseCurrencySsymbolFyFromDateToDateSsymbolGet(ctx, consumerKey, sessionToken, fy, ssymbol, fromDate, toDate, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **ssymbol** | **string**| unique token of ssymbol | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseCurrencySsymbolFyFromDateToDateSsymbolGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseCurrencySsymbolFyFromDateToDateSsymbolGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwiseSsymbol**](transactions-scripwise-ssymbol.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseCurrencySsymbolFySsymbolGet**
> []TransactionsScripwiseSsymbol PortfolioTransactionsScripwiseCurrencySsymbolFySsymbolGet(ctx, consumerKey, sessionToken, fy, ssymbol, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **ssymbol** | **string**| unique token of ssymbol | 
 **optional** | ***StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseCurrencySsymbolFySsymbolGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseCurrencySsymbolFySsymbolGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwiseSsymbol**](transactions-scripwise-ssymbol.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseDerivativeSsymbolFyFromDateToDateSsymbolGet**
> []TransactionsScripwiseSsymbol PortfolioTransactionsScripwiseDerivativeSsymbolFyFromDateToDateSsymbolGet(ctx, consumerKey, sessionToken, fy, ssymbol, fromDate, toDate, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **ssymbol** | **string**| unique token of ssymbol | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseDerivativeSsymbolFyFromDateToDateSsymbolGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseDerivativeSsymbolFyFromDateToDateSsymbolGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwiseSsymbol**](transactions-scripwise-ssymbol.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseDerivativeSsymbolFySsymbolGet**
> []TransactionsScripwiseSsymbol PortfolioTransactionsScripwiseDerivativeSsymbolFySsymbolGet(ctx, consumerKey, sessionToken, fy, ssymbol, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **ssymbol** | **string**| unique token of ssymbol | 
 **optional** | ***StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseDerivativeSsymbolFySsymbolGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseForSSymbolApiPortfolioTransactionsScripwiseDerivativeSsymbolFySsymbolGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwiseSsymbol**](transactions-scripwise-ssymbol.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

