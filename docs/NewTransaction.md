# NewTransaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **float64** | Amount. Required. | 
**Currency** | Pointer to [**Currency**](Currency.md) | &lt;strong&gt;Type:&lt;/strong&gt; Currency&lt;br/&gt; Transaction currency. | [optional] 
**OriginalAmount** | Pointer to **float64** | Original amount | [optional] 
**OriginalCurrency** | Pointer to [**Currency**](Currency.md) | &lt;strong&gt;Type:&lt;/strong&gt; Currency&lt;br/&gt; Currency of the original amount. | [optional] 
**Purpose** | Pointer to **string** | Purpose. Any symbols are allowed. Maximum length is 2000. Optional. Default value: null. | [optional] 
**Counterpart** | Pointer to **string** | Counterpart. Any symbols are allowed. Maximum length is 80. Optional. Default value: null. | [optional] 
**CounterpartIban** | Pointer to **string** | Counterpart IBAN. Optional. Default value: null. | [optional] 
**CounterpartBlz** | Pointer to **string** | Counterpart BLZ. Optional. Default value: null. | [optional] 
**CounterpartBic** | Pointer to **string** | Counterpart BIC. Optional. Default value: null. | [optional] 
**CounterpartAccountNumber** | Pointer to **string** | Counterpart account number. Maximum length is 34. Optional. Default value: null. | [optional] 
**BookingDate** | Pointer to **string** | Booking date in the format &#39;YYYY-MM-DD&#39;.&lt;br/&gt;&lt;br/&gt;If the date lies back more than 10 days from the booking date of the latest transaction that currently exists in the account, then this transaction will be ignored and not imported. If the date depicts a date in the future, then finAPI will deal with it the same way as it does with real transactions during a real update (see fields &#39;bankBookingDate&#39; and &#39;finapiBookingDate&#39; in the Transaction Resource for explanation).&lt;br/&gt;&lt;br/&gt;This field is optional, default value is the current date. | [optional] 
**ValueDate** | Pointer to **string** | Value date in the format &#39;YYYY-MM-DD&#39;. Optional. Default value: Same as the booking date. | [optional] 
**TypeId** | Pointer to **int32** | The transaction type id. It&#39;s usually a number between 1 and 999. You can look up valid transaction in the following document on page 198: &lt;a href&#x3D;&#39;https://www.hbci-zka.de/dokumente/spezifikation_deutsch/fintsv4/FinTS_4.1_Messages_Finanzdatenformate_2014-01-20-FV.pdf&#39; target&#x3D;&#39;_blank&#39;&gt;FinTS Financial Transaction Services&lt;/a&gt;.&lt;br/&gt; For numbers not listed here, the service call might fail. | [optional] 
**CounterpartMandateReference** | Pointer to **string** | The mandate reference of the counterpart. The maximum possible length of this field is 270 characters. | [optional] 
**CounterpartCreditorId** | Pointer to **string** | The creditor ID of the counterpart. Exists only for SEPA direct debit transactions (\&quot;Lastschrift\&quot;). The maximum possible length of this field is 270 characters. | [optional] 
**CounterpartCustomerReference** | Pointer to **string** | The customer reference of the counterpart. The maximum possible length of this field is 270 characters. | [optional] 
**CounterpartDebitorId** | Pointer to **string** | The originator&#39;s identification code. Exists only for SEPA money transfer transactions (\&quot;Überweisung\&quot;). The maximum possible length of this field is 100 characters. | [optional] 
**Type** | Pointer to **string** | Transaction type, according to the bank. If set, this will contain a German term that you can display to the user. Some examples of common values are: \&quot;Lastschrift\&quot;, \&quot;Auslands&amp;uuml;berweisung\&quot;, \&quot;Geb&amp;uuml;hren\&quot;, \&quot;Zinsen\&quot;. The maximum possible length of this field is 270 characters. | [optional] 
**TypeCodeSwift** | Pointer to **string** | SWIFT transaction type code. | [optional] 
**SepaPurposeCode** | Pointer to **string** | SEPA purpose code. | [optional] 

## Methods

### NewNewTransaction

`func NewNewTransaction(amount float64, ) *NewTransaction`

NewNewTransaction instantiates a new NewTransaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewTransactionWithDefaults

`func NewNewTransactionWithDefaults() *NewTransaction`

NewNewTransactionWithDefaults instantiates a new NewTransaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *NewTransaction) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *NewTransaction) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *NewTransaction) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *NewTransaction) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *NewTransaction) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *NewTransaction) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *NewTransaction) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetOriginalAmount

