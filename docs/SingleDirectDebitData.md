# SingleDirectDebitData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebitorName** | **string** | Name of the debitor. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the debitor name. Even if the debitor name does not depict the actual registered account holder of the specified debitor account, the direct debit request might still be successful. | 
**DebitorIban** | **string** | IBAN of the debitor&#39;s account | 
**DebitorBic** | Pointer to **string** | BIC of the debitor&#39;s account. Note: This field is optional if - and only if - the bank connection of the account that you want to transfer money to supports the IBAN-Only direct debit. You can find this out via GET /bankConnections/&lt;id&gt;. If no BIC is given, finAPI will try to recognize it using the given debitorIban value (if it&#39;s given). And then if the result value is not empty, it will be used for the direct debit request independent of whether it is required or not. | [optional] 
**Amount** | **float64** | The amount to transfer. Must be a positive decimal number with at most two decimal places (e.g. 99.99) | 
**Purpose** | Pointer to **string** | The purpose of the transfer transaction | [optional] 
**SepaPurposeCode** | Pointer to **string** | SEPA purpose code, according to ISO 20022, external codes set. | [optional] 
**MandateId** | **string** | Mandate ID that this direct debit order is based on. | 
**MandateDate** | **string** | Date of the mandate that this direct debit order is based on, in the format &#39;YYYY-MM-DD&#39; | 
**CreditorId** | Pointer to **string** | Creditor ID of the source account&#39;s holder | [optional] 
**EndToEndId** | Pointer to **string** | End-To-End ID for the transfer transaction | [optional] 

## Methods

### NewSingleDirectDebitData

`func NewSingleDirectDebitData(debitorName string, debitorIban string, amount float64, mandateId string, mandateDate string, ) *SingleDirectDebitData`

NewSingleDirectDebitData instantiates a new SingleDirectDebitData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleDirectDebitDataWithDefaults

`func NewSingleDirectDebitDataWithDefaults() *SingleDirectDebitData`

NewSingleDirectDebitDataWithDefaults instantiates a new SingleDirectDebitData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDebitorName

`func (o *SingleDirectDebitData) GetDebitorName() string`

GetDebitorName returns the DebitorName field if non-nil, zero value otherwise.

### GetDebitorNameOk

`func (o *SingleDirectDebitData) GetDebitorNameOk() (*string, bool)`

GetDebitorNameOk returns a tuple with the DebitorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitorName

`func (o *SingleDirectDebitData) SetDebitorName(v string)`

SetDebitorName sets DebitorName field to given value.


### GetDebitorIban

`func (o *SingleDirectDebitData) GetDebitorIban() string`

GetDebitorIban returns the DebitorIban field if non-nil, zero value otherwise.

### GetDebitorIbanOk

`func (o *SingleDirectDebitData) GetDebitorIbanOk() (*string, bool)`

GetDebitorIbanOk returns a tuple with the DebitorIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitorIban

`func (o *SingleDirectDebitData) SetDebitorIban(v string)`

SetDebitorIban sets DebitorIban field to given value.


### GetDebitorBic

`func (o *SingleDirectDebitData) GetDebitorBic() string`

GetDebitorBic returns the DebitorBic field if non-nil, zero value otherwise.

### GetDebitorBicOk

`func (o *SingleDirectDebitData) GetDebitorBicOk() (*string, bool)`

GetDebitorBicOk returns a tuple with the DebitorBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebitorBic

`func (o *SingleDirectDebitData) SetDebitorBic(v string)`

SetDebitorBic sets DebitorBic field to given value.

### HasDebitorBic

`func (o *SingleDirectDebitData) HasDebitorBic() bool`

HasDebitorBic returns a boolean if a field has been set.

### GetAmount

`func (o *SingleDirectDebitData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SingleDirectDebitData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SingleDirectDebitData) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *SingleDirectDebitData) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *SingleDirectDebitData) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *SingleDirectDebitData) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *SingleDirectDebitData) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSepaPurposeCode

`func (o *SingleDirectDebitData) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *SingleDirectDebitData) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *SingleDirectDebitData) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *SingleDirectDebitData) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.

### GetMandateId

`func (o *SingleDirectDebitData) GetMandateId() string`

GetMandateId returns the MandateId field if non-nil, zero value otherwise.

### GetMandateIdOk

`func (o *SingleDirectDebitData) GetMandateIdOk() (*string, bool)`

GetMandateIdOk returns a tuple with the MandateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateId

`func (o *SingleDirectDebitData) SetMandateId(v string)`

SetMandateId sets MandateId field to given value.


### GetMandateDate

`func (o *SingleDirectDebitData) GetMandateDate() string`

GetMandateDate returns the MandateDate field if non-nil, zero value otherwise.

### GetMandateDateOk

`func (o *SingleDirectDebitData) GetMandateDateOk() (*string, bool)`

GetMandateDateOk returns a tuple with the MandateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandateDate

`func (o *SingleDirectDebitData) SetMandateDate(v string)`

SetMandateDate sets MandateDate field to given value.


### GetCreditorId

`func (o *SingleDirectDebitData) GetCreditorId() string`

GetCreditorId returns the CreditorId field if non-nil, zero value otherwise.

### GetCreditorIdOk

`func (o *SingleDirectDebitData) GetCreditorIdOk() (*string, bool)`

GetCreditorIdOk returns a tuple with the CreditorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditorId

`func (o *SingleDirectDebitData) SetCreditorId(v string)`

SetCreditorId sets CreditorId field to given value.

### HasCreditorId

`func (o *SingleDirectDebitData) HasCreditorId() bool`

HasCreditorId returns a boolean if a field has been set.

### GetEndToEndId

`func (o *SingleDirectDebitData) GetEndToEndId() string`

GetEndToEndId returns the EndToEndId field if non-nil, zero value otherwise.

### GetEndToEndIdOk

`func (o *SingleDirectDebitData) GetEndToEndIdOk() (*string, bool)`

GetEndToEndIdOk returns a tuple with the EndToEndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndId

`func (o *SingleDirectDebitData) SetEndToEndId(v string)`

SetEndToEndId sets EndToEndId field to given value.

### HasEndToEndId

`func (o *SingleDirectDebitData) HasEndToEndId() bool`

HasEndToEndId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


