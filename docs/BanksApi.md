# \BanksApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAndSearchAllBanks**](BanksApi.md#GetAndSearchAllBanks) | **Get** /api/v1/banks | Get and search all banks
[**GetBank**](BanksApi.md#GetBank) | **Get** /api/v1/banks/{id} | Get a bank
[**GetMultipleBanks**](BanksApi.md#GetMultipleBanks) | **Get** /api/v1/banks/{ids} | Get multiple banks


# **GetAndSearchAllBanks**
> PageableBankList GetAndSearchAllBanks(ctx, optional)
Get and search all banks

Get and search banks from finAPI's database of banks. Must pass the authorized user's access_token, or your client's access_token. You can set optional search criteria to get only those banks that you are interested in. If you do not specify any search criteria, then this service functions as a 'get all' service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAndSearchAllBanksOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAndSearchAllBanksOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []int64**](int64.md)| A comma-separated list of bank identifiers. If specified, then only banks whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **search** | **optional.String**| If specified, then only those banks will be contained in the result whose &#39;name&#39;, &#39;blz&#39;, &#39;bic&#39; or &#39;city&#39; contains the given search string (the matching works case-insensitive). If no banks contain the search string in any of the regarded fields, then the result will be an empty list. Note that you may also pass an IBAN in this field, in which case finAPI will try to find the related bank in its database and regard only this bank for the search. Also note: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must apply to a bank in order for it to get included into the result. | 
 **isSupported** | **optional.Bool**| If specified, then only supported (in case of &#39;true&#39; value) or unsupported (in case of &#39;false&#39; value) banks will be regarded. | 
 **pinsAreVolatile** | **optional.Bool**| If specified, then only those banks will be regarded that have the given value (true or false) for their &#39;pinsAreVolatile&#39; field. | 
 **supportedDataSources** | [**optional.Interface of []string**](string.md)| Comma-separated list of data sources. Possible values: WEB_SCRAPER,FINTS_SERVER. If this parameter is specified, then only those banks will be regarded in the search that support ALL of the given data sources. Note that this does NOT imply that those data sources must be the only data sources that are supported by a bank. | 
 **location** | [**optional.Interface of []string**](string.md)| Comma-separated list of two-letter country codes (ISO 3166 ALPHA-2). If set, then only those banks will be regarded in the search that are located in the specified countries. Notes: Banks which do not have a location set (i.e. international institutes) will ALWAYS be regarded in the search, independent of what you specify for this field. When you pass a country code that doesn&#39;t exist in the ISO 3166 ALPHA-2 standard, then the service will respond with 400 BAD_REQUEST. | 
 **isTestBank** | **optional.Bool**| If specified, then only those banks will be regarded that have the given value (true or false) for their &#39;isTestBank&#39; field. | 
 **page** | **optional.Int32**| Result page that you want to retrieve. | [default to 1]
 **perPage** | **optional.Int32**| Maximum number of records per page. Can be at most 500. NOTE: Due to its validation and visualization, the swagger frontend might show very low performance, or even crashes, when a service responds with a lot of data. It is recommended to use a HTTP client like Postman or DHC instead of our swagger frontend for service calls with large page sizes. | [default to 20]
 **order** | [**optional.Interface of []string**](string.md)| Determines the order of the results. You can order the results by &#39;id&#39;, &#39;name&#39;, &#39;blz&#39;, &#39;bic&#39; or &#39;popularity&#39;. The default order for all services is &#39;id,asc&#39;. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: &#39;/banks?order&#x3D;name,desc&amp;order&#x3D;id,asc&#39; will return banks ordered by &#39;name&#39; (descending), where banks with the same &#39;name&#39; are ordered by &#39;id&#39; (ascending). The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. Please note that ordering by multiple fields is not supported in our swagger frontend, but you can test this feature with any HTTP tool of your choice (e.g. postman or DHC).  | 

### Return type

[**PageableBankList**](PageableBankList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBank**
> Bank GetBank(ctx, id)
Get a bank

Get a single bank from finAPI's database of banks. You have to pass the bank's identifier, and either the authorized user's access_token, or your client's access token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Identifier of requested bank | 

### Return type

[**Bank**](Bank.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMultipleBanks**
> BankList GetMultipleBanks(ctx, ids)
Get multiple banks

Get a list of multiple banks from finAPI's database of banks. You have to pass a list of bank identifiers, and either the authorized user's access_token, or your client's access token. Note that banks whose identifiers do not exist will not be contained in the result (If this applies to all of the given identifiers, then the result will be an empty list).<br/><br/><b>WARNING</b>: This service is deprecated and will be removed at some point. If you want to get multiple banks, please instead use the service 'Get and search all banks' and pass a comma-separated list of identifiers with the parameter 'ids'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **ids** | [**[]int64**](int64.md)| Comma-separated list of identifiers of requested banks | 

### Return type

[**BankList**](BankList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

