# BulkCreateTeamMembersResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamMembers** | [**map[string]CreateTeamMemberResponse**](CreateTeamMemberResponse.md) | The successfully created &#x60;TeamMember&#x60; objects. Each key is the &#x60;idempotency_key&#x60; that maps to the &#x60;CreateTeamMemberRequest&#x60;. | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) | The errors that occurred during the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

