# CreateCatalogImageRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** | A unique string that identifies this CreateCatalogImage request. Keys can be any valid string but must be unique for every CreateCatalogImage request.  See [Idempotency keys](https://developer.squareup.com/docs/basics/api101/idempotency) for more information. | [default to null]
**ObjectId** | **string** | Unique ID of the &#x60;CatalogObject&#x60; to attach this &#x60;CatalogImage&#x60; object to. Leave this field empty to create unattached images, for example if you are building an integration where an image can be attached to catalog items at a later time. | [optional] [default to null]
**Image** | [***CatalogObject**](CatalogObject.md) |  | [default to null]
**IsPrimary** | **bool** | If this is set to &#x60;true&#x60;, the image created will be the primary, or first image of the object referenced by &#x60;object_id&#x60;. If the &#x60;CatalogObject&#x60; already has a primary &#x60;CatalogImage&#x60;, setting this field to &#x60;true&#x60; will replace the primary image. If this is set to &#x60;false&#x60; and you use the Square API version 2021-12-15 or later, the image id will be appended to the list of &#x60;image_ids&#x60; on the object.  With Square API version 2021-12-15 or later, the default value is &#x60;false&#x60;. Otherwise, the effective default value is &#x60;true&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

