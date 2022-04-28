# CheckCategorizationTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionId** | **string** | Identifier of transaction. This can be any arbitrary string that will be passed back in the response so that you can map the results to the given transactions. Note that the identifier must be unique within the given list of transactions. | 
**AccountTypeId** | **int64** | Identifier of account type.&lt;br/&gt;&lt;br/&gt;1 &#x3D; Checking,&lt;br/&gt;2 &#x3D; Savings,&lt;br/&gt;3 &#x3D; CreditCard,&lt;br/&gt;4 &#x3D; Security,&lt;br/&gt;5 &#x3D; Loan,&lt;br/&gt;6 &#x3D; Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),&lt;br/&gt;7 &#x3D; Membership,&lt;br/&gt;8 &#x3D; Bausparen&lt;br/&gt;&lt;br/&gt; | 
**Amount** | **float64** | Amount | 
**Purpose** | Pointer to **string** | Purpose. Any symbols are allowed. Maximum length is 2000. Default value: null. | [optional] 
**Counterpart** | Pointer to **string** | Counterpart. Any symbols are allowed. Maximum length is 80. Default value: null. | [optional] 
**CounterpartIban** | Pointer to **string** | Counterpart IBAN. Default value: null. | [optional] 
**CounterpartAccountNumber** | Pointer to **string** | Counterpart account number. Default value: null. | [optional] 
**CounterpartBlz** | Pointer to **string** | Counterpart BLZ. Default value: null. | [optional] 
**CounterpartBic** | Pointer to **string** | Counterpart BIC. Default value: null. | [optional] 
**McCode** | Pointer to **string** | Merchant category code (for credit card transactions only). May only contain up to 4 digits. Default value: null. | [optional] 
**TypeCodeZka** | Pointer to **string** | ZKA business transaction code which relates to the transaction&#39;s type (Number from 0 through 999). Default value: null. | [optional] 

## Methods

### NewCheckCategorizationTransactionData

`func NewCheckCategorizationTransactionData(transactionId string, accountTypeId int64, amount float64, ) *CheckCategorizationTransactionData`

NewCheckCategorizationTransactionData instantiates a new CheckCategorizationTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckCategorizationTransactionDataWithDefaults

`func NewCheckCategorizationTransactionDataWithDefaults() *CheckCategorizationTransactionData`

NewCheckCategorizationTransactionDataWithDefaults instantiates a new CheckCategorizationTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionId

`func (o *CheckCategorizationTransactionData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *CheckCategorizationTransactionData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *CheckCategorizationTransactionData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.


### GetAccountTypeId

`func (o *CheckCategorizationTransactionData) GetAccountTypeId() int64`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *CheckCategorizationTransactionData) GetAccountTypeIdOk() (*int64, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *CheckCategorizationTransactionData) SetAccountTypeId(v int64)`

SetAccountTypeId sets AccountTypeId field to given value.


### GetAmount

`func (o *CheckCategorizationTransactionData) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CheckCategorizationTransactionData) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CheckCategorizationTransactionData) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetPurpose

`func (o *CheckCategorizationTransactionData) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CheckCategorizationTransactionData) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CheckCategorizationTransactionData) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *CheckCategorizationTransactionData) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetCounterpart

`func (o *CheckCategorizationTransactionData) GetCounterpart() string`

GetCounterpart returns the Counterpart field if non-nil, zero value otherwise.

### GetCounterpartOk

`func (o *CheckCategorizationTransactionData) GetCounterpartOk() (*string, bool)`

GetCounterpartOk returns a tuple with the Counterpart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpart

`func (o *CheckCategorizationTransactionData) SetCounterpart(v string)`

SetCounterpart sets Counterpart field to given value.

### HasCounterpart

`func (o *CheckCategorizationTransactionData) HasCounterpart() bool`

HasCounterpart returns a boolean if a field has been set.

### GetCounterpartIban

`func (o *CheckCategorizationTransactionData) GetCounterpartIban() string`

GetCounterpartIban returns the CounterpartIban field if non-nil, zero value otherwise.

### GetCounterpartIbanOk

`func (o *CheckCategorizationTransactionData) GetCounterpartIbanOk() (*string, bool)`

