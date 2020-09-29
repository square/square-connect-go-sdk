# V1ModifierOption

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The modifier option&#x27;s unique ID. | [optional] [default to null]
**Name** | **string** | The modifier option&#x27;s name. | [optional] [default to null]
**PriceMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**OnByDefault** | **bool** | If true, the modifier option is the default option in a modifier list for which selection_type is SINGLE. | [optional] [default to null]
**Ordinal** | **int32** | Indicates the modifier option&#x27;s list position when displayed in Square Point of Sale and the merchant dashboard. If more than one modifier option in the same modifier list has the same ordinal value, those options are displayed in alphabetical order. | [optional] [default to null]
**ModifierListId** | **string** | The ID of the modifier list the option belongs to. | [optional] [default to null]
**V2Id** | **string** | The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

