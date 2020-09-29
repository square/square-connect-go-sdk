# CatalogInfoResponseLimits

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BatchUpsertMaxObjectsPerBatch** | **int32** | The maximum number of objects that may appear within a single batch in a &#x60;/v2/catalog/batch-upsert&#x60; request. | [optional] [default to null]
**BatchUpsertMaxTotalObjects** | **int32** | The maximum number of objects that may appear across all batches in a &#x60;/v2/catalog/batch-upsert&#x60; request. | [optional] [default to null]
**BatchRetrieveMaxObjectIds** | **int32** | The maximum number of object IDs that may appear in a &#x60;/v2/catalog/batch-retrieve&#x60; request. | [optional] [default to null]
**SearchMaxPageLimit** | **int32** | The maximum number of results that may be returned in a page of a &#x60;/v2/catalog/search&#x60; response. | [optional] [default to null]
**BatchDeleteMaxObjectIds** | **int32** | The maximum number of object IDs that may be included in a single &#x60;/v2/catalog/batch-delete&#x60; request. | [optional] [default to null]
**UpdateItemTaxesMaxItemIds** | **int32** | The maximum number of item IDs that may be included in a single &#x60;/v2/catalog/update-item-taxes&#x60; request. | [optional] [default to null]
**UpdateItemTaxesMaxTaxesToEnable** | **int32** | The maximum number of tax IDs to be enabled that may be included in a single &#x60;/v2/catalog/update-item-taxes&#x60; request. | [optional] [default to null]
**UpdateItemTaxesMaxTaxesToDisable** | **int32** | The maximum number of tax IDs to be disabled that may be included in a single &#x60;/v2/catalog/update-item-taxes&#x60; request. | [optional] [default to null]
**UpdateItemModifierListsMaxItemIds** | **int32** | The maximum number of item IDs that may be included in a single &#x60;/v2/catalog/update-item-modifier-lists&#x60; request. | [optional] [default to null]
**UpdateItemModifierListsMaxModifierListsToEnable** | **int32** | The maximum number of modifier list IDs to be enabled that may be included in a single &#x60;/v2/catalog/update-item-modifier-lists&#x60; request. | [optional] [default to null]
**UpdateItemModifierListsMaxModifierListsToDisable** | **int32** | The maximum number of modifier list IDs to be disabled that may be included in a single &#x60;/v2/catalog/update-item-modifier-lists&#x60; request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

