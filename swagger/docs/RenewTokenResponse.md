# RenewTokenResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The renewed access token. This value might be different from the &#x60;access_token&#x60; you provided in your request. You provide this token in a header with every request to Connect API endpoints. See [Request and response headers](https://developer.squareup.com/docs/api/connect/v2/#requestandresponseheaders) for the format of this header. | [optional] [default to null]
**TokenType** | **string** | This value is always _bearer_. | [optional] [default to null]
**ExpiresAt** | **string** | The date when the &#x60;access_token&#x60; expires, in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format. | [optional] [default to null]
**MerchantId** | **string** | The ID of the authorizing merchant&#x27;s business. | [optional] [default to null]
**SubscriptionId** | **string** | __LEGACY FIELD__. The ID of the merchant subscription associated with the authorization. The ID is only present if the merchant signed up for a subscription during authorization. | [optional] [default to null]
**PlanId** | **string** | __LEGACY FIELD__. The ID of the subscription plan the merchant signed up for. The ID is only present if the merchant signed up for a subscription plan during authorization. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

