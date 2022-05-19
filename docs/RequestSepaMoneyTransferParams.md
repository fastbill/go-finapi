# RequestSepaMoneyTransferParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientName** | Pointer to **string** | Name of the recipient. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the recipient name. Even if the recipient name does not depict the actual registered account holder of the specified recipient account, the money transfer request might still be successful. This field is optional only when you pass a clearing account as the recipient. Otherwise, this field is required. | [optional] 
**RecipientIban** | Pointer to **string** | IBAN of the recipient&#39;s account. This field is optional only when you pass a clearing account as the recipient. Otherwise, this field is required. | [optional] 
**RecipientBic** | Pointer to **string** | BIC of the recipient&#39;s account. Note: This field is optional when you pass a clearing account as the recipient or if the bank connection of the account that you want to transfer money from supports the IBAN-Only money transfer. You can find this out via GET /bankConnections/&lt;id&gt;. If no BIC is given, finAPI will try to recognize it using the given recipientIban value (if it&#39;s given). And then if the result value is not empty, it will be used for the money transfer request independent of whether it is required or not (unless you pass a clearing account, in which case the value will always be ignored). | [optional] 
**ClearingAccountId** | Pointer to **string** | Identifier of a clearing account. If this field is set, then the fields &#39;recipientName&#39;, &#39;recipientIban&#39; and &#39;recipientBic&#39; will be ignored and the recipient account will be the specified clearing account. | [optional] 
**EndToEndId** | Pointer to **string** | End-To-End ID for the transfer transaction | [optional] 
**Amount** | **float64** | The amount to transfer. Must be a positive decimal number with at most two decimal places (e.g. 99.99) | 
**Purpose** | Pointer to **string** | The purpose of the transfer transaction | [optional] 
**SepaPurposeCode** | Pointer to **string** | SEPA purpose code, according to ISO 20022, external codes set. | [optional] 
**AccountId** | **int64** | Identifier of the bank account that you want to transfer money from | 
**BankingPin** | Pointer to **string** | Online banking PIN. Any symbols are allowed. Max length: 170. If a PIN is stored in the bank connection, then this field may remain unset. If finAPI&#39;s Web Form is not required and the field is set though then it will always be used (even if there is some other PIN stored in the bank connection). If you want the user to enter a PIN in finAPI&#39;s Web Form even when a PIN is stored, then just set the field to any value, so that the service recognizes that you wish to use the Web Form flow. | [optional] 
**StoreSecrets** | Pointer to **bool** | Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when executing SEPA orders. Default value is &#39;false&#39;. &lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt; - before you set this field to true, please regard the &#39;pinsAreVolatile&#39; flag of the bank connection that the account belongs to. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;&lt;br/&gt; - this field is ignored in case when the user will need to use finAPI&#39;s Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the &#39;storeSecretsAvailableInWebForm&#39; setting (see Client Configuration). | [optional] [default to false]
**TwoStepProcedureId** | Pointer to **string** | The bank-given ID of the two-step-procedure that should be used for the order. For a list of available two-step-procedures, see the corresponding bank connection (GET /bankConnections). If this field is not set, then the bank connection&#39;s default two-step-procedure will be used. Note that in this case, when the bank connection has no default two-step-procedure set, then the response of the service depends on whether you need to use finAPI&#39;s Web Form or not. If you need to use the Web Form, the user will be prompted to select the two-step-procedure within the Web Form. If you don&#39;t need to use the Web Form, then the service will return an error (passing a value for this field is required in this case). | [optional] 
**ExecutionDate** | Pointer to **string** | Execution date for the money transfer(s), in the format &#39;YYYY-MM-DD&#39;. If not specified, then the current date will be used. | [optional] 
**SingleBooking** | Pointer to **bool** | This field is only regarded when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of &#39;false&#39;), or as single bookings (in case of &#39;true&#39;). Default value is &#39;false&#39;. | [optional] [default to false]
**AdditionalMoneyTransfers** | Pointer to [**[]SingleMoneyTransferRecipientData**](SingleMoneyTransferRecipientData.md) | &lt;strong&gt;Type:&lt;/strong&gt; SingleMoneyTransferRecipientData&lt;br/&gt; In case that you want to submit not just a single money transfer, but do a collective money transfer, use this field to pass a list of additional money transfer orders. The service will then pass a collective money transfer request to the bank, including both the money transfer specified on the top-level, as well as all money transfers specified in this list. The maximum count of money transfers that you can pass (in total) is 15000. Note that you should check the account&#39;s &#39;supportedOrders&#39; field to find out whether or not it is supporting collective money transfers. | [optional] 
**ChallengeResponse** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;multiStepAuthentication&#39; field instead.&lt;br/&gt;&lt;br/&gt;Challenge response. This field should be set only when the previous attempt to request a SEPA money transfer failed with HTTP code 510, i.e. the bank sent a challenge for the user for an additional authentication. In this case, this field must contain the response to the bank&#39;s challenge. Please note that in case of using finAPI&#39;s Web Form you have to leave this field unset and the application will automatically recognize that the user has to input challenge response and then a Web Form will be shown to the user. | [optional] 
**HideTransactionDetailsInWebForm** | Pointer to **bool** | Whether the finAPI Web Form should hide transaction details when prompting the caller for the second factor. Default value is false. | [optional] [default to false]
**MultiStepAuthentication** | Pointer to [**MultiStepAuthenticationCallback**](MultiStepAuthenticationCallback.md) | &lt;strong&gt;Type:&lt;/strong&gt; MultiStepAuthenticationCallback&lt;br/&gt; Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication. | [optional] 
**StorePin** | Pointer to **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;storeSecrets&#39; field instead.&lt;br/&gt;&lt;br/&gt;Whether to store the PIN. If the PIN is stored, it is not required to pass the PIN again when executing SEPA orders. Default value is &#39;false&#39;. &lt;br/&gt;&lt;br/&gt;NOTES:&lt;br/&gt; - before you set this field to true, please regard the &#39;pinsAreVolatile&#39; flag of the bank connection that the account belongs to. Please note that volatile credentials will not be stored, even if provided, to enforce user involvement in the next communication with the bank;&lt;br/&gt; - this field is ignored in case when the user will need to use finAPI&#39;s Web Form. The user will be able to decide whether to store the PIN or not in the Web Form, depending on the &#39;storeSecretsAvailableInWebForm&#39; setting (see Client Configuration). | [optional] [default to false]

