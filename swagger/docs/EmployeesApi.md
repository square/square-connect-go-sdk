# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEmployees**](EmployeesApi.md#ListEmployees) | **Get** /v2/employees | ListEmployees
[**RetrieveEmployee**](EmployeesApi.md#RetrieveEmployee) | **Get** /v2/employees/{id} | RetrieveEmployee

# **ListEmployees**
> ListEmployeesResponse ListEmployees(ctx, optional)
ListEmployees

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EmployeesApiListEmployeesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EmployeesApiListEmployeesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **locationId** | **optional.String**|  | 
 **status** | [**optional.Interface of EmployeeStatus**](.md)| Specifies the EmployeeStatus to filter the employee by. | 
 **limit** | **optional.Int32**| The number of employees to be returned on each page. | 
 **cursor** | **optional.String**| The token required to retrieve the specified page of results. | 

### Return type

[**ListEmployeesResponse**](ListEmployeesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveEmployee**
> RetrieveEmployeeResponse RetrieveEmployee(ctx, id)
RetrieveEmployee

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| UUID for the employee that was requested. | 

### Return type

[**RetrieveEmployeeResponse**](RetrieveEmployeeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

