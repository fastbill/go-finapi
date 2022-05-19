# CreateMoneyTransferParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SingleBooking** | Pointer to **bool** | This field is only relevant when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of &#39;false&#39;), or as single bookings (in case of &#39;true&#39;). Note that it is subject to the bank whether it will regard the field. Default value is &#39;false&#39;. | [optional] [default to false]
**AccountId** | Pointer to **int64** | Identifier of the account that should be used for the money transfer. If you want to do a standalone money transfer (finAPI Payment product, i.e. for an account that is not imported in finAPI) leave this field unset and instead use the field &#39;iban&#39;. | [optional] 
**Iban** | Pointer to **string** | IBAN of the account that should be used for the money transfer. Use this field only if you want to do a standalone money transfer (finAPI Payment product, i.e. for an account that is not imported in finAPI) otherwise, use the &#39;accountId&#39; field and leave this field unset. | [optional] 
**ExecutionDate** | Pointer to **string** | Execution date for the money transfer(s), in the format &#39;YYYY-MM-DD&#39;. May not be in the past. For instant payments, it must be the current date. If not specified, most banks will use the current date as the instructed date for execution. | [optional] 
**MoneyTransfers** | [**[]MoneyTransferOrderParams**](MoneyTransferOrderParams.md) | &lt;strong&gt;Type:&lt;/strong&gt; MoneyTransferOrderParams&lt;br/&gt; List of money transfer orders (may contain at most 15000 items). Please note that collective money transfer may not always be supported. | 
**InstantPayment** | Pointer to **bool** | Whether the order should be submitted to the bank as an instant SEPA order. Default value is &#39;false&#39;.&lt;br/&gt;&lt;br/&gt;NOTE:&lt;br/&gt;&amp;bull; Instant payments can only be submitted if you are self-licensed (and not using the finAPI Web Form) OR via our Web Form from the endpoint &lt;a href&#x3D;&#39;?product&#x3D;web_form_2.0#tag--Payment-Initiation-Services&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt;.&lt;br/&gt;&amp;bull; Submitting an instant payment will work only with interfaces that support it, see BankInterface.paymentCapabilities.sepaInstantMoneyTransfer&lt;br/&gt;&amp;bull; Instant payments work only for a single order, not for collective orders.&lt;br/&gt;&amp;bull; The bank may charge a fee for instant payments, depending on the agreement between the user and the bank.&lt;br/&gt;&amp;bull; The payment might get rejected if the source and/or target account doesn&#39;t support instant payments. | [optional] [default to false]

## Methods

### NewCreateMoneyTransferParams

`func NewCreateMoneyTransferParams(moneyTransfers []MoneyTransferOrderParams, ) *CreateMoneyTransferParams`

NewCreateMoneyTransferParams instantiates a new CreateMoneyTransferParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMoneyTransferParamsWithDefaults

`func NewCreateMoneyTransferParamsWithDefaults() *CreateMoneyTransferParams`

NewCreateMoneyTransferParamsWithDefaults instantiates a new CreateMoneyTransferParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleBooking

`func (o *CreateMoneyTransferParams) GetSingleBooking() bool`

GetSingleBooking returns the SingleBooking field if non-nil, zero value otherwise.

### GetSingleBookingOk

`func (o *CreateMoneyTransferParams) GetSingleBookingOk() (*bool, bool)`

GetSingleBookingOk returns a tuple with the SingleBooking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleBooking

`func (o *CreateMoneyTransferParams) SetSingleBooking(v bool)`

SetSingleBooking sets SingleBooking field to given value.

### HasSingleBooking

`func (o *CreateMoneyTransferParams) HasSingleBooking() bool`

HasSingleBooking returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateMoneyTransferParams) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateMoneyTransferParams) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateMoneyTransferParams) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateMoneyTransferParams) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetIban

`func (o *CreateMoneyTransferParams) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *CreateMoneyTransferParams) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *CreateMoneyTransferParams) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *CreateMoneyTransferParams) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetExecutionDate

`func (o *CreateMoneyTransferParams) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *CreateMoneyTransferParams) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *CreateMoneyTransferParams) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.

### HasExecutionDate

`func (o *CreateMoneyTransferParams) HasExecutionDate() bool`

HasExecutionDate returns a boolean if a field has been set.

### GetMoneyTransfers

`func (o *CreateMoneyTransferParams) GetMoneyTransfers() []MoneyTransferOrderParams`

GetMoneyTransfers returns the MoneyTransfers field if non-nil, zero value otherwise.

### GetMoneyTransfersOk

`func (o *CreateMoneyTransferParams) GetMoneyTransfersOk() (*[]MoneyTransferOrderParams, bool)`

GetMoneyTransfersOk returns a tuple with the MoneyTransfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyTransfers

`func (o *CreateMoneyTransferParams) SetMoneyTransfers(v []MoneyTransferOrderParams)`

SetMoneyTransfers sets MoneyTransfers field to given value.


### GetInstantPayment

`func (o *CreateMoneyTransferParams) GetInstantPayment() bool`

GetInstantPayment returns the InstantPayment field if non-nil, zero value otherwise.

### GetInstantPaymentOk

`func (o *CreateMoneyTransferParams) GetInstantPaymentOk() (*bool, bool)`

GetInstantPaymentOk returns a tuple with the InstantPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantPayment

`func (o *CreateMoneyTransferParams) SetInstantPayment(v bool)`

SetInstantPayment sets InstantPayment field to given value.

### HasInstantPayment

`func (o *CreateMoneyTransferParams) HasInstantPayment() bool`

HasInstantPayment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


