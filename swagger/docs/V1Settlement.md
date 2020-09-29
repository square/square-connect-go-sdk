# V1Settlement

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The settlement&#x27;s unique identifier. | [optional] [default to null]
**Status** | [***V1SettlementStatus**](V1SettlementStatus.md) |  | [optional] [default to null]
**TotalMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**InitiatedAt** | **string** | The time when the settlement was submitted for deposit or withdrawal, in ISO 8601 format. | [optional] [default to null]
**BankAccountId** | **string** | The Square-issued unique identifier for the bank account associated with the settlement. | [optional] [default to null]
**Entries** | [**[]V1SettlementEntry**](V1SettlementEntry.md) | The entries included in this settlement. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

