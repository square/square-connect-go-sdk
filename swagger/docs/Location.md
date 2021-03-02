# Location

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-issued ID of the location. | [optional] [default to null]
**Name** | **string** | The name of the location. This information appears in the dashboard as the nickname. A location name must be unique within a seller account. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**Timezone** | **string** | The [IANA Timezone](https://www.iana.org/time-zones) identifier for the timezone of the location. | [optional] [default to null]
**Capabilities** | [**[]LocationCapability**](LocationCapability.md) | The Square features that are enabled for the location. See [LocationCapability](#type-locationcapability) for possible values. See [LocationCapability](#type-locationcapability) for possible values | [optional] [default to null]
**Status** | [***LocationStatus**](LocationStatus.md) |  | [optional] [default to null]
**CreatedAt** | **string** | The time when the location was created, in RFC 3339 format. | [optional] [default to null]
**MerchantId** | **string** | The ID of the merchant that owns the location. | [optional] [default to null]
**Country** | [***Country**](Country.md) |  | [optional] [default to null]
**LanguageCode** | **string** | The language associated with the location, in [BCP 47 format](https://tools.ietf.org/html/bcp47#appendix-A). | [optional] [default to null]
**Currency** | [***Currency**](Currency.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | The phone number of the location in human readable format. | [optional] [default to null]
**BusinessName** | **string** | The business name of the location This is the name visible to the customers of the location. For example, this name appears on customer receipts. | [optional] [default to null]
**Type_** | [***LocationType**](LocationType.md) |  | [optional] [default to null]
**WebsiteUrl** | **string** | The website URL of the location. | [optional] [default to null]
**BusinessHours** | [***BusinessHours**](BusinessHours.md) |  | [optional] [default to null]
**BusinessEmail** | **string** | The email of the location. This email is visible to the customers of the location. For example, the email appears on customer receipts. | [optional] [default to null]
**Description** | **string** | The description of the location. | [optional] [default to null]
**TwitterUsername** | **string** | The Twitter username of the location without the &#x27;@&#x27; symbol. | [optional] [default to null]
**InstagramUsername** | **string** | The Instagram username of the location without the &#x27;@&#x27; symbol. | [optional] [default to null]
**FacebookUrl** | **string** | The Facebook profile URL of the location. The URL should begin with &#x27;facebook.com/&#x27;. | [optional] [default to null]
**Coordinates** | [***Coordinates**](Coordinates.md) |  | [optional] [default to null]
**LogoUrl** | **string** | The URL of the logo image for the location. The Seller must choose this logo in the Seller dashboard (Receipts section) for the logo to appear on transactions (such as receipts, invoices) that Square generates on behalf of the Seller. This image should have an aspect ratio close to 1:1 and is recommended to be at least 200x200 pixels. | [optional] [default to null]
**PosBackgroundUrl** | **string** | The URL of the Point of Sale background image for the location. | [optional] [default to null]
**Mcc** | **string** | The merchant category code (MCC) of the location, as standardized by ISO 18245. The MCC describes the kind of goods or services sold at the location. | [optional] [default to null]
**FullFormatLogoUrl** | **string** | The URL of a full-format logo image for the location. The Seller must choose this logo in the Seller dashboard (Receipts section) for the logo to appear on transactions (such as receipts, invoices) that Square generates on behalf of the Seller. This image can have an aspect ratio of 2:1 or greater and is recommended to be at least 1280x648 pixels. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

