# InvoicePaymentRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | The Square-generated ID of the payment request in an [invoice](entity:Invoice). | [optional] [default to null]
**RequestMethod** | [***InvoiceRequestMethod**](InvoiceRequestMethod.md) |  | [optional] [default to null]
**RequestType** | [***InvoiceRequestType**](InvoiceRequestType.md) |  | [optional] [default to null]
**DueDate** | **string** | The due date (in the invoice&#x27;s time zone) for the payment request, in &#x60;YYYY-MM-DD&#x60; format. This field is required to create a payment request. If an &#x60;automatic_payment_source&#x60; is defined for the request, Square charges the payment source on this date.  After this date, the invoice becomes overdue. For example, a payment &#x60;due_date&#x60; of 2021-03-09 with a &#x60;timezone&#x60; of America/Los\\_Angeles becomes overdue at midnight on March 9 in America/Los\\_Angeles (which equals a UTC timestamp of 2021-03-10T08:00:00Z). | [optional] [default to null]
**FixedAmountRequestedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**PercentageRequested** | **string** | Specifies the amount for the payment request in percentage:  - When the payment &#x60;request_type&#x60; is &#x60;DEPOSIT&#x60;, it is the percentage of the order&#x27;s total amount. - When the payment &#x60;request_type&#x60; is &#x60;INSTALLMENT&#x60;, it is the percentage of the order&#x27;s total less  the deposit, if requested. The sum of the &#x60;percentage_requested&#x60; in all installment  payment requests must be equal to 100.  You cannot specify this when the payment &#x60;request_type&#x60; is &#x60;BALANCE&#x60; or when the  payment request specifies the &#x60;fixed_amount_requested_money&#x60; field. | [optional] [default to null]
**TippingEnabled** | **bool** | If set to true, the Square-hosted invoice page (the &#x60;public_url&#x60; field of the invoice)  provides a place for the customer to pay a tip.   This field is allowed only on the final payment request   and the payment &#x60;request_type&#x60; must be &#x60;BALANCE&#x60; or &#x60;INSTALLMENT&#x60;. | [optional] [default to null]
**AutomaticPaymentSource** | [***InvoiceAutomaticPaymentSource**](InvoiceAutomaticPaymentSource.md) |  | [optional] [default to null]
**CardId** | **string** | The ID of the credit or debit card on file to charge for the payment request. To get the cards on file for a customer, call [ListCards](api-endpoint:Cards-ListCards) and include the &#x60;customer_id&#x60; of the invoice recipient. | [optional] [default to null]
**Reminders** | [**[]InvoicePaymentReminder**](InvoicePaymentReminder.md) | A list of one or more reminders to send for the payment request. | [optional] [default to null]
**ComputedAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalCompletedAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**RoundingAdjustmentIncludedMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

