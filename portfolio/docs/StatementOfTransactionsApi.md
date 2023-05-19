# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioTransactionsFromDateToDateGet**](StatementOfTransactionsApi.md#PortfolioTransactionsFromDateToDateGet) | **Get** /portfolio/transactions/{fromDate}/{toDate} | Get Transaction Statement
[**PortfolioTransactionsGeneratepdfFromDateToDateGet**](StatementOfTransactionsApi.md#PortfolioTransactionsGeneratepdfFromDateToDateGet) | **Get** /portfolio/transactions/generatepdf/{fromDate}/{toDate} | Get pdf for Transaction Statement
[**PortfolioTransactionsGeneratepdfGet**](StatementOfTransactionsApi.md#PortfolioTransactionsGeneratepdfGet) | **Get** /portfolio/transactions/generatepdf | Get Transaction Statement pdf
[**PortfolioTransactionsGet**](StatementOfTransactionsApi.md#PortfolioTransactionsGet) | **Get** /portfolio/transactions | Get Transaction Statement

# **PortfolioTransactionsFromDateToDateGet**
> []Transactions PortfolioTransactionsFromDateToDateGet(ctx, consumerKey, sessionToken, fromDate, toDate, optional)
Get Transaction Statement

Returns the full DP Transaction Statement of Client for last one month of  Current Financial Year 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***StatementOfTransactionsApiPortfolioTransactionsFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsApiPortfolioTransactionsFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]Transactions**](transactions.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsGeneratepdfFromDateToDateGet**
> []Transactions PortfolioTransactionsGeneratepdfFromDateToDateGet(ctx, consumerKey, sessionToken, fromDate, toDate, optional)
Get pdf for Transaction Statement

Returns the full DP Transaction Statement pdf of Client for last one month of  Current Financial Year 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***StatementOfTransactionsApiPortfolioTransactionsGeneratepdfFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsApiPortfolioTransactionsGeneratepdfFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]Transactions**](transactions.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsGeneratepdfGet**
> InlineResponse2001 PortfolioTransactionsGeneratepdfGet(ctx, consumerKey, sessionToken, optional)
Get Transaction Statement pdf

Returns the full DP Transaction Statement of Client for last one month of  Current Financial Year pdf path 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
 **optional** | ***StatementOfTransactionsApiPortfolioTransactionsGeneratepdfGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsApiPortfolioTransactionsGeneratepdfGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioTransactionsGet**
> []Transactions PortfolioTransactionsGet(ctx, consumerKey, sessionToken, optional)
Get Transaction Statement

Returns the full DP Transaction Statement of Client for last one month of  Current Financial Year 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
 **optional** | ***StatementOfTransactionsApiPortfolioTransactionsGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StatementOfTransactionsApiPortfolioTransactionsGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]Transactions**](transactions.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

