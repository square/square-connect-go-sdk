# V1PageCell

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageId** | **string** | The unique identifier of the page the cell is included on. | [optional] [default to null]
**Row** | **int32** | The row of the cell. Always an integer between 0 and 4, inclusive. | [optional] [default to null]
**Column** | **int32** | The column of the cell. Always an integer between 0 and 4, inclusive. | [optional] [default to null]
**ObjectType** | [***V1PageCellObjectType**](V1PageCellObjectType.md) |  | [optional] [default to null]
**ObjectId** | **string** | The unique identifier of the entity represented in the cell. Not present for cells with an object_type of PLACEHOLDER. | [optional] [default to null]
**PlaceholderType** | [***V1PageCellPlaceholderType**](V1PageCellPlaceholderType.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

