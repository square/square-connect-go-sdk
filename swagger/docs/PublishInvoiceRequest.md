# PublishInvoiceRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **int32** | The version of the [invoice](entity:Invoice) to publish. This must match the current version of the invoice; otherwise, the request is rejected. | [default to null]
**IdempotencyKey** | **string** | A unique string that identifies the &#x60;PublishInvoice&#x60; request. If you do not  provide &#x60;idempotency_key&#x60; (or provide an empty string as the value), the endpoint  treats each request as independent.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

