# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCard**](CardsApi.md#CreateCard) | **Post** /v2/cards | CreateCard
[**DisableCard**](CardsApi.md#DisableCard) | **Post** /v2/cards/{card_id}/disable | DisableCard
[**ListCards**](CardsApi.md#ListCards) | **Get** /v2/cards | ListCards
[**RetrieveCard**](CardsApi.md#RetrieveCard) | **Get** /v2/cards/{card_id} | RetrieveCard

# **CreateCard**
> CreateCardResponse CreateCard(ctx, body)
CreateCard

Adds a card on file to an existing merchant.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateCardRequest**](CreateCardRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateCardResponse**](CreateCardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisableCard**
> DisableCardResponse DisableCard(ctx, body, cardId)
DisableCard

Disables the card, preventing any further updates or charges. Disabling an already disabled card is allowed but has no effect.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DisableCardRequest**](DisableCardRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **cardId** | **string**| Unique ID for the desired Card. | 

### Return type

[**DisableCardResponse**](DisableCardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCards**
> ListCardsResponse ListCards(ctx, optional)
ListCards

Retrieves a list of cards owned by the account making the request. A max of 25 cards will be returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***CardsApiListCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CardsApiListCardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information. | 
 **customerId** | **optional.String**| Limit results to cards associated with the customer supplied. By default, all cards owned by the merchant are returned. | 
 **includeDisabled** | **optional.Bool**| Includes disabled cards. By default, all enabled cards owned by the merchant are returned. | [default to false]
 **referenceId** | **optional.String**| Limit results to cards associated with the reference_id supplied. | 
 **sortOrder** | [**optional.Interface of SortOrder**](.md)| Sorts the returned list by when the card was created with the specified order. This field defaults to ASC. | 

### Return type

[**ListCardsResponse**](ListCardsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveCard**
> RetrieveCardResponse RetrieveCard(ctx, cardId)
RetrieveCard

Retrieves details for a specific Card.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **cardId** | **string**| Unique ID for the desired Card. | 

### Return type

[**RetrieveCardResponse**](RetrieveCardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

