# SubmitStandingOrderParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StandingOrderId** | **int64** | Standing order identifier | 
**Interface** | [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; Bank interface. Possible values:&lt;br&gt;&lt;br&gt;&amp;bull; &lt;code&gt;FINTS_SERVER&lt;/code&gt; - means that finAPI will execute the standing order via the bank&#39;s FinTS interface.&lt;br&gt;&amp;bull; &lt;code&gt;WEB_SCRAPER&lt;/code&gt; - means that finAPI will parse data from the bank&#39;s online banking website.&lt;br&gt;&amp;bull; &lt;code&gt;XS2A&lt;/code&gt; - means that finAPI will execute the standing order via the bank&#39;s XS2A interface.&lt;br/&gt;To determine what interface(s) you can choose to submit a standing order, please refer to the field paymentCapabilities.sepaStandingOrder in BankInterface.&lt;br/&gt;For standalone standing orders in particular, we suggest to always use XS2A if supported, and only use FINTS_SERVER or WEB_SCRAPER as a fallback, because non-XS2A interfaces might require not just a single, but multiple authentications when submitting the standing order.&lt;br/&gt; | 
**LoginCredentials** | Pointer to [**[]LoginCredential**](LoginCredential.md) | &lt;strong&gt;Type:&lt;/strong&gt; LoginCredential&lt;br/&gt; Login credentials. May not be required when the credentials are stored in finAPI, or when the bank interface has no login credentials. | [optional] 
**RedirectUrl** | Pointer to **string** | Must only be passed when the used interface has the property REDIRECT_APPROACH. The user will be redirected to the given URL from the bank&#39;s website after completing the bank login and (possibly) the SCA. | [optional] 
**MultiStepAuthentication** | Pointer to [**MultiStepAuthenticationCallback**](MultiStepAuthenticationCallback.md) | &lt;strong&gt;Type:&lt;/strong&gt; MultiStepAuthenticationCallback&lt;br/&gt; Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication. | [optional] 

## Methods

### NewSubmitStandingOrderParams

`func NewSubmitStandingOrderParams(standingOrderId int64, interface_ BankingInterface, ) *SubmitStandingOrderParams`

NewSubmitStandingOrderParams instantiates a new SubmitStandingOrderParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitStandingOrderParamsWithDefaults

`func NewSubmitStandingOrderParamsWithDefaults() *SubmitStandingOrderParams`

NewSubmitStandingOrderParamsWithDefaults instantiates a new SubmitStandingOrderParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStandingOrderId

`func (o *SubmitStandingOrderParams) GetStandingOrderId() int64`

GetStandingOrderId returns the StandingOrderId field if non-nil, zero value otherwise.

### GetStandingOrderIdOk

`func (o *SubmitStandingOrderParams) GetStandingOrderIdOk() (*int64, bool)`

GetStandingOrderIdOk returns a tuple with the StandingOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandingOrderId

`func (o *SubmitStandingOrderParams) SetStandingOrderId(v int64)`

SetStandingOrderId sets StandingOrderId field to given value.


### GetInterface

`func (o *SubmitStandingOrderParams) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *SubmitStandingOrderParams) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *SubmitStandingOrderParams) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.


### GetLoginCredentials

`func (o *SubmitStandingOrderParams) GetLoginCredentials() []LoginCredential`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *SubmitStandingOrderParams) GetLoginCredentialsOk() (*[]LoginCredential, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *SubmitStandingOrderParams) SetLoginCredentials(v []LoginCredential)`

SetLoginCredentials sets LoginCredentials field to given value.

### HasLoginCredentials

`func (o *SubmitStandingOrderParams) HasLoginCredentials() bool`

HasLoginCredentials returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SubmitStandingOrderParams) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SubmitStandingOrderParams) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SubmitStandingOrderParams) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SubmitStandingOrderParams) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetMultiStepAuthentication

`func (o *SubmitStandingOrderParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback`

GetMultiStepAuthentication returns the MultiStepAuthentication field if non-nil, zero value otherwise.

### GetMultiStepAuthenticationOk

`func (o *SubmitStandingOrderParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool)`

GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStepAuthentication

`func (o *SubmitStandingOrderParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback)`

SetMultiStepAuthentication sets MultiStepAuthentication field to given value.

### HasMultiStepAuthentication

`func (o *SubmitStandingOrderParams) HasMultiStepAuthentication() bool`

HasMultiStepAuthentication returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


