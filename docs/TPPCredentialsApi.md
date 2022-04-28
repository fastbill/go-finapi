# \TPPCredentialsApi

All URIs are relative to *https://sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTppCredential**](TPPCredentialsApi.md#CreateTppCredential) | **Post** /api/v1/tppCredentials | Upload TPP credentials
[**DeleteTppCredential**](TPPCredentialsApi.md#DeleteTppCredential) | **Delete** /api/v1/tppCredentials/{id} | Delete a set of TPP credentials
[**EditTppCredential**](TPPCredentialsApi.md#EditTppCredential) | **Patch** /api/v1/tppCredentials/{id} | Edit a set of TPP credentials
[**GetAllTppCredentials**](TPPCredentialsApi.md#GetAllTppCredentials) | **Get** /api/v1/tppCredentials | Get all TPP credentials
[**GetAndSearchTppAuthenticationGroups**](TPPCredentialsApi.md#GetAndSearchTppAuthenticationGroups) | **Get** /api/v1/tppCredentials/tppAuthenticationGroups | Get all TPP Authentication Groups
[**GetTppCredential**](TPPCredentialsApi.md#GetTppCredential) | **Get** /api/v1/tppCredentials/{id} | Get a set of TPP credentials



## CreateTppCredential

> TppCredentials CreateTppCredential(ctx).TppCredentialsParams(tppCredentialsParams).XRequestId(xRequestId).Execute()

Upload TPP credentials



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
    tppCredentialsParams := *openapiclient.NewTppCredentialsParams(int64(1), "Demo bank TPP credentials") // TppCredentialsParams | Parameters of a new set of TPP credentials
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TPPCredentialsApi.CreateTppCredential(context.Background()).TppCredentialsParams(tppCredentialsParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TPPCredentialsApi.CreateTppCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTppCredential`: TppCredentials
    fmt.Fprintf(os.Stdout, "Response from `TPPCredentialsApi.CreateTppCredential`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTppCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tppCredentialsParams** | [**TppCredentialsParams**](TppCredentialsParams.md) | Parameters of a new set of TPP credentials | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**TppCredentials**](TppCredentials.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTppCredential

> DeleteTppCredential(ctx, id).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Delete a set of TPP credentials



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
    id := int64(789) // int64 | Id of the TPP credentials to delete
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TPPCredentialsApi.DeleteTppCredential(context.Background(), id).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TPPCredentialsApi.DeleteTppCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id of the TPP credentials to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTppCredentialRequest struct via the builder pattern


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


## EditTppCredential

> TppCredentials EditTppCredential(ctx, id).EditTppCredentialParams(editTppCredentialParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()

Edit a set of TPP credentials



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
    id := int64(789) // int64 | Id of the TPP credentials to edit
    editTppCredentialParams := *openapiclient.NewEditTppCredentialParams() // EditTppCredentialParams | New TPP credentials parameters
    xHTTPMethodOverride := "xHTTPMethodOverride_example" // string | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TPPCredentialsApi.EditTppCredential(context.Background(), id).EditTppCredentialParams(editTppCredentialParams).XHTTPMethodOverride(xHTTPMethodOverride).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TPPCredentialsApi.EditTppCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditTppCredential`: TppCredentials
    fmt.Fprintf(os.Stdout, "Response from `TPPCredentialsApi.EditTppCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id of the TPP credentials to edit | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTppCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **editTppCredentialParams** | [**EditTppCredentialParams**](EditTppCredentialParams.md) | New TPP credentials parameters | 
 **xHTTPMethodOverride** | **string** | Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with this header indicating the originally intended HTTP method. POST Requests having this  header set will be treated either as PATCH or DELETE by the finAPI servers. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**TppCredentials**](TppCredentials.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllTppCredentials

> PageableTppCredentialResources GetAllTppCredentials(ctx).Search(search).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get all TPP credentials



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
    search := "search_example" // string | Returns only the TPP credentials belonging to those banks whose 'name', 'blz', or 'bic' contains the given search string (the matching works case-insensitive). Note: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must apply to a bank for it to get included into the result. (optional)
    page := int32(56) // int32 | Result page that you want to retrieve (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can order the results by 'label'. The default order for this service is 'label,asc'. The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TPPCredentialsApi.GetAllTppCredentials(context.Background()).Search(search).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TPPCredentialsApi.GetAllTppCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllTppCredentials`: PageableTppCredentialResources
    fmt.Fprintf(os.Stdout, "Response from `TPPCredentialsApi.GetAllTppCredentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllTppCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **search** | **string** | Returns only the TPP credentials belonging to those banks whose &#39;name&#39;, &#39;blz&#39;, or &#39;bic&#39; contains the given search string (the matching works case-insensitive). Note: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must apply to a bank for it to get included into the result. | 
 **page** | **int32** | Result page that you want to retrieve | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can order the results by &#39;label&#39;. The default order for this service is &#39;label,asc&#39;. The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageableTppCredentialResources**](PageableTppCredentialResources.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAndSearchTppAuthenticationGroups

> PageableTppAuthenticationGroupResources GetAndSearchTppAuthenticationGroups(ctx).Ids(ids).Name(name).BankBlz(bankBlz).BankName(bankName).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get all TPP Authentication Groups



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
    ids := []int64{int64(123)} // []int64 | A comma-separated list of TPP authentication group identifiers. If specified, then only TPP authentication groups whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. (optional)
    name := "name_example" // string | Only the tpp authentication groups with name matching the given one should appear in the result list (optional)
    bankBlz := "bankBlz_example" // string | Search by connected banks: only the banks with BLZ matching the given one should appear in the result list (optional)
    bankName := "bankName_example" // string | Search by connected banks: only the banks with name matching the given one should appear in the result list (optional)
    page := int32(56) // int32 | Result page that you want to retrieve (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can order the results by 'id' and 'name'. The default order for this service is 'id,asc'. The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TPPCredentialsApi.GetAndSearchTppAuthenticationGroups(context.Background()).Ids(ids).Name(name).BankBlz(bankBlz).BankName(bankName).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TPPCredentialsApi.GetAndSearchTppAuthenticationGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAndSearchTppAuthenticationGroups`: PageableTppAuthenticationGroupResources
    fmt.Fprintf(os.Stdout, "Response from `TPPCredentialsApi.GetAndSearchTppAuthenticationGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAndSearchTppAuthenticationGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | A comma-separated list of TPP authentication group identifiers. If specified, then only TPP authentication groups whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **name** | **string** | Only the tpp authentication groups with name matching the given one should appear in the result list | 
 **bankBlz** | **string** | Search by connected banks: only the banks with BLZ matching the given one should appear in the result list | 
 **bankName** | **string** | Search by connected banks: only the banks with name matching the given one should appear in the result list | 
 **page** | **int32** | Result page that you want to retrieve | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can order the results by &#39;id&#39; and &#39;name&#39;. The default order for this service is &#39;id,asc&#39;. The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageableTppAuthenticationGroupResources**](PageableTppAuthenticationGroupResources.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTppCredential

> TppCredentials GetTppCredential(ctx, id).XRequestId(xRequestId).Execute()

Get a set of TPP credentials



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
    id := int64(789) // int64 | Id of the requested TPP credentials
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TPPCredentialsApi.GetTppCredential(context.Background(), id).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TPPCredentialsApi.GetTppCredential``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTppCredential`: TppCredentials
    fmt.Fprintf(os.Stdout, "Response from `TPPCredentialsApi.GetTppCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Id of the requested TPP credentials | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTppCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**TppCredentials**](TppCredentials.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