## Methods

### NewRequestSepaMoneyTransferParams

`func NewRequestSepaMoneyTransferParams(amount float64, accountId int64, ) *RequestSepaMoneyTransferParams`

NewRequestSepaMoneyTransferParams instantiates a new RequestSepaMoneyTransferParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSepaMoneyTransferParamsWithDefaults

`func NewRequestSepaMoneyTransferParamsWithDefaults() *RequestSepaMoneyTransferParams`

NewRequestSepaMoneyTransferParamsWithDefaults instantiates a new RequestSepaMoneyTransferParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientName

`func (o *RequestSepaMoneyTransferParams) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *RequestSepaMoneyTransferParams) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *RequestSepaMoneyTransferParams) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *RequestSepaMoneyTransferParams) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### GetRecipientIban

`func (o *RequestSepaMoneyTransferParams) GetRecipientIban() string`

GetRecipientIban returns the RecipientIban field if non-nil, zero value otherwise.

### GetRecipientIbanOk

`func (o *RequestSepaMoneyTransferParams) GetRecipientIbanOk() (*string, bool)`

GetRecipientIbanOk returns a tuple with the RecipientIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIban

`func (o *RequestSepaMoneyTransferParams) SetRecipientIban(v string)`

SetRecipientIban sets RecipientIban field to given value.

### HasRecipientIban

`func (o *RequestSepaMoneyTransferParams) HasRecipientIban() bool`

HasRecipientIban returns a boolean if a field has been set.

### GetRecipientBic

`func (o *RequestSepaMoneyTransferParams) GetRecipientBic() string`

GetRecipientBic returns the RecipientBic field if non-nil, zero value otherwise.

### GetRecipientBicOk

`func (o *RequestSepaMoneyTransferParams) GetRecipientBicOk() (*string, bool)`

GetRecipientBicOk returns a tuple with the RecipientBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientBic

`func (o *RequestSepaMoneyTransferParams) SetRecipientBic(v string)`

SetRecipientBic sets RecipientBic field to given value.

### HasRecipientBic

`func (o *RequestSepaMoneyTransferParams) HasRecipientBic() bool`

HasRecipientBic returns a boolean if a field has been set.

### GetClearingAccountId

`func (o *RequestSepaMoneyTransferParams) GetClearingAccountId() string`

GetClearingAccountId returns the ClearingAccountId field if non-nil, zero value otherwise.

### GetClearingAccountIdOk

`func (o *RequestSepaMoneyTransferParams) GetClearingAccountIdOk() (*string, bool)`

GetClearingAccountIdOk returns a tuple with the ClearingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingAccountId

`func (o *RequestSepaMoneyTransferParams) SetClearingAccountId(v string)`

SetClearingAccountId sets ClearingAccountId field to given value.

### HasClearingAccountId

`func (o *RequestSepaMoneyTransferParams) HasClearingAccountId() bool`

HasClearingAccountId returns a boolean if a field has been set.

### GetEndToEndId

`func (o *RequestSepaMoneyTransferParams) GetEndToEndId() string`

GetEndToEndId returns the EndToEndId field if non-nil, zero value otherwise.

### GetEndToEndIdOk

`func (o *RequestSepaMoneyTransferParams) GetEndToEndIdOk() (*string, bool)`

GetEndToEndIdOk returns a tuple with the EndToEndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndId

`func (o *RequestSepaMoneyTransferParams) SetEndToEndId(v string)`

SetEndToEndId sets EndToEndId field to given value.

### HasEndToEndId

`func (o *RequestSepaMoneyTransferParams) HasEndToEndId() bool`

HasEndToEndId returns a boolean if a field has been set.

### GetAmount

`func (o *RequestSepaMoneyTransferParams) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *RequestSepaMoneyTransferParams) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *RequestSepaMoneyTransferParams) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *RequestSepaMoneyTransferParams) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *RequestSepaMoneyTransferParams) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *RequestSepaMoneyTransferParams) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *RequestSepaMoneyTransferParams) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSepaPurposeCode

