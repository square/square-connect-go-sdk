# CatalogItemVariation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** | The ID of the &#x60;CatalogItem&#x60; associated with this item variation. | [optional] [default to null]
**Name** | **string** | The item variation&#x27;s name. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points. | [optional] [default to null]
**Sku** | **string** | The item variation&#x27;s SKU, if any. This is a searchable attribute for use in applicable query filters. | [optional] [default to null]
**Upc** | **string** | The universal product code (UPC) of the item variation, if any. This is a searchable attribute for use in applicable query filters.  The value of this attribute should be a number of 12-14 digits long.  This restriction is enforced on the Square Seller Dashboard, Square Point of Sale or Retail Point of Sale apps, where this attribute shows in the GTIN field. If a non-compliant UPC value is assigned to this attribute using the API, the value is not editable on the Seller Dashboard, Square Point of Sale or Retail Point of Sale apps unless it is updated to fit the expected format. | [optional] [default to null]
**Ordinal** | **int32** | The order in which this item variation should be displayed. This value is read-only. On writes, the ordinal for each item variation within a parent &#x60;CatalogItem&#x60; is set according to the item variations&#x27;s position. On reads, the value is not guaranteed to be sequential or unique. | [optional] [default to null]
**PricingType** | [***CatalogPricingType**](CatalogPricingType.md) |  | [optional] [default to null]
**PriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**LocationOverrides** | [**[]ItemVariationLocationOverrides**](ItemVariationLocationOverrides.md) | Per-location price and inventory overrides. | [optional] [default to null]
**TrackInventory** | **bool** | If &#x60;true&#x60;, inventory tracking is active for the variation. | [optional] [default to null]
**InventoryAlertType** | [***InventoryAlertType**](InventoryAlertType.md) |  | [optional] [default to null]
**InventoryAlertThreshold** | **int64** | If the inventory quantity for the variation is less than or equal to this value and &#x60;inventory_alert_type&#x60; is &#x60;LOW_QUANTITY&#x60;, the variation displays an alert in the merchant dashboard.  This value is always an integer. | [optional] [default to null]
**UserData** | **string** | Arbitrary user metadata to associate with the item variation. This attribute value length is of Unicode code points. | [optional] [default to null]
**ServiceDuration** | **int64** | If the &#x60;CatalogItem&#x60; that owns this item variation is of type &#x60;APPOINTMENTS_SERVICE&#x60;, then this is the duration of the service in milliseconds. For example, a 30 minute appointment would have the value &#x60;1800000&#x60;, which is equal to 30 (minutes) * 60 (seconds per minute) * 1000 (milliseconds per second). | [optional] [default to null]
**AvailableForBooking** | **bool** | If the &#x60;CatalogItem&#x60; that owns this item variation is of type &#x60;APPOINTMENTS_SERVICE&#x60;, a bool representing whether this service is available for booking. | [optional] [default to null]
**ItemOptionValues** | [**[]CatalogItemOptionValueForItemVariation**](CatalogItemOptionValueForItemVariation.md) | List of item option values associated with this item variation. Listed in the same order as the item options of the parent item. | [optional] [default to null]
**MeasurementUnitId** | **string** | ID of the ‘CatalogMeasurementUnit’ that is used to measure the quantity sold of this item variation. If left unset, the item will be sold in whole quantities. | [optional] [default to null]
**Sellable** | **bool** | Whether this variation can be sold. The inventory count of a sellable variation indicates  the number of units available for sale. When a variation is both stockable and sellable,  its sellable inventory count can be smaller than or equal to its stockable count. | [optional] [default to null]
**Stockable** | **bool** | Whether stock is counted directly on this variation (TRUE) or only on its components (FALSE). When a variation is both stockable and sellable, the inventory count of a stockable variation keeps track of the number of units of this variation in stock and is not an indicator of the number of units of the variation that can be sold. | [optional] [default to null]
**ImageIds** | **[]string** | The IDs of images associated with this &#x60;CatalogItemVariation&#x60; instance. These images will be shown to customers in Square Online Store. | [optional] [default to null]
**TeamMemberIds** | **[]string** | Tokens of employees that can perform the service represented by this variation. Only valid for variations of type &#x60;APPOINTMENTS_SERVICE&#x60;. | [optional] [default to null]
**StockableConversion** | [***CatalogStockConversion**](CatalogStockConversion.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

