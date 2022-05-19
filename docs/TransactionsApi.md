# \TransactionsApi

All URIs are relative to *https://sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllTransactions**](TransactionsApi.md#DeleteAllTransactions) | **Delete** /api/v1/transactions | Delete all transactions
[**DeleteTransaction**](TransactionsApi.md#DeleteTransaction) | **Delete** /api/v1/transactions/{id} | Delete a transaction
[**EditMultipleTransactions**](TransactionsApi.md#EditMultipleTransactions) | **Patch** /api/v1/transactions | Edit multiple transactions
[**EditMultipleTransactionsDeprecated**](TransactionsApi.md#EditMultipleTransactionsDeprecated) | **Patch** /api/v1/transactions/{ids} | Edit multiple transactions (DEPRECATED)
[**EditTransaction**](TransactionsApi.md#EditTransaction) | **Patch** /api/v1/transactions/{id} | Edit a transaction
[**GetAndSearchAllTransactions**](TransactionsApi.md#GetAndSearchAllTransactions) | **Get** /api/v1/transactions | Get and search all transactions
[**GetMultipleTransactions**](TransactionsApi.md#GetMultipleTransactions) | **Get** /api/v1/transactions/{ids} | Get multiple transactions
[**GetTransaction**](TransactionsApi.md#GetTransaction) | **Get** /api/v1/transactions/{id} | Get a transaction
[**RestoreTransaction**](TransactionsApi.md#RestoreTransaction) | **Post** /api/v1/transactions/{id}/restore | Restore a transaction
[**SplitTransaction**](TransactionsApi.md#SplitTransaction) | **Post** /api/v1/transactions/{id}/split | Split a transaction
[**TriggerCategorization**](TransactionsApi.md#TriggerCategorization) | **Post** /api/v1/transactions/triggerCategorization | Trigger categorization



## DeleteAllTransactions

> IdentifierList DeleteAllTransactions(ctx).MaxDeletionDate(maxDeletionDate).MinImportDate(minImportDate).AccountIds(accountIds).SafeMode(safeMode).RememberDeletion(rememberDeletion).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Delete all transactions



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
    maxDeletionDate := "maxDeletionDate_example" // string | If specified, then only those transactions are being deleted whose 'finapiBookingDate' is equal to or earlier to the given date. The date may not be in future, and must be given in the format 'YYYY-MM-DD'. If not specified, then no date limitation will be in place for the deletion. (optional)
    minImportDate := "minImportDate_example" // string | If specified, then only those transactions are being deleted whose 'importDate' is later than or equal to the given date. The date may not be in future, and must be given in the format 'YYYY-MM-DD'. This is useful e.g. if a bank returns incorrect transactions and then fixes that issue. Then you could put the date when the error was first observed as 'minImportDate'. This would lead to deletion of all transactions after the issue was introduced and allow finAPI to refetch them from scratch. This only works if safeMode is set to false and 'rememberDeletion' is undefined or set to false. You also can not use this parameter alongside 'maxDeletionDate'. (optional)
    accountIds := []int64{int64(123)} // []int64 | A comma-separated list of account identifiers. If specified, then only transactions whose account's identifier is included in this list will be get deleted. The maximum number of identifiers is 1000. (optional)
    safeMode := true // bool | When passing 'true', then only those transactions are being deleted where at least one of the following holds true: <br/>1. The transaction belongs to an account of a test bank <br/>2. The transaction's 'potentialDuplicate' flag is set to TRUE<br/>3. The transaction is an adjusting entry ('Zwischensaldo' transaction) that was added by finAPI<br/> When passing 'false', then finAPI will delete transactions independent of these characteristics. The default value for this parameter is 'true'. (optional) (default to true)
    rememberDeletion := true // bool | When passing 'true', then finAPI will make sure to not re-import deleted transactions on future account updates. When 'false', then deleted transactions might be re-imported. Default value for this parameter is 'false'. (optional) (default to false)
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.DeleteAllTransactions(context.Background()).MaxDeletionDate(maxDeletionDate).MinImportDate(minImportDate).AccountIds(accountIds).SafeMode(safeMode).RememberDeletion(rememberDeletion).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.DeleteAllTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllTransactions`: IdentifierList
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.DeleteAllTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxDeletionDate** | **string** | If specified, then only those transactions are being deleted whose &#39;finapiBookingDate&#39; is equal to or earlier to the given date. The date may not be in future, and must be given in the format &#39;YYYY-MM-DD&#39;. If not specified, then no date limitation will be in place for the deletion. | 
 **minImportDate** | **string** | If specified, then only those transactions are being deleted whose &#39;importDate&#39; is later than or equal to the given date. The date may not be in future, and must be given in the format &#39;YYYY-MM-DD&#39;. This is useful e.g. if a bank returns incorrect transactions and then fixes that issue. Then you could put the date when the error was first observed as &#39;minImportDate&#39;. This would lead to deletion of all transactions after the issue was introduced and allow finAPI to refetch them from scratch. This only works if safeMode is set to false and &#39;rememberDeletion&#39; is undefined or set to false. You also can not use this parameter alongside &#39;maxDeletionDate&#39;. | 
 **accountIds** | **[]int64** | A comma-separated list of account identifiers. If specified, then only transactions whose account&#39;s identifier is included in this list will be get deleted. The maximum number of identifiers is 1000. | 
 **safeMode** | **bool** | When passing &#39;true&#39;, then only those transactions are being deleted where at least one of the following holds true: &lt;br/&gt;1. The transaction belongs to an account of a test bank &lt;br/&gt;2. The transaction&#39;s &#39;potentialDuplicate&#39; flag is set to TRUE&lt;br/&gt;3. The transaction is an adjusting entry (&#39;Zwischensaldo&#39; transaction) that was added by finAPI&lt;br/&gt; When passing &#39;false&#39;, then finAPI will delete transactions independent of these characteristics. The default value for this parameter is &#39;true&#39;. | [default to true]
 **rememberDeletion** | **bool** | When passing &#39;true&#39;, then finAPI will make sure to not re-import deleted transactions on future account updates. When &#39;false&#39;, then deleted transactions might be re-imported. Default value for this parameter is &#39;false&#39;. | [default to false]
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


## DeleteTransaction

> DeleteTransaction(ctx, id).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Delete a transaction



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
    id := int64(789) // int64 | Identifier of transaction
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.DeleteTransaction(context.Background(), id).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.DeleteTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTransactionRequest struct via the builder pattern


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


## EditMultipleTransactions

> IdentifierList EditMultipleTransactions(ctx).UpdateMultipleTransactionsParams(updateMultipleTransactionsParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Edit multiple transactions



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
    updateMultipleTransactionsParams := *openapiclient.NewUpdateMultipleTransactionsParams() // UpdateMultipleTransactionsParams | Update transactions parameters
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.EditMultipleTransactions(context.Background()).UpdateMultipleTransactionsParams(updateMultipleTransactionsParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.EditMultipleTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMultipleTransactions`: IdentifierList
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.EditMultipleTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEditMultipleTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMultipleTransactionsParams** | [**UpdateMultipleTransactionsParams**](UpdateMultipleTransactionsParams.md) | Update transactions parameters | 
 **xHTTPMethodOverride** | **string** | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**IdentifierList**](IdentifierList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditMultipleTransactionsDeprecated

> IdentifierList EditMultipleTransactionsDeprecated(ctx, ids).UpdateTransactionsParams(updateTransactionsParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Edit multiple transactions (DEPRECATED)



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
    ids := []int64{int64(123)} // []int64 | Comma-separated list of identifiers of updated transactions
    updateTransactionsParams := *openapiclient.NewUpdateTransactionsParams() // UpdateTransactionsParams | Update transactions parameters
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.EditMultipleTransactionsDeprecated(context.Background(), ids).UpdateTransactionsParams(updateTransactionsParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.EditMultipleTransactionsDeprecated``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditMultipleTransactionsDeprecated`: IdentifierList
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.EditMultipleTransactionsDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | [**[]int64**](int64.md) | Comma-separated list of identifiers of updated transactions | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditMultipleTransactionsDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransactionsParams** | [**UpdateTransactionsParams**](UpdateTransactionsParams.md) | Update transactions parameters | 
 **xHTTPMethodOverride** | **string** | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**IdentifierList**](IdentifierList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTransaction

> Transaction EditTransaction(ctx, id).UpdateTransactionsParams(updateTransactionsParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Edit a transaction



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
    id := int64(789) // int64 | Identifier of transaction
    updateTransactionsParams := *openapiclient.NewUpdateTransactionsParams() // UpdateTransactionsParams | Update transactions parameters
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.EditTransaction(context.Background(), id).UpdateTransactionsParams(updateTransactionsParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.EditTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.EditTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTransactionsParams** | [**UpdateTransactionsParams**](UpdateTransactionsParams.md) | Update transactions parameters | 
 **xHTTPMethodOverride** | **string** | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAndSearchAllTransactions

> PageableTransactionList GetAndSearchAllTransactions(ctx).View(view).Ids(ids).Search(search).Counterpart(counterpart).Purpose(purpose).AccountIds(accountIds).MinBankBookingDate(minBankBookingDate).MaxBankBookingDate(maxBankBookingDate).MinFinapiBookingDate(minFinapiBookingDate).MaxFinapiBookingDate(maxFinapiBookingDate).MinAmount(minAmount).MaxAmount(maxAmount).Direction(direction).LabelIds(labelIds).CategoryIds(categoryIds).IncludeChildCategories(includeChildCategories).IsNew(isNew).IsPotentialDuplicate(isPotentialDuplicate).IsAdjustingEntry(isAdjustingEntry).MinImportDate(minImportDate).MaxImportDate(maxImportDate).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get and search all transactions



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
    view := "userView" // string | This parameter defines finAPI's logical view on the transactions when querying them: 'bankView' regards only the original transactions as they were received from the bank, without considering how the transactions might have gotten split by the user (see POST /transactions/&lt;id&gt;/split). This means that if a transaction is split into logical sub-transactions, then the service will still regard only the original transaction, and NOT the logical sub-transactions in its processing (though for convenience, the transactions will have the data of their sub-transactions included in the response). 'userView' by contrast regards the transactions as they exist for the user. For transactions that have not been split into logical sub-transactions, there is no difference to the \"bankView\". But for transaction that have been split into logical sub-transactions, the service will ONLY regard these sub-transactions, and not the originally received transaction (though for convenience, the sub-transactions will have the identifier of their original transaction included in the response).
    ids := []int64{int64(123)} // []int64 | A comma-separated list of transaction identifiers. If specified, then only transactions whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. (optional)
    search := "search_example" // string | If specified, then only those transactions will be contained in the result whose 'purpose', 'endToEndReference', or one of the counterpart fields ('counterpart', 'counterpartAccountNumber', 'counterpartIban', 'counterpartBlz', 'counterpartBic', 'mandateReference', 'customerReference', 'creditorId', or 'debitorId') contain the given search string (the matching works case-insensitive). If no transactions contain the search string in any of these fields, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the searched fields for a transaction to get included into the result. (optional)
    counterpart := "counterpart_example" // string | If specified, then only those transactions will be contained in the result whose counterpart fields ('counterpart', 'counterpartAccountNumber', 'counterpartIban', 'counterpartBlz', 'counterpartBic', 'mandateReference', 'customerReference', 'creditorId', or 'debitorId') contain the given search string (the matching works case-insensitive). If no transactions contain the search string in any of the counterpart fields, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the searched fields for a transaction to get included into the result. (optional)
    purpose := "purpose_example" // string | If specified, then only those transactions will be contained in the result whose purpose field contains the given search string (the matching works case-insensitive). If no transactions contain the search string in the purpose field, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the purpose for a transaction to get included into the result. (optional)
    accountIds := []int64{int64(123)} // []int64 | A comma-separated list of account identifiers. If specified, then only transactions that relate to the given accounts will be regarded. If not specified, then all accounts will be regarded. (optional)
    minBankBookingDate := "minBankBookingDate_example" // string | Lower bound for a transaction's booking date as returned by the bank (= original booking date), in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only transactions whose 'bankBookingDate' is equal to or later than the given date will be regarded. (optional)
    maxBankBookingDate := "maxBankBookingDate_example" // string | Upper bound for a transaction's booking date as returned by the bank (= original booking date), in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only transactions whose 'bankBookingDate' is equal to or earlier than the given date will be regarded. (optional)
    minFinapiBookingDate := "minFinapiBookingDate_example" // string | Lower bound for a transaction's booking date as set by finAPI, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). For details about the meaning of the finAPI booking date, please see the field's documentation in the service's response. (optional)
    maxFinapiBookingDate := "maxFinapiBookingDate_example" // string | Upper bound for a transaction's booking date as set by finAPI, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). For details about the meaning of the finAPI booking date, please see the field's documentation in the service's response. (optional)
    minAmount := float64(8.14) // float64 | If specified, then only transactions whose amount is equal to or greater than the given amount will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 (optional)
    maxAmount := float64(8.14) // float64 | If specified, then only transactions whose amount is equal to or less than the given amount will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 (optional)
    direction := "direction_example" // string | If specified, then only transactions with the given direction(s) will be regarded. Use 'income' for regarding only received payments (amount >= 0), 'spending' for regarding only outgoing payments (amount < 0), or 'all' to regard both directions. If not specified, the direction defaults to 'all'. (optional) (default to "all")
    labelIds := []int64{int64(123)} // []int64 | A comma-separated list of label identifiers. If specified, then only transactions that  have been marked with at least one of the given labels will be contained in the result. (optional)
    categoryIds := []int64{int64(123)} // []int64 | A comma-separated list of category identifiers. If specified, then the result will contain only transactions whose category is either one of the target categories, or - but only if the 'includeChildCategories' flag is set to 'true' - whose category is a sub-category of one of the target categories. To include transactions without any category, pass the value '0' as the category ID.<br/><br/>NOTE: If your client is restricted to certain categories (see GET /clientConfiguration, field 'categoryRestrictions'), then you may only specify categories that match your restrictions. Alternatively, you can leave this field unset, in which case finAPI will automatically populate this field with all categories that are defined in your restrictions. (optional)
    includeChildCategories := true // bool | This flag controls how the 'categoryIds' are handled. If set to 'true', then all transactions with the target categories, as well as all transactions with any of their sub-categories will be regarded. If set to 'false', then sub-categories of a category are excluded and only those transactions are regarded whose category matches exactly the target category. The default value for this flag is 'true'.<br/><br/>Note that this field has an effect independent of whether you pass the 'categoryIds' yourself, or whether the 'categoryIds' are populated automatically based on your client's category restrictions (see GET /clientConfiguration, field 'categoryRestrictions'). (optional) (default to true)
    isNew := true // bool | If specified, then only transactions that have their 'isNew' flag set to true/false will be regarded. (optional)
    isPotentialDuplicate := true // bool | If specified, then only transactions that have their 'isPotentialDuplicate' flag set to true/false will be regarded. (optional)
    isAdjustingEntry := true // bool | If specified, then only transactions that have their 'isAdjustingEntry' flag set to true/false will be regarded. (optional)
    minImportDate := "minImportDate_example" // string | Lower bound for a transaction's import date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only transactions whose 'importDate' is equal to or later than the given date will be regarded. (optional)
    maxImportDate := "maxImportDate_example" // string | Upper bound for a transaction's import date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only transactions whose 'importDate' is equal to or earlier than the given date will be regarded. (optional)
    page := int32(56) // int32 | Result page that you want to retrieve. (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can use the following fields for ordering the response: 'id', 'parentId', 'accountId', 'valueDate', 'bankBookingDate', 'finapiBookingDate', 'amount', 'purpose', 'counterpartName', 'counterpartAccountNumber', 'counterpartIban', 'counterpartBlz', 'counterpartBic', 'type', 'primanota', 'category.id', 'category.name', 'isPotentialDuplicate', 'isNew' and 'importDate'. The default order for all services is 'id,asc'. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: '/transactions?order=finapiBookingDate,desc&order=counterpartName' will return the latest transactions first. If there are more transactions on the same day, then these transactions are ordered by the counterpart name (ascending). The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.GetAndSearchAllTransactions(context.Background()).View(view).Ids(ids).Search(search).Counterpart(counterpart).Purpose(purpose).AccountIds(accountIds).MinBankBookingDate(minBankBookingDate).MaxBankBookingDate(maxBankBookingDate).MinFinapiBookingDate(minFinapiBookingDate).MaxFinapiBookingDate(maxFinapiBookingDate).MinAmount(minAmount).MaxAmount(maxAmount).Direction(direction).LabelIds(labelIds).CategoryIds(categoryIds).IncludeChildCategories(includeChildCategories).IsNew(isNew).IsPotentialDuplicate(isPotentialDuplicate).IsAdjustingEntry(isAdjustingEntry).MinImportDate(minImportDate).MaxImportDate(maxImportDate).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetAndSearchAllTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAndSearchAllTransactions`: PageableTransactionList
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetAndSearchAllTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAndSearchAllTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **view** | **string** | This parameter defines finAPI&#39;s logical view on the transactions when querying them: &#39;bankView&#39; regards only the original transactions as they were received from the bank, without considering how the transactions might have gotten split by the user (see POST /transactions/&amp;lt;id&amp;gt;/split). This means that if a transaction is split into logical sub-transactions, then the service will still regard only the original transaction, and NOT the logical sub-transactions in its processing (though for convenience, the transactions will have the data of their sub-transactions included in the response). &#39;userView&#39; by contrast regards the transactions as they exist for the user. For transactions that have not been split into logical sub-transactions, there is no difference to the \&quot;bankView\&quot;. But for transaction that have been split into logical sub-transactions, the service will ONLY regard these sub-transactions, and not the originally received transaction (though for convenience, the sub-transactions will have the identifier of their original transaction included in the response). | 
 **ids** | **[]int64** | A comma-separated list of transaction identifiers. If specified, then only transactions whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **search** | **string** | If specified, then only those transactions will be contained in the result whose &#39;purpose&#39;, &#39;endToEndReference&#39;, or one of the counterpart fields (&#39;counterpart&#39;, &#39;counterpartAccountNumber&#39;, &#39;counterpartIban&#39;, &#39;counterpartBlz&#39;, &#39;counterpartBic&#39;, &#39;mandateReference&#39;, &#39;customerReference&#39;, &#39;creditorId&#39;, or &#39;debitorId&#39;) contain the given search string (the matching works case-insensitive). If no transactions contain the search string in any of these fields, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the searched fields for a transaction to get included into the result. | 
 **counterpart** | **string** | If specified, then only those transactions will be contained in the result whose counterpart fields (&#39;counterpart&#39;, &#39;counterpartAccountNumber&#39;, &#39;counterpartIban&#39;, &#39;counterpartBlz&#39;, &#39;counterpartBic&#39;, &#39;mandateReference&#39;, &#39;customerReference&#39;, &#39;creditorId&#39;, or &#39;debitorId&#39;) contain the given search string (the matching works case-insensitive). If no transactions contain the search string in any of the counterpart fields, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the searched fields for a transaction to get included into the result. | 
 **purpose** | **string** | If specified, then only those transactions will be contained in the result whose purpose field contains the given search string (the matching works case-insensitive). If no transactions contain the search string in the purpose field, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the purpose for a transaction to get included into the result. | 
 **accountIds** | **[]int64** | A comma-separated list of account identifiers. If specified, then only transactions that relate to the given accounts will be regarded. If not specified, then all accounts will be regarded. | 
 **minBankBookingDate** | **string** | Lower bound for a transaction&#39;s booking date as returned by the bank (&#x3D; original booking date), in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only transactions whose &#39;bankBookingDate&#39; is equal to or later than the given date will be regarded. | 
 **maxBankBookingDate** | **string** | Upper bound for a transaction&#39;s booking date as returned by the bank (&#x3D; original booking date), in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only transactions whose &#39;bankBookingDate&#39; is equal to or earlier than the given date will be regarded. | 
 **minFinapiBookingDate** | **string** | Lower bound for a transaction&#39;s booking date as set by finAPI, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). For details about the meaning of the finAPI booking date, please see the field&#39;s documentation in the service&#39;s response. | 
 **maxFinapiBookingDate** | **string** | Upper bound for a transaction&#39;s booking date as set by finAPI, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). For details about the meaning of the finAPI booking date, please see the field&#39;s documentation in the service&#39;s response. | 
 **minAmount** | **float64** | If specified, then only transactions whose amount is equal to or greater than the given amount will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 | 
 **maxAmount** | **float64** | If specified, then only transactions whose amount is equal to or less than the given amount will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 | 
 **direction** | **string** | If specified, then only transactions with the given direction(s) will be regarded. Use &#39;income&#39; for regarding only received payments (amount &gt;&#x3D; 0), &#39;spending&#39; for regarding only outgoing payments (amount &lt; 0), or &#39;all&#39; to regard both directions. If not specified, the direction defaults to &#39;all&#39;. | [default to &quot;all&quot;]
 **labelIds** | **[]int64** | A comma-separated list of label identifiers. If specified, then only transactions that  have been marked with at least one of the given labels will be contained in the result. | 
 **categoryIds** | **[]int64** | A comma-separated list of category identifiers. If specified, then the result will contain only transactions whose category is either one of the target categories, or - but only if the &#39;includeChildCategories&#39; flag is set to &#39;true&#39; - whose category is a sub-category of one of the target categories. To include transactions without any category, pass the value &#39;0&#39; as the category ID.&lt;br/&gt;&lt;br/&gt;NOTE: If your client is restricted to certain categories (see GET /clientConfiguration, field &#39;categoryRestrictions&#39;), then you may only specify categories that match your restrictions. Alternatively, you can leave this field unset, in which case finAPI will automatically populate this field with all categories that are defined in your restrictions. | 
 **includeChildCategories** | **bool** | This flag controls how the &#39;categoryIds&#39; are handled. If set to &#39;true&#39;, then all transactions with the target categories, as well as all transactions with any of their sub-categories will be regarded. If set to &#39;false&#39;, then sub-categories of a category are excluded and only those transactions are regarded whose category matches exactly the target category. The default value for this flag is &#39;true&#39;.&lt;br/&gt;&lt;br/&gt;Note that this field has an effect independent of whether you pass the &#39;categoryIds&#39; yourself, or whether the &#39;categoryIds&#39; are populated automatically based on your client&#39;s category restrictions (see GET /clientConfiguration, field &#39;categoryRestrictions&#39;). | [default to true]
 **isNew** | **bool** | If specified, then only transactions that have their &#39;isNew&#39; flag set to true/false will be regarded. | 
 **isPotentialDuplicate** | **bool** | If specified, then only transactions that have their &#39;isPotentialDuplicate&#39; flag set to true/false will be regarded. | 
 **isAdjustingEntry** | **bool** | If specified, then only transactions that have their &#39;isAdjustingEntry&#39; flag set to true/false will be regarded. | 
 **minImportDate** | **string** | Lower bound for a transaction&#39;s import date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only transactions whose &#39;importDate&#39; is equal to or later than the given date will be regarded. | 
 **maxImportDate** | **string** | Upper bound for a transaction&#39;s import date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only transactions whose &#39;importDate&#39; is equal to or earlier than the given date will be regarded. | 
 **page** | **int32** | Result page that you want to retrieve. | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can use the following fields for ordering the response: &#39;id&#39;, &#39;parentId&#39;, &#39;accountId&#39;, &#39;valueDate&#39;, &#39;bankBookingDate&#39;, &#39;finapiBookingDate&#39;, &#39;amount&#39;, &#39;purpose&#39;, &#39;counterpartName&#39;, &#39;counterpartAccountNumber&#39;, &#39;counterpartIban&#39;, &#39;counterpartBlz&#39;, &#39;counterpartBic&#39;, &#39;type&#39;, &#39;primanota&#39;, &#39;category.id&#39;, &#39;category.name&#39;, &#39;isPotentialDuplicate&#39;, &#39;isNew&#39; and &#39;importDate&#39;. The default order for all services is &#39;id,asc&#39;. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: &#39;/transactions?order&#x3D;finapiBookingDate,desc&amp;order&#x3D;counterpartName&#39; will return the latest transactions first. If there are more transactions on the same day, then these transactions are ordered by the counterpart name (ascending). The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageableTransactionList**](PageableTransactionList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleTransactions

> TransactionList GetMultipleTransactions(ctx, ids).XRequestId(xRequestId).Execute()

Get multiple transactions



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
    ids := []int64{int64(123)} // []int64 | Comma-separated list of identifiers of requested transactions
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.GetMultipleTransactions(context.Background(), ids).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetMultipleTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultipleTransactions`: TransactionList
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetMultipleTransactions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | [**[]int64**](int64.md) | Comma-separated list of identifiers of requested transactions | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**TransactionList**](TransactionList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransaction

> Transaction GetTransaction(ctx, id).XRequestId(xRequestId).Execute()

Get a transaction



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
    id := int64(789) // int64 | Identifier of transaction
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.GetTransaction(context.Background(), id).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.GetTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.GetTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Identifier of transaction | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreTransaction

> Transaction RestoreTransaction(ctx, id).XRequestId(xRequestId).Execute()

Restore a transaction



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
    id := int64(789) // int64 | Transaction identifier
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.RestoreTransaction(context.Background(), id).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.RestoreTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestoreTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.RestoreTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Transaction identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SplitTransaction

> Transaction SplitTransaction(ctx, id).SplitTransactionsParams(splitTransactionsParams).XRequestId(xRequestId).Execute()

Split a transaction



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
    id := int64(789) // int64 | Transaction identifier
    splitTransactionsParams := *openapiclient.NewSplitTransactionsParams([]openapiclient.SubTransactionParams{*openapiclient.NewSubTransactionParams(float64(-99.99))}) // SplitTransactionsParams | Split transactions parameters
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.SplitTransaction(context.Background(), id).SplitTransactionsParams(splitTransactionsParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.SplitTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SplitTransaction`: Transaction
    fmt.Fprintf(os.Stdout, "Response from `TransactionsApi.SplitTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Transaction identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiSplitTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **splitTransactionsParams** | [**SplitTransactionsParams**](SplitTransactionsParams.md) | Split transactions parameters | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Transaction**](Transaction.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TriggerCategorization

> TriggerCategorization(ctx).TriggerCategorizationParams(triggerCategorizationParams).XRequestId(xRequestId).Execute()

Trigger categorization



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
    triggerCategorizationParams := *openapiclient.NewTriggerCategorizationParams() // TriggerCategorizationParams | Trigger categorization parameters
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransactionsApi.TriggerCategorization(context.Background()).TriggerCategorizationParams(triggerCategorizationParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransactionsApi.TriggerCategorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTriggerCategorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **triggerCategorizationParams** | [**TriggerCategorizationParams**](TriggerCategorizationParams.md) | Trigger categorization parameters | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

