# \WebFormsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWebForm**](WebFormsApi.md#GetWebForm) | **Get** /api/v1/webForms/{id} | Get a web form


# **GetWebForm**
> WebForm GetWebForm(ctx, id)
Get a web form

Get a web form of the user that is authorized by the access_token. Must pass the web form's identifier and the user's access_token. <br/><br/>Note that every web form resource is automatically removed from the finAPI system after 24 hours after its creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Identifier of web form | 

### Return type

[**WebForm**](WebForm.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

