# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchChangeInventory**](InventoryApi.md#BatchChangeInventory) | **Post** /v2/inventory/batch-change | BatchChangeInventory
[**BatchRetrieveInventoryChanges**](InventoryApi.md#BatchRetrieveInventoryChanges) | **Post** /v2/inventory/batch-retrieve-changes | BatchRetrieveInventoryChanges
[**BatchRetrieveInventoryCounts**](InventoryApi.md#BatchRetrieveInventoryCounts) | **Post** /v2/inventory/batch-retrieve-counts | BatchRetrieveInventoryCounts
[**RetrieveInventoryAdjustment**](InventoryApi.md#RetrieveInventoryAdjustment) | **Get** /v2/inventory/adjustment/{adjustment_id} | RetrieveInventoryAdjustment
[**RetrieveInventoryChanges**](InventoryApi.md#RetrieveInventoryChanges) | **Get** /v2/inventory/{catalog_object_id}/changes | RetrieveInventoryChanges
[**RetrieveInventoryCount**](InventoryApi.md#RetrieveInventoryCount) | **Get** /v2/inventory/{catalog_object_id} | RetrieveInventoryCount
[**RetrieveInventoryPhysicalCount**](InventoryApi.md#RetrieveInventoryPhysicalCount) | **Get** /v2/inventory/physical-count/{physical_count_id} | RetrieveInventoryPhysicalCount

# **BatchChangeInventory**
> BatchChangeInventoryResponse BatchChangeInventory(ctx, body)
BatchChangeInventory

Applies adjustments and counts to the provided item quantities.  On success: returns the current calculated counts for all objects referenced in the request. On failure: returns a list of related errors.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchChangeInventoryRequest**](BatchChangeInventoryRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchChangeInventoryResponse**](BatchChangeInventoryResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchRetrieveInventoryChanges**
> BatchRetrieveInventoryChangesResponse BatchRetrieveInventoryChanges(ctx, body)
BatchRetrieveInventoryChanges

Returns historical physical counts and adjustments based on the provided filter criteria.  Results are paginated and sorted in ascending order according their `occurred_at` timestamp (oldest first).  BatchRetrieveInventoryChanges is a catch-all query endpoint for queries that cannot be handled by other, simpler endpoints.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchRetrieveInventoryChangesRequest**](BatchRetrieveInventoryChangesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchRetrieveInventoryChangesResponse**](BatchRetrieveInventoryChangesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BatchRetrieveInventoryCounts**
> BatchRetrieveInventoryCountsResponse BatchRetrieveInventoryCounts(ctx, body)
BatchRetrieveInventoryCounts

Returns current counts for the provided [CatalogObject](entity:CatalogObject)s at the requested [Location](entity:Location)s.  Results are paginated and sorted in descending order according to their `calculated_at` timestamp (newest first).  When `updated_after` is specified, only counts that have changed since that time (based on the server timestamp for the most recent change) are returned. This allows clients to perform a \"sync\" operation, for example in response to receiving a Webhook notification.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchRetrieveInventoryCountsRequest**](BatchRetrieveInventoryCountsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchRetrieveInventoryCountsResponse**](BatchRetrieveInventoryCountsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveInventoryAdjustment**
> RetrieveInventoryAdjustmentResponse RetrieveInventoryAdjustment(ctx, adjustmentId)
RetrieveInventoryAdjustment

Returns the [InventoryAdjustment](entity:InventoryAdjustment) object with the provided `adjustment_id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **adjustmentId** | **string**| ID of the [InventoryAdjustment](entity:InventoryAdjustment) to retrieve. | 

### Return type

[**RetrieveInventoryAdjustmentResponse**](RetrieveInventoryAdjustmentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveInventoryChanges**
> RetrieveInventoryChangesResponse RetrieveInventoryChanges(ctx, catalogObjectId, optional)
RetrieveInventoryChanges

Returns a set of physical counts and inventory adjustments for the provided [CatalogObject](entity:CatalogObject) at the requested [Location](entity:Location)s.  Results are paginated and sorted in descending order according to their `occurred_at` timestamp (newest first).  There are no limits on how far back the caller can page. This endpoint can be  used to display recent changes for a specific item. For more sophisticated queries, use a batch endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **catalogObjectId** | **string**| ID of the [CatalogObject](entity:CatalogObject) to retrieve. | 
 **optional** | ***InventoryApiRetrieveInventoryChangesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryApiRetrieveInventoryChangesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationIds** | **optional.String**| The [Location](entity:Location) IDs to look up as a comma-separated list. An empty list queries all locations. | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information. | 

### Return type

[**RetrieveInventoryChangesResponse**](RetrieveInventoryChangesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveInventoryCount**
> RetrieveInventoryCountResponse RetrieveInventoryCount(ctx, catalogObjectId, optional)
RetrieveInventoryCount

Retrieves the current calculated stock count for a given [CatalogObject](entity:CatalogObject) at a given set of [Location](entity:Location)s. Responses are paginated and unsorted. For more sophisticated queries, use a batch endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **catalogObjectId** | **string**| ID of the [CatalogObject](entity:CatalogObject) to retrieve. | 
 **optional** | ***InventoryApiRetrieveInventoryCountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a InventoryApiRetrieveInventoryCountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **locationIds** | **optional.String**| The [Location](entity:Location) IDs to look up as a comma-separated list. An empty list queries all locations. | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information. | 

### Return type

[**RetrieveInventoryCountResponse**](RetrieveInventoryCountResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveInventoryPhysicalCount**
> RetrieveInventoryPhysicalCountResponse RetrieveInventoryPhysicalCount(ctx, physicalCountId)
RetrieveInventoryPhysicalCount

Returns the [InventoryPhysicalCount](entity:InventoryPhysicalCount) object with the provided `physical_count_id`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **physicalCountId** | **string**| ID of the [InventoryPhysicalCount](entity:InventoryPhysicalCount) to retrieve. | 

### Return type

[**RetrieveInventoryPhysicalCountResponse**](RetrieveInventoryPhysicalCountResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

