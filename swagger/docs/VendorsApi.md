# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateVendors**](VendorsApi.md#BulkCreateVendors) | **Post** /v2/vendors/bulk-create | BulkCreateVendors
[**BulkRetrieveVendors**](VendorsApi.md#BulkRetrieveVendors) | **Post** /v2/vendors/bulk-retrieve | BulkRetrieveVendors
[**BulkUpdateVendors**](VendorsApi.md#BulkUpdateVendors) | **Put** /v2/vendors/bulk-update | BulkUpdateVendors
[**CreateVendor**](VendorsApi.md#CreateVendor) | **Post** /v2/vendors/create | CreateVendor
[**RetrieveVendor**](VendorsApi.md#RetrieveVendor) | **Get** /v2/vendors/{vendor_id} | RetrieveVendor
[**SearchVendors**](VendorsApi.md#SearchVendors) | **Post** /v2/vendors/search | SearchVendors
[**UpdateVendor**](VendorsApi.md#UpdateVendor) | **Put** /v2/vendors/{vendor_id} | UpdateVendor

# **BulkCreateVendors**
> BulkCreateVendorsResponse BulkCreateVendors(ctx, body)
BulkCreateVendors

Creates one or more [Vendor](entity:Vendor) objects to represent suppliers to a seller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkCreateVendorsRequest**](BulkCreateVendorsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkCreateVendorsResponse**](BulkCreateVendorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkRetrieveVendors**
> BulkRetrieveVendorsResponse BulkRetrieveVendors(ctx, body)
BulkRetrieveVendors

Retrieves one or more vendors of specified [Vendor](entity:Vendor) IDs.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkRetrieveVendorsRequest**](BulkRetrieveVendorsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkRetrieveVendorsResponse**](BulkRetrieveVendorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUpdateVendors**
> BulkUpdateVendorsResponse BulkUpdateVendors(ctx, body)
BulkUpdateVendors

Updates one or more of existing [Vendor](entity:Vendor) objects as suppliers to a seller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkUpdateVendorsRequest**](BulkUpdateVendorsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkUpdateVendorsResponse**](BulkUpdateVendorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVendor**
> CreateVendorResponse CreateVendor(ctx, body)
CreateVendor

Creates a single [Vendor](entity:Vendor) object to represent a supplier to a seller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateVendorRequest**](CreateVendorRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateVendorResponse**](CreateVendorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveVendor**
> RetrieveVendorResponse RetrieveVendor(ctx, vendorId)
RetrieveVendor

Retrieves the vendor of a specified [Vendor](entity:Vendor) ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **vendorId** | **string**| ID of the [Vendor](entity:Vendor) to retrieve. | 

### Return type

[**RetrieveVendorResponse**](RetrieveVendorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchVendors**
> SearchVendorsResponse SearchVendors(ctx, body)
SearchVendors

Searches for vendors using a filter against supported [Vendor](entity:Vendor) properties and a supported sorter.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchVendorsRequest**](SearchVendorsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchVendorsResponse**](SearchVendorsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVendor**
> UpdateVendorResponse UpdateVendor(ctx, body)
UpdateVendor

Updates an existing [Vendor](entity:Vendor) object as a supplier to a seller.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateVendorRequest**](UpdateVendorRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**UpdateVendorResponse**](UpdateVendorResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

