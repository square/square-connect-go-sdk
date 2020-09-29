# OrderFulfillmentShipmentDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipient** | [***OrderFulfillmentRecipient**](OrderFulfillmentRecipient.md) |  | [optional] [default to null]
**Carrier** | **string** | The shipping carrier being used to ship this fulfillment e.g. UPS, FedEx, USPS, etc. | [optional] [default to null]
**ShippingNote** | **string** | A note with additional information for the shipping carrier. | [optional] [default to null]
**ShippingType** | **string** | A description of the type of shipping product purchased from the carrier. e.g. First Class, Priority, Express | [optional] [default to null]
**TrackingNumber** | **string** | The reference number provided by the carrier to track the shipment&#x27;s progress. | [optional] [default to null]
**TrackingUrl** | **string** | A link to the tracking webpage on the carrier&#x27;s website. | [optional] [default to null]
**PlacedAt** | **string** | The [timestamp](#workingwithdates) indicating when the shipment was requested. Must be in RFC 3339 timestamp format, e.g., \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**InProgressAt** | **string** | The [timestamp](#workingwithdates) indicating when this fulfillment was moved to the &#x60;RESERVED&#x60; state. Indicates that preparation of this shipment has begun. Must be in RFC 3339 timestamp format, e.g., \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**PackagedAt** | **string** | The [timestamp](#workingwithdates) indicating when this fulfillment was moved to the &#x60;PREPARED&#x60; state. Indicates that the fulfillment is packaged. Must be in RFC 3339 timestamp format, e.g., \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**ExpectedShippedAt** | **string** | The [timestamp](#workingwithdates) indicating when the shipment is expected to be delivered to the shipping carrier. Must be in RFC 3339 timestamp format, e.g., \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**ShippedAt** | **string** | The [timestamp](#workingwithdates) indicating when this fulfillment was moved to the &#x60;COMPLETED&#x60;state. Indicates that the fulfillment has been given to the shipping carrier. Must be in RFC 3339 timestamp format, e.g., \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**CanceledAt** | **string** | The [timestamp](#workingwithdates) indicating the shipment was canceled. Must be in RFC 3339 timestamp format, e.g., \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**CancelReason** | **string** | A description of why the shipment was canceled. | [optional] [default to null]
**FailedAt** | **string** | The [timestamp](#workingwithdates) indicating when the shipment failed to be completed. Must be in RFC 3339 timestamp format, e.g., \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**FailureReason** | **string** | A description of why the shipment failed to be completed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

