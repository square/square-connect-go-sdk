# AdditionalRecipientReceivableRefund

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The receivable refund&#x27;s unique ID, issued by Square payments servers. | [default to null]
**ReceivableId** | **string** | The ID of the receivable that the refund was applied to. | [default to null]
**RefundId** | **string** | The ID of the refund that is associated to this receivable refund. | [default to null]
**TransactionLocationId** | **string** | The ID of the location that created the receivable. This is the location ID on the associated transaction. | [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**CreatedAt** | **string** | The time when the refund was created, in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

