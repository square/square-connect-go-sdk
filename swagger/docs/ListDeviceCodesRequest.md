# ListDeviceCodesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Paginating results](https://developer.squareup.com/docs/working-with-apis/pagination) for more information. | [optional] [default to null]
**LocationId** | **string** | If specified, only returns DeviceCodes of the specified location. Returns DeviceCodes of all locations if empty. | [optional] [default to null]
**ProductType** | [***ProductType**](ProductType.md) |  | [optional] [default to null]
**Status** | [**[]DeviceCodeStatus**](DeviceCodeStatus.md) | If specified, returns DeviceCodes with the specified statuses. Returns DeviceCodes of status &#x60;PAIRED&#x60; and &#x60;UNPAIRED&#x60; if empty. See [DeviceCodeStatus](#type-devicecodestatus) for possible values | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

