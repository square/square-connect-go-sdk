# DeviceCheckoutOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | **string** | The unique ID of the device intended for this &#x60;TerminalCheckout&#x60;. A list of &#x60;DeviceCode&#x60; objects can be retrieved from the /v2/devices/codes endpoint. Match a &#x60;DeviceCode.device_id&#x60; value with &#x60;device_id&#x60; to get the associated device code. | [default to null]
**SkipReceiptScreen** | **bool** | Instruct the device to skip the receipt screen. Defaults to false. | [optional] [default to null]
**TipSettings** | [***TipSettings**](TipSettings.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

