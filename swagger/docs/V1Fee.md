# V1Fee

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The fee&#x27;s unique ID. | [optional] [default to null]
**Name** | **string** | The fee&#x27;s name. | [optional] [default to null]
**Rate** | **string** | The rate of the fee, as a string representation of a decimal number. A value of 0.07 corresponds to a rate of 7%. | [optional] [default to null]
**CalculationPhase** | [***V1FeeCalculationPhase**](V1FeeCalculationPhase.md) |  | [optional] [default to null]
**AdjustmentType** | [***V1FeeAdjustmentType**](V1FeeAdjustmentType.md) |  | [optional] [default to null]
**AppliesToCustomAmounts** | **bool** | If true, the fee applies to custom amounts entered into Square Point of Sale that are not associated with a particular item. | [optional] [default to null]
**Enabled** | **bool** | If true, the fee is applied to all appropriate items. If false, the fee is not applied at all. | [optional] [default to null]
**InclusionType** | [***V1FeeInclusionType**](V1FeeInclusionType.md) |  | [optional] [default to null]
**Type_** | [***V1FeeType**](V1FeeType.md) |  | [optional] [default to null]
**V2Id** | **string** | The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

