# WageSetting

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamMemberId** | **string** | The unique ID of the &#x60;TeamMember&#x60; whom this wage setting describes. | [optional] [default to null]
**JobAssignments** | [**[]JobAssignment**](JobAssignment.md) | &lt;b&gt;Required&lt;/b&gt; The ordered list of jobs that the team member is assigned to. The first job assignment is considered the team member&#x27;s \&quot;Primary Job\&quot;. &lt;br&gt; &lt;b&gt;Min Length 1    Max Length 12&lt;/b&gt; | [optional] [default to null]
**IsOvertimeExempt** | **bool** | Whether the team member is exempt from the overtime rules of the seller country. | [optional] [default to null]
**Version** | **int32** | Used for resolving concurrency issues; request will fail if version provided does not match server version at time of request. If not provided, Square executes a blind write, potentially overwriting data from another write. Read about [optimistic concurrency](https://developer.squareup.com/docs/docs/working-with-apis/optimistic-concurrency) in Square APIs for more information. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp in RFC 3339 format describing when the wage setting object was created. Ex: \&quot;2018-10-04T04:00:00-07:00\&quot; or \&quot;2019-02-05T12:00:00Z\&quot; | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp in RFC 3339 format describing when the wage setting object was last updated. Ex: \&quot;2018-10-04T04:00:00-07:00\&quot; or \&quot;2019-02-05T12:00:00Z\&quot; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

