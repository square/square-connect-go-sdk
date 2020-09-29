# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1SubmitBatch**](V1BatchingApi.md#V1SubmitBatch) | **Post** /v1/batch | V1SubmitBatch

# **V1SubmitBatch**
> []BatchResponse V1SubmitBatch(ctx, body)
V1SubmitBatch

Bundle multiple requests to Connect V1 API endpoints as a single request.   The __V1SubmitBatch__ endpoint does not require an access token in the request header. Instead, provide an `access_token` parameter for each request included in the batch.  __V1SubmitBatch__ responds with an array that contains response objects for each of the batched requests. There is no guarantee of the order in which batched requests are performed.   __IMPORTANT__  You cannot include more than 30 requests in a single batch and recursive requests to __V1SubmitBatch__ are not allowed. In other words, none of the requests included in a batch can itself be a request to the __V1SubmitBatch__ endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchRequest**](BatchRequest.md)| The set of API actions to perform. | 

### Return type

[**[]BatchResponse**](BatchResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

