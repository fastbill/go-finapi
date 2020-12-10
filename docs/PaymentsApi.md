# \PaymentsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectDebit**](PaymentsApi.md#CreateDirectDebit) | **Post** /api/v1/payments/directDebits | Create direct debit
[**CreateMoneyTransfer**](PaymentsApi.md#CreateMoneyTransfer) | **Post** /api/v1/payments/moneyTransfers | Create money transfer
[**GetPayments**](PaymentsApi.md#GetPayments) | **Get** /api/v1/payments | Get payments
[**SubmitPayment**](PaymentsApi.md#SubmitPayment) | **Post** /api/v1/payments/submit | Submit payment


# **CreateDirectDebit**
> Payment CreateDirectDebit(ctx, body)
Create direct debit

Create a payment for a single or collective direct debit order.<br/>Note that this service only creates a payment resource in finAPI and there is no bank communication involved.<br/>To submit the direct debit to the bank, please call the submitPayment service afterwards.<br/>Existing direct debits can be retrieved by the getPayments service.<br/><br/>A collective direct debit contains more than one order in the 'directDebits' list. It is specially handled by the bank and can be booked by the bank either as a single booking for each order or as a single cumulative booking. The preferred booking type can be given with the 'singleBooking' flag, but it is not guaranteed each bank will regard this flag.<br/><br/>Note: Prior creation of the payment resource is also necessary if you are using finAPI's web form flow.<br/><br/>Note that this service only works when your client has payments enabled (see client configuration).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateDirectDebitParams**](CreateDirectDebitParams.md)| Create direct debit parameters | 

### Return type

[**Payment**](Payment.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateMoneyTransfer**
> Payment CreateMoneyTransfer(ctx, body)
Create money transfer

Create a payment for a single or collective money transfer order.<br/>Note that this service only creates a payment resource in finAPI and there is no bank communication involved.<br/>To submit the money transfer to the bank, please call the submitPayment service afterwards.<br/>Existing money transfers can be retrieved by the getPayments service.<br/><br/>A collective money transfer contains more than one order in the 'moneyTransfers' list. It is specially handled by the bank and can be booked by the bank either as a single booking for each order or as a single cumulative booking. The preferred booking type can be given with the 'singleBooking' flag, but it is not guaranteed each bank will regard this flag.<br/><br/>Note: Prior creation of the payment resource is also necessary if you are using finAPI's web form flow.<br/><br/>Note that this service only works when your client has payments enabled (see client configuration).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateMoneyTransferParams**](CreateMoneyTransferParams.md)| Create money transfer parameters | 

### Return type

[**Payment**](Payment.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPayments**
> PageablePaymentResources GetPayments(ctx, optional)
Get payments

Get payments of the user that is authorized by the access_token. Must pass the user's access_token.<br/>This service lists all payments created by executing one of the following services:<br/>&bull; createMoneyTransfer<br/>&bull; createDirectDebit<br/>&bull; requestSepaMoneyTransfer (deprecated)<br/>&bull; requestSepaDirectDebit (deprecated)<br/><br/>You can set optional search criteria to get only those payments returned you are interested in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PaymentsApiGetPaymentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PaymentsApiGetPaymentsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []int64**](int64.md)| A comma-separated list of payment identifiers. If specified, then only payments whose identifier is matching any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **accountIds** | [**optional.Interface of []int64**](int64.md)| A comma-separated list of account identifiers. If specified, then only payments that relate to the given account(s) will be regarded. The maximum number of identifiers is 1000. | 
 **minAmount** | **optional.Float64**| If specified, then only those payments are regarded whose (absolute) total amount is equal or greater than the given amount will be regarded. The value must be a positive (absolute) amount. | 
 **maxAmount** | **optional.Float64**| If specified, then only those payments are regarded whose (absolute) total amount is equal or less than the given amount will be regarded. Value must be a positive (absolute) amount. | 
 **status** | [**optional.Interface of []string**](string.md)| A comma-separated list of payment statuses. If provided, then only payments whose status is matching any of the given statuses will be returned. Allowed values &#39;OPEN&#39;, &#39;PENDING&#39;, &#39;SUCCESSFUL&#39;, &#39;NOT_SUCCESSFUL&#39; or &#39;DISCARDED&#39;. Example: &#39;OPEN,PENDING&#39;. | 
 **page** | **optional.Int32**| Result page that you want to retrieve | [default to 1]
 **perPage** | **optional.Int32**| Maximum number of records per page. By default it&#39;s 20. Can be at most 500. NOTE: Due to its validation and visualization, the swagger frontend might show very low performance, or even crashes, when a service responds with a lot of data. It is recommended to use a HTTP client like Postman or DHC instead of our swagger frontend for service calls with large page sizes. | [default to 20]
 **order** | [**optional.Interface of []string**](string.md)| Determines the order of the results. You can use the following fields for ordering the response: &#39;id&#39;, &#39;amount&#39;, &#39;requestDate&#39; and &#39;executionDate&#39;. The default order for all services is &#39;id,asc&#39;. | 

### Return type

[**PageablePaymentResources**](PageablePaymentResources.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitPayment**
> Payment SubmitPayment(ctx, body)
Submit payment

Submit a payment to the bank which was previously created with either the createMoneyTransfer or createDirectDebit service.<br/><br/>Before you submit the payment, please check the given bank interface supports the required payment capabilities. You can do this by calling the accounts service for the account for which you created the payment and check the supported capabilities for the given interface in the 'interfaces' list.<br/>The following capabilities must be available for the given interface to submit a payment:<br/>&bull; SEPA_MONEY_TRANSFER for a single money transfer<br/>&bull; SEPA_COLLECTIVE_MONEY_TRANSFER for a collective money transfer<br/>&bull; SEPA_BASIC_DIRECT_DEBIT for a single direct debit with directDebitType set to BASIC<br/>&bull; SEPA_BASIC_COLLECTIVE_DIRECT_DEBIT for a collective direct debit with directDebitType set to BASIC<br/>&bull; SEPA_B2B_DIRECT_DEBIT for a single direct debit with directDebitType set to B2B<br/>&bull; SEPA_B2B_COLLECTIVE_DIRECT_DEBIT for a collective direct debit with directDebitType set to B2B<br/><br/>Additionally, ‘counterpartBic’ is mandatory for all orders in the payment if the account interface does not have the capability IBAN_ONLY_SEPA_MONEY_TRANSFER for money transfers, respectively IBAN_ONLY_SEPA_DIRECT_DEBIT for direct debits.<br/><br/>If the given interface doesn't have the required capabilities for the account, the payment will be rejected.<br/><br/>Usually banks require a multi-step authentication to authorize the payment. In this case, and if the finAPI web form flow is not used, the service will respond with HTTP code 510 and an error object containing a multiStepAuthentication object which describes the necessary next authentication steps. You must then retry the service call, passing the same arguments plus an additional 'multiStepAuthentication' element.<br/>Please refer to the description of the HTTP 510 error code below and the documentation of the 'MultiStepAuthenticationCallback' response object for details.<br/><br/>NOTE: Depending on your license, this service may respond with HTTP code 451, containing an error message with a identifier of web form in it. In addition to that the response will also have included a 'Location' header, which contains the URL to the web form. In this case, you must forward your user to finAPI's web form. For a detailed explanation of the Web Form Flow, please refer to this article: <a href='https://finapi.zendesk.com/hc/en-us/articles/360002596391' target='_blank'>finAPI's Web Form Flow</a><br/><br/>Note that this service only works when your client has payments enabled (see client configuration).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitPaymentParams**](SubmitPaymentParams.md)| Parameters for payment submission | 

### Return type

[**Payment**](Payment.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

