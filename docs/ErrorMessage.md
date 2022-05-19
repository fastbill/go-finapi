# ErrorMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ErrorDetails**](ErrorDetails.md) | &lt;strong&gt;Type:&lt;/strong&gt; ErrorDetails&lt;br/&gt; List of errors | 
**Date** | **string** | Server date of when the error(s) occurred, in the format YYYY-MM-DD HH:MM:SS.SSS | 
**RequestId** | **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the response header &#39;X-Request-Id&#39; instead.&lt;br/&gt;&lt;br/&gt;ID of the request that caused this error. This is either what you have passed for the header &#39;X-Request-Id&#39;, or an auto-generated ID in case you didn&#39;t pass any value for that header.&lt;br/&gt;&lt;br/&gt; | 
**Endpoint** | **string** | The service endpoint that was called | 
**AuthContext** | **string** | Information about the authorization context of the service call | 
**Bank** | **NullableString** | BLZ and name (in format \&quot;&lt;BLZ&gt; - &lt;name&gt;\&quot;) of a bank that was used for the original request | 

## Methods

### NewErrorMessage

`func NewErrorMessage(errors []ErrorDetails, date string, requestId string, endpoint string, authContext string, bank NullableString, ) *ErrorMessage`

NewErrorMessage instantiates a new ErrorMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorMessageWithDefaults

`func NewErrorMessageWithDefaults() *ErrorMessage`

NewErrorMessageWithDefaults instantiates a new ErrorMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ErrorMessage) GetErrors() []ErrorDetails`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorMessage) GetErrorsOk() (*[]ErrorDetails, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorMessage) SetErrors(v []ErrorDetails)`

SetErrors sets Errors field to given value.


### GetDate

`func (o *ErrorMessage) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ErrorMessage) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ErrorMessage) SetDate(v string)`

SetDate sets Date field to given value.


### GetRequestId

`func (o *ErrorMessage) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorMessage) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorMessage) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetEndpoint

`func (o *ErrorMessage) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ErrorMessage) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ErrorMessage) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAuthContext

`func (o *ErrorMessage) GetAuthContext() string`

GetAuthContext returns the AuthContext field if non-nil, zero value otherwise.

### GetAuthContextOk

`func (o *ErrorMessage) GetAuthContextOk() (*string, bool)`

GetAuthContextOk returns a tuple with the AuthContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthContext

`func (o *ErrorMessage) SetAuthContext(v string)`

SetAuthContext sets AuthContext field to given value.


### GetBank

`func (o *ErrorMessage) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *ErrorMessage) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *ErrorMessage) SetBank(v string)`

SetBank sets Bank field to given value.


### SetBankNil

`func (o *ErrorMessage) SetBankNil(b bool)`

 SetBankNil sets the value for Bank to be an explicit nil

### UnsetBank
`func (o *ErrorMessage) UnsetBank()`

UnsetBank ensures that no value is present for Bank, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