`func (o *NewTransaction) GetOriginalAmount() float64`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *NewTransaction) GetOriginalAmountOk() (*float64, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *NewTransaction) SetOriginalAmount(v float64)`

SetOriginalAmount sets OriginalAmount field to given value.

### HasOriginalAmount

`func (o *NewTransaction) HasOriginalAmount() bool`

HasOriginalAmount returns a boolean if a field has been set.

### GetOriginalCurrency

`func (o *NewTransaction) GetOriginalCurrency() Currency`

GetOriginalCurrency returns the OriginalCurrency field if non-nil, zero value otherwise.

### GetOriginalCurrencyOk

`func (o *NewTransaction) GetOriginalCurrencyOk() (*Currency, bool)`

GetOriginalCurrencyOk returns a tuple with the OriginalCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCurrency

`func (o *NewTransaction) SetOriginalCurrency(v Currency)`

SetOriginalCurrency sets OriginalCurrency field to given value.

### HasOriginalCurrency

`func (o *NewTransaction) HasOriginalCurrency() bool`

HasOriginalCurrency returns a boolean if a field has been set.

### GetPurpose

`func (o *NewTransaction) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *NewTransaction) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *NewTransaction) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *NewTransaction) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetCounterpart

`func (o *NewTransaction) GetCounterpart() string`

GetCounterpart returns the Counterpart field if non-nil, zero value otherwise.

### GetCounterpartOk

`func (o *NewTransaction) GetCounterpartOk() (*string, bool)`

GetCounterpartOk returns a tuple with the Counterpart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpart

`func (o *NewTransaction) SetCounterpart(v string)`

SetCounterpart sets Counterpart field to given value.

### HasCounterpart

`func (o *NewTransaction) HasCounterpart() bool`

HasCounterpart returns a boolean if a field has been set.

### GetCounterpartIban

`func (o *NewTransaction) GetCounterpartIban() string`

GetCounterpartIban returns the CounterpartIban field if non-nil, zero value otherwise.

### GetCounterpartIbanOk

`func (o *NewTransaction) GetCounterpartIbanOk() (*string, bool)`

GetCounterpartIbanOk returns a tuple with the CounterpartIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartIban

`func (o *NewTransaction) SetCounterpartIban(v string)`

SetCounterpartIban sets CounterpartIban field to given value.

### HasCounterpartIban

`func (o *NewTransaction) HasCounterpartIban() bool`

HasCounterpartIban returns a boolean if a field has been set.

### GetCounterpartBlz

`func (o *NewTransaction) GetCounterpartBlz() string`

GetCounterpartBlz returns the CounterpartBlz field if non-nil, zero value otherwise.

### GetCounterpartBlzOk

`func (o *NewTransaction) GetCounterpartBlzOk() (*string, bool)`

GetCounterpartBlzOk returns a tuple with the CounterpartBlz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBlz

`func (o *NewTransaction) SetCounterpartBlz(v string)`

SetCounterpartBlz sets CounterpartBlz field to given value.

### HasCounterpartBlz

`func (o *NewTransaction) HasCounterpartBlz() bool`

HasCounterpartBlz returns a boolean if a field has been set.

### GetCounterpartBic

`func (o *NewTransaction) GetCounterpartBic() string`

GetCounterpartBic returns the CounterpartBic field if non-nil, zero value otherwise.

### GetCounterpartBicOk

`func (o *NewTransaction) GetCounterpartBicOk() (*string, bool)`

GetCounterpartBicOk returns a tuple with the CounterpartBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBic

`func (o *NewTransaction) SetCounterpartBic(v string)`

SetCounterpartBic sets CounterpartBic field to given value.

### HasCounterpartBic

`func (o *NewTransaction) HasCounterpartBic() bool`

HasCounterpartBic returns a boolean if a field has been set.

### GetCounterpartAccountNumber

`func (o *NewTransaction) GetCounterpartAccountNumber() string`

GetCounterpartAccountNumber returns the CounterpartAccountNumber field if non-nil, zero value otherwise.

### GetCounterpartAccountNumberOk

`func (o *NewTransaction) GetCounterpartAccountNumberOk() (*string, bool)`

GetCounterpartAccountNumberOk returns a tuple with the CounterpartAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartAccountNumber

`func (o *NewTransaction) SetCounterpartAccountNumber(v string)`

SetCounterpartAccountNumber sets CounterpartAccountNumber field to given value.

### HasCounterpartAccountNumber

`func (o *NewTransaction) HasCounterpartAccountNumber() bool`

HasCounterpartAccountNumber returns a boolean if a field has been set.

### GetBookingDate

`func (o *NewTransaction) GetBookingDate() string`

GetBookingDate returns the BookingDate field if non-nil, zero value otherwise.

### GetBookingDateOk

`func (o *NewTransaction) GetBookingDateOk() (*string, bool)`

GetBookingDateOk returns a tuple with the BookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookingDate

`func (o *NewTransaction) SetBookingDate(v string)`

SetBookingDate sets BookingDate field to given value.

### HasBookingDate

`func (o *NewTransaction) HasBookingDate() bool`

HasBookingDate returns a boolean if a field has been set.

### GetValueDate

