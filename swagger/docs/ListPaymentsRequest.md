# ListPaymentsRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeginTime** | **string** | Timestamp for the beginning of the reporting period, in RFC 3339 format. Inclusive. Default: The current time minus one year. | [optional] [default to null]
**EndTime** | **string** | Timestamp for the end of the requested reporting period, in RFC 3339 format.  Default: The current time. | [optional] [default to null]
**SortOrder** | **string** | The order in which results are listed. - &#x60;ASC&#x60; - oldest to newest - &#x60;DESC&#x60; - newest to oldest (default). | [optional] [default to null]
**Cursor** | **string** | A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  See [Pagination](https://developer.squareup.com/docs/basics/api101/pagination) for more information. | [optional] [default to null]
**LocationId** | **string** | Limit results to the location supplied. By default, results are returned for the default (main) location associated with the merchant. | [optional] [default to null]
**Total** | **int64** | The exact amount in the total_money for a &#x60;Payment&#x60;. | [optional] [default to null]
**Last4** | **string** | The last 4 digits of &#x60;Payment&#x60; card. | [optional] [default to null]
**CardBrand** | **string** | The brand of &#x60;Payment&#x60; card. For example, &#x60;VISA&#x60; | [optional] [default to null]
**Limit** | **int32** | Maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page.  If the supplied value is greater than 100, at most 100 results will be returned.  Default: &#x60;100&#x60; | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

