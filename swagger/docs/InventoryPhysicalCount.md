# InventoryPhysicalCount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique Square-generated ID for the [InventoryPhysicalCount](entity:InventoryPhysicalCount). | [optional] [default to null]
**ReferenceId** | **string** | An optional ID provided by the application to tie the [InventoryPhysicalCount](entity:InventoryPhysicalCount) to an external system. | [optional] [default to null]
**CatalogObjectId** | **string** | The Square-generated ID of the [CatalogObject](entity:CatalogObject) being tracked. | [optional] [default to null]
**CatalogObjectType** | **string** | The [type](entity:CatalogObjectType) of the [CatalogObject](entity:CatalogObject) being tracked.   The Inventory API supports setting and reading the &#x60;\&quot;catalog_object_type\&quot;: \&quot;ITEM_VARIATION\&quot;&#x60; field value.  In addition, it can also read the &#x60;\&quot;catalog_object_type\&quot;: \&quot;ITEM\&quot;&#x60; field value that is set by the Square Restaurants app. | [optional] [default to null]
**Status** | [***InventoryState**](InventoryState.md) |  | [optional] [default to null]
**State** | [***InventoryState**](InventoryState.md) |  | [optional] [default to null]
**LocationId** | **string** | The Square-generated ID of the [Location](entity:Location) where the related quantity of items is being tracked. | [optional] [default to null]
**Quantity** | **string** | The number of items affected by the physical count as a decimal string. The number can support up to 5 digits after the decimal point. | [optional] [default to null]
**Source** | [***SourceApplication**](SourceApplication.md) |  | [optional] [default to null]
**EmployeeId** | **string** | The Square-generated ID of the [Employee](entity:Employee) responsible for the physical count. | [optional] [default to null]
**TeamMemberId** | **string** | The Square-generated ID of the [Team Member](entity:TeamMember) responsible for the physical count. | [optional] [default to null]
**OccurredAt** | **string** | A client-generated RFC 3339-formatted timestamp that indicates when the physical count was examined. For physical count updates, the &#x60;occurred_at&#x60; timestamp cannot be older than 24 hours or in the future relative to the time of the request. | [optional] [default to null]
**CreatedAt** | **string** | An RFC 3339-formatted timestamp that indicates when the physical count is received. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

