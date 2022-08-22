# GiftCardActivityRedeem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**PaymentId** | **string** | The ID of the payment that represents the gift card redemption. Square populates this field  if the payment was processed by Square. | [optional] [default to null]
**ReferenceId** | **string** | A client-specified ID that associates the gift card activity with an entity in another system.   Applications that use a custom payment processing system can use this field to track information related to an order or payment. | [optional] [default to null]
**Status** | [***GiftCardActivityRedeemStatus**](GiftCardActivityRedeemStatus.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

