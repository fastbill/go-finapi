# PageableTppCredentialResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 
**TppCredentials** | [**[]TppCredentials**](TppCredentials.md) | &lt;strong&gt;Type:&lt;/strong&gt; TppCredentials&lt;br/&gt; List of TPP client credentials | 

## Methods

### NewPageableTppCredentialResources

`func NewPageableTppCredentialResources(paging Paging, tppCredentials []TppCredentials, ) *PageableTppCredentialResources`

NewPageableTppCredentialResources instantiates a new PageableTppCredentialResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableTppCredentialResourcesWithDefaults

`func NewPageableTppCredentialResourcesWithDefaults() *PageableTppCredentialResources`

NewPageableTppCredentialResourcesWithDefaults instantiates a new PageableTppCredentialResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *PageableTppCredentialResources) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableTppCredentialResources) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableTppCredentialResources) SetPaging(v Paging)`

SetPaging sets Paging field to given value.


### GetTppCredentials

`func (o *PageableTppCredentialResources) GetTppCredentials() []TppCredentials`

GetTppCredentials returns the TppCredentials field if non-nil, zero value otherwise.

### GetTppCredentialsOk

`func (o *PageableTppCredentialResources) GetTppCredentialsOk() (*[]TppCredentials, bool)`

GetTppCredentialsOk returns a tuple with the TppCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppCredentials

`func (o *PageableTppCredentialResources) SetTppCredentials(v []TppCredentials)`

SetTppCredentials sets TppCredentials field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


