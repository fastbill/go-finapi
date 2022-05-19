# BankInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; Bank interface. Possible values:&lt;br&gt;&lt;br&gt;&amp;bull; &lt;code&gt;WEB_SCRAPER&lt;/code&gt; - means that finAPI will parse data from the bank&#39;s online banking website.&lt;br&gt;&amp;bull; &lt;code&gt;FINTS_SERVER&lt;/code&gt; - means that finAPI will download data via the bank&#39;s FinTS server.&lt;br&gt;&amp;bull; &lt;code&gt;XS2A&lt;/code&gt; - means that finAPI will download data via the bank&#39;s XS2A interface.&lt;br&gt; | 
**TppAuthenticationGroup** | [**NullableTppAuthenticationGroup**](TppAuthenticationGroup.md) | &lt;strong&gt;Type:&lt;/strong&gt; TppAuthenticationGroup&lt;br/&gt; TPP Authentication Group which the bank interface is connected to | 
**LoginCredentials** | [**[]BankInterfaceLoginField**](BankInterfaceLoginField.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankInterfaceLoginField&lt;br/&gt; Login fields for this interface (in the order that we suggest to show them to the user) | 
**Properties** | [**[]BankInterfaceProperty**](BankInterfaceProperty.md) |  | 
**LoginHint** | **NullableString** | Login hint. Contains a German message for the user that explains what kind of credentials are expected.&lt;br/&gt;&lt;br/&gt;Please note that it is essential to always show the login hint to the user if there is one, as the credentials that finAPI requires for the bank might be different to the credentials that the user knows from his online banking.&lt;br/&gt;&lt;br/&gt;Also note that the contents of this field should always be interpreted as HTML, as the text might contain HTML tags for highlighted words, paragraphs, etc. | 
**Health** | **int32** | The health status of this interface. This is a value between 0 and 100, depicting the percentage of successful communication attempts with the bank via this interface during the latest couple of bank connection imports or updates (across the entire finAPI system). Note that &#39;successful&#39; means that there was no technical error trying to establish a communication with the bank. Non-technical errors (like incorrect credentials) are regarded successful communication attempts. | 
**LastCommunicationAttempt** | **NullableString** | Time of the last communication attempt with this interface during an import, update or connect interface (across the entire finAPI system). The value is returned in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**LastSuccessfulCommunication** | **NullableString** | Time of the last successful communication with this interface during an import, update or connect interface (across the entire finAPI system). The value is returned in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**IsMoneyTransferSupported** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;paymentCapabilities&#39; instead.&lt;br/&gt;&lt;br/&gt;Whether this interface has the general capability to do money transfers.&lt;br/&gt;&lt;br/&gt;Note: It still depends on the specifics of an account whether the account supports money transfer. In general, you should prefer the field AccountInterface.capabilities to determine what is possible for a specific account. This field here is meant to be used mainly for when you want to do standalone money transfers (finAPI Payment product), i.e. when you do not plan to import an account and thus will not have the information about the account&#39;s capabilities. | 
**IsAisSupported** | **bool** | Whether this interface has the general capability to perform Account Information Services (AIS), i.e. if this interface can be used to download accounts, balances and transactions.  | 
**PaymentCapabilities** | [**BankInterfacePaymentCapabilities**](BankInterfacePaymentCapabilities.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankInterfacePaymentCapabilities&lt;br/&gt; The general payment capabilities of this interface. If a capability is &#39;true&#39;, it means that the option is supported, as long as the involved account also supports it (see AccountInterface.capabilities and AccountInterface.paymentCapabilities). If a capability is &#39;false&#39;, the option is not supported for any account. | 
**AisAccountTypes** | [**[]AccountType**](AccountType.md) |  | 

## Methods

### NewBankInterface

`func NewBankInterface(interface_ BankingInterface, tppAuthenticationGroup NullableTppAuthenticationGroup, loginCredentials []BankInterfaceLoginField, properties []BankInterfaceProperty, loginHint NullableString, health int32, lastCommunicationAttempt NullableString, lastSuccessfulCommunication NullableString, isMoneyTransferSupported bool, isAisSupported bool, paymentCapabilities BankInterfacePaymentCapabilities, aisAccountTypes []AccountType, ) *BankInterface`

NewBankInterface instantiates a new BankInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankInterfaceWithDefaults

`func NewBankInterfaceWithDefaults() *BankInterface`

NewBankInterfaceWithDefaults instantiates a new BankInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *BankInterface) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *BankInterface) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *BankInterface) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.


