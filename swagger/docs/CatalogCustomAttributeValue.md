# CatalogCustomAttributeValue

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the custom attribute. | [optional] [default to null]
**StringValue** | **string** | The string value of the custom attribute.  Populated if &#x60;type&#x60; &#x3D; &#x60;STRING&#x60;. | [optional] [default to null]
**CustomAttributeDefinitionId** | **string** | __Read-only.__ The id of the [CatalogCustomAttributeDefinition](#type-CatalogCustomAttributeDefinition) this value belongs to. | [optional] [default to null]
**Type_** | [***CatalogCustomAttributeDefinitionType**](CatalogCustomAttributeDefinitionType.md) |  | [optional] [default to null]
**NumberValue** | **string** | Populated if &#x60;type&#x60; &#x3D; &#x60;NUMBER&#x60;. Contains a string representation of a decimal number, using a &#x60;.&#x60; as the decimal separator. | [optional] [default to null]
**BooleanValue** | **bool** | A &#x60;true&#x60; or &#x60;false&#x60; value. Populated if &#x60;type&#x60; &#x3D; &#x60;BOOLEAN&#x60;. | [optional] [default to null]
**SelectionUidValues** | **[]string** | One or more choices from &#x60;allowed_selections&#x60;. Populated if &#x60;type&#x60; &#x3D; &#x60;SELECTION&#x60;. | [optional] [default to null]
**Key** | **string** | __Read-only.__ A copy of key from the associated &#x60;CatalogCustomAttributeDefinition&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

