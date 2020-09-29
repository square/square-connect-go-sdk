# V1Merchant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The merchant account&#x27;s unique identifier. | [optional] [default to null]
**Name** | **string** | The name associated with the merchant account. | [optional] [default to null]
**Email** | **string** | The email address associated with the merchant account. | [optional] [default to null]
**AccountType** | [***V1MerchantAccountType**](V1MerchantAccountType.md) |  | [optional] [default to null]
**AccountCapabilities** | **[]string** | Capabilities that are enabled for the merchant&#x27;s Square account. Capabilities that are not listed in this array are not enabled for the account. | [optional] [default to null]
**CountryCode** | **string** | The country associated with the merchant account, in ISO 3166-1-alpha-2 format. | [optional] [default to null]
**LanguageCode** | **string** | The language associated with the merchant account, in BCP 47 format. | [optional] [default to null]
**CurrencyCode** | **string** | The currency associated with the merchant account, in ISO 4217 format. For example, the currency code for US dollars is USD. | [optional] [default to null]
**BusinessName** | **string** | The name of the merchant&#x27;s business. | [optional] [default to null]
**BusinessAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**BusinessPhone** | [***V1PhoneNumber**](V1PhoneNumber.md) |  | [optional] [default to null]
**BusinessType** | [***V1MerchantBusinessType**](V1MerchantBusinessType.md) |  | [optional] [default to null]
**ShippingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**LocationDetails** | [***V1MerchantLocationDetails**](V1MerchantLocationDetails.md) |  | [optional] [default to null]
**MarketUrl** | **string** | The URL of the merchant&#x27;s online store. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

