# BankConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Bank connection identifier | 
**BankId** | **int64** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;bank&#39; field instead.&lt;br/&gt;&lt;br/&gt;Identifier of the bank that this connection belongs to.  | 
**Name** | **NullableString** | Custom name for the bank connection. You can set this field with the &#39;Edit a bank connection&#39; service, as well as during the initial import of the bank connection. Maximum length is 64. | 
**BankingUserId** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; in the &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Stored online banking user ID credential. This field may be null for the &#39;demo connection&#39;. If your client has no license for processing banking credentials then a banking user ID will always be &#39;XXXXX&#39; | 
**BankingCustomerId** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; in the &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Stored online banking customer ID credential. If your client has no license for processing banking credentials or if this field contains a value that requires password protection (see field &#39;isCustomerIdPassword&#39; in Bank Resource) then the banking customer ID will always be &#39;XXXXX | 
**BankingPin** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; in the &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Stored online banking PIN. If a PIN is stored, this will always be &#39;XXXXX&#39; | 
**Type** | **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Bank connection type | 
**UpdateStatus** | **string** | Current status of data download (account balances and transactions/securities). The POST /bankConnections/import and POST /bankConnections/&lt;id&gt;/update services will set this flag to IN_PROGRESS before they return. Once the import or update has finished, the status will be changed to READY. | 
**CategorizationStatus** | [**CategorizationStatus**](CategorizationStatus.md) | &lt;strong&gt;Type:&lt;/strong&gt; CategorizationStatus&lt;br/&gt; Current status of transaction categorization. The asynchronous download process that is triggered by a call of the POST /bankConnections/import and POST /bankConnections/&lt;id&gt;/update services (and also by finAPI&#39;s auto update, if enabled) will set this flag to PENDING once the download has finished and a categorization is scheduled for the imported transactions. A separate categorization thread will then start to categorize the transactions (during this process, the status is IN_PROGRESS). When categorization has finished, the status will be (re-)set to READY. Note that the current categorization status should only be queried after the download has finished, i.e. once the download status has switched from IN_PROGRESS to READY. | 
**LastManualUpdate** | [**NullableUpdateResult**](UpdateResult.md) | &lt;strong&gt;Type:&lt;/strong&gt; UpdateResult&lt;br/&gt; THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the corresponding field in &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Result of the last manual update of this bank connection. If no manual update has ever been done so far, then this field will not be set. | 
**LastAutoUpdate** | [**NullableUpdateResult**](UpdateResult.md) | &lt;strong&gt;Type:&lt;/strong&gt; UpdateResult&lt;br/&gt; THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the corresponding field in &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Result of the last auto update of this bank connection (ran by finAPI&#39;s automatic batch update process). If no auto update has ever been done so far, then this field will not be set. | 
**IbanOnlyMoneyTransferSupported** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;capabilities&#39; field of the &#39;interfaces&#39; in the Account resource instead.&lt;br/&gt;&lt;br/&gt;Whether this bank connection accepts money transfer requests where the recipient&#39;s account is defined just by the IBAN (without an additional BIC). This field is re-evaluated each time this bank connection is updated. &lt;br/&gt;See also: /accounts/requestSepaMoneyTransfer | 
**IbanOnlyDirectDebitSupported** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;capabilities&#39; field of the &#39;interfaces&#39; in the Account resource instead.&lt;br/&gt;&lt;br/&gt;Whether this bank connection accepts direct debit requests where the debitor&#39;s account is defined just by the IBAN (without an additional BIC). This field is re-evaluated each time this bank connection is updated. &lt;br/&gt;See also: /accounts/requestSepaDirectDebit | 
**CollectiveMoneyTransferSupported** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;capabilities&#39; field of the &#39;interfaces&#39; in the Account resource instead.&lt;br/&gt;&lt;br/&gt;Whether this bank connection supports submitting collective money transfers. This field is re-evaluated each time this bank connection is updated. &lt;br/&gt;See also: /accounts/requestSepaMoneyTransfer | 
**DefaultTwoStepProcedureId** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the corresponding field in &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;The default two-step-procedure. Must match one of the available &#39;procedureId&#39;s from the &#39;twoStepProcedures&#39; list. When this field is set, then finAPI will automatically try to select the procedure wherever applicable. Note that the list of available procedures of a bank connection may change as a result of an update of the connection, and if this field references a procedure that is no longer available after an update, finAPI will automatically clear the default procedure (set it to null). | 
**TwoStepProcedures** | [**[]TwoStepProcedure**](TwoStepProcedure.md) | &lt;strong&gt;Type:&lt;/strong&gt; TwoStepProcedure&lt;br/&gt; THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the corresponding field in &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Available two-step-procedures for this bank connection, used for submitting a money transfer or direct debit request (see /accounts/requestSepaMoneyTransfer or /requestSepaDirectDebit). The available two-step-procedures are re-evaluated each time this bank connection is updated (/bankConnections/update). This means that this list may change as a result of an update. | 
**Interfaces** | [**[]BankConnectionInterface**](BankConnectionInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankConnectionInterface&lt;br/&gt; Set of interfaces that are connected for this bank connection. | 
**AccountIds** | **[]int64** | Identifiers of the accounts that belong to this bank connection | 
**Owners** | [**[]BankConnectionOwner**](BankConnectionOwner.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankConnectionOwner&lt;br/&gt; Information about the owner(s) of the bank connection | 
**Bank** | [**Bank**](Bank.md) | &lt;strong&gt;Type:&lt;/strong&gt; Bank&lt;br/&gt; Bank that this connection belongs to | 
**FurtherLoginNotRecommended** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;userActionRequired&#39; field of the BankConnectionInterface resource instead.&lt;br/&gt;&lt;br/&gt;This field indicates whether the last communication with the bank failed with an error that requires the user&#39;s attention or multi-step authentication error. If &#39;furtherLoginNotRecommended&#39; is true, finAPI will stop auto updates of this bank connection to mitigate the risk of the user&#39;s bank account getting locked by the bank. Every communication with the bank (via updates, money_transfers, direct debits. etc.) can change the value of this flag. If this field is true, we recommend the user to check his credentials and try a manual update of the bank connection. If the update is successful without any multi-step authentication error, the &#39;furtherLoginNotRecommended&#39; field will be set to false and the bank connection will be reincluded in the next batch update process. A manual update of the bank connection with incorrect credentials or if multi-step authentication error happens will set this field to true and lead to the exclusion of the bank connection from the following batch updates. | 

## Methods

### NewBankConnection

`func NewBankConnection(id int64, bankId int64, name NullableString, bankingUserId NullableString, bankingCustomerId NullableString, bankingPin NullableString, type_ string, updateStatus string, categorizationStatus CategorizationStatus, lastManualUpdate NullableUpdateResult, lastAutoUpdate NullableUpdateResult, ibanOnlyMoneyTransferSupported bool, ibanOnlyDirectDebitSupported bool, collectiveMoneyTransferSupported bool, defaultTwoStepProcedureId NullableString, twoStepProcedures []TwoStepProcedure, interfaces []BankConnectionInterface, accountIds []int64, owners []BankConnectionOwner, bank Bank, furtherLoginNotRecommended bool, ) *BankConnection`

NewBankConnection instantiates a new BankConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankConnectionWithDefaults

`func NewBankConnectionWithDefaults() *BankConnection`

NewBankConnectionWithDefaults instantiates a new BankConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BankConnection) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BankConnection) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BankConnection) SetId(v int64)`

SetId sets Id field to given value.


### GetBankId

`func (o *BankConnection) GetBankId() int64`

GetBankId returns the BankId field if non-nil, zero value otherwise.

### GetBankIdOk

`func (o *BankConnection) GetBankIdOk() (*int64, bool)`

GetBankIdOk returns a tuple with the BankId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankId

`func (o *BankConnection) SetBankId(v int64)`

SetBankId sets BankId field to given value.


### GetName

`func (o *BankConnection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BankConnection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BankConnection) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *BankConnection) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *BankConnection) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetBankingUserId

