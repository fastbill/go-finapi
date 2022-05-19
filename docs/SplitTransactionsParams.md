# SplitTransactionsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubTransactions** | [**[]SubTransactionParams**](SubTransactionParams.md) | &lt;strong&gt;Type:&lt;/strong&gt; SubTransactionParams&lt;br/&gt; List of sub-transactions | 

## Methods

### NewSplitTransactionsParams

`func NewSplitTransactionsParams(subTransactions []SubTransactionParams, ) *SplitTransactionsParams`

NewSplitTransactionsParams instantiates a new SplitTransactionsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplitTransactionsParamsWithDefaults

`func NewSplitTransactionsParamsWithDefaults() *SplitTransactionsParams`

NewSplitTransactionsParamsWithDefaults instantiates a new SplitTransactionsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubTransactions

`func (o *SplitTransactionsParams) GetSubTransactions() []SubTransactionParams`

GetSubTransactions returns the SubTransactions field if non-nil, zero value otherwise.

### GetSubTransactionsOk

`func (o *SplitTransactionsParams) GetSubTransactionsOk() (*[]SubTransactionParams, bool)`

GetSubTransactionsOk returns a tuple with the SubTransactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTransactions

`func (o *SplitTransactionsParams) SetSubTransactions(v []SubTransactionParams)`

SetSubTransactions sets SubTransactions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


