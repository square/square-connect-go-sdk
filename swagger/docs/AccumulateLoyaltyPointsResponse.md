# AccumulateLoyaltyPointsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Event** | [***LoyaltyEvent**](LoyaltyEvent.md) |  | [optional] [default to null]
**Events** | [**[]LoyaltyEvent**](LoyaltyEvent.md) | The resulting loyalty events. If the purchase qualifies for points, the &#x60;ACCUMULATE_POINTS&#x60; event is always included. When using the Orders API, the &#x60;ACCUMULATE_PROMOTION_POINTS&#x60; event is included if the purchase also qualifies for a loyalty promotion. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

