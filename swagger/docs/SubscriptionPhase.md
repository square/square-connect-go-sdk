# SubscriptionPhase

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | The Square-assigned ID of the subscription phase. This field cannot be changed after a &#x60;SubscriptionPhase&#x60; is created. | [optional] [default to null]
**Cadence** | [***SubscriptionCadence**](SubscriptionCadence.md) |  | [default to null]
**Periods** | **int32** | The number of &#x60;cadence&#x60;s the phase lasts. If not set, the phase never ends. Only the last phase can be indefinite. This field cannot be changed after a &#x60;SubscriptionPhase&#x60; is created. | [optional] [default to null]
**RecurringPriceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Ordinal** | **int64** | The position this phase appears in the sequence of phases defined for the plan, indexed from 0. This field cannot be changed after a &#x60;SubscriptionPhase&#x60; is created. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

