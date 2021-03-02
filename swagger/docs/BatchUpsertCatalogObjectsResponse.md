# BatchUpsertCatalogObjectsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Objects** | [**[]CatalogObject**](CatalogObject.md) | The created successfully created CatalogObjects. | [optional] [default to null]
**UpdatedAt** | **string** | The database [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) of this update in RFC 3339 format, e.g., \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**IdMappings** | [**[]CatalogIdMapping**](CatalogIdMapping.md) | The mapping between client and server IDs for this upsert. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

