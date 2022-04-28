# CreateDirectDebitParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SingleBooking** | Pointer to **bool** | This field is only relevant when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of &#39;false&#39;), or as single bookings (in case of &#39;true&#39;). Note that it is subject to the bank whether it will regard the field. Default value is &#39;false&#39;. | [optional] [default to false]
**AccountId** | **int64** | Identifier of the account that should be used for the direct debit. | 
**DirectDebitType** | [**DirectDebitType**](DirectDebitType.md) | &lt;strong&gt;Type:&lt;/strong&gt; DirectDebitType&lt;br/&gt; Type of the direct debit; either &lt;code&gt;BASIC&lt;/code&gt; or &lt;code&gt;B2B&lt;/code&gt; (Business-To-Business). | 
**SequenceType** | [**DirectDebitSequenceType**](DirectDebitSequenceType.md) | &lt;strong&gt;Type:&lt;/strong&gt; DirectDebitSequenceType&lt;br/&gt; Sequence type of the direct debit. Possible values:&lt;br/&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;OOFF&lt;/code&gt; - means that this is a one-time direct debit order&lt;br/&gt;&amp;bull; &lt;code&gt;FRST&lt;/code&gt; - means that this is the first in a row of multiple direct debit orders&lt;br/&gt;&amp;bull; &lt;code&gt;RCUR&lt;/code&gt; - means that this is one (but not the first or final) within a row of multiple direct debit orders&lt;br/&gt;&amp;bull; &lt;code&gt;FNAL&lt;/code&gt; - means that this is the final in a row of multiple direct debit orders&lt;br/&gt;&lt;br/&gt; | 
**DirectDebits** | [**[]DirectDebitOrderParams**](DirectDebitOrderParams.md) | &lt;strong&gt;Type:&lt;/strong&gt; DirectDebitOrderParams&lt;br/&gt; List of direct debit orders (may contain at most 15000 items). Please note that collective direct debit may not always be supported. | 
**ExecutionDate** | **string** | Execution date for the direct debit(s), in the format &#39;YYYY-MM-DD&#39;. May not be in the past. If not specified, most banks will use the current date as the instructed date for execution. | 

## Methods

### NewCreateDirectDebitParams

`func NewCreateDirectDebitParams(accountId int64, directDebitType DirectDebitType, sequenceType DirectDebitSequenceType, directDebits []DirectDebitOrderParams, executionDate string, ) *CreateDirectDebitParams`

NewCreateDirectDebitParams instantiates a new CreateDirectDebitParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDirectDebitParamsWithDefaults

`func NewCreateDirectDebitParamsWithDefaults() *CreateDirectDebitParams`

NewCreateDirectDebitParamsWithDefaults instantiates a new CreateDirectDebitParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSingleBooking

`func (o *CreateDirectDebitParams) GetSingleBooking() bool`

GetSingleBooking returns the SingleBooking field if non-nil, zero value otherwise.

### GetSingleBookingOk

`func (o *CreateDirectDebitParams) GetSingleBookingOk() (*bool, bool)`

GetSingleBookingOk returns a tuple with the SingleBooking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleBooking

`func (o *CreateDirectDebitParams) SetSingleBooking(v bool)`

SetSingleBooking sets SingleBooking field to given value.

### HasSingleBooking

`func (o *CreateDirectDebitParams) HasSingleBooking() bool`

HasSingleBooking returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateDirectDebitParams) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateDirectDebitParams) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateDirectDebitParams) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetDirectDebitType

`func (o *CreateDirectDebitParams) GetDirectDebitType() DirectDebitType`

GetDirectDebitType returns the DirectDebitType field if non-nil, zero value otherwise.

### GetDirectDebitTypeOk

`func (o *CreateDirectDebitParams) GetDirectDebitTypeOk() (*DirectDebitType, bool)`

GetDirectDebitTypeOk returns a tuple with the DirectDebitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebitType

`func (o *CreateDirectDebitParams) SetDirectDebitType(v DirectDebitType)`

SetDirectDebitType sets DirectDebitType field to given value.


### GetSequenceType

`func (o *CreateDirectDebitParams) GetSequenceType() DirectDebitSequenceType`

GetSequenceType returns the SequenceType field if non-nil, zero value otherwise.

### GetSequenceTypeOk

`func (o *CreateDirectDebitParams) GetSequenceTypeOk() (*DirectDebitSequenceType, bool)`

GetSequenceTypeOk returns a tuple with the SequenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequenceType

`func (o *CreateDirectDebitParams) SetSequenceType(v DirectDebitSequenceType)`

SetSequenceType sets SequenceType field to given value.


### GetDirectDebits

`func (o *CreateDirectDebitParams) GetDirectDebits() []DirectDebitOrderParams`

GetDirectDebits returns the DirectDebits field if non-nil, zero value otherwise.

### GetDirectDebitsOk

`func (o *CreateDirectDebitParams) GetDirectDebitsOk() (*[]DirectDebitOrderParams, bool)`

GetDirectDebitsOk returns a tuple with the DirectDebits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectDebits

`func (o *CreateDirectDebitParams) SetDirectDebits(v []DirectDebitOrderParams)`

SetDirectDebits sets DirectDebits field to given value.


### GetExecutionDate

`func (o *CreateDirectDebitParams) GetExecutionDate() string`

GetExecutionDate returns the ExecutionDate field if non-nil, zero value otherwise.

### GetExecutionDateOk

`func (o *CreateDirectDebitParams) GetExecutionDateOk() (*string, bool)`

GetExecutionDateOk returns a tuple with the ExecutionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionDate

`func (o *CreateDirectDebitParams) SetExecutionDate(v string)`

SetExecutionDate sets ExecutionDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


