# \CategoriesApi

All URIs are relative to *https://sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCategory**](CategoriesApi.md#CreateCategory) | **Post** /api/v1/categories | Create a new category
[**DeleteAllCategories**](CategoriesApi.md#DeleteAllCategories) | **Delete** /api/v1/categories | Delete all categories
[**DeleteCategory**](CategoriesApi.md#DeleteCategory) | **Delete** /api/v1/categories/{id} | Delete a category
[**EditCategory**](CategoriesApi.md#EditCategory) | **Patch** /api/v1/categories/{id} | Edit a category
[**GetAndSearchAllCategories**](CategoriesApi.md#GetAndSearchAllCategories) | **Get** /api/v1/categories | Get and search all categories
[**GetCashFlows**](CategoriesApi.md#GetCashFlows) | **Get** /api/v1/categories/cashFlows | Get cash flows
[**GetCategory**](CategoriesApi.md#GetCategory) | **Get** /api/v1/categories/{id} | Get a category
[**GetMultipleCategories**](CategoriesApi.md#GetMultipleCategories) | **Get** /api/v1/categories/{ids} | Get multiple categories
[**TrainCategorization**](CategoriesApi.md#TrainCategorization) | **Post** /api/v1/categories/trainCategorization | Train categorization



## CreateCategory

> Category CreateCategory(ctx).CategoryParams(categoryParams).XRequestId(xRequestId).Execute()

Create a new category



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
    categoryParams := *openapiclient.NewCategoryParams("Sport & Fitness") // CategoryParams | Parameters of the new category
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.CreateCategory(context.Background()).CategoryParams(categoryParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.CreateCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCategory`: Category
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.CreateCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **categoryParams** | [**CategoryParams**](CategoryParams.md) | Parameters of the new category | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Category**](Category.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllCategories

> IdentifierList DeleteAllCategories(ctx).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Delete all categories



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
    resp, r, err := api_client.CategoriesApi.DeleteAllCategories(context.Background()).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.DeleteAllCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllCategories`: IdentifierList
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.DeleteAllCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllCategoriesRequest struct via the builder pattern


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


## DeleteCategory

> DeleteCategory(ctx, id).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Delete a category



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
    id := int64(789) // int64 | Category identifier
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.DeleteCategory(context.Background(), id).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.DeleteCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Category identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCategoryRequest struct via the builder pattern


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


## EditCategory

> Category EditCategory(ctx, id).EditCategoryParams(editCategoryParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Edit a category



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
    id := int64(789) // int64 | Identifier of the category to edit
    editCategoryParams := *openapiclient.NewEditCategoryParams("Sport & Fitness") // EditCategoryParams | New category name
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.EditCategory(context.Background(), id).EditCategoryParams(editCategoryParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.EditCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditCategory`: Category
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.EditCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Identifier of the category to edit | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editCategoryParams** | [**EditCategoryParams**](EditCategoryParams.md) | New category name | 
 **xHTTPMethodOverride** | **string** | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Category**](Category.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAndSearchAllCategories

> PageableCategoryList GetAndSearchAllCategories(ctx).Ids(ids).Search(search).IsCustom(isCustom).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get and search all categories



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
    ids := []int64{int64(123)} // []int64 | A comma-separated list of category identifiers. If specified, then only categories whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. (optional)
    search := "search_example" // string | If specified, then only those categories will be contained in the result whose 'name' contains the given search string (the matching works case-insensitive). If no categories contain the search string in their name, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the name for a category to get included into the result. (optional)
    isCustom := true // bool | If specified, then the result will contain only categories that are either finAPI global (in case of value 'false'), or only categories that have been created by the authorized user (in case of value 'true'). (optional)
    page := int32(56) // int32 | Result page that you want to retrieve. (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can order the results by 'id', 'name' and 'isCustom'. The default order is 'id,asc'. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: '/categories?order=isCustom,desc&order=name' will return all custom categories followed by all default categories. Both groups are ordered ascending by name. The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.GetAndSearchAllCategories(context.Background()).Ids(ids).Search(search).IsCustom(isCustom).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetAndSearchAllCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAndSearchAllCategories`: PageableCategoryList
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetAndSearchAllCategories`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAndSearchAllCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | A comma-separated list of category identifiers. If specified, then only categories whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **search** | **string** | If specified, then only those categories will be contained in the result whose &#39;name&#39; contains the given search string (the matching works case-insensitive). If no categories contain the search string in their name, then the result will be an empty list. NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the name for a category to get included into the result. | 
 **isCustom** | **bool** | If specified, then the result will contain only categories that are either finAPI global (in case of value &#39;false&#39;), or only categories that have been created by the authorized user (in case of value &#39;true&#39;). | 
 **page** | **int32** | Result page that you want to retrieve. | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can order the results by &#39;id&#39;, &#39;name&#39; and &#39;isCustom&#39;. The default order is &#39;id,asc&#39;. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: &#39;/categories?order&#x3D;isCustom,desc&amp;order&#x3D;name&#39; will return all custom categories followed by all default categories. Both groups are ordered ascending by name. The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageableCategoryList**](PageableCategoryList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCashFlows

> CashFlowList GetCashFlows(ctx).Search(search).Counterpart(counterpart).Purpose(purpose).AccountIds(accountIds).MinBankBookingDate(minBankBookingDate).MaxBankBookingDate(maxBankBookingDate).MinFinapiBookingDate(minFinapiBookingDate).MaxFinapiBookingDate(maxFinapiBookingDate).MinAmount(minAmount).MaxAmount(maxAmount).Direction(direction).LabelIds(labelIds).CategoryIds(categoryIds).IsNew(isNew).MinImportDate(minImportDate).MaxImportDate(maxImportDate).IncludeSubCashFlows(includeSubCashFlows).Order(order).XRequestId(xRequestId).Execute()

Get cash flows



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
    search := "search_example" // string | If specified, then only transactions that contain the search term in their purpose or counterpart fields will be contained in the result. Note that the search is case insensitive. (optional)
    counterpart := "counterpart_example" // string | The counterpart is the person or institution that received your payment, or that you made the payment to. If this parameter is specified, then only transactions that contain the given term in one (or more) of their counterpart fields ('counterpartName', 'counterpartAccountNumber', 'counterpartIban', 'counterpartBic' or 'counterpartBlz') will be contained in the result. Note that the search is case insensitive. (optional)
    purpose := "purpose_example" // string | If specified, then only those transactions will be contained in the result whose purpose field contains the given search string. Note that the search is case insensitive.NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the purpose for a transaction to get included into the result. (optional)
    accountIds := []int64{int64(123)} // []int64 | A comma-separated list of account identifiers. If specified, then only transactions that relate to the given accounts will be regarded. If not specified, then all accounts will be regarded. (optional)
    minBankBookingDate := "minBankBookingDate_example" // string | Lower bound for a transaction's booking date as returned by the bank (= original booking date), in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only transactions whose 'bankBookingDate' is equal to or later than the given date will be regarded. (optional)
    maxBankBookingDate := "maxBankBookingDate_example" // string | Upper bound for a transaction's booking date as returned by the bank (= original booking date), in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only transactions whose 'bankBookingDate' is equal to or earlier than the given date will be regarded. (optional)
    minFinapiBookingDate := "minFinapiBookingDate_example" // string | Lower bound for a transaction's booking date as set by finAPI, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). For details about the meaning of the finAPI booking date, please see the field's documentation in the service's response. (optional)
    maxFinapiBookingDate := "maxFinapiBookingDate_example" // string | Upper bound for a transaction's booking date as set by finAPI, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). For details about the meaning of the finAPI booking date, please see the field's documentation in the service's response. (optional)
    minAmount := float64(8.14) // float64 | If specified, then only transactions whose amount is equal to or greater than the given amount will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 (optional)
    maxAmount := float64(8.14) // float64 | If specified, then only transactions whose amount is equal to or less than the given amount will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 (optional)
    direction := "direction_example" // string | If specified, then only transactions with the given direction(s) will be regarded. Use 'income' for regarding only received payments (amount >= 0), 'spending' for regarding only outgoing payments (amount < 0), or 'all' to regard both directions. If not specified, the direction defaults to 'all'. (optional) (default to "all")
    labelIds := []int64{int64(123)} // []int64 | A comma-separated list of label identifiers. If specified, then only transactions that have been marked with at least one of the given labels will be contained in the result. (optional)
    categoryIds := []int64{int64(123)} // []int64 | If specified, then the result will contain only those cash flows that relate to the given categories. Note that the cash flow for a category may include/exclude the cash flows of its sub-categories, depending on the 'includeSubCashFlows' setting. To include the cash flow of not categorized transactions, pass the value '0' as categoryId. Note: When this parameter is NOT set, then the result will contain a cash flow for all categories that have transactions associated to them (this includes the 'null'-category for the cash flow of not categorized transactions), more precisely: transactions that fulfill the filter criteria. Categories that have no associated transactions according to the filter criteria will not appear in the result. However, when you specify this parameter, then all specified categories will have a cash flow entry in the result, even if there are no associated transactions for the category (the cash flow will have income, spending and balance all set to zero). (optional)
    isNew := true // bool | If specified, then only transactions that have their 'isNew' flag set to true/false will be regarded for the cash flow calculations. (optional)
    minImportDate := "minImportDate_example" // string | Lower bound for a transaction's import date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only transactions whose 'importDate' is equal to or later than the given date will be regarded. (optional)
    maxImportDate := "maxImportDate_example" // string | Upper bound for a transaction's import date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only transactions whose 'importDate' is equal to or earlier than the given date will be regarded. (optional)
    includeSubCashFlows := true // bool | If it is true, then the income, spending, balance and count of transactions of a main category results from all transactions that have either this (main) category or any of its subcategories assigned (of course all transactions depends from the other filtering settings); If it is false, then the income, spending, balance and count of transactions of a main category only results from the transactions that have exactly this (main) category assigned. Default value for this parameter is 'true'. (optional) (default to true)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can order the results by 'income', 'spending', 'balance', 'category.id' or 'category.name'. The default order for this service is 'category.id,asc'. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: '/cashFlows?order=income,desc&order=spending,asc&balance,desc' will return as first result the category with the highest income. If two categories have the same income, it returns the category with the highest spending first (because spending is a negative value) and so on. The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.GetCashFlows(context.Background()).Search(search).Counterpart(counterpart).Purpose(purpose).AccountIds(accountIds).MinBankBookingDate(minBankBookingDate).MaxBankBookingDate(maxBankBookingDate).MinFinapiBookingDate(minFinapiBookingDate).MaxFinapiBookingDate(maxFinapiBookingDate).MinAmount(minAmount).MaxAmount(maxAmount).Direction(direction).LabelIds(labelIds).CategoryIds(categoryIds).IsNew(isNew).MinImportDate(minImportDate).MaxImportDate(maxImportDate).IncludeSubCashFlows(includeSubCashFlows).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetCashFlows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCashFlows`: CashFlowList
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetCashFlows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCashFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | If specified, then only transactions that contain the search term in their purpose or counterpart fields will be contained in the result. Note that the search is case insensitive. | 
 **counterpart** | **string** | The counterpart is the person or institution that received your payment, or that you made the payment to. If this parameter is specified, then only transactions that contain the given term in one (or more) of their counterpart fields (&#39;counterpartName&#39;, &#39;counterpartAccountNumber&#39;, &#39;counterpartIban&#39;, &#39;counterpartBic&#39; or &#39;counterpartBlz&#39;) will be contained in the result. Note that the search is case insensitive. | 
 **purpose** | **string** | If specified, then only those transactions will be contained in the result whose purpose field contains the given search string. Note that the search is case insensitive.NOTE: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must be contained in the purpose for a transaction to get included into the result. | 
 **accountIds** | **[]int64** | A comma-separated list of account identifiers. If specified, then only transactions that relate to the given accounts will be regarded. If not specified, then all accounts will be regarded. | 
 **minBankBookingDate** | **string** | Lower bound for a transaction&#39;s booking date as returned by the bank (&#x3D; original booking date), in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only transactions whose &#39;bankBookingDate&#39; is equal to or later than the given date will be regarded. | 
 **maxBankBookingDate** | **string** | Upper bound for a transaction&#39;s booking date as returned by the bank (&#x3D; original booking date), in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only transactions whose &#39;bankBookingDate&#39; is equal to or earlier than the given date will be regarded. | 
 **minFinapiBookingDate** | **string** | Lower bound for a transaction&#39;s booking date as set by finAPI, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). For details about the meaning of the finAPI booking date, please see the field&#39;s documentation in the service&#39;s response. | 
 **maxFinapiBookingDate** | **string** | Upper bound for a transaction&#39;s booking date as set by finAPI, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). For details about the meaning of the finAPI booking date, please see the field&#39;s documentation in the service&#39;s response. | 
 **minAmount** | **float64** | If specified, then only transactions whose amount is equal to or greater than the given amount will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 | 
 **maxAmount** | **float64** | If specified, then only transactions whose amount is equal to or less than the given amount will be regarded. Can contain a positive or negative number with at most two decimal places. Examples: -300.12, or 90.95 | 
 **direction** | **string** | If specified, then only transactions with the given direction(s) will be regarded. Use &#39;income&#39; for regarding only received payments (amount &gt;&#x3D; 0), &#39;spending&#39; for regarding only outgoing payments (amount &lt; 0), or &#39;all&#39; to regard both directions. If not specified, the direction defaults to &#39;all&#39;. | [default to &quot;all&quot;]
 **labelIds** | **[]int64** | A comma-separated list of label identifiers. If specified, then only transactions that have been marked with at least one of the given labels will be contained in the result. | 
 **categoryIds** | **[]int64** | If specified, then the result will contain only those cash flows that relate to the given categories. Note that the cash flow for a category may include/exclude the cash flows of its sub-categories, depending on the &#39;includeSubCashFlows&#39; setting. To include the cash flow of not categorized transactions, pass the value &#39;0&#39; as categoryId. Note: When this parameter is NOT set, then the result will contain a cash flow for all categories that have transactions associated to them (this includes the &#39;null&#39;-category for the cash flow of not categorized transactions), more precisely: transactions that fulfill the filter criteria. Categories that have no associated transactions according to the filter criteria will not appear in the result. However, when you specify this parameter, then all specified categories will have a cash flow entry in the result, even if there are no associated transactions for the category (the cash flow will have income, spending and balance all set to zero). | 
 **isNew** | **bool** | If specified, then only transactions that have their &#39;isNew&#39; flag set to true/false will be regarded for the cash flow calculations. | 
 **minImportDate** | **string** | Lower bound for a transaction&#39;s import date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only transactions whose &#39;importDate&#39; is equal to or later than the given date will be regarded. | 
 **maxImportDate** | **string** | Upper bound for a transaction&#39;s import date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only transactions whose &#39;importDate&#39; is equal to or earlier than the given date will be regarded. | 
 **includeSubCashFlows** | **bool** | If it is true, then the income, spending, balance and count of transactions of a main category results from all transactions that have either this (main) category or any of its subcategories assigned (of course all transactions depends from the other filtering settings); If it is false, then the income, spending, balance and count of transactions of a main category only results from the transactions that have exactly this (main) category assigned. Default value for this parameter is &#39;true&#39;. | [default to true]
 **order** | **[]string** | Determines the order of the results. You can order the results by &#39;income&#39;, &#39;spending&#39;, &#39;balance&#39;, &#39;category.id&#39; or &#39;category.name&#39;. The default order for this service is &#39;category.id,asc&#39;. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: &#39;/cashFlows?order&#x3D;income,desc&amp;order&#x3D;spending,asc&amp;balance,desc&#39; will return as first result the category with the highest income. If two categories have the same income, it returns the category with the highest spending first (because spending is a negative value) and so on. The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**CashFlowList**](CashFlowList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategory

> Category GetCategory(ctx, id).XRequestId(xRequestId).Execute()

Get a category



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
    id := int64(789) // int64 | Category identifier
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.GetCategory(context.Background(), id).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategory`: Category
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Category identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Category**](Category.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleCategories

> CategoryList GetMultipleCategories(ctx, ids).XRequestId(xRequestId).Execute()

Get multiple categories



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
    ids := []int64{int64(123)} // []int64 | Comma-separated list of identifiers of requested categories
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.GetMultipleCategories(context.Background(), ids).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.GetMultipleCategories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultipleCategories`: CategoryList
    fmt.Fprintf(os.Stdout, "Response from `CategoriesApi.GetMultipleCategories`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | [**[]int64**](int64.md) | Comma-separated list of identifiers of requested categories | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleCategoriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**CategoryList**](CategoryList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrainCategorization

> TrainCategorization(ctx).TrainCategorizationData(trainCategorizationData).XRequestId(xRequestId).Execute()

Train categorization



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
    trainCategorizationData := *openapiclient.NewTrainCategorizationData([]openapiclient.TrainCategorizationTransactionData{*openapiclient.NewTrainCategorizationTransactionData(int64(1), float64(-99.99))}, int64(378)) // TrainCategorizationData | Categorization sample
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CategoriesApi.TrainCategorization(context.Background()).TrainCategorizationData(trainCategorizationData).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CategoriesApi.TrainCategorization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTrainCategorizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **trainCategorizationData** | [**TrainCategorizationData**](TrainCategorizationData.md) | Categorization sample | 
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

