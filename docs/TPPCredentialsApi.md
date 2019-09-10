# \TPPCredentialsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTppCredential**](TPPCredentialsApi.md#CreateTppCredential) | **Post** /api/v1/tppCredentials | Create new TPP credentials
[**DeleteTppCredential**](TPPCredentialsApi.md#DeleteTppCredential) | **Delete** /api/v1/tppCredentials/{id} | Delete a set of TPP credentials
[**EditTppCredential**](TPPCredentialsApi.md#EditTppCredential) | **Patch** /api/v1/tppCredentials/{id} | Edit a set of TPP credentials
[**GetAllTppCredentials**](TPPCredentialsApi.md#GetAllTppCredentials) | **Get** /api/v1/tppCredentials | Get all TPP credentials
[**GetAndSearchTppAuthenticationGroups**](TPPCredentialsApi.md#GetAndSearchTppAuthenticationGroups) | **Get** /api/v1/tppCredentials/tppAuthenticationGroups | Get all TPP Authentication Groups
[**GetTppCredential**](TPPCredentialsApi.md#GetTppCredential) | **Get** /api/v1/tppCredentials/{id} | Get a set of TPP credentials


# **CreateTppCredential**
> TppCredentials CreateTppCredential(ctx, body)
Create new TPP credentials

Upload TPP credentials for a TPP Authentication Group. Must pass the <a href='https://finapi.zendesk.com/hc/en-us/articles/115003661827-Difference-between-app-clients-and-mandator-admin-client' target='_blank'>mandator admin client</a>'s access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TppCredentialsParams**](TppCredentialsParams.md)| Parameters of a new set of TPP credentials | 

### Return type

[**TppCredentials**](TppCredentials.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTppCredential**
> DeleteTppCredential(ctx, id)
Delete a set of TPP credentials

Delete a single set of TPP credentials by its id. Must pass the <a href='https://finapi.zendesk.com/hc/en-us/articles/115003661827-Difference-between-app-clients-and-mandator-admin-client' target='_blank'>mandator admin client</a>'s access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Id of the TPP credentials to delete | 

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditTppCredential**
> TppCredentials EditTppCredential(ctx, id, body)
Edit a set of TPP credentials

Edit TPP credentials data. Must pass the <a href='https://finapi.zendesk.com/hc/en-us/articles/115003661827-Difference-between-app-clients-and-mandator-admin-client' target='_blank'>mandator admin client</a>'s access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Id of the TPP credentials to edit | 
  **body** | [**EditTppCredentialParams**](EditTppCredentialParams.md)| New TPP credentials parameters | 

### Return type

[**TppCredentials**](TppCredentials.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllTppCredentials**
> PageableTppCredentialResources GetAllTppCredentials(ctx, optional)
Get all TPP credentials

Get and search all TPP credentials. Must pass the <a href='https://finapi.zendesk.com/hc/en-us/articles/115003661827-Difference-between-app-clients-and-mandator-admin-client' target='_blank'>mandator admin client</a>'s access_token. You can set optional search criteria to get only those TPP credentials that you are interested in. If you do not specify any search criteria, then this service functions as a 'get all' service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAllTppCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAllTppCredentialsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **optional.String**| Returns only the TPP credentials belonging to those banks whose &#39;name&#39;, &#39;blz&#39;, or &#39;bic&#39; contains the given search string (the matching works case-insensitive). Note: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must apply to a bank in order for it to get included into the result. | 
 **page** | **optional.Int32**| Result page that you want to retrieve | [default to 1]
 **perPage** | **optional.Int32**| Maximum number of records per page. Can be at most 500. NOTE: Due to its validation and visualization, the swagger frontend might show very low performance, or even crashes, when a service responds with a lot of data. It is recommended to use a HTTP client like Postman or DHC instead of our swagger frontend for service calls with large page sizes. | [default to 20]

### Return type

[**PageableTppCredentialResources**](PageableTppCredentialResources.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAndSearchTppAuthenticationGroups**
> PageableTppAuthenticationGroupResources GetAndSearchTppAuthenticationGroups(ctx, optional)
Get all TPP Authentication Groups

Get and search across all available TPP authentication groups. Must pass the <a href='https://finapi.zendesk.com/hc/en-us/articles/115003661827-Difference-between-app-clients-and-mandator-admin-client' target='_blank'>mandator admin client</a>'s access_token. You can set optional search criteria to get only those TPP authentication groups that you are interested in. If you do not specify any search criteria, then this service functions as a 'get all' service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAndSearchTppAuthenticationGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAndSearchTppAuthenticationGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| Only the tpp authentication groups with name matching the given one should appear in the result list | 
 **bankBlz** | **optional.String**| Search by connected banks: only the banks with BLZ matching the given one should appear in the result list | 
 **bankName** | **optional.String**| Search by connected banks: only the banks with name matching the given one should appear in the result list | 
 **page** | **optional.Int32**| Result page that you want to retrieve | [default to 1]
 **perPage** | **optional.Int32**| Maximum number of records per page. Can be at most 500. NOTE: Due to its validation and visualization, the swagger frontend might show very low performance, or even crashes, when a service responds with a lot of data. It is recommended to use a HTTP client like Postman or DHC instead of our swagger frontend for service calls with large page sizes. | [default to 20]

### Return type

[**PageableTppAuthenticationGroupResources**](PageableTppAuthenticationGroupResources.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTppCredential**
> TppCredentials GetTppCredential(ctx, id)
Get a set of TPP credentials

Get a single set of TPP credentials by its id. Must pass the <a href='https://finapi.zendesk.com/hc/en-us/articles/115003661827-Difference-between-app-clients-and-mandator-admin-client' target='_blank'>mandator admin client</a>'s access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Id of the requested TPP credentials | 

### Return type

[**TppCredentials**](TppCredentials.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

