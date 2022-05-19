# DeleteConsent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Local** | [**DeleteConsentResult**](DeleteConsentResult.md) | &lt;strong&gt;Type:&lt;/strong&gt; DeleteConsentResult&lt;br/&gt; Result of consent deletion in the finAPI database (local):&lt;br/&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;DELETED&lt;/code&gt; - if there was a stored consent and it was deleted.&lt;br/&gt;&amp;bull; &lt;code&gt;NOT_EXIST&lt;/code&gt; - if there was no stored consent.&lt;br/&gt; | 
**Remote** | [**DeleteConsentResult**](DeleteConsentResult.md) | &lt;strong&gt;Type:&lt;/strong&gt; DeleteConsentResult&lt;br/&gt; Result of consent deletion on the bank&#39;s side (remote):&lt;br/&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;DELETED&lt;/code&gt; - if the consent was successfully deleted on the bank side.&lt;br/&gt;&amp;bull; &lt;code&gt;NOT_SUPPORTED&lt;/code&gt; - if the bank doesn&#39;t support consent deletion.&lt;br/&gt;&amp;bull; &lt;code&gt;NOT_EXIST&lt;/code&gt; - if either there was no remote consent, or there was no local consent (making it impossible to identify any remote consent).&lt;br/&gt;&amp;bull; &lt;code&gt;NOT_DELETED&lt;/code&gt; - if the consent couldn&#39;t get deleted on the bank side.&lt;br/&gt; | 

## Methods

### NewDeleteConsent

`func NewDeleteConsent(local DeleteConsentResult, remote DeleteConsentResult, ) *DeleteConsent`

NewDeleteConsent instantiates a new DeleteConsent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteConsentWithDefaults

`func NewDeleteConsentWithDefaults() *DeleteConsent`

NewDeleteConsentWithDefaults instantiates a new DeleteConsent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocal

`func (o *DeleteConsent) GetLocal() DeleteConsentResult`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *DeleteConsent) GetLocalOk() (*DeleteConsentResult, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *DeleteConsent) SetLocal(v DeleteConsentResult)`

SetLocal sets Local field to given value.


### GetRemote

`func (o *DeleteConsent) GetRemote() DeleteConsentResult`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *DeleteConsent) GetRemoteOk() (*DeleteConsentResult, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *DeleteConsent) SetRemote(v DeleteConsentResult)`

SetRemote sets Remote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


