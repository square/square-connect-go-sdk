# V1PaymentTax

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Name** | **string** | The merchant-defined name of the tax. | [optional] [default to null]
**AppliedMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**Rate** | **string** | The rate of the tax, as a string representation of a decimal number. A value of 0.07 corresponds to a rate of 7%. | [optional] [default to null]
**InclusionType** | [***V1PaymentTaxInclusionType**](V1PaymentTaxInclusionType.md) |  | [optional] [default to null]
**FeeId** | **string** | The ID of the tax, if available. Taxes applied in older versions of Square Register might not have an ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

