# \AccountsApi

All URIs are relative to *https://sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccount**](AccountsApi.md#DeleteAccount) | **Delete** /api/v1/accounts/{id} | Delete an account
[**DeleteAllAccounts**](AccountsApi.md#DeleteAllAccounts) | **Delete** /api/v1/accounts | Delete all accounts
[**EditAccount**](AccountsApi.md#EditAccount) | **Patch** /api/v1/accounts/{id} | Edit an account
[**ExecuteSepaDirectDebit**](AccountsApi.md#ExecuteSepaDirectDebit) | **Post** /api/v1/accounts/executeSepaDirectDebit | Execute SEPA Direct Debit
[**ExecuteSepaMoneyTransfer**](AccountsApi.md#ExecuteSepaMoneyTransfer) | **Post** /api/v1/accounts/executeSepaMoneyTransfer | Execute SEPA Money Transfer
[**GetAccount**](AccountsApi.md#GetAccount) | **Get** /api/v1/accounts/{id} | Get an account
[**GetAndSearchAllAccounts**](AccountsApi.md#GetAndSearchAllAccounts) | **Get** /api/v1/accounts | Get and search all accounts
[**GetDailyBalances**](AccountsApi.md#GetDailyBalances) | **Get** /api/v1/accounts/dailyBalances | Get daily balances
[**GetMultipleAccounts**](AccountsApi.md#GetMultipleAccounts) | **Get** /api/v1/accounts/{ids} | Get multiple accounts
[**RequestSepaDirectDebit**](AccountsApi.md#RequestSepaDirectDebit) | **Post** /api/v1/accounts/requestSepaDirectDebit | Request SEPA Direct Debit
[**RequestSepaMoneyTransfer**](AccountsApi.md#RequestSepaMoneyTransfer) | **Post** /api/v1/accounts/requestSepaMoneyTransfer | Request SEPA Money Transfer



## DeleteAccount

> DeleteAccount(ctx, id).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Delete an account



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
    id := int64(789) // int64 | Identifier of the account to delete
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAccount(context.Background(), id).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Identifier of the account to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xHTTPMethodOverride** | **string** | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllAccounts

> IdentifierList DeleteAllAccounts(ctx).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Delete all accounts



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
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.DeleteAllAccounts(context.Background()).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.DeleteAllAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllAccounts`: IdentifierList
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.DeleteAllAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xHTTPMethodOverride** | **string** | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**IdentifierList**](IdentifierList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditAccount

> Account EditAccount(ctx, id).AccountParams(accountParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Edit an account



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
    id := int64(789) // int64 | Identifier of the account to edit
    accountParams := *openapiclient.NewAccountParams() // AccountParams | New account name and/or type and/or 'isNew' flag
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.EditAccount(context.Background(), id).AccountParams(accountParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.EditAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.EditAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Identifier of the account to edit | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountParams** | [**AccountParams**](AccountParams.md) | New account name and/or type and/or &#39;isNew&#39; flag | 
 **xHTTPMethodOverride** | **string** | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Account**](Account.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteSepaDirectDebit

> PaymentExecutionResponse ExecuteSepaDirectDebit(ctx).ExecuteSepaDirectDebitParams(executeSepaDirectDebitParams).XRequestId(xRequestId).Execute()

Execute SEPA Direct Debit



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
    executeSepaDirectDebitParams := *openapiclient.NewExecuteSepaDirectDebitParams(int64(1)) // ExecuteSepaDirectDebitParams | Parameters for the execution of a SEPA direct debit order
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ExecuteSepaDirectDebit(context.Background()).ExecuteSepaDirectDebitParams(executeSepaDirectDebitParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ExecuteSepaDirectDebit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteSepaDirectDebit`: PaymentExecutionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ExecuteSepaDirectDebit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteSepaDirectDebitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executeSepaDirectDebitParams** | [**ExecuteSepaDirectDebitParams**](ExecuteSepaDirectDebitParams.md) | Parameters for the execution of a SEPA direct debit order | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PaymentExecutionResponse**](PaymentExecutionResponse.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteSepaMoneyTransfer

> PaymentExecutionResponse ExecuteSepaMoneyTransfer(ctx).ExecuteSepaMoneyTransferParams(executeSepaMoneyTransferParams).XRequestId(xRequestId).Execute()

Execute SEPA Money Transfer



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
    executeSepaMoneyTransferParams := *openapiclient.NewExecuteSepaMoneyTransferParams(int64(1)) // ExecuteSepaMoneyTransferParams | Parameters for the execution of a SEPA money transfer order
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.ExecuteSepaMoneyTransfer(context.Background()).ExecuteSepaMoneyTransferParams(executeSepaMoneyTransferParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.ExecuteSepaMoneyTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExecuteSepaMoneyTransfer`: PaymentExecutionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.ExecuteSepaMoneyTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExecuteSepaMoneyTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **executeSepaMoneyTransferParams** | [**ExecuteSepaMoneyTransferParams**](ExecuteSepaMoneyTransferParams.md) | Parameters for the execution of a SEPA money transfer order | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PaymentExecutionResponse**](PaymentExecutionResponse.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccount

> Account GetAccount(ctx, id).XRequestId(xRequestId).Execute()

Get an account



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
    id := int64(789) // int64 | Identifier of requested account
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAccount(context.Background(), id).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccount`: Account
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Identifier of requested account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Account**](Account.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAndSearchAllAccounts

> AccountList GetAndSearchAllAccounts(ctx).Ids(ids).Search(search).AccountTypeIds(accountTypeIds).AccountTypes(accountTypes).BankConnectionIds(bankConnectionIds).MinLastSuccessfulUpdate(minLastSuccessfulUpdate).MaxLastSuccessfulUpdate(maxLastSuccessfulUpdate).MinBalance(minBalance).MaxBalance(maxBalance).XRequestId(xRequestId).Execute()

Get and search all accounts



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
    ids := []int64{int64(123)} // []int64 | A comma-separated list of account identifiers. If specified, then only accounts whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. (optional)
    search := "search_example" // string | If specified, then only those accounts will be contained in the result whose 'accountName', 'iban', 'accountNumber' or 'subAccountNumber' contains the given search string (the matching works case-insensitive). If no accounts contain the search string in any of these fields, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the searched fields for an account to get included into the result. (optional)
    accountTypeIds := []int64{int64(123)} // []int64 | THIS FIELD IS DEPRECATED AND WILL BE REMOVED. Please refer to the 'accountTypes' field instead. A comma-separated list of account type ids. If specified, then only accounts that relate to the given types will be regarded. If not specified, then all accounts will be regarded. (optional)
    accountTypes := []string{"Inner_example"} // []string | A comma-separated list of account types. If specified, then only accounts that relate to the given types will be regarded. If not specified, then all accounts will be regarded. (optional)
    bankConnectionIds := []int64{int64(123)} // []int64 | A comma-separated list of bank connection identifiers. If specified, then only accounts that relate to the given bank connections will be regarded. If not specified, then all accounts will be regarded. (optional)
    minLastSuccessfulUpdate := "minLastSuccessfulUpdate_example" // string | Lower bound for a account's last successful update date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only accounts whose 'lastSuccessfulUpdate' is equal to or later than the given date will be regarded. (optional)
    maxLastSuccessfulUpdate := "maxLastSuccessfulUpdate_example" // string | Upper bound for a account's last successful update date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only accounts whose 'lastSuccessfulUpdate' is equal to or earlier than the given date will be regarded. (optional)
    minBalance := float64(8.14) // float64 | If specified, then only accounts whose balance is equal to or greater than the given balance will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 (optional)
    maxBalance := float64(8.14) // float64 | If specified, then only accounts whose balance is equal to or less than the given balance will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetAndSearchAllAccounts(context.Background()).Ids(ids).Search(search).AccountTypeIds(accountTypeIds).AccountTypes(accountTypes).BankConnectionIds(bankConnectionIds).MinLastSuccessfulUpdate(minLastSuccessfulUpdate).MaxLastSuccessfulUpdate(maxLastSuccessfulUpdate).MinBalance(minBalance).MaxBalance(maxBalance).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetAndSearchAllAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAndSearchAllAccounts`: AccountList
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetAndSearchAllAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAndSearchAllAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | A comma-separated list of account identifiers. If specified, then only accounts whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **search** | **string** | If specified, then only those accounts will be contained in the result whose &#39;accountName&#39;, &#39;iban&#39;, &#39;accountNumber&#39; or &#39;subAccountNumber&#39; contains the given search string (the matching works case-insensitive). If no accounts contain the search string in any of these fields, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the searched fields for an account to get included into the result. | 
 **accountTypeIds** | **[]int64** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED. Please refer to the &#39;accountTypes&#39; field instead. A comma-separated list of account type ids. If specified, then only accounts that relate to the given types will be regarded. If not specified, then all accounts will be regarded. | 
 **accountTypes** | **[]string** | A comma-separated list of account types. If specified, then only accounts that relate to the given types will be regarded. If not specified, then all accounts will be regarded. | 
 **bankConnectionIds** | **[]int64** | A comma-separated list of bank connection identifiers. If specified, then only accounts that relate to the given bank connections will be regarded. If not specified, then all accounts will be regarded. | 
 **minLastSuccessfulUpdate** | **string** | Lower bound for a account&#39;s last successful update date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only accounts whose &#39;lastSuccessfulUpdate&#39; is equal to or later than the given date will be regarded. | 
 **maxLastSuccessfulUpdate** | **string** | Upper bound for a account&#39;s last successful update date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only accounts whose &#39;lastSuccessfulUpdate&#39; is equal to or earlier than the given date will be regarded. | 
 **minBalance** | **float64** | If specified, then only accounts whose balance is equal to or greater than the given balance will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 | 
 **maxBalance** | **float64** | If specified, then only accounts whose balance is equal to or less than the given balance will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**AccountList**](AccountList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDailyBalances

> DailyBalanceList GetDailyBalances(ctx).AccountIds(accountIds).StartDate(startDate).EndDate(endDate).WithProjection(withProjection).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get daily balances



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
    accountIds := []int64{int64(123)} // []int64 | A comma-separated list of (non-security) account identifiers. If no accounts are specified, all (non-security) accounts of the user are regarded. (optional)
    startDate := "startDate_example" // string | A string in the format 'YYYY-MM-DD'. Note that the requested date range [startDate..endDate] may not exceed 1 year (366 days - considering Leap Years too).If startDate is not specified, it defaults to the endDate minus one month. (optional)
    endDate := "endDate_example" // string | A string in the format 'YYYY-MM-DD'. Note that the requested date range [startDate..endDate] may not exceed 1 year (366 days - considering Leap Years too). If endDate is not specified, it defaults to today's date. (optional)
    withProjection := true // bool | Whether finAPI should project the first and last actually existing balance of an account into the past and future. When passing 'true', then the result will always contain a daily balance for every day of the entire requested date range, even for days before the first actually existing balance, resp. after the last actually existing balance. Those days will have the same balance as the day of the first actual balance, resp. last actual balance, i.e. the first/last balance will be infinitely projected into the past/the future. When passing 'false', then the result will contain daily balances only from the day on where the first actual balance exists for any of the regarded accounts, and only up to the day where the last actual balance exists for any of the regarded accounts. Note that when in this case there are no actual balances within the requested date range, then an empty array will be returned. Default value for this parameter is 'true'. (optional) (default to true)
    page := int32(56) // int32 | Result page that you want to retrieve. (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can order the results by 'date', 'balance', 'income' or 'spending'. The default order for this service is 'date,asc'. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: '/accounts/dailyBalances?order=date,desc&order=balance,asc' will return daily balances ordered by 'date' (descending), where items with the same 'date' are ordered by 'balance' (ascending). The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetDailyBalances(context.Background()).AccountIds(accountIds).StartDate(startDate).EndDate(endDate).WithProjection(withProjection).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetDailyBalances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDailyBalances`: DailyBalanceList
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetDailyBalances`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyBalancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountIds** | **[]int64** | A comma-separated list of (non-security) account identifiers. If no accounts are specified, all (non-security) accounts of the user are regarded. | 
 **startDate** | **string** | A string in the format &#39;YYYY-MM-DD&#39;. Note that the requested date range [startDate..endDate] may not exceed 1 year (366 days - considering Leap Years too).If startDate is not specified, it defaults to the endDate minus one month. | 
 **endDate** | **string** | A string in the format &#39;YYYY-MM-DD&#39;. Note that the requested date range [startDate..endDate] may not exceed 1 year (366 days - considering Leap Years too). If endDate is not specified, it defaults to today&#39;s date. | 
 **withProjection** | **bool** | Whether finAPI should project the first and last actually existing balance of an account into the past and future. When passing &#39;true&#39;, then the result will always contain a daily balance for every day of the entire requested date range, even for days before the first actually existing balance, resp. after the last actually existing balance. Those days will have the same balance as the day of the first actual balance, resp. last actual balance, i.e. the first/last balance will be infinitely projected into the past/the future. When passing &#39;false&#39;, then the result will contain daily balances only from the day on where the first actual balance exists for any of the regarded accounts, and only up to the day where the last actual balance exists for any of the regarded accounts. Note that when in this case there are no actual balances within the requested date range, then an empty array will be returned. Default value for this parameter is &#39;true&#39;. | [default to true]
 **page** | **int32** | Result page that you want to retrieve. | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can order the results by &#39;date&#39;, &#39;balance&#39;, &#39;income&#39; or &#39;spending&#39;. The default order for this service is &#39;date,asc&#39;. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: &#39;/accounts/dailyBalances?order&#x3D;date,desc&amp;order&#x3D;balance,asc&#39; will return daily balances ordered by &#39;date&#39; (descending), where items with the same &#39;date&#39; are ordered by &#39;balance&#39; (ascending). The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**DailyBalanceList**](DailyBalanceList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleAccounts

> AccountList GetMultipleAccounts(ctx, ids).XRequestId(xRequestId).Execute()

Get multiple accounts



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
    ids := []int64{int64(123)} // []int64 | Comma-separated list of identifiers of requested accounts
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.GetMultipleAccounts(context.Background(), ids).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.GetMultipleAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultipleAccounts`: AccountList
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.GetMultipleAccounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | [**[]int64**](int64.md) | Comma-separated list of identifiers of requested accounts | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**AccountList**](AccountList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestSepaDirectDebit

> DirectDebitOrderingResponse RequestSepaDirectDebit(ctx).RequestSepaDirectDebitParams(requestSepaDirectDebitParams).XRequestId(xRequestId).Execute()

Request SEPA Direct Debit



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
    requestSepaDirectDebitParams := *openapiclient.NewRequestSepaDirectDebitParams(int64(1), openapiclient.DirectDebitType("B2B"), openapiclient.DirectDebitSequenceType("OOFF"), "2018-01-01", []openapiclient.SingleDirectDebitData{*openapiclient.NewSingleDirectDebitData("Debitor", "DE13700800000061110500", float64(99.99), "1", "2018-01-01")}) // RequestSepaDirectDebitParams | Parameters for a SEPA direct debit request
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.RequestSepaDirectDebit(context.Background()).RequestSepaDirectDebitParams(requestSepaDirectDebitParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.RequestSepaDirectDebit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestSepaDirectDebit`: DirectDebitOrderingResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.RequestSepaDirectDebit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestSepaDirectDebitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestSepaDirectDebitParams** | [**RequestSepaDirectDebitParams**](RequestSepaDirectDebitParams.md) | Parameters for a SEPA direct debit request | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**DirectDebitOrderingResponse**](DirectDebitOrderingResponse.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestSepaMoneyTransfer

> MoneyTransferOrderingResponse RequestSepaMoneyTransfer(ctx).RequestSepaMoneyTransferParams(requestSepaMoneyTransferParams).XRequestId(xRequestId).Execute()

Request SEPA Money Transfer



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
    requestSepaMoneyTransferParams := *openapiclient.NewRequestSepaMoneyTransferParams(float64(99.99), int64(1)) // RequestSepaMoneyTransferParams | Parameters for a SEPA money transfer request
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountsApi.RequestSepaMoneyTransfer(context.Background()).RequestSepaMoneyTransferParams(requestSepaMoneyTransferParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountsApi.RequestSepaMoneyTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestSepaMoneyTransfer`: MoneyTransferOrderingResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountsApi.RequestSepaMoneyTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestSepaMoneyTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestSepaMoneyTransferParams** | [**RequestSepaMoneyTransferParams**](RequestSepaMoneyTransferParams.md) | Parameters for a SEPA money transfer request | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**MoneyTransferOrderingResponse**](MoneyTransferOrderingResponse.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

