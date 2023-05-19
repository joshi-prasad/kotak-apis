# {{classname}}

All URIs are relative to *https://tradeapi.kotaksecurities.com/apim/brokerage/1.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrokerageChargesDetailsPost**](DefaultApi.md#GetBrokerageChargesDetailsPost) | **Post** /getBrokerageChargesDetails | 

# **GetBrokerageChargesDetailsPost**
> GetBrokerageChargesDetailsPost(ctx, ksSesToken, optional)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ksSesToken** | **string**| base64(&#x27;{\&quot;clientId\&quot;:\&quot;your-clientId\&quot;,\&quot;sessionToken\&quot;:\&quot;your-sessionToken\&quot;,\&quot;sourceConsumerKey\&quot;:\&quot;your-sourceConsumerKey\&quot;}&#x27;) | 
 **optional** | ***DefaultApiGetBrokerageChargesDetailsPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetBrokerageChargesDetailsPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of GetBrokerageChargesDetailsBody**](GetBrokerageChargesDetailsBody.md)|  | 
 **userparams** | **optional.**|  | 

### Return type

 (empty response body)

### Authorization

[default](../README.md#default)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

