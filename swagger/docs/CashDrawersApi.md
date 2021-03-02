# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListCashDrawerShiftEvents**](CashDrawersApi.md#ListCashDrawerShiftEvents) | **Get** /v2/cash-drawers/shifts/{shift_id}/events | ListCashDrawerShiftEvents
[**ListCashDrawerShifts**](CashDrawersApi.md#ListCashDrawerShifts) | **Get** /v2/cash-drawers/shifts | ListCashDrawerShifts
[**RetrieveCashDrawerShift**](CashDrawersApi.md#RetrieveCashDrawerShift) | **Get** /v2/cash-drawers/shifts/{shift_id} | RetrieveCashDrawerShift

# **ListCashDrawerShiftEvents**
> ListCashDrawerShiftEventsResponse ListCashDrawerShiftEvents(ctx, locationId, shiftId, optional)
ListCashDrawerShiftEvents

Provides a paginated list of events for a single cash drawer shift.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list cash drawer shifts for. | 
  **shiftId** | **string**| The shift ID. | 
 **optional** | ***CashDrawersApiListCashDrawerShiftEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CashDrawersApiListCashDrawerShiftEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| Number of resources to be returned in a page of results (200 by default, 1000 max). | 
 **cursor** | **optional.String**| Opaque cursor for fetching the next page of results. | 

### Return type

[**ListCashDrawerShiftEventsResponse**](ListCashDrawerShiftEventsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCashDrawerShifts**
> ListCashDrawerShiftsResponse ListCashDrawerShifts(ctx, locationId, optional)
ListCashDrawerShifts

Provides the details for all of the cash drawer shifts for a location in a date range.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to query for a list of cash drawer shifts. | 
 **optional** | ***CashDrawersApiListCashDrawerShiftsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CashDrawersApiListCashDrawerShiftsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sortOrder** | [**optional.Interface of SortOrder**](.md)| The order in which cash drawer shifts are listed in the response, based on their opened_at field. Default value: ASC | 
 **beginTime** | **optional.String**| The inclusive start time of the query on opened_at, in ISO 8601 format. | 
 **endTime** | **optional.String**| The exclusive end date of the query on opened_at, in ISO 8601 format. | 
 **limit** | **optional.Int32**| Number of cash drawer shift events in a page of results (200 by default, 1000 max). | 
 **cursor** | **optional.String**| Opaque cursor for fetching the next page of results. | 

### Return type

[**ListCashDrawerShiftsResponse**](ListCashDrawerShiftsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCashDrawerShift**
> RetrieveCashDrawerShiftResponse RetrieveCashDrawerShift(ctx, locationId, shiftId)
RetrieveCashDrawerShift

Provides the summary details for a single cash drawer shift. See [ListCashDrawerShiftEvents](#endpoint-CashDrawers-ListCashDrawerShiftEvents) for a list of cash drawer shift events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to retrieve cash drawer shifts from. | 
  **shiftId** | **string**| The shift ID. | 

### Return type

[**RetrieveCashDrawerShiftResponse**](RetrieveCashDrawerShiftResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

