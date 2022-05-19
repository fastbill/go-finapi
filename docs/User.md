# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | User identifier | 
**Password** | **string** | User&#39;s password. Please note that some services may return a distorted password (always &#39;XXXXX&#39;). See the documentation of individual services to find out whether the password is returned as plain text or as &#39;XXXXX&#39;. | 
**Email** | **NullableString** | User&#39;s email address | 
**Phone** | **NullableString** | User&#39;s phone number | 
**IsAutoUpdateEnabled** | **bool** | Whether the user&#39;s bank connections will get updated in the course of finAPI&#39;s automatic batch update. Note that the automatic batch update will only update bank connections where all of the following applies:&lt;br&gt;&lt;br&gt; - a valid consent exists OR the PIN is stored in finAPI for the bank connection, and the related bank does not have volatile PINs (see the &#39;isVolatile&#39; flag of the &#39;loginCredentials&#39;)&lt;br/&gt; - the user has accepted the latest Terms and Conditions (this only applies to users whose mandator doesn&#39;t have an AIS license)&lt;br&gt; - the bank connection interface flag &#39;userActionRequired&#39; has to be false&lt;br&gt; - the previous update using the stored credentials did not fail due to the credentials being incorrect (or there was no previous update with the stored credentials)&lt;br&gt; - the bank that the bank connection relates to is included in the automatic batch update (please contact your Sys-Admin for details about the batch update configuration)&lt;br&gt;- at least one of the bank&#39;s supported data sources can be used by finAPI for your client (i.e.: if a bank supports only web scraping, but web scraping is disabled for your client, then bank connections of that bank will not get updated by the automatic batch update)&lt;br&gt;&lt;br&gt;Also note that the automatic batch update must generally be enabled for your client for this field to have any effect.&lt;br/&gt;&lt;br/&gt;WARNING: The automatic update will always download transactions and security positions for any account that it updates, even if the account was previously imported or updated with &#39;skipPositionsDownload&#39; &#x3D; true. | 

## Methods

### NewUser

`func NewUser(id string, password string, email NullableString, phone NullableString, isAutoUpdateEnabled bool, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *User) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *User) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *User) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *User) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *User) SetPhone(v string)`

SetPhone sets Phone field to given value.


### SetPhoneNil

`func (o *User) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *User) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetIsAutoUpdateEnabled

`func (o *User) GetIsAutoUpdateEnabled() bool`

GetIsAutoUpdateEnabled returns the IsAutoUpdateEnabled field if non-nil, zero value otherwise.

### GetIsAutoUpdateEnabledOk

`func (o *User) GetIsAutoUpdateEnabledOk() (*bool, bool)`

GetIsAutoUpdateEnabledOk returns a tuple with the IsAutoUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpdateEnabled

`func (o *User) SetIsAutoUpdateEnabled(v bool)`

SetIsAutoUpdateEnabled sets IsAutoUpdateEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


