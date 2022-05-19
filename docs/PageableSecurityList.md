# PageableSecurityList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Securities** | [**[]Security**](Security.md) | &lt;strong&gt;Type:&lt;/strong&gt; Security&lt;br/&gt; List of securities | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewPageableSecurityList

`func NewPageableSecurityList(securities []Security, paging Paging, ) *PageableSecurityList`

NewPageableSecurityList instantiates a new PageableSecurityList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableSecurityListWithDefaults

`func NewPageableSecurityListWithDefaults() *PageableSecurityList`

NewPageableSecurityListWithDefaults instantiates a new PageableSecurityList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurities

`func (o *PageableSecurityList) GetSecurities() []Security`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *PageableSecurityList) GetSecuritiesOk() (*[]Security, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *PageableSecurityList) SetSecurities(v []Security)`

SetSecurities sets Securities field to given value.


### GetPaging

`func (o *PageableSecurityList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableSecurityList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableSecurityList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


