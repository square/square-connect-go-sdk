# RefundPaymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** |  A unique string that identifies this &#x60;RefundPayment&#x60; request. The key can be any valid string but must be unique for every &#x60;RefundPayment&#x60; request.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency). | [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**AppFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**PaymentId** | **string** | The unique ID of the payment being refunded. | [default to null]
**Reason** | **string** | A description of the reason for the refund. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

