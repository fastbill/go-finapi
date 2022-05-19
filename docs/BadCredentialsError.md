# BadCredentialsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Error type | 
**ErrorDescription** | **string** | Error description | 

## Methods

### NewBadCredentialsError

`func NewBadCredentialsError(error_ string, errorDescription string, ) *BadCredentialsError`

NewBadCredentialsError instantiates a new BadCredentialsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadCredentialsErrorWithDefaults

`func NewBadCredentialsErrorWithDefaults() *BadCredentialsError`

NewBadCredentialsErrorWithDefaults instantiates a new BadCredentialsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *BadCredentialsError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BadCredentialsError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BadCredentialsError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *BadCredentialsError) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *BadCredentialsError) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *BadCredentialsError) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