`func (o *RequestSepaMoneyTransferParams) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *RequestSepaMoneyTransferParams) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *RequestSepaMoneyTransferParams) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *RequestSepaMoneyTransferParams) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.

### GetAccountId

`func (o *RequestSepaMoneyTransferParams) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RequestSepaMoneyTransferParams) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RequestSepaMoneyTransferParams) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetBankingPin

`func (o *RequestSepaMoneyTransferParams) GetBankingPin() string`

GetBankingPin returns the BankingPin field if non-nil, zero value otherwise.

### GetBankingPinOk

`func (o *RequestSepaMoneyTransferParams) GetBankingPinOk() (*string, bool)`

GetBankingPinOk returns a tuple with the BankingPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingPin

`func (o *RequestSepaMoneyTransferParams) SetBankingPin(v string)`

SetBankingPin sets BankingPin field to given value.

### HasBankingPin

`func (o *RequestSepaMoneyTransferParams) HasBankingPin() bool`

HasBankingPin returns a boolean if a field has been set.

### GetStoreSecrets

`func (o *RequestSepaMoneyTransferParams) GetStoreSecrets() bool`

GetStoreSecrets returns the StoreSecrets field if non-nil, zero value otherwise.

### GetStoreSecretsOk

`func (o *RequestSepaMoneyTransferParams) GetStoreSecretsOk() (*bool, bool)`

GetStoreSecretsOk returns a tuple with the StoreSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSecrets

`func (o *RequestSepaMoneyTransferParams) SetStoreSecrets(v bool)`

SetStoreSecrets sets StoreSecrets field to given value.

### HasStoreSecrets

`func (o *RequestSepaMoneyTransferParams) HasStoreSecrets() bool`

HasStoreSecrets returns a boolean if a field has been set.

### GetTwoStepProcedureId

`func (o *RequestSepaMoneyTransferParams) GetTwoStepProcedureId() string`

GetTwoStepProcedureId returns the TwoStepProcedureId field if non-nil, zero value otherwise.

### GetTwoStepProcedureIdOk

`func (o *RequestSepaMoneyTransferParams) GetTwoStepProcedureIdOk() (*string, bool)`

GetTwoStepProcedureIdOk returns a tuple with the TwoStepProcedureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoStepProcedureId

`func (o *RequestSepaMoneyTransferParams) SetTwoStepProcedureId(v string)`

SetTwoStepProcedureId sets TwoStepProcedureId field to given value.

### HasTwoStepProcedureId

`func (o *RequestSepaMoneyTransferParams) HasTwoStepProcedureId() bool`

HasTwoStepProcedureId returns a boolean if a field has been set.

### GetExecutionDate

