# LoyaltyReward

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the loyalty reward. | [optional] [default to null]
**Status** | [***LoyaltyRewardStatus**](LoyaltyRewardStatus.md) |  | [optional] [default to null]
**LoyaltyAccountId** | **string** | The Square-assigned ID of the [loyalty account](entity:LoyaltyAccount) to which the reward belongs. | [default to null]
**RewardTierId** | **string** | The Square-assigned ID of the [reward tier](entity:LoyaltyProgramRewardTier) used to create the reward. | [default to null]
**Points** | **int32** | The number of loyalty points used for the reward. | [optional] [default to null]
**OrderId** | **string** | The Square-assigned ID of the [order](entity:Order) to which the reward is attached. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the reward was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the reward was last updated, in RFC 3339 format. | [optional] [default to null]
**RedeemedAt** | **string** | The timestamp when the reward was redeemed, in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

