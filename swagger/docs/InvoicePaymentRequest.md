# InvoicePaymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | The Square-generated ID of the payment request in an [invoice](#type-invoice). | [optional] [default to null]
**RequestMethod** | [***InvoiceRequestMethod**](InvoiceRequestMethod.md) |  | [optional] [default to null]
**RequestType** | [***InvoiceRequestType**](InvoiceRequestType.md) |  | [optional] [default to null]
**DueDate** | **string** | The due date (in the invoice location&#x27;s time zone) for the payment request.  After this date, the invoice becomes overdue. | [optional] [default to null]
**FixedAmountRequestedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**PercentageRequested** | **string** | Specifies the amount for the payment request in percentage:  - When the payment &#x60;request_type&#x60; is &#x60;DEPOSIT&#x60;, it is the percentage of the order total amount. - When the payment &#x60;request_type&#x60; is &#x60;INSTALLMENT&#x60;, it is the percentage of the order total less  the deposit, if requested. The sum of the &#x60;percentage_requested&#x60; in all installment  payment requests must be equal to 100.  You cannot specify this when the payment &#x60;request_type&#x60; is &#x60;BALANCE&#x60; or when the  payment request specifies the &#x60;fixed_amount_requested_money&#x60; field. | [optional] [default to null]
**TippingEnabled** | **bool** | If set to true, the Square-hosted invoice page (the &#x60;public_url&#x60; field of the invoice)  provides a place for the customer to pay a tip.   This field is allowed only on the final payment request   and the payment &#x60;request_type&#x60; must be &#x60;BALANCE&#x60; or &#x60;INSTALLMENT&#x60;. | [optional] [default to null]
**CardId** | **string** | If the request method is &#x60;CHARGE_CARD_ON_FILE&#x60;, this field provides the  card to charge. | [optional] [default to null]
**Reminders** | [**[]InvoicePaymentReminder**](InvoicePaymentReminder.md) | A list of one or more reminders to send for the payment request. | [optional] [default to null]
**ComputedAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalCompletedAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**RoundingAdjustmentIncludedMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

