# ExecutePasswordChangeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User identifier | 
**Password** | **string** | User&#39;s new password. Minimum length is 6, and maximum length is 128. | 
**PasswordChangeToken** | **string** | Decrypted password change token (the token that you received from the /users/requestPasswordChange service, decrypted with your data decryption key) | 

## Methods

### NewExecutePasswordChangeParams

`func NewExecutePasswordChangeParams(userId string, password string, passwordChangeToken string, ) *ExecutePasswordChangeParams`

NewExecutePasswordChangeParams instantiates a new ExecutePasswordChangeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutePasswordChangeParamsWithDefaults

`func NewExecutePasswordChangeParamsWithDefaults() *ExecutePasswordChangeParams`

NewExecutePasswordChangeParamsWithDefaults instantiates a new ExecutePasswordChangeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *ExecutePasswordChangeParams) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExecutePasswordChangeParams) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExecutePasswordChangeParams) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPassword

`func (o *ExecutePasswordChangeParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ExecutePasswordChangeParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ExecutePasswordChangeParams) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPasswordChangeToken

`func (o *ExecutePasswordChangeParams) GetPasswordChangeToken() string`

GetPasswordChangeToken returns the PasswordChangeToken field if non-nil, zero value otherwise.

### GetPasswordChangeTokenOk

`func (o *ExecutePasswordChangeParams) GetPasswordChangeTokenOk() (*string, bool)`

GetPasswordChangeTokenOk returns a tuple with the PasswordChangeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeToken

`func (o *ExecutePasswordChangeParams) SetPasswordChangeToken(v string)`

SetPasswordChangeToken sets PasswordChangeToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


