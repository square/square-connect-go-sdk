# ObtainTokenResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | A valid OAuth access token. OAuth access tokens are 64 bytes long. Provide the access token in a header with every request to Connect API endpoints. See [OAuth API: Walkthrough](https://developer.squareup.com/docs/oauth-api/walkthrough) for more information. | [optional] [default to null]
**TokenType** | **string** | This value is always _bearer_. | [optional] [default to null]
**ExpiresAt** | **string** | The date when access_token expires, in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format. | [optional] [default to null]
**MerchantId** | **string** | The ID of the authorizing merchant&#x27;s business. | [optional] [default to null]
**SubscriptionId** | **string** | __LEGACY FIELD__. The ID of a subscription plan the merchant signed up for. Only present if the merchant signed up for a subscription during authorization. | [optional] [default to null]
**PlanId** | **string** | __LEGACY FIELD__. The ID of the subscription plan the merchant signed up for. Only present if the merchant signed up for a subscription during authorization. | [optional] [default to null]
**IdToken** | **string** | Then OpenID token belonging to this this person. Only present if the OPENID scope is included in the authorize request. | [optional] [default to null]
**RefreshToken** | **string** | A refresh token. OAuth refresh tokens are 64 bytes long. For more information, see [OAuth access token management](https://developer.squareup.com/docs/authz/oauth/how-it-works#oauth-access-token-management). | [optional] [default to null]
**ShortLived** | **bool** | A boolean indicating the access token is a short-lived access token. The short-lived access token returned in the response will expire in 24 hours. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

