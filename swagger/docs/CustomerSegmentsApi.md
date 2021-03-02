# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCustomerSegments**](CustomerSegmentsApi.md#ListCustomerSegments) | **Get** /v2/customers/segments | ListCustomerSegments
[**RetrieveCustomerSegment**](CustomerSegmentsApi.md#RetrieveCustomerSegment) | **Get** /v2/customers/segments/{segment_id} | RetrieveCustomerSegment

# **ListCustomerSegments**
> ListCustomerSegmentsResponse ListCustomerSegments(ctx, optional)
ListCustomerSegments

Retrieves the list of customer segments of a business.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerSegmentsApiListCustomerSegmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerSegmentsApiListCustomerSegmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| A pagination cursor returned by previous calls to __ListCustomerSegments__. Used to retrieve the next set of query results.  See the [Pagination guide](https://developer.squareup.com/docs/working-with-apis/pagination) for more information. | 

### Return type

[**ListCustomerSegmentsResponse**](ListCustomerSegmentsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCustomerSegment**
> RetrieveCustomerSegmentResponse RetrieveCustomerSegment(ctx, segmentId)
RetrieveCustomerSegment

Retrieves a specific customer segment as identified by the `segment_id` value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **segmentId** | **string**| The Square-issued ID of the customer segment. | 

### Return type

[**RetrieveCustomerSegmentResponse**](RetrieveCustomerSegmentResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

