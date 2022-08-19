# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayout**](PayoutsApi.md#GetPayout) | **Get** /v2/payouts/{payout_id} | GetPayout
[**ListPayoutEntries**](PayoutsApi.md#ListPayoutEntries) | **Get** /v2/payouts/{payout_id}/payout-entries | ListPayoutEntries
[**ListPayouts**](PayoutsApi.md#ListPayouts) | **Get** /v2/payouts | ListPayouts

# **GetPayout**
> GetPayoutResponse GetPayout(ctx, payoutId)
GetPayout

Retrieves details of a specific payout identified by a payout ID.  To call this endpoint, set `PAYOUTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payoutId** | **string**| The ID of the payout to retrieve the information for. | 

### Return type

[**GetPayoutResponse**](GetPayoutResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPayoutEntries**
> ListPayoutEntriesResponse ListPayoutEntries(ctx, payoutId, optional)
ListPayoutEntries

Retrieves a list of all payout entries for a specific payout. To call this endpoint, set `PAYOUTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **payoutId** | **string**| The ID of the payout to retrieve the information for. | 
 **optional** | ***PayoutsApiListPayoutEntriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayoutsApiListPayoutEntriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortOrder** | [**optional.Interface of SortOrder**](.md)| The order in which payout entries are listed. | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). If request parameters change between requests, subsequent results may contain duplicates or missing records. | 
 **limit** | **optional.Int32**| The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page. The default value of 100 is also the maximum allowed value. If the provided value is greater than 100, it is ignored and the default value is used instead. Default: &#x60;100&#x60; | 

### Return type

[**ListPayoutEntriesResponse**](ListPayoutEntriesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPayouts**
> ListPayoutsResponse ListPayouts(ctx, optional)
ListPayouts

Retrieves a list of all payouts for the default location.  You can filter payouts by location ID, status, time range, and order them in ascending or descending order.  To call this endpoint, set `PAYOUTS_READ` for the OAuth scope.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PayoutsApiListPayoutsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PayoutsApiListPayoutsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationId** | **optional.String**| The ID of the location for which to list the payouts.  By default, payouts are returned for the default (main) location associated with the seller. | 
 **status** | [**optional.Interface of PayoutStatus**](.md)| If provided, only payouts with the given status are returned. | 
 **beginTime** | **optional.String**| The timestamp for the beginning of the payout creation time, in RFC 3339 format. Inclusive. Default: The current time minus one year. | 
 **endTime** | **optional.String**| The timestamp for the end of the payout creation time, in RFC 3339 format. Default: The current time. | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)| The order in which payouts are listed. | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). If request parameters change between requests, subsequent results may contain duplicates or missing records. | 
 **limit** | **optional.Int32**| The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page. The default value of 100 is also the maximum allowed value. If the provided value is greater than 100, it is ignored and the default value is used instead. Default: &#x60;100&#x60; | 

### Return type

[**ListPayoutsResponse**](ListPayoutsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

