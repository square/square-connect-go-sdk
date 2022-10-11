# CatalogDiscount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The discount name. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points. | [optional] [default to null]
**DiscountType** | [***CatalogDiscountType**](CatalogDiscountType.md) |  | [optional] [default to null]
**Percentage** | **string** | The percentage of the discount as a string representation of a decimal number, using a &#x60;.&#x60; as the decimal separator and without a &#x60;%&#x60; sign. A value of &#x60;7.5&#x60; corresponds to &#x60;7.5%&#x60;. Specify a percentage of &#x60;0&#x60; if &#x60;discount_type&#x60; is &#x60;VARIABLE_PERCENTAGE&#x60;.  Do not use this field for amount-based or variable discounts. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**PinRequired** | **bool** | Indicates whether a mobile staff member needs to enter their PIN to apply the discount to a payment in the Square Point of Sale app. | [optional] [default to null]
**LabelColor** | **string** | The color of the discount display label in the Square Point of Sale app. This must be a valid hex color code. | [optional] [default to null]
**ModifyTaxBasis** | [***CatalogDiscountModifyTaxBasis**](CatalogDiscountModifyTaxBasis.md) |  | [optional] [default to null]
**MaximumAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

