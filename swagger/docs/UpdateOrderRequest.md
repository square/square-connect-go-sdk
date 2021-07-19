# UpdateOrderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [***Order**](Order.md) |  | [optional] [default to null]
**FieldsToClear** | **[]string** | The [dot notation paths](https://developer.squareup.com/docs/orders-api/manage-orders#on-dot-notation) fields to clear. For example, &#x60;line_items[uid].note&#x60;. For more information, see [Deleting fields](https://developer.squareup.com/docs/orders-api/manage-orders#delete-fields). | [optional] [default to null]
**IdempotencyKey** | **string** | A value you specify that uniquely identifies this update request.  If you are unsure whether a particular update was applied to an order successfully, you can reattempt it with the same idempotency key without worrying about creating duplicate updates to the order. The latest order version is returned.  For more information, see [Idempotency](https://developer.squareup.com/docs/basics/api101/idempotency). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

