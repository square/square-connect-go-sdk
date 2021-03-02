# CardPaymentDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | The card payment&#x27;s current state. The state can be AUTHORIZED, CAPTURED, VOIDED, or FAILED. | [optional] [default to null]
**Card** | [***Card**](Card.md) |  | [optional] [default to null]
**EntryMethod** | **string** | The method used to enter the card&#x27;s details for the payment. The method can be &#x60;KEYED&#x60;, &#x60;SWIPED&#x60;, &#x60;EMV&#x60;, &#x60;ON_FILE&#x60;, or &#x60;CONTACTLESS&#x60;. | [optional] [default to null]
**CvvStatus** | **string** | The status code returned from the Card Verification Value (CVV) check. The code can be &#x60;CVV_ACCEPTED&#x60;, &#x60;CVV_REJECTED&#x60;, or &#x60;CVV_NOT_CHECKED&#x60;. | [optional] [default to null]
**AvsStatus** | **string** | The status code returned from the Address Verification System (AVS) check. The code can be &#x60;AVS_ACCEPTED&#x60;, &#x60;AVS_REJECTED&#x60;, or &#x60;AVS_NOT_CHECKED&#x60;. | [optional] [default to null]
**AuthResultCode** | **string** | The status code returned by the card issuer that describes the payment&#x27;s authorization status. | [optional] [default to null]
**ApplicationIdentifier** | **string** | For EMV payments, the application ID identifies the EMV application used for the payment. | [optional] [default to null]
**ApplicationName** | **string** | For EMV payments, the human-readable name of the EMV application used for the payment. | [optional] [default to null]
**ApplicationCryptogram** | **string** | For EMV payments, the cryptogram generated for the payment. | [optional] [default to null]
**VerificationMethod** | **string** | For EMV payments, the method used to verify the cardholder&#x27;s identity. The method can be &#x60;PIN&#x60;, &#x60;SIGNATURE&#x60;, &#x60;PIN_AND_SIGNATURE&#x60;, &#x60;ON_DEVICE&#x60;, or &#x60;NONE&#x60;. | [optional] [default to null]
**VerificationResults** | **string** | For EMV payments, the results of the cardholder verification. The result can be &#x60;SUCCESS&#x60;, &#x60;FAILURE&#x60;, or &#x60;UNKNOWN&#x60;. | [optional] [default to null]
**StatementDescription** | **string** | The statement description sent to the card networks.  Note: The actual statement description varies and is likely to be truncated and appended with additional information on a per issuer basis. | [optional] [default to null]
**DeviceDetails** | [***DeviceDetails**](DeviceDetails.md) |  | [optional] [default to null]
**CardPaymentTimeline** | [***CardPaymentTimeline**](CardPaymentTimeline.md) |  | [optional] [default to null]
**RefundRequiresCardPresence** | **bool** | Whether the card must be physically present for the payment to be refunded.  If set to &#x60;true&#x60;, the card must be present. | [optional] [default to null]
**Errors** | [**[]ModelError**](Error.md) | Information about errors encountered during the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

