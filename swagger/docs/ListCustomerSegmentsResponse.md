# ListCustomerSegmentsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Segments** | [**[]CustomerSegment**](CustomerSegment.md) | The list of customer segments belonging to the associated Square account. | [optional] [default to null]
**Cursor** | **string** | A pagination cursor to be used in subsequent calls to &#x60;ListCustomerSegments&#x60; to retrieve the next set of query results. The cursor is only present if the request succeeded and additional results are available.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

