# Account

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Account identifier | 
**BankConnectionId** | **int64** | Identifier of the bank connection that this account belongs to | 
**AccountName** | **NullableString** | Account name | 
**Iban** | **NullableString** | Account&#39;s IBAN. Note that this field can change from &#39;null&#39; to a value - or vice versa - any time when the account is being updated. This is subject to changes within the bank&#39;s internal account management. | 
**AccountNumber** | **string** | (National) account number. Note that this value might change whenever the account is updated (for example, leading zeros might be added or removed). | 
**SubAccountNumber** | **NullableString** | Account&#39;s sub-account-number. Note that this field can change from &#39;null&#39; to a value - or vice versa - any time when the account is being updated. This is subject to changes within the bank&#39;s internal account management. | 
**AccountHolderName** | **NullableString** | Name of the account holder | 
**AccountHolderId** | **NullableString** | Bank&#39;s internal identification of the account holder. Note that if your client has no license for processing this field, it will always be &#39;XXXXX&#39; | 
**AccountCurrency** | **NullableString** | Account&#39;s currency | 
**AccountTypeId** | **int64** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;accountType&#39; field instead.&lt;br/&gt;&lt;br/&gt;Identifier of the account&#39;s type. Note that, in general, the type of an account can change any time when the account is being updated. This is subject to changes within the bank&#39;s internal account management. However, if the account&#39;s type has previously been changed explicitly (via the PATCH method), then the explicitly set type will NOT be automatically changed anymore, even if the type has changed on the bank side. &lt;br/&gt;1 &#x3D; Checking,&lt;br/&gt;2 &#x3D; Savings,&lt;br/&gt;3 &#x3D; CreditCard,&lt;br/&gt;4 &#x3D; Security,&lt;br/&gt;5 &#x3D; Loan,&lt;br/&gt;6 &#x3D; Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),&lt;br/&gt;7 &#x3D; Membership,&lt;br/&gt;8 &#x3D; Bausparen&lt;br/&gt; | 
**AccountTypeName** | **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;accountType&#39; field instead.&lt;br/&gt;&lt;br/&gt;Name of the account&#39;s type. | 
**AccountType** | [**AccountType**](AccountType.md) | &lt;strong&gt;Type:&lt;/strong&gt; AccountType&lt;br/&gt; An account type.&lt;br/&gt;&lt;br/&gt;Checking,&lt;br/&gt;Savings,&lt;br/&gt;CreditCard,&lt;br/&gt;Security,&lt;br/&gt;Loan,&lt;br/&gt;Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),&lt;br/&gt;Membership,&lt;br/&gt;Bausparen&lt;br/&gt;&lt;br/&gt; | 
**Balance** | **NullableFloat64** | Current account balance | 
**Overdraft** | **NullableFloat64** | Current overdraft | 
**OverdraftLimit** | **NullableFloat64** | Overdraft limit | 
**AvailableFunds** | **NullableFloat64** | Current available funds. Note that this field is only set if finAPI can make a definite statement about the current available funds. This might not always be the case, for example if there is not enough information available about the overdraft limit and current overdraft. | 
**LastSuccessfulUpdate** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the corresponding field in &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Timestamp of when the account was last successfully updated (or initially imported); more precisely: time when the account data (balance and positions) has been stored into the finAPI databases. The value is returned in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**LastUpdateAttempt** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the corresponding field in &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Timestamp of when the account was last tried to be updated (or initially imported); more precisely: time when the update (or initial import) was triggered. The value is returned in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**IsNew** | **bool** | Indicating whether this account is &#39;new&#39; or not. Any newly imported account will have this flag initially set to true, and remain so until you set it to false (see PATCH /accounts/&lt;id&gt;). How you use this field is up to your interpretation, however it is recommended to set the flag to false for all accounts right after the initial import of the bank connection. This way, you will be able recognize accounts that get newly imported during a later update of the bank connection, by checking for any accounts with the flag set to true right after an update. | 
**Status** | [**AccountStatus**](AccountStatus.md) | &lt;strong&gt;Type:&lt;/strong&gt; AccountStatus&lt;br/&gt; THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;status&#39; field in &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;The current status of the account. Possible values are:&lt;br/&gt;&amp;bull; &lt;code&gt;UPDATED&lt;/code&gt; means that the account is up to date from finAPI&#39;s point of view. This means that no current import/update is running, and the previous import/update could successfully update the account&#39;s data (e.g. transactions and securities), and the bank given balance matched the transaction&#39;s calculated sum, meaning that no adjusting entry (&#39;Zwischensaldo&#39; transaction) was inserted.&lt;br/&gt;&amp;bull; &lt;code&gt;UPDATED_FIXED&lt;/code&gt; means that the account is up to date from finAPI&#39;s point of view (no current import/update is running, and the previous import/update could successfully update the account&#39;s data), BUT there was a deviation in the bank given balance which was fixed by adding an adjusting entry (&#39;Zwischensaldo&#39; transaction).&lt;br/&gt;&amp;bull; &lt;code&gt;DOWNLOAD_IN_PROGRESS&lt;/code&gt; means that the account&#39;s data is currently being imported/updated.&lt;br/&gt;&amp;bull; &lt;code&gt;DOWNLOAD_FAILED&lt;/code&gt; means that the account data could not get successfully imported or updated. Possible reasons: finAPI could not get the account&#39;s balance, or it could not parse all transactions/securities, or some internal error has occurred. Also, it could mean that finAPI could not even get to the point of receiving the account data from the bank server, for example because of incorrect login credentials or a network problem. Note however that when we get a balance and just an empty list of transactions or securities, then this is regarded as valid and successful download. The reason for this is that for some accounts that have little activity, we may actually get no recent transactions but only a balance.&lt;br/&gt;&amp;bull; &lt;code&gt;DEPRECATED&lt;/code&gt; means that the account could no longer get matched with any account from the bank server. This can mean either that the account was terminated by the user and is no longer sent by the bank server, or that finAPI could no longer match it because the account&#39;s data (name, type, iban, account number, etc.) has been changed by the bank. | 
**SupportedOrders** | [**[]SupportedOrder**](SupportedOrder.md) |  | 
**Interfaces** | [**[]AccountInterface**](AccountInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; AccountInterface&lt;br/&gt; Set of interfaces to which this account is connected | 
**ClearingAccounts** | [**[]ClearingAccountData**](ClearingAccountData.md) | &lt;strong&gt;Type:&lt;/strong&gt; ClearingAccountData&lt;br/&gt; THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;&lt;br/&gt;List of clearing accounts that relate to this account. Clearing accounts can be used for money transfers (see field &#39;clearingAccountId&#39; of the &#39;Request SEPA Money Transfer&#39; service). | 
**IsSeized** | **bool** | Whether this account is seized. Note that this information is not received from the bank, but determined by finAPI based on the available account information. | 

## Methods

### NewAccount

`func NewAccount(id int64, bankConnectionId int64, accountName NullableString, iban NullableString, accountNumber string, subAccountNumber NullableString, accountHolderName NullableString, accountHolderId NullableString, accountCurrency NullableString, accountTypeId int64, accountTypeName string, accountType AccountType, balance NullableFloat64, overdraft NullableFloat64, overdraftLimit NullableFloat64, availableFunds NullableFloat64, lastSuccessfulUpdate NullableString, lastUpdateAttempt NullableString, isNew bool, status AccountStatus, supportedOrders []SupportedOrder, interfaces []AccountInterface, clearingAccounts []ClearingAccountData, isSeized bool, ) *Account`

NewAccount instantiates a new Account object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithDefaults

`func NewAccountWithDefaults() *Account`

NewAccountWithDefaults instantiates a new Account object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Account) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Account) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Account) SetId(v int64)`

SetId sets Id field to given value.


### GetBankConnectionId

`func (o *Account) GetBankConnectionId() int64`

GetBankConnectionId returns the BankConnectionId field if non-nil, zero value otherwise.

### GetBankConnectionIdOk

`func (o *Account) GetBankConnectionIdOk() (*int64, bool)`

GetBankConnectionIdOk returns a tuple with the BankConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionId

`func (o *Account) SetBankConnectionId(v int64)`

SetBankConnectionId sets BankConnectionId field to given value.


### GetAccountName

`func (o *Account) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *Account) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *Account) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.