`func (o *RequestSepaMoneyTransferParams) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *RequestSepaMoneyTransferParams) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *RequestSepaMoneyTransferParams) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *RequestSepaMoneyTransferParams) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetSingleBooking

`func (o *RequestSepaMoneyTransferParams) GetSingleBooking() bool`

GetSingleBooking returns the SingleBooking field if non-nil, zero value otherwise.

### GetSingleBookingOk

`func (o *RequestSepaMoneyTransferParams) GetSingleBookingOk() (*bool, bool)`

GetSingleBookingOk returns a tuple with the SingleBooking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleBooking

`func (o *RequestSepaMoneyTransferParams) SetSingleBooking(v bool)`

SetSingleBooking sets SingleBooking field to given value.

### HasSingleBooking

`func (o *RequestSepaMoneyTransferParams) HasSingleBooking() bool`

HasSingleBooking returns a boolean if a field has been set.

### GetAdditionalMoneyTransfers

`func (o *RequestSepaMoneyTransferParams) GetAdditionalMoneyTransfers() []SingleMoneyTransferRecipientData`

GetAdditionalMoneyTransfers returns the AdditionalMoneyTransfers field if non-nil, zero value otherwise.

### GetAdditionalMoneyTransfersOk

`func (o *RequestSepaMoneyTransferParams) GetAdditionalMoneyTransfersOk() (*[]SingleMoneyTransferRecipientData, bool)`

GetAdditionalMoneyTransfersOk returns a tuple with the AdditionalMoneyTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMoneyTransfers

`func (o *RequestSepaMoneyTransferParams) SetAdditionalMoneyTransfers(v []SingleMoneyTransferRecipientData)`

SetAdditionalMoneyTransfers sets AdditionalMoneyTransfers field to given value.

### HasAdditionalMoneyTransfers

`func (o *RequestSepaMoneyTransferParams) HasAdditionalMoneyTransfers() bool`

HasAdditionalMoneyTransfers returns a boolean if a field has been set.

### GetChallengeResponse

`func (o *RequestSepaMoneyTransferParams) GetChallengeResponse() string`

GetChallengeResponse returns the ChallengeResponse field if non-nil, zero value otherwise.

### GetChallengeResponseOk

`func (o *RequestSepaMoneyTransferParams) GetChallengeResponseOk() (*string, bool)`

GetChallengeResponseOk returns a tuple with the ChallengeResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeResponse

`func (o *RequestSepaMoneyTransferParams) SetChallengeResponse(v string)`

SetChallengeResponse sets ChallengeResponse field to given value.

### HasChallengeResponse

`func (o *RequestSepaMoneyTransferParams) HasChallengeResponse() bool`

HasChallengeResponse returns a boolean if a field has been set.

### GetHideTransactionDetailsInWebForm

`func (o *RequestSepaMoneyTransferParams) GetHideTransactionDetailsInWebForm() bool`

GetHideTransactionDetailsInWebForm returns the HideTransactionDetailsInWebForm field if non-nil, zero value otherwise.

### GetHideTransactionDetailsInWebFormOk

`func (o *RequestSepaMoneyTransferParams) GetHideTransactionDetailsInWebFormOk() (*bool, bool)`

GetHideTransactionDetailsInWebFormOk returns a tuple with the HideTransactionDetailsInWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideTransactionDetailsInWebForm

`func (o *RequestSepaMoneyTransferParams) SetHideTransactionDetailsInWebForm(v bool)`

SetHideTransactionDetailsInWebForm sets HideTransactionDetailsInWebForm field to given value.

### HasHideTransactionDetailsInWebForm

`func (o *RequestSepaMoneyTransferParams) HasHideTransactionDetailsInWebForm() bool`

HasHideTransactionDetailsInWebForm returns a boolean if a field has been set.

### GetMultiStepAuthentication

`func (o *RequestSepaMoneyTransferParams) GetMultiStepAuthentication() MultiStepAuthenticationCallback`

GetMultiStepAuthentication returns the MultiStepAuthentication field if non-nil, zero value otherwise.

### GetMultiStepAuthenticationOk

`func (o *RequestSepaMoneyTransferParams) GetMultiStepAuthenticationOk() (*MultiStepAuthenticationCallback, bool)`

GetMultiStepAuthenticationOk returns a tuple with the MultiStepAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiStepAuthentication

`func (o *RequestSepaMoneyTransferParams) SetMultiStepAuthentication(v MultiStepAuthenticationCallback)`

SetMultiStepAuthentication sets MultiStepAuthentication field to given value.

### HasMultiStepAuthentication

`func (o *RequestSepaMoneyTransferParams) HasMultiStepAuthentication() bool`

HasMultiStepAuthentication returns a boolean if a field has been set.

### GetStorePin

`func (o *RequestSepaMoneyTransferParams) GetStorePin() bool`

GetStorePin returns the StorePin field if non-nil, zero value otherwise.

### GetStorePinOk

`func (o *RequestSepaMoneyTransferParams) GetStorePinOk() (*bool, bool)`

GetStorePinOk returns a tuple with the StorePin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePin

`func (o *RequestSepaMoneyTransferParams) SetStorePin(v bool)`

SetStorePin sets StorePin field to given value.

### HasStorePin

`func (o *RequestSepaMoneyTransferParams) HasStorePin() bool`

HasStorePin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


