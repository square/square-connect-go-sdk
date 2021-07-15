# ListPaymentRefundsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeginTime** | **string** | The timestamp for the beginning of the requested reporting period, in RFC 3339 format.  Default: The current time minus one year. | [optional] [default to null]
**EndTime** | **string** | The timestamp for the end of the requested reporting period, in RFC 3339 format.  Default: The current time. | [optional] [default to null]
**SortOrder** | **string** | The order in which results are listed: - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default). | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | [optional] [default to null]
**LocationId** | **string** | Limit results to the location supplied. By default, results are returned for all locations associated with the seller. | [optional] [default to null]
**Status** | **string** | If provided, only refunds with the given status are returned. For a list of refund status values, see [PaymentRefund](entity:PaymentRefund).  Default: If omitted, refunds are returned regardless of their status. | [optional] [default to null]
**SourceType** | **string** | If provided, only refunds with the given source type are returned. - &#x60;CARD&#x60; - List refunds only for payments where &#x60;CARD&#x60; was specified as the payment source.  Default: If omitted, refunds are returned regardless of the source type. | [optional] [default to null]
**Limit** | **int32** | The maximum number of results to be returned in a single page.  It is possible to receive fewer results than the specified limit on a given page.  If the supplied value is greater than 100, no more than 100 results are returned.  Default: 100 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

