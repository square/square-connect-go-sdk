# OrderReturnServiceCharge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique ID that identifies the return service charge only within this order. | [optional] [default to null]
**SourceServiceChargeUid** | **string** | &#x60;uid&#x60; of the Service Charge from the Order containing the original charge of the service charge. &#x60;source_service_charge_uid&#x60; is &#x60;null&#x60; for unlinked returns. | [optional] [default to null]
**Name** | **string** | The name of the service charge. | [optional] [default to null]
**CatalogObjectId** | **string** | The catalog object ID of the associated [CatalogServiceCharge](#type-catalogservicecharge). | [optional] [default to null]
**Percentage** | **string** | The percentage of the service charge, as a string representation of a decimal number. For example, a value of &#x60;\&quot;7.25\&quot;&#x60; corresponds to a percentage of 7.25%.  Exactly one of &#x60;percentage&#x60; or &#x60;amount_money&#x60; should be set. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**AppliedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**TotalTaxMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**CalculationPhase** | [***OrderServiceChargeCalculationPhase**](OrderServiceChargeCalculationPhase.md) |  | [optional] [default to null]
**Taxable** | **bool** | Indicates whether the surcharge can be taxed. Service charges calculated in the &#x60;TOTAL_PHASE&#x60; cannot be marked as taxable. | [optional] [default to null]
**AppliedTaxes** | [**[]OrderLineItemAppliedTax**](OrderLineItemAppliedTax.md) | The list of references to &#x60;OrderReturnTax&#x60; entities applied to the &#x60;OrderReturnServiceCharge&#x60;. Each &#x60;OrderLineItemAppliedTax&#x60; has a &#x60;tax_uid&#x60; that references the &#x60;uid&#x60; of a top-level &#x60;OrderReturnTax&#x60; that is being applied to the &#x60;OrderReturnServiceCharge&#x60;. On reads, the amount applied is populated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

