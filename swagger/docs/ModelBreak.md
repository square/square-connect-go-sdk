# ModelBreak

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID for this object. | [optional] [default to null]
**StartAt** | **string** | RFC 3339; follows the same timezone information as &#x60;Shift&#x60;. Precision up to the minute is respected; seconds are truncated. | [default to null]
**EndAt** | **string** | RFC 3339; follows the same timezone information as &#x60;Shift&#x60;. Precision up to the minute is respected; seconds are truncated. | [optional] [default to null]
**BreakTypeId** | **string** | The &#x60;BreakType&#x60; that this &#x60;Break&#x60; was templated on. | [default to null]
**Name** | **string** | A human-readable name. | [default to null]
**ExpectedDuration** | **string** | Format: RFC-3339 P[n]Y[n]M[n]DT[n]H[n]M[n]S. The expected length of the break. | [default to null]
**IsPaid** | **bool** | Whether this break counts towards time worked for compensation purposes. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

