# CatalogCustomAttributeDefinitionSelectionConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAllowedSelections** | **int32** | The maximum number of selections that can be set. The maximum value for this attribute is 100. The default value is 1. The value can be modified, but changing the value will not affect existing custom attribute values on objects. Clients need to handle custom attributes with more selected values than allowed by this limit. | [optional] [default to null]
**AllowedSelections** | [**[]CatalogCustomAttributeDefinitionSelectionConfigCustomAttributeSelection**](CatalogCustomAttributeDefinitionSelectionConfigCustomAttributeSelection.md) | The set of valid &#x60;CatalogCustomAttributeSelections&#x60;. Up to a maximum of 100 selections can be defined. Can be modified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

