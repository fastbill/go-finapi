# ConnectInterfaceParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankConnectionId** | **int64** | Bank connection identifier | 
**Interface** | [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; The interface to use for connecting with the bank. | 
**SourceInterface** | Pointer to [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; The source interface that should be used as the source of credentials. Set it to one of already existing bank connection&#39;s interfaces and finAPI will try to use the stored credentials of that interface for the current service call. The source interface must fit the following requirements:&lt;br/&gt;- it must have the same set of bank login fields as the main interface (the &#39;interface&#39; parameter);&lt;br/&gt;- it must have stored values for all its bank login fields.&lt;br/&gt;If any of those conditions are not met - the service will throw an appropriate error.&lt;br/&gt;Note: the source interface is ignored if any login credentials are given. | [optional] 
**LoginCredentials** | Pointer to [**[]LoginCredential**](LoginCredential.md) | &lt;strong&gt;Type:&lt;/strong&gt; LoginCredential&lt;br/&gt; Set of login credentials. Must be passed in combination with the &#39;interface&#39; field. | [optional] 
**StoreSecrets** | Pointer to **bool** | Whether to store the secret login fields. If the secret fields are stored, then updates can be triggered without the involvement of the users, as long as the credentials remain valid and the bank consent has not expired. Note that bank consent will be stored regardless of the field value. Default value is false. | [optional] [default to false]
**SkipPositionsDownload** | Pointer to **bool** | Whether to skip the download of transactions and securities or not. If set to true, then finAPI will download just the accounts list with the accounts&#39; information (like account name, number, holder, etc), as well as the accounts&#39; balances (if possible), but skip the download of transactions and securities. Default is false.&lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt;&amp;bull; Setting this flag to true is only meant to be used if A) you generally never download positions, because you are only interested in the accounts list and/or balances, or B) you want to get just the list of accounts in a first step, and then delete unwanted accounts from the bank connection, before you trigger another update with transactions download. This approach allows you to download transactions only for the accounts that you want.&lt;br/&gt;&amp;bull; If you skip the download of transactions and securities during an import or update, you can still download them on a later update (though you might not get all positions at a later point, because the date range in which the bank servers provide this data is usually limited).&lt;br/&gt;&amp;bull; If an account already had any positions imported before an update, and you skip the positions download in the update, then the account&#39;s updated balance might not add up to the set of transactions / security positions. Be aware that certain services (like GET /accounts/dailyBalances) may give incorrect results for accounts in such a state.&lt;br/&gt;&amp;bull; If this bank connection is updated via finAPI&#39;s automatic batch update, then transactions and security positions (of already imported accounts) &lt;u&gt;will&lt;/u&gt; be downloaded in any case!&lt;br/&gt;&amp;bull; For security accounts, skipping the downloading of the securities might result in the account&#39;s balance also not being downloaded.&lt;br/&gt;&amp;bull; For Bausparen accounts, this field is ignored. finAPI will always download transactions for Bausparen accounts.&lt;br/&gt; | [optional] [default to false]
**LoadOwnerData** | Pointer to **bool** | Whether to load information about the bank connection owner(s) - see field &#39;owners&#39;. Default value is &#39;false&#39;.&lt;br/&gt;&lt;br/&gt;NOTE: This feature is supported only by the WEB_SCRAPER interface. | [optional] [default to false]
**AccountTypes** | Pointer to [**[]AccountType**](AccountType.md) |  | [optional] 
**AccountReferences** | Pointer to [**[]AccountReference**](AccountReference.md) | &lt;strong&gt;Type:&lt;/strong&gt; AccountReference&lt;br/&gt; List of accounts for which access is requested from the bank. It must only be passed if the bank interface has the DETAILED_CONSENT property set. | [optional] 
**MultiStepAuthentication** | Pointer to [**MultiStepAuthenticationCallback**](MultiStepAuthenticationCallback.md) | &lt;strong&gt;Type:&lt;/strong&gt; MultiStepAuthenticationCallback&lt;br/&gt; Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication. | [optional] 
**RedirectUrl** | Pointer to **string** | Must only be passed when the used interface has the property REDIRECT_APPROACH. The user will be redirected to the given URL from the bank&#39;s website after completing the bank login and (possibly) the SCA. | [optional] 
**MaxDaysForDownload** | Pointer to **int32** | This setting defines how much of an account&#39;s transactions history will be downloaded whenever a new account is imported. More technically, it depicts the number of days to download transactions for, starting from - and including - the date of the account import. For example, on an account import that happens today, the value 30 would instruct finAPI to download transactions from the past 30 days (including today). The minimum allowed value is 14, the maximum value is 3650. Also possible is the value 0 (which is the default value), in which case there will be no limit to the transactions download and finAPI will try to get all transactions that it can. &lt;br/&gt;&lt;br/&gt;Note:&lt;br/&gt;&amp;bull; There is no guarantee that finAPI will actually download transactions for the entire defined date range, as there may be limitations to the download range (set by the bank or by finAPI, e.g. see ClientConfiguration.transactionImportLimitation). &lt;br/&gt;&amp;bull; This parameter only applies to transactions, not to security positions; For security accounts, finAPI will always download all security positions that it can. &lt;br/&gt;&amp;bull; This setting is stored for each interface individually.&lt;br/&gt;&amp;bull; After an interface has been connected with this setting, there is no way to change the setting for that interface afterwards.&lt;br/&gt;&amp;bull; &lt;b&gt;If you do not limit the download range to a value less than 90 days, the bank is more likely to trigger a strong customer authentication request for the user when finAPI is attempting to download the transactions.&lt;/b&gt; | [optional] [default to 0]

## Methods

### NewConnectInterfaceParams

`func NewConnectInterfaceParams(bankConnectionId int64, interface_ BankingInterface, ) *ConnectInterfaceParams`

NewConnectInterfaceParams instantiates a new ConnectInterfaceParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectInterfaceParamsWithDefaults

`func NewConnectInterfaceParamsWithDefaults() *ConnectInterfaceParams`

NewConnectInterfaceParamsWithDefaults instantiates a new ConnectInterfaceParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankConnectionId

`func (o *ConnectInterfaceParams) GetBankConnectionId() int64`

GetBankConnectionId returns the BankConnectionId field if non-nil, zero value otherwise.

### GetBankConnectionIdOk

`func (o *ConnectInterfaceParams) GetBankConnectionIdOk() (*int64, bool)`

GetBankConnectionIdOk returns a tuple with the BankConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionId

`func (o *ConnectInterfaceParams) SetBankConnectionId(v int64)`

SetBankConnectionId sets BankConnectionId field to given value.


### GetInterface

`func (o *ConnectInterfaceParams) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *ConnectInterfaceParams) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *ConnectInterfaceParams) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.


### GetSourceInterface

`func (o *ConnectInterfaceParams) GetSourceInterface() BankingInterface`

GetSourceInterface returns the SourceInterface field if non-nil, zero value otherwise.

### GetSourceInterfaceOk

`func (o *ConnectInterfaceParams) GetSourceInterfaceOk() (*BankingInterface, bool)`

GetSourceInterfaceOk returns a tuple with the SourceInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInterface

`func (o *ConnectInterfaceParams) SetSourceInterface(v BankingInterface)`

SetSourceInterface sets SourceInterface field to given value.

### HasSourceInterface

`func (o *ConnectInterfaceParams) HasSourceInterface() bool`

HasSourceInterface returns a boolean if a field has been set.

### GetLoginCredentials

`func (o *ConnectInterfaceParams) GetLoginCredentials() []LoginCredential`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *ConnectInterfaceParams) GetLoginCredentialsOk() (*[]LoginCredential, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *ConnectInterfaceParams) SetLoginCredentials(v []LoginCredential)`

SetLoginCredentials sets LoginCredentials field to given value.

### HasLoginCredentials

`func (o *ConnectInterfaceParams) HasLoginCredentials() bool`

HasLoginCredentials returns a boolean if a field has been set.

### GetStoreSecrets

`func (o *ConnectInterfaceParams) GetStoreSecrets() bool`

GetStoreSecrets returns the StoreSecrets field if non-nil, zero value otherwise.

### GetStoreSecretsOk

`func (o *ConnectInterfaceParams) GetStoreSecretsOk() (*bool, bool)`

GetStoreSecretsOk returns a tuple with the StoreSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSecrets

`func (o *ConnectInterfaceParams) SetStoreSecrets(v bool)`

SetStoreSecrets sets StoreSecrets field to given value.

### HasStoreSecrets

`func (o *ConnectInterfaceParams) HasStoreSecrets() bool`

HasStoreSecrets returns a boolean if a field has been set.

### GetSkipPositionsDownload

`func (o *ConnectInterfaceParams) GetSkipPositionsDownload() bool`

GetSkipPositionsDownload returns the SkipPositionsDownload field if non-nil, zero value otherwise.

### GetSkipPositionsDownloadOk

`func (o *ConnectInterfaceParams) GetSkipPositionsDownloadOk() (*bool, bool)`

GetSkipPositionsDownloadOk returns a tuple with the SkipPositionsDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPositionsDownload

`func (o *ConnectInterfaceParams) SetSkipPositionsDownload(v bool)`

SetSkipPositionsDownload sets SkipPositionsDownload field to given value.

### HasSkipPositionsDownload

`func (o *ConnectInterfaceParams) HasSkipPositionsDownload() bool`

HasSkipPositionsDownload returns a boolean if a field has been set.

### GetLoadOwnerData

`func (o *ConnectInterfaceParams) GetLoadOwnerData() bool`

GetLoadOwnerData returns the LoadOwnerData field if non-nil, zero value otherwise.

### GetLoadOwnerDataOk

`func (o *ConnectInterfaceParams) GetLoadOwnerDataOk() (*bool, bool)`

GetLoadOwnerDataOk returns a tuple with the LoadOwnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOwnerData

`func (o *ConnectInterfaceParams) SetLoadOwnerData(v bool)`

SetLoadOwnerData sets LoadOwnerData field to given value.

### HasLoadOwnerData

`func (o *ConnectInterfaceParams) HasLoadOwnerData() bool`

HasLoadOwnerData returns a boolean if a field has been set.

### GetAccountTypes

`func (o *ConnectInterfaceParams) GetAccountTypes() []AccountType`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *ConnectInterfaceParams) GetAccountTypesOk() (*[]AccountType, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *ConnectInterfaceParams) SetAccountTypes(v []AccountType)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *ConnectInterfaceParams) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### GetAccountReferences

