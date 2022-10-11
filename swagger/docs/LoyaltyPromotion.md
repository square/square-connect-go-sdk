# LoyaltyPromotion

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the promotion. | [optional] [default to null]
**Name** | **string** | The name of the promotion. | [default to null]
**Incentive** | [***LoyaltyPromotionIncentive**](LoyaltyPromotionIncentive.md) |  | [default to null]
**AvailableTime** | [***LoyaltyPromotionAvailableTimeData**](LoyaltyPromotionAvailableTimeData.md) |  | [default to null]
**TriggerLimit** | [***LoyaltyPromotionTriggerLimit**](LoyaltyPromotionTriggerLimit.md) |  | [optional] [default to null]
**Status** | [***LoyaltyPromotionStatus**](LoyaltyPromotionStatus.md) |  | [optional] [default to null]
**CreatedAt** | **string** | The timestamp of when the promotion was created, in RFC 3339 format. | [optional] [default to null]
**CanceledAt** | **string** | The timestamp of when the promotion was canceled, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the promotion was last updated, in RFC 3339 format. | [optional] [default to null]
**LoyaltyProgramId** | **string** | The ID of the [loyalty program](entity:LoyaltyProgram) associated with the promotion. | [optional] [default to null]
**MinimumSpendAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

