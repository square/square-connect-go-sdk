# OrderReturnLineItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | A unique ID for this return line-item entry. | [optional] [default to null]
**SourceLineItemUid** | **string** | The &#x60;uid&#x60; of the line item in the original sale order. | [optional] [default to null]
**Name** | **string** | The name of the line item. | [optional] [default to null]
**Quantity** | **string** | The quantity returned, formatted as a decimal number. For example, &#x60;\&quot;3\&quot;&#x60;.  Line items with a &#x60;quantity_unit&#x60; can have non-integer quantities. For example, &#x60;\&quot;1.70000\&quot;&#x60;. | [default to null]
**QuantityUnit** | [***OrderQuantityUnit**](OrderQuantityUnit.md) |  | [optional] [default to null]
**Note** | **string** | The note of the return line item. | [optional] [default to null]
**CatalogObjectId** | **string** | The [CatalogItemVariation](entity:CatalogItemVariation) ID applied to this return line item. | [optional] [default to null]
**VariationName** | **string** | The name of the variation applied to this return line item. | [optional] [default to null]
**ItemType** | [***OrderLineItemItemType**](OrderLineItemItemType.md) |  | [optional] [default to null]
**ReturnModifiers** | [**[]OrderReturnLineItemModifier**](OrderReturnLineItemModifier.md) | The [CatalogModifier](entity:CatalogModifier)s applied to this line item. | [optional] [default to null]
**AppliedTaxes** | [**[]OrderLineItemAppliedTax**](OrderLineItemAppliedTax.md) | The list of references to &#x60;OrderReturnTax&#x60; entities applied to the return line item. Each &#x60;OrderLineItemAppliedTax&#x60; has a &#x60;tax_uid&#x60; that references the &#x60;uid&#x60; of a top-level &#x60;OrderReturnTax&#x60; applied to the return line item. On reads, the applied amount is populated. | [optional] [default to null]
**AppliedDiscounts** | [**[]OrderLineItemAppliedDiscount**](OrderLineItemAppliedDiscount.md) | The list of references to &#x60;OrderReturnDiscount&#x60; entities applied to the return line item. Each &#x60;OrderLineItemAppliedDiscount&#x60; has a &#x60;discount_uid&#x60; that references the &#x60;uid&#x60; of a top-level &#x60;OrderReturnDiscount&#x60; applied to the return line item. On reads, the applied amount is populated. | [optional] [default to null]
**BasePriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**VariationTotalPriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**GrossReturnMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalTaxMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalDiscountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

