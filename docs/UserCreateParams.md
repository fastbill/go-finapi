# UserCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | User&#39;s identifier. Max length is 36 symbols. Only the following symbols are allowed: 0-9, A-Z, a-z, -, _, ., +, @. If not specified, then a unique random value will be generated. | [optional] 
**Password** | Pointer to **string** | User&#39;s password. Minimum length is 6, and maximum length is 128. If not specified, then a unique random value will be generated. | [optional] 
**Email** | Pointer to **string** | User&#39;s email address. Maximum length is 320. | [optional] 
**Phone** | Pointer to **string** | User&#39;s phone number. Maximum length is 50. | [optional] 
**IsAutoUpdateEnabled** | Pointer to **bool** | Whether the user&#39;s bank connections will get updated in the course of finAPI&#39;s automatic batch update. Note that the automatic batch update will only update bank connections where all of the following applies:&lt;br&gt;&lt;br&gt; - a valid consent exists OR the PIN is stored in finAPI for the bank connection, and the related bank does not have volatile PINs (see the &#39;isVolatile&#39; flag of the &#39;loginCredentials&#39;)&lt;br/&gt; - the user has accepted the latest Terms and Conditions (this only applies to users whose mandator doesn&#39;t have an AIS license)&lt;br&gt; - the bank connection interface flag &#39;userActionRequired&#39; has to be false&lt;br&gt; - the previous update using the stored credentials did not fail due to the credentials being incorrect (or there was no previous update with the stored credentials)&lt;br&gt; - the bank that the bank connection relates to is included in the automatic batch update (please contact your Sys-Admin for details about the batch update configuration)&lt;br&gt;- at least one of the bank&#39;s supported data sources can be used by finAPI for your client (i.e.: if a bank supports only web scraping, but web scraping is disabled for your client, then bank connections of that bank will not get updated by the automatic batch update)&lt;br&gt;&lt;br&gt;Also note that the automatic batch update must generally be enabled for your client for this field to have any effect.&lt;br/&gt;&lt;br/&gt;WARNING: The automatic update will always download transactions and security positions for any account that it updates, even if the account was previously imported or updated with &#39;skipPositionsDownload&#39; &#x3D; true.&lt;br/&gt;&lt;br/&gt;Default value is false. | [optional] [default to false]

## Methods

### NewUserCreateParams

`func NewUserCreateParams() *UserCreateParams`

NewUserCreateParams instantiates a new UserCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateParamsWithDefaults

`func NewUserCreateParamsWithDefaults() *UserCreateParams`

NewUserCreateParamsWithDefaults instantiates a new UserCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserCreateParams) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserCreateParams) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserCreateParams) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserCreateParams) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPassword

`func (o *UserCreateParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreateParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreateParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreateParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmail

`func (o *UserCreateParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreateParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreateParams) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCreateParams) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *UserCreateParams) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UserCreateParams) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UserCreateParams) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UserCreateParams) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetIsAutoUpdateEnabled

`func (o *UserCreateParams) GetIsAutoUpdateEnabled() bool`

GetIsAutoUpdateEnabled returns the IsAutoUpdateEnabled field if non-nil, zero value otherwise.

### GetIsAutoUpdateEnabledOk

`func (o *UserCreateParams) GetIsAutoUpdateEnabledOk() (*bool, bool)`

GetIsAutoUpdateEnabledOk returns a tuple with the IsAutoUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoUpdateEnabled

`func (o *UserCreateParams) SetIsAutoUpdateEnabled(v bool)`

SetIsAutoUpdateEnabled sets IsAutoUpdateEnabled field to given value.

### HasIsAutoUpdateEnabled

`func (o *UserCreateParams) HasIsAutoUpdateEnabled() bool`

HasIsAutoUpdateEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


