# SubscriptionTestResult

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A Square-generated unique ID for the subscription test result. | [optional] [default to null]
**StatusCode** | **int32** | The status code returned by the subscription notification URL. | [optional] [default to null]
**Payload** | **string** | An object containing the payload of the test event. For example, a &#x60;payment.created&#x60; event. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp of when the subscription was created, in RFC 3339 format.  For example, \&quot;2016-09-04T23:59:33.123Z\&quot;. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp of when the subscription was updated, in RFC 3339 format. For example, \&quot;2016-09-04T23:59:33.123Z\&quot;. Because a subscription test result is unique, this field is the same as the &#x60;created_at&#x60; field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

