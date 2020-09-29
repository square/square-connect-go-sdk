# V1Timecard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The timecard&#x27;s unique ID. | [optional] [default to null]
**EmployeeId** | **string** | The ID of the employee the timecard is associated with. | [default to null]
**Deleted** | **bool** | If true, the timecard was deleted by the merchant, and it is no longer valid. | [optional] [default to null]
**ClockinTime** | **string** | The clock-in time for the timecard, in ISO 8601 format. | [optional] [default to null]
**ClockoutTime** | **string** | The clock-out time for the timecard, in ISO 8601 format. Provide this value only if importing timecard information from another system. | [optional] [default to null]
**ClockinLocationId** | **string** | The ID of the location the employee clocked in from. We strongly reccomend providing a clockin_location_id. Square uses the clockin_location_id to determine a timecard’s timezone and overtime rules. | [optional] [default to null]
**ClockoutLocationId** | **string** | The ID of the location the employee clocked out from. Provide this value only if importing timecard information from another system. | [optional] [default to null]
**CreatedAt** | **string** | The time when the timecard was created, in ISO 8601 format. | [optional] [default to null]
**UpdatedAt** | **string** | The time when the timecard was most recently updated, in ISO 8601 format. | [optional] [default to null]
**RegularSecondsWorked** | **float64** | The total number of regular (non-overtime) seconds worked in the timecard. | [optional] [default to null]
**OvertimeSecondsWorked** | **float64** | The total number of overtime seconds worked in the timecard. | [optional] [default to null]
**DoubletimeSecondsWorked** | **float64** | The total number of doubletime seconds worked in the timecard. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

