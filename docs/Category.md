# Category

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Category identifier.&lt;br/&gt;&lt;br/&gt;NOTE: Do NOT assume that the identifiers of the global finAPI categories are the same across different finAPI environments. In fact, the identifiers may change whenever a new finAPI version is released, even within the same environment. The identifiers are meant to be used for references within the finAPI services only, but not for hard-coding them in your application. If you need to hard-code the usage of a certain global category within your application, please instead refer to the category name. Also, please make sure to check the &#39;isCustom&#39; flag, which is false for all global categories (if you are not regarding this flag, you might end up referring to a user-specific category, and not the global category). | 
**Name** | **string** | Category name | 
**ParentId** | **NullableInt64** | Identifier of the parent category (if a parent category exists) | 
**ParentName** | **NullableString** | Name of the parent category (if a parent category exists) | 
**IsCustom** | **bool** | Whether the category is a finAPI global category (in which case this field will be false), or the category was created by a user (in which case this field will be true) | 
**Children** | **[]int64** | List of sub-categories identifiers (if any exist) | 

## Methods

### NewCategory

`func NewCategory(id int64, name string, parentId NullableInt64, parentName NullableString, isCustom bool, children []int64, ) *Category`

NewCategory instantiates a new Category object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryWithDefaults

`func NewCategoryWithDefaults() *Category`

NewCategoryWithDefaults instantiates a new Category object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Category) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Category) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Category) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Category) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Category) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Category) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *Category) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Category) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Category) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### SetParentIdNil

`func (o *Category) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *Category) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetParentName

`func (o *Category) GetParentName() string`

GetParentName returns the ParentName field if non-nil, zero value otherwise.

### GetParentNameOk

`func (o *Category) GetParentNameOk() (*string, bool)`

GetParentNameOk returns a tuple with the ParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentName

`func (o *Category) SetParentName(v string)`

SetParentName sets ParentName field to given value.


### SetParentNameNil

`func (o *Category) SetParentNameNil(b bool)`

 SetParentNameNil sets the value for ParentName to be an explicit nil

### UnsetParentName
`func (o *Category) UnsetParentName()`

UnsetParentName ensures that no value is present for ParentName, not even an explicit nil
### GetIsCustom

`func (o *Category) GetIsCustom() bool`

GetIsCustom returns the IsCustom field if non-nil, zero value otherwise.

### GetIsCustomOk

`func (o *Category) GetIsCustomOk() (*bool, bool)`

GetIsCustomOk returns a tuple with the IsCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustom

`func (o *Category) SetIsCustom(v bool)`

SetIsCustom sets IsCustom field to given value.


### GetChildren

`func (o *Category) GetChildren() []int64`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Category) GetChildrenOk() (*[]int64, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Category) SetChildren(v []int64)`

SetChildren sets Children field to given value.


### SetChildrenNil

`func (o *Category) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *Category) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


