# LoyaltyProgramAccrualRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualType** | [***LoyaltyProgramAccrualRuleType**](LoyaltyProgramAccrualRuleType.md) |  | [default to null]
**Points** | **int32** | The number of points that  buyers earn based on the &#x60;accrual_type&#x60;. | [optional] [default to null]
**VisitMinimumAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**SpendAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CatalogObjectId** | **string** | When the accrual rule is item-based or category-based, this field specifies the ID  of the [catalog object](entity:CatalogObject) that buyers can purchase to earn points.  If &#x60;accrual_type&#x60; is &#x60;ITEM_VARIATION&#x60;, the object is an item variation.  If &#x60;accrual_type&#x60; is &#x60;CATEGORY&#x60;, the object is a category. | [optional] [default to null]
**ExcludedCategoryIds** | **[]string** | When the accrual rule is spend-based (&#x60;accrual_type&#x60; is &#x60;SPEND&#x60;), this field  lists the IDs of any &#x60;CATEGORY&#x60; catalog objects that are excluded from points accrual.   You can use the [BatchRetrieveCatalogObjects](api-endpoint:Catalog-BatchRetrieveCatalogObjects)  endpoint to retrieve information about the excluded categories. | [optional] [default to null]
**ExcludedItemVariationIds** | **[]string** | When the accrual rule is spend-based (&#x60;accrual_type&#x60; is &#x60;SPEND&#x60;), this field  lists the IDs of any &#x60;ITEM_VARIATION&#x60; catalog objects that are excluded from points accrual.   You can use the [BatchRetrieveCatalogObjects](api-endpoint:Catalog-BatchRetrieveCatalogObjects)  endpoint to retrieve information about the excluded item variations. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

