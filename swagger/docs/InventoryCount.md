# InventoryCount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogObjectId** | **string** | The Square generated ID of the &#x60;CatalogObject&#x60; being tracked. | [optional] [default to null]
**CatalogObjectType** | **string** | The &#x60;CatalogObjectType&#x60; of the &#x60;CatalogObject&#x60; being tracked. Tracking is only supported for the &#x60;ITEM_VARIATION&#x60; type. | [optional] [default to null]
**Status** | [***InventoryState**](InventoryState.md) |  | [optional] [default to null]
**State** | [***InventoryState**](InventoryState.md) |  | [optional] [default to null]
**LocationId** | **string** | The Square ID of the [Location](#type-location) where the related quantity of items are being tracked. | [optional] [default to null]
**Quantity** | **string** | The number of items affected by the estimated count as a decimal string. Can support up to 5 digits after the decimal point. | [optional] [default to null]
**CalculatedAt** | **string** | A read-only timestamp in RFC 3339 format that indicates when Square received the most recent physical count or adjustment that had an affect on the estimated count. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

