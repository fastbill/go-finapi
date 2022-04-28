# BankConnectionInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; Bank interface. Possible values:&lt;br&gt;&lt;br&gt;&amp;bull; &lt;code&gt;WEB_SCRAPER&lt;/code&gt; - means that finAPI will parse data from the bank&#39;s online banking website.&lt;br&gt;&amp;bull; &lt;code&gt;FINTS_SERVER&lt;/code&gt; - means that finAPI will download data via the bank&#39;s FinTS interface.&lt;br&gt;&amp;bull; &lt;code&gt;XS2A&lt;/code&gt; - means that finAPI will download data via the bank&#39;s XS2A interface.&lt;br&gt; | 
**LoginCredentials** | [**[]LoginCredentialResource**](LoginCredentialResource.md) | &lt;strong&gt;Type:&lt;/strong&gt; LoginCredentialResource&lt;br/&gt; Login fields for this interface (in the order that we suggest to show them to the user), with their currently stored values. Note that this list always contains all existing login fields for this interface, even when there is no stored value for a field (value will be null in such a case). | 
**DefaultTwoStepProcedureId** | **NullableString** | The default two-step-procedure for this interface. Must match one of the available &#39;procedureId&#39;s from the &#39;twoStepProcedures&#39; list. When this field is set, then finAPI will automatically try to select the procedure wherever applicable. Note that the list of available procedures of a bank connection may change as a result of an update of the connection, and if this field references a procedure that is no longer available after an update, finAPI will automatically clear the default procedure (set it to null). | 
**TwoStepProcedures** | [**[]TwoStepProcedure**](TwoStepProcedure.md) | &lt;strong&gt;Type:&lt;/strong&gt; TwoStepProcedure&lt;br/&gt; Available two-step-procedures in this interface, used for submitting a money transfer or direct debit request (see /accounts/requestSepaMoneyTransfer or /requestSepaDirectDebit),or for multi-step-authentication during bank connection import or update. The available two-step-procedures mya be re-evaluated each time this bank connection is updated (/bankConnections/update). This means that this list may change as a result of an update. | 
**AisConsent** | [**NullableBankConsent**](BankConsent.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankConsent&lt;br/&gt; If this field is set, it means that this interface is handing out a consent to finAPI in exchange for the login credentials. finAPI needs to use this consent to get access to the account list and account data (i.e. Account Information Services, AIS). If this field is not set, it means that this interface does not use such consents. | 
**LastManualUpdate** | [**NullableUpdateResult**](UpdateResult.md) | &lt;strong&gt;Type:&lt;/strong&gt; UpdateResult&lt;br/&gt; Result of the last manual update of the associated bank connection using this interface. If no manual update has ever been done so far with this interface, then this field will not be set. | 
**LastAutoUpdate** | [**NullableUpdateResult**](UpdateResult.md) | &lt;strong&gt;Type:&lt;/strong&gt; UpdateResult&lt;br/&gt; Result of the last auto update of the associated bank connection using this interface (ran by finAPI&#39;s automatic batch update process). If no auto update has ever been done so far with this interface, then this field will not be set. | 
**UserActionRequired** | **bool** | This field indicates whether the user&#39;s attention is required for the next update of the given bank connection interface.&lt;br/&gt;If the field is true, finAPI stops auto-updates of this bank connection interface to mitigate the risk of locking the user&#39;s bank account and also of triggering a multi-step authentication that might lead to a notification being sent to the end-user.&lt;br/&gt;If the field is false, the user&#39;s attention might still be required for the next bank update, e.g. because of new Terms and Conditions that have to get approved by the user.(this only applies to users whose mandator doesn&#39;t have an AIS license)&lt;br/&gt;Every communication with the bank (e.g. updating a bank connection, submitting a money transfer or a direct debit, etc.) can change the value of this flag. If the field is true, we recommend to ask the end-user to trigger a manual update of the bank connection interface (using the &#39;Update a bank connection&#39; service). If the update completes successfully without triggering a strong customer authentication or results in storing a valid XS2A consent, this flag will switch to false. The logic about determination of the user&#39;s attention being required might change in time. Please use this as a convenience function to know, when you have to involve the user in the next communication with the bank. Once the flag switches to false, the bank connection interface will be enabled again for the auto-update (if it is configured). | 
**MaxDaysForDownload** | **int32** | This setting defines how much of an account&#39;s transactions history will be downloaded whenever a new account is imported. More technically, it depicts the number of days to download transactions for, starting from - and including - the date of the account import. For example, on an account import that happens today, the value 30 would instruct finAPI to download transactions from the past 30 days (including today). The minimum allowed value is 14, the maximum value is 3650. Also possible is the value 0 (which is the default value), in which case there will be no limit to the transactions download and finAPI will try to get all transactions that it can. &lt;br/&gt;&lt;br/&gt;Note:&lt;br/&gt;&amp;bull; There is no guarantee that finAPI will actually download transactions for the entire defined date range, as there may be limitations to the download range (set by the bank or by finAPI, e.g. see ClientConfiguration.transactionImportLimitation). &lt;br/&gt;&amp;bull; This parameter only applies to transactions, not to security positions; For security accounts, finAPI will always download all security positions that it can. &lt;br/&gt;&amp;bull; This setting is stored for each interface individually.&lt;br/&gt;&amp;bull; After an interface has been connected with this setting, there is no way to change the setting for that interface afterwards.&lt;br/&gt;&amp;bull; &lt;b&gt;If you do not limit the download range to a value less than 90 days, the bank is more likely to trigger a strong customer authentication request for the user when finAPI is attempting to download the transactions.&lt;/b&gt; | 

## Methods

### NewBankConnectionInterface

`func NewBankConnectionInterface(interface_ BankingInterface, loginCredentials []LoginCredentialResource, defaultTwoStepProcedureId NullableString, twoStepProcedures []TwoStepProcedure, aisConsent NullableBankConsent, lastManualUpdate NullableUpdateResult, lastAutoUpdate NullableUpdateResult, userActionRequired bool, maxDaysForDownload int32, ) *BankConnectionInterface`

NewBankConnectionInterface instantiates a new BankConnectionInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankConnectionInterfaceWithDefaults

`func NewBankConnectionInterfaceWithDefaults() *BankConnectionInterface`

NewBankConnectionInterfaceWithDefaults instantiates a new BankConnectionInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *BankConnectionInterface) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *BankConnectionInterface) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *BankConnectionInterface) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.


