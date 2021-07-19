# GiftCardActivityRedeem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountMoney** | [***Money**](Money.md) |  | [default to null]
**PaymentId** | **string** | When the Square Payments API is used, Redeem is not called on the Gift Cards API. However, when Square reads a Redeem activity from the Gift Cards API, developers need to know the associated &#x60;payment_id&#x60;. | [optional] [default to null]
**ReferenceId** | **string** | A client-specified ID to associate an entity, in another system, with this gift card activity. This can be used to track the order or payment related information when the Square Orders API is not being used. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

