# LoyaltyProgram

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the loyalty program. Updates to  the loyalty program do not modify the identifier. | [default to null]
**Status** | [***LoyaltyProgramStatus**](LoyaltyProgramStatus.md) |  | [default to null]
**RewardTiers** | [**[]LoyaltyProgramRewardTier**](LoyaltyProgramRewardTier.md) | The list of rewards for buyers, sorted by ascending points. | [default to null]
**ExpirationPolicy** | [***LoyaltyProgramExpirationPolicy**](LoyaltyProgramExpirationPolicy.md) |  | [optional] [default to null]
**Terminology** | [***LoyaltyProgramTerminology**](LoyaltyProgramTerminology.md) |  | [default to null]
**LocationIds** | **[]string** | The [locations](#type-Location) at which the program is active. | [default to null]
**CreatedAt** | **string** | The timestamp when the program was created, in RFC 3339 format. | [default to null]
**UpdatedAt** | **string** | The timestamp when the reward was last updated, in RFC 3339 format. | [default to null]
**AccrualRules** | [**[]LoyaltyProgramAccrualRule**](LoyaltyProgramAccrualRule.md) | Defines how buyers can earn loyalty points. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

