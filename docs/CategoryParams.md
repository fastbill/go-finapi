# CategoryParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the category. Maximum length is 128. | 
**ParentId** | Pointer to **int64** | Identifier of the parent category, if the new category should be created as a sub-category. Must point to a main category in this case. If the new category should be a main category itself, this field must remain unset. | [optional] 

## Methods

### NewCategoryParams

`func NewCategoryParams(name string, ) *CategoryParams`

NewCategoryParams instantiates a new CategoryParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCategoryParamsWithDefaults

`func NewCategoryParamsWithDefaults() *CategoryParams`

NewCategoryParamsWithDefaults instantiates a new CategoryParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CategoryParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CategoryParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CategoryParams) SetName(v string)`

SetName sets Name field to given value.


### GetParentId

`func (o *CategoryParams) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CategoryParams) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CategoryParams) SetParentId(v int64)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CategoryParams) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


