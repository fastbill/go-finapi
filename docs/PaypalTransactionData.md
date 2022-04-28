# PaypalTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InvoiceNumber** | **NullableString** | Invoice Number. | 
**Fee** | **NullableFloat64** | Fee value. | 
**Net** | **NullableFloat64** | Net value. | 
**AuctionSite** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;&lt;br/&gt; Auction Site. | 

## Methods

### NewPaypalTransactionData

`func NewPaypalTransactionData(invoiceNumber NullableString, fee NullableFloat64, net NullableFloat64, auctionSite NullableString, ) *PaypalTransactionData`

NewPaypalTransactionData instantiates a new PaypalTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalTransactionDataWithDefaults

`func NewPaypalTransactionDataWithDefaults() *PaypalTransactionData`

NewPaypalTransactionDataWithDefaults instantiates a new PaypalTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoiceNumber

`func (o *PaypalTransactionData) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *PaypalTransactionData) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *PaypalTransactionData) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.


### SetInvoiceNumberNil

`func (o *PaypalTransactionData) SetInvoiceNumberNil(b bool)`

 SetInvoiceNumberNil sets the value for InvoiceNumber to be an explicit nil

### UnsetInvoiceNumber
`func (o *PaypalTransactionData) UnsetInvoiceNumber()`

UnsetInvoiceNumber ensures that no value is present for InvoiceNumber, not even an explicit nil
### GetFee

`func (o *PaypalTransactionData) GetFee() float64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *PaypalTransactionData) GetFeeOk() (*float64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *PaypalTransactionData) SetFee(v float64)`

SetFee sets Fee field to given value.


### SetFeeNil

`func (o *PaypalTransactionData) SetFeeNil(b bool)`

 SetFeeNil sets the value for Fee to be an explicit nil

### UnsetFee
`func (o *PaypalTransactionData) UnsetFee()`

UnsetFee ensures that no value is present for Fee, not even an explicit nil
### GetNet

`func (o *PaypalTransactionData) GetNet() float64`

GetNet returns the Net field if non-nil, zero value otherwise.

### GetNetOk

`func (o *PaypalTransactionData) GetNetOk() (*float64, bool)`

GetNetOk returns a tuple with the Net field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet

`func (o *PaypalTransactionData) SetNet(v float64)`

SetNet sets Net field to given value.


### SetNetNil

`func (o *PaypalTransactionData) SetNetNil(b bool)`

 SetNetNil sets the value for Net to be an explicit nil

### UnsetNet
`func (o *PaypalTransactionData) UnsetNet()`

UnsetNet ensures that no value is present for Net, not even an explicit nil
### GetAuctionSite

`func (o *PaypalTransactionData) GetAuctionSite() string`

GetAuctionSite returns the AuctionSite field if non-nil, zero value otherwise.

### GetAuctionSiteOk

`func (o *PaypalTransactionData) GetAuctionSiteOk() (*string, bool)`

GetAuctionSiteOk returns a tuple with the AuctionSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionSite

`func (o *PaypalTransactionData) SetAuctionSite(v string)`

SetAuctionSite sets AuctionSite field to given value.


### SetAuctionSiteNil

`func (o *PaypalTransactionData) SetAuctionSiteNil(b bool)`

 SetAuctionSiteNil sets the value for AuctionSite to be an explicit nil

### UnsetAuctionSite
`func (o *PaypalTransactionData) UnsetAuctionSite()`

UnsetAuctionSite ensures that no value is present for AuctionSite, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


