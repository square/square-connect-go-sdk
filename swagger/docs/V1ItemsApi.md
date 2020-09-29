# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdjustInventory**](V1ItemsApi.md#AdjustInventory) | **Post** /v1/{location_id}/inventory/{variation_id} | AdjustInventory
[**ApplyFee**](V1ItemsApi.md#ApplyFee) | **Put** /v1/{location_id}/items/{item_id}/fees/{fee_id} | ApplyFee
[**ApplyModifierList**](V1ItemsApi.md#ApplyModifierList) | **Put** /v1/{location_id}/items/{item_id}/modifier-lists/{modifier_list_id} | ApplyModifierList
[**CreateCategory**](V1ItemsApi.md#CreateCategory) | **Post** /v1/{location_id}/categories | CreateCategory
[**CreateDiscount**](V1ItemsApi.md#CreateDiscount) | **Post** /v1/{location_id}/discounts | CreateDiscount
[**CreateFee**](V1ItemsApi.md#CreateFee) | **Post** /v1/{location_id}/fees | CreateFee
[**CreateItem**](V1ItemsApi.md#CreateItem) | **Post** /v1/{location_id}/items | CreateItem
[**CreateModifierList**](V1ItemsApi.md#CreateModifierList) | **Post** /v1/{location_id}/modifier-lists | CreateModifierList
[**CreateModifierOption**](V1ItemsApi.md#CreateModifierOption) | **Post** /v1/{location_id}/modifier-lists/{modifier_list_id}/modifier-options | CreateModifierOption
[**CreatePage**](V1ItemsApi.md#CreatePage) | **Post** /v1/{location_id}/pages | CreatePage
[**CreateVariation**](V1ItemsApi.md#CreateVariation) | **Post** /v1/{location_id}/items/{item_id}/variations | CreateVariation
[**DeleteCategory**](V1ItemsApi.md#DeleteCategory) | **Delete** /v1/{location_id}/categories/{category_id} | DeleteCategory
[**DeleteDiscount**](V1ItemsApi.md#DeleteDiscount) | **Delete** /v1/{location_id}/discounts/{discount_id} | DeleteDiscount
[**DeleteFee**](V1ItemsApi.md#DeleteFee) | **Delete** /v1/{location_id}/fees/{fee_id} | DeleteFee
[**DeleteItem**](V1ItemsApi.md#DeleteItem) | **Delete** /v1/{location_id}/items/{item_id} | DeleteItem
[**DeleteModifierList**](V1ItemsApi.md#DeleteModifierList) | **Delete** /v1/{location_id}/modifier-lists/{modifier_list_id} | DeleteModifierList
[**DeleteModifierOption**](V1ItemsApi.md#DeleteModifierOption) | **Delete** /v1/{location_id}/modifier-lists/{modifier_list_id}/modifier-options/{modifier_option_id} | DeleteModifierOption
[**DeletePage**](V1ItemsApi.md#DeletePage) | **Delete** /v1/{location_id}/pages/{page_id} | DeletePage
[**DeletePageCell**](V1ItemsApi.md#DeletePageCell) | **Delete** /v1/{location_id}/pages/{page_id}/cells | DeletePageCell
[**DeleteVariation**](V1ItemsApi.md#DeleteVariation) | **Delete** /v1/{location_id}/items/{item_id}/variations/{variation_id} | DeleteVariation
[**ListCategories**](V1ItemsApi.md#ListCategories) | **Get** /v1/{location_id}/categories | ListCategories
[**ListDiscounts**](V1ItemsApi.md#ListDiscounts) | **Get** /v1/{location_id}/discounts | ListDiscounts
[**ListFees**](V1ItemsApi.md#ListFees) | **Get** /v1/{location_id}/fees | ListFees
[**ListInventory**](V1ItemsApi.md#ListInventory) | **Get** /v1/{location_id}/inventory | ListInventory
[**ListItems**](V1ItemsApi.md#ListItems) | **Get** /v1/{location_id}/items | ListItems
[**ListModifierLists**](V1ItemsApi.md#ListModifierLists) | **Get** /v1/{location_id}/modifier-lists | ListModifierLists
[**ListPages**](V1ItemsApi.md#ListPages) | **Get** /v1/{location_id}/pages | ListPages
[**RemoveFee**](V1ItemsApi.md#RemoveFee) | **Delete** /v1/{location_id}/items/{item_id}/fees/{fee_id} | RemoveFee
[**RemoveModifierList**](V1ItemsApi.md#RemoveModifierList) | **Delete** /v1/{location_id}/items/{item_id}/modifier-lists/{modifier_list_id} | RemoveModifierList
[**RetrieveItem**](V1ItemsApi.md#RetrieveItem) | **Get** /v1/{location_id}/items/{item_id} | RetrieveItem
[**RetrieveModifierList**](V1ItemsApi.md#RetrieveModifierList) | **Get** /v1/{location_id}/modifier-lists/{modifier_list_id} | RetrieveModifierList
[**UpdateCategory**](V1ItemsApi.md#UpdateCategory) | **Put** /v1/{location_id}/categories/{category_id} | UpdateCategory
[**UpdateDiscount**](V1ItemsApi.md#UpdateDiscount) | **Put** /v1/{location_id}/discounts/{discount_id} | UpdateDiscount
[**UpdateFee**](V1ItemsApi.md#UpdateFee) | **Put** /v1/{location_id}/fees/{fee_id} | UpdateFee
[**UpdateItem**](V1ItemsApi.md#UpdateItem) | **Put** /v1/{location_id}/items/{item_id} | UpdateItem
[**UpdateModifierList**](V1ItemsApi.md#UpdateModifierList) | **Put** /v1/{location_id}/modifier-lists/{modifier_list_id} | UpdateModifierList
[**UpdateModifierOption**](V1ItemsApi.md#UpdateModifierOption) | **Put** /v1/{location_id}/modifier-lists/{modifier_list_id}/modifier-options/{modifier_option_id} | UpdateModifierOption
[**UpdatePage**](V1ItemsApi.md#UpdatePage) | **Put** /v1/{location_id}/pages/{page_id} | UpdatePage
[**UpdatePageCell**](V1ItemsApi.md#UpdatePageCell) | **Put** /v1/{location_id}/pages/{page_id}/cells | UpdatePageCell
[**UpdateVariation**](V1ItemsApi.md#UpdateVariation) | **Put** /v1/{location_id}/items/{item_id}/variations/{variation_id} | UpdateVariation
[**UploadItemImage**](V1ItemsApi.md#UploadItemImage) | **Post** /v1/{location_id}/items/{item_id}/image | UploadItemImage

# **AdjustInventory**
> V1InventoryEntry AdjustInventory(ctx, body, locationId, variationId)
AdjustInventory

Adjusts the current available inventory of an item variation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1AdjustInventoryRequest**](V1AdjustInventoryRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **variationId** | **string**| The ID of the variation to adjust inventory information for. | 

### Return type

[**V1InventoryEntry**](V1InventoryEntry.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplyFee**
> V1Item ApplyFee(ctx, locationId, itemId, feeId)
ApplyFee

Associates a fee with an item so the fee is automatically applied to the item in Square Point of Sale.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the fee&#x27;s associated location. | 
  **itemId** | **string**| The ID of the item to add the fee to. | 
  **feeId** | **string**| The ID of the fee to apply. | 

### Return type

[**V1Item**](V1Item.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApplyModifierList**
> V1Item ApplyModifierList(ctx, locationId, modifierListId, itemId)
ApplyModifierList

Associates a modifier list with an item so the associated modifier options can be applied to the item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **modifierListId** | **string**| The ID of the modifier list to apply. | 
  **itemId** | **string**| The ID of the item to add the modifier list to. | 

### Return type

[**V1Item**](V1Item.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCategory**
> V1Category CreateCategory(ctx, body, locationId)
CreateCategory

Creates an item category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Category**](V1Category.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the location to create an item for. | 

### Return type

[**V1Category**](V1Category.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDiscount**
> V1Discount CreateDiscount(ctx, body, locationId)
CreateDiscount

Creates a discount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Discount**](V1Discount.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the location to create an item for. | 

### Return type

[**V1Discount**](V1Discount.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateFee**
> V1Fee CreateFee(ctx, body, locationId)
CreateFee

Creates a fee (tax).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Fee**](V1Fee.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the location to create a fee for. | 

### Return type

[**V1Fee**](V1Fee.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateItem**
> V1Item CreateItem(ctx, body, locationId)
CreateItem

Creates an item and at least one variation for it.    Item-related entities include fields you can use to associate them with entities in a non-Square system.  When you create an item-related entity, you can optionally specify `id`. This value must be unique among all IDs ever specified for the account, including those specified by other applications. You can never reuse an entity ID. If you do not specify an ID, Square generates one for the entity.  Item variations have a `user_data` string that lets you associate arbitrary metadata with the variation. The string cannot exceed 255 characters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Item**](V1Item.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the location to create an item for. | 

### Return type

[**V1Item**](V1Item.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateModifierList**
> V1ModifierList CreateModifierList(ctx, body, locationId)
CreateModifierList

Creates an item modifier list and at least 1 modifier option for it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1ModifierList**](V1ModifierList.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the location to create a modifier list for. | 

### Return type

[**V1ModifierList**](V1ModifierList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateModifierOption**
> V1ModifierOption CreateModifierOption(ctx, body, locationId, modifierListId)
CreateModifierOption

Creates an item modifier option and adds it to a modifier list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1ModifierOption**](V1ModifierOption.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **modifierListId** | **string**| The ID of the modifier list to edit. | 

### Return type

[**V1ModifierOption**](V1ModifierOption.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreatePage**
> V1Page CreatePage(ctx, body, locationId)
CreatePage

Creates a Favorites page in Square Point of Sale.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Page**](V1Page.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the location to create an item for. | 

### Return type

[**V1Page**](V1Page.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateVariation**
> V1Variation CreateVariation(ctx, body, locationId, itemId)
CreateVariation

Creates an item variation for an existing item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Variation**](V1Variation.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **itemId** | **string**| The item&#x27;s ID. | 

### Return type

[**V1Variation**](V1Variation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCategory**
> V1Category DeleteCategory(ctx, locationId, categoryId)
DeleteCategory

Deletes an existing item category.   __DeleteCategory__ returns nothing on success but Connect SDKs map the empty response to an empty `V1DeleteCategoryRequest` object as documented below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **categoryId** | **string**| The ID of the category to delete. | 

### Return type

[**V1Category**](V1Category.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDiscount**
> V1Discount DeleteDiscount(ctx, locationId, discountId)
DeleteDiscount

Deletes an existing discount.   __DeleteDiscount__ returns nothing on success but Connect SDKs map the empty response to an empty `V1DeleteDiscountRequest` object as documented below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **discountId** | **string**| The ID of the discount to delete. | 

### Return type

[**V1Discount**](V1Discount.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteFee**
> V1Fee DeleteFee(ctx, locationId, feeId)
DeleteFee

Deletes an existing fee (tax).   __DeleteFee__ returns nothing on success but Connect SDKs map the empty response to an empty `V1DeleteFeeRequest` object as documented below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the fee&#x27;s associated location. | 
  **feeId** | **string**| The ID of the fee to delete. | 

### Return type

[**V1Fee**](V1Fee.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteItem**
> V1Item DeleteItem(ctx, locationId, itemId)
DeleteItem

Deletes an existing item and all item variations associated with it.   __DeleteItem__ returns nothing on success but Connect SDKs map the empty response to an empty `V1DeleteItemRequest` object as documented below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **itemId** | **string**| The ID of the item to modify. | 

### Return type

[**V1Item**](V1Item.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteModifierList**
> V1ModifierList DeleteModifierList(ctx, locationId, modifierListId)
DeleteModifierList

Deletes an existing item modifier list and all modifier options associated with it.   __DeleteModifierList__ returns nothing on success but Connect SDKs map the empty response to an empty `V1DeleteModifierListRequest` object as documented below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **modifierListId** | **string**| The ID of the modifier list to delete. | 

### Return type

[**V1ModifierList**](V1ModifierList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteModifierOption**
> V1ModifierOption DeleteModifierOption(ctx, locationId, modifierListId, modifierOptionId)
DeleteModifierOption

Deletes an existing item modifier option from a modifier list.   __DeleteModifierOption__ returns nothing on success but Connect SDKs map the empty response to an empty `V1DeleteModifierOptionRequest` object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **modifierListId** | **string**| The ID of the modifier list to delete. | 
  **modifierOptionId** | **string**| The ID of the modifier list to edit. | 

### Return type

[**V1ModifierOption**](V1ModifierOption.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePage**
> V1Page DeletePage(ctx, locationId, pageId)
DeletePage

Deletes an existing Favorites page and all of its cells.   __DeletePage__ returns nothing on success but Connect SDKs map the empty response to an empty `V1DeletePageRequest` object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the Favorites page&#x27;s associated location. | 
  **pageId** | **string**| The ID of the page to delete. | 

### Return type

[**V1Page**](V1Page.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeletePageCell**
> V1Page DeletePageCell(ctx, locationId, pageId, optional)
DeletePageCell

Deletes a cell from a Favorites page in Square Point of Sale.   __DeletePageCell__ returns nothing on success but Connect SDKs map the empty response to an empty `V1DeletePageCellRequest` object as documented below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the Favorites page&#x27;s associated location. | 
  **pageId** | **string**| The ID of the page to delete. | 
 **optional** | ***V1ItemsApiDeletePageCellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ItemsApiDeletePageCellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **row** | **optional.String**| The row of the cell to clear. Always an integer between 0 and 4, inclusive. Row 0 is the top row. | 
 **column** | **optional.String**| The column of the cell to clear. Always an integer between 0 and 4, inclusive. Column 0 is the leftmost column. | 

### Return type

[**V1Page**](V1Page.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteVariation**
> V1Variation DeleteVariation(ctx, locationId, itemId, variationId)
DeleteVariation

Deletes an existing item variation from an item.   __DeleteVariation__ returns nothing on success but Connect SDKs map the empty response to an empty `V1DeleteVariationRequest` object as documented below.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **itemId** | **string**| The ID of the item to delete. | 
  **variationId** | **string**| The ID of the variation to delete. | 

### Return type

[**V1Variation**](V1Variation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCategories**
> []V1Category ListCategories(ctx, locationId)
ListCategories

Lists all the item categories for a given location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list categories for. | 

### Return type

[**[]V1Category**](V1Category.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDiscounts**
> []V1Discount ListDiscounts(ctx, locationId)
ListDiscounts

Lists all the discounts for a given location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list categories for. | 

### Return type

[**[]V1Discount**](V1Discount.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListFees**
> []V1Fee ListFees(ctx, locationId)
ListFees

Lists all the fees (taxes) for a given location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list fees for. | 

### Return type

[**[]V1Fee**](V1Fee.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListInventory**
> []V1InventoryEntry ListInventory(ctx, locationId, optional)
ListInventory

Provides inventory information for all inventory-enabled item variations.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
 **optional** | ***V1ItemsApiListInventoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ItemsApiListInventoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| The maximum number of inventory entries to return in a single response. This value cannot exceed 1000. | 
 **batchToken** | **optional.String**| A pagination cursor to retrieve the next set of results for your original query to the endpoint. | 

### Return type

[**[]V1InventoryEntry**](V1InventoryEntry.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListItems**
> []V1Item ListItems(ctx, locationId, optional)
ListItems

Provides summary information of all items for a given location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list items for. | 
 **optional** | ***V1ItemsApiListItemsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1ItemsApiListItemsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **batchToken** | **optional.String**| A pagination cursor to retrieve the next set of results for your original query to the endpoint. | 

### Return type

[**[]V1Item**](V1Item.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListModifierLists**
> []V1ModifierList ListModifierLists(ctx, locationId)
ListModifierLists

Lists all the modifier lists for a given location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list modifier lists for. | 

### Return type

[**[]V1ModifierList**](V1ModifierList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPages**
> []V1Page ListPages(ctx, locationId)
ListPages

Lists all Favorites pages (in Square Point of Sale) for a given location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list Favorites pages for. | 

### Return type

[**[]V1Page**](V1Page.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveFee**
> V1Item RemoveFee(ctx, locationId, itemId, feeId)
RemoveFee

Removes a fee assocation from an item so the fee is no longer automatically applied to the item in Square Point of Sale.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the fee&#x27;s associated location. | 
  **itemId** | **string**| The ID of the item to add the fee to. | 
  **feeId** | **string**| The ID of the fee to apply. | 

### Return type

[**V1Item**](V1Item.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveModifierList**
> V1Item RemoveModifierList(ctx, locationId, modifierListId, itemId)
RemoveModifierList

Removes a modifier list association from an item so the modifier options from the list can no longer be applied to the item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **modifierListId** | **string**| The ID of the modifier list to remove. | 
  **itemId** | **string**| The ID of the item to remove the modifier list from. | 

### Return type

[**V1Item**](V1Item.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveItem**
> V1Item RetrieveItem(ctx, locationId, itemId)
RetrieveItem

Provides the details for a single item, including associated modifier lists and fees.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **itemId** | **string**| The item&#x27;s ID. | 

### Return type

[**V1Item**](V1Item.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveModifierList**
> V1ModifierList RetrieveModifierList(ctx, locationId, modifierListId)
RetrieveModifierList

Provides the details for a single modifier list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **modifierListId** | **string**| The modifier list&#x27;s ID. | 

### Return type

[**V1ModifierList**](V1ModifierList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCategory**
> V1Category UpdateCategory(ctx, body, locationId, categoryId)
UpdateCategory

Modifies the details of an existing item category.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Category**](V1Category.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the category&#x27;s associated location. | 
  **categoryId** | **string**| The ID of the category to edit. | 

### Return type

[**V1Category**](V1Category.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDiscount**
> V1Discount UpdateDiscount(ctx, body, locationId, discountId)
UpdateDiscount

Modifies the details of an existing discount.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Discount**](V1Discount.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the category&#x27;s associated location. | 
  **discountId** | **string**| The ID of the discount to edit. | 

### Return type

[**V1Discount**](V1Discount.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFee**
> V1Fee UpdateFee(ctx, body, locationId, feeId)
UpdateFee

Modifies the details of an existing fee (tax).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Fee**](V1Fee.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the fee&#x27;s associated location. | 
  **feeId** | **string**| The ID of the fee to edit. | 

### Return type

[**V1Fee**](V1Fee.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateItem**
> V1Item UpdateItem(ctx, body, locationId, itemId)
UpdateItem

Modifies the core details of an existing item.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Item**](V1Item.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **itemId** | **string**| The ID of the item to modify. | 

### Return type

[**V1Item**](V1Item.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateModifierList**
> V1ModifierList UpdateModifierList(ctx, body, locationId, modifierListId)
UpdateModifierList

Modifies the details of an existing item modifier list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1UpdateModifierListRequest**](V1UpdateModifierListRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **modifierListId** | **string**| The ID of the modifier list to edit. | 

### Return type

[**V1ModifierList**](V1ModifierList.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateModifierOption**
> V1ModifierOption UpdateModifierOption(ctx, body, locationId, modifierListId, modifierOptionId)
UpdateModifierOption

Modifies the details of an existing item modifier option.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1ModifierOption**](V1ModifierOption.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **modifierListId** | **string**| The ID of the modifier list to edit. | 
  **modifierOptionId** | **string**| The ID of the modifier list to edit. | 

### Return type

[**V1ModifierOption**](V1ModifierOption.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePage**
> V1Page UpdatePage(ctx, body, locationId, pageId)
UpdatePage

Modifies the details of a Favorites page in Square Point of Sale.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Page**](V1Page.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the Favorites page&#x27;s associated location | 
  **pageId** | **string**| The ID of the page to modify. | 

### Return type

[**V1Page**](V1Page.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePageCell**
> V1Page UpdatePageCell(ctx, body, locationId, pageId)
UpdatePageCell

Modifies a cell of a Favorites page in Square Point of Sale.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1PageCell**](V1PageCell.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the Favorites page&#x27;s associated location. | 
  **pageId** | **string**| The ID of the page the cell belongs to. | 

### Return type

[**V1Page**](V1Page.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateVariation**
> V1Variation UpdateVariation(ctx, body, locationId, itemId, variationId)
UpdateVariation

Modifies the details of an existing item variation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1Variation**](V1Variation.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the item&#x27;s associated location. | 
  **itemId** | **string**| The ID of the item to modify. | 
  **variationId** | **string**| The ID of the variation to modify. | 

### Return type

[**V1Variation**](V1Variation.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UploadItemImage**
> V1ItemImage UploadItemImage(ctx, body, locationId, itemId)
UploadItemImage

Uploads a JPEG or PNG image and sets it as the master image for an item.   If you upload an image for an item that already has a master image, the new image replaces the existing one.  <aside class=\"important\"> Requests to this endpoint use the <code>Content-Type: multipart/form-data</code> header instead of <code>Content-Type: application/json</code>. Square recommends using an HTTP library (such as Requests for Python) that simplifies the process for sending <code>multipart/form-data</code> requests. </aside>  The request body assumes the multipart boundary for the request is set to `BOUNDARY` in the `Content-Type header`, for example: ``` Content-Type: multipart/form-data; boundary=BOUNDARY ``` __Note__: some HTTP libraries will set the multipart boundary for you.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1UploadItemImageRequest**](V1UploadItemImageRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The Square-issued location ID indicating where the item is available. | 
  **itemId** | **string**| The Square-issued ID of the item to modify. | 

### Return type

[**V1ItemImage**](V1ItemImage.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

