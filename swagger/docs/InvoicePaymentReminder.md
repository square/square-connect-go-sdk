# InvoicePaymentReminder

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | A Square-assigned ID that uniquely identifies the reminder within the &#x60;InvoicePaymentRequest&#x60;. | [optional] [default to null]
**RelativeScheduledDays** | **int32** | The number of days before (a negative number) or after (a positive number) the payment request &#x60;due_date&#x60; when the reminder is sent. For example, -3 indicates that the reminder should be sent 3 days before the payment request &#x60;due_date&#x60;. | [optional] [default to null]
**Message** | **string** | The reminder message. | [optional] [default to null]
**Status** | [***InvoicePaymentReminderStatus**](InvoicePaymentReminderStatus.md) |  | [optional] [default to null]
**SentAt** | **string** | If sent, the timestamp when the reminder was sent, in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

