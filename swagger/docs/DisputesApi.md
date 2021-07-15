# {{classname}}

All URIs are relative to *https://connect.squareup.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptDispute**](DisputesApi.md#AcceptDispute) | **Post** /v2/disputes/{dispute_id}/accept | AcceptDispute
[**CreateDisputeEvidenceFile**](DisputesApi.md#CreateDisputeEvidenceFile) | **Post** /v2/disputes/{dispute_id}/evidence-files | CreateDisputeEvidenceFile
[**CreateDisputeEvidenceText**](DisputesApi.md#CreateDisputeEvidenceText) | **Post** /v2/disputes/{dispute_id}/evidence-text | CreateDisputeEvidenceText
[**DeleteDisputeEvidence**](DisputesApi.md#DeleteDisputeEvidence) | **Delete** /v2/disputes/{dispute_id}/evidence/{evidence_id} | DeleteDisputeEvidence
[**ListDisputeEvidence**](DisputesApi.md#ListDisputeEvidence) | **Get** /v2/disputes/{dispute_id}/evidence | ListDisputeEvidence
[**ListDisputes**](DisputesApi.md#ListDisputes) | **Get** /v2/disputes | ListDisputes
[**RetrieveDispute**](DisputesApi.md#RetrieveDispute) | **Get** /v2/disputes/{dispute_id} | RetrieveDispute
[**RetrieveDisputeEvidence**](DisputesApi.md#RetrieveDisputeEvidence) | **Get** /v2/disputes/{dispute_id}/evidence/{evidence_id} | RetrieveDisputeEvidence
[**SubmitEvidence**](DisputesApi.md#SubmitEvidence) | **Post** /v2/disputes/{dispute_id}/submit-evidence | SubmitEvidence

# **AcceptDispute**
> AcceptDisputeResponse AcceptDispute(ctx, disputeId)
AcceptDispute

Accepts the loss on a dispute. Square returns the disputed amount to the cardholder and updates the dispute state to ACCEPTED.  Square debits the disputed amount from the seller’s Square account. If the Square account does not have sufficient funds, Square debits the associated bank account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **disputeId** | **string**| The ID of the dispute you want to accept. | 

### Return type

[**AcceptDisputeResponse**](AcceptDisputeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDisputeEvidenceFile**
> CreateDisputeEvidenceFileResponse CreateDisputeEvidenceFile(ctx, request, imageFile, disputeId)
CreateDisputeEvidenceFile

Uploads a file to use as evidence in a dispute challenge. The endpoint accepts HTTP multipart/form-data file uploads in HEIC, HEIF, JPEG, PDF, PNG, and TIFF formats.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **request** | [**CreateDisputeEvidenceFileRequest**](.md)|  | 
  **imageFile** | ***os.File*****os.File**|  | 
  **disputeId** | **string**| The ID of the dispute you want to upload evidence for. | 

### Return type

[**CreateDisputeEvidenceFileResponse**](CreateDisputeEvidenceFileResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDisputeEvidenceText**
> CreateDisputeEvidenceTextResponse CreateDisputeEvidenceText(ctx, body, disputeId)
CreateDisputeEvidenceText

Uploads text to use as evidence for a dispute challenge.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDisputeEvidenceTextRequest**](CreateDisputeEvidenceTextRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **disputeId** | **string**| The ID of the dispute you want to upload evidence for. | 

### Return type

[**CreateDisputeEvidenceTextResponse**](CreateDisputeEvidenceTextResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDisputeEvidence**
> DeleteDisputeEvidenceResponse DeleteDisputeEvidence(ctx, disputeId, evidenceId)
DeleteDisputeEvidence

Removes specified evidence from a dispute.  Square does not send the bank any evidence that is removed. Also, you cannot remove evidence after submitting it to the bank using [SubmitEvidence](api-endpoint:Disputes-SubmitEvidence).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **disputeId** | **string**| The ID of the dispute you want to remove evidence from. | 
  **evidenceId** | **string**| The ID of the evidence you want to remove. | 

### Return type

[**DeleteDisputeEvidenceResponse**](DeleteDisputeEvidenceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDisputeEvidence**
> ListDisputeEvidenceResponse ListDisputeEvidence(ctx, disputeId, optional)
ListDisputeEvidence

Returns a list of evidence associated with a dispute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **disputeId** | **string**| The ID of the dispute. | 
 **optional** | ***DisputesApiListDisputeEvidenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisputesApiListDisputeEvidenceOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | 

### Return type

[**ListDisputeEvidenceResponse**](ListDisputeEvidenceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDisputes**
> ListDisputesResponse ListDisputes(ctx, optional)
ListDisputes

Returns a list of disputes associated with a particular account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DisputesApiListDisputesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DisputesApiListDisputesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **optional.String**| A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for the original query. For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination). | 
 **states** | [**optional.Interface of DisputeState**](.md)| The dispute states to filter the result. If not specified, the endpoint returns all open disputes (the dispute status is not &#x60;INQUIRY_CLOSED&#x60;, &#x60;WON&#x60;, or &#x60;LOST&#x60;). | 
 **locationId** | **optional.String**| The ID of the location for which to return a list of disputes. If not specified, the endpoint returns all open disputes (the dispute status is not &#x60;INQUIRY_CLOSED&#x60;, &#x60;WON&#x60;, or &#x60;LOST&#x60;) associated with all locations. | 

### Return type

[**ListDisputesResponse**](ListDisputesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveDispute**
> RetrieveDisputeResponse RetrieveDispute(ctx, disputeId)
RetrieveDispute

Returns details about a specific dispute.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **disputeId** | **string**| The ID of the dispute you want more details about. | 

### Return type

[**RetrieveDisputeResponse**](RetrieveDisputeResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RetrieveDisputeEvidence**
> RetrieveDisputeEvidenceResponse RetrieveDisputeEvidence(ctx, disputeId, evidenceId)
RetrieveDisputeEvidence

Returns the evidence metadata specified by the evidence ID in the request URL path  You must maintain a copy of the evidence you upload if you want to reference it later. You cannot download the evidence after you upload it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **disputeId** | **string**| The ID of the dispute that you want to retrieve evidence from. | 
  **evidenceId** | **string**| The ID of the evidence to retrieve. | 

### Return type

[**RetrieveDisputeEvidenceResponse**](RetrieveDisputeEvidenceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitEvidence**
> SubmitEvidenceResponse SubmitEvidence(ctx, disputeId)
SubmitEvidence

Submits evidence to the cardholder's bank.  Before submitting evidence, Square compiles all available evidence. This includes evidence uploaded using the [CreateDisputeEvidenceFile](api-endpoint:Disputes-CreateDisputeEvidenceFile) and [CreateDisputeEvidenceText](api-endpoint:Disputes-CreateDisputeEvidenceText) endpoints and evidence automatically provided by Square, when available.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **disputeId** | **string**| The ID of the dispute that you want to submit evidence for. | 

### Return type

[**SubmitEvidenceResponse**](SubmitEvidenceResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

