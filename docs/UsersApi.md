# \UsersApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Post** /api/v1/users | Create a new user
[**DeleteAuthorizedUser**](UsersApi.md#DeleteAuthorizedUser) | **Delete** /api/v1/users | Delete the authorized user
[**DeleteUnverifiedUser**](UsersApi.md#DeleteUnverifiedUser) | **Delete** /api/v1/users/{userId} | Delete an unverified user
[**EditAuthorizedUser**](UsersApi.md#EditAuthorizedUser) | **Patch** /api/v1/users | Edit the authorized user
[**ExecutePasswordChange**](UsersApi.md#ExecutePasswordChange) | **Post** /api/v1/users/executePasswordChange | Execute password change
[**GetAuthorizedUser**](UsersApi.md#GetAuthorizedUser) | **Get** /api/v1/users | Get the authorized user
[**GetVerificationStatus**](UsersApi.md#GetVerificationStatus) | **Get** /api/v1/users/verificationStatus | Get a user&#39;s verification status
[**RequestPasswordChange**](UsersApi.md#RequestPasswordChange) | **Post** /api/v1/users/requestPasswordChange | Request password change
[**VerifyUser**](UsersApi.md#VerifyUser) | **Post** /api/v1/users/verify/{userId} | Verify a user


# **CreateUser**
> User CreateUser(ctx, body)
Create a new user

<p>Create a new user. Must pass your global (i.e. client) access_token. </p><p>This service returns the user's password as plain text. </p><p>The automatic update of the user's bank connections is disabled by default for any new user. User identifiers are regarded case-insensitive by finAPI.</p><p>Please note that finAPI generally has a restricted set of allowed characters for input fields. You can find the allowed characters <a href = \"https://finapi.zendesk.com/hc/en-us/articles/222013148-What-symbols-are-allowed-in-finAPI-\">here</a>. If a field does not explicitly specify a set of allowed characters, then these are the characters that are allowed for the field. Some fields may specify a different set of characters, in which case this will be documented for the field (like for the 'id' field in this service).</p>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserCreateParams**](UserCreateParams.md)| User&#39;s details | 

### Return type

[**User**](User.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAuthorizedUser**
> DeleteAuthorizedUser(ctx, )
Delete the authorized user

Delete the authorized user. Must pass the user's access_token. ATTENTION: This deletes the user including all of his bank connections, accounts, balance data and transactions! THIS PROCESS CANNOT BE UNDONE! Note that a user cannot get deleted while any of his bank connections are currently busy (in the process of import, update, or transactions categorization). <p>Note: finAPI will send a notification about the deletion of the user to each of your clients that has a user synchronization callback URL set in its configuration. This also includes the client that is performing this request.</p>

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteUnverifiedUser**
> DeleteUnverifiedUser(ctx, userId)
Delete an unverified user

Delete an unverified user. Must pass your global (i.e. client) access_token.<br/><br/>Notes:<br/>&bull; Unverified users can only exist if the field 'isUserAutoVerificationEnabled' (see Client Configuration Resource) is set to 'false' (or had been false at some point in the past).<br/>&bull; finAPI will send a notification about the deletion of the user to each of your clients that has a user synchronization callback URL set in its configuration. This also includes the client that is performing this request.<br/>&bull; finAPI regards user identifiers case-insensitive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**|  | 

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EditAuthorizedUser**
> User EditAuthorizedUser(ctx, body)
Edit the authorized user

Edit the authorized user's data and settings. Must pass the user's access_token. Pass an empty string (but not null) to unset either the email or phone number. At least one field must have a non-null value in the request body. This service returns the user's password as 'XXXXX'.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserUpdateParams**](UserUpdateParams.md)| User&#39;s details | 

### Return type

[**User**](User.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExecutePasswordChange**
> ExecutePasswordChange(ctx, optional)
Execute password change

Change the password of a user. Must pass your global (i.e. client) access_token.<br/><br/>Note: When changing the password of a user, all tokens that have been handed out for that user (for whatever client) will be revoked! Also note that finAPI regards user identifiers case-insensitive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExecutePasswordChangeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExecutePasswordChangeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ExecutePasswordChangeParams**](ExecutePasswordChangeParams.md)|  | 

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAuthorizedUser**
> User GetAuthorizedUser(ctx, )
Get the authorized user

Get the authorized user's data. Must pass the user's access_token. Only the authorized user can get his data, i.e. his access_token must be used. This service returns the user's password as 'XXXXX'.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**User**](User.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetVerificationStatus**
> VerificationStatusResource GetVerificationStatus(ctx, userId)
Get a user's verification status

Get the verification status of the requested user. Must pass your global (i.e. client) access_token. Note that finAPI regards user identifiers case-insensitive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User&#39;s identifier | 

### Return type

[**VerificationStatusResource**](VerificationStatusResource.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestPasswordChange**
> PasswordChangingResource RequestPasswordChange(ctx, optional)
Request password change

Request password change for a user. Must pass your global (i.e. client) access_token. Note that finAPI regards user identifiers case-insensitive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RequestPasswordChangeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RequestPasswordChangeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RequestPasswordChangeParams**](RequestPasswordChangeParams.md)|  | 

### Return type

[**PasswordChangingResource**](PasswordChangingResource.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VerifyUser**
> VerifyUser(ctx, userId)
Verify a user

Verify a user. User verification is only required when your client does not have auto-verification enabled (see field 'isUserAutoVerificationEnabled' in Client Configuration Resource). Must pass your global (i.e. client) access_token. Note that finAPI regards user identifiers case-insensitive.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **userId** | **string**| User&#39;s identifier | 

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

