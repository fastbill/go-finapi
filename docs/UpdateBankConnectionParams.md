# UpdateBankConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankConnectionId** | **int64** | Bank connection identifier | 
**BankingPin** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; and &#39;interface&#39; fields instead. If any of those two fields is used, then the value of this field will be ignored.&lt;br&gt;&lt;br&gt;Online banking PIN. Any symbols are allowed. Max length: 170. If a PIN is stored in the bank connection, then this field may remain unset. If finAPI&#39;s Web Form is not required and the field is set though then it will always be used (even if there is some other PIN stored in the bank connection). If you want the user to enter a PIN in finAPI&#39;s Web Form even when a PIN is stored, then just set the field to any value, so that the service recognizes that you wish to use the Web Form flow. | [optional] 
**StorePin** | Pointer to **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;storeSecrets&#39; field instead.&lt;br&gt;&lt;br&gt;Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when updating this bank connection or when executing orders (like money transfers). Default is false. &lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt; - before you set this field to true, please regard the &#39;pinsAreVolatile&#39; flag of this connection&#39;s bank. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;&lt;br/&gt; - this field is ignored in case when the user will need to use finAPI&#39;s Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the &#39;storeSecretsAvailableInWebForm&#39; setting (see Client Configuration). | [optional] [default to false]
**Interface** | Pointer to [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; The interface to use for connecting with the bank. | [optional] 
**LoginCredentials** | Pointer to [**[]LoginCredential**](LoginCredential.md) | &lt;strong&gt;Type:&lt;/strong&gt; LoginCredential&lt;br/&gt; Set of login credentials. Must be passed in combination with the &#39;interface&#39; field, if the credentials have not been previously stored. The labels that you pass must match with the login credential labels that the respective interface defines. finAPI will combine the given credentials with any credentials that it has stored. You can leave this field unset in case finAPI has stored all required credentials. | [optional] 
**StoreSecrets** | Pointer to **bool** | Whether to store the secret login fields. If the secret fields are stored, then updates can be triggered without the involvement of the users, as long as the credentials remain valid and the bank consent has not expired. Note that bank consent will be stored regardless of the field value. Default value is false. | [optional] [default to false]
**ImportNewAccounts** | Pointer to **bool** | Whether new accounts that have not yet been imported will be imported or not. Default is false. &lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt;&amp;bull; For best performance of the bank connection update, you should not enable this flag unless you really expect new accounts to be available in the connection. It is recommended to let your users tell you through your application when they want the service to look for new accounts.&lt;br/&gt;&amp;bull; If you have imported an interface using a limited set of account types, you would import all other accounts (e.g. security accounts or credit cards) by setting &lt;code&gt;importNewAccounts&lt;/code&gt; to &lt;code&gt;true&lt;/code&gt;. To avoid importing account types that you are not interested in, make sure to keep this field unset (or set to false) for the update. | [optional] [default to false]
**SkipPositionsDownload** | Pointer to **bool** | Whether to skip the download of transactions and securities or not. If set to true, then finAPI will download just the accounts list with the accounts&#39; information (like account name, number, holder, etc), as well as the accounts&#39; balances (if possible), but skip the download of transactions and securities. Default is false.&lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt;&amp;bull; Setting this flag to true is only meant to be used if A) you generally never download positions, because you are only interested in the accounts list and/or balances, or B) you want to get just the list of accounts in a first step, and then delete unwanted accounts from the bank connection, before you trigger another update with transactions download. This approach allows you to download transactions only for the accounts that you want.&lt;br/&gt;&amp;bull; If you skip the download of transactions and securities during an import or update, you can still download them on a later update (though you might not get all positions at a later point, because the date range in which the bank servers provide this data is usually limited).&lt;br/&gt;&amp;bull; If an account already had any positions imported before an update, and you skip the positions download in the update, then the account&#39;s updated balance might not add up to the set of transactions / security positions. Be aware that certain services (like GET /accounts/dailyBalances) may give incorrect results for accounts in such a state.&lt;br/&gt;&amp;bull; If this bank connection is updated via finAPI&#39;s automatic batch update, then transactions and security positions (of already imported accounts) &lt;u&gt;will&lt;/u&gt; be downloaded in any case!&lt;br/&gt;&amp;bull; For security accounts, skipping the downloading of the securities might result in the account&#39;s balance also not being downloaded.&lt;br/&gt;&amp;bull; For Bausparen accounts, this field is ignored. finAPI will always download transactions for Bausparen accounts.&lt;br/&gt; | [optional] [default to false]
**LoadOwnerData** | Pointer to **bool** | Whether to load/refresh information about the bank connection owner(s) - see field &#39;owners&#39;. Default value is &#39;false&#39;. Note that owner data is NOT loaded/refreshed during finAPI&#39;s automatic bank connection update. | [optional] [default to false]
**AccountTypes** | Pointer to [**[]AccountType**](AccountType.md) |  | [optional] 
**AccountReferences** | Pointer to [**[]AccountReference**](AccountReference.md) | &lt;strong&gt;Type:&lt;/strong&gt; AccountReference&lt;br/&gt; List of accounts for which access is requested from the bank. It may only be passed if the bank interface has the DETAILED_CONSENT property set. if omitted, finAPI will use the list of existing accounts. Note that the parameter is still required if you want to import new accounts (i.e. call with importNewAccounts&#x3D;true). | [optional] 
**ChallengeResponse** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;multiStepAuthentication&#39; field instead.&lt;br/&gt;&lt;br/&gt;Challenge response. This field should be set only when the previous attempt of update the bank connection failed with HTTP code 510, i.e. the bank sent a challenge for the user for an additional authentication. In this case, this field must contain the response to the bank&#39;s challenge. Note that in the context of finAPI&#39;s Web Form flow, finAPI will automatically deal with getting the challenge response from the user via the Web Form. | [optional] 
**RedirectUrl** | Pointer to **string** | Must only be passed when the used interface has the property REDIRECT_APPROACH. The user will be redirected to the given URL from the bank&#39;s website after completing the bank login and (possibly) the SCA. | [optional] 
**ForceWebForm** | Pointer to **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/tasks/backgroundUpdate&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;editSavedSettings&#39; field instead.&lt;br/&gt;&lt;br/&gt;If the user has stored credentials in finAPI for the selected bank connection, then the finAPI Web Form will only be shown when the user must be involved for a second authentication, or when the previous connection to the bank via the selected interface had failed. However if you want to provide the Web Form to the user in any case, you can set this field to true. It will force the Web Form flow for the user and allow him to make changes to the data that he has stored in finAPI. Default value is &#39;false&#39;. | [optional] [default to false]
**MultiStepAuthentication** | Pointer to [**MultiStepAuthenticationCallback**](MultiStepAuthenticationCallback.md) | &lt;strong&gt;Type:&lt;/strong&gt; MultiStepAuthenticationCallback&lt;br/&gt; Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication. | [optional] 

## Methods

### NewUpdateBankConnectionParams

`func NewUpdateBankConnectionParams(bankConnectionId int64, ) *UpdateBankConnectionParams`

NewUpdateBankConnectionParams instantiates a new UpdateBankConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBankConnectionParamsWithDefaults

`func NewUpdateBankConnectionParamsWithDefaults() *UpdateBankConnectionParams`

NewUpdateBankConnectionParamsWithDefaults instantiates a new UpdateBankConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankConnectionId

`func (o *UpdateBankConnectionParams) GetBankConnectionId() int64`

GetBankConnectionId returns the BankConnectionId field if non-nil, zero value otherwise.

### GetBankConnectionIdOk

`func (o *UpdateBankConnectionParams) GetBankConnectionIdOk() (*int64, bool)`

GetBankConnectionIdOk returns a tuple with the BankConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionId

`func (o *UpdateBankConnectionParams) SetBankConnectionId(v int64)`

SetBankConnectionId sets BankConnectionId field to given value.


### GetBankingPin

`func (o *UpdateBankConnectionParams) GetBankingPin() string`

GetBankingPin returns the BankingPin field if non-nil, zero value otherwise.

### GetBankingPinOk

`func (o *UpdateBankConnectionParams) GetBankingPinOk() (*string, bool)`

GetBankingPinOk returns a tuple with the BankingPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingPin

`func (o *UpdateBankConnectionParams) SetBankingPin(v string)`

SetBankingPin sets BankingPin field to given value.

### HasBankingPin

`func (o *UpdateBankConnectionParams) HasBankingPin() bool`

HasBankingPin returns a boolean if a field has been set.

### GetStorePin

`func (o *UpdateBankConnectionParams) GetStorePin() bool`

GetStorePin returns the StorePin field if non-nil, zero value otherwise.

### GetStorePinOk

`func (o *UpdateBankConnectionParams) GetStorePinOk() (*bool, bool)`

GetStorePinOk returns a tuple with the StorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePin

`func (o *UpdateBankConnectionParams) SetStorePin(v bool)`

SetStorePin sets StorePin field to given value.

### HasStorePin

`func (o *UpdateBankConnectionParams) HasStorePin() bool`

HasStorePin returns a boolean if a field has been set.

### GetInterface

`func (o *UpdateBankConnectionParams) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *UpdateBankConnectionParams) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *UpdateBankConnectionParams) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *UpdateBankConnectionParams) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetLoginCredentials

`func (o *UpdateBankConnectionParams) GetLoginCredentials() []LoginCredential`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *UpdateBankConnectionParams) GetLoginCredentialsOk() (*[]LoginCredential, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *UpdateBankConnectionParams) SetLoginCredentials(v []LoginCredential)`

SetLoginCredentials sets LoginCredentials field to given value.

### HasLoginCredentials

`func (o *UpdateBankConnectionParams) HasLoginCredentials() bool`

HasLoginCredentials returns a boolean if a field has been set.

### GetStoreSecrets

`func (o *UpdateBankConnectionParams) GetStoreSecrets() bool`

GetStoreSecrets returns the StoreSecrets field if non-nil, zero value otherwise.

### GetStoreSecretsOk

`func (o *UpdateBankConnectionParams) GetStoreSecretsOk() (*bool, bool)`

GetStoreSecretsOk returns a tuple with the StoreSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSecrets

`func (o *UpdateBankConnectionParams) SetStoreSecrets(v bool)`

SetStoreSecrets sets StoreSecrets field to given value.

### HasStoreSecrets

`func (o *UpdateBankConnectionParams) HasStoreSecrets() bool`

HasStoreSecrets returns a boolean if a field has been set.

### GetImportNewAccounts

`func (o *UpdateBankConnectionParams) GetImportNewAccounts() bool`

GetImportNewAccounts returns the ImportNewAccounts field if non-nil, zero value otherwise.

### GetImportNewAccountsOk

`func (o *UpdateBankConnectionParams) GetImportNewAccountsOk() (*bool, bool)`

GetImportNewAccountsOk returns a tuple with the ImportNewAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportNewAccounts

`func (o *UpdateBankConnectionParams) SetImportNewAccounts(v bool)`

SetImportNewAccounts sets ImportNewAccounts field to given value.

### HasImportNewAccounts

`func (o *UpdateBankConnectionParams) HasImportNewAccounts() bool`

HasImportNewAccounts returns a boolean if a field has been set.

### GetSkipPositionsDownload

`func (o *UpdateBankConnectionParams) GetSkipPositionsDownload() bool`

GetSkipPositionsDownload returns the SkipPositionsDownload field if non-nil, zero value otherwise.

### GetSkipPositionsDownloadOk

`func (o *UpdateBankConnectionParams) GetSkipPositionsDownloadOk() (*bool, bool)`

GetSkipPositionsDownloadOk returns a tuple with the SkipPositionsDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPositionsDownload

`func (o *UpdateBankConnectionParams) SetSkipPositionsDownload(v bool)`

SetSkipPositionsDownload sets SkipPositionsDownload field to given value.

### HasSkipPositionsDownload

`func (o *UpdateBankConnectionParams) HasSkipPositionsDownload() bool`

HasSkipPositionsDownload returns a boolean if a field has been set.

### GetLoadOwnerData

`func (o *UpdateBankConnectionParams) GetLoadOwnerData() bool`

GetLoadOwnerData returns the LoadOwnerData field if non-nil, zero value otherwise.

### GetLoadOwnerDataOk

`func (o *UpdateBankConnectionParams) GetLoadOwnerDataOk() (*bool, bool)`

GetLoadOwnerDataOk returns a tuple with the LoadOwnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOwnerData

`func (o *UpdateBankConnectionParams) SetLoadOwnerData(v bool)`

SetLoadOwnerData sets LoadOwnerData field to given value.

### HasLoadOwnerData

`func (o *UpdateBankConnectionParams) HasLoadOwnerData() bool`

HasLoadOwnerData returns a boolean if a field has been set.

### GetAccountTypes

`func (o *UpdateBankConnectionParams) GetAccountTypes() []AccountType`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *UpdateBankConnectionParams) GetAccountTypesOk() (*[]AccountType, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *UpdateBankConnectionParams) SetAccountTypes(v []AccountType)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *UpdateBankConnectionParams) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### GetAccountReferences

