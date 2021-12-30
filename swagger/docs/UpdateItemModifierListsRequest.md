# UpdateItemModifierListsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemIds** | **[]string** | The IDs of the catalog items associated with the CatalogModifierList objects being updated. | [default to null]
**ModifierListsToEnable** | **[]string** | The IDs of the CatalogModifierList objects to enable for the CatalogItem. At least one of &#x60;modifier_lists_to_enable&#x60; or &#x60;modifier_lists_to_disable&#x60; must be specified. | [optional] [default to null]
**ModifierListsToDisable** | **[]string** | The IDs of the CatalogModifierList objects to disable for the CatalogItem. At least one of &#x60;modifier_lists_to_enable&#x60; or &#x60;modifier_lists_to_disable&#x60; must be specified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

