# ListPaymentsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeginTime** | **string** | The timestamp for the beginning of the reporting period, in RFC 3339 format. Inclusive. Default: The current time minus one year. | [optional] [default to null]
**EndTime** | **string** | The timestamp for the end of the reporting period, in RFC 3339 format.  Default: The current time. | [optional] [default to null]
**SortOrder** | **string** | The order in which results are listed: - &#x60;ASC&#x60; - Oldest to newest. - &#x60;DESC&#x60; - Newest to oldest (default). | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | [optional] [default to null]
**LocationId** | **string** | Limit results to the location supplied. By default, results are returned for the default (main) location associated with the seller. | [optional] [default to null]
**Total** | **int64** | The exact amount in the &#x60;total_money&#x60; for a payment. | [optional] [default to null]
**Last4** | **string** | The last four digits of a payment card. | [optional] [default to null]
**CardBrand** | **string** | The brand of the payment card (for example, VISA). | [optional] [default to null]
**Limit** | **int32** | The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page.  The default value of 100 is also the maximum allowed value. If the provided value is  greater than 100, it is ignored and the default value is used instead.  Default: &#x60;100&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

