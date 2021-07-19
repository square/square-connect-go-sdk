# TerminalRefundQueryFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** | &#x60;TerminalRefund&#x60; objects associated with a specific device. If no device is specified, then all &#x60;TerminalRefund&#x60; objects for the signed-in account are displayed. | [optional] [default to null]
**CreatedAt** | [***TimeRange**](TimeRange.md) |  | [optional] [default to null]
**Status** | **string** | Filtered results with the desired status of the &#x60;TerminalRefund&#x60;. Options: &#x60;PENDING&#x60;, &#x60;IN_PROGRESS&#x60;, &#x60;CANCEL_REQUESTED&#x60;, &#x60;CANCELED&#x60;, or &#x60;COMPLETED&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

