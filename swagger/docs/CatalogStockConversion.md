# CatalogStockConversion

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StockableItemVariationId** | **string** | References to the stockable [CatalogItemVariation](entity:CatalogItemVariation) for this stock conversion. Selling, receiving or recounting the non-stockable &#x60;CatalogItemVariation&#x60; defined with a stock conversion results in adjustments of this stockable &#x60;CatalogItemVariation&#x60;. This immutable field must reference a stockable &#x60;CatalogItemVariation&#x60; that shares the parent [CatalogItem](entity:CatalogItem) of the converted &#x60;CatalogItemVariation.&#x60; | [default to null]
**StockableQuantity** | **string** | The quantity of the stockable item variation (as identified by &#x60;stockable_item_variation_id&#x60;) equivalent to the non-stockable item variation quantity (as specified in &#x60;nonstockable_quantity&#x60;) as defined by this stock conversion.  It accepts a decimal number in a string format that can take up to 10 digits before the decimal point and up to 5 digits after the decimal point. | [default to null]
**NonstockableQuantity** | **string** | The converted equivalent quantity of the non-stockable [CatalogItemVariation](entity:CatalogItemVariation) in its measurement unit. The &#x60;stockable_quantity&#x60; value and this &#x60;nonstockable_quantity&#x60; value together define the conversion ratio between stockable item variation and the non-stockable item variation. It accepts a decimal number in a string format that can take up to 10 digits before the decimal point and up to 5 digits after the decimal point. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

