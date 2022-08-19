# LoyaltyAccount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the loyalty account. | [optional] [default to null]
**Mappings** | [**[]LoyaltyAccountMapping**](LoyaltyAccountMapping.md) | The list of mappings that the account is associated with.  Currently, a buyer can only be mapped to a loyalty account using  a phone number. Therefore, the list can only have one mapping. RETIRED at version 2021-05-13. Replaced by the &#x60;mapping&#x60; field. | [optional] [default to null]
**ProgramId** | **string** | The Square-assigned ID of the [loyalty program](entity:LoyaltyProgram) to which the account belongs. | [default to null]
**Balance** | **int32** | The available point balance in the loyalty account. If points are scheduled to expire, they are listed in the &#x60;expiring_point_deadlines&#x60; field.  Your application should be able to handle loyalty accounts that have a negative point balance (&#x60;balance&#x60; is less than 0). This might occur if a seller makes a manual adjustment or as a result of a refund or exchange. | [optional] [default to null]
**LifetimePoints** | **int32** | The total points accrued during the lifetime of the account. | [optional] [default to null]
**CustomerId** | **string** | The Square-assigned ID of the [customer](entity:Customer) that is associated with the account. | [optional] [default to null]
**EnrolledAt** | **string** | The timestamp when the buyer joined the loyalty program, in RFC 3339 format. This field is used to display the **Enrolled On** or **Member Since** date in first-party Square products.  If this field is not set in a &#x60;CreateLoyaltyAccount&#x60; request, Square populates it after the buyer&#x27;s first action on their account  (when &#x60;AccumulateLoyaltyPoints&#x60; or &#x60;CreateLoyaltyReward&#x60; is called). In first-party flows, Square populates the field when the buyer agrees to the terms of service in Square Point of Sale.   This field is typically specified in a &#x60;CreateLoyaltyAccount&#x60; request when creating a loyalty account for a buyer who already interacted with their account.  For example, you would set this field when migrating accounts from an external system. The timestamp in the request can represent a current or previous date and time, but it cannot be set for the future. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the loyalty account was created, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp when the loyalty account was last updated, in RFC 3339 format. | [optional] [default to null]
**Mapping** | [***LoyaltyAccountMapping**](LoyaltyAccountMapping.md) |  | [optional] [default to null]
**ExpiringPointDeadlines** | [**[]LoyaltyAccountExpiringPointDeadline**](LoyaltyAccountExpiringPointDeadline.md) | The schedule for when points expire in the loyalty account balance. This field is present only if the account has points that are scheduled to expire.   The total number of points in this field equals the number of points in the &#x60;balance&#x60; field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

