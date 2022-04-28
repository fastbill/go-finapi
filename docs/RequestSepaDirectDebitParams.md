# RequestSepaDirectDebitParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | Identifier of the bank account to which you want to transfer the money. | 
**BankingPin** | Pointer to **string** | Online banking PIN. Any symbols are allowed. Max length: 170. If a PIN is stored in the bank connection, then this field may remain unset. If finAPI&#39;s Web Form is not required and the field is set though then it will always be used (even if there is some other PIN stored in the bank connection). If you want the user to enter a PIN in finAPI&#39;s Web Form even when a PIN is stored, then just set the field to any value, so that the service recognizes that you wish to use the Web Form flow. | [optional] 
**StoreSecrets** | Pointer to **bool** | Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when executing SEPA orders. Default value is &#39;false&#39;. &lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt; - before you set this field to true, please regard the &#39;pinsAreVolatile&#39; flag of the bank connection that the account belongs to. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;&lt;br/&gt; - this field is ignored in case when the user will need to use finAPI&#39;s Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the &#39;storeSecretsAvailableInWebForm&#39; setting (see Client Configuration). | [optional] [default to false]
**TwoStepProcedureId** | Pointer to **string** | The bank-given ID of the two-step-procedure that should be used for the order. For a list of available two-step-procedures, see the corresponding bank connection (GET /bankConnections). If this field is not set, then the bank connection&#39;s default two-step-procedure will be used. Note that in this case, when the bank connection has no default two-step-procedure set, then the response of the service depends on whether you need to use finAPI&#39;s Web Form or not. If you need to use the Web Form, the user will be prompted to select the two-step-procedure within the Web Form. If you don&#39;t need to use the Web Form, then the service will return an error (passing a value for this field is required in this case). | [optional] 
**DirectDebitType** | [**DirectDebitType**](DirectDebitType.md) | &lt;strong&gt;Type:&lt;/strong&gt; DirectDebitType&lt;br/&gt; Type of the direct debit; either &lt;code&gt;BASIC&lt;/code&gt; or &lt;code&gt;B2B&lt;/code&gt; (Business-To-Business). Please note that an account which supports the basic type must not necessarily support B2B (or vice versa). Check the source account&#39;s &#39;supportedOrders&#39; field to find out which types of direct debit it supports.&lt;br/&gt;&lt;br/&gt; | 
**SequenceType** | [**DirectDebitSequenceType**](DirectDebitSequenceType.md) | &lt;strong&gt;Type:&lt;/strong&gt; DirectDebitSequenceType&lt;br/&gt; Sequence type of the direct debit. Possible values:&lt;br/&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;OOFF&lt;/code&gt; - means that this is a one-time direct debit order&lt;br/&gt;&amp;bull; &lt;code&gt;FRST&lt;/code&gt; - means that this is the first in a row of multiple direct debit orders&lt;br/&gt;&amp;bull; &lt;code&gt;RCUR&lt;/code&gt; - means that this is one (but not the first or final) within a row of multiple direct debit orders&lt;br/&gt;&amp;bull; &lt;code&gt;FNAL&lt;/code&gt; - means that this is the final in a row of multiple direct debit orders&lt;br/&gt;&lt;br/&gt; | 
**ExecutionDate** | **string** | Execution date for the direct debit(s), in the format &#39;YYYY-MM-DD&#39;. | 
**SingleBooking** | Pointer to **bool** | This field is only regarded when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of &#39;false&#39;), or as single bookings (in case of &#39;true&#39;). Default value is &#39;false&#39;. | [optional] [default to false]
**DirectDebits** | [**[]SingleDirectDebitData**](SingleDirectDebitData.md) | &lt;strong&gt;Type:&lt;/strong&gt; SingleDirectDebitData&lt;br/&gt; List of the direct debits that you want to execute (may contain at most 15000 items). Please check the account&#39;s &#39;supportedOrders&#39; field to find out whether you can pass multiple direct debits or just one. | 
**HideTransactionDetailsInWebForm** | Pointer to **bool** | Whether the finAPI Web Form should hide transaction details when prompting the caller for the second factor. Default value is false. | [optional] [default to false]
**MultiStepAuthentication** | Pointer to [**MultiStepAuthenticationCallback**](MultiStepAuthenticationCallback.md) | &lt;strong&gt;Type:&lt;/strong&gt; MultiStepAuthenticationCallback&lt;br/&gt; Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication. | [optional] 
**StorePin** | Pointer to **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;storeSecrets&#39; field instead.&lt;br/&gt;&lt;br/&gt;Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when executing SEPA orders. Default value is &#39;false&#39;. &lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt; - before you set this field to true, please regard the &#39;pinsAreVolatile&#39; flag of the bank connection that the account belongs to. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;&lt;br/&gt; - this field is ignored in case when the user will need to use finAPI&#39;s Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the &#39;storeSecretsAvailableInWebForm&#39; setting (see Client Configuration). | [optional] [default to false]

