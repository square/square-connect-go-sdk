# Customer

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique Square-assigned ID for the customer profile. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the customer profile was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the customer profile was last updated, in RFC 3339 format. | [optional] [default to null]
**Cards** | [**[]Card**](Card.md) | Payment details of cards stored on file for the customer profile. | [optional] [default to null]
**GivenName** | **string** | The given (i.e., first) name associated with the customer profile. | [optional] [default to null]
**FamilyName** | **string** | The family (i.e., last) name associated with the customer profile. | [optional] [default to null]
**Nickname** | **string** | A nickname for the customer profile. | [optional] [default to null]
**CompanyName** | **string** | A business name associated with the customer profile. | [optional] [default to null]
**EmailAddress** | **string** | The email address associated with the customer profile. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | The 11-digit phone number associated with the customer profile. | [optional] [default to null]
**Birthday** | **string** | The birthday associated with the customer profile, in RFC 3339 format. Year is optional, timezone and times are not allowed. For example: &#x60;0000-09-01T00:00:00-00:00&#x60; indicates a birthday on September 1st. &#x60;1998-09-01T00:00:00-00:00&#x60; indications a birthday on September 1st __1998__. | [optional] [default to null]
**ReferenceId** | **string** | An optional, second ID used to associate the customer profile with an entity in another system. | [optional] [default to null]
**Note** | **string** | A custom note associated with the customer profile. | [optional] [default to null]
**Preferences** | [***CustomerPreferences**](CustomerPreferences.md) |  | [optional] [default to null]
**Groups** | [**[]CustomerGroupInfo**](CustomerGroupInfo.md) | The customer groups and segments the customer belongs to. This deprecated field has been replaced with  the dedicated &#x60;group_ids&#x60; for customer groups and the dedicated &#x60;segment_ids&#x60; field for customer segments. You can retrieve information about a given customer group and segment respectively using the Customer Groups API and Customer Segments API. | [optional] [default to null]
**CreationSource** | [***CustomerCreationSource**](CustomerCreationSource.md) |  | [optional] [default to null]
**GroupIds** | **[]string** | The IDs of customer groups the customer belongs to. | [optional] [default to null]
**SegmentIds** | **[]string** | The IDs of segments the customer belongs to. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

