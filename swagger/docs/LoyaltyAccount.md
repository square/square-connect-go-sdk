# LoyaltyAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the loyalty account. | [optional] [default to null]
**Mappings** | [**[]LoyaltyAccountMapping**](LoyaltyAccountMapping.md) | The list of mappings that the account is associated with.  Currently, a buyer can only be mapped to a loyalty account using  a phone number. Therefore, the list can only have one mapping. | [default to null]
**ProgramId** | **string** | The Square-assigned ID of the [loyalty program](#type-LoyaltyProgram) to which the account belongs. | [default to null]
**Balance** | **int32** | The available point balance in the loyalty account.    Your application should be able to handle loyalty accounts that have a negative point balance (&#x60;balance&#x60; is less than 0). This might occur if a seller makes a manual adjustment or as a result of a refund or exchange. | [optional] [default to null]
**LifetimePoints** | **int32** | The total points accrued during the lifetime of the account. | [optional] [default to null]
**CustomerId** | **string** | The Square-assigned ID of the [customer](#type-Customer) that is associated with the account. | [optional] [default to null]
**EnrolledAt** | **string** | The timestamp when enrollment occurred, in RFC 3339 format. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the loyalty account was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the loyalty account was last updated, in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

