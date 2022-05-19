# PageableTppAuthenticationGroupResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TppAuthenticationGroups** | [**[]TppAuthenticationGroup**](TppAuthenticationGroup.md) | &lt;strong&gt;Type:&lt;/strong&gt; TppAuthenticationGroup&lt;br/&gt; List of received TPP authentication groups | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewPageableTppAuthenticationGroupResources

`func NewPageableTppAuthenticationGroupResources(tppAuthenticationGroups []TppAuthenticationGroup, paging Paging, ) *PageableTppAuthenticationGroupResources`

NewPageableTppAuthenticationGroupResources instantiates a new PageableTppAuthenticationGroupResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableTppAuthenticationGroupResourcesWithDefaults

`func NewPageableTppAuthenticationGroupResourcesWithDefaults() *PageableTppAuthenticationGroupResources`

NewPageableTppAuthenticationGroupResourcesWithDefaults instantiates a new PageableTppAuthenticationGroupResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTppAuthenticationGroups

`func (o *PageableTppAuthenticationGroupResources) GetTppAuthenticationGroups() []TppAuthenticationGroup`

GetTppAuthenticationGroups returns the TppAuthenticationGroups field if non-nil, zero value otherwise.

### GetTppAuthenticationGroupsOk

`func (o *PageableTppAuthenticationGroupResources) GetTppAuthenticationGroupsOk() (*[]TppAuthenticationGroup, bool)`

GetTppAuthenticationGroupsOk returns a tuple with the TppAuthenticationGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppAuthenticationGroups

`func (o *PageableTppAuthenticationGroupResources) SetTppAuthenticationGroups(v []TppAuthenticationGroup)`

SetTppAuthenticationGroups sets TppAuthenticationGroups field to given value.


### GetPaging

`func (o *PageableTppAuthenticationGroupResources) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableTppAuthenticationGroupResources) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableTppAuthenticationGroupResources) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