## Methods

### NewRequestSepaDirectDebitParams

`func NewRequestSepaDirectDebitParams(accountId int64, directDebitType DirectDebitType, sequenceType DirectDebitSequenceType, executionDate string, directDebits []SingleDirectDebitData, ) *RequestSepaDirectDebitParams`

NewRequestSepaDirectDebitParams instantiates a new RequestSepaDirectDebitParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSepaDirectDebitParamsWithDefaults

`func NewRequestSepaDirectDebitParamsWithDefaults() *RequestSepaDirectDebitParams`

NewRequestSepaDirectDebitParamsWithDefaults instantiates a new RequestSepaDirectDebitParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RequestSepaDirectDebitParams) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RequestSepaDirectDebitParams) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RequestSepaDirectDebitParams) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetBankingPin

`func (o *RequestSepaDirectDebitParams) GetBankingPin() string`

GetBankingPin returns the BankingPin field if non-nil, zero value otherwise.

### GetBankingPinOk

`func (o *RequestSepaDirectDebitParams) GetBankingPinOk() (*string, bool)`

GetBankingPinOk returns a tuple with the BankingPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingPin

`func (o *RequestSepaDirectDebitParams) SetBankingPin(v string)`

SetBankingPin sets BankingPin field to given value.

### HasBankingPin

`func (o *RequestSepaDirectDebitParams) HasBankingPin() bool`

HasBankingPin returns a boolean if a field has been set.

### GetStoreSecrets

`func (o *RequestSepaDirectDebitParams) GetStoreSecrets() bool`

GetStoreSecrets returns the StoreSecrets field if non-nil, zero value otherwise.

### GetStoreSecretsOk

`func (o *RequestSepaDirectDebitParams) GetStoreSecretsOk() (*bool, bool)`

GetStoreSecretsOk returns a tuple with the StoreSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSecrets

`func (o *RequestSepaDirectDebitParams) SetStoreSecrets(v bool)`

SetStoreSecrets sets StoreSecrets field to given value.

### HasStoreSecrets

`func (o *RequestSepaDirectDebitParams) HasStoreSecrets() bool`

HasStoreSecrets returns a boolean if a field has been set.

### GetTwoStepProcedureId

`func (o *RequestSepaDirectDebitParams) GetTwoStepProcedureId() string`

GetTwoStepProcedureId returns the TwoStepProcedureId field if non-nil, zero value otherwise.

### GetTwoStepProcedureIdOk

`func (o *RequestSepaDirectDebitParams) GetTwoStepProcedureIdOk() (*string, bool)`

GetTwoStepProcedureIdOk returns a tuple with the TwoStepProcedureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoStepProcedureId

`func (o *RequestSepaDirectDebitParams) SetTwoStepProcedureId(v string)`

SetTwoStepProcedureId sets TwoStepProcedureId field to given value.

### HasTwoStepProcedureId

`func (o *RequestSepaDirectDebitParams) HasTwoStepProcedureId() bool`

HasTwoStepProcedureId returns a boolean if a field has been set.

### GetDirectDebitType

`func (o *RequestSepaDirectDebitParams) GetDirectDebitType() DirectDebitType`

GetDirectDebitType returns the DirectDebitType field if non-nil, zero value otherwise.

### GetDirectDebitTypeOk

`func (o *RequestSepaDirectDebitParams) GetDirectDebitTypeOk() (*DirectDebitType, bool)`

GetDirectDebitTypeOk returns a tuple with the DirectDebitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebitType

`func (o *RequestSepaDirectDebitParams) SetDirectDebitType(v DirectDebitType)`

SetDirectDebitType sets DirectDebitType field to given value.


### GetSequenceType

