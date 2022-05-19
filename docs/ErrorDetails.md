# ErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **NullableString** | Error message | 
**Code** | [**ErrorCode**](ErrorCode.md) | &lt;strong&gt;Type:&lt;/strong&gt; ErrorCode&lt;br/&gt; Error code. See the documentation of the individual services for details about what values may be returned. | 
**Type** | [**ErrorType**](ErrorType.md) | &lt;strong&gt;Type:&lt;/strong&gt; ErrorType&lt;br/&gt; Error type. BUSINESS errors depict error messages in the language of the bank (or the preferred language) for the user, e.g. from a bank server. TECHNICAL errors are meant to be read by developers and depict internal errors. | 
**MultiStepAuthentication** | [**NullableMultiStepAuthenticationChallenge**](MultiStepAuthenticationChallenge.md) | &lt;strong&gt;Type:&lt;/strong&gt; MultiStepAuthenticationChallenge&lt;br/&gt; This field is set when a multi-step authentication is required, i.e. when you need to repeat the original service call and provide additional data. The field contains information about what additional data is required. | 

## Methods

### NewErrorDetails

`func NewErrorDetails(message NullableString, code ErrorCode, type_ ErrorType, multiStepAuthentication NullableMultiStepAuthenticationChallenge, ) *ErrorDetails`

NewErrorDetails instantiates a new ErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDetailsWithDefaults

`func NewErrorDetailsWithDefaults() *ErrorDetails`

NewErrorDetailsWithDefaults instantiates a new ErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDetails) SetMessage(v string)`

SetMessage sets Message field to given value.


### SetMessageNil

`func (o *ErrorDetails) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ErrorDetails) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetCode

`func (o *ErrorDetails) GetCode() ErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorDetails) GetCodeOk() (*ErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorDetails) SetCode(v ErrorCode)`

SetCode sets Code field to given value.


### GetType

`func (o *ErrorDetails) GetType() ErrorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorDetails) GetTypeOk() (*ErrorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorDetails) SetType(v ErrorType)`

SetType sets Type field to given value.


### GetMultiStepAuthentication

`func (o *ErrorDetails) GetMultiStepAuthentication() MultiStepAuthenticationChallenge`

GetMultiStepAuthentication returns the MultiStepAuthentication field if non-nil, zero value otherwise.

### GetMultiStepAuthenticationOk

`func (o *ErrorDetails) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationChallenge, bool)`

GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStepAuthentication

`func (o *ErrorDetails) SetMultiStepAuthentication(v MultiStepAuthenticationChallenge)`

SetMultiStepAuthentication sets MultiStepAuthentication field to given value.


### SetMultiStepAuthenticationNil

`func (o *ErrorDetails) SetMultiStepAuthenticationNil(b bool)`

 SetMultiStepAuthenticationNil sets the value for MultiStepAuthentication to be an explicit nil

### UnsetMultiStepAuthentication
`func (o *ErrorDetails) UnsetMultiStepAuthentication()`

UnsetMultiStepAuthentication ensures that no value is present for MultiStepAuthentication, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


