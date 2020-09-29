# Shift

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | UUID for this object | [optional] [default to null]
**EmployeeId** | **string** | The ID of the employee this shift belongs to. DEPRECATED at version 2020-08-26. Use &#x60;team_member_id&#x60; instead | [optional] [default to null]
**LocationId** | **string** | The ID of the location this shift occurred at. Should be based on where the employee clocked in. | [optional] [default to null]
**Timezone** | **string** | Read-only convenience value that is calculated from the location based on &#x60;location_id&#x60;. Format: the IANA Timezone Database identifier for the location timezone. | [optional] [default to null]
**StartAt** | **string** | RFC 3339; shifted to location timezone + offset. Precision up to the minute is respected; seconds are truncated. | [default to null]
**EndAt** | **string** | RFC 3339; shifted to timezone + offset. Precision up to the minute is respected; seconds are truncated. | [optional] [default to null]
**Wage** | [***ShiftWage**](ShiftWage.md) |  | [optional] [default to null]
**Breaks** | [**[]ModelBreak**](Break.md) | A list of any paid or unpaid breaks that were taken during this shift. | [optional] [default to null]
**Status** | [***ShiftStatus**](ShiftStatus.md) |  | [optional] [default to null]
**Version** | **int32** | Used for resolving concurrency issues; request will fail if version provided does not match server version at time of request. If not provided, Square executes a blind write; potentially overwriting data from another write. | [optional] [default to null]
**CreatedAt** | **string** | A read-only timestamp in RFC 3339 format; presented in UTC. | [optional] [default to null]
**UpdatedAt** | **string** | A read-only timestamp in RFC 3339 format; presented in UTC. | [optional] [default to null]
**TeamMemberId** | **string** | The ID of the team member this shift belongs to. Replaced &#x60;employee_id&#x60; at version \&quot;2020-08-26\&quot; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

