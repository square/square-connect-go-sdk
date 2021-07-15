# WageSetting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamMemberId** | **string** | The unique ID of the &#x60;TeamMember&#x60; whom this wage setting describes. | [optional] [default to null]
**JobAssignments** | [**[]JobAssignment**](JobAssignment.md) | Required. The ordered list of jobs that the team member is assigned to. The first job assignment is considered the team member&#x27;s primary job.  The minimum length is 1 and the maximum length is 12. | [optional] [default to null]
**IsOvertimeExempt** | **bool** | Whether the team member is exempt from the overtime rules of the seller&#x27;s country. | [optional] [default to null]
**Version** | **int32** | Used for resolving concurrency issues. The request fails if the version provided does not match the server version at the time of the request. If not provided, Square executes a blind write, potentially overwriting data from another write. For more information, see [optimistic concurrency](https://developer.squareup.com/docs/working-with-apis/optimistic-concurrency). | [optional] [default to null]
**CreatedAt** | **string** | The timestamp, in RFC 3339 format, describing when the wage setting object was created. For example, \&quot;2018-10-04T04:00:00-07:00\&quot; or \&quot;2019-02-05T12:00:00Z\&quot;. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp, in RFC 3339 format, describing when the wage setting object was last updated. For example, \&quot;2018-10-04T04:00:00-07:00\&quot; or \&quot;2019-02-05T12:00:00Z\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

