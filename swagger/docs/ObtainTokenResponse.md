# ObtainTokenResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | A valid OAuth access token. OAuth access tokens are 64 bytes long. Provide the access token in a header with every request to Connect API endpoints. For more information, see [OAuth API: Walkthrough](https://developer.squareup.com/docs/oauth-api/walkthrough). | [optional] [default to null]
**TokenType** | **string** | This value is always _bearer_. | [optional] [default to null]
**ExpiresAt** | **string** | The date when the &#x60;access_token&#x60; expires, in [ISO 8601](http://www.iso.org/iso/home/standards/iso8601.htm) format. | [optional] [default to null]
**MerchantId** | **string** | The ID of the authorizing merchant&#x27;s business. | [optional] [default to null]
**SubscriptionId** | **string** | __LEGACY FIELD__. The ID of a subscription plan the merchant signed up for. The ID is only present if the merchant signed up for a subscription plan during authorization. | [optional] [default to null]
**PlanId** | **string** | __LEGACY FIELD__. The ID of the subscription plan the merchant signed up for. The ID is only present if the merchant signed up for a subscription plan during authorization. | [optional] [default to null]
**IdToken** | **string** | The OpenID token belonging to this person. This token is only present if the OPENID scope is included in the authorization request. | [optional] [default to null]
**RefreshToken** | **string** | A refresh token. OAuth refresh tokens are 64 bytes long. For more information, see [Refresh, Revoke, and Limit the Scope of OAuth Tokens](https://developer.squareup.com/docs/oauth-api/refresh-revoke-limit-scope). | [optional] [default to null]
**ShortLived** | **bool** | A Boolean indicating that the access token is a short-lived access token. The short-lived access token returned in the response expires in 24 hours. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

