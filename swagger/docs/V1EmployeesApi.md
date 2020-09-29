# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmployee**](V1EmployeesApi.md#CreateEmployee) | **Post** /v1/me/employees | CreateEmployee
[**CreateEmployeeRole**](V1EmployeesApi.md#CreateEmployeeRole) | **Post** /v1/me/roles | CreateEmployeeRole
[**CreateTimecard**](V1EmployeesApi.md#CreateTimecard) | **Post** /v1/me/timecards | CreateTimecard
[**DeleteTimecard**](V1EmployeesApi.md#DeleteTimecard) | **Delete** /v1/me/timecards/{timecard_id} | DeleteTimecard
[**ListCashDrawerShifts**](V1EmployeesApi.md#ListCashDrawerShifts) | **Get** /v1/{location_id}/cash-drawer-shifts | ListCashDrawerShifts
[**ListEmployeeRoles**](V1EmployeesApi.md#ListEmployeeRoles) | **Get** /v1/me/roles | ListEmployeeRoles
[**ListEmployees**](V1EmployeesApi.md#ListEmployees) | **Get** /v1/me/employees | ListEmployees
[**ListTimecardEvents**](V1EmployeesApi.md#ListTimecardEvents) | **Get** /v1/me/timecards/{timecard_id}/events | ListTimecardEvents
[**ListTimecards**](V1EmployeesApi.md#ListTimecards) | **Get** /v1/me/timecards | ListTimecards
[**RetrieveCashDrawerShift**](V1EmployeesApi.md#RetrieveCashDrawerShift) | **Get** /v1/{location_id}/cash-drawer-shifts/{shift_id} | RetrieveCashDrawerShift
[**RetrieveEmployee**](V1EmployeesApi.md#RetrieveEmployee) | **Get** /v1/me/employees/{employee_id} | RetrieveEmployee
[**RetrieveEmployeeRole**](V1EmployeesApi.md#RetrieveEmployeeRole) | **Get** /v1/me/roles/{role_id} | RetrieveEmployeeRole
[**RetrieveTimecard**](V1EmployeesApi.md#RetrieveTimecard) | **Get** /v1/me/timecards/{timecard_id} | RetrieveTimecard
[**UpdateEmployee**](V1EmployeesApi.md#UpdateEmployee) | **Put** /v1/me/employees/{employee_id} | UpdateEmployee
[**UpdateEmployeeRole**](V1EmployeesApi.md#UpdateEmployeeRole) | **Put** /v1/me/roles/{role_id} | UpdateEmployeeRole
[**UpdateTimecard**](V1EmployeesApi.md#UpdateTimecard) | **Put** /v1/me/timecards/{timecard_id} | UpdateTimecard

# **CreateEmployee**
> V1Employee CreateEmployee(ctx, body)
CreateEmployee

 Use the CreateEmployee endpoint to add an employee to a Square account. Employees created with the Connect API have an initial status of `INACTIVE`. Inactive employees cannot sign in to Square Point of Sale until they are activated from the Square Dashboard. Employee status cannot be changed with the Connect API.  <aside class=\"important\"> Employee entities cannot be deleted. To disable employee profiles, set the employee's status to <code>INACTIVE</code> </aside>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Employee**](V1Employee.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**V1Employee**](V1Employee.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEmployeeRole**
> V1EmployeeRole CreateEmployeeRole(ctx, body)
CreateEmployeeRole

Creates an employee role you can then assign to employees.  Square accounts can include any number of roles that can be assigned to employees. These roles define the actions and permissions granted to an employee with that role. For example, an employee with a \"Shift Manager\" role might be able to issue refunds in Square Point of Sale, whereas an employee with a \"Clerk\" role might not.  Roles are assigned with the [V1UpdateEmployee](#endpoint-v1updateemployee) endpoint. An employee can have only one role at a time.  If an employee has no role, they have none of the permissions associated with roles. All employees can accept payments with Square Point of Sale.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1EmployeeRole**](V1EmployeeRole.md)| An EmployeeRole object with a name and permissions, and an optional owner flag. | 

### Return type

[**V1EmployeeRole**](V1EmployeeRole.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTimecard**
> V1Timecard CreateTimecard(ctx, body)
CreateTimecard

Creates a timecard for an employee and clocks them in with an `API_CREATE` event and a `clockin_time` set to the current time unless the request provides a different value.  To import timecards from another system (rather than clocking someone in). Specify the `clockin_time` and* `clockout_time` in the request.  Timecards correspond to exactly one shift for a given employee, bounded by the `clockin_time` and `clockout_time` fields. An employee is considered clocked in if they have a timecard that doesn't have a `clockout_time` set. An employee that is currently clocked in cannot be clocked in a second time.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Timecard**](V1Timecard.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**V1Timecard**](V1Timecard.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTimecard**
> interface{} DeleteTimecard(ctx, timecardId)
DeleteTimecard

Deletes a timecard. Timecards can also be deleted through the Square Dashboard. Deleted timecards are still accessible through Connect API endpoints, but cannot be modified. The `deleted` field of the `Timecard` object indicates whether the timecard has been deleted.   __Note__: By default, deleted timecards appear alongside valid timecards in results returned by the [ListTimecards](#endpoint-v1employees-listtimecards) endpoint. To filter deleted timecards, include the `deleted` query parameter in the list request.  Only approved accounts can manage their employees with Square. Unapproved accounts cannot use employee management features with the API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timecardId** | **string**| The ID of the timecard to delete. | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCashDrawerShifts**
> []V1CashDrawerShift ListCashDrawerShifts(ctx, locationId, optional)
ListCashDrawerShifts

Provides the details for all of a location's cash drawer shifts during a date range. The date range you specify cannot exceed 90 days.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list cash drawer shifts for. | 
 **optional** | ***V1EmployeesApiListCashDrawerShiftsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1EmployeesApiListCashDrawerShiftsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | [**optional.Interface of SortOrder**](.md)| The order in which cash drawer shifts are listed in the response, based on their created_at field. Default value: ASC | 
 **beginTime** | **optional.String**| The beginning of the requested reporting period, in ISO 8601 format. Default value: The current time minus 90 days. | 
 **endTime** | **optional.String**| The beginning of the requested reporting period, in ISO 8601 format. Default value: The current time. | 

### Return type

[**[]V1CashDrawerShift**](V1CashDrawerShift.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEmployeeRoles**
> []V1EmployeeRole ListEmployeeRoles(ctx, optional)
ListEmployeeRoles

Provides summary information for all of a business's employee roles.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1EmployeesApiListEmployeeRolesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1EmployeesApiListEmployeeRolesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | [**optional.Interface of SortOrder**](.md)| The order in which employees are listed in the response, based on their created_at field.Default value: ASC | 
 **limit** | **optional.Int32**| The maximum integer number of employee entities to return in a single response. Default 100, maximum 200. | 
 **batchToken** | **optional.String**| A pagination cursor to retrieve the next set of results for your original query to the endpoint. | 

### Return type

[**[]V1EmployeeRole**](V1EmployeeRole.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEmployees**
> []V1Employee ListEmployees(ctx, optional)
ListEmployees

Provides summary information for all of a business's employees.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1EmployeesApiListEmployeesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1EmployeesApiListEmployeesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | [**optional.Interface of SortOrder**](.md)| The order in which employees are listed in the response, based on their created_at field.      Default value: ASC | 
 **beginUpdatedAt** | **optional.String**| If filtering results by their updated_at field, the beginning of the requested reporting period, in ISO 8601 format | 
 **endUpdatedAt** | **optional.String**| If filtering results by there updated_at field, the end of the requested reporting period, in ISO 8601 format. | 
 **beginCreatedAt** | **optional.String**| If filtering results by their created_at field, the beginning of the requested reporting period, in ISO 8601 format. | 
 **endCreatedAt** | **optional.String**| If filtering results by their created_at field, the end of the requested reporting period, in ISO 8601 format. | 
 **status** | [**optional.Interface of V1ListEmployeesRequestStatus**](.md)| If provided, the endpoint returns only employee entities with the specified status (ACTIVE or INACTIVE). | 
 **externalId** | **optional.String**| If provided, the endpoint returns only employee entities with the specified external_id. | 
 **limit** | **optional.Int32**| The maximum integer number of employee entities to return in a single response. Default 100, maximum 200. | 
 **batchToken** | **optional.String**| A pagination cursor to retrieve the next set of results for your original query to the endpoint. | 

### Return type

[**[]V1Employee**](V1Employee.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTimecardEvents**
> []V1TimecardEvent ListTimecardEvents(ctx, timecardId)
ListTimecardEvents

Provides summary information for all events associated with a particular timecard.   <aside> Only approved accounts can manage their employees with Square. Unapproved accounts cannot use employee management features with the API. </aside>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timecardId** | **string**| The ID of the timecard to list events for. | 

### Return type

[**[]V1TimecardEvent**](V1TimecardEvent.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTimecards**
> []V1Timecard ListTimecards(ctx, optional)
ListTimecards

Provides summary information for all of a business's employee timecards.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***V1EmployeesApiListTimecardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1EmployeesApiListTimecardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | [**optional.Interface of SortOrder**](.md)| The order in which timecards are listed in the response, based on their created_at field. | 
 **employeeId** | **optional.String**| If provided, the endpoint returns only timecards for the employee with the specified ID. | 
 **beginClockinTime** | **optional.String**| If filtering results by their clockin_time field, the beginning of the requested reporting period, in ISO 8601 format. | 
 **endClockinTime** | **optional.String**| If filtering results by their clockin_time field, the end of the requested reporting period, in ISO 8601 format. | 
 **beginClockoutTime** | **optional.String**| If filtering results by their clockout_time field, the beginning of the requested reporting period, in ISO 8601 format. | 
 **endClockoutTime** | **optional.String**| If filtering results by their clockout_time field, the end of the requested reporting period, in ISO 8601 format. | 
 **beginUpdatedAt** | **optional.String**| If filtering results by their updated_at field, the beginning of the requested reporting period, in ISO 8601 format. | 
 **endUpdatedAt** | **optional.String**| If filtering results by their updated_at field, the end of the requested reporting period, in ISO 8601 format. | 
 **deleted** | **optional.Bool**| If true, only deleted timecards are returned. If false, only valid timecards are returned.If you don&#x27;t provide this parameter, both valid and deleted timecards are returned. | [default to false]
 **limit** | **optional.Int32**| The maximum integer number of employee entities to return in a single response. Default 100, maximum 200. | 
 **batchToken** | **optional.String**| A pagination cursor to retrieve the next set of results for your original query to the endpoint. | 

### Return type

[**[]V1Timecard**](V1Timecard.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCashDrawerShift**
> V1CashDrawerShift RetrieveCashDrawerShift(ctx, locationId, shiftId)
RetrieveCashDrawerShift

Provides the details for a single cash drawer shift, including all events that occurred during the shift.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list cash drawer shifts for. | 
  **shiftId** | **string**| The shift&#x27;s ID. | 

### Return type

[**V1CashDrawerShift**](V1CashDrawerShift.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveEmployee**
> V1Employee RetrieveEmployee(ctx, employeeId)
RetrieveEmployee

Provides the details for a single employee.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **employeeId** | **string**| The employee&#x27;s ID. | 

### Return type

[**V1Employee**](V1Employee.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveEmployeeRole**
> V1EmployeeRole RetrieveEmployeeRole(ctx, roleId)
RetrieveEmployeeRole

Provides the details for a single employee role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **roleId** | **string**| The role&#x27;s ID. | 

### Return type

[**V1EmployeeRole**](V1EmployeeRole.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveTimecard**
> V1Timecard RetrieveTimecard(ctx, timecardId)
RetrieveTimecard

Provides the details for a single timecard.   <aside> Only approved accounts can manage their employees with Square. Unapproved accounts cannot use employee management features with the API. </aside>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **timecardId** | **string**| The timecard&#x27;s ID. | 

### Return type

[**V1Timecard**](V1Timecard.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmployee**
> V1Employee UpdateEmployee(ctx, body, employeeId)
UpdateEmployee

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Employee**](V1Employee.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **employeeId** | **string**| The ID of the role to modify. | 

### Return type

[**V1Employee**](V1Employee.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEmployeeRole**
> V1EmployeeRole UpdateEmployeeRole(ctx, body, roleId)
UpdateEmployeeRole

Modifies the details of an employee role.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1EmployeeRole**](V1EmployeeRole.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **roleId** | **string**| The ID of the role to modify. | 

### Return type

[**V1EmployeeRole**](V1EmployeeRole.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTimecard**
> V1Timecard UpdateTimecard(ctx, body, timecardId)
UpdateTimecard

Modifies the details of a timecard with an `API_EDIT` event for the timecard. Updating an active timecard with a `clockout_time` clocks the employee out.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Timecard**](V1Timecard.md)| An object containing the fields to POST for the request.
See the corresponding object definition for field details. | 
  **timecardId** | **string**| TThe ID of the timecard to modify. | 

### Return type

[**V1Timecard**](V1Timecard.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

