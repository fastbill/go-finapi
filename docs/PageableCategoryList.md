# PageableCategoryList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | [**[]Category**](Category.md) | &lt;strong&gt;Type:&lt;/strong&gt; Category&lt;br/&gt; Categories | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewPageableCategoryList

`func NewPageableCategoryList(categories []Category, paging Paging, ) *PageableCategoryList`

NewPageableCategoryList instantiates a new PageableCategoryList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableCategoryListWithDefaults

`func NewPageableCategoryListWithDefaults() *PageableCategoryList`

NewPageableCategoryListWithDefaults instantiates a new PageableCategoryList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *PageableCategoryList) GetCategories() []Category`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *PageableCategoryList) GetCategoriesOk() (*[]Category, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *PageableCategoryList) SetCategories(v []Category)`

SetCategories sets Categories field to given value.


### GetPaging

`func (o *PageableCategoryList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableCategoryList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableCategoryList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


