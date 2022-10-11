# SearchVendorsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Errors encountered when the request fails. | [optional] [default to null]
**Vendors** | [**[]Vendor**](Vendor.md) | The [Vendor](entity:Vendor) objects matching the specified search filter. | [optional] [default to null]
**Cursor** | **string** | The pagination cursor to be used in a subsequent request. If unset, this is the final response.  See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

