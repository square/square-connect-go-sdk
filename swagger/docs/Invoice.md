# Invoice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the invoice. | [optional] [default to null]
**Version** | **int32** | The Square-assigned version number, which is incremented each time an update is committed to the invoice. | [optional] [default to null]
**LocationId** | **string** | The ID of the location that this invoice is associated with. This field is required when creating an invoice. | [optional] [default to null]
**OrderId** | **string** | The ID of the [order](#type-order) for which the invoice is created.  This order must be in the &#x60;OPEN&#x60; state and must belong to the &#x60;location_id&#x60; specified for this invoice. This field is required when creating an invoice. | [optional] [default to null]
**PrimaryRecipient** | [***InvoiceRecipient**](InvoiceRecipient.md) |  | [optional] [default to null]
**PaymentRequests** | [**[]InvoicePaymentRequest**](InvoicePaymentRequest.md) | The payment schedule for the invoice, represented by one or more payment requests that define payment settings, such as amount due and due date. You can specify a maximum of 13 payment requests, with up to 12 &#x60;INSTALLMENT&#x60; request types. For more information, see [Payment requests](https://developer.squareup.com/docs/invoices-api/overview#payment-requests).  This field is required when creating an invoice. It must contain at least one payment request. | [optional] [default to null]
**DeliveryMethod** | [***InvoiceDeliveryMethod**](InvoiceDeliveryMethod.md) |  | [optional] [default to null]
**InvoiceNumber** | **string** | A user-friendly invoice number. The value is unique within a location. If not provided when creating an invoice, Square assigns a value. It increments from 1 and padded with zeros making it 7 characters long for example, 0000001, 0000002. | [optional] [default to null]
**Title** | **string** | The title of the invoice. | [optional] [default to null]
**Description** | **string** | The description of the invoice. This is visible to the customer receiving the invoice. | [optional] [default to null]
**ScheduledAt** | **string** | The timestamp when the invoice is scheduled for processing, in RFC 3339 format. After the invoice is published, Square processes the invoice on the specified date, according to the delivery method and payment request settings.  If the field is not set, Square processes the invoice immediately after it is published. | [optional] [default to null]
**PublicUrl** | **string** | The URL of the Square-hosted invoice page. After you publish the invoice using the &#x60;PublishInvoice&#x60; endpoint, Square hosts the invoice page and returns the page URL in the response. | [optional] [default to null]
**NextPaymentAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Status** | [***InvoiceStatus**](InvoiceStatus.md) |  | [optional] [default to null]
**Timezone** | **string** | The time zone of the date values (for example, &#x60;due_date&#x60;) specified in the invoice. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the invoice was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the invoice was last updated, in RFC 3339 format. | [optional] [default to null]
**CustomFields** | [**[]InvoiceCustomField**](InvoiceCustomField.md) | Additional seller-defined fields to render on the invoice. These fields are visible to sellers and buyers on the Square-hosted invoice page and in emailed or PDF copies of invoices. For more information, see [Custom fields](https://developer.squareup.com/docs/invoices-api/overview#custom-fields).  Max: 2 custom fields | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

