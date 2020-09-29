# OrderReturnDiscount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique ID that identifies the return discount only within this order. | [optional] [default to null]
**SourceDiscountUid** | **string** | &#x60;uid&#x60; of the Discount from the Order which contains the original application of this discount. | [optional] [default to null]
**CatalogObjectId** | **string** | The catalog object id referencing [CatalogDiscount](#type-catalogdiscount). | [optional] [default to null]
**Name** | **string** | The discount&#x27;s name. | [optional] [default to null]
**Type_** | [***OrderLineItemDiscountType**](OrderLineItemDiscountType.md) |  | [optional] [default to null]
**Percentage** | **string** | The percentage of the tax, as a string representation of a decimal number. A value of &#x60;7.25&#x60; corresponds to a percentage of 7.25%.  &#x60;percentage&#x60; is not set for amount-based discounts. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**AppliedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Scope** | [***OrderLineItemDiscountScope**](OrderLineItemDiscountScope.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

