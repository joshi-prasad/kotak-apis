# InlineResponse200

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstrumentName** | **string** | Name of the instrument | [optional] [default to null]
**InstrumentToken** | **int32** |  | [optional] [default to null]
**LastUpdatedTime** | **int64** | Last time in epoch format when this data was updated | [optional] [default to null]
**LastTradedTime** | **int64** | Last time in epoch format when this scrip was traded | [optional] [default to null]
**LastPrice** | **float32** | Last traded price | [optional] [default to null]
**LastTradedQuantity** | **int32** | Last traded quantity | [optional] [default to null]
**TotalBuyQuantity** | **int64** | Total bid quantity | [optional] [default to null]
**TotalSellQuantity** | **int64** | Total ask quantity | [optional] [default to null]
**AverageTradedPrice** | **float32** | Average traded price for the day | [optional] [default to null]
**OpenInterest** | **int64** | Open interest in case of derivateive contracts | [optional] [default to null]
**NetChange** | **float32** | Absolute change in price from previous day close | [optional] [default to null]
**DprLow** | **float32** | Lower value of daily price range | [optional] [default to null]
**DprHigh** | **float32** | Higher value of daily price range | [optional] [default to null]
**Ohlc** | [***InlineResponse200Ohlc**](inline_response_200_ohlc.md) |  | [optional] [default to null]
**Depth** | [***InlineResponse200Depth**](inline_response_200_depth.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

