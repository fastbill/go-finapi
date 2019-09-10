# \ClientConfigurationApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EditClientConfiguration**](ClientConfigurationApi.md#EditClientConfiguration) | **Patch** /api/v1/clientConfiguration | Edit client configuration
[**GetClientConfiguration**](ClientConfigurationApi.md#GetClientConfiguration) | **Get** /api/v1/clientConfiguration | Get client configuration


# **EditClientConfiguration**
> ClientConfiguration EditClientConfiguration(ctx, optional)
Edit client configuration

Edit your client's configuration. Must pass your global (i.e. client) access_token.<br/><br/> <b>NOTE</b>: When token validity periods are changed, this only applies to newly requested tokens, and does not change the expiration time of already existing tokens.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EditClientConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EditClientConfigurationOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ClientConfigurationParams**](ClientConfigurationParams.md)| Client configuration parameters | 

### Return type

[**ClientConfiguration**](ClientConfiguration.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClientConfiguration**
> ClientConfiguration GetClientConfiguration(ctx, )
Get client configuration

Get your client's configuration. Must pass your global (i.e. client) access_token.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ClientConfiguration**](ClientConfiguration.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

