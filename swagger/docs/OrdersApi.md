# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BatchRetrieveOrders**](OrdersApi.md#BatchRetrieveOrders) | **Post** /v2/orders/batch-retrieve | BatchRetrieveOrders
[**CalculateOrder**](OrdersApi.md#CalculateOrder) | **Post** /v2/orders/calculate | CalculateOrder
[**CloneOrder**](OrdersApi.md#CloneOrder) | **Post** /v2/orders/clone | CloneOrder
[**CreateOrder**](OrdersApi.md#CreateOrder) | **Post** /v2/orders | CreateOrder
[**PayOrder**](OrdersApi.md#PayOrder) | **Post** /v2/orders/{order_id}/pay | PayOrder
[**RetrieveOrder**](OrdersApi.md#RetrieveOrder) | **Get** /v2/orders/{order_id} | RetrieveOrder
[**SearchOrders**](OrdersApi.md#SearchOrders) | **Post** /v2/orders/search | SearchOrders
[**UpdateOrder**](OrdersApi.md#UpdateOrder) | **Put** /v2/orders/{order_id} | UpdateOrder

# **BatchRetrieveOrders**
> BatchRetrieveOrdersResponse BatchRetrieveOrders(ctx, body)
BatchRetrieveOrders

Retrieves a set of [orders](entity:Order) by their IDs.  If a given order ID does not exist, the ID is ignored instead of generating an error.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BatchRetrieveOrdersRequest**](BatchRetrieveOrdersRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BatchRetrieveOrdersResponse**](BatchRetrieveOrdersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CalculateOrder**
> CalculateOrderResponse CalculateOrder(ctx, body)
CalculateOrder

Enables applications to preview order pricing without creating an order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CalculateOrderRequest**](CalculateOrderRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CalculateOrderResponse**](CalculateOrderResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CloneOrder**
> CloneOrderResponse CloneOrder(ctx, body)
CloneOrder

Creates a new order, in the `DRAFT` state, by duplicating an existing order. The newly created order has  only the core fields (such as line items, taxes, and discounts) copied from the original order.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CloneOrderRequest**](CloneOrderRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CloneOrderResponse**](CloneOrderResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrder**
> CreateOrderResponse CreateOrder(ctx, body)
CreateOrder

Creates a new [order](entity:Order) that can include information about products for purchase and settings to apply to the purchase.  To pay for a created order, see  [Pay for Orders](https://developer.squareup.com/docs/orders-api/pay-for-orders).  You can modify open orders using the [UpdateOrder](api-endpoint:Orders-UpdateOrder) endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateOrderRequest**](CreateOrderRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateOrderResponse**](CreateOrderResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PayOrder**
> PayOrderResponse PayOrder(ctx, body, orderId)
PayOrder

Pay for an [order](entity:Order) using one or more approved [payments](entity:Payment) or settle an order with a total of `0`.  The total of the `payment_ids` listed in the request must be equal to the order total. Orders with a total amount of `0` can be marked as paid by specifying an empty array of `payment_ids` in the request.  To be used with `PayOrder`, a payment must:  - Reference the order by specifying the `order_id` when [creating the payment](api-endpoint:Payments-CreatePayment). Any approved payments that reference the same `order_id` not specified in the `payment_ids` is canceled. - Be approved with [delayed capture](https://developer.squareup.com/docs/payments-api/take-payments#delayed-capture). Using a delayed capture payment with `PayOrder` completes the approved payment.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PayOrderRequest**](PayOrderRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **orderId** | **string**| The ID of the order being paid. | 

### Return type

[**PayOrderResponse**](PayOrderResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveOrder**
> RetrieveOrderResponse RetrieveOrder(ctx, orderId)
RetrieveOrder

Retrieves an [Order](entity:Order) by ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orderId** | **string**| The ID of the order to retrieve. | 

### Return type

[**RetrieveOrderResponse**](RetrieveOrderResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchOrders**
> SearchOrdersResponse SearchOrders(ctx, body)
SearchOrders

Search all orders for one or more locations. Orders include all sales, returns, and exchanges regardless of how or when they entered the Square ecosystem (such as Point of Sale, Invoices, and Connect APIs).  `SearchOrders` requests need to specify which locations to search and define a [SearchOrdersQuery](entity:SearchOrdersQuery) object that controls how to sort or filter the results. Your `SearchOrdersQuery` can:    Set filter criteria.   Set the sort order.   Determine whether to return results as complete `Order` objects or as [OrderEntry](entity:OrderEntry) objects.  Note that details for orders processed with Square Point of Sale while in offline mode might not be transmitted to Square for up to 72 hours. Offline orders have a `created_at` value that reflects the time the order was created, not the time it was subsequently transmitted to Square.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchOrdersRequest**](SearchOrdersRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchOrdersResponse**](SearchOrdersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrder**
> UpdateOrderResponse UpdateOrder(ctx, body, orderId)
UpdateOrder

Updates an open [order](entity:Order) by adding, replacing, or deleting fields. Orders with a `COMPLETED` or `CANCELED` state cannot be updated.  An `UpdateOrder` request requires the following:  - The `order_id` in the endpoint path, identifying the order to update. - The latest `version` of the order to update. - The [sparse order](https://developer.squareup.com/docs/orders-api/manage-orders#sparse-order-objects) containing only the fields to update and the version to which the update is being applied. - If deleting fields, the [dot notation paths](https://developer.squareup.com/docs/orders-api/manage-orders#on-dot-notation) identifying the fields to clear.  To pay for an order, see  [Pay for Orders](https://developer.squareup.com/docs/orders-api/pay-for-orders).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateOrderRequest**](UpdateOrderRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **orderId** | **string**| The ID of the order to update. | 

### Return type

[**UpdateOrderResponse**](UpdateOrderResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

