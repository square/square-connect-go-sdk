# Subscription

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the subscription. | [optional] [default to null]
**LocationId** | **string** | The ID of the location associated with the subscription. | [optional] [default to null]
**PlanId** | **string** | The ID of the associated [subscription plan](entity:CatalogSubscriptionPlan). | [optional] [default to null]
**CustomerId** | **string** | The ID of the associated [customer](entity:Customer) profile. | [optional] [default to null]
**StartDate** | **string** | The start date of the subscription, in YYYY-MM-DD format (for example, 2013-01-15). | [optional] [default to null]
**CanceledDate** | **string** | The subscription cancellation date, in YYYY-MM-DD format (for example, 2013-01-15). On this date, the subscription status changes to &#x60;CANCELED&#x60; and the subscription billing stops. If you don&#x27;t set this field, the subscription plan dictates if and when subscription ends.  You cannot update this field, you can only clear it. | [optional] [default to null]
**ChargedThroughDate** | **string** | The date up to which the customer is invoiced for the subscription, in YYYY-MM-DD format (for example, 2013-01-15).  After the invoice is sent for a given billing period, this date will be the last day of the billing period. For example, suppose for the month of May a customer gets an invoice (or charged the card) on May 1. For the monthly billing scenario, this date is then set to May 31. | [optional] [default to null]
**Status** | [***SubscriptionStatus**](SubscriptionStatus.md) |  | [optional] [default to null]
**TaxPercentage** | **string** | The tax amount applied when billing the subscription. The percentage is expressed in decimal form, using a &#x60;&#x27;.&#x27;&#x60; as the decimal separator and without a &#x60;&#x27;%&#x27;&#x60; sign. For example, a value of &#x60;7.5&#x60; corresponds to 7.5%. | [optional] [default to null]
**InvoiceIds** | **[]string** | The IDs of the [invoices](entity:Invoice) created for the subscription, listed in order when the invoices were created (oldest invoices appear first). | [optional] [default to null]
**PriceOverrideMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Version** | **int64** | The version of the object. When updating an object, the version supplied must match the version in the database, otherwise the write will be rejected as conflicting. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the subscription was created, in RFC 3339 format. | [optional] [default to null]
**CardId** | **string** | The ID of the [customer](entity:Customer) [card](entity:Card) that is charged for the subscription. | [optional] [default to null]
**PaidUntilDate** | **string** | The date up to which the customer is invoiced for the subscription, in YYYY-MM-DD format (for example, 2013-01-15).  After the invoice is sent for a given billing period, this date will be the last day of the billing period. For example, suppose for the month of May a customer gets an invoice (or charged the card) on May 1. For the monthly billing scenario, this date is then set to May 31. | [optional] [default to null]
**Timezone** | **string** | Timezone that will be used in date calculations for the subscription. Defaults to the timezone of the location based on &#x60;location_id&#x60;. Format: the IANA Timezone Database identifier for the location timezone (for example, &#x60;America/Los_Angeles&#x60;). | [optional] [default to null]
**Source** | [***SubscriptionSource**](SubscriptionSource.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

