# TppCredentialsParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TppAuthenticationGroupId** | **int64** | The TPP Authentication Group Id for which the credentials can be used | [default to null]
**Label** | **string** | Optional label to credentials | [default to null]
**TppClientId** | **string** | ID of the TPP accessing the ASPSP API, as provided by the ASPSP as the result of registration | [default to null]
**TppClientSecret** | **string** | Secret of the TPP accessing the ASPSP API, as provided by the ASPSP as the result of registration | [default to null]
**TppApiKey** | **string** | API Key provided by ASPSP  as the result of registration | [optional] [default to null]
**ValidFromDate** | **string** | Credentials \&quot;valid from\&quot; date in the format &#39;YYYY-MM-DD&#39;. Default is today&#39;s date | [optional] [default to null]
**ValidUntilDate** | **string** | Credentials \&quot;valid until\&quot; date in the format &#39;YYYY-MM-DD&#39;. Default is null which means \&quot;indefinite\&quot; (no limit) | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


