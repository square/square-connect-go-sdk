# UpdateCustomerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GivenName** | **string** | The given name (that is, the first name) associated with the customer profile. | [optional] [default to null]
**FamilyName** | **string** | The family name (that is, the last name) associated with the customer profile. | [optional] [default to null]
**CompanyName** | **string** | A business name associated with the customer profile. | [optional] [default to null]
**Nickname** | **string** | A nickname for the customer profile. | [optional] [default to null]
**EmailAddress** | **string** | The email address associated with the customer profile. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | The 11-digit phone number associated with the customer profile. | [optional] [default to null]
**ReferenceId** | **string** | An optional second ID used to associate the customer profile with an entity in another system. | [optional] [default to null]
**Note** | **string** | A custom note associated with the customer profile. | [optional] [default to null]
**Birthday** | **string** | The birthday associated with the customer profile, in RFC 3339 format. The year is optional. The timezone and time are not allowed. For example, &#x60;0000-09-21T00:00:00-00:00&#x60; represents a birthday on September 21 and &#x60;1998-09-21T00:00:00-00:00&#x60; represents a birthday on September 21, 1998.  You can also specify this value in &#x60;YYYY-MM-DD&#x60; format. | [optional] [default to null]
**Version** | **int64** | The current version of the customer profile.   As a best practice, you should include this field to enable [optimistic concurrency](https://developer.squareup.com/docs/working-with-apis/optimistic-concurrency) control. For more information, see [Update a customer profile](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#update-a-customer-profile). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

