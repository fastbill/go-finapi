# UpdateMultipleTransactionsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsNew** | Pointer to **bool** | Whether this transactions should be flagged as &#39;new&#39; or not. Any newly imported transaction will have this flag initially set to true. How you use this field is up to your interpretation. For example, you might want to set it to false once a user has clicked on/seen the transaction. | [optional] 
**IsPotentialDuplicate** | Pointer to **bool** | You can set this field only to &#39;false&#39;. finAPI marks transactions as a potential duplicates  when its internal duplicate detection algorithm is signaling so. Transactions that are flagged as duplicates can be deleted by the user. To prevent the user from deleting original transactions, which might lead to incorrect balances, it is not possible to manually set this flag to &#39;true&#39;. | [optional] 
**CategoryId** | Pointer to **int64** | Identifier of the new category to apply to the transaction. When updating the transaction&#39;s category, the category&#39;s fields &#39;id&#39;, &#39;name&#39;, &#39;parentId&#39;, &#39;parentName&#39;, and &#39;isCustom&#39; will all get updated. To clear the category for the transaction, the categoryId field must be passed with value 0. | [optional] 
**TrainCategorization** | Pointer to **bool** | This field is only regarded when the field &#39;categoryId&#39; is set. It controls whether finAPI&#39;s categorization system should learn from the given categorization(s). If set to &#39;true&#39;, then the user&#39;s categorization rules will be updated so that similar transactions will get categorized accordingly in future. If set to &#39;false&#39;, then the service will simply change the category of the given transaction(s), without updating the user&#39;s categorization rules. The field defaults to &#39;true&#39; if not specified. | [optional] [default to true]
**LabelIds** | Pointer to **[]int64** | Identifiers of labels to apply to the transaction. To clear transactions&#39; labels, pass an empty array of identifiers: &#39;[]&#39; | [optional] 
**Ids** | Pointer to **[]int64** | A comma-separated list of transaction identifiers. If specified, then only transactions whose identifier match any of the given identifiers will be regarded. The maximum number of identifiers is 100. | [optional] 
**AccountIds** | Pointer to **[]int64** | A comma-separated list of account identifiers. If specified, then only transactions whose account&#39;s identifier is included in this list will be regarded. | [optional] 

## Methods

### NewUpdateMultipleTransactionsParams

`func NewUpdateMultipleTransactionsParams() *UpdateMultipleTransactionsParams`

NewUpdateMultipleTransactionsParams instantiates a new UpdateMultipleTransactionsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMultipleTransactionsParamsWithDefaults

`func NewUpdateMultipleTransactionsParamsWithDefaults() *UpdateMultipleTransactionsParams`

NewUpdateMultipleTransactionsParamsWithDefaults instantiates a new UpdateMultipleTransactionsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsNew

`func (o *UpdateMultipleTransactionsParams) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *UpdateMultipleTransactionsParams) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *UpdateMultipleTransactionsParams) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *UpdateMultipleTransactionsParams) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.

### GetIsPotentialDuplicate

`func (o *UpdateMultipleTransactionsParams) GetIsPotentialDuplicate() bool`

GetIsPotentialDuplicate returns the IsPotentialDuplicate field if non-nil, zero value otherwise.

### GetIsPotentialDuplicateOk

`func (o *UpdateMultipleTransactionsParams) GetIsPotentialDuplicateOk() (*bool, bool)`

GetIsPotentialDuplicateOk returns a tuple with the IsPotentialDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPotentialDuplicate

`func (o *UpdateMultipleTransactionsParams) SetIsPotentialDuplicate(v bool)`

SetIsPotentialDuplicate sets IsPotentialDuplicate field to given value.

### HasIsPotentialDuplicate

`func (o *UpdateMultipleTransactionsParams) HasIsPotentialDuplicate() bool`

HasIsPotentialDuplicate returns a boolean if a field has been set.

### GetCategoryId

`func (o *UpdateMultipleTransactionsParams) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *UpdateMultipleTransactionsParams) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *UpdateMultipleTransactionsParams) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *UpdateMultipleTransactionsParams) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetTrainCategorization

`func (o *UpdateMultipleTransactionsParams) GetTrainCategorization() bool`

GetTrainCategorization returns the TrainCategorization field if non-nil, zero value otherwise.

### GetTrainCategorizationOk

`func (o *UpdateMultipleTransactionsParams) GetTrainCategorizationOk() (*bool, bool)`

GetTrainCategorizationOk returns a tuple with the TrainCategorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainCategorization

`func (o *UpdateMultipleTransactionsParams) SetTrainCategorization(v bool)`

SetTrainCategorization sets TrainCategorization field to given value.

### HasTrainCategorization

`func (o *UpdateMultipleTransactionsParams) HasTrainCategorization() bool`

HasTrainCategorization returns a boolean if a field has been set.

### GetLabelIds

`func (o *UpdateMultipleTransactionsParams) GetLabelIds() []int64`

GetLabelIds returns the LabelIds field if non-nil, zero value otherwise.

### GetLabelIdsOk

`func (o *UpdateMultipleTransactionsParams) GetLabelIdsOk() (*[]int64, bool)`

GetLabelIdsOk returns a tuple with the LabelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelIds

`func (o *UpdateMultipleTransactionsParams) SetLabelIds(v []int64)`

SetLabelIds sets LabelIds field to given value.

### HasLabelIds

`func (o *UpdateMultipleTransactionsParams) HasLabelIds() bool`

HasLabelIds returns a boolean if a field has been set.

### GetIds

`func (o *UpdateMultipleTransactionsParams) GetIds() []int64`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *UpdateMultipleTransactionsParams) GetIdsOk() (*[]int64, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *UpdateMultipleTransactionsParams) SetIds(v []int64)`

SetIds sets Ids field to given value.

### HasIds

`func (o *UpdateMultipleTransactionsParams) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetAccountIds

`func (o *UpdateMultipleTransactionsParams) GetAccountIds() []int64`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *UpdateMultipleTransactionsParams) GetAccountIdsOk() (*[]int64, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *UpdateMultipleTransactionsParams) SetAccountIds(v []int64)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *UpdateMultipleTransactionsParams) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


