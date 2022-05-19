# BankInterfaceLoginField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Contains a German label for the input field that you should provide to the user. Also, these labels are used to identify login fields on the API-level, when you pass credentials to the service. | 
**IsSecret** | **bool** | Whether this field has to be treated as a secret. In this case your application should use a password input field instead of a cleartext field. | 
**IsVolatile** | **bool** | This field depicts whether the given credential is volatile. If a field is volatile, it means that the value for the field, which is provided by the user, is valid only for a single authentication and then gets invalidated on bank-side. If a login credential field is volatile, it is not possible to store it in finAPI, as stored credentials will never work for future authentications. | 

## Methods

### NewBankInterfaceLoginField

`func NewBankInterfaceLoginField(label string, isSecret bool, isVolatile bool, ) *BankInterfaceLoginField`

NewBankInterfaceLoginField instantiates a new BankInterfaceLoginField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankInterfaceLoginFieldWithDefaults

`func NewBankInterfaceLoginFieldWithDefaults() *BankInterfaceLoginField`

NewBankInterfaceLoginFieldWithDefaults instantiates a new BankInterfaceLoginField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *BankInterfaceLoginField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BankInterfaceLoginField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BankInterfaceLoginField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetIsSecret

`func (o *BankInterfaceLoginField) GetIsSecret() bool`

GetIsSecret returns the IsSecret field if non-nil, zero value otherwise.

### GetIsSecretOk

`func (o *BankInterfaceLoginField) GetIsSecretOk() (*bool, bool)`

GetIsSecretOk returns a tuple with the IsSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSecret

`func (o *BankInterfaceLoginField) SetIsSecret(v bool)`

SetIsSecret sets IsSecret field to given value.


### GetIsVolatile

`func (o *BankInterfaceLoginField) GetIsVolatile() bool`

GetIsVolatile returns the IsVolatile field if non-nil, zero value otherwise.

### GetIsVolatileOk

`func (o *BankInterfaceLoginField) GetIsVolatileOk() (*bool, bool)`

GetIsVolatileOk returns a tuple with the IsVolatile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVolatile

`func (o *BankInterfaceLoginField) SetIsVolatile(v bool)`

SetIsVolatile sets IsVolatile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


