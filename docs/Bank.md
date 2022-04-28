# Bank

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Bank identifier.&lt;br/&gt;&lt;br/&gt;NOTE: Do NOT assume that the identifiers of banks are the same across different finAPI environments. In fact, the identifiers may change whenever a new finAPI version is released, even within the same environment. The identifiers are meant to be used for references within the finAPI services only, but not for hard-coding them in your application. If you need to hard-code the usage of a certain bank within your application, please instead refer to the BLZ. | 
**Name** | **string** | Name of bank | 
**LoginHint** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginHint&#39; field in the &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Login hint. Contains a German message for the user that explains what kind of credentials are expected.&lt;br/&gt;&lt;br/&gt;Please note that it is strongly recommended to always show the login hint to the user if there is one, as the credentials that finAPI requires for the bank might be different to the credentials that the user knows from the bank&#39;s website.&lt;br/&gt;&lt;br/&gt;Also note that the contents of this field should always be interpreted as HTML, as the text might contain HTML tags for highlighted words, paragraphs, etc. | 
**Bic** | **NullableString** | BIC of bank | 
**Blzs** | **[]string** |  | 
**Blz** | **string** | BLZ of bank | 
**Location** | **NullableString** | Bank location (two-letter country code; ISO 3166 ALPHA-2). Note that when this field is not set, it means that this bank depicts an international institute which is not bound to any specific country. | 
**City** | **NullableString** | City that this bank is located in. Note that this field may not be set for some banks. | 
**IsSupported** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;isMoneyTransferSupported&#39; and &#39;isAisSupported&#39; fields in the &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Whether this bank is supported by finAPI, i.e. whether you can import/update a bank connection of this bank.&lt;br/&gt;&lt;br/&gt;NOTE: this field only regards FINTS_SERVER and WEB_SCRAPER interfaces. XS2A is not considered. | 
**IsTestBank** | **bool** | If true, then this bank does not depict a real bank, but rather a testing endpoint provided by a bank or by finAPI. You probably want to regard these banks only during the development of your application, but not in production. You can filter out these banks in production by making sure that the &#39;isTestBank&#39; parameter is always set to &#39;false&#39; whenever your application is calling the &#39;Get and search all banks&#39; service. | 
**Popularity** | **int32** | Popularity of this bank with your users (mandator-wide, i.e. across all of your clients). The value equals the number of bank connections that are currently imported for this bank across all of your users (which means it is a constantly adjusting value). You can use this field for statistical evaluation, and also for ordering bank search results (see service &#39;Get and search all banks&#39;). | 
**Health** | **int32** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;health&#39; field in the &#39;interfaces&#39; instead. &lt;br/&gt;&lt;br/&gt;The health status of this bank. This is a value between 0 and 100, depicting the percentage of successful communication attempts with this bank during the latest couple of bank connection imports or updates (across the entire finAPI system). Note that &#39;successful&#39; means that there was no technical error trying to establish a communication with the bank. Non-technical errors (like incorrect credentials) are regarded successful communication attempts. | 
**LoginFieldUserId** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; in the &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Label for the user ID login field, as it is called on the bank&#39;s website (e.g. \&quot;Nutzerkennung\&quot;). If this field is set (i.e. not null) then you should prompt your users to enter the required data in a text field which you can label with this field&#39;s value. | 
**LoginFieldCustomerId** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; in the &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Label for the customer ID login field, as it is called on the bank&#39;s website (e.g. \&quot;Kundennummer\&quot;). If this field is set (i.e. not null) then you should prompt your users to enter the required data in a text field which you can label with this field&#39;s value. | 
**LoginFieldPin** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; in the &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Label for the PIN field, as it is called on the bank&#39;s website (mostly \&quot;PIN\&quot;). If this field is set (i.e. not null) then you should prompt your users to enter the required data in a text field which you can label with this field&#39;s value. | 
**PinsAreVolatile** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;isVolatile&#39; field of the &#39;loginCredentials&#39; in the &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Whether the PINs that are used for authentication with the bank are volatile. If the PINs are volatile, it means that the value for the field, which is provided by the user, is valid only for a single authentication and then gets invalidated on bank-side. If pinsAreVolatile is true, it is not possible to store the PIN in finAPI, as a stored PIN will never work for future authentications. | 
**IsCustomerIdPassword** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;isSecret&#39; field of the &#39;loginCredentials&#39; in &#39;interfaces&#39; instead.&lt;br/&gt;&lt;br/&gt;Whether the banking customer ID has to be treated like a password. Certain banks require a second password (besides the PIN) for the user to login. In this case your application should use a password input field when prompting users for their credentials. | 
**SupportedDataSources** | [**[]SupportedDataSource**](SupportedDataSource.md) |  | 
**Interfaces** | [**[]BankInterface**](BankInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankInterface&lt;br/&gt; Set of interfaces that finAPI can use to connect to the bank. Note that this set will be empty for non-supported banks. Note also that the WEB_SCRAPER interface might be disabled for your client (see GET /clientConfiguration). When this is the case, then finAPI will not use the web scraper for data download, and if the web scraper is the only supported interface of this bank, then finAPI will not allow to download any data for this bank at all (for details, see POST /bankConnections/import and POST /bankConnections/update). | 
**BankGroup** | [**NullableBankGroup**](BankGroup.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankGroup&lt;br/&gt; Bank group | 
**LastCommunicationAttempt** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;lastCommunicationAttempt&#39; field in the &#39;interfaces&#39; instead. &lt;br/&gt;&lt;br/&gt;Time of the last communication attempt with this bank during a bank connection import or update (across the entire finAPI system). The value is returned in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**LastSuccessfulCommunication** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;lastSuccessfulCommunication&#39; field in the &#39;interfaces&#39; instead. &lt;br/&gt;&lt;br/&gt;Time of the last successful communication with this bank during a bank connection import or update (across the entire finAPI system). The value is returned in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**IsBeta** | **bool** | Whether this bank is in beta phase. For more details, please refer to the field ClientConfiguration.betaBanksEnabled. | 

## Methods

### NewBank

`func NewBank(id int64, name string, loginHint NullableString, bic NullableString, blzs []string, blz string, location NullableString, city NullableString, isSupported bool, isTestBank bool, popularity int32, health int32, loginFieldUserId NullableString, loginFieldCustomerId NullableString, loginFieldPin NullableString, pinsAreVolatile bool, isCustomerIdPassword bool, supportedDataSources []SupportedDataSource, interfaces []BankInterface, bankGroup NullableBankGroup, lastCommunicationAttempt NullableString, lastSuccessfulCommunication NullableString, isBeta bool, ) *Bank`

NewBank instantiates a new Bank object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankWithDefaults

`func NewBankWithDefaults() *Bank`

NewBankWithDefaults instantiates a new Bank object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bank) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bank) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bank) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *Bank) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bank) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bank) SetName(v string)`

SetName sets Name field to given value.


### GetLoginHint

`func (o *Bank) GetLoginHint() string`

GetLoginHint returns the LoginHint field if non-nil, zero value otherwise.

### GetLoginHintOk

`func (o *Bank) GetLoginHintOk() (*string, bool)`

GetLoginHintOk returns a tuple with the LoginHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginHint

`func (o *Bank) SetLoginHint(v string)`

SetLoginHint sets LoginHint field to given value.


### SetLoginHintNil

`func (o *Bank) SetLoginHintNil(b bool)`

 SetLoginHintNil sets the value for LoginHint to be an explicit nil

### UnsetLoginHint
`func (o *Bank) UnsetLoginHint()`

UnsetLoginHint ensures that no value is present for LoginHint, not even an explicit nil
### GetBic

`func (o *Bank) GetBic() string`

GetBic returns the Bic field if non-nil, zero value otherwise.

### GetBicOk

`func (o *Bank) GetBicOk() (*string, bool)`

GetBicOk returns a tuple with the Bic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBic

`func (o *Bank) SetBic(v string)`

SetBic sets Bic field to given value.


### SetBicNil

`func (o *Bank) SetBicNil(b bool)`

 SetBicNil sets the value for Bic to be an explicit nil

### UnsetBic
`func (o *Bank) UnsetBic()`

UnsetBic ensures that no value is present for Bic, not even an explicit nil
### GetBlzs

`func (o *Bank) GetBlzs() []string`

GetBlzs returns the Blzs field if non-nil, zero value otherwise.

### GetBlzsOk

`func (o *Bank) GetBlzsOk() (*[]string, bool)`

GetBlzsOk returns a tuple with the Blzs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlzs

`func (o *Bank) SetBlzs(v []string)`

SetBlzs sets Blzs field to given value.


### GetBlz

`func (o *Bank) GetBlz() string`

GetBlz returns the Blz field if non-nil, zero value otherwise.

### GetBlzOk

`func (o *Bank) GetBlzOk() (*string, bool)`

GetBlzOk returns a tuple with the Blz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlz

`func (o *Bank) SetBlz(v string)`

SetBlz sets Blz field to given value.


### GetLocation

`func (o *Bank) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Bank) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Bank) SetLocation(v string)`

SetLocation sets Location field to given value.


### SetLocationNil

`func (o *Bank) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *Bank) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetCity

