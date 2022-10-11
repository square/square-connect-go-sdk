# TerminalAction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID for this &#x60;TerminalAction&#x60;. | [optional] [default to null]
**DeviceId** | **string** | The unique Id of the device intended for this &#x60;TerminalAction&#x60;. The Id can be retrieved from /v2/devices api. | [optional] [default to null]
**DeadlineDuration** | **string** | The duration as an RFC 3339 duration, after which the action will be automatically canceled. TerminalActions that are &#x60;PENDING&#x60; will be automatically &#x60;CANCELED&#x60; and have a cancellation reason of &#x60;TIMED_OUT&#x60;  Default: 5 minutes from creation  Maximum: 5 minutes | [optional] [default to null]
**Status** | **string** | The status of the &#x60;TerminalAction&#x60;. Options: &#x60;PENDING&#x60;, &#x60;IN_PROGRESS&#x60;, &#x60;CANCEL_REQUESTED&#x60;, &#x60;CANCELED&#x60;, &#x60;COMPLETED&#x60; | [optional] [default to null]
**CancelReason** | [***ActionCancelReason**](ActionCancelReason.md) |  | [optional] [default to null]
**CreatedAt** | **string** | The time when the &#x60;TerminalAction&#x60; was created as an RFC 3339 timestamp. | [optional] [default to null]
**UpdatedAt** | **string** | The time when the &#x60;TerminalAction&#x60; was last updated as an RFC 3339 timestamp. | [optional] [default to null]
**AppId** | **string** | The ID of the application that created the action. | [optional] [default to null]
**Type_** | [***TerminalActionActionType**](TerminalActionActionType.md) |  | [optional] [default to null]
**SaveCardOptions** | [***SaveCardOptions**](SaveCardOptions.md) |  | [optional] [default to null]
**DeviceMetadata** | [***DeviceMetadata**](DeviceMetadata.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

