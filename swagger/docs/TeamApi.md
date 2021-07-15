# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateTeamMembers**](TeamApi.md#BulkCreateTeamMembers) | **Post** /v2/team-members/bulk-create | BulkCreateTeamMembers
[**BulkUpdateTeamMembers**](TeamApi.md#BulkUpdateTeamMembers) | **Post** /v2/team-members/bulk-update | BulkUpdateTeamMembers
[**CreateTeamMember**](TeamApi.md#CreateTeamMember) | **Post** /v2/team-members | CreateTeamMember
[**RetrieveTeamMember**](TeamApi.md#RetrieveTeamMember) | **Get** /v2/team-members/{team_member_id} | RetrieveTeamMember
[**RetrieveWageSetting**](TeamApi.md#RetrieveWageSetting) | **Get** /v2/team-members/{team_member_id}/wage-setting | RetrieveWageSetting
[**SearchTeamMembers**](TeamApi.md#SearchTeamMembers) | **Post** /v2/team-members/search | SearchTeamMembers
[**UpdateTeamMember**](TeamApi.md#UpdateTeamMember) | **Put** /v2/team-members/{team_member_id} | UpdateTeamMember
[**UpdateWageSetting**](TeamApi.md#UpdateWageSetting) | **Put** /v2/team-members/{team_member_id}/wage-setting | UpdateWageSetting

# **BulkCreateTeamMembers**
> BulkCreateTeamMembersResponse BulkCreateTeamMembers(ctx, body)
BulkCreateTeamMembers

Creates multiple `TeamMember` objects. The created `TeamMember` objects are returned on successful creates. This process is non-transactional and processes as much of the request as possible. If one of the creates in the request cannot be successfully processed, the request is not marked as failed, but the body of the response contains explicit error information for the failed create.  Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#bulk-create-team-members).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkCreateTeamMembersRequest**](BulkCreateTeamMembersRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkCreateTeamMembersResponse**](BulkCreateTeamMembersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **BulkUpdateTeamMembers**
> BulkUpdateTeamMembersResponse BulkUpdateTeamMembers(ctx, body)
BulkUpdateTeamMembers

Updates multiple `TeamMember` objects. The updated `TeamMember` objects are returned on successful updates. This process is non-transactional and processes as much of the request as possible. If one of the updates in the request cannot be successfully processed, the request is not marked as failed, but the body of the response contains explicit error information for the failed update. Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#bulk-update-team-members).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkUpdateTeamMembersRequest**](BulkUpdateTeamMembersRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**BulkUpdateTeamMembersResponse**](BulkUpdateTeamMembersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTeamMember**
> CreateTeamMemberResponse CreateTeamMember(ctx, body)
CreateTeamMember

Creates a single `TeamMember` object. The `TeamMember` object is returned on successful creates. You must provide the following values in your request to this endpoint: - `given_name` - `family_name`  Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#createteammember).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateTeamMemberRequest**](CreateTeamMemberRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateTeamMemberResponse**](CreateTeamMemberResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveTeamMember**
> RetrieveTeamMemberResponse RetrieveTeamMember(ctx, teamMemberId)
RetrieveTeamMember

Retrieves a `TeamMember` object for the given `TeamMember.id`. Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#retrieve-a-team-member).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamMemberId** | **string**| The ID of the team member to retrieve. | 

### Return type

[**RetrieveTeamMemberResponse**](RetrieveTeamMemberResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveWageSetting**
> RetrieveWageSettingResponse RetrieveWageSetting(ctx, teamMemberId)
RetrieveWageSetting

Retrieves a `WageSetting` object for a team member specified by `TeamMember.id`. Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#retrievewagesetting).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamMemberId** | **string**| The ID of the team member for which to retrieve the wage setting. | 

### Return type

[**RetrieveWageSettingResponse**](RetrieveWageSettingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchTeamMembers**
> SearchTeamMembersResponse SearchTeamMembers(ctx, body)
SearchTeamMembers

Returns a paginated list of `TeamMember` objects for a business. The list can be filtered by the following: - location IDs - `status`

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchTeamMembersRequest**](SearchTeamMembersRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchTeamMembersResponse**](SearchTeamMembersResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTeamMember**
> UpdateTeamMemberResponse UpdateTeamMember(ctx, body, teamMemberId)
UpdateTeamMember

Updates a single `TeamMember` object. The `TeamMember` object is returned on successful updates. Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#update-a-team-member).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateTeamMemberRequest**](UpdateTeamMemberRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **teamMemberId** | **string**| The ID of the team member to update. | 

### Return type

[**UpdateTeamMemberResponse**](UpdateTeamMemberResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWageSetting**
> UpdateWageSettingResponse UpdateWageSetting(ctx, body, teamMemberId)
UpdateWageSetting

Creates or updates a `WageSetting` object. The object is created if a `WageSetting` with the specified `team_member_id` does not exist. Otherwise, it fully replaces the `WageSetting` object for the team member. The `WageSetting` is returned on a successful update. Learn about [Troubleshooting the Team API](https://developer.squareup.com/docs/team/troubleshooting#create-or-update-a-wage-setting).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateWageSettingRequest**](UpdateWageSettingRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **teamMemberId** | **string**| The ID of the team member for which to update the &#x60;WageSetting&#x60; object. | 

### Return type

[**UpdateWageSettingResponse**](UpdateWageSettingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

