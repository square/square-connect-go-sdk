# V1PaymentSurcharge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the surcharge. | [optional] [default to null]
**AppliedMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**Rate** | **string** | The amount of the surcharge as a percentage. The percentage is provided as a string representing the decimal equivalent of the percentage. For example, \&quot;0.7\&quot; corresponds to a 7% surcharge. Exactly one of rate or amount_money should be set. | [optional] [default to null]
**AmountMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**Type_** | [***V1PaymentSurchargeType**](V1PaymentSurchargeType.md) |  | [optional] [default to null]
**Taxable** | **bool** | Indicates whether the surcharge is taxable. | [optional] [default to null]
**Taxes** | [**[]V1PaymentTax**](V1PaymentTax.md) | The list of taxes that should be applied to the surcharge. | [optional] [default to null]
**SurchargeId** | **string** | A Square-issued unique identifier associated with the surcharge. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

