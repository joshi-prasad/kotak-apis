# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioPnlFyFromDateToDateGet**](RealisedProfitLossApi.md#PortfolioPnlFyFromDateToDateGet) | **Get** /portfolio/pnl/{fy}/{fromDate}/{toDate} | Get Realised Profit and Loss
[**PortfolioPnlGeneratepdfFyFromDateToDateGet**](RealisedProfitLossApi.md#PortfolioPnlGeneratepdfFyFromDateToDateGet) | **Get** /portfolio/pnl/generatepdf/{fy}/{fromDate}/{toDate} | Get pdf for Realised Profit and Loss
[**PortfolioPnlGeneratepdfGet**](RealisedProfitLossApi.md#PortfolioPnlGeneratepdfGet) | **Get** /portfolio/pnl/generatepdf | Get PDF for Realised Profit and Loss
[**PortfolioPnlGet**](RealisedProfitLossApi.md#PortfolioPnlGet) | **Get** /portfolio/pnl | Get Realised Profit and Loss

# **PortfolioPnlFyFromDateToDateGet**
> []Pnl PortfolioPnlFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, optional)
Get Realised Profit and Loss

Returns the full DP holdings of the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***RealisedProfitLossApiPortfolioPnlFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossApiPortfolioPnlFyFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]Pnl**](pnl.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioPnlGeneratepdfFyFromDateToDateGet**
> InlineResponse2001 PortfolioPnlGeneratepdfFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, optional)
Get pdf for Realised Profit and Loss

Returns pdf path for the full DP holdings of the user

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
 **optional** | ***RealisedProfitLossApiPortfolioPnlGeneratepdfFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossApiPortfolioPnlGeneratepdfFyFromDateToDateGetOpts struct
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

# **PortfolioPnlGeneratepdfGet**
> InlineResponse2001 PortfolioPnlGeneratepdfGet(ctx, consumerKey, sessionToken, optional)
Get PDF for Realised Profit and Loss

Returns Profit loss statement pdf for last one month of Current FY 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
 **optional** | ***RealisedProfitLossApiPortfolioPnlGeneratepdfGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossApiPortfolioPnlGeneratepdfGetOpts struct
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

# **PortfolioPnlGet**
> []Pnl PortfolioPnlGet(ctx, consumerKey, sessionToken, optional)
Get Realised Profit and Loss

Returns Profit loss statement for last one month of Current FY 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
 **optional** | ***RealisedProfitLossApiPortfolioPnlGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RealisedProfitLossApiPortfolioPnlGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]Pnl**](pnl.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

