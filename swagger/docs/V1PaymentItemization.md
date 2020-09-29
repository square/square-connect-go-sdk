# V1PaymentItemization

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The item&#x27;s name. | [optional] [default to null]
**Quantity** | **float64** | The quantity of the item purchased. This can be a decimal value. | [optional] [default to null]
**ItemizationType** | [***V1PaymentItemizationItemizationType**](V1PaymentItemizationItemizationType.md) |  | [optional] [default to null]
**ItemDetail** | [***V1PaymentItemDetail**](V1PaymentItemDetail.md) |  | [optional] [default to null]
**Notes** | **string** | Notes entered by the merchant about the item at the time of payment, if any. | [optional] [default to null]
**ItemVariationName** | **string** | The name of the item variation purchased, if any. | [optional] [default to null]
**TotalMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**SingleQuantityMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**GrossSalesMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**DiscountMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**NetSalesMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**Taxes** | [**[]V1PaymentTax**](V1PaymentTax.md) | All taxes applied to this itemization. | [optional] [default to null]
**Discounts** | [**[]V1PaymentDiscount**](V1PaymentDiscount.md) | All discounts applied to this itemization. | [optional] [default to null]
**Modifiers** | [**[]V1PaymentModifier**](V1PaymentModifier.md) | All modifier options applied to this itemization. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

