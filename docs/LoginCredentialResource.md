# LoginCredentialResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label for this login credential, as we suggest to show it to the user. | 
**Value** | **NullableString** | Stored value for this login credential. Please NOTE:&lt;br/&gt;If your client has no license for processing banking credentials, or if this field contains a value that requires password protection (e.g. PIN), then this field will always be &#39;XXXXX&#39;. | 

## Methods

### NewLoginCredentialResource

`func NewLoginCredentialResource(label string, value NullableString, ) *LoginCredentialResource`

NewLoginCredentialResource instantiates a new LoginCredentialResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginCredentialResourceWithDefaults

`func NewLoginCredentialResourceWithDefaults() *LoginCredentialResource`

NewLoginCredentialResourceWithDefaults instantiates a new LoginCredentialResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *LoginCredentialResource) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LoginCredentialResource) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LoginCredentialResource) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValue

`func (o *LoginCredentialResource) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *LoginCredentialResource) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *LoginCredentialResource) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *LoginCredentialResource) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *LoginCredentialResource) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


