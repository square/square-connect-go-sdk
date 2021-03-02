# CatalogProductSet

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | User-defined name for the product set. For example, \&quot;Clearance Items\&quot; or \&quot;Winter Sale Items\&quot;. | [optional] [default to null]
**ProductIdsAny** | **[]string** |  Unique IDs for any &#x60;CatalogObject&#x60; included in this product set. Any number of these catalog objects can be in an order for a pricing rule to apply.  This can be used with &#x60;product_ids_all&#x60; in a parent &#x60;CatalogProductSet&#x60; to match groups of products for a bulk discount, such as a discount for an entree and side combo.  Only one of &#x60;product_ids_all&#x60;, &#x60;product_ids_any&#x60;, or &#x60;all_products&#x60; can be set.  Max: 500 catalog object IDs. | [optional] [default to null]
**ProductIdsAll** | **[]string** | Unique IDs for any &#x60;CatalogObject&#x60; included in this product set. All objects in this set must be included in an order for a pricing rule to apply.  Only one of &#x60;product_ids_all&#x60;, &#x60;product_ids_any&#x60;, or &#x60;all_products&#x60; can be set.  Max: 500 catalog object IDs. | [optional] [default to null]
**QuantityExact** | **int64** | If set, there must be exactly this many items from &#x60;products_any&#x60; or &#x60;products_all&#x60; in the cart for the discount to apply.  Cannot be combined with either &#x60;quantity_min&#x60; or &#x60;quantity_max&#x60;. | [optional] [default to null]
**QuantityMin** | **int64** | If set, there must be at least this many items from &#x60;products_any&#x60; or &#x60;products_all&#x60; in a cart for the discount to apply. See &#x60;quantity_exact&#x60;. Defaults to 0 if &#x60;quantity_exact&#x60;, &#x60;quantity_min&#x60; and &#x60;quantity_max&#x60; are all unspecified. | [optional] [default to null]
**QuantityMax** | **int64** | If set, the pricing rule will apply to a maximum of this many items from &#x60;products_any&#x60; or &#x60;products_all&#x60;. | [optional] [default to null]
**AllProducts** | **bool** | If set to &#x60;true&#x60;, the product set will include every item in the catalog. Only one of &#x60;product_ids_all&#x60;, &#x60;product_ids_any&#x60;, or &#x60;all_products&#x60; can be set. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

