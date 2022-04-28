# DailyBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Date in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**Balance** | **float64** | Calculated balance at the end of day, across all regarded accounts. Note that the balance may not always add up to the income/spending of the day. This happens when a bank reports a balance that includes transactions which the bank didn&#39;t (yet) deliver. In any case, it is recommended to rely on the balance rather than on calculations based on the income/spending. | 
**Income** | **float64** | The sum of income of the given day, based on the &#39;transactions&#39;, across all regarded accounts. | 
**Spending** | **float64** | The sum of spending of the given day, based on the &#39;transactions&#39;, across all regarded accounts. Note that this is an absolute (i.e. positive) value. | 
**InternalAdjustingEntries** | **float64** | Sometimes finAPI may detect deviations between the bank reported account balance and the set of transactions received from the bank. This is an expected behaviour when an account has not been updated for a while, as banks provide only a limited history of transactions. In such cases, finAPI adds an adjusting entry (see the field Transaction.isAdjustingEntry), which will be contained in the &#39;transactions&#39; list, just as any other transaction.&lt;br/&gt;&lt;br/&gt;However, if an account was regularly updated and gaps in the transaction history are not expected, then finAPI will fix such deviations by adding an internal adjusting entry. These internal entries are not visible in the API and will not be contained in the &#39;transactions&#39; list, and thus also not regarded for the calculations of &#39;income&#39; and &#39;spending&#39;. They are however regarded for the calculation of the &#39;balance&#39;.&lt;br/&gt;&lt;br/&gt;As long as you don&#39;t do your own balance calculations, you do not need to regard this field here; The &#39;balance&#39; will always be correct. But if you do your own calculations, then you should not only regard the &#39;income&#39; and &#39;spending&#39;, but this field as well.&lt;br/&gt;&lt;br/&gt;Note that unlike the &#39;income&#39; and &#39;spending&#39;, this field can have a positive or negative value. | 
**Transactions** | **[]int64** | Identifiers of the transactions for the given day | 

## Methods

### NewDailyBalance

`func NewDailyBalance(date string, balance float64, income float64, spending float64, internalAdjustingEntries float64, transactions []int64, ) *DailyBalance`

NewDailyBalance instantiates a new DailyBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyBalanceWithDefaults

`func NewDailyBalanceWithDefaults() *DailyBalance`

NewDailyBalanceWithDefaults instantiates a new DailyBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DailyBalance) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DailyBalance) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DailyBalance) SetDate(v string)`

SetDate sets Date field to given value.


### GetBalance

`func (o *DailyBalance) GetBalance() float64`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *DailyBalance) GetBalanceOk() (*float64, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *DailyBalance) SetBalance(v float64)`

SetBalance sets Balance field to given value.


### GetIncome

`func (o *DailyBalance) GetIncome() float64`

GetIncome returns the Income field if non-nil, zero value otherwise.

### GetIncomeOk

`func (o *DailyBalance) GetIncomeOk() (*float64, bool)`

GetIncomeOk returns a tuple with the Income field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncome

`func (o *DailyBalance) SetIncome(v float64)`

SetIncome sets Income field to given value.


### GetSpending

`func (o *DailyBalance) GetSpending() float64`

GetSpending returns the Spending field if non-nil, zero value otherwise.

### GetSpendingOk

`func (o *DailyBalance) GetSpendingOk() (*float64, bool)`

GetSpendingOk returns a tuple with the Spending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpending

`func (o *DailyBalance) SetSpending(v float64)`

SetSpending sets Spending field to given value.


### GetInternalAdjustingEntries

`func (o *DailyBalance) GetInternalAdjustingEntries() float64`

GetInternalAdjustingEntries returns the InternalAdjustingEntries field if non-nil, zero value otherwise.

### GetInternalAdjustingEntriesOk

`func (o *DailyBalance) GetInternalAdjustingEntriesOk() (*float64, bool)`

GetInternalAdjustingEntriesOk returns a tuple with the InternalAdjustingEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalAdjustingEntries

`func (o *DailyBalance) SetInternalAdjustingEntries(v float64)`

SetInternalAdjustingEntries sets InternalAdjustingEntries field to given value.


### GetTransactions

`func (o *DailyBalance) GetTransactions() []int64`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *DailyBalance) GetTransactionsOk() (*[]int64, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *DailyBalance) SetTransactions(v []int64)`

SetTransactions sets Transactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