`func (o *BankConnection) GetBankingUserId() string`

GetBankingUserId returns the BankingUserId field if non-nil, zero value otherwise.

### GetBankingUserIdOk

`func (o *BankConnection) GetBankingUserIdOk() (*string, bool)`

GetBankingUserIdOk returns a tuple with the BankingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingUserId

`func (o *BankConnection) SetBankingUserId(v string)`

SetBankingUserId sets BankingUserId field to given value.


### SetBankingUserIdNil

`func (o *BankConnection) SetBankingUserIdNil(b bool)`

 SetBankingUserIdNil sets the value for BankingUserId to be an explicit nil

### UnsetBankingUserId
`func (o *BankConnection) UnsetBankingUserId()`

UnsetBankingUserId ensures that no value is present for BankingUserId, not even an explicit nil
### GetBankingCustomerId

`func (o *BankConnection) GetBankingCustomerId() string`

GetBankingCustomerId returns the BankingCustomerId field if non-nil, zero value otherwise.

### GetBankingCustomerIdOk

`func (o *BankConnection) GetBankingCustomerIdOk() (*string, bool)`

GetBankingCustomerIdOk returns a tuple with the BankingCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingCustomerId

`func (o *BankConnection) SetBankingCustomerId(v string)`

SetBankingCustomerId sets BankingCustomerId field to given value.


