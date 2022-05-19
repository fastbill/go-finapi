# PageablePaymentResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Payments** | [**[]Payment**](Payment.md) | &lt;strong&gt;Type:&lt;/strong&gt; Payment&lt;br/&gt; List of received account payments | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewPageablePaymentResources

`func NewPageablePaymentResources(payments []Payment, paging Paging, ) *PageablePaymentResources`

NewPageablePaymentResources instantiates a new PageablePaymentResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageablePaymentResourcesWithDefaults

`func NewPageablePaymentResourcesWithDefaults() *PageablePaymentResources`

NewPageablePaymentResourcesWithDefaults instantiates a new PageablePaymentResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayments

`func (o *PageablePaymentResources) GetPayments() []Payment`

GetPayments returns the Payments field if non-nil, zero value otherwise.

### GetPaymentsOk

`func (o *PageablePaymentResources) GetPaymentsOk() (*[]Payment, bool)`

GetPaymentsOk returns a tuple with the Payments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayments

`func (o *PageablePaymentResources) SetPayments(v []Payment)`

SetPayments sets Payments field to given value.


### GetPaging

`func (o *PageablePaymentResources) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageablePaymentResources) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageablePaymentResources) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


