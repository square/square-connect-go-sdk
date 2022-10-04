# ListDisputeEvidenceResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Evidence** | [**[]DisputeEvidence**](DisputeEvidence.md) | The list of evidence previously uploaded to the specified dispute. | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) | Information about errors encountered during the request. | [optional] [default to null]
**Cursor** | **string** | The pagination cursor to be used in a subsequent request. If unset, this is the final response. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

