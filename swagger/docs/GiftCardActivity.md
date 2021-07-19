# GiftCardActivity

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique ID of the gift card activity. | [optional] [default to null]
**Type_** | [***GiftCardActivityType**](GiftCardActivityType.md) |  | [default to null]
**LocationId** | **string** | The ID of the location at which the activity occurred. | [default to null]
**CreatedAt** | **string** | The timestamp when the gift card activity was created, in RFC 3339 format. | [optional] [default to null]
**GiftCardId** | **string** | The gift card ID. The ID is not required if a GAN is present. | [optional] [default to null]
**GiftCardGan** | **string** | The gift card GAN. The GAN is not required if &#x60;gift_card_id&#x60; is present. | [optional] [default to null]
**GiftCardBalanceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**LoadActivityDetails** | [***GiftCardActivityLoad**](GiftCardActivityLoad.md) |  | [optional] [default to null]
**ActivateActivityDetails** | [***GiftCardActivityActivate**](GiftCardActivityActivate.md) |  | [optional] [default to null]
**RedeemActivityDetails** | [***GiftCardActivityRedeem**](GiftCardActivityRedeem.md) |  | [optional] [default to null]
**ClearBalanceActivityDetails** | [***GiftCardActivityClearBalance**](GiftCardActivityClearBalance.md) |  | [optional] [default to null]
**DeactivateActivityDetails** | [***GiftCardActivityDeactivate**](GiftCardActivityDeactivate.md) |  | [optional] [default to null]
**AdjustIncrementActivityDetails** | [***GiftCardActivityAdjustIncrement**](GiftCardActivityAdjustIncrement.md) |  | [optional] [default to null]
**AdjustDecrementActivityDetails** | [***GiftCardActivityAdjustDecrement**](GiftCardActivityAdjustDecrement.md) |  | [optional] [default to null]
**RefundActivityDetails** | [***GiftCardActivityRefund**](GiftCardActivityRefund.md) |  | [optional] [default to null]
**UnlinkedActivityRefundActivityDetails** | [***GiftCardActivityUnlinkedActivityRefund**](GiftCardActivityUnlinkedActivityRefund.md) |  | [optional] [default to null]
**ImportActivityDetails** | [***GiftCardActivityImport**](GiftCardActivityImport.md) |  | [optional] [default to null]
**BlockActivityDetails** | [***GiftCardActivityBlock**](GiftCardActivityBlock.md) |  | [optional] [default to null]
**UnblockActivityDetails** | [***GiftCardActivityUnblock**](GiftCardActivityUnblock.md) |  | [optional] [default to null]
**ImportReversalActivityDetails** | [***GiftCardActivityImportReversal**](GiftCardActivityImportReversal.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

