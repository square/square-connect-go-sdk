# PaymentLink

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the payment link. | [optional] [default to null]
**Version** | **int32** | The Square-assigned version number, which is incremented each time an update is committed to the payment link. | [default to null]
**Description** | **string** | The optional description of the &#x60;payment_link&#x60; object. It is primarily for use by your application and is not used anywhere. | [optional] [default to null]
**OrderId** | **string** | The ID of the order associated with the payment link. | [optional] [default to null]
**CheckoutOptions** | [***CheckoutOptions**](CheckoutOptions.md) |  | [optional] [default to null]
**PrePopulatedData** | [***PrePopulatedData**](PrePopulatedData.md) |  | [optional] [default to null]
**Url** | **string** | The URL of the payment link. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the payment link was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the payment link was last updated, in RFC 3339 format. | [optional] [default to null]
**PaymentNote** | **string** | An optional note. After Square processes the payment, this note is added to the resulting &#x60;Payment&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

