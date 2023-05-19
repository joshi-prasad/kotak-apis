# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioTransactionsScripwiseCommodityFyFromDateToDateInstrumentTokenGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseCommodityFyFromDateToDateInstrumentTokenGet) | **Get** /portfolio/transactions-scripwise/commodity/{fy}/{fromDate}/{toDate}/{instrumentToken} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseCommodityFyInstrumentTokenGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseCommodityFyInstrumentTokenGet) | **Get** /portfolio/transactions-scripwise/commodity/{fy}/{instrumentToken} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseCurrencyFyFromDateToDateInstrumentTokenGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseCurrencyFyFromDateToDateInstrumentTokenGet) | **Get** /portfolio/transactions-scripwise/currency/{fy}/{fromDate}/{toDate}/{instrumentToken} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseCurrencyFyInstrumentTokenGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseCurrencyFyInstrumentTokenGet) | **Get** /portfolio/transactions-scripwise/currency/{fy}/{instrumentToken} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseDerivativeFyFromDateToDateInstrumentTokenGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseDerivativeFyFromDateToDateInstrumentTokenGet) | **Get** /portfolio/transactions-scripwise/derivative/{fy}/{fromDate}/{toDate}/{instrumentToken} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseDerivativeFyInstrumentTokenGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseDerivativeFyInstrumentTokenGet) | **Get** /portfolio/transactions-scripwise/derivative/{fy}/{instrumentToken} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseEquityFyFromDateToDateInstrumentTokenGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseEquityFyFromDateToDateInstrumentTokenGet) | **Get** /portfolio/transactions-scripwise/equity/{fy}/{fromDate}/{toDate}/{instrumentToken} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseEquityFyInstrumentTokenGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseEquityFyInstrumentTokenGet) | **Get** /portfolio/transactions-scripwise/equity/{fy}/{instrumentToken} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseMutualfundFyFromDateToDateMfcodeGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseMutualfundFyFromDateToDateMfcodeGet) | **Get** /portfolio/transactions-scripwise/mutualfund/{fy}/{fromDate}/{toDate}/{mfcode} | Get Scripwie Transaction Statement
[**PortfolioTransactionsScripwiseMutualfundFyMfcodeGet**](StatementOfTransactionsScripwiseApi.md#PortfolioTransactionsScripwiseMutualfundFyMfcodeGet) | **Get** /portfolio/transactions-scripwise/mutualfund/{fy}/{mfcode} | Get Scripwie Transaction Statement

# **PortfolioTransactionsScripwiseCommodityFyFromDateToDateInstrumentTokenGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseCommodityFyFromDateToDateInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, instrumentToken, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
  **instrumentToken** | **int32**| unique token of instrument | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseCommodityFyFromDateToDateInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseCommodityFyFromDateToDateInstrumentTokenGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseCommodityFyInstrumentTokenGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseCommodityFyInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, instrumentToken, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **instrumentToken** | **int32**| unique token of instrument | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseCommodityFyInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseCommodityFyInstrumentTokenGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseCurrencyFyFromDateToDateInstrumentTokenGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseCurrencyFyFromDateToDateInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, instrumentToken, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
  **instrumentToken** | **int32**| unique token of instrument | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseCurrencyFyFromDateToDateInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseCurrencyFyFromDateToDateInstrumentTokenGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseCurrencyFyInstrumentTokenGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseCurrencyFyInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, instrumentToken, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **instrumentToken** | **int32**| unique token of instrument | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseCurrencyFyInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseCurrencyFyInstrumentTokenGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseDerivativeFyFromDateToDateInstrumentTokenGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseDerivativeFyFromDateToDateInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, instrumentToken, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
  **instrumentToken** | **int32**| unique token of instrument | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseDerivativeFyFromDateToDateInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseDerivativeFyFromDateToDateInstrumentTokenGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseDerivativeFyInstrumentTokenGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseDerivativeFyInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, instrumentToken, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **instrumentToken** | **int32**| unique token of instrument | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseDerivativeFyInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseDerivativeFyInstrumentTokenGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseEquityFyFromDateToDateInstrumentTokenGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseEquityFyFromDateToDateInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, instrumentToken, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
  **instrumentToken** | **int32**| unique token of instrument | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseEquityFyFromDateToDateInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseEquityFyFromDateToDateInstrumentTokenGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseEquityFyInstrumentTokenGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseEquityFyInstrumentTokenGet(ctx, consumerKey, sessionToken, fy, instrumentToken, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **instrumentToken** | **int32**| unique token of instrument | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseEquityFyInstrumentTokenGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseEquityFyInstrumentTokenGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseMutualfundFyFromDateToDateMfcodeGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseMutualfundFyFromDateToDateMfcodeGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, mfcode, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
  **mfcode** | **int32**| This is commconScripCode for MF scrip in output of Holdings API | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseMutualfundFyFromDateToDateMfcodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseMutualfundFyFromDateToDateMfcodeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsScripwiseMutualfundFyMfcodeGet**
> []TransactionsScripwise PortfolioTransactionsScripwiseMutualfundFyMfcodeGet(ctx, consumerKey, sessionToken, fy, mfcode, optional)
Get Scripwie Transaction Statement

Returns the full Transaction Statement of Client for given details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **mfcode** | **int32**| This is commconScripCode for MF scrip in output of Holdings API | 
 **optional** | ***StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseMutualfundFyMfcodeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsScripwiseApiPortfolioTransactionsScripwiseMutualfundFyMfcodeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]TransactionsScripwise**](transactions-scripwise.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

