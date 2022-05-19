# ImportBankConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankId** | **int64** | Bank Identifier | 
**Name** | Pointer to **string** | Custom name for the bank connection. Maximum length is 64. If you do not want to set a name, you can leave this field unset. | [optional] 
**BankingUserId** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; and &#39;interface&#39; fields instead. If any of those two fields is used, then the value of this field will be ignored.&lt;br/&gt;&lt;br/&gt;Online banking user ID credential. Max length: 170. NOTES:&lt;br/&gt;- if you import the &#39;demo connection&#39;, this field can be left unset;&lt;br/&gt; - if the user will need to enter his credentials in finAPI&#39;s Web Form, this field can contain any value. It will be ignored. | [optional] 
**BankingCustomerId** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; and &#39;interface&#39; fields instead. If any of those two fields is used, then the value of this field will be ignored.&lt;br/&gt;&lt;br/&gt;Online banking customer ID credential (for most banks this field can remain unset). Max length: 170. NOTES:&lt;br/&gt;- if the user will need to enter his credentials in finAPI&#39;s Web Form, this field can contain any value. It will be ignored. | [optional] 
**BankingPin** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; and &#39;interface&#39; fields instead. If any of those two fields is used, then the value of this field will be ignored.&lt;br/&gt;&lt;br/&gt;Online banking PIN. Max length: 170. Any symbols are allowed. NOTES:&lt;br/&gt;- if you import the &#39;demo connection&#39;, this field can be left unset;&lt;br/&gt; - if the user will need to enter his credentials in finAPI&#39;s Web Form, this field can be left unset or contain any value. It will be ignored. | [optional] 
**StorePin** | Pointer to **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;storeSecrets&#39; field instead.&lt;br/&gt;&lt;br/&gt;Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when updating this bank connection or when executing orders (like money transfers). Default is false. &lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt; - before you set this field to true, please regard the &#39;pinsAreVolatile&#39; flag of this connection&#39;s bank. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;&lt;br/&gt; - this field is ignored in case when the user will need to use finAPI&#39;s Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the &#39;storeSecretsAvailableInWebForm&#39; setting (see Client Configuration). | [optional] [default to false]
**Interface** | Pointer to [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; The interface to use for connecting with the bank. | [optional] 
**LoginCredentials** | Pointer to [**[]LoginCredential**](LoginCredential.md) | &lt;strong&gt;Type:&lt;/strong&gt; LoginCredential&lt;br/&gt; Set of login credentials. Must be passed in combination with the &#39;interface&#39; field. | [optional] 
**StoreSecrets** | Pointer to **bool** | Whether to store the secret login fields. If the secret fields are stored, then updates can be triggered without the involvement of the users, as long as the credentials remain valid and the bank consent has not expired. Note that bank consent will be stored regardless of the field value. Default value is false. | [optional] [default to false]
**SkipPositionsDownload** | Pointer to **bool** | Whether to skip the download of transactions and securities or not. If set to true, then finAPI will download just the accounts list with the accounts&#39; information (like account name, number, holder, etc), as well as the accounts&#39; balances (if possible), but skip the download of transactions and securities. Default is false.&lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt;&amp;bull; Setting this flag to true is only meant to be used if A) you generally never download positions, because you are only interested in the accounts list and/or balances, or B) you want to get just the list of accounts in a first step, and then delete unwanted accounts from the bank connection, before you trigger another update with transactions download. This approach allows you to download transactions only for the accounts that you want.&lt;br/&gt;&amp;bull; If you skip the download of transactions and securities during an import or update, you can still download them on a later update (though you might not get all positions at a later point, because the date range in which the bank servers provide this data is usually limited).&lt;br/&gt;&amp;bull; If an account already had any positions imported before an update, and you skip the positions download in the update, then the account&#39;s updated balance might not add up to the set of transactions / security positions. Be aware that certain services (like GET /accounts/dailyBalances) may give incorrect results for accounts in such a state.&lt;br/&gt;&amp;bull; If this bank connection is updated via finAPI&#39;s automatic batch update, then transactions and security positions (of already imported accounts) &lt;u&gt;will&lt;/u&gt; be downloaded in any case!&lt;br/&gt;&amp;bull; For security accounts, skipping the downloading of the securities might result in the account&#39;s balance also not being downloaded.&lt;br/&gt;&amp;bull; For Bausparen accounts, this field is ignored. finAPI will always download transactions for Bausparen accounts.&lt;br/&gt; | [optional] [default to false]
**LoadOwnerData** | Pointer to **bool** | Whether to load information about the bank connection owner(s) - see field &#39;owners&#39;. Default value is &#39;false&#39;.&lt;br/&gt;&lt;br/&gt;NOTE: This feature is supported only by the WEB_SCRAPER interface. | [optional] [default to false]
**MaxDaysForDownload** | Pointer to **int32** | This setting defines how much of an account&#39;s transactions history will be downloaded whenever a new account is imported. More technically, it depicts the number of days to download transactions for, starting from - and including - the date of the account import. For example, on an account import that happens today, the value 30 would instruct finAPI to download transactions from the past 30 days (including today). The minimum allowed value is 14, the maximum value is 3650. Also possible is the value 0 (which is the default value), in which case there will be no limit to the transactions download and finAPI will try to get all transactions that it can. &lt;br/&gt;&lt;br/&gt;Note:&lt;br/&gt;&amp;bull; There is no guarantee that finAPI will actually download transactions for the entire defined date range, as there may be limitations to the download range (set by the bank or by finAPI, e.g. see ClientConfiguration.transactionImportLimitation). &lt;br/&gt;&amp;bull; This parameter only applies to transactions, not to security positions; For security accounts, finAPI will always download all security positions that it can. &lt;br/&gt;&amp;bull; This setting is stored for each interface individually.&lt;br/&gt;&amp;bull; After an interface has been connected with this setting, there is no way to change the setting for that interface afterwards.&lt;br/&gt;&amp;bull; &lt;b&gt;If you do not limit the download range to a value less than 90 days, the bank is more likely to trigger a strong customer authentication request for the user when finAPI is attempting to download the transactions.&lt;/b&gt; | [optional] [default to 0]
**AccountTypes** | Pointer to [**[]AccountType**](AccountType.md) |  | [optional] 
**AccountTypeIds** | Pointer to **[]int64** |  | [optional] 
**AccountReferences** | Pointer to [**[]AccountReference**](AccountReference.md) | &lt;strong&gt;Type:&lt;/strong&gt; AccountReference&lt;br/&gt; List of accounts for which access is requested from the bank. It must only be passed if the bank interface has the DETAILED_CONSENT property set. | [optional] 
**ChallengeResponse** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;multiStepAuthentication&#39; field instead.&lt;br/&gt;&lt;br/&gt;Challenge response. This field should be set only when the previous attempt of importing the bank connection failed with HTTP code 510, i.e. the bank sent a challenge for the user for an additional authentication. In this case, this field must contain the response to the bank&#39;s challenge. Note that in the context of finAPI&#39;s Web Form flow, finAPI will automatically deal with getting the challenge response from the user via the Web Form. | [optional] 
**MultiStepAuthentication** | Pointer to [**MultiStepAuthenticationCallback**](MultiStepAuthenticationCallback.md) | &lt;strong&gt;Type:&lt;/strong&gt; MultiStepAuthenticationCallback&lt;br/&gt; Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication. | [optional] 
**RedirectUrl** | Pointer to **string** | Must only be passed when the used interface has the property REDIRECT_APPROACH. The user will be redirected to the given URL from the bank&#39;s website after completing the bank login and (possibly) the SCA. | [optional] 

## Methods

### NewImportBankConnectionParams

`func NewImportBankConnectionParams(bankId int64, ) *ImportBankConnectionParams`

NewImportBankConnectionParams instantiates a new ImportBankConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportBankConnectionParamsWithDefaults

`func NewImportBankConnectionParamsWithDefaults() *ImportBankConnectionParams`

NewImportBankConnectionParamsWithDefaults instantiates a new ImportBankConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankId

`func (o *ImportBankConnectionParams) GetBankId() int64`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *ImportBankConnectionParams) GetBankIdOk() (*int64, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *ImportBankConnectionParams) SetBankId(v int64)`

SetBankId sets BankId field to given value.


### GetName

`func (o *ImportBankConnectionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportBankConnectionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportBankConnectionParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ImportBankConnectionParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBankingUserId

`func (o *ImportBankConnectionParams) GetBankingUserId() string`

GetBankingUserId returns the BankingUserId field if non-nil, zero value otherwise.

### GetBankingUserIdOk

`func (o *ImportBankConnectionParams) GetBankingUserIdOk() (*string, bool)`

GetBankingUserIdOk returns a tuple with the BankingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingUserId

`func (o *ImportBankConnectionParams) SetBankingUserId(v string)`

SetBankingUserId sets BankingUserId field to given value.

### HasBankingUserId

`func (o *ImportBankConnectionParams) HasBankingUserId() bool`

HasBankingUserId returns a boolean if a field has been set.

### GetBankingCustomerId

`func (o *ImportBankConnectionParams) GetBankingCustomerId() string`

GetBankingCustomerId returns the BankingCustomerId field if non-nil, zero value otherwise.

### GetBankingCustomerIdOk

`func (o *ImportBankConnectionParams) GetBankingCustomerIdOk() (*string, bool)`

GetBankingCustomerIdOk returns a tuple with the BankingCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingCustomerId

`func (o *ImportBankConnectionParams) SetBankingCustomerId(v string)`

SetBankingCustomerId sets BankingCustomerId field to given value.

### HasBankingCustomerId

`func (o *ImportBankConnectionParams) HasBankingCustomerId() bool`

HasBankingCustomerId returns a boolean if a field has been set.

### GetBankingPin

`func (o *ImportBankConnectionParams) GetBankingPin() string`

GetBankingPin returns the BankingPin field if non-nil, zero value otherwise.

### GetBankingPinOk

`func (o *ImportBankConnectionParams) GetBankingPinOk() (*string, bool)`

GetBankingPinOk returns a tuple with the BankingPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingPin

`func (o *ImportBankConnectionParams) SetBankingPin(v string)`

SetBankingPin sets BankingPin field to given value.

### HasBankingPin

`func (o *ImportBankConnectionParams) HasBankingPin() bool`

HasBankingPin returns a boolean if a field has been set.

### GetStorePin

`func (o *ImportBankConnectionParams) GetStorePin() bool`

GetStorePin returns the StorePin field if non-nil, zero value otherwise.

### GetStorePinOk

`func (o *ImportBankConnectionParams) GetStorePinOk() (*bool, bool)`

GetStorePinOk returns a tuple with the StorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePin

`func (o *ImportBankConnectionParams) SetStorePin(v bool)`

SetStorePin sets StorePin field to given value.

### HasStorePin

`func (o *ImportBankConnectionParams) HasStorePin() bool`

HasStorePin returns a boolean if a field has been set.

### GetInterface

`func (o *ImportBankConnectionParams) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *ImportBankConnectionParams) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *ImportBankConnectionParams) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *ImportBankConnectionParams) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetLoginCredentials

`func (o *ImportBankConnectionParams) GetLoginCredentials() []LoginCredential`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *ImportBankConnectionParams) GetLoginCredentialsOk() (*[]LoginCredential, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *ImportBankConnectionParams) SetLoginCredentials(v []LoginCredential)`

SetLoginCredentials sets LoginCredentials field to given value.

### HasLoginCredentials

`func (o *ImportBankConnectionParams) HasLoginCredentials() bool`

HasLoginCredentials returns a boolean if a field has been set.

### GetStoreSecrets

`func (o *ImportBankConnectionParams) GetStoreSecrets() bool`

GetStoreSecrets returns the StoreSecrets field if non-nil, zero value otherwise.

### GetStoreSecretsOk

`func (o *ImportBankConnectionParams) GetStoreSecretsOk() (*bool, bool)`

GetStoreSecretsOk returns a tuple with the StoreSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSecrets

`func (o *ImportBankConnectionParams) SetStoreSecrets(v bool)`

SetStoreSecrets sets StoreSecrets field to given value.

### HasStoreSecrets

`func (o *ImportBankConnectionParams) HasStoreSecrets() bool`

HasStoreSecrets returns a boolean if a field has been set.

### GetSkipPositionsDownload

`func (o *ImportBankConnectionParams) GetSkipPositionsDownload() bool`

GetSkipPositionsDownload returns the SkipPositionsDownload field if non-nil, zero value otherwise.

### GetSkipPositionsDownloadOk

`func (o *ImportBankConnectionParams) GetSkipPositionsDownloadOk() (*bool, bool)`

GetSkipPositionsDownloadOk returns a tuple with the SkipPositionsDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipPositionsDownload

`func (o *ImportBankConnectionParams) SetSkipPositionsDownload(v bool)`

SetSkipPositionsDownload sets SkipPositionsDownload field to given value.

### HasSkipPositionsDownload

`func (o *ImportBankConnectionParams) HasSkipPositionsDownload() bool`

HasSkipPositionsDownload returns a boolean if a field has been set.

### GetLoadOwnerData

`func (o *ImportBankConnectionParams) GetLoadOwnerData() bool`

GetLoadOwnerData returns the LoadOwnerData field if non-nil, zero value otherwise.

### GetLoadOwnerDataOk

`func (o *ImportBankConnectionParams) GetLoadOwnerDataOk() (*bool, bool)`

GetLoadOwnerDataOk returns a tuple with the LoadOwnerData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadOwnerData

`func (o *ImportBankConnectionParams) SetLoadOwnerData(v bool)`

SetLoadOwnerData sets LoadOwnerData field to given value.

### HasLoadOwnerData

`func (o *ImportBankConnectionParams) HasLoadOwnerData() bool`

HasLoadOwnerData returns a boolean if a field has been set.

### GetMaxDaysForDownload

`func (o *ImportBankConnectionParams) GetMaxDaysForDownload() int32`

GetMaxDaysForDownload returns the MaxDaysForDownload field if non-nil, zero value otherwise.

### GetMaxDaysForDownloadOk

`func (o *ImportBankConnectionParams) GetMaxDaysForDownloadOk() (*int32, bool)`

GetMaxDaysForDownloadOk returns a tuple with the MaxDaysForDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDaysForDownload

`func (o *ImportBankConnectionParams) SetMaxDaysForDownload(v int32)`

SetMaxDaysForDownload sets MaxDaysForDownload field to given value.

### HasMaxDaysForDownload

`func (o *ImportBankConnectionParams) HasMaxDaysForDownload() bool`

HasMaxDaysForDownload returns a boolean if a field has been set.

### GetAccountTypes

`func (o *ImportBankConnectionParams) GetAccountTypes() []AccountType`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *ImportBankConnectionParams) GetAccountTypesOk() (*[]AccountType, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *ImportBankConnectionParams) SetAccountTypes(v []AccountType)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *ImportBankConnectionParams) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### GetAccountTypeIds

