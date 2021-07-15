# RevokeTokenRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The Square issued ID for your application, available from the [developer dashboard](https://developer.squareup.com/apps). | [optional] [default to null]
**AccessToken** | **string** | The access token of the merchant whose token you want to revoke. Do not provide a value for merchant_id if you provide this parameter. | [optional] [default to null]
**MerchantId** | **string** | The ID of the merchant whose token you want to revoke. Do not provide a value for access_token if you provide this parameter. | [optional] [default to null]
**RevokeOnlyAccessToken** | **bool** | If &#x60;true&#x60;, terminate the given single access token, but do not terminate the entire authorization. Default: &#x60;false&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

