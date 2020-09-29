# V1Tender

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The tender&#x27;s unique ID. | [optional] [default to null]
**Type_** | [***V1TenderType**](V1TenderType.md) |  | [optional] [default to null]
**Name** | **string** | A human-readable description of the tender. | [optional] [default to null]
**EmployeeId** | **string** | The ID of the employee that processed the tender. | [optional] [default to null]
**ReceiptUrl** | **string** | The URL of the receipt for the tender. | [optional] [default to null]
**CardBrand** | [***V1TenderCardBrand**](V1TenderCardBrand.md) |  | [optional] [default to null]
**PanSuffix** | **string** | The last four digits of the provided credit card&#x27;s account number. | [optional] [default to null]
**EntryMethod** | [***V1TenderEntryMethod**](V1TenderEntryMethod.md) |  | [optional] [default to null]
**PaymentNote** | **string** | Notes entered by the merchant about the tender at the time of payment, if any. Typically only present for tender with the type: OTHER. | [optional] [default to null]
**TotalMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**TenderedMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**TenderedAt** | **string** | The time when the tender was created, in ISO 8601 format. | [optional] [default to null]
**SettledAt** | **string** | The time when the tender was settled, in ISO 8601 format. | [optional] [default to null]
**ChangeBackMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**IsExchange** | **bool** | Indicates whether or not the tender is associated with an exchange. If is_exchange is true, the tender represents the value of goods returned in an exchange not the actual money paid. The exchange value reduces the tender amounts needed to pay for items purchased in the exchange. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

