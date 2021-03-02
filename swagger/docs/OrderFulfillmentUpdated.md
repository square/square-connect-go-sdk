# OrderFulfillmentUpdated

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | The order&#x27;s unique ID. | [optional] [default to null]
**Version** | **int32** | Version number which is incremented each time an update is committed to the order. Orders that were not created through the API will not include a version and thus cannot be updated.  [Read more about working with versions](https://developer.squareup.com/docs/orders-api/manage-orders#update-orders) | [optional] [default to null]
**LocationId** | **string** | The ID of the merchant location this order is associated with. | [optional] [default to null]
**State** | [***OrderState**](OrderState.md) |  | [optional] [default to null]
**CreatedAt** | **string** | Timestamp for when the order was created in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | Timestamp for when the order was last updated in RFC 3339 format. | [optional] [default to null]
**FulfillmentUpdate** | [**[]OrderFulfillmentUpdatedUpdate**](OrderFulfillmentUpdatedUpdate.md) | The fulfillments that were updated with this version change. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

