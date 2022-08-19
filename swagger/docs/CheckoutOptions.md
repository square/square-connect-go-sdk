# CheckoutOptions

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowTipping** | **bool** | Indicates whether the payment allows tipping. | [optional] [default to null]
**CustomFields** | [**[]CustomField**](CustomField.md) | The custom fields requesting information from the buyer. | [optional] [default to null]
**SubscriptionPlanId** | **string** | The ID of the subscription plan for the buyer to pay and subscribe.  For more information, see [Subscription Plan Checkout](https://developer.squareup.com/docs/checkout-api/subscription-plan-checkout). | [optional] [default to null]
**RedirectUrl** | **string** | The confirmation page URL to redirect the buyer to after Square processes the payment. | [optional] [default to null]
**MerchantSupportEmail** | **string** | The email address that buyers can use to contact the seller. | [optional] [default to null]
**AskForShippingAddress** | **bool** | Indicates whether to include the address fields in the payment form. | [optional] [default to null]
**AcceptedPaymentMethods** | [***AcceptedPaymentMethods**](AcceptedPaymentMethods.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

