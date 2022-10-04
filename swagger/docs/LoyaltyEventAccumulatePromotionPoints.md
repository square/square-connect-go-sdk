# LoyaltyEventAccumulatePromotionPoints

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoyaltyProgramId** | **string** | The Square-assigned ID of the [loyalty program](entity:LoyaltyProgram). | [optional] [default to null]
**LoyaltyPromotionId** | **string** | The Square-assigned ID of the [loyalty promotion](entity:LoyaltyPromotion). | [optional] [default to null]
**Points** | **int32** | The number of points earned by the event. | [default to null]
**OrderId** | **string** | The ID of the [order](entity:Order) for which the buyer earned the promotion points. Only applications that use the Orders API to process orders can trigger this event. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

