# VendorCreatedWebhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The ID of a seller associated with this webhook event. | [optional] [default to null]
**AccountId** | **string** | The ID of a target platform account associated with this webhook event, if the event is associated with the platform account. | [optional] [default to null]
**LocationId** | **string** | The ID of a location associated with the webhook event, if the event is associated with the location of the seller. | [optional] [default to null]
**Type_** | **string** | The type of this webhook event. The value is &#x60;\&quot;vendor.created\&quot;.&#x60; | [optional] [default to null]
**EventId** | **string** | A unique ID for this webhoook event. | [optional] [default to null]
**CreatedAt** | **string** | The RFC 3339-formatted time when the underlying webhook event data object is created. | [optional] [default to null]
**Data** | [***VendorCreatedWebhookData**](VendorCreatedWebhookData.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

