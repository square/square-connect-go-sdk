# CatalogTax

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The tax&#x27;s name. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points. | [optional] [default to null]
**CalculationPhase** | [***TaxCalculationPhase**](TaxCalculationPhase.md) |  | [optional] [default to null]
**InclusionType** | [***TaxInclusionType**](TaxInclusionType.md) |  | [optional] [default to null]
**Percentage** | **string** | The percentage of the tax in decimal form, using a &#x60;&#x27;.&#x27;&#x60; as the decimal separator and without a &#x60;&#x27;%&#x27;&#x60; sign. A value of &#x60;7.5&#x60; corresponds to 7.5%. For a location-specific tax rate, contact the tax authority of the location or a tax consultant. | [optional] [default to null]
**AppliesToCustomAmounts** | **bool** | If &#x60;true&#x60;, the fee applies to custom amounts entered into the Square Point of Sale app that are not associated with a particular &#x60;CatalogItem&#x60;. | [optional] [default to null]
**Enabled** | **bool** | A Boolean flag to indicate whether the tax is displayed as enabled (&#x60;true&#x60;) in the Square Point of Sale app or not (&#x60;false&#x60;). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

