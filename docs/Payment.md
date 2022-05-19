# Payment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Payment identifier | 
**AccountId** | **NullableInt64** | Identifier of the account to which this payment relates. This field is only set if it was specified upon creation of the payment. | 
**Iban** | **NullableString** | IBAN of the account to which this payment relates. This field is only set if it was specified upon creation of the payment. | 
**Type** | [**PaymentType**](PaymentType.md) | &lt;strong&gt;Type:&lt;/strong&gt; PaymentType&lt;br/&gt; Payment type | 
**Amount** | **float64** | Total money amount of the payment order(s), as absolute value | 
**OrderCount** | **int32** | Total count of orders included in this payment | 
**Status** | [**OrderInitiationStatus**](OrderInitiationStatus.md) | &lt;strong&gt;Type:&lt;/strong&gt; OrderInitiationStatus&lt;br/&gt; Current payment status:&lt;br/&gt; &amp;bull; OPEN: means that this payment has been created in finAPI, but not yet submitted to the bank.&lt;br/&gt; &amp;bull; PENDING: means that this payment has been requested at the bank,  but has not been confirmed yet.&lt;br/&gt; &amp;bull; SUCCESSFUL: means that this payment has been successfully initiated.&lt;br/&gt; &amp;bull; NOT_SUCCESSFUL: means that this payment could not be initiated successfully.&lt;br/&gt; &amp;bull; DISCARDED: means that this payment was discarded, either because another payment was requested for the same account before this payment was initiated and the bank does not support this, or because the user has aborted the initiation (when using finAPI&#39;s Web Form). | 
**BankMessage** | **NullableString** | The bank&#39;s response to the most recent request for this payment. Possible requests are: Initial submission of the payment, execution request or subsequent status checks. Note that this field may not always (or never) be set. Also, as long as the payment has not reached its final status, this field can always change. | 
**RequestDate** | **NullableString** | Time of when finAPI submitted this payment to the bank, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time) | 
**ExecutionDate** | **NullableString** | Time of when the execution of this payment has completed, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time).&lt;br/&gt;&lt;br/&gt;Note:&lt;br/&gt;&amp;bull; When the execution of a payment has completed, it does not necessarily mean that the payment was successful. Please refer to the payment &#39;status&#39; for its final status.&lt;br/&gt;&amp;bull; The execution date may deviate from the date when the bank will actually book the payment (for example if the &#39;instructedExecutionDate&#39; is in the future). | 
**InstructedExecutionDate** | **NullableString** | The date that was specified as &#39;executionDate&#39; upon creation of the payment, in the format &#39;YYYY-MM-DD&#39;. This field may not be set if no &#39;executionDate&#39; was specified upon payment creation. | 

## Methods

### NewPayment

`func NewPayment(id int64, accountId NullableInt64, iban NullableString, type_ PaymentType, amount float64, orderCount int32, status OrderInitiationStatus, bankMessage NullableString, requestDate NullableString, executionDate NullableString, instructedExecutionDate NullableString, ) *Payment`

NewPayment instantiates a new Payment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentWithDefaults

`func NewPaymentWithDefaults() *Payment`

NewPaymentWithDefaults instantiates a new Payment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Payment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Payment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Payment) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Payment) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Payment) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Payment) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### SetAccountIdNil

`func (o *Payment) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *Payment) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetIban

`func (o *Payment) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *Payment) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *Payment) SetIban(v string)`

SetIban sets Iban field to given value.


### SetIbanNil

`func (o *Payment) SetIbanNil(b bool)`

 SetIbanNil sets the value for Iban to be an explicit nil

### UnsetIban
`func (o *Payment) UnsetIban()`

UnsetIban ensures that no value is present for Iban, not even an explicit nil
### GetType

`func (o *Payment) GetType() PaymentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Payment) GetTypeOk() (*PaymentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Payment) SetType(v PaymentType)`

SetType sets Type field to given value.


### GetAmount

`func (o *Payment) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payment) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payment) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetOrderCount

`func (o *Payment) GetOrderCount() int32`

GetOrderCount returns the OrderCount field if non-nil, zero value otherwise.

### GetOrderCountOk

`func (o *Payment) GetOrderCountOk() (*int32, bool)`

GetOrderCountOk returns a tuple with the OrderCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCount

`func (o *Payment) SetOrderCount(v int32)`

SetOrderCount sets OrderCount field to given value.


### GetStatus

`func (o *Payment) GetStatus() OrderInitiationStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payment) GetStatusOk() (*OrderInitiationStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payment) SetStatus(v OrderInitiationStatus)`

SetStatus sets Status field to given value.


### GetBankMessage

`func (o *Payment) GetBankMessage() string`

GetBankMessage returns the BankMessage field if non-nil, zero value otherwise.

### GetBankMessageOk

`func (o *Payment) GetBankMessageOk() (*string, bool)`

GetBankMessageOk returns a tuple with the BankMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankMessage

`func (o *Payment) SetBankMessage(v string)`

SetBankMessage sets BankMessage field to given value.


### SetBankMessageNil

`func (o *Payment) SetBankMessageNil(b bool)`

 SetBankMessageNil sets the value for BankMessage to be an explicit nil

### UnsetBankMessage
`func (o *Payment) UnsetBankMessage()`

UnsetBankMessage ensures that no value is present for BankMessage, not even an explicit nil
### GetRequestDate

`func (o *Payment) GetRequestDate() string`

GetRequestDate returns the RequestDate field if non-nil, zero value otherwise.

### GetRequestDateOk

`func (o *Payment) GetRequestDateOk() (*string, bool)`

GetRequestDateOk returns a tuple with the RequestDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestDate

`func (o *Payment) SetRequestDate(v string)`

SetRequestDate sets RequestDate field to given value.


### SetRequestDateNil

`func (o *Payment) SetRequestDateNil(b bool)`

 SetRequestDateNil sets the value for RequestDate to be an explicit nil

### UnsetRequestDate
`func (o *Payment) UnsetRequestDate()`

UnsetRequestDate ensures that no value is present for RequestDate, not even an explicit nil
### GetExecutionDate

`func (o *Payment) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *Payment) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *Payment) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.


### SetExecutionDateNil

`func (o *Payment) SetExecutionDateNil(b bool)`

 SetExecutionDateNil sets the value for ExecutionDate to be an explicit nil

### UnsetExecutionDate
`func (o *Payment) UnsetExecutionDate()`

UnsetExecutionDate ensures that no value is present for ExecutionDate, not even an explicit nil
### GetInstructedExecutionDate

`func (o *Payment) GetInstructedExecutionDate() string`

GetInstructedExecutionDate returns the InstructedExecutionDate field if non-nil, zero value otherwise.

### GetInstructedExecutionDateOk

`func (o *Payment) GetInstructedExecutionDateOk() (*string, bool)`

GetInstructedExecutionDateOk returns a tuple with the InstructedExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructedExecutionDate

`func (o *Payment) SetInstructedExecutionDate(v string)`

SetInstructedExecutionDate sets InstructedExecutionDate field to given value.


### SetInstructedExecutionDateNil

`func (o *Payment) SetInstructedExecutionDateNil(b bool)`

 SetInstructedExecutionDateNil sets the value for InstructedExecutionDate to be an explicit nil

### UnsetInstructedExecutionDate
`func (o *Payment) UnsetInstructedExecutionDate()`

UnsetInstructedExecutionDate ensures that no value is present for InstructedExecutionDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


