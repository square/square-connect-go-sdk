# V1Payment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The payment&#x27;s unique identifier. | [optional] [default to null]
**MerchantId** | **string** | The unique identifier of the merchant that took the payment. | [optional] [default to null]
**CreatedAt** | **string** | The time when the payment was created, in ISO 8601 format. Reflects the time of the first payment if the object represents an incomplete partial payment, and the time of the last or complete payment otherwise. | [optional] [default to null]
**CreatorId** | **string** | The unique identifier of the Square account that took the payment. | [optional] [default to null]
**Device** | [***Device**](Device.md) |  | [optional] [default to null]
**PaymentUrl** | **string** | The URL of the payment&#x27;s detail page in the merchant dashboard. The merchant must be signed in to the merchant dashboard to view this page. | [optional] [default to null]
**ReceiptUrl** | **string** | The URL of the receipt for the payment. Note that for split tender payments, this URL corresponds to the receipt for the first tender listed in the payment&#x27;s tender field. Each Tender object has its own receipt_url field you can use to get the other receipts associated with a split tender payment. | [optional] [default to null]
**InclusiveTaxMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**AdditiveTaxMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**TaxMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**TipMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**DiscountMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**TotalCollectedMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**ProcessingFeeMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**NetTotalMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**RefundedMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**SwedishRoundingMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**GrossSalesMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**NetSalesMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**InclusiveTax** | [**[]V1PaymentTax**](V1PaymentTax.md) | All of the inclusive taxes associated with the payment. | [optional] [default to null]
**AdditiveTax** | [**[]V1PaymentTax**](V1PaymentTax.md) | All of the additive taxes associated with the payment. | [optional] [default to null]
**Tender** | [**[]V1Tender**](V1Tender.md) | All of the tenders associated with the payment. | [optional] [default to null]
**Refunds** | [**[]V1Refund**](V1Refund.md) | All of the refunds applied to the payment. Note that the value of all refunds on a payment can exceed the value of all tenders if a merchant chooses to refund money to a tender after previously accepting returned goods as part of an exchange. | [optional] [default to null]
**Itemizations** | [**[]V1PaymentItemization**](V1PaymentItemization.md) | The items purchased in the payment. | [optional] [default to null]
**SurchargeMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**Surcharges** | [**[]V1PaymentSurcharge**](V1PaymentSurcharge.md) | A list of all surcharges associated with the payment. | [optional] [default to null]
**IsPartial** | **bool** | Indicates whether or not the payment is only partially paid for. If true, this payment will have the tenders collected so far, but the itemizations will be empty until the payment is completed. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

