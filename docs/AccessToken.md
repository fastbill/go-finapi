# AccessToken

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | Access token. Token has a length of up to 128 characters. | [default to null]
**RefreshToken** | **string** | Refresh token. Only set in case of grant_type&#x3D;&#39;password&#39;. Token has a length of up to 128 characters. | [optional] [default to null]
**ExpiresIn** | **int32** | Expiration time in seconds. A value of 0 means that the token never expires (unless it is explicitly invalidated, e.g. by revocation, or when a user gets locked). | [default to null]
**TokenType** | **string** | Token type (it&#39;s always &#39;bearer&#39;) | [default to null]
**Scope** | **string** | Requested scopes (it&#39;s always &#39;all&#39;) | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


