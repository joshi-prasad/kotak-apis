# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/portfolio/1.1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PortfolioPriceForUpdateFyFromDateToDateGet**](UpdatePriceStockApi.md#PortfolioPriceForUpdateFyFromDateToDateGet) | **Get** /portfolio/price-for-update/{fy}/{fromDate}/{toDate} | Get update price stock Report by Details
[**PortfolioPriceForUpdateScripwiseFyFromDateToDatePost**](UpdatePriceStockApi.md#PortfolioPriceForUpdateScripwiseFyFromDateToDatePost) | **Post** /portfolio/price-for-update/scripwise/{fy}/{fromDate}/{toDate} | Update price stock details
[**PortfolioPriceForUpdateScripwisePost**](UpdatePriceStockApi.md#PortfolioPriceForUpdateScripwisePost) | **Post** /portfolio/price-for-update/scripwise | Update price stock details with order type

# **PortfolioPriceForUpdateFyFromDateToDateGet**
> []InlineResponse2004 PortfolioPriceForUpdateFyFromDateToDateGet(ctx, consumerKey, sessionToken, fy, fromDate, toDate, type_, optional)
Get update price stock Report by Details

returns update price stock Report Details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
  **type_** | **string**| type of stock transfer | 
 **optional** | ***UpdatePriceStockApiPortfolioPriceForUpdateFyFromDateToDateGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdatePriceStockApiPortfolioPriceForUpdateFyFromDateToDateGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **userId** | **optional.String**| user Id for your application | 
 **clientAccount** | **optional.String**| Client account for which data requested | 

### Return type

[**[]InlineResponse2004**](inline_response_200_4.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioPriceForUpdateScripwiseFyFromDateToDatePost**
> []InlineResponse2005 PortfolioPriceForUpdateScripwiseFyFromDateToDatePost(ctx, consumerKey, sessionToken, fy, fromDate, toDate, type_, optional)
Update price stock details

Update price stock details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
  **fy** | **string**| Financial Year : Format - YYYY-YYYY | 
  **fromDate** | **string**| From Date : Format - YYYY-MM-DD | 
  **toDate** | **string**| To Date : Format - YYYY-MM-DD | 
  **type_** | **string**| type of stock transfer | 
 **optional** | ***UpdatePriceStockApiPortfolioPriceForUpdateScripwiseFyFromDateToDatePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdatePriceStockApiPortfolioPriceForUpdateScripwiseFyFromDateToDatePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | [**optional.Interface of []FromDateToDateBody**](fromDate_toDate_body.md)|  | 
 **userId** | **optional.**| user Id for your application | 
 **clientAccount** | **optional.**| Client account for which data requested | 

### Return type

[**[]InlineResponse2005**](inline_response_200_5.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PortfolioPriceForUpdateScripwisePost**
> []InlineResponse2005 PortfolioPriceForUpdateScripwisePost(ctx, consumerKey, sessionToken, optional)
Update price stock details with order type

Update price stock details with order type (restriction for 20 entries)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**| Unique ID for your application | 
  **sessionToken** | **string**| Session ID for your application | 
 **optional** | ***UpdatePriceStockApiPortfolioPriceForUpdateScripwisePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a UpdatePriceStockApiPortfolioPriceForUpdateScripwisePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []PriceforupdateScripwiseBody**](priceforupdate_scripwise_body.md)|  | 
 **userId** | **optional.**| user Id for your application | 
 **clientAccount** | **optional.**| Client account for which data requested | 

### Return type

[**[]InlineResponse2005**](inline_response_200_5.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

