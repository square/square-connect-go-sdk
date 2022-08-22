# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkUpsertCustomerCustomAttributes**](CustomerCustomAttributesApi.md#BulkUpsertCustomerCustomAttributes) | **Post** /v2/customers/custom-attributes/bulk-upsert | BulkUpsertCustomerCustomAttributes
[**CreateCustomerCustomAttributeDefinition**](CustomerCustomAttributesApi.md#CreateCustomerCustomAttributeDefinition) | **Post** /v2/customers/custom-attribute-definitions | CreateCustomerCustomAttributeDefinition
[**DeleteCustomerCustomAttribute**](CustomerCustomAttributesApi.md#DeleteCustomerCustomAttribute) | **Delete** /v2/customers/{customer_id}/custom-attributes/{key} | DeleteCustomerCustomAttribute
[**DeleteCustomerCustomAttributeDefinition**](CustomerCustomAttributesApi.md#DeleteCustomerCustomAttributeDefinition) | **Delete** /v2/customers/custom-attribute-definitions/{key} | DeleteCustomerCustomAttributeDefinition
[**ListCustomerCustomAttributeDefinitions**](CustomerCustomAttributesApi.md#ListCustomerCustomAttributeDefinitions) | **Get** /v2/customers/custom-attribute-definitions | ListCustomerCustomAttributeDefinitions
[**ListCustomerCustomAttributes**](CustomerCustomAttributesApi.md#ListCustomerCustomAttributes) | **Get** /v2/customers/{customer_id}/custom-attributes | ListCustomerCustomAttributes
[**RetrieveCustomerCustomAttribute**](CustomerCustomAttributesApi.md#RetrieveCustomerCustomAttribute) | **Get** /v2/customers/{customer_id}/custom-attributes/{key} | RetrieveCustomerCustomAttribute
[**RetrieveCustomerCustomAttributeDefinition**](CustomerCustomAttributesApi.md#RetrieveCustomerCustomAttributeDefinition) | **Get** /v2/customers/custom-attribute-definitions/{key} | RetrieveCustomerCustomAttributeDefinition
[**UpdateCustomerCustomAttributeDefinition**](CustomerCustomAttributesApi.md#UpdateCustomerCustomAttributeDefinition) | **Put** /v2/customers/custom-attribute-definitions/{key} | UpdateCustomerCustomAttributeDefinition
[**UpsertCustomerCustomAttribute**](CustomerCustomAttributesApi.md#UpsertCustomerCustomAttribute) | **Post** /v2/customers/{customer_id}/custom-attributes/{key} | UpsertCustomerCustomAttribute

# **BulkUpsertCustomerCustomAttributes**
> BulkUpsertCustomerCustomAttributesResponse BulkUpsertCustomerCustomAttributes(ctx, body)
BulkUpsertCustomerCustomAttributes

Creates or updates [custom attributes](entity:CustomAttribute) for customer profiles as a bulk operation.  Use this endpoint to set the value of one or more custom attributes for one or more customer profiles. A custom attribute is based on a custom attribute definition in a Square seller account, which is created using the [CreateCustomerCustomAttributeDefinition](api-endpoint:CustomerCustomAttributes-CreateCustomerCustomAttributeDefinition) endpoint.  This `BulkUpsertCustomerCustomAttributes` endpoint accepts a map of 1 to 25 individual upsert  requests and returns a map of individual upsert responses. Each upsert request has a unique ID  and provides a customer ID and custom attribute. Each upsert response is returned with the ID  of the corresponding request.  To create or update a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkUpsertCustomerCustomAttributesRequest**](BulkUpsertCustomerCustomAttributesRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkUpsertCustomerCustomAttributesResponse**](BulkUpsertCustomerCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomerCustomAttributeDefinition**
> CreateCustomerCustomAttributeDefinitionResponse CreateCustomerCustomAttributeDefinition(ctx, body)
CreateCustomerCustomAttributeDefinition

Creates a customer-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account. Use this endpoint to define a custom attribute that can be associated with customer profiles.  A custom attribute definition specifies the `key`, `visibility`, `schema`, and other properties for a custom attribute. After the definition is created, you can call [UpsertCustomerCustomAttribute](api-endpoint:CustomerCustomAttributes-UpsertCustomerCustomAttribute) or [BulkUpsertCustomerCustomAttributes](api-endpoint:CustomerCustomAttributes-BulkUpsertCustomerCustomAttributes) to set the custom attribute for customer profiles in the seller's Customer Directory.  Sellers can view all custom attributes in exported customer data, including those set to `VISIBILITY_HIDDEN`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCustomerCustomAttributeDefinitionRequest**](CreateCustomerCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateCustomerCustomAttributeDefinitionResponse**](CreateCustomerCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomerCustomAttribute**
> DeleteCustomerCustomAttributeResponse DeleteCustomerCustomAttribute(ctx, customerId, key)
DeleteCustomerCustomAttribute

Deletes a [custom attribute](entity:CustomAttribute) associated with a customer profile.  To delete a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The ID of the target [customer profile](entity:Customer). | 
  **key** | **string**| The key of the custom attribute to delete. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 

### Return type

[**DeleteCustomerCustomAttributeResponse**](DeleteCustomerCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomerCustomAttributeDefinition**
> DeleteCustomerCustomAttributeDefinitionResponse DeleteCustomerCustomAttributeDefinition(ctx, key)
DeleteCustomerCustomAttributeDefinition

Deletes a customer-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.  Deleting a custom attribute definition also deletes the corresponding custom attribute from all customer profiles in the seller's Customer Directory.  Only the definition owner can delete a custom attribute definition.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to delete. | 

### Return type

[**DeleteCustomerCustomAttributeDefinitionResponse**](DeleteCustomerCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerCustomAttributeDefinitions**
> ListCustomerCustomAttributeDefinitionsResponse ListCustomerCustomAttributeDefinitions(ctx, optional)
ListCustomerCustomAttributeDefinitions

Lists the customer-related [custom attribute definitions](entity:CustomAttributeDefinition) that belong to a Square seller account.  When all response pages are retrieved, the results include all custom attribute definitions that are visible to the requesting application, including those that are created by other applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomerCustomAttributesApiListCustomerCustomAttributeDefinitionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerCustomAttributesApiListCustomerCustomAttributeDefinitionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory. The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint. Provide this cursor to retrieve the next page of results for your original request. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 

### Return type

[**ListCustomerCustomAttributeDefinitionsResponse**](ListCustomerCustomAttributeDefinitionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomerCustomAttributes**
> ListCustomerCustomAttributesResponse ListCustomerCustomAttributes(ctx, customerId, optional)
ListCustomerCustomAttributes

Lists the [custom attributes](entity:CustomAttribute) associated with a customer profile.  You can use the `with_definitions` query parameter to also retrieve custom attribute definitions in the same call.  When all response pages are retrieved, the results include all custom attributes that are visible to the requesting application, including those that are owned by other applications and set to `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The ID of the target [customer profile](entity:Customer). | 
 **optional** | ***CustomerCustomAttributesApiListCustomerCustomAttributesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerCustomAttributesApiListCustomerCustomAttributesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. This limit is advisory. The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100. The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint. Provide this cursor to retrieve the next page of results for your original request. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **withDefinitions** | **optional.Bool**| Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the &#x60;definition&#x60; field of each custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of each custom attribute, information about the data type, or other definition details. The default value is &#x60;false&#x60;. | [default to false]

### Return type

[**ListCustomerCustomAttributesResponse**](ListCustomerCustomAttributesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCustomerCustomAttribute**
> RetrieveCustomerCustomAttributeResponse RetrieveCustomerCustomAttribute(ctx, customerId, key, optional)
RetrieveCustomerCustomAttribute

Retrieves a [custom attribute](entity:CustomAttribute) associated with a customer profile.  You can use the `with_definition` query parameter to also retrieve the custom attribute definition in the same call.  To retrieve a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The ID of the target [customer profile](entity:Customer). | 
  **key** | **string**| The key of the custom attribute to retrieve. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 
 **optional** | ***CustomerCustomAttributesApiRetrieveCustomerCustomAttributeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerCustomAttributesApiRetrieveCustomerCustomAttributeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **withDefinition** | **optional.Bool**| Indicates whether to return the [custom attribute definition](entity:CustomAttributeDefinition) in the &#x60;definition&#x60; field of the custom attribute. Set this parameter to &#x60;true&#x60; to get the name and description of the custom attribute, information about the data type, or other definition details. The default value is &#x60;false&#x60;. | [default to false]
 **version** | **optional.Int32**| The current version of the custom attribute, which is used for strongly consistent reads to guarantee that you receive the most up-to-date data. When included in the request, Square returns the specified version or a higher version if one exists. If the specified version is higher than the current version, Square returns a &#x60;BAD_REQUEST&#x60; error. | 

### Return type

[**RetrieveCustomerCustomAttributeResponse**](RetrieveCustomerCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCustomerCustomAttributeDefinition**
> RetrieveCustomerCustomAttributeDefinitionResponse RetrieveCustomerCustomAttributeDefinition(ctx, key, optional)
RetrieveCustomerCustomAttributeDefinition

Retrieves a customer-related [custom attribute definition](entity:CustomAttributeDefinition) from a Square seller account.  To retrieve a custom attribute definition created by another application, the `visibility` setting must be `VISIBILITY_READ_ONLY` or `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| The key of the custom attribute definition to retrieve. If the requesting application is not the definition owner, you must use the qualified key. | 
 **optional** | ***CustomerCustomAttributesApiRetrieveCustomerCustomAttributeDefinitionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomerCustomAttributesApiRetrieveCustomerCustomAttributeDefinitionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int32**| The current version of the custom attribute definition, which is used for strongly consistent reads to guarantee that you receive the most up-to-date data. When included in the request, Square returns the specified version or a higher version if one exists. If the specified version is higher than the current version, Square returns a &#x60;BAD_REQUEST&#x60; error. | 

### Return type

[**RetrieveCustomerCustomAttributeDefinitionResponse**](RetrieveCustomerCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomerCustomAttributeDefinition**
> UpdateCustomerCustomAttributeDefinitionResponse UpdateCustomerCustomAttributeDefinition(ctx, body, key)
UpdateCustomerCustomAttributeDefinition

Updates a customer-related [custom attribute definition](entity:CustomAttributeDefinition) for a Square seller account.  Use this endpoint to update the following fields: `name`, `description`, `visibility`, or the `schema` for a `Selection` data type.  Only the definition owner can update a custom attribute definition. Note that sellers can view all custom attributes in exported customer data, including those set to `VISIBILITY_HIDDEN`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCustomerCustomAttributeDefinitionRequest**](UpdateCustomerCustomAttributeDefinitionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **key** | **string**| The key of the custom attribute definition to update. | 

### Return type

[**UpdateCustomerCustomAttributeDefinitionResponse**](UpdateCustomerCustomAttributeDefinitionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpsertCustomerCustomAttribute**
> UpsertCustomerCustomAttributeResponse UpsertCustomerCustomAttribute(ctx, body, customerId, key)
UpsertCustomerCustomAttribute

Creates or updates a [custom attribute](entity:CustomAttribute) for a customer profile.  Use this endpoint to set the value of a custom attribute for a specified customer profile. A custom attribute is based on a custom attribute definition in a Square seller account, which is created using the [CreateCustomerCustomAttributeDefinition](api-endpoint:CustomerCustomAttributes-CreateCustomerCustomAttributeDefinition) endpoint.  To create or update a custom attribute owned by another application, the `visibility` setting must be `VISIBILITY_READ_WRITE_VALUES`. Note that seller-defined custom attributes (also known as custom fields) are always set to `VISIBILITY_READ_WRITE_VALUES`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpsertCustomerCustomAttributeRequest**](UpsertCustomerCustomAttributeRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **customerId** | **string**| The ID of the target [customer profile](entity:Customer). | 
  **key** | **string**| The key of the custom attribute to create or update. This key must match the &#x60;key&#x60; of a custom attribute definition in the Square seller account. If the requesting application is not the definition owner, you must use the qualified key. | 

### Return type

[**UpsertCustomerCustomAttributeResponse**](UpsertCustomerCustomAttributeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

