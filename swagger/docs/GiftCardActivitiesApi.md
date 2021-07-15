# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGiftCardActivity**](GiftCardActivitiesApi.md#CreateGiftCardActivity) | **Post** /v2/gift-cards/activities | CreateGiftCardActivity
[**ListGiftCardActivities**](GiftCardActivitiesApi.md#ListGiftCardActivities) | **Get** /v2/gift-cards/activities | ListGiftCardActivities

# **CreateGiftCardActivity**
> CreateGiftCardActivityResponse CreateGiftCardActivity(ctx, body)
CreateGiftCardActivity

Creates a gift card activity. For more information, see  [GiftCardActivity](https://developer.squareup.com/docs/gift-cards/using-gift-cards-api#giftcardactivity) and  [Using activated gift cards](https://developer.squareup.com/docs/gift-cards/using-gift-cards-api#using-activated-gift-cards).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateGiftCardActivityRequest**](CreateGiftCardActivityRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateGiftCardActivityResponse**](CreateGiftCardActivityResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListGiftCardActivities**
> ListGiftCardActivitiesResponse ListGiftCardActivities(ctx, optional)
ListGiftCardActivities

Lists gift card activities. By default, you get gift card activities for all gift cards in the seller's account. You can optionally specify query parameters to filter the list. For example, you can get a list of gift card activities for a gift card, for all gift cards in a specific region, or for activities within a time window.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GiftCardActivitiesApiListGiftCardActivitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GiftCardActivitiesApiListGiftCardActivitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **giftCardId** | **optional.String**| If you provide a gift card ID, the endpoint returns activities that belong  to the specified gift card. Otherwise, the endpoint returns all gift card activities for  the seller. | 
 **type_** | **optional.String**| If you provide a type, the endpoint returns gift card activities of this type.  Otherwise, the endpoint returns all types of gift card activities. | 
 **locationId** | **optional.String**| If you provide a location ID, the endpoint returns gift card activities for that location.  Otherwise, the endpoint returns gift card activities for all locations. | 
 **beginTime** | **optional.String**| The timestamp for the beginning of the reporting period, in RFC 3339 format. Inclusive. Default: The current time minus one year. | 
 **endTime** | **optional.String**| The timestamp for the end of the reporting period, in RFC 3339 format. Inclusive. Default: The current time. | 
 **limit** | **optional.Int32**| If you provide a limit value, the endpoint returns the specified number  of results (or less) per page. A maximum value is 100. The default value is 50. | 
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. If you do not provide the cursor, the call returns the first page of the results. | 
 **sortOrder** | **optional.String**| The order in which the endpoint returns the activities, based on &#x60;created_at&#x60;. - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default). | 

### Return type

[**ListGiftCardActivitiesResponse**](ListGiftCardActivitiesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

