# LoyaltyProgramAccrualRuleSpendData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**ExcludedCategoryIds** | **[]string** | The IDs of any &#x60;CATEGORY&#x60; catalog objects that are excluded from points accrual.  You can use the [BatchRetrieveCatalogObjects](api-endpoint:Catalog-BatchRetrieveCatalogObjects) endpoint to retrieve information about the excluded categories. | [optional] [default to null]
**ExcludedItemVariationIds** | **[]string** | The IDs of any &#x60;ITEM_VARIATION&#x60; catalog objects that are excluded from points accrual.  You can use the [BatchRetrieveCatalogObjects](api-endpoint:Catalog-BatchRetrieveCatalogObjects) endpoint to retrieve information about the excluded item variations. | [optional] [default to null]
**TaxMode** | [***LoyaltyProgramAccrualRuleTaxMode**](LoyaltyProgramAccrualRuleTaxMode.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

