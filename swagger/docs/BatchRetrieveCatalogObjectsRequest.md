# BatchRetrieveCatalogObjectsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectIds** | **[]string** | The IDs of the CatalogObjects to be retrieved. | [default to null]
**IncludeRelatedObjects** | **bool** | If &#x60;true&#x60;, the response will include additional objects that are related to the requested objects, as follows:  If the &#x60;objects&#x60; field of the response contains a CatalogItem, its associated CatalogCategory objects, CatalogTax objects, CatalogImage objects and CatalogModifierLists will be returned in the &#x60;related_objects&#x60; field of the response. If the &#x60;objects&#x60; field of the response contains a CatalogItemVariation, its parent CatalogItem will be returned in the &#x60;related_objects&#x60; field of the response. | [optional] [default to null]
**CatalogVersion** | **int64** | The specific version of the catalog objects to be included in the response.  This allows you to retrieve historical versions of objects. The specified version value is matched against the [CatalogObject](entity:CatalogObject)s&#x27; &#x60;version&#x60; attribute. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