### GetTppAuthenticationGroup

`func (o *BankInterface) GetTppAuthenticationGroup() TppAuthenticationGroup`

GetTppAuthenticationGroup returns the TppAuthenticationGroup field if non-nil, zero value otherwise.

### GetTppAuthenticationGroupOk

`func (o *BankInterface) GetTppAuthenticationGroupOk() (*TppAuthenticationGroup, bool)`

GetTppAuthenticationGroupOk returns a tuple with the TppAuthenticationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTppAuthenticationGroup

`func (o *BankInterface) SetTppAuthenticationGroup(v TppAuthenticationGroup)`

SetTppAuthenticationGroup sets TppAuthenticationGroup field to given value.


### SetTppAuthenticationGroupNil

`func (o *BankInterface) SetTppAuthenticationGroupNil(b bool)`

 SetTppAuthenticationGroupNil sets the value for TppAuthenticationGroup to be an explicit nil

### UnsetTppAuthenticationGroup
`func (o *BankInterface) UnsetTppAuthenticationGroup()`

UnsetTppAuthenticationGroup ensures that no value is present for TppAuthenticationGroup, not even an explicit nil
### GetLoginCredentials

`func (o *BankInterface) GetLoginCredentials() []BankInterfaceLoginField`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *BankInterface) GetLoginCredentialsOk() (*[]BankInterfaceLoginField, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *BankInterface) SetLoginCredentials(v []BankInterfaceLoginField)`

SetLoginCredentials sets LoginCredentials field to given value.


### GetProperties

`func (o *BankInterface) GetProperties() []BankInterfaceProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *BankInterface) GetPropertiesOk() (*[]BankInterfaceProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *BankInterface) SetProperties(v []BankInterfaceProperty)`

SetProperties sets Properties field to given value.


### GetLoginHint

`func (o *BankInterface) GetLoginHint() string`

GetLoginHint returns the LoginHint field if non-nil, zero value otherwise.

### GetLoginHintOk

`func (o *BankInterface) GetLoginHintOk() (*string, bool)`

GetLoginHintOk returns a tuple with the LoginHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHint

`func (o *BankInterface) SetLoginHint(v string)`

SetLoginHint sets LoginHint field to given value.


### SetLoginHintNil

`func (o *BankInterface) SetLoginHintNil(b bool)`

 SetLoginHintNil sets the value for LoginHint to be an explicit nil

### UnsetLoginHint
`func (o *BankInterface) UnsetLoginHint()`

UnsetLoginHint ensures that no value is present for LoginHint, not even an explicit nil
### GetHealth

`func (o *BankInterface) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *BankInterface) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *BankInterface) SetHealth(v int32)`

SetHealth sets Health field to given value.


### GetLastCommunicationAttempt

`func (o *BankInterface) GetLastCommunicationAttempt() string`

GetLastCommunicationAttempt returns the LastCommunicationAttempt field if non-nil, zero value otherwise.

### GetLastCommunicationAttemptOk

`func (o *BankInterface) GetLastCommunicationAttemptOk() (*string, bool)`

GetLastCommunicationAttemptOk returns a tuple with the LastCommunicationAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicationAttempt

`func (o *BankInterface) SetLastCommunicationAttempt(v string)`

SetLastCommunicationAttempt sets LastCommunicationAttempt field to given value.


### SetLastCommunicationAttemptNil

`func (o *BankInterface) SetLastCommunicationAttemptNil(b bool)`

 SetLastCommunicationAttemptNil sets the value for LastCommunicationAttempt to be an explicit nil

### UnsetLastCommunicationAttempt
`func (o *BankInterface) UnsetLastCommunicationAttempt()`

UnsetLastCommunicationAttempt ensures that no value is present for LastCommunicationAttempt, not even an explicit nil
### GetLastSuccessfulCommunication

`func (o *BankInterface) GetLastSuccessfulCommunication() string`

GetLastSuccessfulCommunication returns the LastSuccessfulCommunication field if non-nil, zero value otherwise.

### GetLastSuccessfulCommunicationOk

`func (o *BankInterface) GetLastSuccessfulCommunicationOk() (*string, bool)`

GetLastSuccessfulCommunicationOk returns a tuple with the LastSuccessfulCommunication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulCommunication

`func (o *BankInterface) SetLastSuccessfulCommunication(v string)`

SetLastSuccessfulCommunication sets LastSuccessfulCommunication field to given value.


### SetLastSuccessfulCommunicationNil

`func (o *BankInterface) SetLastSuccessfulCommunicationNil(b bool)`

 SetLastSuccessfulCommunicationNil sets the value for LastSuccessfulCommunication to be an explicit nil

### UnsetLastSuccessfulCommunication
`func (o *BankInterface) UnsetLastSuccessfulCommunication()`

UnsetLastSuccessfulCommunication ensures that no value is present for LastSuccessfulCommunication, not even an explicit nil
### GetIsMoneyTransferSupported

`func (o *BankInterface) GetIsMoneyTransferSupported() bool`

GetIsMoneyTransferSupported returns the IsMoneyTransferSupported field if non-nil, zero value otherwise.

### GetIsMoneyTransferSupportedOk

`func (o *BankInterface) GetIsMoneyTransferSupportedOk() (*bool, bool)`

GetIsMoneyTransferSupportedOk returns a tuple with the IsMoneyTransferSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMoneyTransferSupported

`func (o *BankInterface) SetIsMoneyTransferSupported(v bool)`

SetIsMoneyTransferSupported sets IsMoneyTransferSupported field to given value.


### GetIsAisSupported

`func (o *BankInterface) GetIsAisSupported() bool`

GetIsAisSupported returns the IsAisSupported field if non-nil, zero value otherwise.

### GetIsAisSupportedOk

`func (o *BankInterface) GetIsAisSupportedOk() (*bool, bool)`

GetIsAisSupportedOk returns a tuple with the IsAisSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAisSupported

`func (o *BankInterface) SetIsAisSupported(v bool)`

SetIsAisSupported sets IsAisSupported field to given value.


### GetPaymentCapabilities

`func (o *BankInterface) GetPaymentCapabilities() BankInterfacePaymentCapabilities`

GetPaymentCapabilities returns the PaymentCapabilities field if non-nil, zero value otherwise.

### GetPaymentCapabilitiesOk

`func (o *BankInterface) GetPaymentCapabilitiesOk() (*BankInterfacePaymentCapabilities, bool)`

GetPaymentCapabilitiesOk returns a tuple with the PaymentCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCapabilities

`func (o *BankInterface) SetPaymentCapabilities(v BankInterfacePaymentCapabilities)`

SetPaymentCapabilities sets PaymentCapabilities field to given value.


### GetAisAccountTypes

`func (o *BankInterface) GetAisAccountTypes() []AccountType`

GetAisAccountTypes returns the AisAccountTypes field if non-nil, zero value otherwise.

### GetAisAccountTypesOk

`func (o *BankInterface) GetAisAccountTypesOk() (*[]AccountType, bool)`

GetAisAccountTypesOk returns a tuple with the AisAccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAisAccountTypes

`func (o *BankInterface) SetAisAccountTypes(v []AccountType)`

SetAisAccountTypes sets AisAccountTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