### SetBankingCustomerIdNil

`func (o *BankConnection) SetBankingCustomerIdNil(b bool)`

 SetBankingCustomerIdNil sets the value for BankingCustomerId to be an explicit nil

### UnsetBankingCustomerId
`func (o *BankConnection) UnsetBankingCustomerId()`

UnsetBankingCustomerId ensures that no value is present for BankingCustomerId, not even an explicit nil
### GetBankingPin

`func (o *BankConnection) GetBankingPin() string`

GetBankingPin returns the BankingPin field if non-nil, zero value otherwise.

### GetBankingPinOk

`func (o *BankConnection) GetBankingPinOk() (*string, bool)`

GetBankingPinOk returns a tuple with the BankingPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingPin

`func (o *BankConnection) SetBankingPin(v string)`

SetBankingPin sets BankingPin field to given value.


### SetBankingPinNil

`func (o *BankConnection) SetBankingPinNil(b bool)`

 SetBankingPinNil sets the value for BankingPin to be an explicit nil

### UnsetBankingPin
`func (o *BankConnection) UnsetBankingPin()`

UnsetBankingPin ensures that no value is present for BankingPin, not even an explicit nil
### GetType

`func (o *BankConnection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BankConnection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BankConnection) SetType(v string)`

SetType sets Type field to given value.


### GetUpdateStatus

