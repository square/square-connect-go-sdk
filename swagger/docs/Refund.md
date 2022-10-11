# Refund

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The refund&#x27;s unique ID. | [default to null]
**LocationId** | **string** | The ID of the refund&#x27;s associated location. | [default to null]
**TransactionId** | **string** | The ID of the transaction that the refunded tender is part of. | [optional] [default to null]
**TenderId** | **string** | The ID of the refunded tender. | [default to null]
**CreatedAt** | **string** | The timestamp for when the refund was created, in RFC 3339 format. | [optional] [default to null]
**Reason** | **string** | The reason for the refund being issued. | [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**Status** | [***RefundStatus**](RefundStatus.md) |  | [default to null]
**ProcessingFeeMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**AdditionalRecipients** | [**[]AdditionalRecipient**](AdditionalRecipient.md) | Additional recipients (other than the merchant) receiving a portion of this refund. For example, fees assessed on a refund of a purchase by a third party integration. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

