# PaymentOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autocomplete** | **bool** | Indicates whether the &#x60;Payment&#x60; objects created from this &#x60;TerminalCheckout&#x60; are automatically &#x60;COMPLETED&#x60; or left in an &#x60;APPROVED&#x60; state for later modification. | [optional] [default to null]
**DelayDuration** | **string** | The duration of time after the payment&#x27;s creation when Square automatically cancels the payment. This automatic cancellation applies only to payments that do not reach a terminal state (COMPLETED, CANCELED, or FAILED) before the &#x60;delay_duration&#x60; time period.  This parameter should be specified as a time duration, in RFC 3339 format, with a minimum value of 1 minute.  Note: This feature is only supported for card payments. This parameter can only be set for a delayed capture payment (&#x60;autocomplete&#x3D;false&#x60;). Default: - Card-present payments: \&quot;PT36H\&quot; (36 hours) from the creation time. - Card-not-present payments: \&quot;P7D\&quot; (7 days) from the creation time. | [optional] [default to null]
**AcceptPartialAuthorization** | **bool** | If set to &#x60;true&#x60; and charging a Square Gift Card, a payment might be returned with &#x60;amount_money&#x60; equal to less than what was requested. For example, a request for $20 when charging a Square Gift Card with a balance of $5 results in an APPROVED payment of $5. You might choose to prompt the buyer for an additional payment to cover the remainder or cancel the Gift Card payment.  This field cannot be &#x60;true&#x60; when &#x60;autocomplete &#x3D; true&#x60;. This field cannot be &#x60;true&#x60; when an &#x60;order_id&#x60; isn&#x27;t specified.  For more information, see [Take Partial Payments](https://developer.squareup.com/docs/payments-api/take-payments/card-payments/partial-payments-with-gift-cards).  Default: false | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

