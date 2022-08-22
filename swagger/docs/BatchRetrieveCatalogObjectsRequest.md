# BatchRetrieveCatalogObjectsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectIds** | **[]string** | The IDs of the CatalogObjects to be retrieved. | [default to null]
**IncludeRelatedObjects** | **bool** | If &#x60;true&#x60;, the response will include additional objects that are related to the requested objects. Related objects are defined as any objects referenced by ID by the results in the &#x60;objects&#x60; field of the response. These objects are put in the &#x60;related_objects&#x60; field. Setting this to &#x60;true&#x60; is helpful when the objects are needed for immediate display to a user. This process only goes one level deep. Objects referenced by the related objects will not be included. For example,  if the &#x60;objects&#x60; field of the response contains a CatalogItem, its associated CatalogCategory objects, CatalogTax objects, CatalogImage objects and CatalogModifierLists will be returned in the &#x60;related_objects&#x60; field of the response. If the &#x60;objects&#x60; field of the response contains a CatalogItemVariation, its parent CatalogItem will be returned in the &#x60;related_objects&#x60; field of the response.  Default value: &#x60;false&#x60; | [optional] [default to null]
**CatalogVersion** | **int64** | The specific version of the catalog objects to be included in the response.  This allows you to retrieve historical versions of objects. The specified version value is matched against the [CatalogObject](entity:CatalogObject)s&#x27; &#x60;version&#x60; attribute. If not included, results will be from the current version of the catalog. | [optional] [default to null]
**IncludeDeletedObjects** | **bool** | Indicates whether to include (&#x60;true&#x60;) or not (&#x60;false&#x60;) in the response deleted objects, namely, those with the &#x60;is_deleted&#x60; attribute set to &#x60;true&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