`func (o *UpdateBankConnectionParams) GetAccountReferences() []AccountReference`

GetAccountReferences returns the AccountReferences field if non-nil, zero value otherwise.

### GetAccountReferencesOk

`func (o *UpdateBankConnectionParams) GetAccountReferencesOk() (*[]AccountReference, bool)`

GetAccountReferencesOk returns a tuple with the AccountReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReferences

`func (o *UpdateBankConnectionParams) SetAccountReferences(v []AccountReference)`

SetAccountReferences sets AccountReferences field to given value.

### HasAccountReferences

`func (o *UpdateBankConnectionParams) HasAccountReferences() bool`

HasAccountReferences returns a boolean if a field has been set.

### GetChallengeResponse

`func (o *UpdateBankConnectionParams) GetChallengeResponse() string`

GetChallengeResponse returns the ChallengeResponse field if non-nil, zero value otherwise.

### GetChallengeResponseOk

`func (o *UpdateBankConnectionParams) GetChallengeResponseOk() (*string, bool)`

GetChallengeResponseOk returns a tuple with the ChallengeResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeResponse

`func (o *UpdateBankConnectionParams) SetChallengeResponse(v string)`

SetChallengeResponse sets ChallengeResponse field to given value.

### HasChallengeResponse

`func (o *UpdateBankConnectionParams) HasChallengeResponse() bool`

