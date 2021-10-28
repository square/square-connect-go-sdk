# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGroupToCustomer**](CustomersApi.md#AddGroupToCustomer) | **Put** /v2/customers/{customer_id}/groups/{group_id} | AddGroupToCustomer
[**CreateCustomer**](CustomersApi.md#CreateCustomer) | **Post** /v2/customers | CreateCustomer
[**CreateCustomerCard**](CustomersApi.md#CreateCustomerCard) | **Post** /v2/customers/{customer_id}/cards | CreateCustomerCard
[**DeleteCustomer**](CustomersApi.md#DeleteCustomer) | **Delete** /v2/customers/{customer_id} | DeleteCustomer
[**DeleteCustomerCard**](CustomersApi.md#DeleteCustomerCard) | **Delete** /v2/customers/{customer_id}/cards/{card_id} | DeleteCustomerCard
[**ListCustomers**](CustomersApi.md#ListCustomers) | **Get** /v2/customers | ListCustomers
[**RemoveGroupFromCustomer**](CustomersApi.md#RemoveGroupFromCustomer) | **Delete** /v2/customers/{customer_id}/groups/{group_id} | RemoveGroupFromCustomer
[**RetrieveCustomer**](CustomersApi.md#RetrieveCustomer) | **Get** /v2/customers/{customer_id} | RetrieveCustomer
[**SearchCustomers**](CustomersApi.md#SearchCustomers) | **Post** /v2/customers/search | SearchCustomers
[**UpdateCustomer**](CustomersApi.md#UpdateCustomer) | **Put** /v2/customers/{customer_id} | UpdateCustomer

# **AddGroupToCustomer**
> AddGroupToCustomerResponse AddGroupToCustomer(ctx, customerId, groupId)
AddGroupToCustomer

Adds a group membership to a customer.  The customer is identified by the `customer_id` value and the customer group is identified by the `group_id` value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The ID of the customer to add to a group. | 
  **groupId** | **string**| The ID of the customer group to add the customer to. | 

### Return type

[**AddGroupToCustomerResponse**](AddGroupToCustomerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomer**
> CreateCustomerResponse CreateCustomer(ctx, body)
CreateCustomer

Creates a new customer for a business.  You must provide at least one of the following values in your request to this endpoint:  - `given_name` - `family_name` - `company_name` - `email_address` - `phone_number`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCustomerRequest**](CreateCustomerRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateCustomerResponse**](CreateCustomerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCustomerCard**
> CreateCustomerCardResponse CreateCustomerCard(ctx, body, customerId)
CreateCustomerCard

Adds a card on file to an existing customer.  As with charges, calls to `CreateCustomerCard` are idempotent. Multiple calls with the same card nonce return the same card record that was created with the provided nonce during the _first_ call.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCustomerCardRequest**](CreateCustomerCardRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **customerId** | **string**| The Square ID of the customer profile the card is linked to. | 

### Return type

[**CreateCustomerCardResponse**](CreateCustomerCardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomer**
> DeleteCustomerResponse DeleteCustomer(ctx, customerId, optional)
DeleteCustomer

Deletes a customer profile from a business. This operation also unlinks any associated cards on file.   As a best practice, you should include the `version` field in the request to enable [optimistic concurrency](https://developer.squareup.com/docs/working-with-apis/optimistic-concurrency) control. The value must be set to the current version of the customer profile.   To delete a customer profile that was created by merging existing profiles, you must use the ID of the newly created profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The ID of the customer to delete. | 
 **optional** | ***CustomersApiDeleteCustomerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiDeleteCustomerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | **optional.Int64**| The current version of the customer profile.  As a best practice, you should include this parameter to enable [optimistic concurrency](https://developer.squareup.com/docs/working-with-apis/optimistic-concurrency) control.  For more information, see [Delete a customer profile](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#delete-customer-profile). | 

### Return type

[**DeleteCustomerResponse**](DeleteCustomerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCustomerCard**
> DeleteCustomerCardResponse DeleteCustomerCard(ctx, customerId, cardId)
DeleteCustomerCard

Removes a card on file from a customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The ID of the customer that the card on file belongs to. | 
  **cardId** | **string**| The ID of the card on file to delete. | 

### Return type

[**DeleteCustomerCardResponse**](DeleteCustomerCardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCustomers**
> ListCustomersResponse ListCustomers(ctx, optional)
ListCustomers

Lists customer profiles associated with a Square account.  Under normal operating conditions, newly created or updated customer profiles become available for the listing operation in well under 30 seconds. Occasionally, propagation of the new or updated profiles can take closer to one minute or longer, especially during network incidents and outages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CustomersApiListCustomersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CustomersApiListCustomersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **limit** | **optional.Int32**| The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results. The limit is ignored if it is less than 1 or greater than 100. The default value is 100.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **sortField** | [**optional.Interface of CustomerSortField**](.md)| Indicates how customers should be sorted.  The default value is &#x60;DEFAULT&#x60;. | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)| Indicates whether customers should be sorted in ascending (&#x60;ASC&#x60;) or descending (&#x60;DESC&#x60;) order.  The default value is &#x60;ASC&#x60;. | 

### Return type

[**ListCustomersResponse**](ListCustomersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveGroupFromCustomer**
> RemoveGroupFromCustomerResponse RemoveGroupFromCustomer(ctx, customerId, groupId)
RemoveGroupFromCustomer

Removes a group membership from a customer.  The customer is identified by the `customer_id` value and the customer group is identified by the `group_id` value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The ID of the customer to remove from the group. | 
  **groupId** | **string**| The ID of the customer group to remove the customer from. | 

### Return type

[**RemoveGroupFromCustomerResponse**](RemoveGroupFromCustomerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCustomer**
> RetrieveCustomerResponse RetrieveCustomer(ctx, customerId)
RetrieveCustomer

Returns details for a single customer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **customerId** | **string**| The ID of the customer to retrieve. | 

### Return type

[**RetrieveCustomerResponse**](RetrieveCustomerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchCustomers**
> SearchCustomersResponse SearchCustomers(ctx, body)
SearchCustomers

Searches the customer profiles associated with a Square account using a supported query filter.  Calling `SearchCustomers` without any explicit query filter returns all customer profiles ordered alphabetically based on `given_name` and `family_name`.  Under normal operating conditions, newly created or updated customer profiles become available for the search operation in well under 30 seconds. Occasionally, propagation of the new or updated profiles can take closer to one minute or longer, especially during network incidents and outages.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchCustomersRequest**](SearchCustomersRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchCustomersResponse**](SearchCustomersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomer**
> UpdateCustomerResponse UpdateCustomer(ctx, body, customerId)
UpdateCustomer

Updates a customer profile. To change an attribute, specify the new value. To remove an attribute, specify the value as an empty string or empty object.  As a best practice, you should include the `version` field in the request to enable [optimistic concurrency](https://developer.squareup.com/docs/working-with-apis/optimistic-concurrency) control. The value must be set to the current version of the customer profile.  To update a customer profile that was created by merging existing profiles, you must use the ID of the newly created profile.  You cannot use this endpoint to change cards on file. To make changes, use the [Cards API](api:Cards) or [Gift Cards API](api:GiftCards).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateCustomerRequest**](UpdateCustomerRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **customerId** | **string**| The ID of the customer to update. | 

### Return type

[**UpdateCustomerResponse**](UpdateCustomerResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

