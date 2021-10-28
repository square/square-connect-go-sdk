# Shift

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID for this object. | [optional] [default to null]
**EmployeeId** | **string** | The ID of the employee this shift belongs to. DEPRECATED at version 2020-08-26. Use &#x60;team_member_id&#x60; instead. | [optional] [default to null]
**LocationId** | **string** | The ID of the location this shift occurred at. The location should be based on where the employee clocked in. | [optional] [default to null]
**Timezone** | **string** | The read-only convenience value that is calculated from the location based on the &#x60;location_id&#x60;. Format: the IANA timezone database identifier for the location timezone. | [optional] [default to null]
**StartAt** | **string** | RFC 3339; shifted to the location timezone + offset. Precision up to the minute is respected; seconds are truncated. | [default to null]
**EndAt** | **string** | RFC 3339; shifted to the timezone + offset. Precision up to the minute is respected; seconds are truncated. | [optional] [default to null]
**Wage** | [***ShiftWage**](ShiftWage.md) |  | [optional] [default to null]
**Breaks** | [**[]ModelBreak**](Break.md) | A list of all the paid or unpaid breaks that were taken during this shift. | [optional] [default to null]
**Status** | [***ShiftStatus**](ShiftStatus.md) |  | [optional] [default to null]
**Version** | **int32** | Used for resolving concurrency issues. The request fails if the version provided does not match the server version at the time of the request. If not provided, Square executes a blind write; potentially overwriting data from another write. | [optional] [default to null]
**CreatedAt** | **string** | A read-only timestamp in RFC 3339 format; presented in UTC. | [optional] [default to null]
**UpdatedAt** | **string** | A read-only timestamp in RFC 3339 format; presented in UTC. | [optional] [default to null]
**TeamMemberId** | **string** | The ID of the team member this shift belongs to. Replaced &#x60;employee_id&#x60; at version \&quot;2020-08-26\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

