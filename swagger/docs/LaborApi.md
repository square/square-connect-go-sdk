# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBreakType**](LaborApi.md#CreateBreakType) | **Post** /v2/labor/break-types | CreateBreakType
[**CreateShift**](LaborApi.md#CreateShift) | **Post** /v2/labor/shifts | CreateShift
[**DeleteBreakType**](LaborApi.md#DeleteBreakType) | **Delete** /v2/labor/break-types/{id} | DeleteBreakType
[**DeleteShift**](LaborApi.md#DeleteShift) | **Delete** /v2/labor/shifts/{id} | DeleteShift
[**GetBreakType**](LaborApi.md#GetBreakType) | **Get** /v2/labor/break-types/{id} | GetBreakType
[**GetEmployeeWage**](LaborApi.md#GetEmployeeWage) | **Get** /v2/labor/employee-wages/{id} | GetEmployeeWage
[**GetShift**](LaborApi.md#GetShift) | **Get** /v2/labor/shifts/{id} | GetShift
[**GetTeamMemberWage**](LaborApi.md#GetTeamMemberWage) | **Get** /v2/labor/team-member-wages/{id} | GetTeamMemberWage
[**ListBreakTypes**](LaborApi.md#ListBreakTypes) | **Get** /v2/labor/break-types | ListBreakTypes
[**ListEmployeeWages**](LaborApi.md#ListEmployeeWages) | **Get** /v2/labor/employee-wages | ListEmployeeWages
[**ListTeamMemberWages**](LaborApi.md#ListTeamMemberWages) | **Get** /v2/labor/team-member-wages | ListTeamMemberWages
[**ListWorkweekConfigs**](LaborApi.md#ListWorkweekConfigs) | **Get** /v2/labor/workweek-configs | ListWorkweekConfigs
[**SearchShifts**](LaborApi.md#SearchShifts) | **Post** /v2/labor/shifts/search | SearchShifts
[**UpdateBreakType**](LaborApi.md#UpdateBreakType) | **Put** /v2/labor/break-types/{id} | UpdateBreakType
[**UpdateShift**](LaborApi.md#UpdateShift) | **Put** /v2/labor/shifts/{id} | UpdateShift
[**UpdateWorkweekConfig**](LaborApi.md#UpdateWorkweekConfig) | **Put** /v2/labor/workweek-configs/{id} | UpdateWorkweekConfig

# **CreateBreakType**
> CreateBreakTypeResponse CreateBreakType(ctx, body)
CreateBreakType

Creates a new `BreakType`.  A `BreakType` is a template for creating `Break` objects. You must provide the following values in your request to this endpoint:  - `location_id` - `break_name` - `expected_duration` - `is_paid`  You can only have 3 `BreakType` instances per location. If you attempt to add a 4th `BreakType` for a location, an `INVALID_REQUEST_ERROR` \"Exceeded limit of 3 breaks per location.\" is returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBreakTypeRequest**](CreateBreakTypeRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateBreakTypeResponse**](CreateBreakTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateShift**
> CreateShiftResponse CreateShift(ctx, body)
CreateShift

Creates a new `Shift`.  A `Shift` represents a complete work day for a single employee. You must provide the following values in your request to this endpoint:  - `location_id` - `employee_id` - `start_at`  An attempt to create a new `Shift` can result in a `BAD_REQUEST` error when: - The `status` of the new `Shift` is `OPEN` and the employee has another shift with an `OPEN` status. - The `start_at` date is in the future - the `start_at` or `end_at` overlaps another shift for the same employee - If `Break`s are set in the request, a break `start_at` must not be before the `Shift.start_at`. A break `end_at` must not be after the `Shift.end_at`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateShiftRequest**](CreateShiftRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateShiftResponse**](CreateShiftResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBreakType**
> DeleteBreakTypeResponse DeleteBreakType(ctx, id)
DeleteBreakType

Deletes an existing `BreakType`.  A `BreakType` can be deleted even if it is referenced from a `Shift`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| UUID for the &#x60;BreakType&#x60; being deleted. | 

### Return type

[**DeleteBreakTypeResponse**](DeleteBreakTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteShift**
> DeleteShiftResponse DeleteShift(ctx, id)
DeleteShift

Deletes a `Shift`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| UUID for the &#x60;Shift&#x60; being deleted. | 

### Return type

[**DeleteShiftResponse**](DeleteShiftResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBreakType**
> GetBreakTypeResponse GetBreakType(ctx, id)
GetBreakType

Returns a single `BreakType` specified by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| UUID for the &#x60;BreakType&#x60; being retrieved. | 

### Return type

[**GetBreakTypeResponse**](GetBreakTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEmployeeWage**
> GetEmployeeWageResponse GetEmployeeWage(ctx, id)
GetEmployeeWage

Returns a single `EmployeeWage` specified by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| UUID for the &#x60;EmployeeWage&#x60; being retrieved. | 

### Return type

[**GetEmployeeWageResponse**](GetEmployeeWageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetShift**
> GetShiftResponse GetShift(ctx, id)
GetShift

Returns a single `Shift` specified by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| UUID for the &#x60;Shift&#x60; being retrieved. | 

### Return type

[**GetShiftResponse**](GetShiftResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTeamMemberWage**
> GetTeamMemberWageResponse GetTeamMemberWage(ctx, id)
GetTeamMemberWage

Returns a single `TeamMemberWage` specified by id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| UUID for the &#x60;TeamMemberWage&#x60; being retrieved. | 

### Return type

[**GetTeamMemberWageResponse**](GetTeamMemberWageResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBreakTypes**
> ListBreakTypesResponse ListBreakTypes(ctx, optional)
ListBreakTypes

Returns a paginated list of `BreakType` instances for a business.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaborApiListBreakTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaborApiListBreakTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationId** | **optional.String**| Filter Break Types returned to only those that are associated with the specified location. | 
 **limit** | **optional.Int32**| Maximum number of Break Types to return per page. Can range between 1 and 200. The default is the maximum at 200. | 
 **cursor** | **optional.String**| Pointer to the next page of Break Type results to fetch. | 

### Return type

[**ListBreakTypesResponse**](ListBreakTypesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEmployeeWages**
> ListEmployeeWagesResponse ListEmployeeWages(ctx, optional)
ListEmployeeWages

Returns a paginated list of `EmployeeWage` instances for a business.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaborApiListEmployeeWagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaborApiListEmployeeWagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employeeId** | **optional.String**| Filter wages returned to only those that are associated with the specified employee. | 
 **limit** | **optional.Int32**| Maximum number of Employee Wages to return per page. Can range between 1 and 200. The default is the maximum at 200. | 
 **cursor** | **optional.String**| Pointer to the next page of Employee Wage results to fetch. | 

### Return type

[**ListEmployeeWagesResponse**](ListEmployeeWagesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTeamMemberWages**
> ListTeamMemberWagesResponse ListTeamMemberWages(ctx, optional)
ListTeamMemberWages

Returns a paginated list of `TeamMemberWage` instances for a business.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaborApiListTeamMemberWagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaborApiListTeamMemberWagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **teamMemberId** | **optional.String**| Filter wages returned to only those that are associated with the specified team member. | 
 **limit** | **optional.Int32**| Maximum number of Team Member Wages to return per page. Can range between 1 and 200. The default is the maximum at 200. | 
 **cursor** | **optional.String**| Pointer to the next page of Employee Wage results to fetch. | 

### Return type

[**ListTeamMemberWagesResponse**](ListTeamMemberWagesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListWorkweekConfigs**
> ListWorkweekConfigsResponse ListWorkweekConfigs(ctx, optional)
ListWorkweekConfigs

Returns a list of `WorkweekConfig` instances for a business.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***LaborApiListWorkweekConfigsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LaborApiListWorkweekConfigsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| Maximum number of Workweek Configs to return per page. | 
 **cursor** | **optional.String**| Pointer to the next page of Workweek Config results to fetch. | 

### Return type

[**ListWorkweekConfigsResponse**](ListWorkweekConfigsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchShifts**
> SearchShiftsResponse SearchShifts(ctx, body)
SearchShifts

Returns a paginated list of `Shift` records for a business. The list to be returned can be filtered by: - Location IDs **and** - employee IDs **and** - shift status (`OPEN`, `CLOSED`) **and** - shift start **and** - shift end **and** - work day details  The list can be sorted by: - `start_at` - `end_at` - `created_at` - `updated_at`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchShiftsRequest**](SearchShiftsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchShiftsResponse**](SearchShiftsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBreakType**
> UpdateBreakTypeResponse UpdateBreakType(ctx, body, id)
UpdateBreakType

Updates an existing `BreakType`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBreakTypeRequest**](UpdateBreakTypeRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **id** | **string**| UUID for the &#x60;BreakType&#x60; being updated. | 

### Return type

[**UpdateBreakTypeResponse**](UpdateBreakTypeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateShift**
> UpdateShiftResponse UpdateShift(ctx, body, id)
UpdateShift

Updates an existing `Shift`.  When adding a `Break` to a `Shift`, any earlier `Breaks` in the `Shift` have the `end_at` property set to a valid RFC-3339 datetime string.  When closing a `Shift`, all `Break` instances in the shift must be complete with `end_at` set on each `Break`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateShiftRequest**](UpdateShiftRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **id** | **string**| ID of the object being updated. | 

### Return type

[**UpdateShiftResponse**](UpdateShiftResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWorkweekConfig**
> UpdateWorkweekConfigResponse UpdateWorkweekConfig(ctx, body, id)
UpdateWorkweekConfig

Updates a `WorkweekConfig`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateWorkweekConfigRequest**](UpdateWorkweekConfigRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **id** | **string**| UUID for the &#x60;WorkweekConfig&#x60; object being updated. | 

### Return type

[**UpdateWorkweekConfigResponse**](UpdateWorkweekConfigResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