`func (o *BankConnection) GetUpdateStatus() string`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *BankConnection) GetUpdateStatusOk() (*string, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *BankConnection) SetUpdateStatus(v string)`

SetUpdateStatus sets UpdateStatus field to given value.


### GetCategorizationStatus

`func (o *BankConnection) GetCategorizationStatus() CategorizationStatus`

GetCategorizationStatus returns the CategorizationStatus field if non-nil, zero value otherwise.

### GetCategorizationStatusOk

`func (o *BankConnection) GetCategorizationStatusOk() (*CategorizationStatus, bool)`

GetCategorizationStatusOk returns a tuple with the CategorizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorizationStatus

`func (o *BankConnection) SetCategorizationStatus(v CategorizationStatus)`

SetCategorizationStatus sets CategorizationStatus field to given value.


### GetLastManualUpdate

`func (o *BankConnection) GetLastManualUpdate() UpdateResult`

GetLastManualUpdate returns the LastManualUpdate field if non-nil, zero value otherwise.

### GetLastManualUpdateOk

`func (o *BankConnection) GetLastManualUpdateOk() (*UpdateResult, bool)`

GetLastManualUpdateOk returns a tuple with the LastManualUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastManualUpdate

`func (o *BankConnection) SetLastManualUpdate(v UpdateResult)`

SetLastManualUpdate sets LastManualUpdate field to given value.


### SetLastManualUpdateNil

`func (o *BankConnection) SetLastManualUpdateNil(b bool)`

 SetLastManualUpdateNil sets the value for LastManualUpdate to be an explicit nil

### UnsetLastManualUpdate
`func (o *BankConnection) UnsetLastManualUpdate()`

UnsetLastManualUpdate ensures that no value is present for LastManualUpdate, not even an explicit nil
### GetLastAutoUpdate

`func (o *BankConnection) GetLastAutoUpdate() UpdateResult`

GetLastAutoUpdate returns the LastAutoUpdate field if non-nil, zero value otherwise.

### GetLastAutoUpdateOk

`func (o *BankConnection) GetLastAutoUpdateOk() (*UpdateResult, bool)`

GetLastAutoUpdateOk returns a tuple with the LastAutoUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAutoUpdate

`func (o *BankConnection) SetLastAutoUpdate(v UpdateResult)`

SetLastAutoUpdate sets LastAutoUpdate field to given value.


### SetLastAutoUpdateNil

`func (o *BankConnection) SetLastAutoUpdateNil(b bool)`

 SetLastAutoUpdateNil sets the value for LastAutoUpdate to be an explicit nil

### UnsetLastAutoUpdate
`func (o *BankConnection) UnsetLastAutoUpdate()`

UnsetLastAutoUpdate ensures that no value is present for LastAutoUpdate, not even an explicit nil
### GetIbanOnlyMoneyTransferSupported

`func (o *BankConnection) GetIbanOnlyMoneyTransferSupported() bool`

GetIbanOnlyMoneyTransferSupported returns the IbanOnlyMoneyTransferSupported field if non-nil, zero value otherwise.

### GetIbanOnlyMoneyTransferSupportedOk

`func (o *BankConnection) GetIbanOnlyMoneyTransferSupportedOk() (*bool, bool)`

GetIbanOnlyMoneyTransferSupportedOk returns a tuple with the IbanOnlyMoneyTransferSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanOnlyMoneyTransferSupported

`func (o *BankConnection) SetIbanOnlyMoneyTransferSupported(v bool)`

SetIbanOnlyMoneyTransferSupported sets IbanOnlyMoneyTransferSupported field to given value.


### GetIbanOnlyDirectDebitSupported

`func (o *BankConnection) GetIbanOnlyDirectDebitSupported() bool`

GetIbanOnlyDirectDebitSupported returns the IbanOnlyDirectDebitSupported field if non-nil, zero value otherwise.

### GetIbanOnlyDirectDebitSupportedOk

`func (o *BankConnection) GetIbanOnlyDirectDebitSupportedOk() (*bool, bool)`

GetIbanOnlyDirectDebitSupportedOk returns a tuple with the IbanOnlyDirectDebitSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanOnlyDirectDebitSupported

`func (o *BankConnection) SetIbanOnlyDirectDebitSupported(v bool)`

SetIbanOnlyDirectDebitSupported sets IbanOnlyDirectDebitSupported field to given value.


### GetCollectiveMoneyTransferSupported

`func (o *BankConnection) GetCollectiveMoneyTransferSupported() bool`

GetCollectiveMoneyTransferSupported returns the CollectiveMoneyTransferSupported field if non-nil, zero value otherwise.

### GetCollectiveMoneyTransferSupportedOk

`func (o *BankConnection) GetCollectiveMoneyTransferSupportedOk() (*bool, bool)`

GetCollectiveMoneyTransferSupportedOk returns a tuple with the CollectiveMoneyTransferSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectiveMoneyTransferSupported

`func (o *BankConnection) SetCollectiveMoneyTransferSupported(v bool)`

SetCollectiveMoneyTransferSupported sets CollectiveMoneyTransferSupported field to given value.


### GetDefaultTwoStepProcedureId

`func (o *BankConnection) GetDefaultTwoStepProcedureId() string`

GetDefaultTwoStepProcedureId returns the DefaultTwoStepProcedureId field if non-nil, zero value otherwise.

### GetDefaultTwoStepProcedureIdOk

`func (o *BankConnection) GetDefaultTwoStepProcedureIdOk() (*string, bool)`

GetDefaultTwoStepProcedureIdOk returns a tuple with the DefaultTwoStepProcedureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTwoStepProcedureId

`func (o *BankConnection) SetDefaultTwoStepProcedureId(v string)`

SetDefaultTwoStepProcedureId sets DefaultTwoStepProcedureId field to given value.


### SetDefaultTwoStepProcedureIdNil

`func (o *BankConnection) SetDefaultTwoStepProcedureIdNil(b bool)`

 SetDefaultTwoStepProcedureIdNil sets the value for DefaultTwoStepProcedureId to be an explicit nil

### UnsetDefaultTwoStepProcedureId
`func (o *BankConnection) UnsetDefaultTwoStepProcedureId()`

UnsetDefaultTwoStepProcedureId ensures that no value is present for DefaultTwoStepProcedureId, not even an explicit nil
### GetTwoStepProcedures

`func (o *BankConnection) GetTwoStepProcedures() []TwoStepProcedure`

GetTwoStepProcedures returns the TwoStepProcedures field if non-nil, zero value otherwise.

### GetTwoStepProceduresOk

`func (o *BankConnection) GetTwoStepProceduresOk() (*[]TwoStepProcedure, bool)`

GetTwoStepProceduresOk returns a tuple with the TwoStepProcedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoStepProcedures

`func (o *BankConnection) SetTwoStepProcedures(v []TwoStepProcedure)`

SetTwoStepProcedures sets TwoStepProcedures field to given value.


### GetInterfaces

`func (o *BankConnection) GetInterfaces() []BankConnectionInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *BankConnection) GetInterfacesOk() (*[]BankConnectionInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *BankConnection) SetInterfaces(v []BankConnectionInterface)`

SetInterfaces sets Interfaces field to given value.


### GetAccountIds

`func (o *BankConnection) GetAccountIds() []int64`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *BankConnection) GetAccountIdsOk() (*[]int64, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *BankConnection) SetAccountIds(v []int64)`

