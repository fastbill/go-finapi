# EditTppCredentialParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TppAuthenticationGroupId** | Pointer to **int64** | The TPP Authentication Group Id for which the credentials can be used | [optional] 
**Label** | Pointer to **string** | Label for credentials | [optional] 
**TppClientId** | Pointer to **string** | ID of the TPP accessing the ASPSP API, as provided by the ASPSP as the result of registration | [optional] 
**TppClientSecret** | Pointer to **string** | Secret of the TPP accessing the ASPSP API, as provided by the ASPSP as the result of registration | [optional] 
**TppApiKey** | Pointer to **string** | API Key provided by ASPSP  as the result of registration | [optional] 
**TppName** | Pointer to **string** | TPP name | [optional] 
**ValidFromDate** | Pointer to **string** | Credentials \&quot;valid from\&quot; date in the format &#39;YYYY-MM-DD&#39;. Default is today&#39;s date | [optional] 
**ValidUntilDate** | Pointer to **string** | Credentials \&quot;valid until\&quot; date in the format &#39;YYYY-MM-DD&#39;. Default is null which means \&quot;indefinite\&quot; (no limit) | [optional] 

## Methods

### NewEditTppCredentialParams

`func NewEditTppCredentialParams() *EditTppCredentialParams`

NewEditTppCredentialParams instantiates a new EditTppCredentialParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditTppCredentialParamsWithDefaults

`func NewEditTppCredentialParamsWithDefaults() *EditTppCredentialParams`

NewEditTppCredentialParamsWithDefaults instantiates a new EditTppCredentialParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTppAuthenticationGroupId

`func (o *EditTppCredentialParams) GetTppAuthenticationGroupId() int64`

GetTppAuthenticationGroupId returns the TppAuthenticationGroupId field if non-nil, zero value otherwise.

### GetTppAuthenticationGroupIdOk

`func (o *EditTppCredentialParams) GetTppAuthenticationGroupIdOk() (*int64, bool)`

GetTppAuthenticationGroupIdOk returns a tuple with the TppAuthenticationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppAuthenticationGroupId

`func (o *EditTppCredentialParams) SetTppAuthenticationGroupId(v int64)`

SetTppAuthenticationGroupId sets TppAuthenticationGroupId field to given value.

### HasTppAuthenticationGroupId

`func (o *EditTppCredentialParams) HasTppAuthenticationGroupId() bool`

HasTppAuthenticationGroupId returns a boolean if a field has been set.

### GetLabel

`func (o *EditTppCredentialParams) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EditTppCredentialParams) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EditTppCredentialParams) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EditTppCredentialParams) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTppClientId

`func (o *EditTppCredentialParams) GetTppClientId() string`

GetTppClientId returns the TppClientId field if non-nil, zero value otherwise.

### GetTppClientIdOk

`func (o *EditTppCredentialParams) GetTppClientIdOk() (*string, bool)`

GetTppClientIdOk returns a tuple with the TppClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppClientId

`func (o *EditTppCredentialParams) SetTppClientId(v string)`

SetTppClientId sets TppClientId field to given value.

### HasTppClientId

`func (o *EditTppCredentialParams) HasTppClientId() bool`

HasTppClientId returns a boolean if a field has been set.

### GetTppClientSecret

`func (o *EditTppCredentialParams) GetTppClientSecret() string`

GetTppClientSecret returns the TppClientSecret field if non-nil, zero value otherwise.

### GetTppClientSecretOk

`func (o *EditTppCredentialParams) GetTppClientSecretOk() (*string, bool)`

GetTppClientSecretOk returns a tuple with the TppClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppClientSecret

`func (o *EditTppCredentialParams) SetTppClientSecret(v string)`

SetTppClientSecret sets TppClientSecret field to given value.

### HasTppClientSecret

`func (o *EditTppCredentialParams) HasTppClientSecret() bool`

HasTppClientSecret returns a boolean if a field has been set.

### GetTppApiKey

`func (o *EditTppCredentialParams) GetTppApiKey() string`

GetTppApiKey returns the TppApiKey field if non-nil, zero value otherwise.

### GetTppApiKeyOk

`func (o *EditTppCredentialParams) GetTppApiKeyOk() (*string, bool)`

GetTppApiKeyOk returns a tuple with the TppApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppApiKey

`func (o *EditTppCredentialParams) SetTppApiKey(v string)`

SetTppApiKey sets TppApiKey field to given value.

### HasTppApiKey

`func (o *EditTppCredentialParams) HasTppApiKey() bool`

HasTppApiKey returns a boolean if a field has been set.

### GetTppName

`func (o *EditTppCredentialParams) GetTppName() string`

GetTppName returns the TppName field if non-nil, zero value otherwise.

### GetTppNameOk

`func (o *EditTppCredentialParams) GetTppNameOk() (*string, bool)`

GetTppNameOk returns a tuple with the TppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppName

`func (o *EditTppCredentialParams) SetTppName(v string)`

SetTppName sets TppName field to given value.

### HasTppName

`func (o *EditTppCredentialParams) HasTppName() bool`

HasTppName returns a boolean if a field has been set.

### GetValidFromDate

`func (o *EditTppCredentialParams) GetValidFromDate() string`

GetValidFromDate returns the ValidFromDate field if non-nil, zero value otherwise.

### GetValidFromDateOk

`func (o *EditTppCredentialParams) GetValidFromDateOk() (*string, bool)`

GetValidFromDateOk returns a tuple with the ValidFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFromDate

`func (o *EditTppCredentialParams) SetValidFromDate(v string)`

SetValidFromDate sets ValidFromDate field to given value.

### HasValidFromDate

`func (o *EditTppCredentialParams) HasValidFromDate() bool`

HasValidFromDate returns a boolean if a field has been set.

### GetValidUntilDate

`func (o *EditTppCredentialParams) GetValidUntilDate() string`

GetValidUntilDate returns the ValidUntilDate field if non-nil, zero value otherwise.

### GetValidUntilDateOk

`func (o *EditTppCredentialParams) GetValidUntilDateOk() (*string, bool)`

GetValidUntilDateOk returns a tuple with the ValidUntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntilDate

`func (o *EditTppCredentialParams) SetValidUntilDate(v string)`

SetValidUntilDate sets ValidUntilDate field to given value.

### HasValidUntilDate

`func (o *EditTppCredentialParams) HasValidUntilDate() bool`

HasValidUntilDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


