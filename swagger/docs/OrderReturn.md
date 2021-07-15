# OrderReturn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | A unique ID that identifies the return only within this order. | [optional] [default to null]
**SourceOrderId** | **string** | An order that contains the original sale of these return line items. This is unset for unlinked returns. | [optional] [default to null]
**ReturnLineItems** | [**[]OrderReturnLineItem**](OrderReturnLineItem.md) | A collection of line items that are being returned. | [optional] [default to null]
**ReturnServiceCharges** | [**[]OrderReturnServiceCharge**](OrderReturnServiceCharge.md) | A collection of service charges that are being returned. | [optional] [default to null]
**ReturnTaxes** | [**[]OrderReturnTax**](OrderReturnTax.md) | A collection of references to taxes being returned for an order, including the total applied tax amount to be returned. The taxes must reference a top-level tax ID from the source order. | [optional] [default to null]
**ReturnDiscounts** | [**[]OrderReturnDiscount**](OrderReturnDiscount.md) | A collection of references to discounts being returned for an order, including the total applied discount amount to be returned. The discounts must reference a top-level discount ID from the source order. | [optional] [default to null]
**RoundingAdjustment** | [***OrderRoundingAdjustment**](OrderRoundingAdjustment.md) |  | [optional] [default to null]
**ReturnAmounts** | [***OrderMoneyAmounts**](OrderMoneyAmounts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