GetCounterpartIbanOk returns a tuple with the CounterpartIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartIban

`func (o *CheckCategorizationTransactionData) SetCounterpartIban(v string)`

SetCounterpartIban sets CounterpartIban field to given value.

### HasCounterpartIban

`func (o *CheckCategorizationTransactionData) HasCounterpartIban() bool`

HasCounterpartIban returns a boolean if a field has been set.

### GetCounterpartAccountNumber

`func (o *CheckCategorizationTransactionData) GetCounterpartAccountNumber() string`

GetCounterpartAccountNumber returns the CounterpartAccountNumber field if non-nil, zero value otherwise.

### GetCounterpartAccountNumberOk

`func (o *CheckCategorizationTransactionData) GetCounterpartAccountNumberOk() (*string, bool)`

GetCounterpartAccountNumberOk returns a tuple with the CounterpartAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartAccountNumber

`func (o *CheckCategorizationTransactionData) SetCounterpartAccountNumber(v string)`

SetCounterpartAccountNumber sets CounterpartAccountNumber field to given value.

### HasCounterpartAccountNumber

`func (o *CheckCategorizationTransactionData) HasCounterpartAccountNumber() bool`

HasCounterpartAccountNumber returns a boolean if a field has been set.

### GetCounterpartBlz

`func (o *CheckCategorizationTransactionData) GetCounterpartBlz() string`

GetCounterpartBlz returns the CounterpartBlz field if non-nil, zero value otherwise.

### GetCounterpartBlzOk

`func (o *CheckCategorizationTransactionData) GetCounterpartBlzOk() (*string, bool)`

GetCounterpartBlzOk returns a tuple with the CounterpartBlz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBlz

`func (o *CheckCategorizationTransactionData) SetCounterpartBlz(v string)`

SetCounterpartBlz sets CounterpartBlz field to given value.

### HasCounterpartBlz

`func (o *CheckCategorizationTransactionData) HasCounterpartBlz() bool`

HasCounterpartBlz returns a boolean if a field has been set.

### GetCounterpartBic

`func (o *CheckCategorizationTransactionData) GetCounterpartBic() string`

GetCounterpartBic returns the CounterpartBic field if non-nil, zero value otherwise.

### GetCounterpartBicOk

`func (o *CheckCategorizationTransactionData) GetCounterpartBicOk() (*string, bool)`

GetCounterpartBicOk returns a tuple with the CounterpartBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBic

`func (o *CheckCategorizationTransactionData) SetCounterpartBic(v string)`

SetCounterpartBic sets CounterpartBic field to given value.

### HasCounterpartBic

`func (o *CheckCategorizationTransactionData) HasCounterpartBic() bool`

HasCounterpartBic returns a boolean if a field has been set.

### GetMcCode

`func (o *CheckCategorizationTransactionData) GetMcCode() string`

GetMcCode returns the McCode field if non-nil, zero value otherwise.

### GetMcCodeOk

`func (o *CheckCategorizationTransactionData) GetMcCodeOk() (*string, bool)`

GetMcCodeOk returns a tuple with the McCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcCode

`func (o *CheckCategorizationTransactionData) SetMcCode(v string)`

SetMcCode sets McCode field to given value.

### HasMcCode

`func (o *CheckCategorizationTransactionData) HasMcCode() bool`

HasMcCode returns a boolean if a field has been set.

### GetTypeCodeZka

`func (o *CheckCategorizationTransactionData) GetTypeCodeZka() string`

GetTypeCodeZka returns the TypeCodeZka field if non-nil, zero value otherwise.

### GetTypeCodeZkaOk

`func (o *CheckCategorizationTransactionData) GetTypeCodeZkaOk() (*string, bool)`

GetTypeCodeZkaOk returns a tuple with the TypeCodeZka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCodeZka

`func (o *CheckCategorizationTransactionData) SetTypeCodeZka(v string)`

SetTypeCodeZka sets TypeCodeZka field to given value.

### HasTypeCodeZka

`func (o *CheckCategorizationTransactionData) HasTypeCodeZka() bool`

HasTypeCodeZka returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


