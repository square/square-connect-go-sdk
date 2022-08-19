# ListCustomerCustomAttributesResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomAttributes** | [**[]CustomAttribute**](CustomAttribute.md) | The retrieved custom attributes. If &#x60;with_definitions&#x60; was set to &#x60;true&#x60; in the request, the custom attribute definition is returned in the &#x60;definition&#x60; field of each custom attribute.  If no custom attributes are found, Square returns an empty object (&#x60;{}&#x60;). | [optional] [default to null]
**Cursor** | **string** | The cursor to use in your next call to this endpoint to retrieve the next page of results for your original request. This field is present only if the request succeeded and additional results are available. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

