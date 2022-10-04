# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1CreateRefund**](V1TransactionsApi.md#V1CreateRefund) | **Post** /v1/{location_id}/refunds | V1CreateRefund
[**V1ListPayments**](V1TransactionsApi.md#V1ListPayments) | **Get** /v1/{location_id}/payments | V1ListPayments
[**V1ListRefunds**](V1TransactionsApi.md#V1ListRefunds) | **Get** /v1/{location_id}/refunds | V1ListRefunds
[**V1ListSettlements**](V1TransactionsApi.md#V1ListSettlements) | **Get** /v1/{location_id}/settlements | V1ListSettlements
[**V1RetrievePayment**](V1TransactionsApi.md#V1RetrievePayment) | **Get** /v1/{location_id}/payments/{payment_id} | V1RetrievePayment
[**V1RetrieveSettlement**](V1TransactionsApi.md#V1RetrieveSettlement) | **Get** /v1/{location_id}/settlements/{settlement_id} | V1RetrieveSettlement

# **V1CreateRefund**
> V1Refund V1CreateRefund(ctx, body, locationId)
V1CreateRefund

Issues a refund for a previously processed payment. You must issue a refund within 60 days of the associated payment.  You cannot issue a partial refund for a split tender payment. You must instead issue a full or partial refund for a particular tender, by providing the applicable tender id to the V1CreateRefund endpoint. Issuing a full refund for a split tender payment refunds all tenders associated with the payment.  Issuing a refund for a card payment is not reversible. For development purposes, you can create fake cash payments in Square Point of Sale and refund them.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1CreateRefundRequest**](V1CreateRefundRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the original payment&#x27;s associated location. | 

### Return type

[**V1Refund**](V1Refund.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ListPayments**
> []V1Payment V1ListPayments(ctx, locationId, optional)
V1ListPayments

Provides summary information for all payments taken for a given Square account during a date range. Date ranges cannot exceed 1 year in length. See Date ranges for details of inclusive and exclusive dates.  *Note**: Details for payments processed with Square Point of Sale while in offline mode may not be transmitted to Square for up to 72 hours. Offline payments have a `created_at` value that reflects the time the payment was originally processed, not the time it was subsequently transmitted to Square. Consequently, the ListPayments endpoint might list an offline payment chronologically between online payments that were seen in a previous request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list payments for. If you specify me, this endpoint returns payments aggregated from all of the business&#x27;s locations. | 
 **optional** | ***V1TransactionsApiV1ListPaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1TransactionsApiV1ListPaymentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | [**optional.Interface of SortOrder**](.md)| The order in which payments are listed in the response. | 
 **beginTime** | **optional.String**| The beginning of the requested reporting period, in ISO 8601 format. If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint returns an error. Default value: The current time minus one year. | 
 **endTime** | **optional.String**| The end of the requested reporting period, in ISO 8601 format. If this value is more than one year greater than begin_time, this endpoint returns an error. Default value: The current time. | 
 **limit** | **optional.Int32**| The maximum number of payments to return in a single response. This value cannot exceed 200. | 
 **batchToken** | **optional.String**| A pagination cursor to retrieve the next set of results for your original query to the endpoint. | 
 **includePartial** | **optional.Bool**| Indicates whether or not to include partial payments in the response. Partial payments will have the tenders collected so far, but the itemizations will be empty until the payment is completed. | [default to false]

### Return type

[**[]V1Payment**](V1Payment.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ListRefunds**
> []V1Refund V1ListRefunds(ctx, locationId, optional)
V1ListRefunds

Provides the details for all refunds initiated by a merchant or any of the merchant's mobile staff during a date range. Date ranges cannot exceed one year in length.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list refunds for. | 
 **optional** | ***V1TransactionsApiV1ListRefundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1TransactionsApiV1ListRefundsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | [**optional.Interface of SortOrder**](.md)| The order in which payments are listed in the response. | 
 **beginTime** | **optional.String**| The beginning of the requested reporting period, in ISO 8601 format. If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint returns an error. Default value: The current time minus one year. | 
 **endTime** | **optional.String**| The end of the requested reporting period, in ISO 8601 format. If this value is more than one year greater than begin_time, this endpoint returns an error. Default value: The current time. | 
 **limit** | **optional.Int32**| The approximate number of refunds to return in a single response. Default: 100. Max: 200. Response may contain more results than the prescribed limit when refunds are made simultaneously to multiple tenders in a payment or when refunds are generated in an exchange to account for the value of returned goods. | 
 **batchToken** | **optional.String**| A pagination cursor to retrieve the next set of results for your original query to the endpoint. | 

### Return type

[**[]V1Refund**](V1Refund.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1ListSettlements**
> []V1Settlement V1ListSettlements(ctx, locationId, optional)
V1ListSettlements

Provides summary information for all deposits and withdrawals initiated by Square to a linked bank account during a date range. Date ranges cannot exceed one year in length.  *Note**: the ListSettlements endpoint does not provide entry information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list settlements for. If you specify me, this endpoint returns settlements aggregated from all of the business&#x27;s locations. | 
 **optional** | ***V1TransactionsApiV1ListSettlementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1TransactionsApiV1ListSettlementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | [**optional.Interface of SortOrder**](.md)| The order in which settlements are listed in the response. | 
 **beginTime** | **optional.String**| The beginning of the requested reporting period, in ISO 8601 format. If this value is before January 1, 2013 (2013-01-01T00:00:00Z), this endpoint returns an error. Default value: The current time minus one year. | 
 **endTime** | **optional.String**| The end of the requested reporting period, in ISO 8601 format. If this value is more than one year greater than begin_time, this endpoint returns an error. Default value: The current time. | 
 **limit** | **optional.Int32**| The maximum number of settlements to return in a single response. This value cannot exceed 200. | 
 **status** | [**optional.Interface of V1ListSettlementsRequestStatus**](.md)| Provide this parameter to retrieve only settlements with a particular status (SENT or FAILED). | 
 **batchToken** | **optional.String**| A pagination cursor to retrieve the next set of results for your original query to the endpoint. | 

### Return type

[**[]V1Settlement**](V1Settlement.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1RetrievePayment**
> V1Payment V1RetrievePayment(ctx, locationId, paymentId)
V1RetrievePayment

Provides comprehensive information for a single payment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the payment&#x27;s associated location. | 
  **paymentId** | **string**| The Square-issued payment ID. payment_id comes from Payment objects returned by the List Payments endpoint, Settlement objects returned by the List Settlements endpoint, or Refund objects returned by the List Refunds endpoint. | 

### Return type

[**V1Payment**](V1Payment.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1RetrieveSettlement**
> V1Settlement V1RetrieveSettlement(ctx, locationId, settlementId)
V1RetrieveSettlement

Provides comprehensive information for a single settlement.  The returned `Settlement` objects include an `entries` field that lists the transactions that contribute to the settlement total. Most settlement entries correspond to a payment payout, but settlement entries are also generated for less common events, like refunds, manual adjustments, or chargeback holds.  Square initiates its regular deposits as indicated in the [Deposit Options with Square](https://squareup.com/help/us/en/article/3807) help article. Details for a regular deposit are usually not available from Connect API endpoints before 10 p.m. PST the same day.  Square does not know when an initiated settlement **completes**, only whether it has failed. A completed settlement is typically reflected in a bank account within 3 business days, but in exceptional cases it may take longer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the settlements&#x27;s associated location. | 
  **settlementId** | **string**| The settlement&#x27;s Square-issued ID. You obtain this value from Settlement objects returned by the List Settlements endpoint. | 

### Return type

[**V1Settlement**](V1Settlement.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

