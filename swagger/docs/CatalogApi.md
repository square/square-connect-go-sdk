# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchDeleteCatalogObjects**](CatalogApi.md#BatchDeleteCatalogObjects) | **Post** /v2/catalog/batch-delete | BatchDeleteCatalogObjects
[**BatchRetrieveCatalogObjects**](CatalogApi.md#BatchRetrieveCatalogObjects) | **Post** /v2/catalog/batch-retrieve | BatchRetrieveCatalogObjects
[**BatchUpsertCatalogObjects**](CatalogApi.md#BatchUpsertCatalogObjects) | **Post** /v2/catalog/batch-upsert | BatchUpsertCatalogObjects
[**CatalogInfo**](CatalogApi.md#CatalogInfo) | **Get** /v2/catalog/info | CatalogInfo
[**CreateCatalogImage**](CatalogApi.md#CreateCatalogImage) | **Post** /v2/catalog/images | CreateCatalogImage
[**DeleteCatalogObject**](CatalogApi.md#DeleteCatalogObject) | **Delete** /v2/catalog/object/{object_id} | DeleteCatalogObject
[**ListCatalog**](CatalogApi.md#ListCatalog) | **Get** /v2/catalog/list | ListCatalog
[**RetrieveCatalogObject**](CatalogApi.md#RetrieveCatalogObject) | **Get** /v2/catalog/object/{object_id} | RetrieveCatalogObject
[**SearchCatalogItems**](CatalogApi.md#SearchCatalogItems) | **Post** /v2/catalog/search-catalog-items | SearchCatalogItems
[**SearchCatalogObjects**](CatalogApi.md#SearchCatalogObjects) | **Post** /v2/catalog/search | SearchCatalogObjects
[**UpdateCatalogImage**](CatalogApi.md#UpdateCatalogImage) | **Put** /v2/catalog/images/{image_id} | UpdateCatalogImage
[**UpdateItemModifierLists**](CatalogApi.md#UpdateItemModifierLists) | **Post** /v2/catalog/update-item-modifier-lists | UpdateItemModifierLists
[**UpdateItemTaxes**](CatalogApi.md#UpdateItemTaxes) | **Post** /v2/catalog/update-item-taxes | UpdateItemTaxes
[**UpsertCatalogObject**](CatalogApi.md#UpsertCatalogObject) | **Post** /v2/catalog/object | UpsertCatalogObject

# **BatchDeleteCatalogObjects**
> BatchDeleteCatalogObjectsResponse BatchDeleteCatalogObjects(ctx, body)
BatchDeleteCatalogObjects

Deletes a set of [CatalogItem](entity:CatalogItem)s based on the provided list of target IDs and returns a set of successfully deleted IDs in the response. Deletion is a cascading event such that all children of the targeted object are also deleted. For example, deleting a CatalogItem will also delete all of its [CatalogItemVariation](entity:CatalogItemVariation) children.  `BatchDeleteCatalogObjects` succeeds even if only a portion of the targeted IDs can be deleted. The response will only include IDs that were actually deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchDeleteCatalogObjectsRequest**](BatchDeleteCatalogObjectsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchDeleteCatalogObjectsResponse**](BatchDeleteCatalogObjectsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchRetrieveCatalogObjects**
> BatchRetrieveCatalogObjectsResponse BatchRetrieveCatalogObjects(ctx, body)
BatchRetrieveCatalogObjects

Returns a set of objects based on the provided ID. Each [CatalogItem](entity:CatalogItem) returned in the set includes all of its child information including: all of its [CatalogItemVariation](entity:CatalogItemVariation) objects, references to its [CatalogModifierList](entity:CatalogModifierList) objects, and the ids of any [CatalogTax](entity:CatalogTax) objects that apply to it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchRetrieveCatalogObjectsRequest**](BatchRetrieveCatalogObjectsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchRetrieveCatalogObjectsResponse**](BatchRetrieveCatalogObjectsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchUpsertCatalogObjects**
> BatchUpsertCatalogObjectsResponse BatchUpsertCatalogObjects(ctx, body)
BatchUpsertCatalogObjects

Creates or updates up to 10,000 target objects based on the provided list of objects. The target objects are grouped into batches and each batch is inserted/updated in an all-or-nothing manner. If an object within a batch is malformed in some way, or violates a database constraint, the entire batch containing that item will be disregarded. However, other batches in the same request may still succeed. Each batch may contain up to 1,000 objects, and batches will be processed in order as long as the total object count for the request (items, variations, modifier lists, discounts, and taxes) is no more than 10,000.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchUpsertCatalogObjectsRequest**](BatchUpsertCatalogObjectsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchUpsertCatalogObjectsResponse**](BatchUpsertCatalogObjectsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CatalogInfo**
> CatalogInfoResponse CatalogInfo(ctx, )
CatalogInfo

Retrieves information about the Square Catalog API, such as batch size limits that can be used by the `BatchUpsertCatalogObjects` endpoint.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**CatalogInfoResponse**](CatalogInfoResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCatalogImage**
> CreateCatalogImageResponse CreateCatalogImage(ctx, request, imageFile)
CreateCatalogImage

Uploads an image file to be represented by a [CatalogImage](entity:CatalogImage) object that can be linked to an existing [CatalogObject](entity:CatalogObject) instance. The resulting `CatalogImage` is unattached to any `CatalogObject` if the `object_id` is not specified.  This `CreateCatalogImage` endpoint accepts HTTP multipart/form-data requests with a JSON part and an image file part in JPEG, PJPEG, PNG, or GIF format. The maximum file size is 15MB.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**CreateCatalogImageRequest**](.md)|  | 
  **imageFile** | ***os.File*****os.File**|  | 

### Return type

[**CreateCatalogImageResponse**](CreateCatalogImageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCatalogObject**
> DeleteCatalogObjectResponse DeleteCatalogObject(ctx, objectId)
DeleteCatalogObject

Deletes a single [CatalogObject](entity:CatalogObject) based on the provided ID and returns the set of successfully deleted IDs in the response. Deletion is a cascading event such that all children of the targeted object are also deleted. For example, deleting a [CatalogItem](entity:CatalogItem) will also delete all of its [CatalogItemVariation](entity:CatalogItemVariation) children.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The ID of the catalog object to be deleted. When an object is deleted, other objects in the graph that depend on that object will be deleted as well (for example, deleting a catalog item will delete its catalog item variations). | 

### Return type

[**DeleteCatalogObjectResponse**](DeleteCatalogObjectResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCatalog**
> ListCatalogResponse ListCatalog(ctx, optional)
ListCatalog

Returns a list of all [CatalogObject](entity:CatalogObject)s of the specified types in the catalog.   The `types` parameter is specified as a comma-separated list of the [CatalogObjectType](entity:CatalogObjectType) values,  for example, \"`ITEM`, `ITEM_VARIATION`, `MODIFIER`, `MODIFIER_LIST`, `CATEGORY`, `DISCOUNT`, `TAX`, `IMAGE`\".  __Important:__ ListCatalog does not return deleted catalog items. To retrieve deleted catalog items, use [SearchCatalogObjects](api-endpoint:Catalog-SearchCatalogObjects) and set the `include_deleted_objects` attribute value to `true`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CatalogApiListCatalogOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiListCatalogOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| The pagination cursor returned in the previous response. Leave unset for an initial request. The page size is currently set to be 100. See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information. | 
 **types** | **optional.String**| An optional case-insensitive, comma-separated list of object types to retrieve.  The valid values are defined in the [CatalogObjectType](entity:CatalogObjectType) enum, for example, &#x60;ITEM&#x60;, &#x60;ITEM_VARIATION&#x60;, &#x60;CATEGORY&#x60;, &#x60;DISCOUNT&#x60;, &#x60;TAX&#x60;, &#x60;MODIFIER&#x60;, &#x60;MODIFIER_LIST&#x60;, &#x60;IMAGE&#x60;, etc.  If this is unspecified, the operation returns objects of all the top level types at the version of the Square API used to make the request. Object types that are nested onto other object types are not included in the defaults.  At the current API version the default object types are: ITEM, CATEGORY, TAX, DISCOUNT, MODIFIER_LIST,  PRICING_RULE, PRODUCT_SET, TIME_PERIOD, MEASUREMENT_UNIT, SUBSCRIPTION_PLAN, ITEM_OPTION, CUSTOM_ATTRIBUTE_DEFINITION, QUICK_AMOUNT_SETTINGS. | 
 **catalogVersion** | **optional.Int64**| The specific version of the catalog objects to be included in the response.  This allows you to retrieve historical versions of objects. The specified version value is matched against the [CatalogObject](entity:CatalogObject)s&#x27; &#x60;version&#x60; attribute.  If not included, results will be from the current version of the catalog. | 

### Return type

[**ListCatalogResponse**](ListCatalogResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCatalogObject**
> RetrieveCatalogObjectResponse RetrieveCatalogObject(ctx, objectId, optional)
RetrieveCatalogObject

Returns a single [CatalogItem](entity:CatalogItem) as a [CatalogObject](entity:CatalogObject) based on the provided ID. The returned object includes all of the relevant [CatalogItem](entity:CatalogItem) information including: [CatalogItemVariation](entity:CatalogItemVariation) children, references to its [CatalogModifierList](entity:CatalogModifierList) objects, and the ids of any [CatalogTax](entity:CatalogTax) objects that apply to it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **objectId** | **string**| The object ID of any type of catalog objects to be retrieved. | 
 **optional** | ***CatalogApiRetrieveCatalogObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CatalogApiRetrieveCatalogObjectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeRelatedObjects** | **optional.Bool**| If &#x60;true&#x60;, the response will include additional objects that are related to the requested objects. Related objects are defined as any objects referenced by ID by the results in the &#x60;objects&#x60; field of the response. These objects are put in the &#x60;related_objects&#x60; field. Setting this to &#x60;true&#x60; is helpful when the objects are needed for immediate display to a user. This process only goes one level deep. Objects referenced by the related objects will not be included. For example,  if the &#x60;objects&#x60; field of the response contains a CatalogItem, its associated CatalogCategory objects, CatalogTax objects, CatalogImage objects and CatalogModifierLists will be returned in the &#x60;related_objects&#x60; field of the response. If the &#x60;objects&#x60; field of the response contains a CatalogItemVariation, its parent CatalogItem will be returned in the &#x60;related_objects&#x60; field of the response.  Default value: &#x60;false&#x60; | [default to false]
 **catalogVersion** | **optional.Int64**| Requests objects as of a specific version of the catalog. This allows you to retrieve historical versions of objects. The value to retrieve a specific version of an object can be found in the version field of [CatalogObject](entity:CatalogObject)s. If not included, results will be from the current version of the catalog. | 

### Return type

[**RetrieveCatalogObjectResponse**](RetrieveCatalogObjectResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCatalogItems**
> SearchCatalogItemsResponse SearchCatalogItems(ctx, body)
SearchCatalogItems

Searches for catalog items or item variations by matching supported search attribute values, including custom attribute values, against one or more of the specified query filters.  This (`SearchCatalogItems`) endpoint differs from the [SearchCatalogObjects](api-endpoint:Catalog-SearchCatalogObjects) endpoint in the following aspects:  - `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects. - `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not. - `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does. - The both endpoints use different call conventions, including the query filter formats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCatalogItemsRequest**](SearchCatalogItemsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchCatalogItemsResponse**](SearchCatalogItemsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCatalogObjects**
> SearchCatalogObjectsResponse SearchCatalogObjects(ctx, body)
SearchCatalogObjects

Searches for [CatalogObject](entity:CatalogObject) of any type by matching supported search attribute values, excluding custom attribute values on items or item variations, against one or more of the specified query filters.  This (`SearchCatalogObjects`) endpoint differs from the [SearchCatalogItems](api-endpoint:Catalog-SearchCatalogItems) endpoint in the following aspects:  - `SearchCatalogItems` can only search for items or item variations, whereas `SearchCatalogObjects` can search for any type of catalog objects. - `SearchCatalogItems` supports the custom attribute query filters to return items or item variations that contain custom attribute values, where `SearchCatalogObjects` does not. - `SearchCatalogItems` does not support the `include_deleted_objects` filter to search for deleted items or item variations, whereas `SearchCatalogObjects` does. - The both endpoints have different call conventions, including the query filter formats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCatalogObjectsRequest**](SearchCatalogObjectsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchCatalogObjectsResponse**](SearchCatalogObjectsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCatalogImage**
> UpdateCatalogImageResponse UpdateCatalogImage(ctx, request, imageFile, imageId)
UpdateCatalogImage

Uploads a new image file to replace the existing one in the specified [CatalogImage](entity:CatalogImage) object.   This `UpdateCatalogImage` endpoint accepts HTTP multipart/form-data requests with a JSON part and an image file part in JPEG, PJPEG, PNG, or GIF format. The maximum file size is 15MB.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**UpdateCatalogImageRequest**](.md)|  | 
  **imageFile** | ***os.File*****os.File**|  | 
  **imageId** | **string**| The ID of the &#x60;CatalogImage&#x60; object to update the encapsulated image file. | 

### Return type

[**UpdateCatalogImageResponse**](UpdateCatalogImageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateItemModifierLists**
> UpdateItemModifierListsResponse UpdateItemModifierLists(ctx, body)
UpdateItemModifierLists

Updates the [CatalogModifierList](entity:CatalogModifierList) objects that apply to the targeted [CatalogItem](entity:CatalogItem) without having to perform an upsert on the entire item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateItemModifierListsRequest**](UpdateItemModifierListsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**UpdateItemModifierListsResponse**](UpdateItemModifierListsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateItemTaxes**
> UpdateItemTaxesResponse UpdateItemTaxes(ctx, body)
UpdateItemTaxes

Updates the [CatalogTax](entity:CatalogTax) objects that apply to the targeted [CatalogItem](entity:CatalogItem) without having to perform an upsert on the entire item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateItemTaxesRequest**](UpdateItemTaxesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**UpdateItemTaxesResponse**](UpdateItemTaxesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertCatalogObject**
> UpsertCatalogObjectResponse UpsertCatalogObject(ctx, body)
UpsertCatalogObject

Creates or updates the target [CatalogObject](entity:CatalogObject).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertCatalogObjectRequest**](UpsertCatalogObjectRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**UpsertCatalogObjectResponse**](UpsertCatalogObjectResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

