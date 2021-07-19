# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLocation**](LocationsApi.md#CreateLocation) | **Post** /v2/locations | CreateLocation
[**ListLocations**](LocationsApi.md#ListLocations) | **Get** /v2/locations | ListLocations
[**RetrieveLocation**](LocationsApi.md#RetrieveLocation) | **Get** /v2/locations/{location_id} | RetrieveLocation
[**UpdateLocation**](LocationsApi.md#UpdateLocation) | **Put** /v2/locations/{location_id} | UpdateLocation

# **CreateLocation**
> CreateLocationResponse CreateLocation(ctx, body)
CreateLocation

Creates a location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateLocationRequest**](CreateLocationRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateLocationResponse**](CreateLocationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLocations**
> ListLocationsResponse ListLocations(ctx, )
ListLocations

Provides information of all locations of a business.  Many Square API endpoints require a `location_id` parameter. The `id` field of the [`Location`](entity:Location) objects returned by this endpoint correspond to that `location_id` parameter.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListLocationsResponse**](ListLocationsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLocation**
> RetrieveLocationResponse RetrieveLocation(ctx, locationId)
RetrieveLocation

Retrieves details of a location. You can specify \"main\"  as the location ID to retrieve details of the  main location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to retrieve. If you specify the string \&quot;main\&quot;, then the endpoint returns the main location. | 

### Return type

[**RetrieveLocationResponse**](RetrieveLocationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLocation**
> UpdateLocationResponse UpdateLocation(ctx, body, locationId)
UpdateLocation

Updates a location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateLocationRequest**](UpdateLocationRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the location to update. | 

### Return type

[**UpdateLocationResponse**](UpdateLocationResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

