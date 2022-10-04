# CatalogModifier

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The modifier name.  This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points. | [optional] [default to null]
**PriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Ordinal** | **int32** | Determines where this &#x60;CatalogModifier&#x60; appears in the &#x60;CatalogModifierList&#x60;. | [optional] [default to null]
**ModifierListId** | **string** | The ID of the &#x60;CatalogModifierList&#x60; associated with this modifier. | [optional] [default to null]
**ImageIds** | **[]string** | The IDs of images associated with this &#x60;CatalogModifier&#x60; instance. Currently these images are not displayed by Square, but are free to be displayed in 3rd party applications. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

