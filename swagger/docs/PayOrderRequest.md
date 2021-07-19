# PayOrderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** | A value you specify that uniquely identifies this request among requests you have sent. If you are unsure whether a particular payment request was completed successfully, you can reattempt it with the same idempotency key without worrying about duplicate payments.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency). | [default to null]
**OrderVersion** | **int32** | The version of the order being paid. If not supplied, the latest version will be paid. | [optional] [default to null]
**PaymentIds** | **[]string** | The IDs of the [payments](entity:Payment) to collect. The payment total must match the order total. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

