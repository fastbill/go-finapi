# CashFlow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | [**NullableCategory**](Category.md) | &lt;strong&gt;Type:&lt;/strong&gt; Category&lt;br/&gt; Category of this cash flow. When null, then this is the cash flow of transactions that do not have a category. | 
**Income** | **float64** | The total calculated income for the given category | 
**Spending** | **float64** | The total calculated spending for the given category | 
**Balance** | **float64** | The calculated balance for the given category | 
**CountIncomeTransactions** | **int32** | The total count of income transactions for the given category | 
**CountSpendingTransactions** | **int32** | The total count of spending transactions for the given category | 
**CountAllTransactions** | **int32** | The total count of all transactions for the given category | 

## Methods

### NewCashFlow

`func NewCashFlow(category NullableCategory, income float64, spending float64, balance float64, countIncomeTransactions int32, countSpendingTransactions int32, countAllTransactions int32, ) *CashFlow`

NewCashFlow instantiates a new CashFlow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashFlowWithDefaults

`func NewCashFlowWithDefaults() *CashFlow`

NewCashFlowWithDefaults instantiates a new CashFlow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CashFlow) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CashFlow) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CashFlow) SetCategory(v Category)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *CashFlow) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CashFlow) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetIncome

`func (o *CashFlow) GetIncome() float64`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *CashFlow) GetIncomeOk() (*float64, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *CashFlow) SetIncome(v float64)`

SetIncome sets Income field to given value.


### GetSpending

`func (o *CashFlow) GetSpending() float64`

GetSpending returns the Spending field if non-nil, zero value otherwise.

### GetSpendingOk

`func (o *CashFlow) GetSpendingOk() (*float64, bool)`

GetSpendingOk returns a tuple with the Spending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpending

`func (o *CashFlow) SetSpending(v float64)`

SetSpending sets Spending field to given value.


### GetBalance

`func (o *CashFlow) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *CashFlow) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *CashFlow) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetCountIncomeTransactions

`func (o *CashFlow) GetCountIncomeTransactions() int32`

GetCountIncomeTransactions returns the CountIncomeTransactions field if non-nil, zero value otherwise.

### GetCountIncomeTransactionsOk

`func (o *CashFlow) GetCountIncomeTransactionsOk() (*int32, bool)`

GetCountIncomeTransactionsOk returns a tuple with the CountIncomeTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountIncomeTransactions

`func (o *CashFlow) SetCountIncomeTransactions(v int32)`

SetCountIncomeTransactions sets CountIncomeTransactions field to given value.


### GetCountSpendingTransactions

`func (o *CashFlow) GetCountSpendingTransactions() int32`

GetCountSpendingTransactions returns the CountSpendingTransactions field if non-nil, zero value otherwise.

### GetCountSpendingTransactionsOk

`func (o *CashFlow) GetCountSpendingTransactionsOk() (*int32, bool)`

GetCountSpendingTransactionsOk returns a tuple with the CountSpendingTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountSpendingTransactions

`func (o *CashFlow) SetCountSpendingTransactions(v int32)`

SetCountSpendingTransactions sets CountSpendingTransactions field to given value.


### GetCountAllTransactions

`func (o *CashFlow) GetCountAllTransactions() int32`

GetCountAllTransactions returns the CountAllTransactions field if non-nil, zero value otherwise.

### GetCountAllTransactionsOk

`func (o *CashFlow) GetCountAllTransactionsOk() (*int32, bool)`

GetCountAllTransactionsOk returns a tuple with the CountAllTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountAllTransactions

`func (o *CashFlow) SetCountAllTransactions(v int32)`

SetCountAllTransactions sets CountAllTransactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


