# ShiftFilter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationIds** | **[]string** | Fetch shifts for the specified location. | [optional] [default to null]
**EmployeeIds** | **[]string** | Fetch shifts for the specified employees. DEPRECATED at version 2020-08-26. Use &#x60;team_member_ids&#x60; instead. | [optional] [default to null]
**Status** | [***ShiftFilterStatus**](ShiftFilterStatus.md) |  | [optional] [default to null]
**Start** | [***TimeRange**](TimeRange.md) |  | [optional] [default to null]
**End** | [***TimeRange**](TimeRange.md) |  | [optional] [default to null]
**Workday** | [***ShiftWorkday**](ShiftWorkday.md) |  | [optional] [default to null]
**TeamMemberIds** | **[]string** | Fetch shifts for the specified team members. Replaced &#x60;employee_ids&#x60; at version \&quot;2020-08-26\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

