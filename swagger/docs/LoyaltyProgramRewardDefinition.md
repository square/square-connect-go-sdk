# LoyaltyProgramRewardDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | [***LoyaltyProgramRewardDefinitionScope**](LoyaltyProgramRewardDefinitionScope.md) |  | [default to null]
**DiscountType** | [***LoyaltyProgramRewardDefinitionType**](LoyaltyProgramRewardDefinitionType.md) |  | [default to null]
**PercentageDiscount** | **string** | Present if &#x60;discount_type&#x60; is &#x60;FIXED_PERCENTAGE&#x60;. For example, a 7.25% off discount will be represented as \&quot;7.25\&quot;. | [optional] [default to null]
**CatalogObjectIds** | **[]string** | A list of [catalog object](#type-CatalogObject) ids to which this reward can be applied. They are either all item-variation ids or category ids, depending on the &#x60;type&#x60; field. | [optional] [default to null]
**FixedDiscountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**MaxDiscountMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

