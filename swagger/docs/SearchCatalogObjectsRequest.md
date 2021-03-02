# SearchCatalogObjectsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | The pagination cursor returned in the previous response. Leave unset for an initial request. See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information. | [optional] [default to null]
**ObjectTypes** | [**[]CatalogObjectType**](CatalogObjectType.md) | The desired set of object types to appear in the search results. | [optional] [default to null]
**IncludeDeletedObjects** | **bool** | If &#x60;true&#x60;, deleted objects will be included in the results. Deleted objects will have their &#x60;is_deleted&#x60; field set to &#x60;true&#x60;. | [optional] [default to null]
**IncludeRelatedObjects** | **bool** | If &#x60;true&#x60;, the response will include additional objects that are related to the requested object, as follows:  If a CatalogItem is returned in the object field of the response, its associated CatalogCategory, CatalogTax objects, CatalogImage objects and CatalogModifierList objects will be included in the &#x60;related_objects&#x60; field of the response.  If a CatalogItemVariation is returned in the object field of the response, its parent CatalogItem will be included in the &#x60;related_objects&#x60; field of the response. | [optional] [default to null]
**BeginTime** | **string** | Return objects modified after this [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates), in RFC 3339 format, e.g., &#x60;2016-09-04T23:59:33.123Z&#x60;. The timestamp is exclusive - objects with a timestamp equal to &#x60;begin_time&#x60; will not be included in the response. | [optional] [default to null]
**Query** | [***CatalogQuery**](CatalogQuery.md) |  | [optional] [default to null]
**Limit** | **int32** | A limit on the number of results to be returned in a single page. The limit is advisory - the implementation may return more or fewer results. If the supplied limit is negative, zero, or is higher than the maximum limit of 1,000, it will be ignored. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

