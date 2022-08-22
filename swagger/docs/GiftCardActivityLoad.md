# GiftCardActivityLoad

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**OrderId** | **string** | The ID of the [order](entity:Order) that contains the &#x60;GIFT_CARD&#x60; line item.  Applications that use the Square Orders API to process orders must specify the order ID in the  [CreateGiftCardActivity](api-endpoint:GiftCardActivities-CreateGiftCardActivity) request. | [optional] [default to null]
**LineItemUid** | **string** | The UID of the &#x60;GIFT_CARD&#x60; line item in the order that represents the additional funds for the gift card.  Applications that use the Square Orders API to process orders must specify the line item UID in the [CreateGiftCardActivity](api-endpoint:GiftCardActivities-CreateGiftCardActivity) request. | [optional] [default to null]
**ReferenceId** | **string** | A client-specified ID that associates the gift card activity with an entity in another system.   Applications that use a custom order processing system can use this field to track information related to  an order or payment. | [optional] [default to null]
**BuyerPaymentInstrumentIds** | **[]string** | The payment instrument IDs used to process the order for the additional funds, such as a credit card ID  or bank account ID.   Applications that use a custom order processing system must specify payment instrument IDs in  the [CreateGiftCardActivity](api-endpoint:GiftCardActivities-CreateGiftCardActivity) request. Square uses this information to perform compliance checks.   For applications that use the Square Orders API to process payments, Square has the necessary  instrument IDs to perform compliance checks. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

