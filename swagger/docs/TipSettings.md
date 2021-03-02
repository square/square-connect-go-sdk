# TipSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowTipping** | **bool** | Indicates whether tipping is enabled for this checkout. Defaults to false. | [optional] [default to null]
**SeparateTipScreen** | **bool** | Indicates whether tip options should be presented on their own screen before presenting the signature screen during card payment. Defaults to false. | [optional] [default to null]
**CustomTipField** | **bool** | Indicates whether custom tip amounts are allowed during the checkout flow. Defaults to false. | [optional] [default to null]
**TipPercentages** | **[]int32** | A list of tip percentages that should be presented during the checkout flow. Specified as up to 3 non-negative integers from 0 to 100 (inclusive). Defaults to [15, 20, 25] | [optional] [default to null]
**SmartTipping** | **bool** | Enables the \&quot;Smart Tip Amounts\&quot; behavior. Exact tipping options depend on the region the Square seller is active in.  In the United States and Canada, tipping options will be presented in whole dollar amounts for payments under 10 USD/CAD respectively.  If set to true, the tip_percentages settings is ignored. Defaults to false.  To learn more about smart tipping, see [Accept Tips with the Square App](https://squareup.com/help/us/en/article/5069-accept-tips-with-the-square-app) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

