# UserIdentifiersList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeletedUsers** | **[]string** |  | 
**NonDeletedUsers** | **[]string** |  | 

## Methods

### NewUserIdentifiersList

`func NewUserIdentifiersList(deletedUsers []string, nonDeletedUsers []string, ) *UserIdentifiersList`

NewUserIdentifiersList instantiates a new UserIdentifiersList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentifiersListWithDefaults

`func NewUserIdentifiersListWithDefaults() *UserIdentifiersList`

NewUserIdentifiersListWithDefaults instantiates a new UserIdentifiersList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletedUsers

`func (o *UserIdentifiersList) GetDeletedUsers() []string`

GetDeletedUsers returns the DeletedUsers field if non-nil, zero value otherwise.

### GetDeletedUsersOk

`func (o *UserIdentifiersList) GetDeletedUsersOk() (*[]string, bool)`

GetDeletedUsersOk returns a tuple with the DeletedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedUsers

`func (o *UserIdentifiersList) SetDeletedUsers(v []string)`

SetDeletedUsers sets DeletedUsers field to given value.


### GetNonDeletedUsers

`func (o *UserIdentifiersList) GetNonDeletedUsers() []string`

GetNonDeletedUsers returns the NonDeletedUsers field if non-nil, zero value otherwise.

### GetNonDeletedUsersOk

`func (o *UserIdentifiersList) GetNonDeletedUsersOk() (*[]string, bool)`

GetNonDeletedUsersOk returns a tuple with the NonDeletedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonDeletedUsers

`func (o *UserIdentifiersList) SetNonDeletedUsers(v []string)`

SetNonDeletedUsers sets NonDeletedUsers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


