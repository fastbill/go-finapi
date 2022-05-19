# PageableUserInfoList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]UserInfo**](UserInfo.md) | &lt;strong&gt;Type:&lt;/strong&gt; UserInfo&lt;br/&gt; List of users information | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewPageableUserInfoList

`func NewPageableUserInfoList(users []UserInfo, paging Paging, ) *PageableUserInfoList`

NewPageableUserInfoList instantiates a new PageableUserInfoList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableUserInfoListWithDefaults

`func NewPageableUserInfoListWithDefaults() *PageableUserInfoList`

NewPageableUserInfoListWithDefaults instantiates a new PageableUserInfoList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *PageableUserInfoList) GetUsers() []UserInfo`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *PageableUserInfoList) GetUsersOk() (*[]UserInfo, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *PageableUserInfoList) SetUsers(v []UserInfo)`

SetUsers sets Users field to given value.


### GetPaging

`func (o *PageableUserInfoList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableUserInfoList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableUserInfoList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


