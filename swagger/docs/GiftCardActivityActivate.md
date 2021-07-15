# GiftCardActivityActivate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**OrderId** | **string** | The ID of the order associated with the activity.  This is required if your application uses the Square Orders API. | [optional] [default to null]
**LineItemUid** | **string** | The &#x60;line_item_uid&#x60; of the gift card line item in an order.  This is required if your application uses the Square Orders API. | [optional] [default to null]
**ReferenceId** | **string** | If your application does not use the Square Orders API, you can optionally use this field  to associate the gift card activity with a client-side entity. | [optional] [default to null]
**BuyerPaymentInstrumentIds** | **[]string** | Required if your application does not use the Square Orders API.  This is a list of client-provided payment instrument IDs.  Square uses this information to perform compliance checks. If you use the Square Orders API, Square has the necessary instrument IDs to perform necessary  compliance checks. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

