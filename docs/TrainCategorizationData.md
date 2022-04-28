# TrainCategorizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionData** | [**[]TrainCategorizationTransactionData**](TrainCategorizationTransactionData.md) | &lt;strong&gt;Type:&lt;/strong&gt; TrainCategorizationTransactionData&lt;br/&gt; Set of transaction data (at most 100 transactions at once) | 
**CategoryId** | **int64** | Category identifier | 

## Methods

### NewTrainCategorizationData

`func NewTrainCategorizationData(transactionData []TrainCategorizationTransactionData, categoryId int64, ) *TrainCategorizationData`

NewTrainCategorizationData instantiates a new TrainCategorizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainCategorizationDataWithDefaults

`func NewTrainCategorizationDataWithDefaults() *TrainCategorizationData`

NewTrainCategorizationDataWithDefaults instantiates a new TrainCategorizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionData

`func (o *TrainCategorizationData) GetTransactionData() []TrainCategorizationTransactionData`

GetTransactionData returns the TransactionData field if non-nil, zero value otherwise.

### GetTransactionDataOk

`func (o *TrainCategorizationData) GetTransactionDataOk() (*[]TrainCategorizationTransactionData, bool)`

GetTransactionDataOk returns a tuple with the TransactionData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionData

`func (o *TrainCategorizationData) SetTransactionData(v []TrainCategorizationTransactionData)`

SetTransactionData sets TransactionData field to given value.


### GetCategoryId

`func (o *TrainCategorizationData) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *TrainCategorizationData) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *TrainCategorizationData) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


