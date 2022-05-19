# AccountReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **string** | The account&#39;s IBAN | 

## Methods

### NewAccountReference

`func NewAccountReference(iban string, ) *AccountReference`

NewAccountReference instantiates a new AccountReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountReferenceWithDefaults

`func NewAccountReferenceWithDefaults() *AccountReference`

NewAccountReferenceWithDefaults instantiates a new AccountReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *AccountReference) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *AccountReference) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *AccountReference) SetIban(v string)`

SetIban sets Iban field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


