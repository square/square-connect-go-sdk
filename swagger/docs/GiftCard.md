# GiftCard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Square-assigned ID of the gift card. | [optional] [default to null]
**Type_** | [***GiftCardType**](GiftCardType.md) |  | [default to null]
**GanSource** | [***GiftCardGanSource**](GiftCardGANSource.md) |  | [optional] [default to null]
**State** | [***GiftCardStatus**](GiftCardStatus.md) |  | [optional] [default to null]
**BalanceMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Gan** | **string** | The gift card account number (GAN). Buyers can use the GAN to make purchases or check  the gift card balance. | [optional] [default to null]
**CreatedAt** | **string** | The timestamp when the gift card was created, in RFC 3339 format.  In the case of a digital gift card, it is the time when you create a card  (using the Square Point of Sale application, Seller Dashboard, or Gift Cards API).   In the case of a plastic gift card, it is the time when Square associates the card with the  seller at the time of activation. | [optional] [default to null]
**CustomerIds** | **[]string** | The IDs of the [customer profiles](entity:Customer) to whom this gift card is linked. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

