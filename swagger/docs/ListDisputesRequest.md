# ListDisputesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | [optional] [default to null]
**States** | [**[]DisputeState**](DisputeState.md) | The dispute states used to filter the result. If not specified, the endpoint returns all disputes. See [DisputeState](#type-disputestate) for possible values | [optional] [default to null]
**LocationId** | **string** | The ID of the location for which to return a list of disputes. If not specified, the endpoint returns disputes associated with all locations. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

