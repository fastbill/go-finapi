# TwoStepProcedure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcedureId** | **string** | Bank-given ID of the procedure | 
**ProcedureName** | **string** | Bank-given name of the procedure | 
**ProcedureChallengeType** | **NullableString** | The challenge type of the procedure. Possible values are:&lt;br/&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;TEXT&lt;/code&gt; - the challenge will be a text that contains instructions for the user on how to proceed with the authorization.&lt;br/&gt;&amp;bull; &lt;code&gt;PHOTO&lt;/code&gt; - the challenge will contain a BASE-64 string depicting a photo (or any kind of QR-code-like data) that must be shown to the user.&lt;br/&gt;&amp;bull; &lt;code&gt;FLICKER_CODE&lt;/code&gt; - the challenge will contain a BASE-64 string depicting a flicker code animation that must be shown to the user.&lt;br/&gt;&lt;br/&gt;Note that this challenge type information does not originate from the bank, but is determined by finAPI internally. There is no guarantee that the determined challenge type is correct. Note also that this field may not be set, meaning that finAPI could not determine the challenge type of the procedure. | 
**ImplicitExecute** | **bool** | If &#39;true&#39;, then there is no need for your client to pass back anything to finAPI to continue the authorization when using this procedure. The authorization will be dealt with directly between the user, finAPI, and the bank. | 

## Methods

### NewTwoStepProcedure

`func NewTwoStepProcedure(procedureId string, procedureName string, procedureChallengeType NullableString, implicitExecute bool, ) *TwoStepProcedure`

NewTwoStepProcedure instantiates a new TwoStepProcedure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwoStepProcedureWithDefaults

`func NewTwoStepProcedureWithDefaults() *TwoStepProcedure`

NewTwoStepProcedureWithDefaults instantiates a new TwoStepProcedure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcedureId

`func (o *TwoStepProcedure) GetProcedureId() string`

GetProcedureId returns the ProcedureId field if non-nil, zero value otherwise.

### GetProcedureIdOk

`func (o *TwoStepProcedure) GetProcedureIdOk() (*string, bool)`

GetProcedureIdOk returns a tuple with the ProcedureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedureId

`func (o *TwoStepProcedure) SetProcedureId(v string)`

SetProcedureId sets ProcedureId field to given value.


### GetProcedureName

`func (o *TwoStepProcedure) GetProcedureName() string`

GetProcedureName returns the ProcedureName field if non-nil, zero value otherwise.

### GetProcedureNameOk

`func (o *TwoStepProcedure) GetProcedureNameOk() (*string, bool)`

GetProcedureNameOk returns a tuple with the ProcedureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedureName

`func (o *TwoStepProcedure) SetProcedureName(v string)`

SetProcedureName sets ProcedureName field to given value.


### GetProcedureChallengeType

`func (o *TwoStepProcedure) GetProcedureChallengeType() string`

GetProcedureChallengeType returns the ProcedureChallengeType field if non-nil, zero value otherwise.

### GetProcedureChallengeTypeOk

`func (o *TwoStepProcedure) GetProcedureChallengeTypeOk() (*string, bool)`

GetProcedureChallengeTypeOk returns a tuple with the ProcedureChallengeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcedureChallengeType

`func (o *TwoStepProcedure) SetProcedureChallengeType(v string)`

SetProcedureChallengeType sets ProcedureChallengeType field to given value.


### SetProcedureChallengeTypeNil

`func (o *TwoStepProcedure) SetProcedureChallengeTypeNil(b bool)`

 SetProcedureChallengeTypeNil sets the value for ProcedureChallengeType to be an explicit nil

### UnsetProcedureChallengeType
`func (o *TwoStepProcedure) UnsetProcedureChallengeType()`

UnsetProcedureChallengeType ensures that no value is present for ProcedureChallengeType, not even an explicit nil
### GetImplicitExecute

`func (o *TwoStepProcedure) GetImplicitExecute() bool`

GetImplicitExecute returns the ImplicitExecute field if non-nil, zero value otherwise.

### GetImplicitExecuteOk

`func (o *TwoStepProcedure) GetImplicitExecuteOk() (*bool, bool)`

GetImplicitExecuteOk returns a tuple with the ImplicitExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitExecute

`func (o *TwoStepProcedure) SetImplicitExecute(v bool)`

SetImplicitExecute sets ImplicitExecute field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


