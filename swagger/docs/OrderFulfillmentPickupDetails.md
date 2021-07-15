# OrderFulfillmentPickupDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | [***OrderFulfillmentRecipient**](OrderFulfillmentRecipient.md) |  | [optional] [default to null]
**ExpiresAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates)  indicating when this fulfillment expires if it is not accepted. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). The expiration time can only be set up to 7 days in the future. If &#x60;expires_at&#x60; is not set, this pickup fulfillment is automatically accepted when  placed. | [optional] [default to null]
**AutoCompleteDuration** | **string** | The duration of time after which an open and accepted pickup fulfillment is automatically moved to the &#x60;COMPLETED&#x60; state. The duration must be in RFC 3339 format (for example, \&quot;P1W3D\&quot;).  If not set, this pickup fulfillment remains accepted until it is canceled or completed. | [optional] [default to null]
**ScheduleType** | [***OrderFulfillmentPickupDetailsScheduleType**](OrderFulfillmentPickupDetailsScheduleType.md) |  | [optional] [default to null]
**PickupAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates)  that represents the start of the pickup window. Must be in RFC 3339 timestamp format, e.g.,  \&quot;2016-09-04T23:59:33.123Z\&quot;.  For fulfillments with the schedule type &#x60;ASAP&#x60;, this is automatically set to the current time plus the expected duration to prepare the fulfillment. | [optional] [default to null]
**PickupWindowDuration** | **string** | The window of time in which the order should be picked up after the &#x60;pickup_at&#x60; timestamp. Must be in RFC 3339 duration format, e.g., \&quot;P1W3D\&quot;. Can be used as an informational guideline for merchants. | [optional] [default to null]
**PrepTimeDuration** | **string** | The duration of time it takes to prepare this fulfillment. The duration must be in RFC 3339 format (for example, \&quot;P1W3D\&quot;). | [optional] [default to null]
**Note** | **string** | A note meant to provide additional instructions about the pickup fulfillment displayed in the Square Point of Sale application and set by the API. | [optional] [default to null]
**PlacedAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates)  indicating when the fulfillment was placed. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**AcceptedAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates)  indicating when the fulfillment was accepted. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**RejectedAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates)  indicating when the fulfillment was rejected. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**ReadyAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates)  indicating when the fulfillment is marked as ready for pickup. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**ExpiredAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the fulfillment expired. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**PickedUpAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the fulfillment was picked up by the recipient. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**CanceledAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates)  indicating when the fulfillment was canceled. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**CancelReason** | **string** | A description of why the pickup was canceled. The maximum length: 100 characters. | [optional] [default to null]
**IsCurbsidePickup** | **bool** | If set to &#x60;true&#x60;, indicates that this pickup order is for curbside pickup, not in-store pickup. | [optional] [default to null]
**CurbsidePickupDetails** | [***OrderFulfillmentPickupDetailsCurbsidePickupDetails**](OrderFulfillmentPickupDetailsCurbsidePickupDetails.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

