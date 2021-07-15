# OrderUpdated

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | The order&#x27;s unique ID. | [optional] [default to null]
**Version** | **int32** | The version number, which is incremented each time an update is committed to the order. Orders that were not created through the API do not include a version number and therefore cannot be updated.  [Read more about working with versions.](https://developer.squareup.com/docs/orders-api/manage-orders#update-orders) | [optional] [default to null]
**LocationId** | **string** | The ID of the seller location that this order is associated with. | [optional] [default to null]
**State** | [***OrderState**](OrderState.md) |  | [optional] [default to null]
**CreatedAt** | **string** | The timestamp for when the order was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp for when the order was last updated, in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

