# Invoice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the invoice. | [optional] [default to null]
**Version** | **int32** | The version number, which is incremented each time an update is committed to the invoice. | [optional] [default to null]
**LocationId** | **string** | The ID of the location that this invoice is associated with. This field is required when creating an invoice. | [optional] [default to null]
**OrderId** | **string** | The ID of the [order](#type-order) for which the invoice is created.  This order must be in the &#x60;OPEN&#x60; state and must belong to the &#x60;location_id&#x60; specified for this invoice. This field is required when creating an invoice. | [optional] [default to null]
**PrimaryRecipient** | [***InvoiceRecipient**](InvoiceRecipient.md) |  | [optional] [default to null]
**PaymentRequests** | [**[]InvoicePaymentRequest**](InvoicePaymentRequest.md) | An array of &#x60;InvoicePaymentRequest&#x60; objects. Each object defines a payment request in an invoice payment schedule. It provides information such as when and how Square processes payments. You can specify maximum of nine payment requests. All all the payment requests must specify the same &#x60;request_method&#x60;.  This field is required when creating an invoice. | [optional] [default to null]
**InvoiceNumber** | **string** | A user-friendly invoice number. The value is unique within a location. If not provided when creating an invoice, Square assigns a value. It increments from 1 and padded with zeros making it 7 characters long for example, 0000001, 0000002. | [optional] [default to null]
**Title** | **string** | The title of the invoice. | [optional] [default to null]
**Description** | **string** | The description of the invoice. This is visible the customer receiving the invoice. | [optional] [default to null]
**ScheduledAt** | **string** | The timestamp when the invoice is scheduled for processing, in RFC 3339 format. At the specified time, depending on the &#x60;request_method&#x60;, Square sends the invoice to the customer&#x27;s email address or charge the customer&#x27;s card on file.  If the field is not set, Square processes the invoice immediately after publication. | [optional] [default to null]
**PublicUrl** | **string** | The URL of the Square-hosted invoice page. After you publish the invoice using the &#x60;PublishInvoice&#x60; endpoint, Square hosts the invoice page and returns the page URL in the response. | [optional] [default to null]
**NextPaymentAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Status** | [***InvoiceStatus**](InvoiceStatus.md) |  | [optional] [default to null]
**Timezone** | **string** | The time zone of the date values (for example, &#x60;due_date&#x60;) specified in the invoice. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the invoice was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the invoice was last updated, in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

