# Location

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A short generated string of letters and numbers that uniquely identifies this location instance. | [optional] [default to null]
**Name** | **string** | The name of the location. This information appears in the Seller Dashboard as the nickname. A location name must be unique within a seller account. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**Timezone** | **string** | The [IANA time zone](https://www.iana.org/time-zones) identifier for the time zone of the location. For example, &#x60;America/Los_Angeles&#x60;. | [optional] [default to null]
**Capabilities** | [**[]LocationCapability**](LocationCapability.md) | The Square features that are enabled for the location. See [LocationCapability](entity:LocationCapability) for possible values. See [LocationCapability](#type-locationcapability) for possible values | [optional] [default to null]
**Status** | [***LocationStatus**](LocationStatus.md) |  | [optional] [default to null]
**CreatedAt** | **string** | The time when the location was created, in RFC 3339 format. For more information, see [Working with Dates](https://developer.squareup.com/docs/build-basics/working-with-dates). | [optional] [default to null]
**MerchantId** | **string** | The ID of the merchant that owns the location. | [optional] [default to null]
**Country** | [***Country**](Country.md) |  | [optional] [default to null]
**LanguageCode** | **string** | The language associated with the location, in [BCP 47 format](https://tools.ietf.org/html/bcp47#appendix-A). For more information, see [Location language code](https://developer.squareup.com/docs/locations-api#location-language-code). | [optional] [default to null]
**Currency** | [***Currency**](Currency.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | The phone number of the location. For example, &#x60;+1 855-700-6000&#x60;. | [optional] [default to null]
**BusinessName** | **string** | The name of the location&#x27;s overall business. This name is present on receipts and other customer-facing branding. | [optional] [default to null]
**Type_** | [***LocationType**](LocationType.md) |  | [optional] [default to null]
**WebsiteUrl** | **string** | The website URL of the location.  For example, &#x60;https://squareup.com&#x60;. | [optional] [default to null]
**BusinessHours** | [***BusinessHours**](BusinessHours.md) |  | [optional] [default to null]
**BusinessEmail** | **string** | The email address of the location. This can be unique to the location and is not always the email address for the business owner or administrator. | [optional] [default to null]
**Description** | **string** | The description of the location. For example, &#x60;Main Street location&#x60;. | [optional] [default to null]
**TwitterUsername** | **string** | The Twitter username of the location without the &#x27;@&#x27; symbol. For example, &#x60;Square&#x60;. | [optional] [default to null]
**InstagramUsername** | **string** | The Instagram username of the location without the &#x27;@&#x27; symbol. For example, &#x60;square&#x60;. | [optional] [default to null]
**FacebookUrl** | **string** | The Facebook profile URL of the location. The URL should begin with &#x27;facebook.com/&#x27;. For example, &#x60;https://www.facebook.com/square&#x60;. | [optional] [default to null]
**Coordinates** | [***Coordinates**](Coordinates.md) |  | [optional] [default to null]
**LogoUrl** | **string** | The URL of the logo image for the location. When configured in the Seller Dashboard (Receipts section), the logo appears on transactions (such as receipts and invoices) that Square generates on behalf of the seller. This image should have a roughly square (1:1) aspect ratio and should be at least 200x200 pixels. | [optional] [default to null]
**PosBackgroundUrl** | **string** | The URL of the Point of Sale background image for the location. | [optional] [default to null]
**Mcc** | **string** | A four-digit number that describes the kind of goods or services sold at the location. The [merchant category code (MCC)](https://developer.squareup.com/docs/locations-api#initialize-a-merchant-category-code) of the location as standardized by ISO 18245. For example, &#x60;5045&#x60;, for a location that sells computer goods and software. | [optional] [default to null]
**FullFormatLogoUrl** | **string** | The URL of a full-format logo image for the location. When configured in the Seller Dashboard (Receipts section), the logo appears on transactions (such as receipts and invoices) that Square generates on behalf of the seller. This image can be wider than it is tall and should be at least 1280x648 pixels. | [optional] [default to null]
**TaxIds** | [***TaxIds**](TaxIds.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

