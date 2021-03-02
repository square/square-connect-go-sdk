# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelBooking**](BookingsApi.md#CancelBooking) | **Post** /v2/bookings/{booking_id}/cancel | CancelBooking
[**CreateBooking**](BookingsApi.md#CreateBooking) | **Post** /v2/bookings | CreateBooking
[**ListTeamMemberBookingProfiles**](BookingsApi.md#ListTeamMemberBookingProfiles) | **Get** /v2/bookings/team-member-booking-profiles | ListTeamMemberBookingProfiles
[**RetrieveBooking**](BookingsApi.md#RetrieveBooking) | **Get** /v2/bookings/{booking_id} | RetrieveBooking
[**RetrieveBusinessBookingProfile**](BookingsApi.md#RetrieveBusinessBookingProfile) | **Get** /v2/bookings/business-booking-profile | RetrieveBusinessBookingProfile
[**RetrieveTeamMemberBookingProfile**](BookingsApi.md#RetrieveTeamMemberBookingProfile) | **Get** /v2/bookings/team-member-booking-profiles/{team_member_id} | RetrieveTeamMemberBookingProfile
[**SearchAvailability**](BookingsApi.md#SearchAvailability) | **Post** /v2/bookings/availability/search | SearchAvailability
[**UpdateBooking**](BookingsApi.md#UpdateBooking) | **Put** /v2/bookings/{booking_id} | UpdateBooking

# **CancelBooking**
> CancelBookingResponse CancelBooking(ctx, body, bookingId)
CancelBooking

Cancels an existing booking.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CancelBookingRequest**](CancelBookingRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **bookingId** | **string**| The ID of the [Booking](#type-booking) object representing the to-be-cancelled booking. | 

### Return type

[**CancelBookingResponse**](CancelBookingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBooking**
> CreateBookingResponse CreateBooking(ctx, body)
CreateBooking

Creates a booking.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBookingRequest**](CreateBookingRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**CreateBookingResponse**](CreateBookingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTeamMemberBookingProfiles**
> ListTeamMemberBookingProfilesResponse ListTeamMemberBookingProfiles(ctx, optional)
ListTeamMemberBookingProfiles

Lists booking profiles for team members.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BookingsApiListTeamMemberBookingProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BookingsApiListTeamMemberBookingProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bookableOnly** | **optional.Bool**| Indicates whether to include only bookable team members in the returned result (&#x60;true&#x60;) or not (&#x60;false&#x60;). | [default to false]
 **limit** | **optional.Int32**| The maximum number of results to return. | 
 **cursor** | **optional.String**| The cursor for paginating through the results. | 
 **locationId** | **optional.String**| Indicates whether to include only team members enabled at the given location in the returned result. | 

### Return type

[**ListTeamMemberBookingProfilesResponse**](ListTeamMemberBookingProfilesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveBooking**
> RetrieveBookingResponse RetrieveBooking(ctx, bookingId)
RetrieveBooking

Retrieves a booking.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bookingId** | **string**| The ID of the [Booking](#type-booking) object representing the to-be-retrieved booking. | 

### Return type

[**RetrieveBookingResponse**](RetrieveBookingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveBusinessBookingProfile**
> RetrieveBusinessBookingProfileResponse RetrieveBusinessBookingProfile(ctx, )
RetrieveBusinessBookingProfile

Retrieves a seller's booking profile.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RetrieveBusinessBookingProfileResponse**](RetrieveBusinessBookingProfileResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveTeamMemberBookingProfile**
> RetrieveTeamMemberBookingProfileResponse RetrieveTeamMemberBookingProfile(ctx, teamMemberId)
RetrieveTeamMemberBookingProfile

Retrieves a team member's booking profile.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **teamMemberId** | **string**| The ID of the team member to retrieve. | 

### Return type

[**RetrieveTeamMemberBookingProfileResponse**](RetrieveTeamMemberBookingProfileResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchAvailability**
> SearchAvailabilityResponse SearchAvailability(ctx, body)
SearchAvailability

Searches for availabilities for booking.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchAvailabilityRequest**](SearchAvailabilityRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchAvailabilityResponse**](SearchAvailabilityResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBooking**
> UpdateBookingResponse UpdateBooking(ctx, body, bookingId)
UpdateBooking

Updates a booking.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBookingRequest**](UpdateBookingRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **bookingId** | **string**| The ID of the [Booking](#type-booking) object representing the to-be-updated booking. | 

### Return type

[**UpdateBookingResponse**](UpdateBookingResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

