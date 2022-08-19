# Merchant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-issued ID of the merchant. | [optional] [default to null]
**BusinessName** | **string** | The name of the merchant&#x27;s overall business. | [optional] [default to null]
**Country** | [***Country**](Country.md) |  | [default to null]
**LanguageCode** | **string** | The code indicating the [language preferences](https://developer.squareup.com/docs/build-basics/general-considerations/language-preferences) of the merchant, in [BCP 47 format](https://tools.ietf.org/html/bcp47#appendix-A). For example, &#x60;en-US&#x60; or &#x60;fr-CA&#x60;. | [optional] [default to null]
**Currency** | [***Currency**](Currency.md) |  | [optional] [default to null]
**Status** | [***MerchantStatus**](MerchantStatus.md) |  | [optional] [default to null]
**MainLocationId** | **string** | The ID of the [main &#x60;Location&#x60;](https://developer.squareup.com/docs/locations-api#about-the-main-location) for this merchant. | [optional] [default to null]
**CreatedAt** | **string** | The time when the merchant was created, in RFC 3339 format.    For more information, see [Working with Dates](https://developer.squareup.com/docs/build-basics/working-with-dates). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