SetAccountIds sets AccountIds field to given value.


### GetOwners

`func (o *BankConnection) GetOwners() []BankConnectionOwner`

GetOwners returns the Owners field if non-nil, zero value otherwise.

### GetOwnersOk

`func (o *BankConnection) GetOwnersOk() (*[]BankConnectionOwner, bool)`

GetOwnersOk returns a tuple with the Owners field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwners

`func (o *BankConnection) SetOwners(v []BankConnectionOwner)`

SetOwners sets Owners field to given value.


### SetOwnersNil

`func (o *BankConnection) SetOwnersNil(b bool)`

 SetOwnersNil sets the value for Owners to be an explicit nil

### UnsetOwners
`func (o *BankConnection) UnsetOwners()`

UnsetOwners ensures that no value is present for Owners, not even an explicit nil
### GetBank

`func (o *BankConnection) GetBank() Bank`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *BankConnection) GetBankOk() (*Bank, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *BankConnection) SetBank(v Bank)`

SetBank sets Bank field to given value.


### GetFurtherLoginNotRecommended

`func (o *BankConnection) GetFurtherLoginNotRecommended() bool`

GetFurtherLoginNotRecommended returns the FurtherLoginNotRecommended field if non-nil, zero value otherwise.

### GetFurtherLoginNotRecommendedOk

`func (o *BankConnection) GetFurtherLoginNotRecommendedOk() (*bool, bool)`

GetFurtherLoginNotRecommendedOk returns a tuple with the FurtherLoginNotRecommended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFurtherLoginNotRecommended

`func (o *BankConnection) SetFurtherLoginNotRecommended(v bool)`

SetFurtherLoginNotRecommended sets FurtherLoginNotRecommended field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


