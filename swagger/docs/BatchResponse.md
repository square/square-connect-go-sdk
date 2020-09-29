# BatchResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusCode** | **int32** | HTTP status code for the response | [optional] [default to null]
**Body** | **map[string]string** | The body of the response (if any). | [optional] [default to null]
**Headers** | **map[string]string** | Contains any important headers for the response, indexed by header name. For example, if the response includes a pagination header, the header value is available from &#x60;headers[\&quot;Link\&quot;]&#x60;. | [optional] [default to null]
**RequestId** | **string** | The value provided in a request for &#x60;request_id&#x60; in the corresponding &#x60;BatchRequest&#x60; (if any). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

