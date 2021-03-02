# UpdateInvoiceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | [***Invoice**](Invoice.md) |  | [default to null]
**IdempotencyKey** | **string** | A unique string that identifies the &#x60;UpdateInvoice&#x60; request. If you do not provide &#x60;idempotency_key&#x60; (or provide an empty string as the value), the endpoint treats each request as independent.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency). | [optional] [default to null]
**FieldsToClear** | **[]string** | The list of fields to clear. For examples, see [Update an invoice](https://developer.squareup.com/docs/invoices-api/overview#update-an-invoice). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

