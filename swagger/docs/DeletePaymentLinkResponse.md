# DeletePaymentLinkResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) |  | [optional] [default to null]
**Id** | **string** | The ID of the link that is deleted. | [optional] [default to null]
**CancelledOrderId** | **string** | The ID of the order that is canceled. When a payment link is deleted, Square updates the the &#x60;state&#x60; (of the order that the checkout link created) to CANCELED. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

