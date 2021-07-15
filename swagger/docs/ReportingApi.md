# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAdditionalRecipientReceivableRefunds**](ReportingApi.md#ListAdditionalRecipientReceivableRefunds) | **Get** /v2/locations/{location_id}/additional-recipient-receivable-refunds | ListAdditionalRecipientReceivableRefunds
[**ListAdditionalRecipientReceivables**](ReportingApi.md#ListAdditionalRecipientReceivables) | **Get** /v2/locations/{location_id}/additional-recipient-receivables | ListAdditionalRecipientReceivables

# **ListAdditionalRecipientReceivableRefunds**
> ListAdditionalRecipientReceivableRefundsResponse ListAdditionalRecipientReceivableRefunds(ctx, locationId, optional)
ListAdditionalRecipientReceivableRefunds

Returns a list of refunded transactions (across all possible originating locations) relating to monies credited to the provided location ID by another Square account using the `additional_recipients` field in a transaction.  Max results per [page](https://developer.squareup.com/docs/working-with-apis/pagination): 50

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list AdditionalRecipientReceivableRefunds for. | 
 **optional** | ***ReportingApiListAdditionalRecipientReceivableRefundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiListAdditionalRecipientReceivableRefundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beginTime** | **optional.String**| The beginning of the requested reporting period, in RFC 3339 format.  See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.  Default value: The current time minus one year. | 
 **endTime** | **optional.String**| The end of the requested reporting period, in RFC 3339 format.  See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.  Default value: The current time. | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)| The order in which results are listed in the response (&#x60;ASC&#x60; for oldest first, &#x60;DESC&#x60; for newest first).  Default value: &#x60;DESC&#x60; | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Paginating results](https://developer.squareup.com/docs/working-with-apis/pagination) for more information. | 

### Return type

[**ListAdditionalRecipientReceivableRefundsResponse**](ListAdditionalRecipientReceivableRefundsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAdditionalRecipientReceivables**
> ListAdditionalRecipientReceivablesResponse ListAdditionalRecipientReceivables(ctx, locationId, optional)
ListAdditionalRecipientReceivables

Returns a list of receivables (across all possible sending locations) representing monies credited to the provided location ID by another Square account using the `additional_recipients` field in a transaction.  Max results per [page](https://developer.squareup.com/docs/working-with-apis/pagination): 50

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list AdditionalRecipientReceivables for. | 
 **optional** | ***ReportingApiListAdditionalRecipientReceivablesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ReportingApiListAdditionalRecipientReceivablesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **beginTime** | **optional.String**| The beginning of the requested reporting period, in RFC 3339 format.  See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.  Default value: The current time minus one year. | 
 **endTime** | **optional.String**| The end of the requested reporting period, in RFC 3339 format.  See [Date ranges](https://developer.squareup.com/docs/build-basics/working-with-dates) for details on date inclusivity/exclusivity.  Default value: The current time. | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)| The order in which results are listed in the response (&#x60;ASC&#x60; for oldest first, &#x60;DESC&#x60; for newest first).  Default value: &#x60;DESC&#x60; | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Paginating results](https://developer.squareup.com/docs/working-with-apis/pagination) for more information. | 

### Return type

[**ListAdditionalRecipientReceivablesResponse**](ListAdditionalRecipientReceivablesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

