# SecurityList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Securities** | [**[]Security**](Security.md) | &lt;strong&gt;Type:&lt;/strong&gt; Security&lt;br/&gt; List of securities | 

## Methods

### NewSecurityList

`func NewSecurityList(securities []Security, ) *SecurityList`

NewSecurityList instantiates a new SecurityList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityListWithDefaults

`func NewSecurityListWithDefaults() *SecurityList`

NewSecurityListWithDefaults instantiates a new SecurityList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurities

`func (o *SecurityList) GetSecurities() []Security`

GetSecurities returns the Securities field if non-nil, zero value otherwise.

### GetSecuritiesOk

`func (o *SecurityList) GetSecuritiesOk() (*[]Security, bool)`

GetSecuritiesOk returns a tuple with the Securities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurities

`func (o *SecurityList) SetSecurities(v []Security)`

SetSecurities sets Securities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


