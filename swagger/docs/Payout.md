# Payout

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID for the payout. | [default to null]
**Status** | [***PayoutStatus**](PayoutStatus.md) |  | [optional] [default to null]
**LocationId** | **string** | The ID of the location associated with the payout. | [default to null]
**CreatedAt** | **string** | The timestamp of when the payout was created and submitted for deposit to the seller&#x27;s banking destination, in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | The timestamp of when the payout was last updated, in RFC 3339 format. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Destination** | [***Destination**](Destination.md) |  | [optional] [default to null]
**Version** | **int32** | The version number, which is incremented each time an update is made to this payout record. The version number helps developers receive event notifications or feeds out of order. | [optional] [default to null]
**Type_** | [***PayoutType**](PayoutType.md) |  | [optional] [default to null]
**PayoutFee** | [**[]PayoutFee**](PayoutFee.md) | A list of transfer fees and any taxes on the fees assessed by Square for this payout. | [optional] [default to null]
**ArrivalDate** | **string** | The calendar date, in ISO 8601 format (YYYY-MM-DD), when the payout is due to arrive in the seller’s banking destination. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