### GetLoginCredentials

`func (o *BankConnectionInterface) GetLoginCredentials() []LoginCredentialResource`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *BankConnectionInterface) GetLoginCredentialsOk() (*[]LoginCredentialResource, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *BankConnectionInterface) SetLoginCredentials(v []LoginCredentialResource)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetDefaultTwoStepProcedureId

`func (o *BankConnectionInterface) GetDefaultTwoStepProcedureId() string`

GetDefaultTwoStepProcedureId returns the DefaultTwoStepProcedureId field if non-nil, zero value otherwise.

### GetDefaultTwoStepProcedureIdOk

`func (o *BankConnectionInterface) GetDefaultTwoStepProcedureIdOk() (*string, bool)`

GetDefaultTwoStepProcedureIdOk returns a tuple with the DefaultTwoStepProcedureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTwoStepProcedureId

`func (o *BankConnectionInterface) SetDefaultTwoStepProcedureId(v string)`

SetDefaultTwoStepProcedureId sets DefaultTwoStepProcedureId field to given value.


### SetDefaultTwoStepProcedureIdNil

`func (o *BankConnectionInterface) SetDefaultTwoStepProcedureIdNil(b bool)`

 SetDefaultTwoStepProcedureIdNil sets the value for DefaultTwoStepProcedureId to be an explicit nil

### UnsetDefaultTwoStepProcedureId
`func (o *BankConnectionInterface) UnsetDefaultTwoStepProcedureId()`

UnsetDefaultTwoStepProcedureId ensures that no value is present for DefaultTwoStepProcedureId, not even an explicit nil
### GetTwoStepProcedures

`func (o *BankConnectionInterface) GetTwoStepProcedures() []TwoStepProcedure`

GetTwoStepProcedures returns the TwoStepProcedures field if non-nil, zero value otherwise.

### GetTwoStepProceduresOk

`func (o *BankConnectionInterface) GetTwoStepProceduresOk() (*[]TwoStepProcedure, bool)`

GetTwoStepProceduresOk returns a tuple with the TwoStepProcedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoStepProcedures

