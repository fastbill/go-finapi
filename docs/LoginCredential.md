# LoginCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The login field label, as defined by the respective bank interface. | 
**Value** | **string** | The value for the login field | 

## Methods

### NewLoginCredential

`func NewLoginCredential(label string, value string, ) *LoginCredential`

NewLoginCredential instantiates a new LoginCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginCredentialWithDefaults

`func NewLoginCredentialWithDefaults() *LoginCredential`

NewLoginCredentialWithDefaults instantiates a new LoginCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *LoginCredential) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LoginCredential) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LoginCredential) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *LoginCredential) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LoginCredential) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LoginCredential) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


