# AccountInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interface** | [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; Bank interface. Possible values:&lt;br&gt;&lt;br&gt;&amp;bull; &lt;code&gt;WEB_SCRAPER&lt;/code&gt; - finAPI will parse account data from the bank&#39;s online banking website.&lt;br&gt;&amp;bull; &lt;code&gt;FINTS_SERVER&lt;/code&gt; - finAPI will download account data via the bank&#39;s FinTS interface.&lt;br&gt;&amp;bull; &lt;code&gt;XS2A&lt;/code&gt; - finAPI will download account data via the bank&#39;s XS2A interface.&lt;br&gt; | 
**Status** | [**AccountStatus**](AccountStatus.md) | &lt;strong&gt;Type:&lt;/strong&gt; AccountStatus&lt;br/&gt; The current status of the account from the perspective of this interface. Possible values are:&lt;br/&gt;&amp;bull; &lt;code&gt;UPDATED&lt;/code&gt; means that the account is up to date from finAPI&#39;s point of view. This means that no current import/update is running, and the previous import/update had successfully updated the account&#39;s data (e.g. transactions and securities), and the bank given balance matched the transaction&#39;s calculated sum, meaning that no adjusting entry (&#39;Zwischensaldo&#39; transaction) was inserted.&lt;br/&gt;&amp;bull; &lt;code&gt;UPDATED_FIXED&lt;/code&gt; means that the account is up to date from finAPI&#39;s point of view (no current import/update is running, and the previous import/update had successfully updated the account&#39;s data), BUT there was a deviation in the bank given balance which was fixed by adding an adjusting entry (&#39;Zwischensaldo&#39; transaction).&lt;br/&gt;&amp;bull; &lt;code&gt;DOWNLOAD_IN_PROGRESS&lt;/code&gt; means that the account&#39;s data is currently being imported/updated.&lt;br/&gt;&amp;bull; &lt;code&gt;DOWNLOAD_FAILED&lt;/code&gt; means that the account data was not successfully imported or updated. Possible reasons: finAPI could not get the account&#39;s balance, or it could not parse all transactions/securities, or some internal error has occurred. Also, it could mean that finAPI could not even get to the point of receiving the account data from the bank server, for example because of incorrect login credentials or a network problem. Note however that when we get a balance and just an empty list of transactions or securities, then this is regarded as valid and successful download. The reason for this is that for some accounts that have little activity, we may actually get no recent transactions but only a balance.&lt;br/&gt;&amp;bull; &lt;code&gt;DEPRECATED&lt;/code&gt; means that the account could no longer be matched with any account from the bank server. This can mean either that the account was terminated by the user and is no longer sent by the bank server, or that finAPI could no longer match it because the account&#39;s data (name, type, iban, account number, etc.) has been changed by the bank. | 
**Capabilities** | [**[]AccountCapability**](AccountCapability.md) |  | 
**PaymentCapabilities** | [**AccountInterfacePaymentCapabilities**](AccountInterfacePaymentCapabilities.md) | &lt;strong&gt;Type:&lt;/strong&gt; AccountInterfacePaymentCapabilities&lt;br/&gt; The payment capabilities of this account. | 
**LastSuccessfulUpdate** | **NullableString** | Timestamp of when the account was last successfully updated using this interface (or initially imported); more precisely: time when the account data (balance and positions) has been stored into the finAPI databases. The value is returned in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**LastUpdateAttempt** | **NullableString** | Timestamp of when the account was last tried to be updated using this interface (or initially imported); more precisely: time when the update (or initial import) was triggered. The value is returned in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 

## Methods

### NewAccountInterface

`func NewAccountInterface(interface_ BankingInterface, status AccountStatus, capabilities []AccountCapability, paymentCapabilities AccountInterfacePaymentCapabilities, lastSuccessfulUpdate NullableString, lastUpdateAttempt NullableString, ) *AccountInterface`

NewAccountInterface instantiates a new AccountInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountInterfaceWithDefaults

`func NewAccountInterfaceWithDefaults() *AccountInterface`

NewAccountInterfaceWithDefaults instantiates a new AccountInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterface

`func (o *AccountInterface) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *AccountInterface) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *AccountInterface) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.


### GetStatus

`func (o *AccountInterface) GetStatus() AccountStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccountInterface) GetStatusOk() (*AccountStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccountInterface) SetStatus(v AccountStatus)`

SetStatus sets Status field to given value.


### GetCapabilities

`func (o *AccountInterface) GetCapabilities() []AccountCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AccountInterface) GetCapabilitiesOk() (*[]AccountCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AccountInterface) SetCapabilities(v []AccountCapability)`

SetCapabilities sets Capabilities field to given value.


### GetPaymentCapabilities

`func (o *AccountInterface) GetPaymentCapabilities() AccountInterfacePaymentCapabilities`

GetPaymentCapabilities returns the PaymentCapabilities field if non-nil, zero value otherwise.

### GetPaymentCapabilitiesOk

`func (o *AccountInterface) GetPaymentCapabilitiesOk() (*AccountInterfacePaymentCapabilities, bool)`

GetPaymentCapabilitiesOk returns a tuple with the PaymentCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentCapabilities

`func (o *AccountInterface) SetPaymentCapabilities(v AccountInterfacePaymentCapabilities)`

SetPaymentCapabilities sets PaymentCapabilities field to given value.


### GetLastSuccessfulUpdate

`func (o *AccountInterface) GetLastSuccessfulUpdate() string`

GetLastSuccessfulUpdate returns the LastSuccessfulUpdate field if non-nil, zero value otherwise.

### GetLastSuccessfulUpdateOk

`func (o *AccountInterface) GetLastSuccessfulUpdateOk() (*string, bool)`

GetLastSuccessfulUpdateOk returns a tuple with the LastSuccessfulUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulUpdate

`func (o *AccountInterface) SetLastSuccessfulUpdate(v string)`

SetLastSuccessfulUpdate sets LastSuccessfulUpdate field to given value.


### SetLastSuccessfulUpdateNil

`func (o *AccountInterface) SetLastSuccessfulUpdateNil(b bool)`

 SetLastSuccessfulUpdateNil sets the value for LastSuccessfulUpdate to be an explicit nil

### UnsetLastSuccessfulUpdate
`func (o *AccountInterface) UnsetLastSuccessfulUpdate()`

UnsetLastSuccessfulUpdate ensures that no value is present for LastSuccessfulUpdate, not even an explicit nil
### GetLastUpdateAttempt

`func (o *AccountInterface) GetLastUpdateAttempt() string`

GetLastUpdateAttempt returns the LastUpdateAttempt field if non-nil, zero value otherwise.

### GetLastUpdateAttemptOk

`func (o *AccountInterface) GetLastUpdateAttemptOk() (*string, bool)`

GetLastUpdateAttemptOk returns a tuple with the LastUpdateAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateAttempt

`func (o *AccountInterface) SetLastUpdateAttempt(v string)`

SetLastUpdateAttempt sets LastUpdateAttempt field to given value.


### SetLastUpdateAttemptNil

`func (o *AccountInterface) SetLastUpdateAttemptNil(b bool)`

 SetLastUpdateAttemptNil sets the value for LastUpdateAttempt to be an explicit nil

### UnsetLastUpdateAttempt
`func (o *AccountInterface) UnsetLastUpdateAttempt()`

UnsetLastUpdateAttempt ensures that no value is present for LastUpdateAttempt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


