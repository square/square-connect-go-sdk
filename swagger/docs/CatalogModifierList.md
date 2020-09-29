# CatalogModifierList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name for the &#x60;CatalogModifierList&#x60; instance. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points. | [optional] [default to null]
**Ordinal** | **int32** | Determines where this modifier list appears in a list of &#x60;CatalogModifierList&#x60; values. | [optional] [default to null]
**SelectionType** | [***CatalogModifierListSelectionType**](CatalogModifierListSelectionType.md) |  | [optional] [default to null]
**Modifiers** | [**[]CatalogObject**](CatalogObject.md) | The options included in the &#x60;CatalogModifierList&#x60;. You must include at least one &#x60;CatalogModifier&#x60;. Each CatalogObject must have type &#x60;MODIFIER&#x60; and contain &#x60;CatalogModifier&#x60; data. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

