# ListCashDrawerShiftEventsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]CashDrawerShiftEvent**](CashDrawerShiftEvent.md) | All of the events (payments, refunds, etc.) for a cash drawer during the shift. | [optional] [default to null]
**Cursor** | **string** | Opaque cursor for fetching the next page. Cursor is not present in the last page of results. | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

