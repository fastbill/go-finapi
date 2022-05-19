# SubmitPaymentParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **int64** | Payment identifier | 
**Interface** | [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; Bank interface. Possible values:&lt;br&gt;&lt;br&gt;&amp;bull; &lt;code&gt;FINTS_SERVER&lt;/code&gt; - means that finAPI will execute the payment via the bank&#39;s FinTS interface.&lt;br&gt;&amp;bull; &lt;code&gt;WEB_SCRAPER&lt;/code&gt; - means that finAPI will parse data from the bank&#39;s online banking website.&lt;br&gt;&amp;bull; &lt;code&gt;XS2A&lt;/code&gt; - means that finAPI will execute the payment via the bank&#39;s XS2A interface.Please note that XS2A doesn&#39;t support direct debits yet. &lt;br/&gt;To determine what interface(s) you can choose to submit a payment, please refer to the field AccountInterface.capabilities of the account that is related to the payment, or if this is a standalone payment without a related account imported in finAPI, refer to the field BankInterface.isMoneyTransferSupported.&lt;br/&gt;For standalone money transfers (finAPI Payment product) in particular, we suggest to always use XS2A if supported, and only use FINTS_SERVER or WEB_SCRAPER as a fallback, because non-XS2A interfaces might require not just a single, but multiple authentications when submitting the payment.&lt;br/&gt; | 
**LoginCredentials** | Pointer to [**[]LoginCredential**](LoginCredential.md) | &lt;strong&gt;Type:&lt;/strong&gt; LoginCredential&lt;br/&gt; Login credentials. May not be required when the credentials are stored in finAPI, or when the bank interface has no login credentials. | [optional] 
**RedirectUrl** | Pointer to **string** | Must only be passed when the used interface has the property REDIRECT_APPROACH. The user will be redirected to the given URL from the bank&#39;s website after completing the bank login and (possibly) the SCA. | [optional] 
**MultiStepAuthentication** | Pointer to [**MultiStepAuthenticationCallback**](MultiStepAuthenticationCallback.md) | &lt;strong&gt;Type:&lt;/strong&gt; MultiStepAuthenticationCallback&lt;br/&gt; Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication. | [optional] 
**HideTransactionDetailsInWebForm** | Pointer to **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;hidePaymentSummary&#39; field defined &lt;a href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/profiles&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; instead&lt;br/&gt;&lt;br/&gt;Whether the finAPI Web Form should hide transaction details when prompting the caller for the second factor. Default value is false. | [optional] [default to false]
**ForceWebForm** | Pointer to **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;&lt;br/&gt;If the user has stored credentials in finAPI for the account (resp. bank connection) that relates to the payment, then the finAPI Web Form will only be shown when the user must be involved for a second authentication, or when the previous connection to the bank via the selected interface had failed. However if you want to provide the Web Form to the user in any case, you can set this field to true. It will force the Web Form flow for the user and allow him to make changes to the data that he has stored in finAPI. Default value is &#39;false&#39;.&lt;br/&gt;Note that this flag is irrelevant when submitting a standalone payment, as in this case there is no related data stored in finAPI. | [optional] [default to false]

## Methods

### NewSubmitPaymentParams

`func NewSubmitPaymentParams(paymentId int64, interface_ BankingInterface, ) *SubmitPaymentParams`

NewSubmitPaymentParams instantiates a new SubmitPaymentParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitPaymentParamsWithDefaults

`func NewSubmitPaymentParamsWithDefaults() *SubmitPaymentParams`

NewSubmitPaymentParamsWithDefaults instantiates a new SubmitPaymentParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *SubmitPaymentParams) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *SubmitPaymentParams) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *SubmitPaymentParams) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.


### GetInterface

`func (o *SubmitPaymentParams) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *SubmitPaymentParams) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *SubmitPaymentParams) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.


### GetLoginCredentials

`func (o *SubmitPaymentParams) GetLoginCredentials() []LoginCredential`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *SubmitPaymentParams) GetLoginCredentialsOk() (*[]LoginCredential, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *SubmitPaymentParams) SetLoginCredentials(v []LoginCredential)`

SetLoginCredentials sets LoginCredentials field to given value.

### HasLoginCredentials

`func (o *SubmitPaymentParams) HasLoginCredentials() bool`

HasLoginCredentials returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SubmitPaymentParams) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SubmitPaymentParams) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SubmitPaymentParams) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SubmitPaymentParams) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetMultiStepAuthentication

`func (o *SubmitPaymentParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback`

GetMultiStepAuthentication returns the MultiStepAuthentication field if non-nil, zero value otherwise.

### GetMultiStepAuthenticationOk

`func (o *SubmitPaymentParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool)`

GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStepAuthentication

`func (o *SubmitPaymentParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback)`

SetMultiStepAuthentication sets MultiStepAuthentication field to given value.

### HasMultiStepAuthentication

`func (o *SubmitPaymentParams) HasMultiStepAuthentication() bool`

HasMultiStepAuthentication returns a boolean if a field has been set.

### GetHideTransactionDetailsInWebForm

`func (o *SubmitPaymentParams) GetHideTransactionDetailsInWebForm() bool`

GetHideTransactionDetailsInWebForm returns the HideTransactionDetailsInWebForm field if non-nil, zero value otherwise.

### GetHideTransactionDetailsInWebFormOk

`func (o *SubmitPaymentParams) GetHideTransactionDetailsInWebFormOk() (*bool, bool)`

GetHideTransactionDetailsInWebFormOk returns a tuple with the HideTransactionDetailsInWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTransactionDetailsInWebForm

`func (o *SubmitPaymentParams) SetHideTransactionDetailsInWebForm(v bool)`

SetHideTransactionDetailsInWebForm sets HideTransactionDetailsInWebForm field to given value.

### HasHideTransactionDetailsInWebForm

`func (o *SubmitPaymentParams) HasHideTransactionDetailsInWebForm() bool`

HasHideTransactionDetailsInWebForm returns a boolean if a field has been set.

### GetForceWebForm

`func (o *SubmitPaymentParams) GetForceWebForm() bool`

GetForceWebForm returns the ForceWebForm field if non-nil, zero value otherwise.

### GetForceWebFormOk

`func (o *SubmitPaymentParams) GetForceWebFormOk() (*bool, bool)`

GetForceWebFormOk returns a tuple with the ForceWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceWebForm

`func (o *SubmitPaymentParams) SetForceWebForm(v bool)`

SetForceWebForm sets ForceWebForm field to given value.

### HasForceWebForm

`func (o *SubmitPaymentParams) HasForceWebForm() bool`

HasForceWebForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


