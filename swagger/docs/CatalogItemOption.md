# CatalogItemOption

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The item option&#x27;s display name for the seller. Must be unique across all item options. This is a searchable attribute for use in applicable query filters. | [optional] [default to null]
**DisplayName** | **string** | The item option&#x27;s display name for the customer. This is a searchable attribute for use in applicable query filters. | [optional] [default to null]
**Description** | **string** | The item option&#x27;s human-readable description. Displayed in the Square Point of Sale app for the seller and in the Online Store or on receipts for the buyer. This is a searchable attribute for use in applicable query filters. | [optional] [default to null]
**ShowColors** | **bool** | If true, display colors for entries in &#x60;values&#x60; when present. | [optional] [default to null]
**Values** | [**[]CatalogObject**](CatalogObject.md) | A list of CatalogObjects containing the &#x60;CatalogItemOptionValue&#x60;s for this item. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