### SetAccountNameNil

`func (o *Account) SetAccountNameNil(b bool)`

 SetAccountNameNil sets the value for AccountName to be an explicit nil

### UnsetAccountName
`func (o *Account) UnsetAccountName()`

UnsetAccountName ensures that no value is present for AccountName, not even an explicit nil
### GetIban

`func (o *Account) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Account) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Account) SetIban(v string)`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *Account) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *Account) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetAccountNumber

`func (o *Account) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Account) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Account) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.


### GetSubAccountNumber

`func (o *Account) GetSubAccountNumber() string`

GetSubAccountNumber returns the SubAccountNumber field if non-nil, zero value otherwise.

### GetSubAccountNumberOk

`func (o *Account) GetSubAccountNumberOk() (*string, bool)`

GetSubAccountNumberOk returns a tuple with the SubAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccountNumber

`func (o *Account) SetSubAccountNumber(v string)`

SetSubAccountNumber sets SubAccountNumber field to given value.


### SetSubAccountNumberNil

`func (o *Account) SetSubAccountNumberNil(b bool)`

 SetSubAccountNumberNil sets the value for SubAccountNumber to be an explicit nil

### UnsetSubAccountNumber
`func (o *Account) UnsetSubAccountNumber()`

UnsetSubAccountNumber ensures that no value is present for SubAccountNumber, not even an explicit nil
### GetAccountHolderName

`func (o *Account) GetAccountHolderName() string`

GetAccountHolderName returns the AccountHolderName field if non-nil, zero value otherwise.

### GetAccountHolderNameOk

`func (o *Account) GetAccountHolderNameOk() (*string, bool)`

GetAccountHolderNameOk returns a tuple with the AccountHolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderName

`func (o *Account) SetAccountHolderName(v string)`

SetAccountHolderName sets AccountHolderName field to given value.


### SetAccountHolderNameNil

`func (o *Account) SetAccountHolderNameNil(b bool)`

 SetAccountHolderNameNil sets the value for AccountHolderName to be an explicit nil

### UnsetAccountHolderName
`func (o *Account) UnsetAccountHolderName()`

UnsetAccountHolderName ensures that no value is present for AccountHolderName, not even an explicit nil
### GetAccountHolderId

`func (o *Account) GetAccountHolderId() string`

GetAccountHolderId returns the AccountHolderId field if non-nil, zero value otherwise.

### GetAccountHolderIdOk

`func (o *Account) GetAccountHolderIdOk() (*string, bool)`

GetAccountHolderIdOk returns a tuple with the AccountHolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolderId

`func (o *Account) SetAccountHolderId(v string)`

SetAccountHolderId sets AccountHolderId field to given value.


### SetAccountHolderIdNil

`func (o *Account) SetAccountHolderIdNil(b bool)`

 SetAccountHolderIdNil sets the value for AccountHolderId to be an explicit nil

### UnsetAccountHolderId
`func (o *Account) UnsetAccountHolderId()`

UnsetAccountHolderId ensures that no value is present for AccountHolderId, not even an explicit nil
### GetAccountCurrency

`func (o *Account) GetAccountCurrency() string`

GetAccountCurrency returns the AccountCurrency field if non-nil, zero value otherwise.

### GetAccountCurrencyOk

`func (o *Account) GetAccountCurrencyOk() (*string, bool)`

GetAccountCurrencyOk returns a tuple with the AccountCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountCurrency

`func (o *Account) SetAccountCurrency(v string)`

SetAccountCurrency sets AccountCurrency field to given value.


### SetAccountCurrencyNil

`func (o *Account) SetAccountCurrencyNil(b bool)`

 SetAccountCurrencyNil sets the value for AccountCurrency to be an explicit nil

### UnsetAccountCurrency
`func (o *Account) UnsetAccountCurrency()`

UnsetAccountCurrency ensures that no value is present for AccountCurrency, not even an explicit nil
### GetAccountTypeId

`func (o *Account) GetAccountTypeId() int64`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *Account) GetAccountTypeIdOk() (*int64, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *Account) SetAccountTypeId(v int64)`

