# \StandingOrdersApi

All URIs are relative to *https://sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStandingOrder**](StandingOrdersApi.md#CreateStandingOrder) | **Post** /api/v1/standingOrders | Create a standing order
[**GetStandingOrders**](StandingOrdersApi.md#GetStandingOrders) | **Get** /api/v1/standingOrders | Get standing orders
[**SubmitStandingOrder**](StandingOrdersApi.md#SubmitStandingOrder) | **Post** /api/v1/standingOrders/submit | Submit standing order



## CreateStandingOrder

> StandingOrder CreateStandingOrder(ctx).CreateStandingOrderParams(createStandingOrderParams).XRequestId(xRequestId).Execute()

Create a standing order



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
    createStandingOrderParams := *openapiclient.NewCreateStandingOrderParams("Max Mustermann", "DE13700800000061110500", float64(99.99), openapiclient.Currency("AED"), "2023-01-01", openapiclient.StandingOrderFrequency("MONTHLY")) // CreateStandingOrderParams | Create standing order parameters
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StandingOrdersApi.CreateStandingOrder(context.Background()).CreateStandingOrderParams(createStandingOrderParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandingOrdersApi.CreateStandingOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStandingOrder`: StandingOrder
    fmt.Fprintf(os.Stdout, "Response from `StandingOrdersApi.CreateStandingOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStandingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createStandingOrderParams** | [**CreateStandingOrderParams**](CreateStandingOrderParams.md) | Create standing order parameters | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**StandingOrder**](StandingOrder.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStandingOrders

> PageableStandingOrderResources GetStandingOrders(ctx).Ids(ids).AccountIds(accountIds).MinAmount(minAmount).MaxAmount(maxAmount).Status(status).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get standing orders



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
    ids := []int64{int64(123)} // []int64 | A comma-separated list of standing order identifiers. If specified, then only standing orders whose identifier is matching any of the given identifiers will be regarded. The maximum number of identifiers is 1000. (optional)
    accountIds := []int64{int64(123)} // []int64 | A comma-separated list of account identifiers. If specified, then only standing orders that relate to the given account(s) will be regarded. The maximum number of identifiers is 1000. (optional)
    minAmount := float64(8.14) // float64 | If specified, then only those standing orders are regarded whose (absolute) total amount is equal or greater than the given amount will be regarded. The value must be a positive (absolute) amount. (optional)
    maxAmount := float64(8.14) // float64 | If specified, then only those standing orders are regarded whose (absolute) total amount is equal or less than the given amount will be regarded. Value must be a positive (absolute) amount. (optional)
    status := []string{"Status_example"} // []string | A comma-separated list of standing order statuses. If provided, then only standing orders whose status is matching any of the given statuses will be returned. Allowed values: 'OPEN', 'PENDING', 'SUCCESSFUL', 'NOT_SUCCESSFUL' or 'DISCARDED'. Example: 'OPEN,PENDING'. (optional)
    page := int32(56) // int32 | Result page that you want to retrieve (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can use the following fields for ordering the response: 'id', 'amount', 'requestDate', 'requestCompletionDate'. The default order for all services is 'id,asc'. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StandingOrdersApi.GetStandingOrders(context.Background()).Ids(ids).AccountIds(accountIds).MinAmount(minAmount).MaxAmount(maxAmount).Status(status).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandingOrdersApi.GetStandingOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStandingOrders`: PageableStandingOrderResources
    fmt.Fprintf(os.Stdout, "Response from `StandingOrdersApi.GetStandingOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStandingOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | A comma-separated list of standing order identifiers. If specified, then only standing orders whose identifier is matching any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **accountIds** | **[]int64** | A comma-separated list of account identifiers. If specified, then only standing orders that relate to the given account(s) will be regarded. The maximum number of identifiers is 1000. | 
 **minAmount** | **float64** | If specified, then only those standing orders are regarded whose (absolute) total amount is equal or greater than the given amount will be regarded. The value must be a positive (absolute) amount. | 
 **maxAmount** | **float64** | If specified, then only those standing orders are regarded whose (absolute) total amount is equal or less than the given amount will be regarded. Value must be a positive (absolute) amount. | 
 **status** | **[]string** | A comma-separated list of standing order statuses. If provided, then only standing orders whose status is matching any of the given statuses will be returned. Allowed values: &#39;OPEN&#39;, &#39;PENDING&#39;, &#39;SUCCESSFUL&#39;, &#39;NOT_SUCCESSFUL&#39; or &#39;DISCARDED&#39;. Example: &#39;OPEN,PENDING&#39;. | 
 **page** | **int32** | Result page that you want to retrieve | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can use the following fields for ordering the response: &#39;id&#39;, &#39;amount&#39;, &#39;requestDate&#39;, &#39;requestCompletionDate&#39;. The default order for all services is &#39;id,asc&#39;. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageableStandingOrderResources**](PageableStandingOrderResources.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitStandingOrder

> StandingOrder SubmitStandingOrder(ctx).SubmitStandingOrderParams(submitStandingOrderParams).PSUIPAddress(pSUIPAddress).PSUDeviceOS(pSUDeviceOS).PSUUserAgent(pSUUserAgent).XRequestId(xRequestId).Execute()

Submit standing order



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
    submitStandingOrderParams := *openapiclient.NewSubmitStandingOrderParams(int64(1), openapiclient.BankingInterface("WEB_SCRAPER")) // SubmitStandingOrderParams | Parameters for standing order submission
    pSUIPAddress := "pSUIPAddress_example" // string | The IP address of the user's device. This header will be forwarded to the bank if the 'XS2A' interface is used. (optional)
    pSUDeviceOS := "pSUDeviceOS_example" // string | The user's device and/or operating system identification. This header will be forwarded to the bank if the 'XS2A' interface is used. (optional)
    pSUUserAgent := "pSUUserAgent_example" // string | The user's web browser or other client device identification. This header will be forwarded to the bank if the 'XS2A' interface is used. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.StandingOrdersApi.SubmitStandingOrder(context.Background()).SubmitStandingOrderParams(submitStandingOrderParams).PSUIPAddress(pSUIPAddress).PSUDeviceOS(pSUDeviceOS).PSUUserAgent(pSUUserAgent).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StandingOrdersApi.SubmitStandingOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitStandingOrder`: StandingOrder
    fmt.Fprintf(os.Stdout, "Response from `StandingOrdersApi.SubmitStandingOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitStandingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **submitStandingOrderParams** | [**SubmitStandingOrderParams**](SubmitStandingOrderParams.md) | Parameters for standing order submission | 
 **pSUIPAddress** | **string** | The IP address of the user&#39;s device. This header will be forwarded to the bank if the &#39;XS2A&#39; interface is used. | 
 **pSUDeviceOS** | **string** | The user&#39;s device and/or operating system identification. This header will be forwarded to the bank if the &#39;XS2A&#39; interface is used. | 
 **pSUUserAgent** | **string** | The user&#39;s web browser or other client device identification. This header will be forwarded to the bank if the &#39;XS2A&#39; interface is used. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**StandingOrder**](StandingOrder.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

