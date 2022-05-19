# PageableTransactionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]Transaction**](Transaction.md) | &lt;strong&gt;Type:&lt;/strong&gt; Transaction&lt;br/&gt; Array of transactions (for the requested page) | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 
**Income** | **float64** | The total income of all transactions (across all pages) | 
**Spending** | **float64** | The total spending of all transactions (across all pages) | 
**Balance** | **float64** | The total sum of all transactions (across all pages) | 

## Methods

### NewPageableTransactionList

`func NewPageableTransactionList(transactions []Transaction, paging Paging, income float64, spending float64, balance float64, ) *PageableTransactionList`

NewPageableTransactionList instantiates a new PageableTransactionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableTransactionListWithDefaults

`func NewPageableTransactionListWithDefaults() *PageableTransactionList`

NewPageableTransactionListWithDefaults instantiates a new PageableTransactionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *PageableTransactionList) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *PageableTransactionList) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *PageableTransactionList) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.


### GetPaging

`func (o *PageableTransactionList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableTransactionList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableTransactionList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.


### GetIncome

`func (o *PageableTransactionList) GetIncome() float64`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *PageableTransactionList) GetIncomeOk() (*float64, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *PageableTransactionList) SetIncome(v float64)`

SetIncome sets Income field to given value.


### GetSpending

`func (o *PageableTransactionList) GetSpending() float64`

GetSpending returns the Spending field if non-nil, zero value otherwise.

### GetSpendingOk

`func (o *PageableTransactionList) GetSpendingOk() (*float64, bool)`

GetSpendingOk returns a tuple with the Spending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpending

`func (o *PageableTransactionList) SetSpending(v float64)`

SetSpending sets Spending field to given value.


### GetBalance

`func (o *PageableTransactionList) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *PageableTransactionList) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *PageableTransactionList) SetBalance(v float64)`

SetBalance sets Balance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


