# CreateSubscriptionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** | A unique string that identifies this &#x60;CreateSubscription&#x60; request. If you do not provide a unique string (or provide an empty string as the value), the endpoint treats each request as independent.  For more information, see [Idempotency keys](https://developer.squareup.com/docs/working-with-apis/idempotency). | [default to null]
**LocationId** | **string** | The ID of the location the subscription is associated with. | [default to null]
**PlanId** | **string** | The ID of the subscription plan. For more information, see  [Subscription Plan Overview](https://developer.squareup.com/docs/subscriptions/overview). | [default to null]
**CustomerId** | **string** | The ID of the [customer](#type-customer) profile. | [default to null]
**StartDate** | **string** | The start date of the subscription, in YYYY-MM-DD format. For example, 2013-01-15. If the start date is left empty, the subscription begins  immediately. | [optional] [default to null]
**CanceledDate** | **string** | The date when the subscription should be canceled, in  YYYY-MM-DD format (for example, 2025-02-29). This overrides the plan configuration  if it comes before the date the subscription would otherwise end. | [optional] [default to null]
**TaxPercentage** | **string** | The tax to add when billing the subscription. The percentage is expressed in decimal form, using a &#x60;&#x27;.&#x27;&#x60; as the decimal separator and without a &#x60;&#x27;%&#x27;&#x60; sign. For example, a value of 7.5 corresponds to 7.5%. | [optional] [default to null]
**PriceOverrideMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CardId** | **string** | The ID of the [customer](#type-customer) [card](#type-card) to charge. If not specified, Square sends an invoice via email. For an example to create a customer and add a card on file, see [Subscriptions Walkthrough](https://developer.squareup.com/docs/subscriptions-api/walkthrough). | [optional] [default to null]
**Timezone** | **string** | The timezone that is used in date calculations for the subscription. If unset, defaults to the location timezone. If a timezone is not configured for the location, defaults to \&quot;America/New_York\&quot;. Format: the IANA Timezone Database identifier for the location timezone. For a list of time zones, see [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

