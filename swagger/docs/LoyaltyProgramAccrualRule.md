# LoyaltyProgramAccrualRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualType** | [***LoyaltyProgramAccrualRuleType**](LoyaltyProgramAccrualRuleType.md) |  | [default to null]
**Points** | **int32** | The number of points that  buyers earn based on the &#x60;accrual_type&#x60;. | [optional] [default to null]
**VisitMinimumAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**SpendAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CatalogObjectId** | **string** | The ID of the [catalog object](#type-CatalogObject) to purchase to earn the number of points defined by the rule. This is either an item variation or a category, depending on the type. This is defined on &#x60;ITEM_VARIATION&#x60; rules and &#x60;CATEGORY&#x60; rules. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

