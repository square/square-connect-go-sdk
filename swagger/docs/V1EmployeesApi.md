# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmployee**](V1EmployeesApi.md#CreateEmployee) | **Post** /v1/me/employees | CreateEmployee
[**CreateEmployeeRole**](V1EmployeesApi.md#CreateEmployeeRole) | **Post** /v1/me/roles | CreateEmployeeRole
[**ListEmployeeRoles**](V1EmployeesApi.md#ListEmployeeRoles) | **Get** /v1/me/roles | ListEmployeeRoles
[**ListEmployees**](V1EmployeesApi.md#ListEmployees) | **Get** /v1/me/employees | ListEmployees
[**RetrieveEmployee**](V1EmployeesApi.md#RetrieveEmployee) | **Get** /v1/me/employees/{employee_id} | RetrieveEmployee
[**RetrieveEmployeeRole**](V1EmployeesApi.md#RetrieveEmployeeRole) | **Get** /v1/me/roles/{role_id} | RetrieveEmployeeRole
[**UpdateEmployee**](V1EmployeesApi.md#UpdateEmployee) | **Put** /v1/me/employees/{employee_id} | UpdateEmployee
[**UpdateEmployeeRole**](V1EmployeesApi.md#UpdateEmployeeRole) | **Put** /v1/me/roles/{role_id} | UpdateEmployeeRole

# **CreateEmployee**
> V1Employee CreateEmployee(ctx, body)
CreateEmployee

 Use the CreateEmployee endpoint to add an employee to a Square account. Employees created with the Connect API have an initial status of `INACTIVE`. Inactive employees cannot sign in to Square Point of Sale until they are activated from the Square Dashboard. Employee status cannot be changed with the Connect API.  Employee entities cannot be deleted. To disable employee profiles, set the employee's status to <code>INACTIVE</code>

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

Creates an employee role you can then assign to employees.  Square accounts can include any number of roles that can be assigned to employees. These roles define the actions and permissions granted to an employee with that role. For example, an employee with a \"Shift Manager\" role might be able to issue refunds in Square Point of Sale, whereas an employee with a \"Clerk\" role might not.  Roles are assigned with the [V1UpdateEmployee](api-endpoint:V1Employees-UpdateEmployeeRole) endpoint. An employee can have only one role at a time.  If an employee has no role, they have none of the permissions associated with roles. All employees can accept payments with Square Point of Sale.

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

