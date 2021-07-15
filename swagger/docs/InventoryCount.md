# InventoryCount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogObjectId** | **string** | The Square-generated ID of the [CatalogObject](entity:CatalogObject) being tracked. | [optional] [default to null]
**CatalogObjectType** | **string** | The [type](entity:CatalogObjectType) of the [CatalogObject](entity:CatalogObject) being tracked. Tracking is only supported for the &#x60;ITEM_VARIATION&#x60; type. | [optional] [default to null]
**Status** | [***InventoryState**](InventoryState.md) |  | [optional] [default to null]
**State** | [***InventoryState**](InventoryState.md) |  | [optional] [default to null]
**LocationId** | **string** | The Square-generated ID of the [Location](entity:Location) where the related quantity of items is being tracked. | [optional] [default to null]
**Quantity** | **string** | The number of items affected by the estimated count as a decimal string. Can support up to 5 digits after the decimal point. | [optional] [default to null]
**CalculatedAt** | **string** | An RFC 3339-formatted timestamp that indicates when the most recent physical count or adjustment affecting the estimated count is received. | [optional] [default to null]
**IsEstimated** | **bool** | Whether the inventory count is for composed variation (TRUE) or not (FALSE). If true, the inventory count will not be present in the response of any of these endpoints: [BatchChangeInventory](api-endpoint:Inventory-BatchChangeInventory),  [BatchRetrieveInventoryChanges](api-endpoint:Inventory-BatchRetrieveInventoryChanges),  [BatchRetrieveInventoryCounts](api-endpoint:Inventory-BatchRetrieveInventoryCounts), and  [RetrieveInventoryChanges](api-endpoint:Inventory-RetrieveInventoryChanges). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

