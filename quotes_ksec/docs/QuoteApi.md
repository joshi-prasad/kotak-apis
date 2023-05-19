# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/quotes/v1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DepthInstrumentsInstrumentTokensGet**](QuoteApi.md#DepthInstrumentsInstrumentTokensGet) | **Get** /depth/instruments/{instrumentTokens} | Get market details quote
[**InstrumentsInstrumentTokensGet**](QuoteApi.md#InstrumentsInstrumentTokensGet) | **Get** /instruments/{instrumentTokens} | Get full details
[**LtpInstrumentsInstrumentTokensGet**](QuoteApi.md#LtpInstrumentsInstrumentTokensGet) | **Get** /ltp/instruments/{instrumentTokens} | Get LTP quote
[**OhlcInstrumentsInstrumentTokensGet**](QuoteApi.md#OhlcInstrumentsInstrumentTokensGet) | **Get** /ohlc/instruments/{instrumentTokens} | Get OHLC quote

# **DepthInstrumentsInstrumentTokensGet**
> InlineResponse200 DepthInstrumentsInstrumentTokensGet(ctx, consumerKey, sessionToken, instrumentTokens)
Get market details quote

Returns market depth details for an array of scrips

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **instrumentTokens** | **string**| Instrument token of the scrip for which quote | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **InstrumentsInstrumentTokensGet**
> InstrumentsInstrumentTokensGet(ctx, consumerKey, sessionToken, instrumentTokens)
Get full details

Get full details

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **instrumentTokens** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LtpInstrumentsInstrumentTokensGet**
> InlineResponse2002 LtpInstrumentsInstrumentTokensGet(ctx, consumerKey, sessionToken, instrumentTokens)
Get LTP quote

Returns the LTP for an array of scrips

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **instrumentTokens** | **string**| Instrument token of the scrip for which quote | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OhlcInstrumentsInstrumentTokensGet**
> InlineResponse2001 OhlcInstrumentsInstrumentTokensGet(ctx, consumerKey, sessionToken, instrumentTokens)
Get OHLC quote

Returns the OHLC quote details for an array of scrips

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **consumerKey** | **string**|  | 
  **sessionToken** | **string**|  | 
  **instrumentTokens** | **string**| Instrument token of the scrip for which quote | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