`func (o *ConnectInterfaceParams) GetAccountReferences() []AccountReference`

GetAccountReferences returns the AccountReferences field if non-nil, zero value otherwise.

### GetAccountReferencesOk

`func (o *ConnectInterfaceParams) GetAccountReferencesOk() (*[]AccountReference, bool)`

GetAccountReferencesOk returns a tuple with the AccountReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReferences

`func (o *ConnectInterfaceParams) SetAccountReferences(v []AccountReference)`

SetAccountReferences sets AccountReferences field to given value.

### HasAccountReferences

`func (o *ConnectInterfaceParams) HasAccountReferences() bool`

HasAccountReferences returns a boolean if a field has been set.

### GetMultiStepAuthentication

`func (o *ConnectInterfaceParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback`

GetMultiStepAuthentication returns the MultiStepAuthentication field if non-nil, zero value otherwise.

### GetMultiStepAuthenticationOk

`func (o *ConnectInterfaceParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool)`

GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStepAuthentication

`func (o *ConnectInterfaceParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback)`

SetMultiStepAuthentication sets MultiStepAuthentication field to given value.

### HasMultiStepAuthentication

`func (o *ConnectInterfaceParams) HasMultiStepAuthentication() bool`

HasMultiStepAuthentication returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *ConnectInterfaceParams) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ConnectInterfaceParams) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ConnectInterfaceParams) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ConnectInterfaceParams) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetMaxDaysForDownload

`func (o *ConnectInterfaceParams) GetMaxDaysForDownload() int32`

GetMaxDaysForDownload returns the MaxDaysForDownload field if non-nil, zero value otherwise.

### GetMaxDaysForDownloadOk

`func (o *ConnectInterfaceParams) GetMaxDaysForDownloadOk() (*int32, bool)`

GetMaxDaysForDownloadOk returns a tuple with the MaxDaysForDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDaysForDownload

`func (o *ConnectInterfaceParams) SetMaxDaysForDownload(v int32)`

SetMaxDaysForDownload sets MaxDaysForDownload field to given value.

### HasMaxDaysForDownload

`func (o *ConnectInterfaceParams) HasMaxDaysForDownload() bool`

HasMaxDaysForDownload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


