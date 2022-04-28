# \MandatorAdministrationApi

All URIs are relative to *https://sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeClientCredentials**](MandatorAdministrationApi.md#ChangeClientCredentials) | **Post** /api/v1/mandatorAdmin/changeClientCredentials | Change client credentials
[**CreateIbanRules**](MandatorAdministrationApi.md#CreateIbanRules) | **Post** /api/v1/mandatorAdmin/ibanRules | Create IBAN rules
[**CreateKeywordRules**](MandatorAdministrationApi.md#CreateKeywordRules) | **Post** /api/v1/mandatorAdmin/keywordRules | Create keyword rules
[**DeleteIbanRules**](MandatorAdministrationApi.md#DeleteIbanRules) | **Post** /api/v1/mandatorAdmin/ibanRules/delete | Delete IBAN rules
[**DeleteKeywordRules**](MandatorAdministrationApi.md#DeleteKeywordRules) | **Post** /api/v1/mandatorAdmin/keywordRules/delete | Delete keyword rules
[**DeleteUsers**](MandatorAdministrationApi.md#DeleteUsers) | **Post** /api/v1/mandatorAdmin/deleteUsers | Delete users
[**GetIbanRuleList**](MandatorAdministrationApi.md#GetIbanRuleList) | **Get** /api/v1/mandatorAdmin/ibanRules | Get IBAN rules
[**GetKeywordRuleList**](MandatorAdministrationApi.md#GetKeywordRuleList) | **Get** /api/v1/mandatorAdmin/keywordRules | Get keyword rules
[**GetUserList**](MandatorAdministrationApi.md#GetUserList) | **Get** /api/v1/mandatorAdmin/getUserList | Get user list



## ChangeClientCredentials

> ChangeClientCredentials(ctx).ChangeClientCredentialsParams(changeClientCredentialsParams).XRequestId(xRequestId).Execute()

Change client credentials



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
    changeClientCredentialsParams := *openapiclient.NewChangeClientCredentialsParams("01234567-89ab-cdef-0123-456789abcdef", "01234567-89ab-cdef-0123-456789abcdef", "01234567-89ab-cdef-0123-456789abcdef") // ChangeClientCredentialsParams | Parameters for changing client credentials
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MandatorAdministrationApi.ChangeClientCredentials(context.Background()).ChangeClientCredentialsParams(changeClientCredentialsParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MandatorAdministrationApi.ChangeClientCredentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangeClientCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **changeClientCredentialsParams** | [**ChangeClientCredentialsParams**](ChangeClientCredentialsParams.md) | Parameters for changing client credentials | 
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


## CreateIbanRules

> IbanRuleList CreateIbanRules(ctx).IbanRulesParams(ibanRulesParams).XRequestId(xRequestId).Execute()

Create IBAN rules



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
    ibanRulesParams := *openapiclient.NewIbanRulesParams([]openapiclient.IbanRuleParams{*openapiclient.NewIbanRuleParams("DE89370400440532013000", int64(378), openapiclient.CategorizationRuleDirection("Income"))}) // IbanRulesParams | IBAN rule definitions
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MandatorAdministrationApi.CreateIbanRules(context.Background()).IbanRulesParams(ibanRulesParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MandatorAdministrationApi.CreateIbanRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIbanRules`: IbanRuleList
    fmt.Fprintf(os.Stdout, "Response from `MandatorAdministrationApi.CreateIbanRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIbanRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ibanRulesParams** | [**IbanRulesParams**](IbanRulesParams.md) | IBAN rule definitions | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**IbanRuleList**](IbanRuleList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKeywordRules

> KeywordRuleList CreateKeywordRules(ctx).KeywordRulesParams(keywordRulesParams).XRequestId(xRequestId).Execute()

Create keyword rules



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
    keywordRulesParams := *openapiclient.NewKeywordRulesParams([]openapiclient.KeywordRuleParams{*openapiclient.NewKeywordRuleParams(int64(378), openapiclient.CategorizationRuleDirection("Income"), []string{"Keywords_example"})}) // KeywordRulesParams | Keyword rule definitions
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MandatorAdministrationApi.CreateKeywordRules(context.Background()).KeywordRulesParams(keywordRulesParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MandatorAdministrationApi.CreateKeywordRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateKeywordRules`: KeywordRuleList
    fmt.Fprintf(os.Stdout, "Response from `MandatorAdministrationApi.CreateKeywordRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeywordRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keywordRulesParams** | [**KeywordRulesParams**](KeywordRulesParams.md) | Keyword rule definitions | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**KeywordRuleList**](KeywordRuleList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIbanRules

> IdentifierList DeleteIbanRules(ctx).IbanRuleIdentifiersParams(ibanRuleIdentifiersParams).XRequestId(xRequestId).Execute()

Delete IBAN rules



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
    ibanRuleIdentifiersParams := *openapiclient.NewIbanRuleIdentifiersParams([]int64{int64(123)}) // IbanRuleIdentifiersParams | List of IBAN rules identifiers.The maximum number of identifiers is 100.
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MandatorAdministrationApi.DeleteIbanRules(context.Background()).IbanRuleIdentifiersParams(ibanRuleIdentifiersParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MandatorAdministrationApi.DeleteIbanRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteIbanRules`: IdentifierList
    fmt.Fprintf(os.Stdout, "Response from `MandatorAdministrationApi.DeleteIbanRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIbanRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ibanRuleIdentifiersParams** | [**IbanRuleIdentifiersParams**](IbanRuleIdentifiersParams.md) | List of IBAN rules identifiers.The maximum number of identifiers is 100. | 
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


## DeleteKeywordRules

> IdentifierList DeleteKeywordRules(ctx).KeywordRuleIdentifiersParams(keywordRuleIdentifiersParams).XRequestId(xRequestId).Execute()

Delete keyword rules



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
    keywordRuleIdentifiersParams := *openapiclient.NewKeywordRuleIdentifiersParams([]int64{int64(123)}) // KeywordRuleIdentifiersParams | List of keyword rule identifiers.The maximum number of identifiers is 100.
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MandatorAdministrationApi.DeleteKeywordRules(context.Background()).KeywordRuleIdentifiersParams(keywordRuleIdentifiersParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MandatorAdministrationApi.DeleteKeywordRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteKeywordRules`: IdentifierList
    fmt.Fprintf(os.Stdout, "Response from `MandatorAdministrationApi.DeleteKeywordRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteKeywordRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **keywordRuleIdentifiersParams** | [**KeywordRuleIdentifiersParams**](KeywordRuleIdentifiersParams.md) | List of keyword rule identifiers.The maximum number of identifiers is 100. | 
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


## DeleteUsers

> UserIdentifiersList DeleteUsers(ctx).UserIdentifiersParams(userIdentifiersParams).XRequestId(xRequestId).Execute()

Delete users



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
    userIdentifiersParams := *openapiclient.NewUserIdentifiersParams([]string{"UserIds_example"}) // UserIdentifiersParams | List of user identifiers
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MandatorAdministrationApi.DeleteUsers(context.Background()).UserIdentifiersParams(userIdentifiersParams).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MandatorAdministrationApi.DeleteUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteUsers`: UserIdentifiersList
    fmt.Fprintf(os.Stdout, "Response from `MandatorAdministrationApi.DeleteUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userIdentifiersParams** | [**UserIdentifiersParams**](UserIdentifiersParams.md) | List of user identifiers | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**UserIdentifiersList**](UserIdentifiersList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIbanRuleList

> PageableIbanRuleList GetIbanRuleList(ctx).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get IBAN rules



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
    page := int32(56) // int32 | Result page that you want to retrieve (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can order the results by 'id'. The default order for this service is 'id,asc'. The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MandatorAdministrationApi.GetIbanRuleList(context.Background()).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MandatorAdministrationApi.GetIbanRuleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIbanRuleList`: PageableIbanRuleList
    fmt.Fprintf(os.Stdout, "Response from `MandatorAdministrationApi.GetIbanRuleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIbanRuleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Result page that you want to retrieve | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can order the results by &#39;id&#39;. The default order for this service is &#39;id,asc&#39;. The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageableIbanRuleList**](PageableIbanRuleList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeywordRuleList

> PageableKeywordRuleList GetKeywordRuleList(ctx).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get keyword rules



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
    page := int32(56) // int32 | Result page that you want to retrieve (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can order the results by 'id'. The default order for this service is 'id,asc'. The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MandatorAdministrationApi.GetKeywordRuleList(context.Background()).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MandatorAdministrationApi.GetKeywordRuleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetKeywordRuleList`: PageableKeywordRuleList
    fmt.Fprintf(os.Stdout, "Response from `MandatorAdministrationApi.GetKeywordRuleList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetKeywordRuleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Result page that you want to retrieve | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can order the results by &#39;id&#39;. The default order for this service is &#39;id,asc&#39;. The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageableKeywordRuleList**](PageableKeywordRuleList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserList

> PageableUserInfoList GetUserList(ctx).MinRegistrationDate(minRegistrationDate).MaxRegistrationDate(maxRegistrationDate).MinDeletionDate(minDeletionDate).MaxDeletionDate(maxDeletionDate).MinLastActiveDate(minLastActiveDate).MaxLastActiveDate(maxLastActiveDate).IncludeMonthlyStats(includeMonthlyStats).MonthlyStatsStartDate(monthlyStatsStartDate).MonthlyStatsEndDate(monthlyStatsEndDate).MinBankConnectionCountInMonthlyStats(minBankConnectionCountInMonthlyStats).UserId(userId).IsDeleted(isDeleted).IsLocked(isLocked).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get user list



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
    minRegistrationDate := "minRegistrationDate_example" // string | Lower bound for a user's registration date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only users whose 'registrationDate' is equal to or later than the given date will be regarded. (optional)
    maxRegistrationDate := "maxRegistrationDate_example" // string | Upper bound for a user's registration date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only users whose 'registrationDate' is equal to or earlier than the given date will be regarded. (optional)
    minDeletionDate := "minDeletionDate_example" // string | Lower bound for a user's deletion date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only users whose 'deletionDate' is not null, and is equal to or later than the given date will be regarded. (optional)
    maxDeletionDate := "maxDeletionDate_example" // string | Upper bound for a user's deletion date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only users whose 'deletionDate' is null, or is equal to or earlier than the given date will be regarded. (optional)
    minLastActiveDate := "minLastActiveDate_example" // string | Lower bound for a user's last active date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only users whose 'lastActiveDate' is not null, and is equal to or later than the given date will be regarded. (optional)
    maxLastActiveDate := "maxLastActiveDate_example" // string | Upper bound for a user's last active date, in the format 'YYYY-MM-DD' (e.g. '2016-01-01'). If specified, then only users whose 'lastActiveDate' is null, or is equal to or earlier than the given date will be regarded. (optional)
    includeMonthlyStats := true // bool | Whether to include the 'monthlyStats' for the returned users. If not specified, then the field defaults to 'false'. (optional) (default to false)
    monthlyStatsStartDate := "monthlyStatsStartDate_example" // string | Minimum bound for the monthly stats (=oldest month that should be included). Must be passed in the format 'YYYY-MM'. If not specified, then the monthly stats will go back up to the first month in which the user existed (date of the user's registration). Note that this field is only regarded if 'includeMonthlyStats' = true. (optional)
    monthlyStatsEndDate := "monthlyStatsEndDate_example" // string | Maximum bound for the monthly stats (=latest month that should be included). Must be passed in the format 'YYYY-MM'. If not specified, then the monthly stats will go up to either the current month (for active users), or up to the month of deletion (for deleted users). Note that this field is only regarded if  'includeMonthlyStats' = true. (optional)
    minBankConnectionCountInMonthlyStats := int32(56) // int32 | A value of X means that the service will return only those users which had at least X bank connections imported at any time within the returned monthly stats set. This field is only regarded when 'includeMonthlyStats' is set to 'true'. The default value for this field is 0. (optional) (default to 0)
    userId := "userId_example" // string | The identifier of a user to search for. If specified, then only the user with the given id will be regarded. If no user can be found for the passed userId (because the user was deleted or his username was misspelled), then the result list will be empty. (optional)
    isDeleted := true // bool | If NOT specified, then the service will regard both active and deleted users in the search. If set to 'true', then ONLY deleted users will be regarded. If set to 'false', then ONLY active users will be regarded. (optional)
    isLocked := true // bool | If NOT specified, then the service will regard both locked and not locked users in the search. If set to 'true', then ONLY locked users will be regarded. If set to 'false', then ONLY not locked users will be regarded. (optional)
    page := int32(56) // int32 | Result page that you want to retrieve (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can order the results by 'userId'. The default order for this service is 'userId,asc'. The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MandatorAdministrationApi.GetUserList(context.Background()).MinRegistrationDate(minRegistrationDate).MaxRegistrationDate(maxRegistrationDate).MinDeletionDate(minDeletionDate).MaxDeletionDate(maxDeletionDate).MinLastActiveDate(minLastActiveDate).MaxLastActiveDate(maxLastActiveDate).IncludeMonthlyStats(includeMonthlyStats).MonthlyStatsStartDate(monthlyStatsStartDate).MonthlyStatsEndDate(monthlyStatsEndDate).MinBankConnectionCountInMonthlyStats(minBankConnectionCountInMonthlyStats).UserId(userId).IsDeleted(isDeleted).IsLocked(isLocked).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MandatorAdministrationApi.GetUserList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserList`: PageableUserInfoList
    fmt.Fprintf(os.Stdout, "Response from `MandatorAdministrationApi.GetUserList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minRegistrationDate** | **string** | Lower bound for a user&#39;s registration date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only users whose &#39;registrationDate&#39; is equal to or later than the given date will be regarded. | 
 **maxRegistrationDate** | **string** | Upper bound for a user&#39;s registration date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only users whose &#39;registrationDate&#39; is equal to or earlier than the given date will be regarded. | 
 **minDeletionDate** | **string** | Lower bound for a user&#39;s deletion date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only users whose &#39;deletionDate&#39; is not null, and is equal to or later than the given date will be regarded. | 
 **maxDeletionDate** | **string** | Upper bound for a user&#39;s deletion date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only users whose &#39;deletionDate&#39; is null, or is equal to or earlier than the given date will be regarded. | 
 **minLastActiveDate** | **string** | Lower bound for a user&#39;s last active date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only users whose &#39;lastActiveDate&#39; is not null, and is equal to or later than the given date will be regarded. | 
 **maxLastActiveDate** | **string** | Upper bound for a user&#39;s last active date, in the format &#39;YYYY-MM-DD&#39; (e.g. &#39;2016-01-01&#39;). If specified, then only users whose &#39;lastActiveDate&#39; is null, or is equal to or earlier than the given date will be regarded. | 
 **includeMonthlyStats** | **bool** | Whether to include the &#39;monthlyStats&#39; for the returned users. If not specified, then the field defaults to &#39;false&#39;. | [default to false]
 **monthlyStatsStartDate** | **string** | Minimum bound for the monthly stats (&#x3D;oldest month that should be included). Must be passed in the format &#39;YYYY-MM&#39;. If not specified, then the monthly stats will go back up to the first month in which the user existed (date of the user&#39;s registration). Note that this field is only regarded if &#39;includeMonthlyStats&#39; &#x3D; true. | 
 **monthlyStatsEndDate** | **string** | Maximum bound for the monthly stats (&#x3D;latest month that should be included). Must be passed in the format &#39;YYYY-MM&#39;. If not specified, then the monthly stats will go up to either the current month (for active users), or up to the month of deletion (for deleted users). Note that this field is only regarded if  &#39;includeMonthlyStats&#39; &#x3D; true. | 
 **minBankConnectionCountInMonthlyStats** | **int32** | A value of X means that the service will return only those users which had at least X bank connections imported at any time within the returned monthly stats set. This field is only regarded when &#39;includeMonthlyStats&#39; is set to &#39;true&#39;. The default value for this field is 0. | [default to 0]
 **userId** | **string** | The identifier of a user to search for. If specified, then only the user with the given id will be regarded. If no user can be found for the passed userId (because the user was deleted or his username was misspelled), then the result list will be empty. | 
 **isDeleted** | **bool** | If NOT specified, then the service will regard both active and deleted users in the search. If set to &#39;true&#39;, then ONLY deleted users will be regarded. If set to &#39;false&#39;, then ONLY active users will be regarded. | 
 **isLocked** | **bool** | If NOT specified, then the service will regard both locked and not locked users in the search. If set to &#39;true&#39;, then ONLY locked users will be regarded. If set to &#39;false&#39;, then ONLY not locked users will be regarded. | 
 **page** | **int32** | Result page that you want to retrieve | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can order the results by &#39;userId&#39;. The default order for this service is &#39;userId,asc&#39;. The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageableUserInfoList**](PageableUserInfoList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

