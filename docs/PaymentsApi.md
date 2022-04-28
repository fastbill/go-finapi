# \PaymentsApi

All URIs are relative to *https://sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDirectDebit**](PaymentsApi.md#CreateDirectDebit) | **Post** /api/v1/payments/directDebits | Create direct debit
[**CreateMoneyTransfer**](PaymentsApi.md#CreateMoneyTransfer) | **Post** /api/v1/payments/moneyTransfers | Create money transfer
[**GetPayments**](PaymentsApi.md#GetPayments) | **Get** /api/v1/payments | Get payments
[**SubmitPayment**](PaymentsApi.md#SubmitPayment) | **Post** /api/v1/payments/submit | Submit payment



## CreateDirectDebit

> Payment CreateDirectDebit(ctx).CreateDirectDebitParams(createDirectDebitParams).XRequestId(xRequestId).Execute()

Create direct debit



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createDirectDebitParams := *openapiclient.NewCreateDirectDebitParams(int64(1), openapiclient.DirectDebitType("B2B"), openapiclient.DirectDebitSequenceType("OOFF"), []openapiclient.DirectDebitOrderParams{*openapiclient.NewDirectDebitOrderParams("Max Mustermann", "DE13700800000061110500", float64(99.99), "1", "2018-01-01", "DE02ABC01234567890")}, "2018-01-01") // CreateDirectDebitParams | Create direct debit parameters
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.CreateDirectDebit(context.Background()).CreateDirectDebitParams(createDirectDebitParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.CreateDirectDebit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDirectDebit`: Payment
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.CreateDirectDebit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDirectDebitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDirectDebitParams** | [**CreateDirectDebitParams**](CreateDirectDebitParams.md) | Create direct debit parameters | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Payment**](Payment.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMoneyTransfer

> Payment CreateMoneyTransfer(ctx).CreateMoneyTransferParams(createMoneyTransferParams).XRequestId(xRequestId).Execute()

Create money transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createMoneyTransferParams := *openapiclient.NewCreateMoneyTransferParams([]openapiclient.MoneyTransferOrderParams{*openapiclient.NewMoneyTransferOrderParams("Max Mustermann", "DE13700800000061110500", float64(99.99))}) // CreateMoneyTransferParams | Create money transfer parameters
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.CreateMoneyTransfer(context.Background()).CreateMoneyTransferParams(createMoneyTransferParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.CreateMoneyTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMoneyTransfer`: Payment
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.CreateMoneyTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMoneyTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMoneyTransferParams** | [**CreateMoneyTransferParams**](CreateMoneyTransferParams.md) | Create money transfer parameters | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Payment**](Payment.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPayments

> PageablePaymentResources GetPayments(ctx).Ids(ids).AccountIds(accountIds).MinAmount(minAmount).MaxAmount(maxAmount).Status(status).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get payments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ids := []int64{int64(123)} // []int64 | A comma-separated list of payment identifiers. If specified, then only payments whose identifier is matching any of the given identifiers will be regarded. The maximum number of identifiers is 1000. (optional)
    accountIds := []int64{int64(123)} // []int64 | A comma-separated list of account identifiers. If specified, then only payments that relate to the given account(s) will be regarded. The maximum number of identifiers is 1000. (optional)
    minAmount := float64(8.14) // float64 | If specified, then only those payments are regarded whose (absolute) total amount is equal or greater than the given amount will be regarded. The value must be a positive (absolute) amount. (optional)
    maxAmount := float64(8.14) // float64 | If specified, then only those payments are regarded whose (absolute) total amount is equal or less than the given amount will be regarded. Value must be a positive (absolute) amount. (optional)
    status := []string{"Inner_example"} // []string | A comma-separated list of payment statuses. If provided, then only payments whose status is matching any of the given statuses will be returned. Allowed values: 'OPEN', 'PENDING', 'SUCCESSFUL', 'NOT_SUCCESSFUL' or 'DISCARDED'. Example: 'OPEN,PENDING'. (optional)
    page := int32(56) // int32 | Result page that you want to retrieve (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can use the following fields for ordering the response: 'id', 'amount', 'requestDate', 'executionDate' and 'instructedExecutionDate'. The default order for all services is 'id,asc'. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.GetPayments(context.Background()).Ids(ids).AccountIds(accountIds).MinAmount(minAmount).MaxAmount(maxAmount).Status(status).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.GetPayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayments`: PageablePaymentResources
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.GetPayments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPaymentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | A comma-separated list of payment identifiers. If specified, then only payments whose identifier is matching any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **accountIds** | **[]int64** | A comma-separated list of account identifiers. If specified, then only payments that relate to the given account(s) will be regarded. The maximum number of identifiers is 1000. | 
 **minAmount** | **float64** | If specified, then only those payments are regarded whose (absolute) total amount is equal or greater than the given amount will be regarded. The value must be a positive (absolute) amount. | 
 **maxAmount** | **float64** | If specified, then only those payments are regarded whose (absolute) total amount is equal or less than the given amount will be regarded. Value must be a positive (absolute) amount. | 
 **status** | **[]string** | A comma-separated list of payment statuses. If provided, then only payments whose status is matching any of the given statuses will be returned. Allowed values: &#39;OPEN&#39;, &#39;PENDING&#39;, &#39;SUCCESSFUL&#39;, &#39;NOT_SUCCESSFUL&#39; or &#39;DISCARDED&#39;. Example: &#39;OPEN,PENDING&#39;. | 
 **page** | **int32** | Result page that you want to retrieve | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can use the following fields for ordering the response: &#39;id&#39;, &#39;amount&#39;, &#39;requestDate&#39;, &#39;executionDate&#39; and &#39;instructedExecutionDate&#39;. The default order for all services is &#39;id,asc&#39;. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageablePaymentResources**](PageablePaymentResources.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitPayment

> Payment SubmitPayment(ctx).SubmitPaymentParams(submitPaymentParams).PSUIPAddress(pSUIPAddress).PSUDeviceOS(pSUDeviceOS).PSUUserAgent(pSUUserAgent).XRequestId(xRequestId).Execute()

Submit payment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    submitPaymentParams := *openapiclient.NewSubmitPaymentParams(int64(1), openapiclient.BankingInterface("WEB_SCRAPER")) // SubmitPaymentParams | Parameters for payment submission
    pSUIPAddress := "pSUIPAddress_example" // string | The IP address of the user's device. This header will be forwarded to the bank if the 'XS2A' interface is used. (optional)
    pSUDeviceOS := "pSUDeviceOS_example" // string | The user's device and/or operating system identification. This header will be forwarded to the bank if the 'XS2A' interface is used. (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | The user's web browser or other client device identification. This header will be forwarded to the bank if the 'XS2A' interface is used. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PaymentsApi.SubmitPayment(context.Background()).SubmitPaymentParams(submitPaymentParams).PSUIPAddress(pSUIPAddress).PSUDeviceOS(pSUDeviceOS).PSUUserAgent(pSUUserAgent).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PaymentsApi.SubmitPayment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitPayment`: Payment
    fmt.Fprintf(os.Stdout, "Response from `PaymentsApi.SubmitPayment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitPaymentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitPaymentParams** | [**SubmitPaymentParams**](SubmitPaymentParams.md) | Parameters for payment submission | 
 **pSUIPAddress** | **string** | The IP address of the user&#39;s device. This header will be forwarded to the bank if the &#39;XS2A&#39; interface is used. | 
 **pSUDeviceOS** | **string** | The user&#39;s device and/or operating system identification. This header will be forwarded to the bank if the &#39;XS2A&#39; interface is used. | 
 **pSUUserAgent** | **string** | The user&#39;s web browser or other client device identification. This header will be forwarded to the bank if the &#39;XS2A&#39; interface is used. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Payment**](Payment.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

