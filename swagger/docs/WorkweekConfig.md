# WorkweekConfig

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID for this object. | [optional] [default to null]
**StartOfWeek** | [***Weekday**](Weekday.md) |  | [default to null]
**StartOfDayLocalTime** | **string** | The local time at which a business week starts. Represented as a string in &#x60;HH:MM&#x60; format (&#x60;HH:MM:SS&#x60; is also accepted, but seconds are truncated). | [default to null]
**Version** | **int32** | Used for resolving concurrency issues. The request fails if the version provided does not match the server version at the time of the request. If not provided, Square executes a blind write; potentially overwriting data from another write. | [optional] [default to null]
**CreatedAt** | **string** | A read-only timestamp in RFC 3339 format; presented in UTC. | [optional] [default to null]
**UpdatedAt** | **string** | A read-only timestamp in RFC 3339 format; presented in UTC. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

