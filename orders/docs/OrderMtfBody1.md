# OrderMtfBody1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentToken** | **int64** | Instrument token of the scrip to be traded | [optional] [default to null]
**TransactionType** | **string** | Transaction Type - BUY or SELL | [optional] [default to null]
**Quantity** | **int64** | Order quantity - specified in same unit as quoted in market depth | [optional] [default to null]
**Price** | **float32** | Order Price, non zero positive for limit order and zero for market order | [optional] [default to null]
**Validity** | **string** | Validity of the order - GFD, IOC etc | [optional] [default to null]
**Variety** | **string** | Variety of the order - REGULAR, AMO, SQUAREOFF - for Super Multiple Orders etc | [optional] [default to null]
**DisclosedQuantity** | **int32** | Quantity to be disclosed in order | [optional] [default to null]
**TriggerPrice** | **float32** | Trigger price, required for stoploss or supermultiple order | [optional] [default to null]
**Tag** | **string** | Tag for this order | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

