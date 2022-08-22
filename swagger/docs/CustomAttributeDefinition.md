# CustomAttributeDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The identifier of the custom attribute definition and its corresponding custom attributes. This value can be a simple key, which is the key that is provided when the custom attribute definition is created, or a qualified key, if the requesting application is not the definition owner. The qualified key consists of the application ID of the custom attribute definition owner followed by the simple key that was provided when the definition was created. It has the format application_id:simple key.  The value for a simple key can contain up to 60 alphanumeric characters, periods (.), underscores (_), and hyphens (-).  This field can not be changed after the custom attribute definition is created. This field is required when creating a definition and must be unique per application, seller, and resource type. | [optional] [default to null]
**Schema** | [***interface{}**](interface{}.md) | The JSON schema for the custom attribute definition. For more information about the schema, see [Custom Attributes Overview](https://developer.squareup.com/docs/devtools/customattributes/overview). | [optional] [default to null]
**Name** | **string** | The name of the custom attribute definition for API and seller-facing UI purposes. The name must be unique within the seller and application pair. This field is required if the &#x60;visibility&#x60; field is &#x60;VISIBILITY_READ_ONLY&#x60; or &#x60;VISIBILITY_READ_WRITE_VALUES&#x60;. | [optional] [default to null]
**Description** | **string** | Seller-oriented description of the custom attribute definition, including any constraints that the seller should observe. May be displayed as a tooltip in Square UIs. This field is required if the &#x60;visibility&#x60; field is &#x60;VISIBILITY_READ_ONLY&#x60; or &#x60;VISIBILITY_READ_WRITE_VALUES&#x60;. | [optional] [default to null]
**Visibility** | [***CustomAttributeDefinitionVisibility**](CustomAttributeDefinitionVisibility.md) |  | [optional] [default to null]
**Version** | **int32** | Read only. The current version of the custom attribute definition. The value is incremented each time the custom attribute definition is updated. When updating a custom attribute definition, you can provide this field and specify the current version of the custom attribute definition to enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency).  On writes, this field must be set to the latest version. Stale writes are rejected.  This field can also be used to enforce strong consistency for reads. For more information about strong consistency for reads, see [Custom Attributes Overview](https://developer.squareup.com/docs/devtools/customattributes/overview). | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp that indicates when the custom attribute definition was created or most recently updated, in RFC 3339 format. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp that indicates when the custom attribute definition was created, in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
