# ChangeClientCredentialsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | client_id of the client that you want to change the secret for | 
**OldClientSecret** | **string** | Old (&#x3D;current) client_secret | 
**NewClientSecret** | **string** | New client_secret. Required length is 36. Allowed symbols: Digits (0 through 9), lower-case and upper-case letters (A through Z), and the dash symbol (\&quot;-\&quot;). | 

## Methods

### NewChangeClientCredentialsParams

`func NewChangeClientCredentialsParams(clientId string, oldClientSecret string, newClientSecret string, ) *ChangeClientCredentialsParams`

NewChangeClientCredentialsParams instantiates a new ChangeClientCredentialsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeClientCredentialsParamsWithDefaults

`func NewChangeClientCredentialsParamsWithDefaults() *ChangeClientCredentialsParams`

NewChangeClientCredentialsParamsWithDefaults instantiates a new ChangeClientCredentialsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ChangeClientCredentialsParams) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ChangeClientCredentialsParams) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ChangeClientCredentialsParams) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetOldClientSecret

`func (o *ChangeClientCredentialsParams) GetOldClientSecret() string`

GetOldClientSecret returns the OldClientSecret field if non-nil, zero value otherwise.

### GetOldClientSecretOk

`func (o *ChangeClientCredentialsParams) GetOldClientSecretOk() (*string, bool)`

GetOldClientSecretOk returns a tuple with the OldClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldClientSecret

`func (o *ChangeClientCredentialsParams) SetOldClientSecret(v string)`

SetOldClientSecret sets OldClientSecret field to given value.


### GetNewClientSecret

`func (o *ChangeClientCredentialsParams) GetNewClientSecret() string`

GetNewClientSecret returns the NewClientSecret field if non-nil, zero value otherwise.

### GetNewClientSecretOk

`func (o *ChangeClientCredentialsParams) GetNewClientSecretOk() (*string, bool)`

GetNewClientSecretOk returns a tuple with the NewClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewClientSecret

`func (o *ChangeClientCredentialsParams) SetNewClientSecret(v string)`

SetNewClientSecret sets NewClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


