# BankAccountPaymentDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankName** | **string** | The name of the bank associated with the bank account. | [optional] [default to null]
**TransferType** | **string** | The type of the bank transfer. The type can be &#x60;ACH&#x60; or &#x60;UNKNOWN&#x60;. | [optional] [default to null]
**AccountOwnershipType** | **string** | The ownership type of the bank account performing the transfer. The type can be &#x60;INDIVIDUAL&#x60;, &#x60;COMPANY&#x60;, or &#x60;UNKNOWN&#x60;. | [optional] [default to null]
**Fingerprint** | **string** | Uniquely identifies the bank account for this seller and can be used to determine if payments are from the same bank account. | [optional] [default to null]
**Country** | **string** | The two-letter ISO code representing the country the bank account is located in. | [optional] [default to null]
**StatementDescription** | **string** | The statement description as sent to the bank. | [optional] [default to null]
**AchDetails** | [***AchDetails**](ACHDetails.md) |  | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) | Information about errors encountered during the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

