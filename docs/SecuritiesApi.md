# \SecuritiesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAndSearchAllSecurities**](SecuritiesApi.md#GetAndSearchAllSecurities) | **Get** /api/v1/securities | Get and search all securities
[**GetMultipleSecurities**](SecuritiesApi.md#GetMultipleSecurities) | **Get** /api/v1/securities/{ids} | Get multiple securities
[**GetSecurity**](SecuritiesApi.md#GetSecurity) | **Get** /api/v1/securities/{id} | Get a security


# **GetAndSearchAllSecurities**
> PageableSecurityList GetAndSearchAllSecurities(ctx, optional)
Get and search all securities

Get securities of the user that is authorized by the access_token. Must pass the user's access_token. You can set optional search criteria to get only those securities that you are interested in. If you do not specify any search criteria, then this service functions as a 'get all' service.<p>Note: Whenever a security account is being updated, its security positions will be internally re-created, meaning that the identifier of a security position might change over time.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAndSearchAllSecuritiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAndSearchAllSecuritiesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []int64**](int64.md)| A comma-separated list of security identifiers. If specified, then only securities whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **search** | **optional.String**| If specified, then only those securities will be contained in the result whose &#39;name&#39;, &#39;isin&#39; or &#39;wkn&#39; contains the given search string (the matching works case-insensitive). If no securities contain the search string in any of these fields, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the searched fields in order for a security to get included into the result. | 
 **accountIds** | [**optional.Interface of []int64**](int64.md)| Comma-separated list of identifiers of accounts | 
 **page** | **optional.Int32**| Result page that you want to retrieve. | [default to 1]
 **perPage** | **optional.Int32**| Maximum number of records per page. Can be at most 500. NOTE: Due to its validation and visualization, the swagger frontend might show very low performance, or even crashes, when a service responds with a lot of data. It is recommended to use a HTTP client like Postman or DHC instead of our swagger frontend for service calls with large page sizes. | [default to 20]
 **order** | [**optional.Interface of []string**](string.md)| Determines the order of the results. You can order the results by next fields: &#39;id&#39;, &#39;name&#39;, &#39;isin&#39;, &#39;wkn&#39;, &#39;quote&#39;, &#39;quantityNominal&#39;, &#39;marketValue&#39; and &#39;entryQuote&#39;. The default order for all services is &#39;id,asc&#39;. You can also order by multiple properties. In that case the order of the parameters passed is important. The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. Please note that ordering by multiple fields is not supported in our swagger frontend, but you can test this feature with any HTTP tool of your choice (e.g. postman or DHC).  | 

### Return type

[**PageableSecurityList**](PageableSecurityList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMultipleSecurities**
> SecurityList GetMultipleSecurities(ctx, ids)
Get multiple securities

Get a list of multiple securities of the user that is authorized by the access_token. Must pass the securities' identifiers and the user's access_token. Securities whose identifiers do not exist or do not relate to the authorized user will not be contained in the result (If this applies to all of the given identifiers, then the result will be an empty list). <p>Note: Whenever a security account is being updated, its security positions will be internally re-created, meaning that the identifier of a security position might change over time.</p><p>WARNING: This service is deprecated and will be removed at some point. If you want to get multiple securities, please instead use the service 'Get and search all securities' and pass a comma-separated list of identifiers as a parameter 'ids'.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]int64**](int64.md)| Comma-separated list of identifiers of requested securities | 

### Return type

[**SecurityList**](SecurityList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSecurity**
> Security GetSecurity(ctx, id)
Get a security

Get a single security for a specific user. Must pass the security's identifier and the user's access_token. <p>Note: Whenever a security account is being updated, its security positions will be internally re-created, meaning that the identifier of a security position might change over time.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Security identifier | 

### Return type

[**Security**](Security.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

