# V1Variation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The item variation&#x27;s unique ID. | [optional] [default to null]
**Name** | **string** | The item variation&#x27;s name. | [optional] [default to null]
**ItemId** | **string** | The ID of the variation&#x27;s associated item. | [optional] [default to null]
**Ordinal** | **int32** | Indicates the variation&#x27;s list position when displayed in Square Point of Sale and the merchant dashboard. If more than one variation for the same item has the same ordinal value, those variations are displayed in alphabetical order | [optional] [default to null]
**PricingType** | [***V1VariationPricingType**](V1VariationPricingType.md) |  | [optional] [default to null]
**PriceMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**Sku** | **string** | The item variation&#x27;s SKU, if any. | [optional] [default to null]
**TrackInventory** | **bool** | If true, inventory tracking is active for the variation. | [optional] [default to null]
**InventoryAlertType** | [***V1VariationInventoryAlertType**](V1VariationInventoryAlertType.md) |  | [optional] [default to null]
**InventoryAlertThreshold** | **int32** | If the inventory quantity for the variation is less than or equal to this value and inventory_alert_type is LOW_QUANTITY, the variation displays an alert in the merchant dashboard. | [optional] [default to null]
**UserData** | **string** | Arbitrary metadata associated with the variation. Cannot exceed 255 characters. | [optional] [default to null]
**V2Id** | **string** | The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

