# PayoutPaidWebhook

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** | The ID of the target merchant associated with the event. | [optional] [default to null]
**LocationId** | **string** | The ID of the target location associated with the event. | [optional] [default to null]
**Type_** | **string** | The type of event this represents, &#x60;\&quot;payout.paid\&quot;&#x60;. | [optional] [default to null]
**EventId** | **string** | A unique ID for the webhook event. | [optional] [default to null]
**CreatedAt** | **string** | Timestamp of when the webhook event was verified, in RFC 3339 format. | [optional] [default to null]
**Data** | [***PayoutPaidWebhookData**](PayoutPaidWebhookData.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

