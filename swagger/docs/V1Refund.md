# V1Refund

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | [***V1RefundType**](V1RefundType.md) |  | [optional] [default to null]
**Reason** | **string** | The merchant-specified reason for the refund. | [optional] [default to null]
**RefundedMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedProcessingFeeMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedTaxMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedAdditiveTaxMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedAdditiveTax** | [**[]V1PaymentTax**](V1PaymentTax.md) | All of the additive taxes associated with the refund. | [optional] [default to null]
**RefundedInclusiveTaxMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedInclusiveTax** | [**[]V1PaymentTax**](V1PaymentTax.md) | All of the inclusive taxes associated with the refund. | [optional] [default to null]
**RefundedTipMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedDiscountMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedSurchargeMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedSurcharges** | [**[]V1PaymentSurcharge**](V1PaymentSurcharge.md) | A list of all surcharges associated with the refund. | [optional] [default to null]
**CreatedAt** | **string** | The time when the merchant initiated the refund for Square to process, in ISO 8601 format. | [optional] [default to null]
**ProcessedAt** | **string** | The time when Square processed the refund on behalf of the merchant, in ISO 8601 format. | [optional] [default to null]
**PaymentId** | **string** | A Square-issued ID associated with the refund. For single-tender refunds, payment_id is the ID of the original payment ID. For split-tender refunds, payment_id is the ID of the original tender. For exchange-based refunds (is_exchange &#x3D;&#x3D; true), payment_id is the ID of the original payment ID even if the payment includes other tenders. | [optional] [default to null]
**MerchantId** | **string** |  | [optional] [default to null]
**IsExchange** | **bool** | Indicates whether or not the refund is associated with an exchange. If is_exchange is true, the refund reflects the value of goods returned in the exchange not the total money refunded. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

