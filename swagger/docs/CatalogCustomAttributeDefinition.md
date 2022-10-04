# CatalogCustomAttributeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | [***CatalogCustomAttributeDefinitionType**](CatalogCustomAttributeDefinitionType.md) |  | [default to null]
**Name** | **string** |  The name of this definition for API and seller-facing UI purposes. The name must be unique within the (merchant, application) pair. Required. May not be empty and may not exceed 255 characters. Can be modified after creation. | [default to null]
**Description** | **string** | Seller-oriented description of the meaning of this Custom Attribute, any constraints that the seller should observe, etc. May be displayed as a tooltip in Square UIs. | [optional] [default to null]
**SourceApplication** | [***SourceApplication**](SourceApplication.md) |  | [optional] [default to null]
**AllowedObjectTypes** | [**[]CatalogObjectType**](CatalogObjectType.md) | The set of Catalog Object Types that this Custom Attribute may be applied to. Currently, only &#x60;ITEM&#x60; and &#x60;ITEM_VARIATION&#x60; are allowed. At least one type must be included. See [CatalogObjectType](#type-catalogobjecttype) for possible values | [default to null]
**SellerVisibility** | [***CatalogCustomAttributeDefinitionSellerVisibility**](CatalogCustomAttributeDefinitionSellerVisibility.md) |  | [optional] [default to null]
**AppVisibility** | [***CatalogCustomAttributeDefinitionAppVisibility**](CatalogCustomAttributeDefinitionAppVisibility.md) |  | [optional] [default to null]
**StringConfig** | [***CatalogCustomAttributeDefinitionStringConfig**](CatalogCustomAttributeDefinitionStringConfig.md) |  | [optional] [default to null]
**NumberConfig** | [***CatalogCustomAttributeDefinitionNumberConfig**](CatalogCustomAttributeDefinitionNumberConfig.md) |  | [optional] [default to null]
**SelectionConfig** | [***CatalogCustomAttributeDefinitionSelectionConfig**](CatalogCustomAttributeDefinitionSelectionConfig.md) |  | [optional] [default to null]
**CustomAttributeUsageCount** | **int32** | The number of custom attributes that reference this custom attribute definition. Set by the server in response to a ListCatalog request with &#x60;include_counts&#x60; set to &#x60;true&#x60;.  If the actual count is greater than 100, &#x60;custom_attribute_usage_count&#x60; will be set to &#x60;100&#x60;. | [optional] [default to null]
**Key** | **string** | The name of the desired custom attribute key that can be used to access the custom attribute value on catalog objects. Cannot be modified after the custom attribute definition has been created. Must be between 1 and 60 characters, and may only contain the characters &#x60;[a-zA-Z0-9_-]&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

