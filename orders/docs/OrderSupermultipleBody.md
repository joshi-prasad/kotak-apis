# OrderSupermultipleBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | Order ID of the order to be modified | [default to null]
**Quantity** | **int32** | Order quantity - specified in same unit as quoted in market depth | [optional] [default to null]
**Price** | **float32** | Order Price, non zero positive for limit order and zero for market order | [optional] [default to null]
**DisclosedQuantity** | **int32** | Quantity to be disclosed in order | [optional] [default to null]
**OriginalOpenQuantity** | **int32** | Original open quantity of order | [optional] [default to null]
**OriginalPrice** | **float32** | Original order price input by client/user. | [optional] [default to null]
**TriggerPrice** | **float32** | Trigger price, required for stoploss or supermultiple order | [optional] [default to null]
**Validity** | **string** | Validity of the order - GFD, IOC etc | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

