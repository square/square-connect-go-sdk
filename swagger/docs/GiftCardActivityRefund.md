# GiftCardActivityRefund

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RedeemActivityId** | **string** | The ID of the refunded &#x60;REDEEM&#x60; gift card activity. Square populates this field if the  &#x60;payment_id&#x60; in the corresponding [RefundPayment](api-endpoint:Refunds-RefundPayment) request  represents a redemption made by the same gift card.  Applications that use a custom payment processing system can use this field in a  [CreateGiftCardActivity](api-endpoint:GiftCardActivities-CreateGiftCardActivity)  request to link a refund with a &#x60;REDEEM&#x60; activity for the same gift card. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**ReferenceId** | **string** | A client-specified ID that associates the gift card activity with an entity in another system.   Applications that use a custom payment processing system can use this field to track information related to an order or payment. | [optional] [default to null]
**PaymentId** | **string** | The ID of the refunded payment. Square populates this field if the refund is for a  payment processed by Square. The payment source can be the same gift card or a cross-tender payment from a  credit card or a different gift card. Cross-tender payments can only be refunded from Square Point of Sale  or other Square products. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