`func (o *BankConnectionInterface) SetTwoStepProcedures(v []TwoStepProcedure)`

SetTwoStepProcedures sets TwoStepProcedures field to given value.


### GetAisConsent

`func (o *BankConnectionInterface) GetAisConsent() BankConsent`

GetAisConsent returns the AisConsent field if non-nil, zero value otherwise.

### GetAisConsentOk

`func (o *BankConnectionInterface) GetAisConsentOk() (*BankConsent, bool)`

GetAisConsentOk returns a tuple with the AisConsent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAisConsent

`func (o *BankConnectionInterface) SetAisConsent(v BankConsent)`

SetAisConsent sets AisConsent field to given value.


### SetAisConsentNil

`func (o *BankConnectionInterface) SetAisConsentNil(b bool)`

 SetAisConsentNil sets the value for AisConsent to be an explicit nil

### UnsetAisConsent
`func (o *BankConnectionInterface) UnsetAisConsent()`

UnsetAisConsent ensures that no value is present for AisConsent, not even an explicit nil
### GetLastManualUpdate

`func (o *BankConnectionInterface) GetLastManualUpdate() UpdateResult`

GetLastManualUpdate returns the LastManualUpdate field if non-nil, zero value otherwise.

### GetLastManualUpdateOk

`func (o *BankConnectionInterface) GetLastManualUpdateOk() (*UpdateResult, bool)`

GetLastManualUpdateOk returns a tuple with the LastManualUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastManualUpdate

`func (o *BankConnectionInterface) SetLastManualUpdate(v UpdateResult)`

SetLastManualUpdate sets LastManualUpdate field to given value.


### SetLastManualUpdateNil

`func (o *BankConnectionInterface) SetLastManualUpdateNil(b bool)`

 SetLastManualUpdateNil sets the value for LastManualUpdate to be an explicit nil

### UnsetLastManualUpdate
`func (o *BankConnectionInterface) UnsetLastManualUpdate()`

UnsetLastManualUpdate ensures that no value is present for LastManualUpdate, not even an explicit nil
### GetLastAutoUpdate

`func (o *BankConnectionInterface) GetLastAutoUpdate() UpdateResult`

GetLastAutoUpdate returns the LastAutoUpdate field if non-nil, zero value otherwise.

### GetLastAutoUpdateOk

`func (o *BankConnectionInterface) GetLastAutoUpdateOk() (*UpdateResult, bool)`

GetLastAutoUpdateOk returns a tuple with the LastAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAutoUpdate

`func (o *BankConnectionInterface) SetLastAutoUpdate(v UpdateResult)`

SetLastAutoUpdate sets LastAutoUpdate field to given value.


### SetLastAutoUpdateNil

`func (o *BankConnectionInterface) SetLastAutoUpdateNil(b bool)`

 SetLastAutoUpdateNil sets the value for LastAutoUpdate to be an explicit nil

### UnsetLastAutoUpdate
`func (o *BankConnectionInterface) UnsetLastAutoUpdate()`

UnsetLastAutoUpdate ensures that no value is present for LastAutoUpdate, not even an explicit nil
### GetUserActionRequired

`func (o *BankConnectionInterface) GetUserActionRequired() bool`

GetUserActionRequired returns the UserActionRequired field if non-nil, zero value otherwise.

### GetUserActionRequiredOk

`func (o *BankConnectionInterface) GetUserActionRequiredOk() (*bool, bool)`

GetUserActionRequiredOk returns a tuple with the UserActionRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserActionRequired

`func (o *BankConnectionInterface) SetUserActionRequired(v bool)`

SetUserActionRequired sets UserActionRequired field to given value.


### GetMaxDaysForDownload

`func (o *BankConnectionInterface) GetMaxDaysForDownload() int32`

GetMaxDaysForDownload returns the MaxDaysForDownload field if non-nil, zero value otherwise.

### GetMaxDaysForDownloadOk

`func (o *BankConnectionInterface) GetMaxDaysForDownloadOk() (*int32, bool)`

GetMaxDaysForDownloadOk returns a tuple with the MaxDaysForDownload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDaysForDownload

`func (o *BankConnectionInterface) SetMaxDaysForDownload(v int32)`

SetMaxDaysForDownload sets MaxDaysForDownload field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