`func (o *ImportBankConnectionParams) GetAccountTypeIds() []int64`

GetAccountTypeIds returns the AccountTypeIds field if non-nil, zero value otherwise.

### GetAccountTypeIdsOk

`func (o *ImportBankConnectionParams) GetAccountTypeIdsOk() (*[]int64, bool)`

GetAccountTypeIdsOk returns a tuple with the AccountTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeIds

`func (o *ImportBankConnectionParams) SetAccountTypeIds(v []int64)`

SetAccountTypeIds sets AccountTypeIds field to given value.

### HasAccountTypeIds

`func (o *ImportBankConnectionParams) HasAccountTypeIds() bool`

HasAccountTypeIds returns a boolean if a field has been set.

### GetAccountReferences

`func (o *ImportBankConnectionParams) GetAccountReferences() []AccountReference`

GetAccountReferences returns the AccountReferences field if non-nil, zero value otherwise.

### GetAccountReferencesOk

`func (o *ImportBankConnectionParams) GetAccountReferencesOk() (*[]AccountReference, bool)`

GetAccountReferencesOk returns a tuple with the AccountReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountReferences

`func (o *ImportBankConnectionParams) SetAccountReferences(v []AccountReference)`

SetAccountReferences sets AccountReferences field to given value.

### HasAccountReferences

