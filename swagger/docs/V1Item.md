# V1Item

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The item&#x27;s ID. Must be unique among all entity IDs ever provided on behalf of the merchant. You can never reuse an ID. This value can include alphanumeric characters, dashes (-), and underscores (_). | [optional] [default to null]
**Name** | **string** | The item&#x27;s name. | [optional] [default to null]
**Description** | **string** | The item&#x27;s description. | [optional] [default to null]
**Type_** | [***V1ItemType**](V1ItemType.md) |  | [optional] [default to null]
**Color** | [***V1ItemColor**](V1ItemColor.md) |  | [optional] [default to null]
**Abbreviation** | **string** | The text of the item&#x27;s display label in Square Point of Sale. Only up to the first five characters of the string are used. | [optional] [default to null]
**Visibility** | [***V1ItemVisibility**](V1ItemVisibility.md) |  | [optional] [default to null]
**AvailableOnline** | **bool** | If true, the item can be added to shipping orders from the merchant&#x27;s online store. | [optional] [default to null]
**MasterImage** | [***V1ItemImage**](V1ItemImage.md) |  | [optional] [default to null]
**Category** | [***V1Category**](V1Category.md) |  | [optional] [default to null]
**Variations** | [**[]V1Variation**](V1Variation.md) | The item&#x27;s variations. You must specify at least one variation. | [optional] [default to null]
**ModifierLists** | [**[]V1ModifierList**](V1ModifierList.md) | The modifier lists that apply to the item, if any. | [optional] [default to null]
**Fees** | [**[]V1Fee**](V1Fee.md) | The fees that apply to the item, if any. | [optional] [default to null]
**Taxable** | **bool** | Deprecated. This field is not used. | [optional] [default to null]
**CategoryId** | **string** | The ID of the item&#x27;s category, if any. | [optional] [default to null]
**AvailableForPickup** | **bool** | If true, the item can be added to pickup orders from the merchant&#x27;s online store. Default value: false | [optional] [default to null]
**V2Id** | **string** | The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

