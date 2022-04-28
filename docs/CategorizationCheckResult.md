# CategorizationCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | The transaction identifier. The transactionId of the transaction that was passed to the service as input. This is not an actual ID of a stored transaction in finAPI, as the checkCategorization service doesn&#39;t store any data. | 
**Category** | [**NullableCategory**](Category.md) | &lt;strong&gt;Type:&lt;/strong&gt; Category&lt;br/&gt; A category. The determined transaction category for the given transactionId. This can be null, if the categorization algorithm fails to find a matching rule. | 

## Methods

### NewCategorizationCheckResult

`func NewCategorizationCheckResult(transactionId string, category NullableCategory, ) *CategorizationCheckResult`

NewCategorizationCheckResult instantiates a new CategorizationCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategorizationCheckResultWithDefaults

`func NewCategorizationCheckResultWithDefaults() *CategorizationCheckResult`

NewCategorizationCheckResultWithDefaults instantiates a new CategorizationCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *CategorizationCheckResult) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CategorizationCheckResult) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CategorizationCheckResult) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetCategory

`func (o *CategorizationCheckResult) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CategorizationCheckResult) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CategorizationCheckResult) SetCategory(v Category)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *CategorizationCheckResult) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CategorizationCheckResult) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


