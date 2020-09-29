# V1CashDrawerShift

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The shift&#x27;s unique ID. | [optional] [default to null]
**EventType** | [***V1CashDrawerShiftEventType**](V1CashDrawerShiftEventType.md) |  | [optional] [default to null]
**OpenedAt** | **string** | The time when the shift began, in ISO 8601 format. | [optional] [default to null]
**EndedAt** | **string** | The time when the shift ended, in ISO 8601 format. | [optional] [default to null]
**ClosedAt** | **string** | The time when the shift was closed, in ISO 8601 format. | [optional] [default to null]
**EmployeeIds** | **[]string** | The IDs of all employees that were logged into Square Register at some point during the cash drawer shift. | [optional] [default to null]
**OpeningEmployeeId** | **string** | The ID of the employee that started the cash drawer shift. | [optional] [default to null]
**EndingEmployeeId** | **string** | The ID of the employee that ended the cash drawer shift. | [optional] [default to null]
**ClosingEmployeeId** | **string** | The ID of the employee that closed the cash drawer shift by auditing the cash drawer&#x27;s contents. | [optional] [default to null]
**Description** | **string** | A description of the cash drawer shift. | [optional] [default to null]
**StartingCashMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**CashPaymentMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**CashRefundsMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**CashPaidInMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**CashPaidOutMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**ExpectedCashMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**ClosedCashMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**Device** | [***Device**](Device.md) |  | [optional] [default to null]
**Events** | [**[]V1CashDrawerEvent**](V1CashDrawerEvent.md) | All of the events (payments, refunds, and so on) that involved the cash drawer during the shift. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

