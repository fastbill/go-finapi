# TppCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | A certificate identifier. | 
**CertificateType** | [**TppCertificateType**](TppCertificateType.md) | &lt;strong&gt;Type:&lt;/strong&gt; TppCertificateType&lt;br/&gt; Type of certificate. | 
**Label** | **string** | Label of certificate. | 
**ValidFrom** | **string** | Valid from date in the format &#39;YYYY-MM-DD&#39;. | 
**ValidUntil** | **string** | Valid until date in the format &#39;YYYY-MM-DD&#39;. | 

## Methods

### NewTppCertificate

`func NewTppCertificate(id int64, certificateType TppCertificateType, label string, validFrom string, validUntil string, ) *TppCertificate`

NewTppCertificate instantiates a new TppCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTppCertificateWithDefaults

`func NewTppCertificateWithDefaults() *TppCertificate`

NewTppCertificateWithDefaults instantiates a new TppCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TppCertificate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TppCertificate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TppCertificate) SetId(v int64)`

SetId sets Id field to given value.


### GetCertificateType

`func (o *TppCertificate) GetCertificateType() TppCertificateType`

GetCertificateType returns the CertificateType field if non-nil, zero value otherwise.

### GetCertificateTypeOk

`func (o *TppCertificate) GetCertificateTypeOk() (*TppCertificateType, bool)`

GetCertificateTypeOk returns a tuple with the CertificateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateType

`func (o *TppCertificate) SetCertificateType(v TppCertificateType)`

SetCertificateType sets CertificateType field to given value.


### GetLabel

`func (o *TppCertificate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TppCertificate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TppCertificate) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValidFrom

`func (o *TppCertificate) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *TppCertificate) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *TppCertificate) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.


### GetValidUntil

`func (o *TppCertificate) GetValidUntil() string`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *TppCertificate) GetValidUntilOk() (*string, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *TppCertificate) SetValidUntil(v string)`

SetValidUntil sets ValidUntil field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