HasChallengeResponse returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *UpdateBankConnectionParams) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *UpdateBankConnectionParams) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *UpdateBankConnectionParams) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *UpdateBankConnectionParams) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetForceWebForm

`func (o *UpdateBankConnectionParams) GetForceWebForm() bool`

GetForceWebForm returns the ForceWebForm field if non-nil, zero value otherwise.

### GetForceWebFormOk

`func (o *UpdateBankConnectionParams) GetForceWebFormOk() (*bool, bool)`

GetForceWebFormOk returns a tuple with the ForceWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceWebForm

`func (o *UpdateBankConnectionParams) SetForceWebForm(v bool)`

SetForceWebForm sets ForceWebForm field to given value.

### HasForceWebForm

`func (o *UpdateBankConnectionParams) HasForceWebForm() bool`

HasForceWebForm returns a boolean if a field has been set.

### GetMultiStepAuthentication

`func (o *UpdateBankConnectionParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback`

GetMultiStepAuthentication returns the MultiStepAuthentication field if non-nil, zero value otherwise.

### GetMultiStepAuthenticationOk

`func (o *UpdateBankConnectionParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool)`

GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStepAuthentication

`func (o *UpdateBankConnectionParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback)`

SetMultiStepAuthentication sets MultiStepAuthentication field to given value.

### HasMultiStepAuthentication

`func (o *UpdateBankConnectionParams) HasMultiStepAuthentication() bool`

HasMultiStepAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


