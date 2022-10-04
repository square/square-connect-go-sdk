# WebhookSubscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A Square-generated unique ID for the subscription. | [optional] [default to null]
**Name** | **string** | The name of this subscription. | [optional] [default to null]
**Enabled** | **bool** | Indicates whether the subscription is enabled (&#x60;true&#x60;) or not (&#x60;false&#x60;). | [optional] [default to null]
**EventTypes** | **[]string** | The event types associated with this subscription. | [optional] [default to null]
**NotificationUrl** | **string** | The URL to which webhooks are sent. | [optional] [default to null]
**ApiVersion** | **string** | The API version of the subscription. This field is optional for &#x60;CreateWebhookSubscription&#x60;.  The value defaults to the API version used by the application. | [optional] [default to null]
**SignatureKey** | **string** | The Square-generated signature key used to validate the origin of the webhook event. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp of when the subscription was created, in RFC 3339 format. For example, \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp of when the subscription was last updated, in RFC 3339 format. For example, \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

