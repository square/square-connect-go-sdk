# GiftCardActivityLoad

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**OrderId** | **string** | The &#x60;order_id&#x60; of the order associated with the activity. It is populated along with &#x60;line_item_uid&#x60; and is required if using the Square Orders API. | [optional] [default to null]
**LineItemUid** | **string** | The &#x60;line_item_uid&#x60; of the gift card’s line item in the order associated with the activity. It is populated along with &#x60;order_id&#x60; and is required if using the Square Orders API. | [optional] [default to null]
**ReferenceId** | **string** | A client-specified ID to associate an entity, in another system, with this gift card activity. This can be used to track the order or payment related information when the Square Orders API is not being used. | [optional] [default to null]
**BuyerPaymentInstrumentIds** | **[]string** | If you are not using the Orders API, this field is required because it is used to identify a buyer  to perform compliance checks. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

