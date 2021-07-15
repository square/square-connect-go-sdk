# OrderFulfillmentShipmentDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | [***OrderFulfillmentRecipient**](OrderFulfillmentRecipient.md) |  | [optional] [default to null]
**Carrier** | **string** | The shipping carrier being used to ship this fulfillment (such as UPS, FedEx, or USPS). | [optional] [default to null]
**ShippingNote** | **string** | A note with additional information for the shipping carrier. | [optional] [default to null]
**ShippingType** | **string** | A description of the type of shipping product purchased from the carrier (such as First Class, Priority, or Express). | [optional] [default to null]
**TrackingNumber** | **string** | The reference number provided by the carrier to track the shipment&#x27;s progress. | [optional] [default to null]
**TrackingUrl** | **string** | A link to the tracking webpage on the carrier&#x27;s website. | [optional] [default to null]
**PlacedAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the shipment was requested. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**InProgressAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when this fulfillment was moved to the &#x60;RESERVED&#x60; state, which  indicates that preparation of this shipment has begun. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**PackagedAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when this fulfillment was moved to the &#x60;PREPARED&#x60; state, which indicates that the fulfillment is packaged. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**ExpectedShippedAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the shipment is expected to be delivered to the shipping carrier. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**ShippedAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when this fulfillment was moved to the &#x60;COMPLETED&#x60; state, which indicates that the fulfillment has been given to the shipping carrier. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**CanceledAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating the shipment was canceled. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**CancelReason** | **string** | A description of why the shipment was canceled. | [optional] [default to null]
**FailedAt** | **string** | The [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) indicating when the shipment failed to be completed. The timestamp must be in RFC 3339 format (for example, \&quot;2016-09-04T23:59:33.123Z\&quot;). | [optional] [default to null]
**FailureReason** | **string** | A description of why the shipment failed to be completed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

