# \MocksAndTestsApi

All URIs are relative to *https://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCategorization**](MocksAndTestsApi.md#CheckCategorization) | **Post** /api/v1/tests/checkCategorization | Check categorization
[**MockBatchUpdate**](MocksAndTestsApi.md#MockBatchUpdate) | **Post** /api/v1/tests/mockBatchUpdate | Mock batch update


# **CheckCategorization**
> CategorizationCheckResults CheckCategorization(ctx, body)
Check categorization

This service can be used to check the categorization for a given set of transactions, without the need of having the transactions actually imported in finAPI. The result of the categorization is the same as if the transactions were actually imported (the service regards the user-specific categorization rules of the user that is authorized by the access_token). Must pass the user's access_token.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CheckCategorizationData**](CheckCategorizationData.md)| Transactions data | 

### Return type

[**CategorizationCheckResults**](CategorizationCheckResults.md)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MockBatchUpdate**
> MockBatchUpdate(ctx, body)
Mock batch update

This service can be used to mock an update of one or several bank connections by letting you simulate finAPI's communication with a bank server. More specifically, you can provide custom balances and transactions for existing accounts and finAPI will import that data into the accounts as if the data had been delivered by a real bank server during a real update. The idea of this service is to allow you to create accounts with specific data in them so that you can test your application in different scenarios.<br/><br/>You can also test your application's reception and processing of push notifications with this service, by enabling the 'triggerNotifications' flag in your request. When this flag is enabled, finAPI will send notifications to your application based on the notification rules that are set up for the user and on the data you provided in the request, the same way as it works with finAPI's real automatic batch update process.<br/><br/>Note that this service behaves mostly like calling the bank connection update service, meaning that it returns immediately after having asynchronously started the update process, and also meaning that you have to check the status of the updated bank connections and accounts to find out when the update has finished and what the result is. As you can update several bank connections at once, this service is closer to how finAPI's automatic batch updates work as it is to the manual update service though. Because of this, the result of the mocked bank connection updates will be stored in the 'lastAutoUpdate' field of the bank connections and not in the 'lastManualUpdate' field. Also, just like with the real batch update, any bank connection that you use with this service must have a PIN stored (even though it is not actually forwarded to any bank server).<br/><br/>Also note that this service may be called only when the user's automatic bank connection updates are disabled, to make sure that the mock updates cannot intervene with a real update (please see the User field 'isAutoUpdateEnabled'). Also, it is currently not possible to mock data for security accounts with this service, as you can only pass transactions, but not security positions.<br/><br/>Please be aware that you will 'mess up' the accounts when using this service, meaning that when you perform a real update of accounts that you have previously updated with this service, finAPI might detect inconsistencies in the data that exists in its database and the data that is reported by the bank server, and try to fix this with the insertion of an adjusting entry ('Zwischensaldo' transaction). Also, new real transactions might not get imported as finAPI could match them to mocked transactions. <b>THIS SERVICE IS MEANT FOR TESTING PURPOSES DURING DEVELOPMENT OF YOUR APPLICATION ONLY!</b> This is why it will work only on the sandbox or alpha environments. Calling it on the live environment will result in <b>403 Forbidden</b>.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MockBatchUpdateParams**](MockBatchUpdateParams.md)| Data for mock bank connection updates | 

### Return type

 (empty response body)

### Authorization

[finapi_auth](../README.md#finapi_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

