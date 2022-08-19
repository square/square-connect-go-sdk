# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGiftCard**](GiftCardsApi.md#CreateGiftCard) | **Post** /v2/gift-cards | CreateGiftCard
[**LinkCustomerToGiftCard**](GiftCardsApi.md#LinkCustomerToGiftCard) | **Post** /v2/gift-cards/{gift_card_id}/link-customer | LinkCustomerToGiftCard
[**ListGiftCards**](GiftCardsApi.md#ListGiftCards) | **Get** /v2/gift-cards | ListGiftCards
[**RetrieveGiftCard**](GiftCardsApi.md#RetrieveGiftCard) | **Get** /v2/gift-cards/{id} | RetrieveGiftCard
[**RetrieveGiftCardFromGAN**](GiftCardsApi.md#RetrieveGiftCardFromGAN) | **Post** /v2/gift-cards/from-gan | RetrieveGiftCardFromGAN
[**RetrieveGiftCardFromNonce**](GiftCardsApi.md#RetrieveGiftCardFromNonce) | **Post** /v2/gift-cards/from-nonce | RetrieveGiftCardFromNonce
[**UnlinkCustomerFromGiftCard**](GiftCardsApi.md#UnlinkCustomerFromGiftCard) | **Post** /v2/gift-cards/{gift_card_id}/unlink-customer | UnlinkCustomerFromGiftCard

# **CreateGiftCard**
> CreateGiftCardResponse CreateGiftCard(ctx, body)
CreateGiftCard

Creates a digital gift card or registers a physical (plastic) gift card. After the gift card  is created, you must call [CreateGiftCardActivity](api-endpoint:GiftCardActivities-CreateGiftCardActivity)  to activate the card with an initial balance before it can be used for payment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateGiftCardRequest**](CreateGiftCardRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateGiftCardResponse**](CreateGiftCardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkCustomerToGiftCard**
> LinkCustomerToGiftCardResponse LinkCustomerToGiftCard(ctx, body, giftCardId)
LinkCustomerToGiftCard

Links a customer to a gift card, which is also referred to as adding a card on file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LinkCustomerToGiftCardRequest**](LinkCustomerToGiftCardRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **giftCardId** | **string**| The ID of the gift card to be linked. | 

### Return type

[**LinkCustomerToGiftCardResponse**](LinkCustomerToGiftCardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGiftCards**
> ListGiftCardsResponse ListGiftCards(ctx, optional)
ListGiftCards

Lists all gift cards. You can specify optional filters to retrieve  a subset of the gift cards. Results are sorted by `created_at` in ascending order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GiftCardsApiListGiftCardsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GiftCardsApiListGiftCardsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| If a [type](entity:GiftCardType) is provided, the endpoint returns gift cards of the specified type. Otherwise, the endpoint returns gift cards of all types. | 
 **state** | **optional.String**| If a [state](entity:GiftCardStatus) is provided, the endpoint returns the gift cards in the specified state. Otherwise, the endpoint returns the gift cards of all states. | 
 **limit** | **optional.Int32**| If a limit is provided, the endpoint returns only the specified number of results per page. The maximum value is 50. The default value is 30. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. If a cursor is not provided, the endpoint returns the first page of the results.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination). | 
 **customerId** | **optional.String**| If a customer ID is provided, the endpoint returns only the gift cards linked to the specified customer. | 

### Return type

[**ListGiftCardsResponse**](ListGiftCardsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveGiftCard**
> RetrieveGiftCardResponse RetrieveGiftCard(ctx, id)
RetrieveGiftCard

Retrieves a gift card using the gift card ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the gift card to retrieve. | 

### Return type

[**RetrieveGiftCardResponse**](RetrieveGiftCardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveGiftCardFromGAN**
> RetrieveGiftCardFromGanResponse RetrieveGiftCardFromGAN(ctx, body)
RetrieveGiftCardFromGAN

Retrieves a gift card using the gift card account number (GAN).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RetrieveGiftCardFromGanRequest**](RetrieveGiftCardFromGanRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**RetrieveGiftCardFromGanResponse**](RetrieveGiftCardFromGANResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveGiftCardFromNonce**
> RetrieveGiftCardFromNonceResponse RetrieveGiftCardFromNonce(ctx, body)
RetrieveGiftCardFromNonce

Retrieves a gift card using a secure payment token that represents the gift card.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RetrieveGiftCardFromNonceRequest**](RetrieveGiftCardFromNonceRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**RetrieveGiftCardFromNonceResponse**](RetrieveGiftCardFromNonceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkCustomerFromGiftCard**
> UnlinkCustomerFromGiftCardResponse UnlinkCustomerFromGiftCard(ctx, body, giftCardId)
UnlinkCustomerFromGiftCard

Unlinks a customer from a gift card, which is also referred to as removing a card on file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UnlinkCustomerFromGiftCardRequest**](UnlinkCustomerFromGiftCardRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **giftCardId** | **string**| The ID of the gift card to be unlinked. | 

### Return type

[**UnlinkCustomerFromGiftCardResponse**](UnlinkCustomerFromGiftCardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

