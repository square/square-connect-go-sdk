# CloneOrderRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** | The ID of the order to clone. | [default to null]
**Version** | **int32** | An optional order version for concurrency protection.  If a version is provided, it must match the latest stored version of the order to clone.  If a version is not provided, the API clones the latest version. | [optional] [default to null]
**IdempotencyKey** | **string** | A value you specify that uniquely identifies this clone request.  If you are unsure whether a particular order was cloned successfully, you can reattempt the call with the same idempotency key without worrying about creating duplicate cloned orders. The originally cloned order is returned.  For more information, see [Idempotency](https://developer.squareup.com/docs/basics/api101/idempotency). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

