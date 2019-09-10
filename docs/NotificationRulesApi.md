# \NotificationRulesApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNotificationRule**](NotificationRulesApi.md#CreateNotificationRule) | **Post** /api/v1/notificationRules | Create a new notification rule
[**DeleteAllNotificationRules**](NotificationRulesApi.md#DeleteAllNotificationRules) | **Delete** /api/v1/notificationRules | Delete all notification rules
[**DeleteNotificationRule**](NotificationRulesApi.md#DeleteNotificationRule) | **Delete** /api/v1/notificationRules/{id} | Delete a notification rule
[**GetAndSearchAllNotificationRules**](NotificationRulesApi.md#GetAndSearchAllNotificationRules) | **Get** /api/v1/notificationRules | Get and search all notification rules
[**GetNotificationRule**](NotificationRulesApi.md#GetNotificationRule) | **Get** /api/v1/notificationRules/{id} | Get a notification rule


# **CreateNotificationRule**
> NotificationRule CreateNotificationRule(ctx, body)
Create a new notification rule

Create a new notification rule for a specific user. Must pass the user's access_token.<br/><br/>Setting up notification rules for a user allows your client application to get notified about changes in the user's data, e.g. when new transactions were downloaded, an account's balance has changed, or the user's banking credentials are no longer correct. Note that currently, this feature is implemented only for finAPI's automatic batch update, i.e. notification rules are only relevant when the user has activated the automatic updates (and when the automatic batch update is activated in general for your client).<br/><br/>There are different kinds of notification rules. The kind of a rule is depicted by the 'triggerEvent'. The trigger event specifies what data you have to pass when creating a rule (specifically, the contents of the 'params' field), on which events finAPI will send notifications to your client application, as well as what data is contained in those notifications. The specifics of the different trigger events are documented in the following article on our Dev Portal: <a href='https://finapi.zendesk.com/hc/en-us/articles/232324608-How-to-create-notification-rules-and-receive-notifications' target='_blank'>How to create notification rules and receive notifications</a>

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**NotificationRuleParams**](NotificationRuleParams.md)| Notification rule parameters | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteAllNotificationRules**
> IdentifierList DeleteAllNotificationRules(ctx, )
Delete all notification rules

Delete all notification rules of the user that is authorized by the access_token. Must pass the user's access_token. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**IdentifierList**](IdentifierList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteNotificationRule**
> DeleteNotificationRule(ctx, id)
Delete a notification rule

Delete a single notification rule of the user that is authorized by the access_token. Must pass the notification rule's identifier and the user's access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Identifier of the notification rule to delete | 

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAndSearchAllNotificationRules**
> NotificationRuleList GetAndSearchAllNotificationRules(ctx, optional)
Get and search all notification rules

Get notification rules of the user that is authorized by the access_token. Must pass the user's access_token. You can set optional search criteria to get only those notification rules that you are interested in. If you do not specify any search criteria, then this service functions as a 'get all' service.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***GetAndSearchAllNotificationRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetAndSearchAllNotificationRulesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | [**optional.Interface of []int64**](int64.md)| A comma-separated list of notification rule identifiers. If specified, then only notification rules whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 1000. | 
 **triggerEvent** | **optional.String**| If specified, then only notification rules with given trigger event will be regarded. | 
 **includeDetails** | **optional.Bool**| If specified, then only notification rules that include or not include details will be regarded. | 

### Return type

[**NotificationRuleList**](NotificationRuleList.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNotificationRule**
> NotificationRule GetNotificationRule(ctx, id)
Get a notification rule

Get a single notification rule of the user that is authorized by the access_token. Must pass the notification rule's identifier and the user's access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| Identifier of requested notification rule | 

### Return type

[**NotificationRule**](NotificationRule.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

