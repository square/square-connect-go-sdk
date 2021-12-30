# CustomAttributeFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomAttributeDefinitionId** | **string** | A query expression to filter items or item variations by matching their custom attributes&#x27; &#x60;custom_attribute_definition_id&#x60; property value against the the specified id. Exactly one of &#x60;custom_attribute_definition_id&#x60; or &#x60;key&#x60; must be specified. | [optional] [default to null]
**Key** | **string** | A query expression to filter items or item variations by matching their custom attributes&#x27; &#x60;key&#x60; property value against the specified key. Exactly one of &#x60;custom_attribute_definition_id&#x60; or &#x60;key&#x60; must be specified. | [optional] [default to null]
**StringFilter** | **string** | A query expression to filter items or item variations by matching their custom attributes&#x27; &#x60;string_value&#x60;  property value against the specified text. Exactly one of &#x60;string_filter&#x60;, &#x60;number_filter&#x60;, &#x60;selection_uids_filter&#x60;, or &#x60;bool_filter&#x60; must be specified. | [optional] [default to null]
**NumberFilter** | [***ModelRange**](Range.md) |  | [optional] [default to null]
**SelectionUidsFilter** | **[]string** | A query expression to filter items or item variations by matching  their custom attributes&#x27; &#x60;selection_uid_values&#x60; values against the specified selection uids. Exactly one of &#x60;string_filter&#x60;, &#x60;number_filter&#x60;, &#x60;selection_uids_filter&#x60;, or &#x60;bool_filter&#x60; must be specified. | [optional] [default to null]
**BoolFilter** | **bool** | A query expression to filter items or item variations by matching their custom attributes&#x27; &#x60;boolean_value&#x60; property values against the specified Boolean expression. Exactly one of &#x60;string_filter&#x60;, &#x60;number_filter&#x60;, &#x60;selection_uids_filter&#x60;, or &#x60;bool_filter&#x60; must be specified. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

