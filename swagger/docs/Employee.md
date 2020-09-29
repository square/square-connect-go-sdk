# Employee

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | UUID for this object. | [optional] [default to null]
**FirstName** | **string** | The employee&#x27;s first name. | [optional] [default to null]
**LastName** | **string** | The employee&#x27;s last name. | [optional] [default to null]
**Email** | **string** | The employee&#x27;s email address | [optional] [default to null]
**PhoneNumber** | **string** | The employee&#x27;s phone number in E.164 format, i.e. \&quot;+12125554250\&quot; | [optional] [default to null]
**LocationIds** | **[]string** | A list of location IDs where this employee has access to. | [optional] [default to null]
**Status** | [***EmployeeStatus**](EmployeeStatus.md) |  | [optional] [default to null]
**IsOwner** | **bool** | Whether this employee is the owner of the merchant. Each merchant has one owner employee, and that employee has full authority over the account. | [optional] [default to null]
**CreatedAt** | **string** | A read-only timestamp in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | A read-only timestamp in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

