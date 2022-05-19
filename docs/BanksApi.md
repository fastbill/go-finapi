# \BanksApi

All URIs are relative to *https://sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAndSearchAllBanks**](BanksApi.md#GetAndSearchAllBanks) | **Get** /api/v1/banks | Get and search all banks
[**GetBank**](BanksApi.md#GetBank) | **Get** /api/v1/banks/{id} | Get a bank
[**GetMultipleBanks**](BanksApi.md#GetMultipleBanks) | **Get** /api/v1/banks/{ids} | Get multiple banks



## GetAndSearchAllBanks

> PageableBankList GetAndSearchAllBanks(ctx).Ids(ids).Search(search).IsSupported(isSupported).PinsAreVolatile(pinsAreVolatile).SupportedDataSources(supportedDataSources).SupportedInterfaces(supportedInterfaces).Location(location).TppAuthenticationGroupIds(tppAuthenticationGroupIds).IsTestBank(isTestBank).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()

Get and search all banks



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
    ids := []int64{int64(123)} // []int64 | A comma-separated list of bank identifiers. If specified, then only banks whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. (optional)
    search := "search_example" // string | If specified, then only those banks will be contained in the result whose 'name', 'blz', 'bic' or 'city' contains the given search string (the matching works case-insensitive). If no banks contain the search string in any of the regarded fields, then the result will be an empty list. Note that you may also pass an IBAN in this field, in which case finAPI will try to detect the related bank and regard only this bank for the search (The IBAN may not contain spaces). Also note: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must apply to a bank for it to get included into the result. (optional)
    isSupported := true // bool | THIS FIELD IS DEPRECATED AND WILL BE REMOVED. Please refer to the 'supportedInterfaces' field instead. If specified, then only supported (in case of 'true' value) or unsupported (in case of 'false' value) banks will be regarded. (optional)
    pinsAreVolatile := true // bool | THIS FIELD IS DEPRECATED AND WILL BE REMOVED. If specified, then only those banks will be regarded that have the given value (true or false) for their 'pinsAreVolatile' field. (optional)
    supportedDataSources := []string{"Inner_example"} // []string | THIS FIELD IS DEPRECATED AND WILL BE REMOVED. Please refer to the 'supportedInterfaces' field instead. Comma-separated list of data sources. Possible values: WEB_SCRAPER,FINTS_SERVER. If this parameter is specified, then only those banks will be regarded in the search that support ALL of the given data sources. Note that this does NOT imply that those data sources must be the only data sources that are supported by a bank. (optional)
    supportedInterfaces := []string{"Inner_example"} // []string | Comma-separated list of bank interfaces. Possible values: FINTS_SERVER,WEB_SCRAPER,XS2A. If this parameter is specified, then all the banks that support at least one of the given interfaces will be returned. Note that this does NOT imply that those interfaces must be the only ones that are supported by a bank. (optional)
    location := []string{"Inner_example"} // []string | Comma-separated list of two-letter country codes (ISO 3166 ALPHA-2), for example: DE, AT. If set, then only those banks will be regarded in the search that are located in the specified countries. Notes: Banks which do not have a location set (i.e. international institutes) will ALWAYS be regarded in the search, independent of what you specify for this field. When you pass a country code that doesn't exist in the ISO 3166 ALPHA-2 standard, then the service will respond with 400 BAD_REQUEST. (optional)
    tppAuthenticationGroupIds := []int64{int64(123)} // []int64 | A comma-separated list of TPP authentication group identifiers. If specified, then only banks who have at least one interface belonging to one of the given groups will be regarded. The maximum number of identifiers is 1000. (optional)
    isTestBank := true // bool | If specified, then only those banks will be regarded that have the given value (true or false) for their 'isTestBank' field. (optional)
    page := int32(56) // int32 | Result page that you want to retrieve. (optional) (default to 1)
    perPage := int32(56) // int32 | Maximum number of records per page. By default it's 20. Can be at most 500. (optional) (default to 20)
    order := []string{"Inner_example"} // []string | Determines the order of the results. You can order the results by 'id', 'name', 'blz', 'bic' or 'popularity'. The default order for all services is 'id,asc'. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: '/banks?order=name,desc&order=id,asc' will return banks ordered by 'name' (descending), where banks with the same 'name' are ordered by 'id' (ascending). The general format is: 'property[,asc|desc]', with 'asc' being the default value. (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BanksApi.GetAndSearchAllBanks(context.Background()).Ids(ids).Search(search).IsSupported(isSupported).PinsAreVolatile(pinsAreVolatile).SupportedDataSources(supportedDataSources).SupportedInterfaces(supportedInterfaces).Location(location).TppAuthenticationGroupIds(tppAuthenticationGroupIds).IsTestBank(isTestBank).Page(page).PerPage(perPage).Order(order).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BanksApi.GetAndSearchAllBanks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAndSearchAllBanks`: PageableBankList
    fmt.Fprintf(os.Stdout, "Response from `BanksApi.GetAndSearchAllBanks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAndSearchAllBanksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]int64** | A comma-separated list of bank identifiers. If specified, then only banks whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **search** | **string** | If specified, then only those banks will be contained in the result whose &#39;name&#39;, &#39;blz&#39;, &#39;bic&#39; or &#39;city&#39; contains the given search string (the matching works case-insensitive). If no banks contain the search string in any of the regarded fields, then the result will be an empty list. Note that you may also pass an IBAN in this field, in which case finAPI will try to detect the related bank and regard only this bank for the search (The IBAN may not contain spaces). Also note: If the given search string consists of several terms (separated by whitespace), then ALL of these terms must apply to a bank for it to get included into the result. | 
 **isSupported** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED. Please refer to the &#39;supportedInterfaces&#39; field instead. If specified, then only supported (in case of &#39;true&#39; value) or unsupported (in case of &#39;false&#39; value) banks will be regarded. | 
 **pinsAreVolatile** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED. If specified, then only those banks will be regarded that have the given value (true or false) for their &#39;pinsAreVolatile&#39; field. | 
 **supportedDataSources** | **[]string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED. Please refer to the &#39;supportedInterfaces&#39; field instead. Comma-separated list of data sources. Possible values: WEB_SCRAPER,FINTS_SERVER. If this parameter is specified, then only those banks will be regarded in the search that support ALL of the given data sources. Note that this does NOT imply that those data sources must be the only data sources that are supported by a bank. | 
 **supportedInterfaces** | **[]string** | Comma-separated list of bank interfaces. Possible values: FINTS_SERVER,WEB_SCRAPER,XS2A. If this parameter is specified, then all the banks that support at least one of the given interfaces will be returned. Note that this does NOT imply that those interfaces must be the only ones that are supported by a bank. | 
 **location** | **[]string** | Comma-separated list of two-letter country codes (ISO 3166 ALPHA-2), for example: DE, AT. If set, then only those banks will be regarded in the search that are located in the specified countries. Notes: Banks which do not have a location set (i.e. international institutes) will ALWAYS be regarded in the search, independent of what you specify for this field. When you pass a country code that doesn&#39;t exist in the ISO 3166 ALPHA-2 standard, then the service will respond with 400 BAD_REQUEST. | 
 **tppAuthenticationGroupIds** | **[]int64** | A comma-separated list of TPP authentication group identifiers. If specified, then only banks who have at least one interface belonging to one of the given groups will be regarded. The maximum number of identifiers is 1000. | 
 **isTestBank** | **bool** | If specified, then only those banks will be regarded that have the given value (true or false) for their &#39;isTestBank&#39; field. | 
 **page** | **int32** | Result page that you want to retrieve. | [default to 1]
 **perPage** | **int32** | Maximum number of records per page. By default it&#39;s 20. Can be at most 500. | [default to 20]
 **order** | **[]string** | Determines the order of the results. You can order the results by &#39;id&#39;, &#39;name&#39;, &#39;blz&#39;, &#39;bic&#39; or &#39;popularity&#39;. The default order for all services is &#39;id,asc&#39;. You can also order by multiple properties. In that case the order of the parameters passed is important. Example: &#39;/banks?order&#x3D;name,desc&amp;order&#x3D;id,asc&#39; will return banks ordered by &#39;name&#39; (descending), where banks with the same &#39;name&#39; are ordered by &#39;id&#39; (ascending). The general format is: &#39;property[,asc|desc]&#39;, with &#39;asc&#39; being the default value. | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**PageableBankList**](PageableBankList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBank

> Bank GetBank(ctx, id).XRequestId(xRequestId).Execute()

Get a bank



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
    id := int64(789) // int64 | Identifier of requested bank
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BanksApi.GetBank(context.Background(), id).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BanksApi.GetBank``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBank`: Bank
    fmt.Fprintf(os.Stdout, "Response from `BanksApi.GetBank`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | Identifier of requested bank | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**Bank**](Bank.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMultipleBanks

> BankList GetMultipleBanks(ctx, ids).XRequestId(xRequestId).Execute()

Get multiple banks



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
    ids := []int64{int64(123)} // []int64 | Comma-separated list of identifiers of requested banks
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BanksApi.GetMultipleBanks(context.Background(), ids).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BanksApi.GetMultipleBanks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMultipleBanks`: BankList
    fmt.Fprintf(os.Stdout, "Response from `BanksApi.GetMultipleBanks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ids** | [**[]int64**](int64.md) | Comma-separated list of identifiers of requested banks | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMultipleBanksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 

### Return type

[**BankList**](BankList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth), [finapi_auth](../README.md#finapi_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

