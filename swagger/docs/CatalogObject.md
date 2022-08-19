# CatalogObject

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | [***CatalogObjectType**](CatalogObjectType.md) |  | [default to null]
**Id** | **string** | An identifier to reference this object in the catalog. When a new &#x60;CatalogObject&#x60; is inserted, the client should set the id to a temporary identifier starting with a \&quot;&#x60;#&#x60;\&quot; character. Other objects being inserted or updated within the same request may use this identifier to refer to the new object.  When the server receives the new object, it will supply a unique identifier that replaces the temporary identifier for all future references. | [default to null]
**UpdatedAt** | **string** | Last modification [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) in RFC 3339 format, e.g., &#x60;\&quot;2016-08-15T23:59:33.123Z\&quot;&#x60; would indicate the UTC time (denoted by &#x60;Z&#x60;) of August 15, 2016 at 23:59:33 and 123 milliseconds. | [optional] [default to null]
**Version** | **int64** | The version of the object. When updating an object, the version supplied must match the version in the database, otherwise the write will be rejected as conflicting. | [optional] [default to null]
**IsDeleted** | **bool** | If &#x60;true&#x60;, the object has been deleted from the database. Must be &#x60;false&#x60; for new objects being inserted. When deleted, the &#x60;updated_at&#x60; field will equal the deletion time. | [optional] [default to null]
**CustomAttributeValues** | [**map[string]CatalogCustomAttributeValue**](CatalogCustomAttributeValue.md) | A map (key-value pairs) of application-defined custom attribute values. The value of a key-value pair is a [CatalogCustomAttributeValue](entity:CatalogCustomAttributeValue) object. The key is the &#x60;key&#x60; attribute value defined in the associated [CatalogCustomAttributeDefinition](entity:CatalogCustomAttributeDefinition) object defined by the application making the request.  If the &#x60;CatalogCustomAttributeDefinition&#x60; object is defined by another application, the &#x60;CatalogCustomAttributeDefinition&#x60;&#x27;s key attribute value is prefixed by the defining application ID. For example, if the &#x60;CatalogCustomAttributeDefinition&#x60; has a &#x60;key&#x60; attribute of &#x60;\&quot;cocoa_brand\&quot;&#x60; and the defining application ID is &#x60;\&quot;abcd1234\&quot;&#x60;, the key in the map is &#x60;\&quot;abcd1234:cocoa_brand\&quot;&#x60; if the application making the request is different from the application defining the custom attribute definition. Otherwise, the key used in the map is simply &#x60;\&quot;cocoa_brand\&quot;&#x60;.  Application-defined custom attributes are set at a global (location-independent) level. Custom attribute values are intended to store additional information about a catalog object or associations with an entity in another system. Do not use custom attributes to store any sensitive information (personally identifiable information, card details, etc.). | [optional] [default to null]
**CatalogV1Ids** | [**[]CatalogV1Id**](CatalogV1Id.md) | The Connect v1 IDs for this object at each location where it is present, where they differ from the object&#x27;s Connect V2 ID. The field will only be present for objects that have been created or modified by legacy APIs. | [optional] [default to null]
**PresentAtAllLocations** | **bool** | If &#x60;true&#x60;, this object is present at all locations (including future locations), except where specified in the &#x60;absent_at_location_ids&#x60; field. If &#x60;false&#x60;, this object is not present at any locations (including future locations), except where specified in the &#x60;present_at_location_ids&#x60; field. If not specified, defaults to &#x60;true&#x60;. | [optional] [default to null]
**PresentAtLocationIds** | **[]string** | A list of locations where the object is present, even if &#x60;present_at_all_locations&#x60; is &#x60;false&#x60;. This can include locations that are deactivated. | [optional] [default to null]
**AbsentAtLocationIds** | **[]string** | A list of locations where the object is not present, even if &#x60;present_at_all_locations&#x60; is &#x60;true&#x60;. This can include locations that are deactivated. | [optional] [default to null]
**ImageId** | **string** | Identifies the &#x60;CatalogImage&#x60; attached to this &#x60;CatalogObject&#x60;. | [optional] [default to null]
**ItemData** | [***CatalogItem**](CatalogItem.md) |  | [optional] [default to null]
**CategoryData** | [***CatalogCategory**](CatalogCategory.md) |  | [optional] [default to null]
**ItemVariationData** | [***CatalogItemVariation**](CatalogItemVariation.md) |  | [optional] [default to null]
**TaxData** | [***CatalogTax**](CatalogTax.md) |  | [optional] [default to null]
**DiscountData** | [***CatalogDiscount**](CatalogDiscount.md) |  | [optional] [default to null]
**ModifierListData** | [***CatalogModifierList**](CatalogModifierList.md) |  | [optional] [default to null]
**ModifierData** | [***CatalogModifier**](CatalogModifier.md) |  | [optional] [default to null]
**TimePeriodData** | [***CatalogTimePeriod**](CatalogTimePeriod.md) |  | [optional] [default to null]
**ProductSetData** | [***CatalogProductSet**](CatalogProductSet.md) |  | [optional] [default to null]
**PricingRuleData** | [***CatalogPricingRule**](CatalogPricingRule.md) |  | [optional] [default to null]
**ImageData** | [***CatalogImage**](CatalogImage.md) |  | [optional] [default to null]
**MeasurementUnitData** | [***CatalogMeasurementUnit**](CatalogMeasurementUnit.md) |  | [optional] [default to null]
**SubscriptionPlanData** | [***CatalogSubscriptionPlan**](CatalogSubscriptionPlan.md) |  | [optional] [default to null]
**ItemOptionData** | [***CatalogItemOption**](CatalogItemOption.md) |  | [optional] [default to null]
**ItemOptionValueData** | [***CatalogItemOptionValue**](CatalogItemOptionValue.md) |  | [optional] [default to null]
**CustomAttributeDefinitionData** | [***CatalogCustomAttributeDefinition**](CatalogCustomAttributeDefinition.md) |  | [optional] [default to null]
**QuickAmountsSettingsData** | [***CatalogQuickAmountsSettings**](CatalogQuickAmountsSettings.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

