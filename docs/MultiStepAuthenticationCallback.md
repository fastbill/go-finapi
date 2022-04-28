# MultiStepAuthenticationCallback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Hash that was returned in the previous multi-step authentication error. | 
**ChallengeResponse** | Pointer to **string** | Challenge response. Must be set when the previous multi-step authentication error had status &#39;CHALLENGE_RESPONSE_REQUIRED&#39;. | [optional] 
**TwoStepProcedureId** | Pointer to **string** | The bank-given ID of the two-step-procedure that should be used for authentication. Must be set when the previous multi-step authentication error had status &#39;TWO_STEP_PROCEDURE_REQUIRED&#39;. | [optional] 
**RedirectCallback** | Pointer to **string** | Must be passed when the previous multi-step authentication error had status &#39;REDIRECT_REQUIRED&#39;. The value must consist of the complete query parameter list that was contained in the received redirect from the bank. | [optional] 
**DecoupledCallback** | Pointer to **bool** | Must be passed when the previous multi-step authentication error had status &#39;DECOUPLED_AUTH_REQUIRED&#39; or &#39;DECOUPLED_AUTH_IN_PROGRESS&#39;. The field represents the state of the decoupled authentication meaning that when it&#39;s set to &#39;true&#39;, the end-user has completed the authentication process on bank&#39;s side.&lt;br/&gt;&lt;br/&gt;Please note: Don&#39;t repeat the service call too frequently. Some banks limit the amount of requests per minute. Our suggestion is to repeat the service call for the decoupled approach every 5 seconds. | [optional] 

## Methods

### NewMultiStepAuthenticationCallback

`func NewMultiStepAuthenticationCallback(hash string, ) *MultiStepAuthenticationCallback`

NewMultiStepAuthenticationCallback instantiates a new MultiStepAuthenticationCallback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiStepAuthenticationCallbackWithDefaults

`func NewMultiStepAuthenticationCallbackWithDefaults() *MultiStepAuthenticationCallback`

NewMultiStepAuthenticationCallbackWithDefaults instantiates a new MultiStepAuthenticationCallback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *MultiStepAuthenticationCallback) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *MultiStepAuthenticationCallback) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *MultiStepAuthenticationCallback) SetHash(v string)`

SetHash sets Hash field to given value.


### GetChallengeResponse

`func (o *MultiStepAuthenticationCallback) GetChallengeResponse() string`

GetChallengeResponse returns the ChallengeResponse field if non-nil, zero value otherwise.

### GetChallengeResponseOk

`func (o *MultiStepAuthenticationCallback) GetChallengeResponseOk() (*string, bool)`

GetChallengeResponseOk returns a tuple with the ChallengeResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeResponse

`func (o *MultiStepAuthenticationCallback) SetChallengeResponse(v string)`

SetChallengeResponse sets ChallengeResponse field to given value.

### HasChallengeResponse

`func (o *MultiStepAuthenticationCallback) HasChallengeResponse() bool`

HasChallengeResponse returns a boolean if a field has been set.

### GetTwoStepProcedureId

`func (o *MultiStepAuthenticationCallback) GetTwoStepProcedureId() string`

GetTwoStepProcedureId returns the TwoStepProcedureId field if non-nil, zero value otherwise.

### GetTwoStepProcedureIdOk

`func (o *MultiStepAuthenticationCallback) GetTwoStepProcedureIdOk() (*string, bool)`

GetTwoStepProcedureIdOk returns a tuple with the TwoStepProcedureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoStepProcedureId

`func (o *MultiStepAuthenticationCallback) SetTwoStepProcedureId(v string)`

SetTwoStepProcedureId sets TwoStepProcedureId field to given value.

### HasTwoStepProcedureId

`func (o *MultiStepAuthenticationCallback) HasTwoStepProcedureId() bool`

HasTwoStepProcedureId returns a boolean if a field has been set.

### GetRedirectCallback

`func (o *MultiStepAuthenticationCallback) GetRedirectCallback() string`

GetRedirectCallback returns the RedirectCallback field if non-nil, zero value otherwise.

### GetRedirectCallbackOk

`func (o *MultiStepAuthenticationCallback) GetRedirectCallbackOk() (*string, bool)`

GetRedirectCallbackOk returns a tuple with the RedirectCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectCallback

`func (o *MultiStepAuthenticationCallback) SetRedirectCallback(v string)`

SetRedirectCallback sets RedirectCallback field to given value.

### HasRedirectCallback

`func (o *MultiStepAuthenticationCallback) HasRedirectCallback() bool`

HasRedirectCallback returns a boolean if a field has been set.

### GetDecoupledCallback

`func (o *MultiStepAuthenticationCallback) GetDecoupledCallback() bool`

GetDecoupledCallback returns the DecoupledCallback field if non-nil, zero value otherwise.

### GetDecoupledCallbackOk

`func (o *MultiStepAuthenticationCallback) GetDecoupledCallbackOk() (*bool, bool)`

GetDecoupledCallbackOk returns a tuple with the DecoupledCallback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecoupledCallback

`func (o *MultiStepAuthenticationCallback) SetDecoupledCallback(v bool)`

SetDecoupledCallback sets DecoupledCallback field to given value.

### HasDecoupledCallback

`func (o *MultiStepAuthenticationCallback) HasDecoupledCallback() bool`

HasDecoupledCallback returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


