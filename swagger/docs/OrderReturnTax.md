# OrderReturnTax

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique ID that identifies the return tax only within this order. | [optional] [default to null]
**SourceTaxUid** | **string** | &#x60;uid&#x60; of the Tax from the Order which contains the original charge of this tax. | [optional] [default to null]
**CatalogObjectId** | **string** | The catalog object id referencing [CatalogTax](#type-catalogtax). | [optional] [default to null]
**Name** | **string** | The tax&#x27;s name. | [optional] [default to null]
**Type_** | [***OrderLineItemTaxType**](OrderLineItemTaxType.md) |  | [optional] [default to null]
**Percentage** | **string** | The percentage of the tax, as a string representation of a decimal number. For example, a value of &#x60;\&quot;7.25\&quot;&#x60; corresponds to a percentage of 7.25%. | [optional] [default to null]
**AppliedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Scope** | [***OrderLineItemTaxScope**](OrderLineItemTaxScope.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

