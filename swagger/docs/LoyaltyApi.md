# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccumulateLoyaltyPoints**](LoyaltyApi.md#AccumulateLoyaltyPoints) | **Post** /v2/loyalty/accounts/{account_id}/accumulate | AccumulateLoyaltyPoints
[**AdjustLoyaltyPoints**](LoyaltyApi.md#AdjustLoyaltyPoints) | **Post** /v2/loyalty/accounts/{account_id}/adjust | AdjustLoyaltyPoints
[**CalculateLoyaltyPoints**](LoyaltyApi.md#CalculateLoyaltyPoints) | **Post** /v2/loyalty/programs/{program_id}/calculate | CalculateLoyaltyPoints
[**CreateLoyaltyAccount**](LoyaltyApi.md#CreateLoyaltyAccount) | **Post** /v2/loyalty/accounts | CreateLoyaltyAccount
[**CreateLoyaltyReward**](LoyaltyApi.md#CreateLoyaltyReward) | **Post** /v2/loyalty/rewards | CreateLoyaltyReward
[**DeleteLoyaltyReward**](LoyaltyApi.md#DeleteLoyaltyReward) | **Delete** /v2/loyalty/rewards/{reward_id} | DeleteLoyaltyReward
[**ListLoyaltyPrograms**](LoyaltyApi.md#ListLoyaltyPrograms) | **Get** /v2/loyalty/programs | ListLoyaltyPrograms
[**RedeemLoyaltyReward**](LoyaltyApi.md#RedeemLoyaltyReward) | **Post** /v2/loyalty/rewards/{reward_id}/redeem | RedeemLoyaltyReward
[**RetrieveLoyaltyAccount**](LoyaltyApi.md#RetrieveLoyaltyAccount) | **Get** /v2/loyalty/accounts/{account_id} | RetrieveLoyaltyAccount
[**RetrieveLoyaltyProgram**](LoyaltyApi.md#RetrieveLoyaltyProgram) | **Get** /v2/loyalty/programs/{program_id} | RetrieveLoyaltyProgram
[**RetrieveLoyaltyReward**](LoyaltyApi.md#RetrieveLoyaltyReward) | **Get** /v2/loyalty/rewards/{reward_id} | RetrieveLoyaltyReward
[**SearchLoyaltyAccounts**](LoyaltyApi.md#SearchLoyaltyAccounts) | **Post** /v2/loyalty/accounts/search | SearchLoyaltyAccounts
[**SearchLoyaltyEvents**](LoyaltyApi.md#SearchLoyaltyEvents) | **Post** /v2/loyalty/events/search | SearchLoyaltyEvents
[**SearchLoyaltyRewards**](LoyaltyApi.md#SearchLoyaltyRewards) | **Post** /v2/loyalty/rewards/search | SearchLoyaltyRewards

# **AccumulateLoyaltyPoints**
> AccumulateLoyaltyPointsResponse AccumulateLoyaltyPoints(ctx, body, accountId)
AccumulateLoyaltyPoints

Adds points to a loyalty account.  - If you are using the Orders API to manage orders, you only provide the `order_id`.  The endpoint reads the order to compute points to add to the buyer's account. - If you are not using the Orders API to manage orders,  you first perform a client-side computation to compute the points.   For spend-based and visit-based programs, you can first call  [CalculateLoyaltyPoints](api-endpoint:Loyalty-CalculateLoyaltyPoints) to compute the points   that you provide to this endpoint.   __Note:__ The country of the seller's Square account determines whether tax is included in the purchase amount when accruing points for spend-based and visit-based programs.  For more information, see [Availability of Square Loyalty](https://developer.squareup.com/docs/loyalty-api/overview#loyalty-market-availability).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AccumulateLoyaltyPointsRequest**](AccumulateLoyaltyPointsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **accountId** | **string**| The [loyalty account](entity:LoyaltyAccount) ID to which to add the points. | 

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
  **accountId** | **string**| The ID of the [loyalty account](entity:LoyaltyAccount) in which to adjust the points. | 

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

Calculates the points a purchase earns.  - If you are using the Orders API to manage orders, you provide `order_id` in the request. The  endpoint calculates the points by reading the order. - If you are not using the Orders API to manage orders, you provide the purchase amount in  the request for the endpoint to calculate the points.  An application might call this endpoint to show the points that a buyer can earn with the  specific purchase.  __Note:__ The country of the seller's Square account determines whether tax is included in the purchase amount when accruing points for spend-based and visit-based programs.  For more information, see [Availability of Square Loyalty](https://developer.squareup.com/docs/loyalty-api/overview#loyalty-market-availability).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CalculateLoyaltyPointsRequest**](CalculateLoyaltyPointsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **programId** | **string**| The [loyalty program](entity:LoyaltyProgram) ID, which defines the rules for accruing points. | 

### Return type

[**CalculateLoyaltyPointsResponse**](CalculateLoyaltyPointsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
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

Searches for loyalty rewards in a loyalty account.   In the current implementation, the endpoint supports search by the reward `status`.  If you know a reward ID, use the  [RetrieveLoyaltyReward](api-endpoint:Loyalty-RetrieveLoyaltyReward) endpoint.  Search results are sorted by `updated_at` in descending order.

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

