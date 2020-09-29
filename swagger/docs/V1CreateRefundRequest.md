# V1CreateRefundRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** | The ID of the payment to refund. If you are creating a &#x60;PARTIAL&#x60; refund for a split tender payment, instead provide the id of the particular tender you want to refund. | [default to null]
**Type_** | [***V1CreateRefundRequestType**](V1CreateRefundRequestType.md) |  | [default to null]
**Reason** | **string** | The reason for the refund. | [default to null]
**RefundedMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RequestIdempotenceKey** | **string** | An optional key to ensure idempotence if you issue the same PARTIAL refund request more than once. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