`func (o *ImportBankConnectionParams) HasAccountReferences() bool`

HasAccountReferences returns a boolean if a field has been set.

### GetChallengeResponse

`func (o *ImportBankConnectionParams) GetChallengeResponse() string`

GetChallengeResponse returns the ChallengeResponse field if non-nil, zero value otherwise.

### GetChallengeResponseOk

`func (o *ImportBankConnectionParams) GetChallengeResponseOk() (*string, bool)`

GetChallengeResponseOk returns a tuple with the ChallengeResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeResponse

`func (o *ImportBankConnectionParams) SetChallengeResponse(v string)`

SetChallengeResponse sets ChallengeResponse field to given value.

### HasChallengeResponse

`func (o *ImportBankConnectionParams) HasChallengeResponse() bool`

HasChallengeResponse returns a boolean if a field has been set.

### GetMultiStepAuthentication

`func (o *ImportBankConnectionParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback`

GetMultiStepAuthentication returns the MultiStepAuthentication field if non-nil, zero value otherwise.

### GetMultiStepAuthenticationOk

`func (o *ImportBankConnectionParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool)`

GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStepAuthentication

`func (o *ImportBankConnectionParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback)`

SetMultiStepAuthentication sets MultiStepAuthentication field to given value.

### HasMultiStepAuthentication

`func (o *ImportBankConnectionParams) HasMultiStepAuthentication() bool`

HasMultiStepAuthentication returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *ImportBankConnectionParams) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ImportBankConnectionParams) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ImportBankConnectionParams) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ImportBankConnectionParams) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


