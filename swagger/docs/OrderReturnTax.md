# OrderReturnTax

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | A unique ID that identifies the returned tax only within this order. | [optional] [default to null]
**SourceTaxUid** | **string** | The tax &#x60;uid&#x60; from the order that contains the original tax charge. | [optional] [default to null]
**CatalogObjectId** | **string** | The catalog object ID referencing [CatalogTax](entity:CatalogTax). | [optional] [default to null]
**CatalogVersion** | **int64** | The version of the catalog object that this tax references. | [optional] [default to null]
**Name** | **string** | The tax&#x27;s name. | [optional] [default to null]
**Type_** | [***OrderLineItemTaxType**](OrderLineItemTaxType.md) |  | [optional] [default to null]
**Percentage** | **string** | The percentage of the tax, as a string representation of a decimal number. For example, a value of &#x60;\&quot;7.25\&quot;&#x60; corresponds to a percentage of 7.25%. | [optional] [default to null]
**AppliedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Scope** | [***OrderLineItemTaxScope**](OrderLineItemTaxScope.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