`func (o *RequestSepaDirectDebitParams) GetSequenceType() DirectDebitSequenceType`

GetSequenceType returns the SequenceType field if non-nil, zero value otherwise.

### GetSequenceTypeOk

`func (o *RequestSepaDirectDebitParams) GetSequenceTypeOk() (*DirectDebitSequenceType, bool)`

GetSequenceTypeOk returns a tuple with the SequenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceType

`func (o *RequestSepaDirectDebitParams) SetSequenceType(v DirectDebitSequenceType)`

SetSequenceType sets SequenceType field to given value.


### GetExecutionDate

`func (o *RequestSepaDirectDebitParams) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *RequestSepaDirectDebitParams) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *RequestSepaDirectDebitParams) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.


### GetSingleBooking

`func (o *RequestSepaDirectDebitParams) GetSingleBooking() bool`

GetSingleBooking returns the SingleBooking field if non-nil, zero value otherwise.

### GetSingleBookingOk

`func (o *RequestSepaDirectDebitParams) GetSingleBookingOk() (*bool, bool)`

GetSingleBookingOk returns a tuple with the SingleBooking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleBooking

`func (o *RequestSepaDirectDebitParams) SetSingleBooking(v bool)`

SetSingleBooking sets SingleBooking field to given value.

### HasSingleBooking

`func (o *RequestSepaDirectDebitParams) HasSingleBooking() bool`

HasSingleBooking returns a boolean if a field has been set.

### GetDirectDebits

`func (o *RequestSepaDirectDebitParams) GetDirectDebits() []SingleDirectDebitData`

GetDirectDebits returns the DirectDebits field if non-nil, zero value otherwise.

### GetDirectDebitsOk

`func (o *RequestSepaDirectDebitParams) GetDirectDebitsOk() (*[]SingleDirectDebitData, bool)`

GetDirectDebitsOk returns a tuple with the DirectDebits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebits

`func (o *RequestSepaDirectDebitParams) SetDirectDebits(v []SingleDirectDebitData)`

SetDirectDebits sets DirectDebits field to given value.


### GetHideTransactionDetailsInWebForm

`func (o *RequestSepaDirectDebitParams) GetHideTransactionDetailsInWebForm() bool`

GetHideTransactionDetailsInWebForm returns the HideTransactionDetailsInWebForm field if non-nil, zero value otherwise.

### GetHideTransactionDetailsInWebFormOk

`func (o *RequestSepaDirectDebitParams) GetHideTransactionDetailsInWebFormOk() (*bool, bool)`

GetHideTransactionDetailsInWebFormOk returns a tuple with the HideTransactionDetailsInWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTransactionDetailsInWebForm

`func (o *RequestSepaDirectDebitParams) SetHideTransactionDetailsInWebForm(v bool)`

SetHideTransactionDetailsInWebForm sets HideTransactionDetailsInWebForm field to given value.

### HasHideTransactionDetailsInWebForm

`func (o *RequestSepaDirectDebitParams) HasHideTransactionDetailsInWebForm() bool`

HasHideTransactionDetailsInWebForm returns a boolean if a field has been set.

### GetMultiStepAuthentication

`func (o *RequestSepaDirectDebitParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback`

GetMultiStepAuthentication returns the MultiStepAuthentication field if non-nil, zero value otherwise.

### GetMultiStepAuthenticationOk

`func (o *RequestSepaDirectDebitParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool)`

GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStepAuthentication

`func (o *RequestSepaDirectDebitParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback)`

SetMultiStepAuthentication sets MultiStepAuthentication field to given value.

### HasMultiStepAuthentication

`func (o *RequestSepaDirectDebitParams) HasMultiStepAuthentication() bool`

HasMultiStepAuthentication returns a boolean if a field has been set.

### GetStorePin

`func (o *RequestSepaDirectDebitParams) GetStorePin() bool`

GetStorePin returns the StorePin field if non-nil, zero value otherwise.

### GetStorePinOk

`func (o *RequestSepaDirectDebitParams) GetStorePinOk() (*bool, bool)`

GetStorePinOk returns a tuple with the StorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePin

`func (o *RequestSepaDirectDebitParams) SetStorePin(v bool)`

SetStorePin sets StorePin field to given value.

### HasStorePin

`func (o *RequestSepaDirectDebitParams) HasStorePin() bool`

HasStorePin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


