# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccumulateLoyaltyPoints**](LoyaltyApi.md#AccumulateLoyaltyPoints) | **Post** /v2/loyalty/accounts/{account_id}/accumulate | AccumulateLoyaltyPoints
[**AdjustLoyaltyPoints**](LoyaltyApi.md#AdjustLoyaltyPoints) | **Post** /v2/loyalty/accounts/{account_id}/adjust | AdjustLoyaltyPoints
[**CalculateLoyaltyPoints**](LoyaltyApi.md#CalculateLoyaltyPoints) | **Post** /v2/loyalty/programs/{program_id}/calculate | CalculateLoyaltyPoints
[**CancelLoyaltyPromotion**](LoyaltyApi.md#CancelLoyaltyPromotion) | **Post** /v2/loyalty/programs/{program_id}/promotions/{promotion_id}/cancel | CancelLoyaltyPromotion
[**CreateLoyaltyAccount**](LoyaltyApi.md#CreateLoyaltyAccount) | **Post** /v2/loyalty/accounts | CreateLoyaltyAccount
[**CreateLoyaltyPromotion**](LoyaltyApi.md#CreateLoyaltyPromotion) | **Post** /v2/loyalty/programs/{program_id}/promotions | CreateLoyaltyPromotion
[**CreateLoyaltyReward**](LoyaltyApi.md#CreateLoyaltyReward) | **Post** /v2/loyalty/rewards | CreateLoyaltyReward
[**DeleteLoyaltyReward**](LoyaltyApi.md#DeleteLoyaltyReward) | **Delete** /v2/loyalty/rewards/{reward_id} | DeleteLoyaltyReward
[**ListLoyaltyPrograms**](LoyaltyApi.md#ListLoyaltyPrograms) | **Get** /v2/loyalty/programs | ListLoyaltyPrograms
[**ListLoyaltyPromotions**](LoyaltyApi.md#ListLoyaltyPromotions) | **Get** /v2/loyalty/programs/{program_id}/promotions | ListLoyaltyPromotions
[**RedeemLoyaltyReward**](LoyaltyApi.md#RedeemLoyaltyReward) | **Post** /v2/loyalty/rewards/{reward_id}/redeem | RedeemLoyaltyReward
[**RetrieveLoyaltyAccount**](LoyaltyApi.md#RetrieveLoyaltyAccount) | **Get** /v2/loyalty/accounts/{account_id} | RetrieveLoyaltyAccount
[**RetrieveLoyaltyProgram**](LoyaltyApi.md#RetrieveLoyaltyProgram) | **Get** /v2/loyalty/programs/{program_id} | RetrieveLoyaltyProgram
[**RetrieveLoyaltyPromotion**](LoyaltyApi.md#RetrieveLoyaltyPromotion) | **Get** /v2/loyalty/programs/{program_id}/promotions/{promotion_id} | RetrieveLoyaltyPromotion
[**RetrieveLoyaltyReward**](LoyaltyApi.md#RetrieveLoyaltyReward) | **Get** /v2/loyalty/rewards/{reward_id} | RetrieveLoyaltyReward
[**SearchLoyaltyAccounts**](LoyaltyApi.md#SearchLoyaltyAccounts) | **Post** /v2/loyalty/accounts/search | SearchLoyaltyAccounts
[**SearchLoyaltyEvents**](LoyaltyApi.md#SearchLoyaltyEvents) | **Post** /v2/loyalty/events/search | SearchLoyaltyEvents
[**SearchLoyaltyRewards**](LoyaltyApi.md#SearchLoyaltyRewards) | **Post** /v2/loyalty/rewards/search | SearchLoyaltyRewards

# **AccumulateLoyaltyPoints**
> AccumulateLoyaltyPointsResponse AccumulateLoyaltyPoints(ctx, body, accountId)
AccumulateLoyaltyPoints

Adds points earned from a purchase to a [loyalty account](entity:LoyaltyAccount).  - If you are using the Orders API to manage orders, provide the `order_id`. Square reads the order to compute the points earned from both the base loyalty program and an associated [loyalty promotion](entity:LoyaltyPromotion). For purchases that qualify for multiple accrual rules, Square computes points based on the accrual rule that grants the most points. For purchases that qualify for multiple promotions, Square computes points based on the most recently created promotion. A purchase must first qualify for program points to be eligible for promotion points.  - If you are not using the Orders API to manage orders, provide `points` with the number of points to add. You must first perform a client-side computation of the points earned from the loyalty program and loyalty promotion. For spend-based and visit-based programs, you can call [CalculateLoyaltyPoints](api-endpoint:Loyalty-CalculateLoyaltyPoints) to compute the points earned from the base loyalty program. For information about computing points earned from a loyalty promotion, see [Calculating promotion points](https://developer.squareup.com/docs/loyalty-api/loyalty-promotions#calculate-promotion-points).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccumulateLoyaltyPointsRequest**](AccumulateLoyaltyPointsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **accountId** | **string**| The ID of the target [loyalty account](entity:LoyaltyAccount). | 

### Return type

[**AccumulateLoyaltyPointsResponse**](AccumulateLoyaltyPointsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdjustLoyaltyPoints**
> AdjustLoyaltyPointsResponse AdjustLoyaltyPoints(ctx, body, accountId)
AdjustLoyaltyPoints

Adds points to or subtracts points from a buyer's account.   Use this endpoint only when you need to manually adjust points. Otherwise, in your application flow, you call  [AccumulateLoyaltyPoints](api-endpoint:Loyalty-AccumulateLoyaltyPoints)  to add points when a buyer pays for the purchase.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdjustLoyaltyPointsRequest**](AdjustLoyaltyPointsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **accountId** | **string**| The ID of the target [loyalty account](entity:LoyaltyAccount). | 

### Return type

[**AdjustLoyaltyPointsResponse**](AdjustLoyaltyPointsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CalculateLoyaltyPoints**
> CalculateLoyaltyPointsResponse CalculateLoyaltyPoints(ctx, body, programId)
CalculateLoyaltyPoints

Calculates the number of points a buyer can earn from a purchase. Applications might call this endpoint to display the points to the buyer.  - If you are using the Orders API to manage orders, provide the `order_id` and (optional) `loyalty_account_id`. Square reads the order to compute the points earned from the base loyalty program and an associated [loyalty promotion](entity:LoyaltyPromotion).  - If you are not using the Orders API to manage orders, provide `transaction_amount_money` with the purchase amount. Square uses this amount to calculate the points earned from the base loyalty program, but not points earned from a loyalty promotion. For spend-based and visit-based programs, the `tax_mode` setting of the accrual rule indicates how taxes should be treated for loyalty points accrual. If the purchase qualifies for program points, call [ListLoyaltyPromotions](api-endpoint:Loyalty-ListLoyaltyPromotions) and perform a client-side computation to calculate whether the purchase also qualifies for promotion points. For more information, see [Calculating promotion points](https://developer.squareup.com/docs/loyalty-api/loyalty-promotions#calculate-promotion-points).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CalculateLoyaltyPointsRequest**](CalculateLoyaltyPointsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **programId** | **string**| The ID of the [loyalty program](entity:LoyaltyProgram), which defines the rules for accruing points. | 

### Return type

[**CalculateLoyaltyPointsResponse**](CalculateLoyaltyPointsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CancelLoyaltyPromotion**
> CancelLoyaltyPromotionResponse CancelLoyaltyPromotion(ctx, promotionId, programId)
CancelLoyaltyPromotion

Cancels a loyalty promotion. Use this endpoint to cancel an `ACTIVE` promotion earlier than the end date, cancel an `ACTIVE` promotion when an end date is not specified, or cancel a `SCHEDULED` promotion. Because updating a promotion is not supported, you can also use this endpoint to cancel a promotion before you create a new one.  This endpoint sets the loyalty promotion to the `CANCELED` state

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **promotionId** | **string**| The ID of the [loyalty promotion](entity:LoyaltyPromotion) to cancel. You can cancel a promotion that has an &#x60;ACTIVE&#x60; or &#x60;SCHEDULED&#x60; status. | 
  **programId** | **string**| The ID of the base [loyalty program](entity:LoyaltyProgram). | 

### Return type

[**CancelLoyaltyPromotionResponse**](CancelLoyaltyPromotionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoyaltyAccount**
> CreateLoyaltyAccountResponse CreateLoyaltyAccount(ctx, body)
CreateLoyaltyAccount

Creates a loyalty account. To create a loyalty account, you must provide the `program_id` and a `mapping` with the `phone_number` of the buyer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateLoyaltyAccountRequest**](CreateLoyaltyAccountRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateLoyaltyAccountResponse**](CreateLoyaltyAccountResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoyaltyPromotion**
> CreateLoyaltyPromotionResponse CreateLoyaltyPromotion(ctx, body, programId)
CreateLoyaltyPromotion

Creates a loyalty promotion for a [loyalty program](entity:LoyaltyProgram). A loyalty promotion enables buyers to earn points in addition to those earned from the base loyalty program.  This endpoint sets the loyalty promotion to the `ACTIVE` or `SCHEDULED` status, depending on the `available_time` setting. A loyalty program can have a maximum of 10 loyalty promotions with an `ACTIVE` or `SCHEDULED` status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateLoyaltyPromotionRequest**](CreateLoyaltyPromotionRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **programId** | **string**| The ID of the [loyalty program](entity:LoyaltyProgram) to associate with the promotion. To get the program ID, call [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram) using the &#x60;main&#x60; keyword. | 

### Return type

[**CreateLoyaltyPromotionResponse**](CreateLoyaltyPromotionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLoyaltyReward**
> CreateLoyaltyRewardResponse CreateLoyaltyReward(ctx, body)
CreateLoyaltyReward

Creates a loyalty reward. In the process, the endpoint does following:  - Uses the `reward_tier_id` in the request to determine the number of points  to lock for this reward.  - If the request includes `order_id`, it adds the reward and related discount to the order.   After a reward is created, the points are locked and  not available for the buyer to redeem another reward.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateLoyaltyRewardRequest**](CreateLoyaltyRewardRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateLoyaltyRewardResponse**](CreateLoyaltyRewardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLoyaltyReward**
> DeleteLoyaltyRewardResponse DeleteLoyaltyReward(ctx, rewardId)
DeleteLoyaltyReward

Deletes a loyalty reward by doing the following:  - Returns the loyalty points back to the loyalty account. - If an order ID was specified when the reward was created  (see [CreateLoyaltyReward](api-endpoint:Loyalty-CreateLoyaltyReward)),  it updates the order by removing the reward and related  discounts.  You cannot delete a reward that has reached the terminal state (REDEEMED).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rewardId** | **string**| The ID of the [loyalty reward](entity:LoyaltyReward) to delete. | 

### Return type

[**DeleteLoyaltyRewardResponse**](DeleteLoyaltyRewardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoyaltyPrograms**
> ListLoyaltyProgramsResponse ListLoyaltyPrograms(ctx, )
ListLoyaltyPrograms

Returns a list of loyalty programs in the seller's account. Loyalty programs define how buyers can earn points and redeem points for rewards. Square sellers can have only one loyalty program, which is created and managed from the Seller Dashboard. For more information, see [Loyalty Program Overview](https://developer.squareup.com/docs/loyalty/overview).   Replaced with [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram) when used with the keyword `main`.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ListLoyaltyProgramsResponse**](ListLoyaltyProgramsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLoyaltyPromotions**
> ListLoyaltyPromotionsResponse ListLoyaltyPromotions(ctx, programId, optional)
ListLoyaltyPromotions

Lists the loyalty promotions associated with a [loyalty program](entity:LoyaltyProgram). Results are sorted by the `created_at` date in descending order (newest to oldest).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **programId** | **string**| The ID of the base [loyalty program](entity:LoyaltyProgram). To get the program ID, call [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram) using the &#x60;main&#x60; keyword. | 
 **optional** | ***LoyaltyApiListLoyaltyPromotionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LoyaltyApiListLoyaltyPromotionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**optional.Interface of LoyaltyPromotionStatus**](.md)| The status to filter the results by. If a status is provided, only loyalty promotions with the specified status are returned. Otherwise, all loyalty promotions associated with the loyalty program are returned. | 
 **cursor** | **optional.String**| The cursor returned in the paged response from the previous call to this endpoint. Provide this cursor to retrieve the next page of results for your original request. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 
 **limit** | **optional.Int32**| The maximum number of results to return in a single paged response. The minimum value is 1 and the maximum value is 30. The default value is 30. For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | 

### Return type

[**ListLoyaltyPromotionsResponse**](ListLoyaltyPromotionsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RedeemLoyaltyReward**
> RedeemLoyaltyRewardResponse RedeemLoyaltyReward(ctx, body, rewardId)
RedeemLoyaltyReward

Redeems a loyalty reward.  The endpoint sets the reward to the `REDEEMED` terminal state.   If you are using your own order processing system (not using the  Orders API), you call this endpoint after the buyer paid for the  purchase.  After the reward reaches the terminal state, it cannot be deleted.  In other words, points used for the reward cannot be returned  to the account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RedeemLoyaltyRewardRequest**](RedeemLoyaltyRewardRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **rewardId** | **string**| The ID of the [loyalty reward](entity:LoyaltyReward) to redeem. | 

### Return type

[**RedeemLoyaltyRewardResponse**](RedeemLoyaltyRewardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLoyaltyAccount**
> RetrieveLoyaltyAccountResponse RetrieveLoyaltyAccount(ctx, accountId)
RetrieveLoyaltyAccount

Retrieves a loyalty account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **accountId** | **string**| The ID of the [loyalty account](entity:LoyaltyAccount) to retrieve. | 

### Return type

[**RetrieveLoyaltyAccountResponse**](RetrieveLoyaltyAccountResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLoyaltyProgram**
> RetrieveLoyaltyProgramResponse RetrieveLoyaltyProgram(ctx, programId)
RetrieveLoyaltyProgram

Retrieves the loyalty program in a seller's account, specified by the program ID or the keyword `main`.   Loyalty programs define how buyers can earn points and redeem points for rewards. Square sellers can have only one loyalty program, which is created and managed from the Seller Dashboard. For more information, see [Loyalty Program Overview](https://developer.squareup.com/docs/loyalty/overview).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **programId** | **string**| The ID of the loyalty program or the keyword &#x60;main&#x60;. Either value can be used to retrieve the single loyalty program that belongs to the seller. | 

### Return type

[**RetrieveLoyaltyProgramResponse**](RetrieveLoyaltyProgramResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLoyaltyPromotion**
> RetrieveLoyaltyPromotionResponse RetrieveLoyaltyPromotion(ctx, promotionId, programId)
RetrieveLoyaltyPromotion

Retrieves a loyalty promotion.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **promotionId** | **string**| The ID of the [loyalty promotion](entity:LoyaltyPromotion) to retrieve. | 
  **programId** | **string**| The ID of the base [loyalty program](entity:LoyaltyProgram). To get the program ID, call [RetrieveLoyaltyProgram](api-endpoint:Loyalty-RetrieveLoyaltyProgram) using the &#x60;main&#x60; keyword. | 

### Return type

[**RetrieveLoyaltyPromotionResponse**](RetrieveLoyaltyPromotionResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveLoyaltyReward**
> RetrieveLoyaltyRewardResponse RetrieveLoyaltyReward(ctx, rewardId)
RetrieveLoyaltyReward

Retrieves a loyalty reward.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **rewardId** | **string**| The ID of the [loyalty reward](entity:LoyaltyReward) to retrieve. | 

### Return type

[**RetrieveLoyaltyRewardResponse**](RetrieveLoyaltyRewardResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLoyaltyAccounts**
> SearchLoyaltyAccountsResponse SearchLoyaltyAccounts(ctx, body)
SearchLoyaltyAccounts

Searches for loyalty accounts in a loyalty program.    You can search for a loyalty account using the phone number or customer ID associated with the account. To return all loyalty accounts, specify an empty `query` object or omit it entirely.    Search results are sorted by `created_at` in ascending order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchLoyaltyAccountsRequest**](SearchLoyaltyAccountsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchLoyaltyAccountsResponse**](SearchLoyaltyAccountsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLoyaltyEvents**
> SearchLoyaltyEventsResponse SearchLoyaltyEvents(ctx, body)
SearchLoyaltyEvents

Searches for loyalty events.  A Square loyalty program maintains a ledger of events that occur during the lifetime of a  buyer's loyalty account. Each change in the point balance  (for example, points earned, points redeemed, and points expired) is  recorded in the ledger. Using this endpoint, you can search the ledger for events.  Search results are sorted by `created_at` in descending order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchLoyaltyEventsRequest**](SearchLoyaltyEventsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchLoyaltyEventsResponse**](SearchLoyaltyEventsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchLoyaltyRewards**
> SearchLoyaltyRewardsResponse SearchLoyaltyRewards(ctx, body)
SearchLoyaltyRewards

Searches for loyalty rewards. This endpoint accepts a request with no query filters and returns results for all loyalty accounts.  If you include a `query` object, `loyalty_account_id` is required and `status` is  optional.  If you know a reward ID, use the  [RetrieveLoyaltyReward](api-endpoint:Loyalty-RetrieveLoyaltyReward) endpoint.  Search results are sorted by `updated_at` in descending order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchLoyaltyRewardsRequest**](SearchLoyaltyRewardsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchLoyaltyRewardsResponse**](SearchLoyaltyRewardsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

