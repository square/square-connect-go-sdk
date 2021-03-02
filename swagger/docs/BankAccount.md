# BankAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique, Square-issued identifier for the bank account. | [default to null]
**AccountNumberSuffix** | **string** | The last few digits of the account number. | [default to null]
**Country** | [***Country**](Country.md) |  | [default to null]
**Currency** | [***Currency**](Currency.md) |  | [default to null]
**AccountType** | [***BankAccountType**](BankAccountType.md) |  | [default to null]
**HolderName** | **string** | Name of the account holder. This name must match the name  on the targeted bank account record. | [default to null]
**PrimaryBankIdentificationNumber** | **string** | Primary identifier for the bank. For more information, see  [Bank Accounts API](https://developer.squareup.com/docs/bank-accounts-api). | [default to null]
**SecondaryBankIdentificationNumber** | **string** | Secondary identifier for the bank. For more information, see  [Bank Accounts API](https://developer.squareup.com/docs/bank-accounts-api). | [optional] [default to null]
**DebitMandateReferenceId** | **string** | Reference identifier that will be displayed to UK bank account owners when collecting direct debit authorization. Only required for UK bank accounts. | [optional] [default to null]
**ReferenceId** | **string** | Client-provided identifier for linking the banking account to an entity in a third-party system (for example, a bank account number or a user identifier). | [optional] [default to null]
**LocationId** | **string** | The location to which the bank account belongs. | [optional] [default to null]
**Status** | [***BankAccountStatus**](BankAccountStatus.md) |  | [default to null]
**Creditable** | **bool** | Indicates whether it is possible for Square to send money to this bank account. | [default to null]
**Debitable** | **bool** | Indicates whether it is possible for Square to take money from this  bank account. | [default to null]
**Fingerprint** | **string** | A Square-assigned, unique identifier for the bank account based on the account information. The account fingerprint can be used to compare account entries and determine if the they represent the same real-world bank account. | [optional] [default to null]
**Version** | **int32** | The current version of the &#x60;BankAccount&#x60;. | [optional] [default to null]
**BankName** | **string** | Read only. Name of actual financial institution.  For example \&quot;Bank of America\&quot;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

