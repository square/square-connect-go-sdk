# PayoutEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID for the payout entry. | [default to null]
**PayoutId** | **string** | The ID of the payout entries’ associated payout. | [default to null]
**EffectiveAt** | **string** | The timestamp of when the payout entry affected the balance, in RFC 3339 format. | [optional] [default to null]
**Type_** | [***ActivityType**](ActivityType.md) |  | [optional] [default to null]
**GrossAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**FeeAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**NetAmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TypeAutomaticSavingsDetails** | [***PaymentBalanceActivityAutomaticSavingsDetail**](PaymentBalanceActivityAutomaticSavingsDetail.md) |  | [optional] [default to null]
**TypeAutomaticSavingsReversedDetails** | [***PaymentBalanceActivityAutomaticSavingsReversedDetail**](PaymentBalanceActivityAutomaticSavingsReversedDetail.md) |  | [optional] [default to null]
**TypeChargeDetails** | [***PaymentBalanceActivityChargeDetail**](PaymentBalanceActivityChargeDetail.md) |  | [optional] [default to null]
**TypeDepositFeeDetails** | [***PaymentBalanceActivityDepositFeeDetail**](PaymentBalanceActivityDepositFeeDetail.md) |  | [optional] [default to null]
**TypeDisputeDetails** | [***PaymentBalanceActivityDisputeDetail**](PaymentBalanceActivityDisputeDetail.md) |  | [optional] [default to null]
**TypeFeeDetails** | [***PaymentBalanceActivityFeeDetail**](PaymentBalanceActivityFeeDetail.md) |  | [optional] [default to null]
**TypeFreeProcessingDetails** | [***PaymentBalanceActivityFreeProcessingDetail**](PaymentBalanceActivityFreeProcessingDetail.md) |  | [optional] [default to null]
**TypeHoldAdjustmentDetails** | [***PaymentBalanceActivityHoldAdjustmentDetail**](PaymentBalanceActivityHoldAdjustmentDetail.md) |  | [optional] [default to null]
**TypeOpenDisputeDetails** | [***PaymentBalanceActivityOpenDisputeDetail**](PaymentBalanceActivityOpenDisputeDetail.md) |  | [optional] [default to null]
**TypeOtherDetails** | [***PaymentBalanceActivityOtherDetail**](PaymentBalanceActivityOtherDetail.md) |  | [optional] [default to null]
**TypeOtherAdjustmentDetails** | [***PaymentBalanceActivityOtherAdjustmentDetail**](PaymentBalanceActivityOtherAdjustmentDetail.md) |  | [optional] [default to null]
**TypeRefundDetails** | [***PaymentBalanceActivityRefundDetail**](PaymentBalanceActivityRefundDetail.md) |  | [optional] [default to null]
**TypeReleaseAdjustmentDetails** | [***PaymentBalanceActivityReleaseAdjustmentDetail**](PaymentBalanceActivityReleaseAdjustmentDetail.md) |  | [optional] [default to null]
**TypeReserveHoldDetails** | [***PaymentBalanceActivityReserveHoldDetail**](PaymentBalanceActivityReserveHoldDetail.md) |  | [optional] [default to null]
**TypeReserveReleaseDetails** | [***PaymentBalanceActivityReserveReleaseDetail**](PaymentBalanceActivityReserveReleaseDetail.md) |  | [optional] [default to null]
**TypeSquareCapitalPaymentDetails** | [***PaymentBalanceActivitySquareCapitalPaymentDetail**](PaymentBalanceActivitySquareCapitalPaymentDetail.md) |  | [optional] [default to null]
**TypeSquareCapitalReversedPaymentDetails** | [***PaymentBalanceActivitySquareCapitalReversedPaymentDetail**](PaymentBalanceActivitySquareCapitalReversedPaymentDetail.md) |  | [optional] [default to null]
**TypeTaxOnFeeDetails** | [***PaymentBalanceActivityTaxOnFeeDetail**](PaymentBalanceActivityTaxOnFeeDetail.md) |  | [optional] [default to null]
**TypeThirdPartyFeeDetails** | [***PaymentBalanceActivityThirdPartyFeeDetail**](PaymentBalanceActivityThirdPartyFeeDetail.md) |  | [optional] [default to null]
**TypeThirdPartyFeeRefundDetails** | [***PaymentBalanceActivityThirdPartyFeeRefundDetail**](PaymentBalanceActivityThirdPartyFeeRefundDetail.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

