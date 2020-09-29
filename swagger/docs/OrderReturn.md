# OrderReturn

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique ID that identifies the return only within this order. | [optional] [default to null]
**SourceOrderId** | **string** | Order which contains the original sale of these returned line items. This will be unset for unlinked returns. | [optional] [default to null]
**ReturnLineItems** | [**[]OrderReturnLineItem**](OrderReturnLineItem.md) | Collection of line items which are being returned. | [optional] [default to null]
**ReturnServiceCharges** | [**[]OrderReturnServiceCharge**](OrderReturnServiceCharge.md) | Collection of service charges which are being returned. | [optional] [default to null]
**ReturnTaxes** | [**[]OrderReturnTax**](OrderReturnTax.md) | Collection of references to taxes being returned for an order, including the total applied tax amount to be returned. The taxes must reference a top-level tax ID from the source order. | [optional] [default to null]
**ReturnDiscounts** | [**[]OrderReturnDiscount**](OrderReturnDiscount.md) | Collection of references to discounts being returned for an order, including the total applied discount amount to be returned. The discounts must reference a top-level discount ID from the source order. | [optional] [default to null]
**RoundingAdjustment** | [***OrderRoundingAdjustment**](OrderRoundingAdjustment.md) |  | [optional] [default to null]
**ReturnAmounts** | [***OrderMoneyAmounts**](OrderMoneyAmounts.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

