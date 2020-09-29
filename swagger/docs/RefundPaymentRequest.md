# RefundPaymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** |  A unique string that identifies this RefundPayment request. Key can be any valid string but must be unique for every RefundPayment request.  For more information, see [Idempotency keys](https://developer.squareup.com/docs/working-with-apis/idempotency). | [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**AppFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**PaymentId** | **string** | Unique ID of the payment being refunded. | [default to null]
**Reason** | **string** | A description of the reason for the refund. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

