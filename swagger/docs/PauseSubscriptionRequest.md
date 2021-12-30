# PauseSubscriptionRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PauseEffectiveDate** | **string** | The &#x60;YYYY-MM-DD&#x60;-formatted date when the scheduled &#x60;PAUSE&#x60; action takes place on the subscription.  When this date is unspecified or falls within the current billing cycle, the subscription is paused on the starting date of the next billing cycle. | [optional] [default to null]
**PauseCycleDuration** | **int64** | The number of billing cycles the subscription will be paused before it is reactivated.   When this is set, a &#x60;RESUME&#x60; action is also scheduled to take place on the subscription at  the end of the specified pause cycle duration. In this case, neither &#x60;resume_effective_date&#x60;  nor &#x60;resume_change_timing&#x60; may be specified. | [optional] [default to null]
**ResumeEffectiveDate** | **string** | The date when the subscription is reactivated by a scheduled &#x60;RESUME&#x60; action.  This date must be at least one billing cycle ahead of &#x60;pause_effective_date&#x60;. | [optional] [default to null]
**ResumeChangeTiming** | [***ChangeTiming**](ChangeTiming.md) |  | [optional] [default to null]
**PauseReason** | **string** | The user-provided reason to pause the subscription. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

