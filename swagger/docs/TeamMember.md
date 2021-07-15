# TeamMember

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID for the team member. | [optional] [default to null]
**ReferenceId** | **string** | A second ID used to associate the team member with an entity in another system. | [optional] [default to null]
**IsOwner** | **bool** | Whether the team member is the owner of the Square account. | [optional] [default to null]
**Status** | [***TeamMemberStatus**](TeamMemberStatus.md) |  | [optional] [default to null]
**GivenName** | **string** | The given name (that is, the first name) associated with the team member. | [optional] [default to null]
**FamilyName** | **string** | The family name (that is, the last name) associated with the team member. | [optional] [default to null]
**EmailAddress** | **string** | The email address associated with the team member. | [optional] [default to null]
**PhoneNumber** | **string** | The team member&#x27;s phone number, in E.164 format. For example: +14155552671 - the country code is 1 for US +551155256325 - the country code is 55 for BR | [optional] [default to null]
**CreatedAt** | **string** | The timestamp, in RFC 3339 format, describing when the team member was created. For example, \&quot;2018-10-04T04:00:00-07:00\&quot; or \&quot;2019-02-05T12:00:00Z\&quot;. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp, in RFC 3339 format, describing when the team member was last updated. For example, \&quot;2018-10-04T04:00:00-07:00\&quot; or \&quot;2019-02-05T12:00:00Z\&quot;. | [optional] [default to null]
**AssignedLocations** | [***TeamMemberAssignedLocations**](TeamMemberAssignedLocations.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

