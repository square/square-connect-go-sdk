# Vendor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique Square-generated ID for the [Vendor](entity:Vendor). This field is required when attempting to update a [Vendor](entity:Vendor). | [optional] [default to null]
**CreatedAt** | **string** | An RFC 3339-formatted timestamp that indicates when the [Vendor](entity:Vendor) was created. | [optional] [default to null]
**UpdatedAt** | **string** | An RFC 3339-formatted timestamp that indicates when the [Vendor](entity:Vendor) was last updated. | [optional] [default to null]
**Name** | **string** | The name of the [Vendor](entity:Vendor). This field is required when attempting to create or update a [Vendor](entity:Vendor). | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**Contacts** | [**[]VendorContact**](VendorContact.md) | The contacts of the [Vendor](entity:Vendor). | [optional] [default to null]
**AccountNumber** | **string** | The account number of the [Vendor](entity:Vendor). | [optional] [default to null]
**Note** | **string** | A note detailing information about the [Vendor](entity:Vendor). | [optional] [default to null]
**Version** | **int32** | The version of the [Vendor](entity:Vendor). | [optional] [default to null]
**Status** | [***VendorStatus**](VendorStatus.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

