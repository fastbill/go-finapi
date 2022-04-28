# BankConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**BankConsentStatus**](BankConsentStatus.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankConsentStatus&lt;br/&gt; Status of the consent. If &lt;code&gt;PRESENT&lt;/code&gt;, it means that finAPI has a consent stored and can use it to connect to the bank. If &lt;code&gt;NOT_PRESENT&lt;/code&gt;, finAPI will need to acquire a consent when connecting to the bank, which may require login credentials (either passed by the client, or stored in finAPI), and/or user involvement. Note that even when a consent is &lt;code&gt;PRESENT&lt;/code&gt;, it may no longer be valid and finAPI will still have to acquire a new consent. | 
**ExpiresAt** | **NullableString** | Expiration time of the consent in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 

## Methods

### NewBankConsent

`func NewBankConsent(status BankConsentStatus, expiresAt NullableString, ) *BankConsent`

NewBankConsent instantiates a new BankConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankConsentWithDefaults

`func NewBankConsentWithDefaults() *BankConsent`

NewBankConsentWithDefaults instantiates a new BankConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BankConsent) GetStatus() BankConsentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BankConsent) GetStatusOk() (*BankConsentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BankConsent) SetStatus(v BankConsentStatus)`

SetStatus sets Status field to given value.


### GetExpiresAt

`func (o *BankConsent) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *BankConsent) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *BankConsent) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### SetExpiresAtNil

`func (o *BankConsent) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *BankConsent) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


