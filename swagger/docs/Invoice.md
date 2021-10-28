# Invoice

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the invoice. | [optional] [default to null]
**Version** | **int32** | The Square-assigned version number, which is incremented each time an update is committed to the invoice. | [optional] [default to null]
**LocationId** | **string** | The ID of the location that this invoice is associated with.   If specified in a &#x60;CreateInvoice&#x60; request, the value must match the &#x60;location_id&#x60; of the associated order. | [optional] [default to null]
**OrderId** | **string** | The ID of the [order](entity:Order) for which the invoice is created.  This field is required when creating an invoice, and the order must be in the &#x60;OPEN&#x60; state.  To view the line items and other information for the associated order, call the  [RetrieveOrder](api-endpoint:Orders-RetrieveOrder) endpoint using the order ID. | [optional] [default to null]
**PrimaryRecipient** | [***InvoiceRecipient**](InvoiceRecipient.md) |  | [optional] [default to null]
**PaymentRequests** | [**[]InvoicePaymentRequest**](InvoicePaymentRequest.md) | The payment schedule for the invoice, represented by one or more payment requests that define payment settings, such as amount due and due date. An invoice supports the following payment request combinations: - One balance - One deposit with one balance - 2–12 installments  - One deposit with 2–12 installments  This field is required when creating an invoice. It must contain at least one payment request.  All payment requests for the invoice must equal the total order amount. For more information, see  [Payment requests](https://developer.squareup.com/docs/invoices-api/overview#payment-requests).  Adding &#x60;INSTALLMENT&#x60; payment requests to an invoice requires an  [Invoices Plus subscription](https://developer.squareup.com/docs/invoices-api/overview#invoices-plus-subscription). | [optional] [default to null]
**DeliveryMethod** | [***InvoiceDeliveryMethod**](InvoiceDeliveryMethod.md) |  | [optional] [default to null]
**InvoiceNumber** | **string** | A user-friendly invoice number. The value is unique within a location. If not provided when creating an invoice, Square assigns a value. It increments from 1 and padded with zeros making it 7 characters long (for example, 0000001 and 0000002). | [optional] [default to null]
**Title** | **string** | The title of the invoice. | [optional] [default to null]
**Description** | **string** | The description of the invoice. This is visible to the customer receiving the invoice. | [optional] [default to null]
**ScheduledAt** | **string** | The timestamp when the invoice is scheduled for processing, in RFC 3339 format. After the invoice is published, Square processes the invoice on the specified date, according to the delivery method and payment request settings.  If the field is not set, Square processes the invoice immediately after it is published. | [optional] [default to null]
**PublicUrl** | **string** | The URL of the Square-hosted invoice page. After you publish the invoice using the &#x60;PublishInvoice&#x60; endpoint, Square hosts the invoice page and returns the page URL in the response. | [optional] [default to null]
**NextPaymentAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Status** | [***InvoiceStatus**](InvoiceStatus.md) |  | [optional] [default to null]
**Timezone** | **string** | The time zone used to interpret calendar dates on the invoice, such as &#x60;due_date&#x60;. When an invoice is created, this field is set to the &#x60;timezone&#x60; specified for the seller location. The value cannot be changed.  For example, a payment &#x60;due_date&#x60; of 2021-03-09 with a &#x60;timezone&#x60; of America/Los\\_Angeles becomes overdue at midnight on March 9 in America/Los\\_Angeles (which equals a UTC timestamp of 2021-03-10T08:00:00Z). | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the invoice was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the invoice was last updated, in RFC 3339 format. | [optional] [default to null]
**AcceptedPaymentMethods** | [***InvoiceAcceptedPaymentMethods**](InvoiceAcceptedPaymentMethods.md) |  | [optional] [default to null]
**CustomFields** | [**[]InvoiceCustomField**](InvoiceCustomField.md) | Additional seller-defined fields to render on the invoice. These fields are visible to sellers and buyers on the Square-hosted invoice page and in emailed or PDF copies of invoices. For more information, see [Custom fields](https://developer.squareup.com/docs/invoices-api/overview#custom-fields).  Adding custom fields to an invoice requires an  [Invoices Plus subscription](https://developer.squareup.com/docs/invoices-api/overview#invoices-plus-subscription).  Max: 2 custom fields | [optional] [default to null]
**SubscriptionId** | **string** | The ID of the [subscription](entity:Subscription) associated with the invoice. This field is present only on subscription billing invoices. | [optional] [default to null]
**SaleOrServiceDate** | **string** | The date of the sale or the date that the service is rendered, in &#x60;YYYY-MM-DD&#x60; format. This field can be used to specify a past or future date which is displayed on the invoice. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