`func (o *Bank) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Bank) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Bank) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *Bank) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Bank) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetIsSupported

`func (o *Bank) GetIsSupported() bool`

GetIsSupported returns the IsSupported field if non-nil, zero value otherwise.

### GetIsSupportedOk

`func (o *Bank) GetIsSupportedOk() (*bool, bool)`

GetIsSupportedOk returns a tuple with the IsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSupported

`func (o *Bank) SetIsSupported(v bool)`

SetIsSupported sets IsSupported field to given value.


### GetIsTestBank

`func (o *Bank) GetIsTestBank() bool`

GetIsTestBank returns the IsTestBank field if non-nil, zero value otherwise.

### GetIsTestBankOk

`func (o *Bank) GetIsTestBankOk() (*bool, bool)`

GetIsTestBankOk returns a tuple with the IsTestBank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTestBank

`func (o *Bank) SetIsTestBank(v bool)`

SetIsTestBank sets IsTestBank field to given value.


### GetPopularity

`func (o *Bank) GetPopularity() int32`

GetPopularity returns the Popularity field if non-nil, zero value otherwise.

### GetPopularityOk

`func (o *Bank) GetPopularityOk() (*int32, bool)`

GetPopularityOk returns a tuple with the Popularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopularity

`func (o *Bank) SetPopularity(v int32)`

SetPopularity sets Popularity field to given value.


### GetHealth

`func (o *Bank) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *Bank) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *Bank) SetHealth(v int32)`

SetHealth sets Health field to given value.


### GetLoginFieldUserId

`func (o *Bank) GetLoginFieldUserId() string`

GetLoginFieldUserId returns the LoginFieldUserId field if non-nil, zero value otherwise.

### GetLoginFieldUserIdOk

`func (o *Bank) GetLoginFieldUserIdOk() (*string, bool)`

GetLoginFieldUserIdOk returns a tuple with the LoginFieldUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFieldUserId

`func (o *Bank) SetLoginFieldUserId(v string)`

SetLoginFieldUserId sets LoginFieldUserId field to given value.


### SetLoginFieldUserIdNil

`func (o *Bank) SetLoginFieldUserIdNil(b bool)`

 SetLoginFieldUserIdNil sets the value for LoginFieldUserId to be an explicit nil

### UnsetLoginFieldUserId
`func (o *Bank) UnsetLoginFieldUserId()`

UnsetLoginFieldUserId ensures that no value is present for LoginFieldUserId, not even an explicit nil
### GetLoginFieldCustomerId

`func (o *Bank) GetLoginFieldCustomerId() string`

GetLoginFieldCustomerId returns the LoginFieldCustomerId field if non-nil, zero value otherwise.

### GetLoginFieldCustomerIdOk

`func (o *Bank) GetLoginFieldCustomerIdOk() (*string, bool)`

GetLoginFieldCustomerIdOk returns a tuple with the LoginFieldCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFieldCustomerId

`func (o *Bank) SetLoginFieldCustomerId(v string)`

SetLoginFieldCustomerId sets LoginFieldCustomerId field to given value.


### SetLoginFieldCustomerIdNil

`func (o *Bank) SetLoginFieldCustomerIdNil(b bool)`

 SetLoginFieldCustomerIdNil sets the value for LoginFieldCustomerId to be an explicit nil

### UnsetLoginFieldCustomerId
`func (o *Bank) UnsetLoginFieldCustomerId()`

UnsetLoginFieldCustomerId ensures that no value is present for LoginFieldCustomerId, not even an explicit nil
### GetLoginFieldPin

`func (o *Bank) GetLoginFieldPin() string`

GetLoginFieldPin returns the LoginFieldPin field if non-nil, zero value otherwise.

### GetLoginFieldPinOk

`func (o *Bank) GetLoginFieldPinOk() (*string, bool)`

GetLoginFieldPinOk returns a tuple with the LoginFieldPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginFieldPin

`func (o *Bank) SetLoginFieldPin(v string)`

SetLoginFieldPin sets LoginFieldPin field to given value.


### SetLoginFieldPinNil

`func (o *Bank) SetLoginFieldPinNil(b bool)`

 SetLoginFieldPinNil sets the value for LoginFieldPin to be an explicit nil

### UnsetLoginFieldPin
`func (o *Bank) UnsetLoginFieldPin()`

UnsetLoginFieldPin ensures that no value is present for LoginFieldPin, not even an explicit nil
### GetPinsAreVolatile

`func (o *Bank) GetPinsAreVolatile() bool`

GetPinsAreVolatile returns the PinsAreVolatile field if non-nil, zero value otherwise.

### GetPinsAreVolatileOk

`func (o *Bank) GetPinsAreVolatileOk() (*bool, bool)`

GetPinsAreVolatileOk returns a tuple with the PinsAreVolatile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinsAreVolatile

`func (o *Bank) SetPinsAreVolatile(v bool)`

SetPinsAreVolatile sets PinsAreVolatile field to given value.


### GetIsCustomerIdPassword

`func (o *Bank) GetIsCustomerIdPassword() bool`

GetIsCustomerIdPassword returns the IsCustomerIdPassword field if non-nil, zero value otherwise.

### GetIsCustomerIdPasswordOk

`func (o *Bank) GetIsCustomerIdPasswordOk() (*bool, bool)`

GetIsCustomerIdPasswordOk returns a tuple with the IsCustomerIdPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomerIdPassword

`func (o *Bank) SetIsCustomerIdPassword(v bool)`

SetIsCustomerIdPassword sets IsCustomerIdPassword field to given value.


### GetSupportedDataSources

`func (o *Bank) GetSupportedDataSources() []SupportedDataSource`

GetSupportedDataSources returns the SupportedDataSources field if non-nil, zero value otherwise.

### GetSupportedDataSourcesOk

`func (o *Bank) GetSupportedDataSourcesOk() (*[]SupportedDataSource, bool)`

GetSupportedDataSourcesOk returns a tuple with the SupportedDataSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDataSources

`func (o *Bank) SetSupportedDataSources(v []SupportedDataSource)`

SetSupportedDataSources sets SupportedDataSources field to given value.


### GetInterfaces

`func (o *Bank) GetInterfaces() []BankInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Bank) GetInterfacesOk() (*[]BankInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Bank) SetInterfaces(v []BankInterface)`

SetInterfaces sets Interfaces field to given value.


### GetBankGroup

`func (o *Bank) GetBankGroup() BankGroup`

GetBankGroup returns the BankGroup field if non-nil, zero value otherwise.

### GetBankGroupOk

`func (o *Bank) GetBankGroupOk() (*BankGroup, bool)`

GetBankGroupOk returns a tuple with the BankGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankGroup

`func (o *Bank) SetBankGroup(v BankGroup)`

SetBankGroup sets BankGroup field to given value.


### SetBankGroupNil

`func (o *Bank) SetBankGroupNil(b bool)`

 SetBankGroupNil sets the value for BankGroup to be an explicit nil

### UnsetBankGroup
`func (o *Bank) UnsetBankGroup()`

UnsetBankGroup ensures that no value is present for BankGroup, not even an explicit nil
### GetLastCommunicationAttempt

`func (o *Bank) GetLastCommunicationAttempt() string`

GetLastCommunicationAttempt returns the LastCommunicationAttempt field if non-nil, zero value otherwise.

### GetLastCommunicationAttemptOk

`func (o *Bank) GetLastCommunicationAttemptOk() (*string, bool)`

GetLastCommunicationAttemptOk returns a tuple with the LastCommunicationAttempt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicationAttempt

`func (o *Bank) SetLastCommunicationAttempt(v string)`

SetLastCommunicationAttempt sets LastCommunicationAttempt field to given value.


### SetLastCommunicationAttemptNil

`func (o *Bank) SetLastCommunicationAttemptNil(b bool)`

 SetLastCommunicationAttemptNil sets the value for LastCommunicationAttempt to be an explicit nil

### UnsetLastCommunicationAttempt
`func (o *Bank) UnsetLastCommunicationAttempt()`

UnsetLastCommunicationAttempt ensures that no value is present for LastCommunicationAttempt, not even an explicit nil
### GetLastSuccessfulCommunication

`func (o *Bank) GetLastSuccessfulCommunication() string`

GetLastSuccessfulCommunication returns the LastSuccessfulCommunication field if non-nil, zero value otherwise.

### GetLastSuccessfulCommunicationOk

`func (o *Bank) GetLastSuccessfulCommunicationOk() (*string, bool)`

GetLastSuccessfulCommunicationOk returns a tuple with the LastSuccessfulCommunication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessfulCommunication

`func (o *Bank) SetLastSuccessfulCommunication(v string)`

SetLastSuccessfulCommunication sets LastSuccessfulCommunication field to given value.


### SetLastSuccessfulCommunicationNil

`func (o *Bank) SetLastSuccessfulCommunicationNil(b bool)`

 SetLastSuccessfulCommunicationNil sets the value for LastSuccessfulCommunication to be an explicit nil

### UnsetLastSuccessfulCommunication
`func (o *Bank) UnsetLastSuccessfulCommunication()`

UnsetLastSuccessfulCommunication ensures that no value is present for LastSuccessfulCommunication, not even an explicit nil
### GetIsBeta

`func (o *Bank) GetIsBeta() bool`

GetIsBeta returns the IsBeta field if non-nil, zero value otherwise.

### GetIsBetaOk

`func (o *Bank) GetIsBetaOk() (*bool, bool)`

GetIsBetaOk returns a tuple with the IsBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBeta

`func (o *Bank) SetIsBeta(v bool)`

SetIsBeta sets IsBeta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


