# \LabelsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLabel**](LabelsApi.md#CreateLabel) | **Post** /api/v1/labels | Create a new label
[**DeleteAllLabels**](LabelsApi.md#DeleteAllLabels) | **Delete** /api/v1/labels | Delete all labels
[**DeleteLabel**](LabelsApi.md#DeleteLabel) | **Delete** /api/v1/labels/{id} | Delete a label
[**EditLabel**](LabelsApi.md#EditLabel) | **Patch** /api/v1/labels/{id} | Edit a label
[**GetAndSearchAllLabels**](LabelsApi.md#GetAndSearchAllLabels) | **Get** /api/v1/labels | Get and search all labels
[**GetLabel**](LabelsApi.md#GetLabel) | **Get** /api/v1/labels/{id} | Get a label
[**GetMultipleLabels**](LabelsApi.md#GetMultipleLabels) | **Get** /api/v1/labels/{ids} | Get multiple labels


# **CreateLabel**
> Label CreateLabel(ctx, body)
Create a new label

Create a new label for a specific user. Must pass the new label's name and the user's access_token.<br/><br/>Users can create labels to flag transactions (see method PATCH /transactions), with the goal of collecting and getting an overview of all transactions of a certain 'type'. In this sense, labels are similar to transaction categories. However, labels are supposed to depict more of an implicit meaning of a transaction. For instance, a user might want to assign a flag to a transaction that reminds him that he can offset it against tax. At the same time, the category of the transactions might be something like 'insurance', which is a more 'fact-based', or 'objective' way of typing the transaction. Despite this semantic difference between categories and labels, there is also the difference that a transaction can be assigned multiple labels at the same time (while in contrast it can have just a single category).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**LabelParams**](LabelParams.md)| Label&#39;s name | 

### Return type

[**Label**](Label.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllLabels**
> IdentifierList DeleteAllLabels(ctx, )
Delete all labels

Delete all labels of the user that is authorized by the access_token. Must pass the user's access_token.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdentifierList**](IdentifierList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLabel**
> DeleteLabel(ctx, id)
Delete a label

Delete a single label of the user that is authorized by the access_token. Must pass the label's identifier and the user's access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Identifier of the label to delete | 

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditLabel**
> Label EditLabel(ctx, id, body)
Edit a label

Change the name of a label of the user that is authorized by the access_token. Must pass the label's identifier, the label's new name and the user's access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Label&#39;s identifier | 
  **body** | [**LabelParams**](LabelParams.md)| Label&#39;s new name | 

### Return type

[**Label**](Label.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAndSearchAllLabels**
> PageableLabelList GetAndSearchAllLabels(ctx, optional)
Get and search all labels

Get labels of the user that is authorized by the access_token. Must pass the user's access_token. You can set optional search criteria to get only those labels that you are interested in. If you do not specify any search criteria, then this service functions as a 'get all' service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAndSearchAllLabelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAndSearchAllLabelsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []int64**](int64.md)| A comma-separated list of label identifiers. If specified, then only labels whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **search** | **optional.String**| If specified, then only those labels will be contained in the result whose &#39;name&#39; contains the given search string (the matching works case-insensitive). If no labels contain the search string in their name, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the name in order for a label to get included into the result. | 
 **page** | **optional.Int32**| Result page that you want to retrieve | [default to 1]
 **perPage** | **optional.Int32**| Maximum number of records per page. Can be at most 500. NOTE: Due to its validation and visualization, the swagger frontend might show very low performance, or even crashes, when a service responds with a lot of data. It is recommended to use a HTTP client like Postman or DHC instead of our swagger frontend for service calls with large page sizes. | [default to 20]
 **order** | [**optional.Interface of []string**](string.md)| Determines the order of the results. You can order the results by &#39;id&#39; or &#39;name&#39;. The default order for all services is &#39;id,asc&#39;. Since both fields (id and name) are unique, ordering by multiple fields is pointless. The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value.  | 

### Return type

[**PageableLabelList**](PageableLabelList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLabel**
> Label GetLabel(ctx, id)
Get a label

Get a single label of the user that is authorized by the access_token. Must pass the label's identifier and the user's access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Identifier of requested label | 

### Return type

[**Label**](Label.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMultipleLabels**
> LabelList GetMultipleLabels(ctx, ids)
Get multiple labels

Get a list of multiple labels of the user that is authorized by the access_token.Must pass the labels' identifiers and the user's access_token. Identifiers that do not exist or do not relate to the authorized user will not be contained in the result (If this applies to all of the given identifiers, then the result will be an empty list). WARNING: This service is deprecated and will be removed at some point. If you want to get multiple labels, please instead use the service 'Get all labels' and pass a comma-separated list of identifiers as a parameter 'ids'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]int64**](int64.md)| Comma-separated list of identifiers of requested labels | 

### Return type

[**LabelList**](LabelList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

