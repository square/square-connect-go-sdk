# CreateCustomerRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** | The idempotency key for the request. See the [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency) guide for more information. | [optional] [default to null]
**GivenName** | **string** | The given (i.e., first) name associated with the customer profile. | [optional] [default to null]
**FamilyName** | **string** | The family (i.e., last) name associated with the customer profile. | [optional] [default to null]
**CompanyName** | **string** | A business name associated with the customer profile. | [optional] [default to null]
**Nickname** | **string** | A nickname for the customer profile. | [optional] [default to null]
**EmailAddress** | **string** | The email address associated with the customer profile. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]
**PhoneNumber** | **string** | The 11-digit phone number associated with the customer profile. | [optional] [default to null]
**ReferenceId** | **string** | An optional, second ID used to associate the customer profile with an entity in another system. | [optional] [default to null]
**Note** | **string** | A custom note associated with the customer profile. | [optional] [default to null]
**Birthday** | **string** | The birthday associated with the customer profile, in RFC 3339 format. Year is optional, timezone and times are not allowed. For example: &#x60;0000-09-01T00:00:00-00:00&#x60; indicates a birthday on September 1st. &#x60;1998-09-01T00:00:00-00:00&#x60; indications a birthday on September 1st __1998__. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

