# BankInterfacePaymentCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SepaMoneyTransfer** | **bool** | Capability to do a (single) SEPA money transfer. | 
**SepaInstantMoneyTransfer** | **bool** | Capability to do a (single) SEPA instant money transfer. | 
**SepaCollectiveMoneyTransfer** | **bool** | Capability to do a collective SEPA money transfer. | 
**SepaFutureDatedMoneyTransfer** | **bool** | Capability to do a SEPA money transfer with a future execution date. | 
**SepaStandingOrder** | **bool** | Capability to do a SEPA standing order. | 

## Methods

### NewBankInterfacePaymentCapabilities

`func NewBankInterfacePaymentCapabilities(sepaMoneyTransfer bool, sepaInstantMoneyTransfer bool, sepaCollectiveMoneyTransfer bool, sepaFutureDatedMoneyTransfer bool, sepaStandingOrder bool, ) *BankInterfacePaymentCapabilities`

NewBankInterfacePaymentCapabilities instantiates a new BankInterfacePaymentCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBankInterfacePaymentCapabilitiesWithDefaults

`func NewBankInterfacePaymentCapabilitiesWithDefaults() *BankInterfacePaymentCapabilities`

NewBankInterfacePaymentCapabilitiesWithDefaults instantiates a new BankInterfacePaymentCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSepaMoneyTransfer

`func (o *BankInterfacePaymentCapabilities) GetSepaMoneyTransfer() bool`

GetSepaMoneyTransfer returns the SepaMoneyTransfer field if non-nil, zero value otherwise.

### GetSepaMoneyTransferOk

`func (o *BankInterfacePaymentCapabilities) GetSepaMoneyTransferOk() (*bool, bool)`

GetSepaMoneyTransferOk returns a tuple with the SepaMoneyTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaMoneyTransfer

`func (o *BankInterfacePaymentCapabilities) SetSepaMoneyTransfer(v bool)`

SetSepaMoneyTransfer sets SepaMoneyTransfer field to given value.


### GetSepaInstantMoneyTransfer

`func (o *BankInterfacePaymentCapabilities) GetSepaInstantMoneyTransfer() bool`

GetSepaInstantMoneyTransfer returns the SepaInstantMoneyTransfer field if non-nil, zero value otherwise.

### GetSepaInstantMoneyTransferOk

`func (o *BankInterfacePaymentCapabilities) GetSepaInstantMoneyTransferOk() (*bool, bool)`

GetSepaInstantMoneyTransferOk returns a tuple with the SepaInstantMoneyTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaInstantMoneyTransfer

`func (o *BankInterfacePaymentCapabilities) SetSepaInstantMoneyTransfer(v bool)`

SetSepaInstantMoneyTransfer sets SepaInstantMoneyTransfer field to given value.


### GetSepaCollectiveMoneyTransfer

`func (o *BankInterfacePaymentCapabilities) GetSepaCollectiveMoneyTransfer() bool`

GetSepaCollectiveMoneyTransfer returns the SepaCollectiveMoneyTransfer field if non-nil, zero value otherwise.

### GetSepaCollectiveMoneyTransferOk

`func (o *BankInterfacePaymentCapabilities) GetSepaCollectiveMoneyTransferOk() (*bool, bool)`

GetSepaCollectiveMoneyTransferOk returns a tuple with the SepaCollectiveMoneyTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaCollectiveMoneyTransfer

`func (o *BankInterfacePaymentCapabilities) SetSepaCollectiveMoneyTransfer(v bool)`

SetSepaCollectiveMoneyTransfer sets SepaCollectiveMoneyTransfer field to given value.


### GetSepaFutureDatedMoneyTransfer

`func (o *BankInterfacePaymentCapabilities) GetSepaFutureDatedMoneyTransfer() bool`

GetSepaFutureDatedMoneyTransfer returns the SepaFutureDatedMoneyTransfer field if non-nil, zero value otherwise.

### GetSepaFutureDatedMoneyTransferOk

`func (o *BankInterfacePaymentCapabilities) GetSepaFutureDatedMoneyTransferOk() (*bool, bool)`

GetSepaFutureDatedMoneyTransferOk returns a tuple with the SepaFutureDatedMoneyTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaFutureDatedMoneyTransfer

`func (o *BankInterfacePaymentCapabilities) SetSepaFutureDatedMoneyTransfer(v bool)`

SetSepaFutureDatedMoneyTransfer sets SepaFutureDatedMoneyTransfer field to given value.


### GetSepaStandingOrder

`func (o *BankInterfacePaymentCapabilities) GetSepaStandingOrder() bool`

GetSepaStandingOrder returns the SepaStandingOrder field if non-nil, zero value otherwise.

### GetSepaStandingOrderOk

`func (o *BankInterfacePaymentCapabilities) GetSepaStandingOrderOk() (*bool, bool)`

GetSepaStandingOrderOk returns a tuple with the SepaStandingOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaStandingOrder

`func (o *BankInterfacePaymentCapabilities) SetSepaStandingOrder(v bool)`

SetSepaStandingOrder sets SepaStandingOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


