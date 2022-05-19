# PageableBankList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Banks** | [**[]Bank**](Bank.md) | &lt;strong&gt;Type:&lt;/strong&gt; Bank&lt;br/&gt; Banks | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewPageableBankList

`func NewPageableBankList(banks []Bank, paging Paging, ) *PageableBankList`

NewPageableBankList instantiates a new PageableBankList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableBankListWithDefaults

`func NewPageableBankListWithDefaults() *PageableBankList`

NewPageableBankListWithDefaults instantiates a new PageableBankList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBanks

`func (o *PageableBankList) GetBanks() []Bank`

GetBanks returns the Banks field if non-nil, zero value otherwise.

### GetBanksOk

`func (o *PageableBankList) GetBanksOk() (*[]Bank, bool)`

GetBanksOk returns a tuple with the Banks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanks

`func (o *PageableBankList) SetBanks(v []Bank)`

SetBanks sets Banks field to given value.


### GetPaging

`func (o *PageableBankList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableBankList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableBankList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


