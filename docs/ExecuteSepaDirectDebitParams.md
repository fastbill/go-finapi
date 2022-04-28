# ExecuteSepaDirectDebitParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | Identifier of the bank account that you want to transfer money to | 
**BankingTan** | Pointer to **string** | Banking TAN that the user received from the bank for executing the direct debit order. The field is required if you are licensed to perform SEPA direct debits yourself. Otherwise, i.e. when finAPI&#39;s Web Form flow is required, the Web Form will deal with executing the service itself. | [optional] 

## Methods

### NewExecuteSepaDirectDebitParams

`func NewExecuteSepaDirectDebitParams(accountId int64, ) *ExecuteSepaDirectDebitParams`

NewExecuteSepaDirectDebitParams instantiates a new ExecuteSepaDirectDebitParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteSepaDirectDebitParamsWithDefaults

`func NewExecuteSepaDirectDebitParamsWithDefaults() *ExecuteSepaDirectDebitParams`

NewExecuteSepaDirectDebitParamsWithDefaults instantiates a new ExecuteSepaDirectDebitParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ExecuteSepaDirectDebitParams) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ExecuteSepaDirectDebitParams) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ExecuteSepaDirectDebitParams) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetBankingTan

`func (o *ExecuteSepaDirectDebitParams) GetBankingTan() string`

GetBankingTan returns the BankingTan field if non-nil, zero value otherwise.

### GetBankingTanOk

`func (o *ExecuteSepaDirectDebitParams) GetBankingTanOk() (*string, bool)`

GetBankingTanOk returns a tuple with the BankingTan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingTan

`func (o *ExecuteSepaDirectDebitParams) SetBankingTan(v string)`

SetBankingTan sets BankingTan field to given value.

### HasBankingTan

`func (o *ExecuteSepaDirectDebitParams) HasBankingTan() bool`

HasBankingTan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


