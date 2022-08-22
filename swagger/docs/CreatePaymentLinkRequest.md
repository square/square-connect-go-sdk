# CreatePaymentLinkRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** | A unique string that identifies this &#x60;CreatePaymentLinkRequest&#x60; request. If you do not provide a unique string (or provide an empty string as the value), the endpoint treats each request as independent.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency). | [optional] [default to null]
**Description** | **string** | A description of the payment link. You provide this optional description that is useful in your application context. It is not used anywhere. | [optional] [default to null]
**QuickPay** | [***QuickPay**](QuickPay.md) |  | [optional] [default to null]
**Order** | [***Order**](Order.md) |  | [optional] [default to null]
**CheckoutOptions** | [***CheckoutOptions**](CheckoutOptions.md) |  | [optional] [default to null]
**PrePopulatedData** | [***PrePopulatedData**](PrePopulatedData.md) |  | [optional] [default to null]
**PaymentNote** | **string** | A note for the payment. After processing the payment, Square adds this note to the resulting &#x60;Payment&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

