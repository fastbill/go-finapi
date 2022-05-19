# UpdateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | [**UpdateResultStatus**](UpdateResultStatus.md) | &lt;strong&gt;Type:&lt;/strong&gt; UpdateResultStatus&lt;br/&gt; Note that &#39;OK&#39; just means that finAPI could successfully connect and log in to the bank server. However, it does not necessarily mean that all accounts could be updated successfully. For the latter, please refer to the &#39;status&#39; field of the Account resource. | 
**ErrorMessage** | **NullableString** | In case the update result is not &lt;code&gt;OK&lt;/code&gt;, this field may contain an error message with details about why the update failed (it is not guaranteed that a message is available though). In case the update result is &lt;code&gt;OK&lt;/code&gt;, the field will always be null. | 
**ErrorType** | [**NullableErrorType**](ErrorType.md) | &lt;strong&gt;Type:&lt;/strong&gt; ErrorType&lt;br/&gt; In case the update result is not &lt;code&gt;OK&lt;/code&gt;, this field contains the type of the error that occurred. BUSINESS means that the bank server responded with a non-technical error message for the user. TECHNICAL means that some internal error has occurred in finAPI or at the bank server. | 
**Timestamp** | **string** | Time of the update. The value is returned in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 

## Methods

### NewUpdateResult

`func NewUpdateResult(result UpdateResultStatus, errorMessage NullableString, errorType NullableErrorType, timestamp string, ) *UpdateResult`

NewUpdateResult instantiates a new UpdateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateResultWithDefaults

`func NewUpdateResultWithDefaults() *UpdateResult`

NewUpdateResultWithDefaults instantiates a new UpdateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *UpdateResult) GetResult() UpdateResultStatus`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *UpdateResult) GetResultOk() (*UpdateResultStatus, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *UpdateResult) SetResult(v UpdateResultStatus)`

SetResult sets Result field to given value.


### GetErrorMessage

`func (o *UpdateResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *UpdateResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *UpdateResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *UpdateResult) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *UpdateResult) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetErrorType

`func (o *UpdateResult) GetErrorType() ErrorType`

GetErrorType returns the ErrorType field if non-nil, zero value otherwise.

### GetErrorTypeOk

`func (o *UpdateResult) GetErrorTypeOk() (*ErrorType, bool)`

GetErrorTypeOk returns a tuple with the ErrorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorType

`func (o *UpdateResult) SetErrorType(v ErrorType)`

SetErrorType sets ErrorType field to given value.


### SetErrorTypeNil

`func (o *UpdateResult) SetErrorTypeNil(b bool)`

 SetErrorTypeNil sets the value for ErrorType to be an explicit nil

### UnsetErrorType
`func (o *UpdateResult) UnsetErrorType()`

UnsetErrorType ensures that no value is present for ErrorType, not even an explicit nil
### GetTimestamp

`func (o *UpdateResult) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UpdateResult) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UpdateResult) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


