# TransactionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Transactions** | [**[]Transaction**](Transaction.md) | &lt;strong&gt;Type:&lt;/strong&gt; Transaction&lt;br/&gt; List of transactions | 

## Methods

### NewTransactionList

`func NewTransactionList(transactions []Transaction, ) *TransactionList`

NewTransactionList instantiates a new TransactionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionListWithDefaults

`func NewTransactionListWithDefaults() *TransactionList`

NewTransactionListWithDefaults instantiates a new TransactionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactions

`func (o *TransactionList) GetTransactions() []Transaction`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *TransactionList) GetTransactionsOk() (*[]Transaction, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *TransactionList) SetTransactions(v []Transaction)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


