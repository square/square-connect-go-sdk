# BreakType

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID for this object. | [optional] [default to null]
**LocationId** | **string** | The ID of the business location this type of break applies to. | [default to null]
**BreakName** | **string** | A human-readable name for this type of break. The name is displayed to employees in Square products. | [default to null]
**ExpectedDuration** | **string** | Format: RFC-3339 P[n]Y[n]M[n]DT[n]H[n]M[n]S. The expected length of this break. Precision less than minutes is truncated. | [default to null]
**IsPaid** | **bool** | Whether this break counts towards time worked for compensation purposes. | [default to null]
**Version** | **int32** | Used for resolving concurrency issues. The request fails if the version provided does not match the server version at the time of the request. If a value is not provided, Square&#x27;s servers execute a \&quot;blind\&quot; write; potentially overwriting another writer&#x27;s data. | [optional] [default to null]
**CreatedAt** | **string** | A read-only timestamp in RFC 3339 format. | [optional] [default to null]
**UpdatedAt** | **string** | A read-only timestamp in RFC 3339 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

