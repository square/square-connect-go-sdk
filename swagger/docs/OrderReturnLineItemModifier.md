# OrderReturnLineItemModifier

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | A unique ID that identifies the return modifier only within this order. | [optional] [default to null]
**SourceModifierUid** | **string** | The modifier &#x60;uid&#x60; from the order&#x27;s line item that contains the original sale of this line item modifier. | [optional] [default to null]
**CatalogObjectId** | **string** | The catalog object ID referencing [CatalogModifier](entity:CatalogModifier). | [optional] [default to null]
**Name** | **string** | The name of the item modifier. | [optional] [default to null]
**BasePriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalPriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

