# \AuthorizationApi

All URIs are relative to *https://sandbox.finapi.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetToken**](AuthorizationApi.md#GetToken) | **Post** /oauth/token | Get tokens
[**RevokeToken**](AuthorizationApi.md#RevokeToken) | **Post** /oauth/revoke | Revoke a token



## GetToken

> AccessToken GetToken(ctx).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).XRequestId(xRequestId).RefreshToken(refreshToken).Username(username).Password(password).Execute()

Get tokens



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
    grantType := "grantType_example" // string | Determines the required type of authorization:password - authorize a user; client_credentials - authorize a client;refresh_token - refresh a user's access_token.
    clientId := "clientId_example" // string | Client identifier
    clientSecret := "clientSecret_example" // string | Client secret
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)
    refreshToken := "refreshToken_example" // string | Refresh token. Required for grant_type=refresh_token only. (optional)
    username := "username_example" // string | User identifier. Required for grant_type=password only. (optional)
    password := "password_example" // string | User password. Required for grant_type=password only. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationApi.GetToken(context.Background()).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).XRequestId(xRequestId).RefreshToken(refreshToken).Username(username).Password(password).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.GetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetToken`: AccessToken
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationApi.GetToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | Determines the required type of authorization:password - authorize a user; client_credentials - authorize a client;refresh_token - refresh a user&#39;s access_token. | 
 **clientId** | **string** | Client identifier | 
 **clientSecret** | **string** | Client secret | 
 **xRequestId** | **string** | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don&#39;t pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name &#39;X-Request-Id&#39;. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. | 
 **refreshToken** | **string** | Refresh token. Required for grant_type&#x3D;refresh_token only. | 
 **username** | **string** | User identifier. Required for grant_type&#x3D;password only. | 
 **password** | **string** | User password. Required for grant_type&#x3D;password only. | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevokeToken

> RevokeToken(ctx).Token(token).TokenTypeHint(tokenTypeHint).XRequestId(xRequestId).Execute()

Revoke a token



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
    token := "token_example" // string | The token that the client wants to get revoked
    tokenTypeHint := "tokenTypeHint_example" // string | A hint about the type of the token submitted for revocation (optional)
    xRequestId := "xRequestId_example" // string | With any API call, you can pass a request ID. The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. If you don't pass a request ID for a call, finAPI will generate a random ID internally. The request ID is always returned back in the response of a service, as a header with name 'X-Request-Id'. We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuthorizationApi.RevokeToken(context.Background()).Token(token).TokenTypeHint(tokenTypeHint).XRequestId(xRequestId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationApi.RevokeToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | The token that the client wants to get revoked | 
 **tokenTypeHint** | **string** | A hint about the type of the token submitted for revocation | 
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

