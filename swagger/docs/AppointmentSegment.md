# AppointmentSegment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationMinutes** | **int32** | The time span in minutes of an appointment segment. | [optional] [default to null]
**ServiceVariationId** | **string** | The ID of the [CatalogItemVariation](entity:CatalogItemVariation) object representing the service booked in this segment. | [optional] [default to null]
**TeamMemberId** | **string** | The ID of the [TeamMember](entity:TeamMember) object representing the team member booked in this segment. | [default to null]
**ServiceVariationVersion** | **int64** | The current version of the item variation representing the service booked in this segment. | [optional] [default to null]
**IntermissionMinutes** | **int32** | Time between the end of this segment and the beginning of the subsequent segment. | [optional] [default to null]
**AnyTeamMember** | **bool** | Whether the customer accepts any team member, instead of a specific one, to serve this segment. | [optional] [default to null]
**ResourceIds** | **[]string** | The IDs of the seller-accessible resources used for this appointment segment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

