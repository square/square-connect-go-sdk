# OrderReturnDiscount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | A unique ID that identifies the returned discount only within this order. | [optional] [default to null]
**SourceDiscountUid** | **string** | The discount &#x60;uid&#x60; from the order that contains the original application of this discount. | [optional] [default to null]
**CatalogObjectId** | **string** | The catalog object ID referencing [CatalogDiscount](entity:CatalogDiscount). | [optional] [default to null]
**CatalogVersion** | **int64** | The version of the catalog object that this discount references. | [optional] [default to null]
**Name** | **string** | The discount&#x27;s name. | [optional] [default to null]
**Type_** | [***OrderLineItemDiscountType**](OrderLineItemDiscountType.md) |  | [optional] [default to null]
**Percentage** | **string** | The percentage of the tax, as a string representation of a decimal number. A value of &#x60;\&quot;7.25\&quot;&#x60; corresponds to a percentage of 7.25%.  &#x60;percentage&#x60; is not set for amount-based discounts. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**AppliedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Scope** | [***OrderLineItemDiscountScope**](OrderLineItemDiscountScope.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

