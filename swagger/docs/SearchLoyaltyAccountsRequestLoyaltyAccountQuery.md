# SearchLoyaltyAccountsRequestLoyaltyAccountQuery

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mappings** | [**[]LoyaltyAccountMapping**](LoyaltyAccountMapping.md) | The set of mappings to use in the loyalty account search.    This cannot be combined with &#x60;customer_ids&#x60;.    Max: 30 mappings | [optional] [default to null]
**CustomerIds** | **[]string** | The set of customer IDs to use in the loyalty account search.    This cannot be combined with &#x60;mappings&#x60;.    Max: 30 customer IDs | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

