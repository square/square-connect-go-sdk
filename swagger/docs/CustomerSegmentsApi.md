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
 **cursor** | **optional.String**| A pagination cursor returned by previous calls to &#x60;ListCustomerSegments&#x60;. This cursor is used to retrieve the next set of query results.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **limit** | **optional.Int32**| The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results. If the specified limit is less than 1 or greater than 50, Square returns a &#x60;400 VALUE_TOO_LOW&#x60; or &#x60;400 VALUE_TOO_HIGH&#x60; error. The default value is 50.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 

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

