# MockAccountData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | Account identifier | 
**AccountBalance** | **float64** | The balance that this account should be set to.&lt;br/&gt;&lt;br/&gt;&lt;b&gt;NOTE&lt;/b&gt;:&lt;br/&gt;&amp;bull; If the specified balance does not add up to the account&#39;s current balance plus the sum of the &#39;newTransactions&#39;, then finAPI will fix the balance deviation with the insertion of an adjusting entry (&#39;Zwischensaldo&#39; transaction).&lt;br/&gt;&amp;bull; This service is not calculating exchange rates for transactions, so if &#39;newTransactions&#39; contains any transactions with a currency different to the account&#39;s currency, then the balance deviation might get calculated incorrectly. | 
**NewTransactions** | Pointer to [**[]NewTransaction**](NewTransaction.md) | &lt;strong&gt;Type:&lt;/strong&gt; NewTransaction&lt;br/&gt; New transactions that should be imported into the account (at most 1000 transactions at once). Please make sure that the value you pass in the &#39;accountBalance&#39; field equals the current account balance plus the sum of the new transactions that you pass here, otherwise finAPI will detect a deviation in the balance and might add an adjusting entry (&#39;Zwischensaldo&#39; transaction). &lt;br/&gt;Please also note that it is not guaranteed that all transactions that you pass here will actually get imported. More specifically, finAPI will ignore any transactions whose booking date is prior to the date of the last successful account update minus 10 days (which is the default &#39;update window&#39; that finAPI uses when importing new transactions). Also, finAPI will ignore any transactions that are considered duplicates of already existing transactions within the update window. This is the case for instance when you try to import a new transaction with the same booking date and same amount as an already existing transaction. In such cases, you might get an adjusting entry too (&#39;Zwischensaldo&#39; transaction), as your given balance might not add up to the transactions that will exist in the account after the update. | [optional] 

## Methods

### NewMockAccountData

`func NewMockAccountData(accountId int64, accountBalance float64, ) *MockAccountData`

NewMockAccountData instantiates a new MockAccountData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMockAccountDataWithDefaults

`func NewMockAccountDataWithDefaults() *MockAccountData`

NewMockAccountDataWithDefaults instantiates a new MockAccountData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *MockAccountData) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *MockAccountData) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *MockAccountData) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetAccountBalance

`func (o *MockAccountData) GetAccountBalance() float64`

GetAccountBalance returns the AccountBalance field if non-nil, zero value otherwise.

### GetAccountBalanceOk

`func (o *MockAccountData) GetAccountBalanceOk() (*float64, bool)`

GetAccountBalanceOk returns a tuple with the AccountBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountBalance

`func (o *MockAccountData) SetAccountBalance(v float64)`

SetAccountBalance sets AccountBalance field to given value.


### GetNewTransactions

`func (o *MockAccountData) GetNewTransactions() []NewTransaction`

GetNewTransactions returns the NewTransactions field if non-nil, zero value otherwise.

### GetNewTransactionsOk

`func (o *MockAccountData) GetNewTransactionsOk() (*[]NewTransaction, bool)`

GetNewTransactionsOk returns a tuple with the NewTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTransactions

`func (o *MockAccountData) SetNewTransactions(v []NewTransaction)`

SetNewTransactions sets NewTransactions field to given value.

### HasNewTransactions

`func (o *MockAccountData) HasNewTransactions() bool`

HasNewTransactions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


