# CashDrawerShift

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The shift unique ID. | [optional] [default to null]
**State** | [***CashDrawerShiftState**](CashDrawerShiftState.md) |  | [optional] [default to null]
**OpenedAt** | **string** | The time when the shift began, in ISO 8601 format. | [optional] [default to null]
**EndedAt** | **string** | The time when the shift ended, in ISO 8601 format. | [optional] [default to null]
**ClosedAt** | **string** | The time when the shift was closed, in ISO 8601 format. | [optional] [default to null]
**EmployeeIds** | **[]string** | The IDs of all employees that were logged into Square Point of Sale at any point while the cash drawer shift was open. | [optional] [default to null]
**OpeningEmployeeId** | **string** | The ID of the employee that started the cash drawer shift. | [optional] [default to null]
**EndingEmployeeId** | **string** | The ID of the employee that ended the cash drawer shift. | [optional] [default to null]
**ClosingEmployeeId** | **string** | The ID of the employee that closed the cash drawer shift by auditing the cash drawer contents. | [optional] [default to null]
**Description** | **string** | The free-form text description of a cash drawer by an employee. | [optional] [default to null]
**OpenedCashMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CashPaymentMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CashRefundsMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CashPaidInMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CashPaidOutMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ExpectedCashMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ClosedCashMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Device** | [***CashDrawerDevice**](CashDrawerDevice.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

