# CreateRefundRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdempotencyKey** | **string** | A value you specify that uniquely identifies this refund among refunds you&#x27;ve created for the tender.  If you&#x27;re unsure whether a particular refund succeeded, you can reattempt it with the same idempotency key without worrying about duplicating the refund.  See [Idempotency keys](https://developer.squareup.com/docs/working-with-apis/idempotency) for more information. | [default to null]
**TenderId** | **string** | The ID of the tender to refund.  A [&#x60;Transaction&#x60;](entity:Transaction) has one or more &#x60;tenders&#x60; (i.e., methods of payment) associated with it, and you refund each tender separately with the Connect API. | [default to null]
**Reason** | **string** | A description of the reason for the refund.  Default value: &#x60;Refund via API&#x60; | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

