# CreateOrderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | [***Order**](Order.md) |  | [optional] [default to null]
**LocationId** | **string** | The ID of the business location to associate the order with. | [optional] [default to null]
**IdempotencyKey** | **string** | A value you specify that uniquely identifies this order among orders you&#x27;ve created.  If you&#x27;re unsure whether a particular order was created successfully, you can reattempt it with the same idempotency key without worrying about creating duplicate orders.  See [Idempotency](https://developer.squareup.com/docs/basics/api101/idempotency) for more information. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

