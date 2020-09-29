# AuthorizeRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The Square-issued ID of the application requesting permissions. | [default to null]
**Scope** | [**[]OAuthPermission**](OAuthPermission.md) | __OPTIONAL__  A space-separated list of the permissions the application is requesting. Default: \&quot;&#x60;MERCHANT_PROFILE_READ PAYMENTS_READ SETTLEMENTS_READ BANK_ACCOUNTS_READ&#x60;\&quot; See [OAuthPermission](#type-oauthpermission) for possible values | [optional] [default to null]
**Locale** | **string** | __OPTIONAL__  The locale to present the permission request form in. Square detects the appropriate locale automatically. Only provide this value if the application can definitively determine the preferred locale.  Currently supported values: &#x60;en-US&#x60;, &#x60;en-CA&#x60;, &#x60;es-US&#x60;, &#x60;fr-CA&#x60;, &#x60;ja-JP&#x60;. | [optional] [default to null]
**Session** | **bool** | If &#x60;false&#x60;, the user must log in to their Square account to view the Permission Request form, even if they already have a valid user session. Default: &#x60;true&#x60; | [optional] [default to null]
**State** | **string** | __OPTIONAL__  When provided, &#x60;state&#x60; is passed along to the configured Redirect URL after the Permission Request form is submitted. You can include state and verify its value to help protect against cross-site request forgery. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

