# PaymentExecutionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessMessage** | **NullableString** | Technical message from the bank server, confirming the success of the request. Typically, you would not want to present this message to the user. Note that this field may not be set. However if it is not set, it does not necessarily mean that there was an error in processing the request. | 
**WarnMessage** | **NullableString** | In some cases, a bank server may accept the requested order, but return a warn message. This message may be of technical nature, but could also be of interest to the user. | 
**PaymentId** | **int64** | Payment identifier. Can be used to retrieve the status of the payment (see &#39;Get payments&#39; service). | 

## Methods

### NewPaymentExecutionResponse

`func NewPaymentExecutionResponse(successMessage NullableString, warnMessage NullableString, paymentId int64, ) *PaymentExecutionResponse`

NewPaymentExecutionResponse instantiates a new PaymentExecutionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentExecutionResponseWithDefaults

`func NewPaymentExecutionResponseWithDefaults() *PaymentExecutionResponse`

NewPaymentExecutionResponseWithDefaults instantiates a new PaymentExecutionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessMessage

`func (o *PaymentExecutionResponse) GetSuccessMessage() string`

GetSuccessMessage returns the SuccessMessage field if non-nil, zero value otherwise.

### GetSuccessMessageOk

`func (o *PaymentExecutionResponse) GetSuccessMessageOk() (*string, bool)`

GetSuccessMessageOk returns a tuple with the SuccessMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessMessage

`func (o *PaymentExecutionResponse) SetSuccessMessage(v string)`

SetSuccessMessage sets SuccessMessage field to given value.


### SetSuccessMessageNil

`func (o *PaymentExecutionResponse) SetSuccessMessageNil(b bool)`

 SetSuccessMessageNil sets the value for SuccessMessage to be an explicit nil

### UnsetSuccessMessage
`func (o *PaymentExecutionResponse) UnsetSuccessMessage()`

UnsetSuccessMessage ensures that no value is present for SuccessMessage, not even an explicit nil
### GetWarnMessage

`func (o *PaymentExecutionResponse) GetWarnMessage() string`

GetWarnMessage returns the WarnMessage field if non-nil, zero value otherwise.

### GetWarnMessageOk

`func (o *PaymentExecutionResponse) GetWarnMessageOk() (*string, bool)`

GetWarnMessageOk returns a tuple with the WarnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnMessage

`func (o *PaymentExecutionResponse) SetWarnMessage(v string)`

SetWarnMessage sets WarnMessage field to given value.


### SetWarnMessageNil

`func (o *PaymentExecutionResponse) SetWarnMessageNil(b bool)`

 SetWarnMessageNil sets the value for WarnMessage to be an explicit nil

### UnsetWarnMessage
`func (o *PaymentExecutionResponse) UnsetWarnMessage()`

UnsetWarnMessage ensures that no value is present for WarnMessage, not even an explicit nil
### GetPaymentId

`func (o *PaymentExecutionResponse) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaymentExecutionResponse) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaymentExecutionResponse) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


