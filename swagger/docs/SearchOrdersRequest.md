# SearchOrdersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationIds** | **[]string** | The location IDs for the orders to query. All locations must belong to the same merchant.  Min: 1 location IDs.  Max: 10 location IDs. | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query. See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information. | [optional] [default to null]
**Query** | [***SearchOrdersQuery**](SearchOrdersQuery.md) |  | [optional] [default to null]
**Limit** | **int32** | Maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page.  Default: &#x60;500&#x60; | [optional] [default to null]
**ReturnEntries** | **bool** | Boolean that controls the format of the search results. If &#x60;true&#x60;, SearchOrders will return [&#x60;OrderEntry&#x60;](#type-orderentry) objects. If &#x60;false&#x60;, SearchOrders will return complete Order objects.  Default: &#x60;false&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