SetAccountTypeId sets AccountTypeId field to given value.


### GetAccountTypeName

`func (o *Account) GetAccountTypeName() string`

GetAccountTypeName returns the AccountTypeName field if non-nil, zero value otherwise.

### GetAccountTypeNameOk

`func (o *Account) GetAccountTypeNameOk() (*string, bool)`

GetAccountTypeNameOk returns a tuple with the AccountTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeName

`func (o *Account) SetAccountTypeName(v string)`

SetAccountTypeName sets AccountTypeName field to given value.


### GetAccountType

`func (o *Account) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Account) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Account) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.


### GetBalance

`func (o *Account) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Account) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Account) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### SetBalanceNil

`func (o *Account) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *Account) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetOverdraft

`func (o *Account) GetOverdraft() float64`

GetOverdraft returns the Overdraft field if non-nil, zero value otherwise.

### GetOverdraftOk

`func (o *Account) GetOverdraftOk() (*float64, bool)`

GetOverdraftOk returns a tuple with the Overdraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraft

`func (o *Account) SetOverdraft(v float64)`

SetOverdraft sets Overdraft field to given value.


### SetOverdraftNil

`func (o *Account) SetOverdraftNil(b bool)`

 SetOverdraftNil sets the value for Overdraft to be an explicit nil

### UnsetOverdraft
`func (o *Account) UnsetOverdraft()`

UnsetOverdraft ensures that no value is present for Overdraft, not even an explicit nil
### GetOverdraftLimit

`func (o *Account) GetOverdraftLimit() float64`

GetOverdraftLimit returns the OverdraftLimit field if non-nil, zero value otherwise.

### GetOverdraftLimitOk

`func (o *Account) GetOverdraftLimitOk() (*float64, bool)`

GetOverdraftLimitOk returns a tuple with the OverdraftLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverdraftLimit

`func (o *Account) SetOverdraftLimit(v float64)`

SetOverdraftLimit sets OverdraftLimit field to given value.


### SetOverdraftLimitNil

`func (o *Account) SetOverdraftLimitNil(b bool)`

 SetOverdraftLimitNil sets the value for OverdraftLimit to be an explicit nil

### UnsetOverdraftLimit
`func (o *Account) UnsetOverdraftLimit()`

UnsetOverdraftLimit ensures that no value is present for OverdraftLimit, not even an explicit nil
### GetAvailableFunds

`func (o *Account) GetAvailableFunds() float64`

GetAvailableFunds returns the AvailableFunds field if non-nil, zero value otherwise.

### GetAvailableFundsOk

`func (o *Account) GetAvailableFundsOk() (*float64, bool)`

GetAvailableFundsOk returns a tuple with the AvailableFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFunds

`func (o *Account) SetAvailableFunds(v float64)`

SetAvailableFunds sets AvailableFunds field to given value.


### SetAvailableFundsNil

`func (o *Account) SetAvailableFundsNil(b bool)`

 SetAvailableFundsNil sets the value for AvailableFunds to be an explicit nil

### UnsetAvailableFunds
`func (o *Account) UnsetAvailableFunds()`

UnsetAvailableFunds ensures that no value is present for AvailableFunds, not even an explicit nil
### GetLastSuccessfulUpdate

`func (o *Account) GetLastSuccessfulUpdate() string`

GetLastSuccessfulUpdate returns the LastSuccessfulUpdate field if non-nil, zero value otherwise.

### GetLastSuccessfulUpdateOk

`func (o *Account) GetLastSuccessfulUpdateOk() (*string, bool)`

GetLastSuccessfulUpdateOk returns a tuple with the LastSuccessfulUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulUpdate

`func (o *Account) SetLastSuccessfulUpdate(v string)`

SetLastSuccessfulUpdate sets LastSuccessfulUpdate field to given value.


### SetLastSuccessfulUpdateNil

`func (o *Account) SetLastSuccessfulUpdateNil(b bool)`

 SetLastSuccessfulUpdateNil sets the value for LastSuccessfulUpdate to be an explicit nil

### UnsetLastSuccessfulUpdate
`func (o *Account) UnsetLastSuccessfulUpdate()`

UnsetLastSuccessfulUpdate ensures that no value is present for LastSuccessfulUpdate, not even an explicit nil
### GetLastUpdateAttempt

`func (o *Account) GetLastUpdateAttempt() string`

GetLastUpdateAttempt returns the LastUpdateAttempt field if non-nil, zero value otherwise.

### GetLastUpdateAttemptOk

`func (o *Account) GetLastUpdateAttemptOk() (*string, bool)`

GetLastUpdateAttemptOk returns a tuple with the LastUpdateAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAttempt

`func (o *Account) SetLastUpdateAttempt(v string)`

SetLastUpdateAttempt sets LastUpdateAttempt field to given value.


### SetLastUpdateAttemptNil

`func (o *Account) SetLastUpdateAttemptNil(b bool)`

 SetLastUpdateAttemptNil sets the value for LastUpdateAttempt to be an explicit nil

### UnsetLastUpdateAttempt
`func (o *Account) UnsetLastUpdateAttempt()`

UnsetLastUpdateAttempt ensures that no value is present for LastUpdateAttempt, not even an explicit nil
### GetIsNew

`func (o *Account) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *Account) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *Account) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.


