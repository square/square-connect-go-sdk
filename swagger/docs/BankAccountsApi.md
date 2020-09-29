# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBankAccount**](BankAccountsApi.md#GetBankAccount) | **Get** /v2/bank-accounts/{bank_account_id} | GetBankAccount
[**GetBankAccountByV1Id**](BankAccountsApi.md#GetBankAccountByV1Id) | **Get** /v2/bank-accounts/by-v1-id/{v1_bank_account_id} | GetBankAccountByV1Id
[**ListBankAccounts**](BankAccountsApi.md#ListBankAccounts) | **Get** /v2/bank-accounts | ListBankAccounts

# **GetBankAccount**
> GetBankAccountResponse GetBankAccount(ctx, bankAccountId)
GetBankAccount

Returns details of a [BankAccount](#type-bankaccount)  linked to a Square account. For more information, see  [Bank Accounts API](https://developer.squareup.com/docs/docs/bank-accounts-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bankAccountId** | **string**| Square-issued ID of the desired &#x60;BankAccount&#x60;. | 

### Return type

[**GetBankAccountResponse**](GetBankAccountResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBankAccountByV1Id**
> GetBankAccountByV1IdResponse GetBankAccountByV1Id(ctx, v1BankAccountId)
GetBankAccountByV1Id

Returns details of a [BankAccount](#type-bankaccount) identified by V1 bank account ID.  For more information, see  [Retrieve a bank account by using an ID issued by V1 Bank Accounts API](https://developer.squareup.com/docs/docs/bank-accounts-api#retrieve-a-bank-account-by-using-an-id-issued-by-the-v1-bank-accounts-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **v1BankAccountId** | **string**| Connect V1 ID of the desired &#x60;BankAccount&#x60;. For more information, see  [Retrieve a bank account by using an ID issued by V1 Bank Accounts API](https://developer.squareup.com/docs/docs/bank-accounts-api#retrieve-a-bank-account-by-using-an-id-issued-by-v1-bank-accounts-api). | 

### Return type

[**GetBankAccountByV1IdResponse**](GetBankAccountByV1IdResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListBankAccounts**
> ListBankAccountsResponse ListBankAccounts(ctx, optional)
ListBankAccounts

Returns a list of [BankAccount](#type-bankaccount) objects linked to a Square account.  For more information, see  [Bank Accounts API](https://developer.squareup.com/docs/docs/bank-accounts-api).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BankAccountsApiListBankAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BankAccountsApiListBankAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| The pagination cursor returned by a previous call to this endpoint. Use it in the next &#x60;ListBankAccounts&#x60; request to retrieve the next set  of results.  See the [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination) guide for more information. | 
 **limit** | **optional.Int32**| Upper limit on the number of bank accounts to return in the response.  Currently, 1000 is the largest supported limit. You can specify a limit  of up to 1000 bank accounts. This is also the default limit. | 
 **locationId** | **optional.String**| Location ID. You can specify this optional filter  to retrieve only the linked bank accounts belonging to a specific location. | 

### Return type

[**ListBankAccountsResponse**](ListBankAccountsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

