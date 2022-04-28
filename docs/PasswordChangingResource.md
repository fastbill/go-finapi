# PasswordChangingResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User identifier | 
**UserEmail** | **NullableString** | User&#39;s email, encrypted. Decrypt with your data decryption key. If the user has no email set, then this field will be null. | 
**PasswordChangeToken** | **string** | Encrypted password change token. Decrypt this token with your data decryption key, and pass the decrypted token to the /users/executePasswordChange service to set a new password for the user. | 

## Methods

### NewPasswordChangingResource

`func NewPasswordChangingResource(userId string, userEmail NullableString, passwordChangeToken string, ) *PasswordChangingResource`

NewPasswordChangingResource instantiates a new PasswordChangingResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordChangingResourceWithDefaults

`func NewPasswordChangingResourceWithDefaults() *PasswordChangingResource`

NewPasswordChangingResourceWithDefaults instantiates a new PasswordChangingResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *PasswordChangingResource) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PasswordChangingResource) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PasswordChangingResource) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetUserEmail

`func (o *PasswordChangingResource) GetUserEmail() string`

GetUserEmail returns the UserEmail field if non-nil, zero value otherwise.

### GetUserEmailOk

`func (o *PasswordChangingResource) GetUserEmailOk() (*string, bool)`

GetUserEmailOk returns a tuple with the UserEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEmail

`func (o *PasswordChangingResource) SetUserEmail(v string)`

SetUserEmail sets UserEmail field to given value.


### SetUserEmailNil

`func (o *PasswordChangingResource) SetUserEmailNil(b bool)`

 SetUserEmailNil sets the value for UserEmail to be an explicit nil

### UnsetUserEmail
`func (o *PasswordChangingResource) UnsetUserEmail()`

UnsetUserEmail ensures that no value is present for UserEmail, not even an explicit nil
### GetPasswordChangeToken

`func (o *PasswordChangingResource) GetPasswordChangeToken() string`

GetPasswordChangeToken returns the PasswordChangeToken field if non-nil, zero value otherwise.

### GetPasswordChangeTokenOk

`func (o *PasswordChangingResource) GetPasswordChangeTokenOk() (*string, bool)`

GetPasswordChangeTokenOk returns a tuple with the PasswordChangeToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordChangeToken

`func (o *PasswordChangingResource) SetPasswordChangeToken(v string)`

SetPasswordChangeToken sets PasswordChangeToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


