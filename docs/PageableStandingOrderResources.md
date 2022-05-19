# PageableStandingOrderResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandingOrders** | [**[]StandingOrder**](StandingOrder.md) | &lt;strong&gt;Type:&lt;/strong&gt; StandingOrder&lt;br/&gt; List of standing orders | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewPageableStandingOrderResources

`func NewPageableStandingOrderResources(standingOrders []StandingOrder, paging Paging, ) *PageableStandingOrderResources`

NewPageableStandingOrderResources instantiates a new PageableStandingOrderResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableStandingOrderResourcesWithDefaults

`func NewPageableStandingOrderResourcesWithDefaults() *PageableStandingOrderResources`

NewPageableStandingOrderResourcesWithDefaults instantiates a new PageableStandingOrderResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandingOrders

`func (o *PageableStandingOrderResources) GetStandingOrders() []StandingOrder`

GetStandingOrders returns the StandingOrders field if non-nil, zero value otherwise.

### GetStandingOrdersOk

`func (o *PageableStandingOrderResources) GetStandingOrdersOk() (*[]StandingOrder, bool)`

GetStandingOrdersOk returns a tuple with the StandingOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandingOrders

`func (o *PageableStandingOrderResources) SetStandingOrders(v []StandingOrder)`

SetStandingOrders sets StandingOrders field to given value.


### GetPaging

`func (o *PageableStandingOrderResources) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableStandingOrderResources) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableStandingOrderResources) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


