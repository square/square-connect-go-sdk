# AdditionalRecipientReceivable

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The additional recipient receivable&#x27;s unique ID, issued by Square payments servers. | [default to null]
**TransactionId** | **string** | The ID of the transaction that the additional recipient receivable was applied to. | [default to null]
**TransactionLocationId** | **string** | The ID of the location that created the receivable. This is the location ID on the associated transaction. | [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**CreatedAt** | **string** | The time when the additional recipient receivable was created, in RFC 3339 format. | [optional] [default to null]
**Refunds** | [**[]AdditionalRecipientReceivableRefund**](AdditionalRecipientReceivableRefund.md) | Any refunds of the receivable that have been applied. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

