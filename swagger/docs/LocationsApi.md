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

Creates a [location](https://developer.squareup.com/docs/locations-api). Creating new locations allows for separate configuration of receipt layouts, item prices,  and sales reports. Developers can use locations to separate sales activity via applications  that integrate with Square from sales activity elsewhere in a seller's account.  Locations created programmatically with the Locations API will last forever and  are visible to the seller for their own management, so ensure that  each location has a sensible and unique name.

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

Provides details about all of the seller's locations, including those with an inactive status.

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

Retrieves details of a single location. Specify \"main\" as the location ID to retrieve details of the [main location](https://developer.squareup.com/docs/locations-api#about-the-main-location).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to retrieve. Specify the string \&quot;main\&quot; to return the main location. | 

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

