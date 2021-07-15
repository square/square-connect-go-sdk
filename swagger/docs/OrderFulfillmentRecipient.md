# OrderFulfillmentRecipient

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerId** | **string** | The customer ID of the customer associated with the fulfillment.  If &#x60;customer_id&#x60; is provided, the fulfillment recipient&#x27;s &#x60;display_name&#x60;, &#x60;email_address&#x60;, and &#x60;phone_number&#x60; are automatically populated from the targeted customer profile. If these fields are set in the request, the request values overrides the information from the customer profile. If the targeted customer profile does not contain the necessary information and these fields are left unset, the request results in an error. | [optional] [default to null]
**DisplayName** | **string** | The display name of the fulfillment recipient.  If provided, the display name overrides the value pulled from the customer profile indicated by &#x60;customer_id&#x60;. | [optional] [default to null]
**EmailAddress** | **string** | The email address of the fulfillment recipient.  If provided, the email address overrides the value pulled from the customer profile indicated by &#x60;customer_id&#x60;. | [optional] [default to null]
**PhoneNumber** | **string** | The phone number of the fulfillment recipient.  If provided, the phone number overrides the value pulled from the customer profile indicated by &#x60;customer_id&#x60;. | [optional] [default to null]
**Address** | [***Address**](Address.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

