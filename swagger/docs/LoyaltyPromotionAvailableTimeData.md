# LoyaltyPromotionAvailableTimeData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StartDate** | **string** | The date that the promotion starts, in &#x60;YYYY-MM-DD&#x60; format. Square populates this field based on the provided &#x60;time_periods&#x60;. | [optional] [default to null]
**EndDate** | **string** | The date that the promotion ends, in &#x60;YYYY-MM-DD&#x60; format. Square populates this field based on the provided &#x60;time_periods&#x60;. If an end date is not specified, an &#x60;ACTIVE&#x60; promotion remains available until it is canceled. | [optional] [default to null]
**TimePeriods** | **[]string** | A list of [iCalendar (RFC 5545) events](https://tools.ietf.org/html/rfc5545#section-3.6.1) (&#x60;VEVENT&#x60;). Each event represents an available time period per day or days of the week.  A day can have a maximum of one available time period.  Only &#x60;DTSTART&#x60;, &#x60;DURATION&#x60;, and &#x60;RRULE&#x60; are supported. &#x60;DTSTART&#x60; and &#x60;DURATION&#x60; are required and timestamps must be in local (unzoned) time format. Include &#x60;RRULE&#x60; to specify recurring promotions, an end date (using the &#x60;UNTIL&#x60; keyword), or both. For more information, see [Available time](https://developer.squareup.com/docs/loyalty-api/loyalty-promotions#available-time).  Note that &#x60;BEGIN:VEVENT&#x60; and &#x60;END:VEVENT&#x60; are optional in a &#x60;CreateLoyaltyPromotion&#x60; request but are always included in the response. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

