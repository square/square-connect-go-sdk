# CatalogItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The item&#x27;s name. This is a searchable attribute for use in applicable query filters, its value must not be empty, and the length is of Unicode code points. | [optional] [default to null]
**Description** | **string** | The item&#x27;s description. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points. | [optional] [default to null]
**Abbreviation** | **string** | The text of the item&#x27;s display label in the Square Point of Sale app. Only up to the first five characters of the string are used. This attribute is searchable, and its value length is of Unicode code points. | [optional] [default to null]
**LabelColor** | **string** | The color of the item&#x27;s display label in the Square Point of Sale app. This must be a valid hex color code. | [optional] [default to null]
**AvailableOnline** | **bool** | If &#x60;true&#x60;, the item can be added to shipping orders from the merchant&#x27;s online store. | [optional] [default to null]
**AvailableForPickup** | **bool** | If &#x60;true&#x60;, the item can be added to pickup orders from the merchant&#x27;s online store. | [optional] [default to null]
**AvailableElectronically** | **bool** | If &#x60;true&#x60;, the item can be added to electronically fulfilled orders from the merchant&#x27;s online store. | [optional] [default to null]
**CategoryId** | **string** | The ID of the item&#x27;s category, if any. | [optional] [default to null]
**TaxIds** | **[]string** | A set of IDs indicating the taxes enabled for this item. When updating an item, any taxes listed here will be added to the item. Taxes may also be added to or deleted from an item using &#x60;UpdateItemTaxes&#x60;. | [optional] [default to null]
**ModifierListInfo** | [**[]CatalogItemModifierListInfo**](CatalogItemModifierListInfo.md) | A set of &#x60;CatalogItemModifierListInfo&#x60; objects representing the modifier lists that apply to this item, along with the overrides and min and max limits that are specific to this item. Modifier lists may also be added to or deleted from an item using &#x60;UpdateItemModifierLists&#x60;. | [optional] [default to null]
**Variations** | [**[]CatalogObject**](CatalogObject.md) | A list of CatalogObjects containing the &#x60;CatalogItemVariation&#x60;s for this item. | [optional] [default to null]
**ProductType** | [***CatalogItemProductType**](CatalogItemProductType.md) |  | [optional] [default to null]
**SkipModifierScreen** | **bool** | If &#x60;false&#x60;, the Square Point of Sale app will present the &#x60;CatalogItem&#x60;&#x27;s details screen immediately, allowing the merchant to choose &#x60;CatalogModifier&#x60;s before adding the item to the cart.  This is the default behavior.  If &#x60;true&#x60;, the Square Point of Sale app will immediately add the item to the cart with the pre-selected modifiers, and merchants can edit modifiers by drilling down onto the item&#x27;s details.  Third-party clients are encouraged to implement similar behaviors. | [optional] [default to null]
**ItemOptions** | [**[]CatalogItemOptionForItem**](CatalogItemOptionForItem.md) | List of item options IDs for this item. Used to manage and group item variations in a specified order.  Maximum: 6 item options. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

