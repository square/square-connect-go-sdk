# RefundPaymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** |  A unique string that identifies this &#x60;RefundPayment&#x60; request. The key can be any valid string but must be unique for every &#x60;RefundPayment&#x60; request.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency). | [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**AppFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**PaymentId** | **string** | The unique ID of the payment being refunded. Must be provided and non-empty. | [optional] [default to null]
**Reason** | **string** | A description of the reason for the refund. | [optional] [default to null]
**PaymentVersionToken** | **string** |  Used for optimistic concurrency. This opaque token identifies the current &#x60;Payment&#x60;  version that the caller expects. If the server has a different version of the Payment,  the update fails and a response with a VERSION_MISMATCH error is returned.  If the versions match, or the field is not provided, the refund proceeds as normal. | [optional] [default to null]
**TeamMemberId** | **string** | An optional [TeamMember](entity:TeamMember) ID to associate with this refund. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

