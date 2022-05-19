# TppCertificateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**TppCertificateType**](TppCertificateType.md) | &lt;strong&gt;Type:&lt;/strong&gt; TppCertificateType&lt;br/&gt; The type of the submitted certificate | 
**PublicKey** | **string** | A certificate (public key) | 
**PrivateKey** | **string** | A private key in PKCS #8 or PKCS #1 format. PKCS #1/#8 private keys are typically exchanged in the PEM base64-encoded format (https://support.quovadisglobal.com/kb/a37/what-is-pem-format.aspx)&lt;br/&gt;&lt;br/&gt;NOTE: The certificate should have one of the following headers:&lt;br/&gt;- &#39;-----BEGIN RSA PRIVATE KEY-----&#39;&lt;br/&gt;- &#39;-----BEGIN PRIVATE KEY-----&#39;&lt;br/&gt;- &#39;-----BEGIN ENCRYPTED PRIVATE KEY-----&#39;&lt;br/&gt;Any other header denotes that the private key is neither in PKCS #8 nor in PKCS #1 formats!&lt;br/&gt;&lt;br/&gt;Also, bear in mind that if the private key is in PKCS #1 encrypted format, the encryption information must be provided with explicitly separated lines (the JSON must contain \&quot;\\n\&quot; at the end of each line), such as in the example below:&lt;br/&gt;-----BEGIN RSA PRIVATE KEY-----&lt;br/&gt;Proc-Type: 4,ENCRYPTED&lt;br/&gt;DEK-Info: AES-256-CBC,BFA11F426E7D634BC621C77A72B804DB&lt;br/&gt;...&lt;br/&gt;-----END RSA PRIVATE KEY----- | 
**Passphrase** | Pointer to **string** | Optional passphrase for the private key | [optional] 
**CaPublicKey** | Pointer to **string** | A certificate (public key) of the certificate authority (CA) that signed the certificate. Required in certain cases to build the PKI path between Access and the bank&#39;s API when banks do not possess intermediate TLS certificates while placing the trust chain. | [optional] 
**Label** | **string** | A label for the certificate | 
**ValidFromDate** | Pointer to **string** | Start day of the certificate&#39;s validity, in the format &#39;YYYY-MM-DD&#39;. Default is the passed certificate validFrom date | [optional] 
**ValidUntilDate** | Pointer to **string** | Expiration day of the certificate&#39;s validity, in the format &#39;YYYY-MM-DD&#39;. Default is the passed certificate validUntil date | [optional] 

## Methods

### NewTppCertificateParams

`func NewTppCertificateParams(type_ TppCertificateType, publicKey string, privateKey string, label string, ) *TppCertificateParams`

NewTppCertificateParams instantiates a new TppCertificateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTppCertificateParamsWithDefaults

`func NewTppCertificateParamsWithDefaults() *TppCertificateParams`

NewTppCertificateParamsWithDefaults instantiates a new TppCertificateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TppCertificateParams) GetType() TppCertificateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TppCertificateParams) GetTypeOk() (*TppCertificateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TppCertificateParams) SetType(v TppCertificateType)`

SetType sets Type field to given value.


### GetPublicKey

`func (o *TppCertificateParams) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *TppCertificateParams) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *TppCertificateParams) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetPrivateKey

`func (o *TppCertificateParams) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *TppCertificateParams) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *TppCertificateParams) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetPassphrase

`func (o *TppCertificateParams) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *TppCertificateParams) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *TppCertificateParams) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *TppCertificateParams) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.

### GetCaPublicKey

`func (o *TppCertificateParams) GetCaPublicKey() string`

GetCaPublicKey returns the CaPublicKey field if non-nil, zero value otherwise.

### GetCaPublicKeyOk

`func (o *TppCertificateParams) GetCaPublicKeyOk() (*string, bool)`

GetCaPublicKeyOk returns a tuple with the CaPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaPublicKey

`func (o *TppCertificateParams) SetCaPublicKey(v string)`

SetCaPublicKey sets CaPublicKey field to given value.

### HasCaPublicKey

`func (o *TppCertificateParams) HasCaPublicKey() bool`

HasCaPublicKey returns a boolean if a field has been set.

### GetLabel

`func (o *TppCertificateParams) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *TppCertificateParams) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *TppCertificateParams) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetValidFromDate

`func (o *TppCertificateParams) GetValidFromDate() string`

GetValidFromDate returns the ValidFromDate field if non-nil, zero value otherwise.

### GetValidFromDateOk

`func (o *TppCertificateParams) GetValidFromDateOk() (*string, bool)`

GetValidFromDateOk returns a tuple with the ValidFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFromDate

`func (o *TppCertificateParams) SetValidFromDate(v string)`

SetValidFromDate sets ValidFromDate field to given value.

### HasValidFromDate

`func (o *TppCertificateParams) HasValidFromDate() bool`

HasValidFromDate returns a boolean if a field has been set.

### GetValidUntilDate

`func (o *TppCertificateParams) GetValidUntilDate() string`

GetValidUntilDate returns the ValidUntilDate field if non-nil, zero value otherwise.

### GetValidUntilDateOk

`func (o *TppCertificateParams) GetValidUntilDateOk() (*string, bool)`

GetValidUntilDateOk returns a tuple with the ValidUntilDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntilDate

`func (o *TppCertificateParams) SetValidUntilDate(v string)`

SetValidUntilDate sets ValidUntilDate field to given value.

### HasValidUntilDate

`func (o *TppCertificateParams) HasValidUntilDate() bool`

HasValidUntilDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


