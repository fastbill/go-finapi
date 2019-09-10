# \PaymentsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayments**](PaymentsApi.md#GetPayments) | **Get** /api/v1/payments | Get payments


# **GetPayments**
> PageablePaymentResources GetPayments(ctx, optional)
Get payments

Get payments of the user that is authorized by the access_token. <p>Note: For requesting / executing payments, please refer to the 'Accounts' section of the API.</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetPaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPaymentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []int64**](int64.md)| A comma-separated list of payment identifiers. If specified, then only payments whose identifier is matching any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **accountIds** | [**optional.Interface of []int64**](int64.md)| A comma-separated list of account identifiers. If specified, then only payments that relate to the given account(s) will be regarded. The maximum number of identifiers is 1000. | 
 **minAmount** | **optional.Float32**| If specified, then only those payments are regarded whose (absolute) total amount is equal or greater than the given amount will be regarded. The value must be a positive (absolute) amount. | 
 **maxAmount** | **optional.Float32**| If specified, then only those payments are regarded whose (absolute) total amount is equal or less than the given amount will be regarded. Value must be a positive (absolute) amount. | 
 **page** | **optional.Int32**| Result page that you want to retrieve | [default to 1]
 **perPage** | **optional.Int32**| Maximum number of records per page. Can be at most 500. NOTE: Due to its validation and visualization, the swagger frontend might show very low performance, or even crashes, when a service responds with a lot of data. It is recommended to use a HTTP client like Postman or DHC instead of our swagger frontend for service calls with large page sizes. | [default to 20]
 **order** | [**optional.Interface of []string**](string.md)| Determines the order of the results. You can use the following fields for ordering the response: &#39;id&#39;, &#39;amount&#39;, &#39;requestDate&#39; and &#39;executionDate&#39;. The default order for all services is &#39;id,asc&#39;. | 

### Return type

[**PageablePaymentResources**](PageablePaymentResources.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

