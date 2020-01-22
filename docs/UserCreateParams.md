# UserCreateParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User&#39;s identifier. Max length is 36 symbols. Only the following symbols are allowed: 0-9, A-Z, a-z, -, _, ., +, @. If not specified, then a unique random value will be generated. | [optional] [default to null]
**Password** | **string** | User&#39;s password. Minimum length is 6, and maximum length is 128. If not specified, then a unique random value will be generated. | [optional] [default to null]
**Email** | **string** | User&#39;s email address. Maximum length is 320. | [optional] [default to null]
**Phone** | **string** | User&#39;s phone number. Maximum length is 50. | [optional] [default to null]
**IsAutoUpdateEnabled** | **bool** | Whether the user&#39;s bank connections will get updated in the course of finAPI&#39;s automatic batch update. Note that the automatic batch update will only update bank connections where all of the following applies:&lt;br&gt;&lt;br&gt; - the PIN is stored in finAPI for the bank connection&lt;br&gt; - the user has accepted the latest Terms and Conditions (this only applies to users whose mandator has PISP license or doesn&#39;t have a license at all)&lt;br&gt; - the user has allowed finAPI to use his old stored credentials (this only applies to users which had stored credentials before introducing a web form feature and whose mandator has PISP license or doesn&#39;t have a license at all)&lt;br&gt; - the previous update using the stored credentials did not fail due to the credentials being incorrect (or there was no previous update with the stored credentials)&lt;br&gt; - the bank that the bank connection relates to is included in the automatic batch update (please contact your Sys-Admin for details about the batch update configuration)&lt;br&gt;&lt;br&gt;Also note that the automatic batch update must generally be enabled for your client in order for this field to have any effect.&lt;br/&gt;&lt;br/&gt;WARNING: The automatic update will always download transactions and security positions for any account that it updates! This means that the user will no longer be able to download just the balances for his accounts once the automatic update has run (The &#39;skipPositionsDownload&#39; flag in the bankConnections/update service can no longer be set to true).&lt;br/&gt;&lt;br/&gt;If not specified, then the automatic update will be disabled by default (false). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