`func (o *NewTransaction) GetValueDate() string`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *NewTransaction) GetValueDateOk() (*string, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *NewTransaction) SetValueDate(v string)`

SetValueDate sets ValueDate field to given value.

### HasValueDate

`func (o *NewTransaction) HasValueDate() bool`

HasValueDate returns a boolean if a field has been set.

### GetTypeId

`func (o *NewTransaction) GetTypeId() int32`

GetTypeId returns the TypeId field if non-nil, zero value otherwise.

### GetTypeIdOk

`func (o *NewTransaction) GetTypeIdOk() (*int32, bool)`

GetTypeIdOk returns a tuple with the TypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeId

`func (o *NewTransaction) SetTypeId(v int32)`

SetTypeId sets TypeId field to given value.

### HasTypeId

`func (o *NewTransaction) HasTypeId() bool`

HasTypeId returns a boolean if a field has been set.

### GetCounterpartMandateReference

`func (o *NewTransaction) GetCounterpartMandateReference() string`

GetCounterpartMandateReference returns the CounterpartMandateReference field if non-nil, zero value otherwise.

### GetCounterpartMandateReferenceOk

`func (o *NewTransaction) GetCounterpartMandateReferenceOk() (*string, bool)`

GetCounterpartMandateReferenceOk returns a tuple with the CounterpartMandateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartMandateReference

`func (o *NewTransaction) SetCounterpartMandateReference(v string)`

SetCounterpartMandateReference sets CounterpartMandateReference field to given value.

### HasCounterpartMandateReference

`func (o *NewTransaction) HasCounterpartMandateReference() bool`

HasCounterpartMandateReference returns a boolean if a field has been set.

### GetCounterpartCreditorId

`func (o *NewTransaction) GetCounterpartCreditorId() string`

GetCounterpartCreditorId returns the CounterpartCreditorId field if non-nil, zero value otherwise.

### GetCounterpartCreditorIdOk

`func (o *NewTransaction) GetCounterpartCreditorIdOk() (*string, bool)`

GetCounterpartCreditorIdOk returns a tuple with the CounterpartCreditorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartCreditorId

`func (o *NewTransaction) SetCounterpartCreditorId(v string)`

SetCounterpartCreditorId sets CounterpartCreditorId field to given value.

### HasCounterpartCreditorId

`func (o *NewTransaction) HasCounterpartCreditorId() bool`

HasCounterpartCreditorId returns a boolean if a field has been set.

### GetCounterpartCustomerReference

`func (o *NewTransaction) GetCounterpartCustomerReference() string`

GetCounterpartCustomerReference returns the CounterpartCustomerReference field if non-nil, zero value otherwise.

### GetCounterpartCustomerReferenceOk

`func (o *NewTransaction) GetCounterpartCustomerReferenceOk() (*string, bool)`

GetCounterpartCustomerReferenceOk returns a tuple with the CounterpartCustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartCustomerReference

`func (o *NewTransaction) SetCounterpartCustomerReference(v string)`

SetCounterpartCustomerReference sets CounterpartCustomerReference field to given value.

### HasCounterpartCustomerReference

`func (o *NewTransaction) HasCounterpartCustomerReference() bool`

HasCounterpartCustomerReference returns a boolean if a field has been set.

### GetCounterpartDebitorId

`func (o *NewTransaction) GetCounterpartDebitorId() string`

GetCounterpartDebitorId returns the CounterpartDebitorId field if non-nil, zero value otherwise.

### GetCounterpartDebitorIdOk

`func (o *NewTransaction) GetCounterpartDebitorIdOk() (*string, bool)`

GetCounterpartDebitorIdOk returns a tuple with the CounterpartDebitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartDebitorId

`func (o *NewTransaction) SetCounterpartDebitorId(v string)`

SetCounterpartDebitorId sets CounterpartDebitorId field to given value.

### HasCounterpartDebitorId

`func (o *NewTransaction) HasCounterpartDebitorId() bool`

HasCounterpartDebitorId returns a boolean if a field has been set.

### GetType

`func (o *NewTransaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewTransaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewTransaction) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NewTransaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTypeCodeSwift

`func (o *NewTransaction) GetTypeCodeSwift() string`

GetTypeCodeSwift returns the TypeCodeSwift field if non-nil, zero value otherwise.

### GetTypeCodeSwiftOk

`func (o *NewTransaction) GetTypeCodeSwiftOk() (*string, bool)`

GetTypeCodeSwiftOk returns a tuple with the TypeCodeSwift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCodeSwift

`func (o *NewTransaction) SetTypeCodeSwift(v string)`

SetTypeCodeSwift sets TypeCodeSwift field to given value.

### HasTypeCodeSwift

`func (o *NewTransaction) HasTypeCodeSwift() bool`

HasTypeCodeSwift returns a boolean if a field has been set.

### GetSepaPurposeCode

`func (o *NewTransaction) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *NewTransaction) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *NewTransaction) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.

### HasSepaPurposeCode

`func (o *NewTransaction) HasSepaPurposeCode() bool`

HasSepaPurposeCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


