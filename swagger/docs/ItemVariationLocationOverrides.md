# ItemVariationLocationOverrides

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | **string** | The ID of the &#x60;Location&#x60;. This can include locations that are deactivated. | [optional] [default to null]
**PriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**PricingType** | [***CatalogPricingType**](CatalogPricingType.md) |  | [optional] [default to null]
**TrackInventory** | **bool** | If &#x60;true&#x60;, inventory tracking is active for the &#x60;CatalogItemVariation&#x60; at this &#x60;Location&#x60;. | [optional] [default to null]
**InventoryAlertType** | [***InventoryAlertType**](InventoryAlertType.md) |  | [optional] [default to null]
**InventoryAlertThreshold** | **int64** | If the inventory quantity for the variation is less than or equal to this value and &#x60;inventory_alert_type&#x60; is &#x60;LOW_QUANTITY&#x60;, the variation displays an alert in the merchant dashboard.  This value is always an integer. | [optional] [default to null]
**SoldOut** | **bool** | Indicates whether the overridden item variation is sold out at the specified location.  When inventory tracking is enabled on the item variation either globally or at the specified location, the item variation is automatically marked as sold out when its inventory count reaches zero. The seller can manually set the item variation as sold out even when the inventory count is greater than zero. Attempts by an application to set this attribute are ignored. Regardless how the sold-out status is set, applications should treat its inventory count as zero when this attribute value is &#x60;true&#x60;. | [optional] [default to null]
**SoldOutValidUntil** | **string** | The seller-assigned timestamp, of the RFC 3339 format, to indicate when this sold-out variation becomes available again at the specified location. Attempts by an application to set this attribute are ignored. When the current time is later than this attribute value, the affected item variation is no longer sold out. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

