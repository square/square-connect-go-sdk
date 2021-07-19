# Customer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique Square-assigned ID for the customer profile. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the customer profile was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the customer profile was last updated, in RFC 3339 format. | [optional] [default to null]
**Cards** | [**[]Card**](Card.md) | Payment details of the credit, debit, and gift cards stored on file for the customer profile.   DEPRECATED at version 2021-06-16. Replaced by calling [ListCards](api-endpoint:Cards-ListCards) (for credit and debit cards on file)  or [ListGiftCards](api-endpoint:GiftCards-ListGiftCards) (for gift cards on file) and including the &#x60;customer_id&#x60; query parameter.  For more information, see [Migrate to the Cards API and Gift Cards API](https://developer.squareup.com/docs/customers-api/use-the-api/integrate-with-other-services#migrate-customer-cards). | [optional] [default to null]
**GivenName** | **string** | The given (i.e., first) name associated with the customer profile. | [optional] [default to null]
**FamilyName** | **string** | The family (i.e., last) name associated with the customer profile. | [optional] [default to null]
**Nickname** | **string** | A nickname for the customer profile. | [optional] [default to null]
**CompanyName** | **string** | A business name associated with the customer profile. | [optional] [default to null]
**EmailAddress** | **string** | The email address associated with the customer profile. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | The 11-digit phone number associated with the customer profile. | [optional] [default to null]
**Birthday** | **string** | The birthday associated with the customer profile, in RFC 3339 format. The year is optional. The timezone and time are not allowed. For example, &#x60;0000-09-21T00:00:00-00:00&#x60; represents a birthday on September 21 and &#x60;1998-09-21T00:00:00-00:00&#x60; represents a birthday on September 21, 1998. | [optional] [default to null]
**ReferenceId** | **string** | An optional second ID used to associate the customer profile with an entity in another system. | [optional] [default to null]
**Note** | **string** | A custom note associated with the customer profile. | [optional] [default to null]
**Preferences** | [***CustomerPreferences**](CustomerPreferences.md) |  | [optional] [default to null]
**Groups** | [**[]CustomerGroupInfo**](CustomerGroupInfo.md) | The customer groups and segments the customer belongs to. This deprecated field has been replaced with  the dedicated &#x60;group_ids&#x60; for customer groups and the dedicated &#x60;segment_ids&#x60; field for customer segments. You can retrieve information about a given customer group and segment respectively using the Customer Groups API and Customer Segments API. | [optional] [default to null]
**CreationSource** | [***CustomerCreationSource**](CustomerCreationSource.md) |  | [optional] [default to null]
**GroupIds** | **[]string** | The IDs of customer groups the customer belongs to. | [optional] [default to null]
**SegmentIds** | **[]string** | The IDs of segments the customer belongs to. | [optional] [default to null]
**Version** | **int64** | The Square-assigned version number of the customer profile. The version number is incremented each time an update is committed to the customer profile, except for changes to customer segment membership and cards on file. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

