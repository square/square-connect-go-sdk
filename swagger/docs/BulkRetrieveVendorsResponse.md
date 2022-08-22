# BulkRetrieveVendorsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Responses** | [**map[string]RetrieveVendorResponse**](RetrieveVendorResponse.md) | The set of [RetrieveVendorResponse](entity:RetrieveVendorResponse) objects encapsulating successfully retrieved [Vendor](entity:Vendor) objects or error responses for failed attempts. The set is represented by  a collection of &#x60;Vendor&#x60;-ID/&#x60;Vendor&#x60;-object or &#x60;Vendor&#x60;-ID/error-object pairs. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

