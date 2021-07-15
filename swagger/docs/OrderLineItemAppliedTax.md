# OrderLineItemAppliedTax

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | A unique ID that identifies the applied tax only within this order. | [optional] [default to null]
**TaxUid** | **string** | The &#x60;uid&#x60; of the tax for which this applied tax represents. It must reference a tax present in the &#x60;order.taxes&#x60; field.  This field is immutable. To change which taxes apply to a line item, delete and add a new &#x60;OrderLineItemAppliedTax&#x60;. | [default to null]
**AppliedMoney** | [***Money**](Money.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

