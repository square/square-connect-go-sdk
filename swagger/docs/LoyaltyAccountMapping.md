# LoyaltyAccountMapping

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the mapping. | [optional] [default to null]
**Type_** | [***LoyaltyAccountMappingType**](LoyaltyAccountMappingType.md) |  | [optional] [default to null]
**Value** | **string** | The mapping value, which is used with &#x60;type&#x60; to represent a phone number mapping. The value can be a phone number in E.164 format. For example, \&quot;+14155551111\&quot;. RETIRED at version 2021-05-13. When specifying a mapping, use the &#x60;phone_number&#x60; field instead. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the mapping was created, in RFC 3339 format. | [optional] [default to null]
**PhoneNumber** | **string** | The phone number of the buyer, in E.164 format. For example, \&quot;+14155551111\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