### GetStatus

`func (o *Account) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Account) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Account) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.


### GetSupportedOrders

`func (o *Account) GetSupportedOrders() []SupportedOrder`

GetSupportedOrders returns the SupportedOrders field if non-nil, zero value otherwise.

### GetSupportedOrdersOk

`func (o *Account) GetSupportedOrdersOk() (*[]SupportedOrder, bool)`

GetSupportedOrdersOk returns a tuple with the SupportedOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOrders

`func (o *Account) SetSupportedOrders(v []SupportedOrder)`

SetSupportedOrders sets SupportedOrders field to given value.


### GetInterfaces

`func (o *Account) GetInterfaces() []AccountInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Account) GetInterfacesOk() (*[]AccountInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Account) SetInterfaces(v []AccountInterface)`

SetInterfaces sets Interfaces field to given value.


### GetClearingAccounts

`func (o *Account) GetClearingAccounts() []ClearingAccountData`

GetClearingAccounts returns the ClearingAccounts field if non-nil, zero value otherwise.

### GetClearingAccountsOk

`func (o *Account) GetClearingAccountsOk() (*[]ClearingAccountData, bool)`

GetClearingAccountsOk returns a tuple with the ClearingAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingAccounts

`func (o *Account) SetClearingAccounts(v []ClearingAccountData)`

SetClearingAccounts sets ClearingAccounts field to given value.


### GetIsSeized

`func (o *Account) GetIsSeized() bool`

GetIsSeized returns the IsSeized field if non-nil, zero value otherwise.

### GetIsSeizedOk

`func (o *Account) GetIsSeizedOk() (*bool, bool)`

GetIsSeizedOk returns a tuple with the IsSeized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSeized

`func (o *Account) SetIsSeized(v bool)`

SetIsSeized sets IsSeized field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


