# CreateSubscriptionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** | A unique string that identifies this &#x60;CreateSubscription&#x60; request. If you do not provide a unique string (or provide an empty string as the value), the endpoint treats each request as independent.  For more information, see [Idempotency keys](https://developer.squareup.com/docs/working-with-apis/idempotency). | [optional] [default to null]
**LocationId** | **string** | The ID of the location the subscription is associated with. | [default to null]
**PlanId** | **string** | The ID of the subscription plan created using the Catalog API. For more information, see [Set Up and Manage a Subscription Plan](https://developer.squareup.com/docs/subscriptions-api/setup-plan) and  [Subscriptions Walkthrough](https://developer.squareup.com/docs/subscriptions-api/walkthrough). | [default to null]
**CustomerId** | **string** | The ID of the [customer](entity:Customer) subscribing to the subscription plan. | [default to null]
**StartDate** | **string** | The &#x60;YYYY-MM-DD&#x60;-formatted date to start the subscription.  If it is unspecified, the subscription starts immediately. | [optional] [default to null]
**CanceledDate** | **string** | The &#x60;YYYY-MM-DD&#x60;-formatted date when the newly created subscription is scheduled for cancellation.   This date overrides the cancellation date set in the plan configuration. If the cancellation date is earlier than the end date of a subscription cycle, the subscription stops at the canceled date and the subscriber is sent a prorated invoice at the beginning of the canceled cycle.   When the subscription plan of the newly created subscription has a fixed number of cycles and the &#x60;canceled_date&#x60; occurs before the subscription plan expires, the specified &#x60;canceled_date&#x60; sets the date when the subscription  stops through the end of the last cycle. | [optional] [default to null]
**TaxPercentage** | **string** | The tax to add when billing the subscription. The percentage is expressed in decimal form, using a &#x60;&#x27;.&#x27;&#x60; as the decimal separator and without a &#x60;&#x27;%&#x27;&#x60; sign. For example, a value of 7.5 corresponds to 7.5%. | [optional] [default to null]
**PriceOverrideMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CardId** | **string** | The ID of the [subscriber&#x27;s](entity:Customer) [card](entity:Card) to charge. If it is not specified, the subscriber receives an invoice via email. For an example to create a customer profile for a subscriber and add a card on file, see [Subscriptions Walkthrough](https://developer.squareup.com/docs/subscriptions-api/walkthrough). | [optional] [default to null]
**Timezone** | **string** | The timezone that is used in date calculations for the subscription. If unset, defaults to the location timezone. If a timezone is not configured for the location, defaults to \&quot;America/New_York\&quot;. Format: the IANA Timezone Database identifier for the location timezone. For a list of time zones, see [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones). | [optional] [default to null]
**Source** | [***SubscriptionSource**](SubscriptionSource.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

