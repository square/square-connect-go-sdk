# CalculateLoyaltyPointsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Points** | **int32** | The number of points that the buyer can earn from the base loyalty program. | [optional] [default to null]
**PromotionPoints** | **int32** | The number of points that the buyer can earn from a loyalty promotion. To be eligible to earn promotion points, the purchase must first qualify for program points. When &#x60;order_id&#x60; is not provided in the request, this value is always 0. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

