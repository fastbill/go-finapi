# TppCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The credential identifier. | 
**Label** | **string** | Label of tpp client credentials set. | 
**TppAuthenticationGroupId** | **int64** | TPP Authentication Group ID. | 
**ValidFrom** | **string** | Valid from date in the format &#39;YYYY-MM-DD&#39;. | 
**ValidUntil** | **NullableString** | Valid until date in the format &#39;YYYY-MM-DD&#39;. | 

## Methods

### NewTppCredentials

`func NewTppCredentials(id int64, label string, tppAuthenticationGroupId int64, validFrom string, validUntil NullableString, ) *TppCredentials`

NewTppCredentials instantiates a new TppCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTppCredentialsWithDefaults

`func NewTppCredentialsWithDefaults() *TppCredentials`

NewTppCredentialsWithDefaults instantiates a new TppCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TppCredentials) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TppCredentials) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TppCredentials) SetId(v int64)`

SetId sets Id field to given value.


### GetLabel

`func (o *TppCredentials) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TppCredentials) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TppCredentials) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTppAuthenticationGroupId

`func (o *TppCredentials) GetTppAuthenticationGroupId() int64`

GetTppAuthenticationGroupId returns the TppAuthenticationGroupId field if non-nil, zero value otherwise.

### GetTppAuthenticationGroupIdOk

`func (o *TppCredentials) GetTppAuthenticationGroupIdOk() (*int64, bool)`

GetTppAuthenticationGroupIdOk returns a tuple with the TppAuthenticationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppAuthenticationGroupId

`func (o *TppCredentials) SetTppAuthenticationGroupId(v int64)`

SetTppAuthenticationGroupId sets TppAuthenticationGroupId field to given value.


### GetValidFrom

`func (o *TppCredentials) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *TppCredentials) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *TppCredentials) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.


### GetValidUntil

`func (o *TppCredentials) GetValidUntil() string`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *TppCredentials) GetValidUntilOk() (*string, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *TppCredentials) SetValidUntil(v string)`

SetValidUntil sets ValidUntil field to given value.


### SetValidUntilNil

`func (o *TppCredentials) SetValidUntilNil(b bool)`

 SetValidUntilNil sets the value for ValidUntil to be an explicit nil

### UnsetValidUntil
`func (o *TppCredentials) UnsetValidUntil()`

UnsetValidUntil ensures that no value is present for ValidUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


