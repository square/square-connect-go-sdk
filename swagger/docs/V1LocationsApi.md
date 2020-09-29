# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLocations**](V1LocationsApi.md#ListLocations) | **Get** /v1/me/locations | ListLocations
[**RetrieveBusiness**](V1LocationsApi.md#RetrieveBusiness) | **Get** /v1/me | RetrieveBusiness

# **ListLocations**
> []V1Merchant ListLocations(ctx, )
ListLocations

Provides details for all business locations associated with a Square account, including the Square-assigned object ID for the location.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]V1Merchant**](V1Merchant.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveBusiness**
> V1Merchant RetrieveBusiness(ctx, )
RetrieveBusiness

Get the general information for a business.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**V1Merchant**](V1Merchant.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

