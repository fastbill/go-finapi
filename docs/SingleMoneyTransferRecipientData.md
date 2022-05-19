# SingleMoneyTransferRecipientData

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

## Methods

### NewSingleMoneyTransferRecipientData

`func NewSingleMoneyTransferRecipientData(amount float64, ) *SingleMoneyTransferRecipientData`

NewSingleMoneyTransferRecipientData instantiates a new SingleMoneyTransferRecipientData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleMoneyTransferRecipientDataWithDefaults

`func NewSingleMoneyTransferRecipientDataWithDefaults() *SingleMoneyTransferRecipientData`

NewSingleMoneyTransferRecipientDataWithDefaults instantiates a new SingleMoneyTransferRecipientData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipientName

`func (o *SingleMoneyTransferRecipientData) GetRecipientName() string`

GetRecipientName returns the RecipientName field if non-nil, zero value otherwise.

### GetRecipientNameOk

`func (o *SingleMoneyTransferRecipientData) GetRecipientNameOk() (*string, bool)`

GetRecipientNameOk returns a tuple with the RecipientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientName

`func (o *SingleMoneyTransferRecipientData) SetRecipientName(v string)`

SetRecipientName sets RecipientName field to given value.

### HasRecipientName

`func (o *SingleMoneyTransferRecipientData) HasRecipientName() bool`

HasRecipientName returns a boolean if a field has been set.

### GetRecipientIban

`func (o *SingleMoneyTransferRecipientData) GetRecipientIban() string`

GetRecipientIban returns the RecipientIban field if non-nil, zero value otherwise.

### GetRecipientIbanOk

`func (o *SingleMoneyTransferRecipientData) GetRecipientIbanOk() (*string, bool)`

GetRecipientIbanOk returns a tuple with the RecipientIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientIban

`func (o *SingleMoneyTransferRecipientData) SetRecipientIban(v string)`

SetRecipientIban sets RecipientIban field to given value.

### HasRecipientIban

`func (o *SingleMoneyTransferRecipientData) HasRecipientIban() bool`

HasRecipientIban returns a boolean if a field has been set.

### GetRecipientBic

`func (o *SingleMoneyTransferRecipientData) GetRecipientBic() string`

GetRecipientBic returns the RecipientBic field if non-nil, zero value otherwise.

### GetRecipientBicOk

`func (o *SingleMoneyTransferRecipientData) GetRecipientBicOk() (*string, bool)`

GetRecipientBicOk returns a tuple with the RecipientBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientBic

`func (o *SingleMoneyTransferRecipientData) SetRecipientBic(v string)`

SetRecipientBic sets RecipientBic field to given value.

### HasRecipientBic

`func (o *SingleMoneyTransferRecipientData) HasRecipientBic() bool`

HasRecipientBic returns a boolean if a field has been set.

### GetClearingAccountId

`func (o *SingleMoneyTransferRecipientData) GetClearingAccountId() string`

GetClearingAccountId returns the ClearingAccountId field if non-nil, zero value otherwise.

### GetClearingAccountIdOk

`func (o *SingleMoneyTransferRecipientData) GetClearingAccountIdOk() (*string, bool)`

GetClearingAccountIdOk returns a tuple with the ClearingAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearingAccountId

`func (o *SingleMoneyTransferRecipientData) SetClearingAccountId(v string)`

SetClearingAccountId sets ClearingAccountId field to given value.

### HasClearingAccountId

`func (o *SingleMoneyTransferRecipientData) HasClearingAccountId() bool`

HasClearingAccountId returns a boolean if a field has been set.

### GetEndToEndId

`func (o *SingleMoneyTransferRecipientData) GetEndToEndId() string`

GetEndToEndId returns the EndToEndId field if non-nil, zero value otherwise.

### GetEndToEndIdOk

`func (o *SingleMoneyTransferRecipientData) GetEndToEndIdOk() (*string, bool)`

GetEndToEndIdOk returns a tuple with the EndToEndId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndId

`func (o *SingleMoneyTransferRecipientData) SetEndToEndId(v string)`

SetEndToEndId sets EndToEndId field to given value.

### HasEndToEndId

`func (o *SingleMoneyTransferRecipientData) HasEndToEndId() bool`

HasEndToEndId returns a boolean if a field has been set.

### GetAmount

`func (o *SingleMoneyTransferRecipientData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SingleMoneyTransferRecipientData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SingleMoneyTransferRecipientData) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *SingleMoneyTransferRecipientData) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *SingleMoneyTransferRecipientData) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *SingleMoneyTransferRecipientData) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *SingleMoneyTransferRecipientData) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSepaPurposeCode

`func (o *SingleMoneyTransferRecipientData) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *SingleMoneyTransferRecipientData) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *SingleMoneyTransferRecipientData) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *SingleMoneyTransferRecipientData) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


