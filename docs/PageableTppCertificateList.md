# PageableTppCertificateList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 
**TppCertificates** | [**[]TppCertificate**](TppCertificate.md) | &lt;strong&gt;Type:&lt;/strong&gt; TppCertificate&lt;br/&gt; List of certificates | 

## Methods

### NewPageableTppCertificateList

`func NewPageableTppCertificateList(paging Paging, tppCertificates []TppCertificate, ) *PageableTppCertificateList`

NewPageableTppCertificateList instantiates a new PageableTppCertificateList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableTppCertificateListWithDefaults

`func NewPageableTppCertificateListWithDefaults() *PageableTppCertificateList`

NewPageableTppCertificateListWithDefaults instantiates a new PageableTppCertificateList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaging

`func (o *PageableTppCertificateList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableTppCertificateList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableTppCertificateList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.


### GetTppCertificates

`func (o *PageableTppCertificateList) GetTppCertificates() []TppCertificate`

GetTppCertificates returns the TppCertificates field if non-nil, zero value otherwise.

### GetTppCertificatesOk

`func (o *PageableTppCertificateList) GetTppCertificatesOk() (*[]TppCertificate, bool)`

GetTppCertificatesOk returns a tuple with the TppCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppCertificates

`func (o *PageableTppCertificateList) SetTppCertificates(v []TppCertificate)`

SetTppCertificates sets TppCertificates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


