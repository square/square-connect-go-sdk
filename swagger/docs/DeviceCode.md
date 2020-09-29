# DeviceCode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique id for this device code. | [optional] [default to null]
**Name** | **string** | An optional user-defined name for the device code. | [optional] [default to null]
**Code** | **string** | The unique code that can be used to login. | [optional] [default to null]
**DeviceId** | **string** | The unique id of the device that used this code. Populated when the device is paired up. | [optional] [default to null]
**ProductType** | [***ProductType**](ProductType.md) |  | [default to null]
**LocationId** | **string** | The location assigned to this code. | [optional] [default to null]
**Status** | [***DeviceCodeStatus**](DeviceCodeStatus.md) |  | [optional] [default to null]
**PairBy** | **string** | When this DeviceCode will expire and no longer login. Timestamp in RFC 3339 format. | [optional] [default to null]
**CreatedAt** | **string** | When this DeviceCode was created. Timestamp in RFC 3339 format. | [optional] [default to null]
**StatusChangedAt** | **string** | When this DeviceCode&#x27;s status was last changed. Timestamp in RFC 3339 format. | [optional] [default to null]
**PairedAt** | **string** | When this DeviceCode was paired. Timestamp in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

