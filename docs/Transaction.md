# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Transaction identifier | 
**ParentId** | **NullableInt64** | Parent transaction identifier | 
**AccountId** | **int64** | Account identifier | 
**ValueDate** | **string** | Value date in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**BankBookingDate** | **string** | Bank booking date in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**FinapiBookingDate** | **string** | finAPI Booking date in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). NOTE: In some cases, banks may deliver transactions that are booked in future, but already included in the current account balance. To keep the account balance consistent with the set of transactions, such \&quot;future transactions\&quot; will be imported with their finapiBookingDate set to the current date (i.e.: date of import). The finapiBookingDate will automatically get adjusted towards the bankBookingDate each time the associated bank account is updated. Example: A transaction is imported on July, 3rd, with a bank reported booking date of July, 6th. The transaction will be imported with its finapiBookingDate set to July, 3rd. Then, on July 4th, the associated account is updated. During this update, the transaction&#39;s finapiBookingDate will be automatically adjusted to July 4th. This adjustment of the finapiBookingDate takes place on each update until the bank account is updated on July 6th or later, in which case the transaction&#39;s finapiBookingDate will be adjusted to its final value, July 6th.&lt;br/&gt; The finapiBookingDate is the date that is used by the finAPI PFM services. E.g. when you calculate the spendings of an account for the current month, and have a transaction with finapiBookingDate in the current month but bankBookingDate at the beginning of the next month, then this transaction is included in the calculations (as the bank has this transaction&#39;s amount included in the current account balance as well). | 
**Amount** | **float64** | Transaction amount | 
**Currency** | [**NullableCurrency**](Currency.md) | &lt;strong&gt;Type:&lt;/strong&gt; Currency&lt;br/&gt; Transaction currency in ISO 4217 format.This field can be null if not explicitly provided the bank. In this case it can be assumed as account’s currency. | 
**Purpose** | **NullableString** | Transaction purpose. Maximum length: 2000 | 
**CounterpartName** | **NullableString** | Counterpart name. Maximum length: 80 | 
**CounterpartAccountNumber** | **NullableString** | Counterpart account number | 
**CounterpartIban** | **NullableString** | Counterpart IBAN | 
**CounterpartBlz** | **NullableString** | Counterpart BLZ | 
**CounterpartBic** | **NullableString** | Counterpart BIC | 
**CounterpartBankName** | **NullableString** | Counterpart Bank name | 
**CounterpartMandateReference** | **NullableString** | The mandate reference of the counterpart | 
**CounterpartCustomerReference** | **NullableString** | The customer reference of the counterpart | 
**CounterpartCreditorId** | **NullableString** | The creditor ID of the counterpart. Exists only for SEPA direct debit transactions (\&quot;Lastschrift\&quot;). | 
**CounterpartDebitorId** | **NullableString** | The originator&#39;s identification code. Exists only for SEPA money transfer transactions (\&quot;Überweisung\&quot;). | 
**Type** | **NullableString** | Transaction type, according to the bank. If set, this will contain a German term that you can display to the user. Some examples of common values are: \&quot;Lastschrift\&quot;, \&quot;Auslands&amp;uuml;berweisung\&quot;, \&quot;Geb&amp;uuml;hren\&quot;, \&quot;Zinsen\&quot;. The maximum possible length of this field is 255 characters. | 
**TypeCodeZka** | **NullableString** | ZKA business transaction code which relates to the transaction&#39;s type. Possible values range from 1 through 999. If no information about the ZKA type code is available, then this field will be null. | 
**TypeCodeSwift** | **NullableString** | SWIFT transaction type code. If no information about the SWIFT code is available, then this field will be null. | 
**SepaPurposeCode** | **NullableString** | SEPA purpose code, according to ISO 20022 | 
**BankTransactionCode** | **NullableString** | Bank transaction code, according to ISO 20022 | 
**Primanota** | **NullableString** | Transaction primanota (bank side identification number) | 
**Category** | [**NullableCategory**](Category.md) | &lt;strong&gt;Type:&lt;/strong&gt; Category&lt;br/&gt; Transaction category, if any is assigned. Note: Recently imported transactions that have currently no category assigned might still get categorized by the background categorization process. To check the status of the background categorization, see GET /bankConnections. Manual category assignments to a transaction will remove the transaction from the background categorization process (i.e. the background categorization process will never overwrite a manual category assignment). | 
**Labels** | [**[]Label**](Label.md) | &lt;strong&gt;Type:&lt;/strong&gt; Label&lt;br/&gt; Array of assigned labels | 
**IsPotentialDuplicate** | **bool** | While finAPI uses a well-elaborated algorithm for uniquely identifying transactions, there is still the possibility that during an account update, a transaction that was imported previously may be imported a second time as a new transaction. For example, this can happen if some transaction data changes on the bank server side. However, finAPI also includes an algorithm of identifying such \&quot;potential duplicate\&quot; transactions. If this field is set to true, it means that finAPI detected a similar transaction that might actually be the same. It is recommended to communicate this information to the end user, and give him an option to delete the transaction in case he confirms that it really is a duplicate. | 
**IsAdjustingEntry** | **bool** | Indicating whether this transaction is an adjusting entry (&#39;Zwischensaldo&#39;).&lt;br/&gt;&lt;br/&gt;Adjusting entries do not originate from the bank, but are added by finAPI during an account update when the bank reported account balance does not add up to the set of transactions that finAPI receives for the account. In this case, the adjusting entry will fix the deviation between the balance and the received transactions so that both adds up again.&lt;br/&gt;&lt;br/&gt;Possible causes for such deviations are:&lt;br/&gt;- Inconsistencies in how the bank calculates the balance, for instance when not yet booked transactions are already included in the balance, but not included in the set of transactions&lt;br/&gt;- Gaps in the transaction history that finAPI receives, for instance because the account has not been updated for a while and older transactions are no longer available | 
**IsNew** | **bool** | Indicating whether this transaction is &#39;new&#39; or not. Any newly imported transaction will have this flag initially set to true. How you use this field is up to your interpretation. For example, you might want to set it to false once a user has clicked on/seen the transaction. You can change this flag to &#39;false&#39; with the PATCH method. | 
**ImportDate** | **string** | Date of transaction import, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**Children** | **[]int64** | Sub-transactions identifiers (if this transaction is split) | 
**PaypalData** | [**NullablePaypalTransactionData**](PaypalTransactionData.md) | &lt;strong&gt;Type:&lt;/strong&gt; PaypalTransactionData&lt;br/&gt; Additional, PayPal-specific transaction data. | 
**EndToEndReference** | **NullableString** | End-To-End reference | 
**CompensationAmount** | **NullableFloat64** | Compensation Amount. Sum of reimbursement of out-of-pocket expenses plus processing brokerage in case of a national return / refund debit as well as an optional interest equalisation. Exists predominantly for SEPA direct debit returns. | 
**OriginalAmount** | **NullableFloat64** | Original Amount of the original direct debit. Exists predominantly for SEPA direct debit returns. | 
**OriginalCurrency** | [**NullableCurrency**](Currency.md) | &lt;strong&gt;Type:&lt;/strong&gt; Currency&lt;br/&gt; Currency of the original amount in ISO 4217 format. This field can be null if not explicitly provided the bank. In this case it can be assumed as account’s currency. | 
**DifferentDebitor** | **NullableString** | Payer&#39;s/debtor&#39;s reference party (in the case of a credit transfer) or payee&#39;s/creditor&#39;s reference party (in the case of a direct debit) | 
**DifferentCreditor** | **NullableString** | Payee&#39;s/creditor&#39;s reference party (in the case of a credit transfer) or payer&#39;s/debtor&#39;s reference party (in the case of a direct debit) | 

## Methods

### NewTransaction

`func NewTransaction(id int64, parentId NullableInt64, accountId int64, valueDate string, bankBookingDate string, finapiBookingDate string, amount float64, currency NullableCurrency, purpose NullableString, counterpartName NullableString, counterpartAccountNumber NullableString, counterpartIban NullableString, counterpartBlz NullableString, counterpartBic NullableString, counterpartBankName NullableString, counterpartMandateReference NullableString, counterpartCustomerReference NullableString, counterpartCreditorId NullableString, counterpartDebitorId NullableString, type_ NullableString, typeCodeZka NullableString, typeCodeSwift NullableString, sepaPurposeCode NullableString, bankTransactionCode NullableString, primanota NullableString, category NullableCategory, labels []Label, isPotentialDuplicate bool, isAdjustingEntry bool, isNew bool, importDate string, children []int64, paypalData NullablePaypalTransactionData, endToEndReference NullableString, compensationAmount NullableFloat64, originalAmount NullableFloat64, originalCurrency NullableCurrency, differentDebitor NullableString, differentCreditor NullableString, ) *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v int64)`

SetId sets Id field to given value.


### GetParentId

`func (o *Transaction) GetParentId() int64`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Transaction) GetParentIdOk() (*int64, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Transaction) SetParentId(v int64)`

SetParentId sets ParentId field to given value.


### SetParentIdNil

`func (o *Transaction) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *Transaction) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetAccountId

`func (o *Transaction) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Transaction) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Transaction) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetValueDate

`func (o *Transaction) GetValueDate() string`

GetValueDate returns the ValueDate field if non-nil, zero value otherwise.

### GetValueDateOk

`func (o *Transaction) GetValueDateOk() (*string, bool)`

GetValueDateOk returns a tuple with the ValueDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueDate

`func (o *Transaction) SetValueDate(v string)`

SetValueDate sets ValueDate field to given value.


### GetBankBookingDate

`func (o *Transaction) GetBankBookingDate() string`

GetBankBookingDate returns the BankBookingDate field if non-nil, zero value otherwise.

### GetBankBookingDateOk

`func (o *Transaction) GetBankBookingDateOk() (*string, bool)`

GetBankBookingDateOk returns a tuple with the BankBookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankBookingDate

`func (o *Transaction) SetBankBookingDate(v string)`

SetBankBookingDate sets BankBookingDate field to given value.


### GetFinapiBookingDate

`func (o *Transaction) GetFinapiBookingDate() string`

GetFinapiBookingDate returns the FinapiBookingDate field if non-nil, zero value otherwise.

### GetFinapiBookingDateOk

`func (o *Transaction) GetFinapiBookingDateOk() (*string, bool)`

GetFinapiBookingDateOk returns a tuple with the FinapiBookingDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinapiBookingDate

`func (o *Transaction) SetFinapiBookingDate(v string)`

SetFinapiBookingDate sets FinapiBookingDate field to given value.


### GetAmount

`func (o *Transaction) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v float64)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *Transaction) GetCurrency() Currency`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Transaction) GetCurrencyOk() (*Currency, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Transaction) SetCurrency(v Currency)`

SetCurrency sets Currency field to given value.


### SetCurrencyNil

`func (o *Transaction) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *Transaction) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetPurpose

`func (o *Transaction) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *Transaction) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *Transaction) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### SetPurposeNil

`func (o *Transaction) SetPurposeNil(b bool)`

 SetPurposeNil sets the value for Purpose to be an explicit nil

### UnsetPurpose
`func (o *Transaction) UnsetPurpose()`

UnsetPurpose ensures that no value is present for Purpose, not even an explicit nil
### GetCounterpartName

`func (o *Transaction) GetCounterpartName() string`

GetCounterpartName returns the CounterpartName field if non-nil, zero value otherwise.

### GetCounterpartNameOk

`func (o *Transaction) GetCounterpartNameOk() (*string, bool)`

GetCounterpartNameOk returns a tuple with the CounterpartName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartName

`func (o *Transaction) SetCounterpartName(v string)`

SetCounterpartName sets CounterpartName field to given value.


### SetCounterpartNameNil

`func (o *Transaction) SetCounterpartNameNil(b bool)`

 SetCounterpartNameNil sets the value for CounterpartName to be an explicit nil

### UnsetCounterpartName
`func (o *Transaction) UnsetCounterpartName()`

UnsetCounterpartName ensures that no value is present for CounterpartName, not even an explicit nil
### GetCounterpartAccountNumber

`func (o *Transaction) GetCounterpartAccountNumber() string`

GetCounterpartAccountNumber returns the CounterpartAccountNumber field if non-nil, zero value otherwise.

### GetCounterpartAccountNumberOk

`func (o *Transaction) GetCounterpartAccountNumberOk() (*string, bool)`

GetCounterpartAccountNumberOk returns a tuple with the CounterpartAccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartAccountNumber

`func (o *Transaction) SetCounterpartAccountNumber(v string)`

SetCounterpartAccountNumber sets CounterpartAccountNumber field to given value.


### SetCounterpartAccountNumberNil

`func (o *Transaction) SetCounterpartAccountNumberNil(b bool)`

 SetCounterpartAccountNumberNil sets the value for CounterpartAccountNumber to be an explicit nil

### UnsetCounterpartAccountNumber
`func (o *Transaction) UnsetCounterpartAccountNumber()`

UnsetCounterpartAccountNumber ensures that no value is present for CounterpartAccountNumber, not even an explicit nil
### GetCounterpartIban

`func (o *Transaction) GetCounterpartIban() string`

GetCounterpartIban returns the CounterpartIban field if non-nil, zero value otherwise.

### GetCounterpartIbanOk

`func (o *Transaction) GetCounterpartIbanOk() (*string, bool)`

GetCounterpartIbanOk returns a tuple with the CounterpartIban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartIban

`func (o *Transaction) SetCounterpartIban(v string)`

SetCounterpartIban sets CounterpartIban field to given value.


### SetCounterpartIbanNil

`func (o *Transaction) SetCounterpartIbanNil(b bool)`

 SetCounterpartIbanNil sets the value for CounterpartIban to be an explicit nil

### UnsetCounterpartIban
`func (o *Transaction) UnsetCounterpartIban()`

UnsetCounterpartIban ensures that no value is present for CounterpartIban, not even an explicit nil
### GetCounterpartBlz

`func (o *Transaction) GetCounterpartBlz() string`

GetCounterpartBlz returns the CounterpartBlz field if non-nil, zero value otherwise.

### GetCounterpartBlzOk

`func (o *Transaction) GetCounterpartBlzOk() (*string, bool)`

GetCounterpartBlzOk returns a tuple with the CounterpartBlz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBlz

`func (o *Transaction) SetCounterpartBlz(v string)`

SetCounterpartBlz sets CounterpartBlz field to given value.


### SetCounterpartBlzNil

`func (o *Transaction) SetCounterpartBlzNil(b bool)`

 SetCounterpartBlzNil sets the value for CounterpartBlz to be an explicit nil

### UnsetCounterpartBlz
`func (o *Transaction) UnsetCounterpartBlz()`

UnsetCounterpartBlz ensures that no value is present for CounterpartBlz, not even an explicit nil
### GetCounterpartBic

`func (o *Transaction) GetCounterpartBic() string`

GetCounterpartBic returns the CounterpartBic field if non-nil, zero value otherwise.

### GetCounterpartBicOk

`func (o *Transaction) GetCounterpartBicOk() (*string, bool)`

GetCounterpartBicOk returns a tuple with the CounterpartBic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBic

`func (o *Transaction) SetCounterpartBic(v string)`

SetCounterpartBic sets CounterpartBic field to given value.


### SetCounterpartBicNil

`func (o *Transaction) SetCounterpartBicNil(b bool)`

 SetCounterpartBicNil sets the value for CounterpartBic to be an explicit nil

### UnsetCounterpartBic
`func (o *Transaction) UnsetCounterpartBic()`

UnsetCounterpartBic ensures that no value is present for CounterpartBic, not even an explicit nil
### GetCounterpartBankName

`func (o *Transaction) GetCounterpartBankName() string`

GetCounterpartBankName returns the CounterpartBankName field if non-nil, zero value otherwise.

### GetCounterpartBankNameOk

`func (o *Transaction) GetCounterpartBankNameOk() (*string, bool)`

GetCounterpartBankNameOk returns a tuple with the CounterpartBankName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartBankName

`func (o *Transaction) SetCounterpartBankName(v string)`

SetCounterpartBankName sets CounterpartBankName field to given value.


### SetCounterpartBankNameNil

`func (o *Transaction) SetCounterpartBankNameNil(b bool)`

 SetCounterpartBankNameNil sets the value for CounterpartBankName to be an explicit nil

### UnsetCounterpartBankName
`func (o *Transaction) UnsetCounterpartBankName()`

UnsetCounterpartBankName ensures that no value is present for CounterpartBankName, not even an explicit nil
### GetCounterpartMandateReference

`func (o *Transaction) GetCounterpartMandateReference() string`

GetCounterpartMandateReference returns the CounterpartMandateReference field if non-nil, zero value otherwise.

### GetCounterpartMandateReferenceOk

`func (o *Transaction) GetCounterpartMandateReferenceOk() (*string, bool)`

GetCounterpartMandateReferenceOk returns a tuple with the CounterpartMandateReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartMandateReference

`func (o *Transaction) SetCounterpartMandateReference(v string)`

SetCounterpartMandateReference sets CounterpartMandateReference field to given value.


### SetCounterpartMandateReferenceNil

`func (o *Transaction) SetCounterpartMandateReferenceNil(b bool)`

 SetCounterpartMandateReferenceNil sets the value for CounterpartMandateReference to be an explicit nil

### UnsetCounterpartMandateReference
`func (o *Transaction) UnsetCounterpartMandateReference()`

UnsetCounterpartMandateReference ensures that no value is present for CounterpartMandateReference, not even an explicit nil
### GetCounterpartCustomerReference

`func (o *Transaction) GetCounterpartCustomerReference() string`

GetCounterpartCustomerReference returns the CounterpartCustomerReference field if non-nil, zero value otherwise.

### GetCounterpartCustomerReferenceOk

`func (o *Transaction) GetCounterpartCustomerReferenceOk() (*string, bool)`

GetCounterpartCustomerReferenceOk returns a tuple with the CounterpartCustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartCustomerReference

`func (o *Transaction) SetCounterpartCustomerReference(v string)`

SetCounterpartCustomerReference sets CounterpartCustomerReference field to given value.


### SetCounterpartCustomerReferenceNil

`func (o *Transaction) SetCounterpartCustomerReferenceNil(b bool)`

 SetCounterpartCustomerReferenceNil sets the value for CounterpartCustomerReference to be an explicit nil

### UnsetCounterpartCustomerReference
`func (o *Transaction) UnsetCounterpartCustomerReference()`

UnsetCounterpartCustomerReference ensures that no value is present for CounterpartCustomerReference, not even an explicit nil
### GetCounterpartCreditorId

`func (o *Transaction) GetCounterpartCreditorId() string`

GetCounterpartCreditorId returns the CounterpartCreditorId field if non-nil, zero value otherwise.

### GetCounterpartCreditorIdOk

`func (o *Transaction) GetCounterpartCreditorIdOk() (*string, bool)`

GetCounterpartCreditorIdOk returns a tuple with the CounterpartCreditorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartCreditorId

`func (o *Transaction) SetCounterpartCreditorId(v string)`

SetCounterpartCreditorId sets CounterpartCreditorId field to given value.


### SetCounterpartCreditorIdNil

`func (o *Transaction) SetCounterpartCreditorIdNil(b bool)`

 SetCounterpartCreditorIdNil sets the value for CounterpartCreditorId to be an explicit nil

### UnsetCounterpartCreditorId
`func (o *Transaction) UnsetCounterpartCreditorId()`

UnsetCounterpartCreditorId ensures that no value is present for CounterpartCreditorId, not even an explicit nil
### GetCounterpartDebitorId

`func (o *Transaction) GetCounterpartDebitorId() string`

GetCounterpartDebitorId returns the CounterpartDebitorId field if non-nil, zero value otherwise.

### GetCounterpartDebitorIdOk

`func (o *Transaction) GetCounterpartDebitorIdOk() (*string, bool)`

GetCounterpartDebitorIdOk returns a tuple with the CounterpartDebitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterpartDebitorId

`func (o *Transaction) SetCounterpartDebitorId(v string)`

SetCounterpartDebitorId sets CounterpartDebitorId field to given value.


### SetCounterpartDebitorIdNil

`func (o *Transaction) SetCounterpartDebitorIdNil(b bool)`

 SetCounterpartDebitorIdNil sets the value for CounterpartDebitorId to be an explicit nil

### UnsetCounterpartDebitorId
`func (o *Transaction) UnsetCounterpartDebitorId()`

UnsetCounterpartDebitorId ensures that no value is present for CounterpartDebitorId, not even an explicit nil
### GetType

`func (o *Transaction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v string)`

SetType sets Type field to given value.


### SetTypeNil

`func (o *Transaction) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *Transaction) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetTypeCodeZka

`func (o *Transaction) GetTypeCodeZka() string`

GetTypeCodeZka returns the TypeCodeZka field if non-nil, zero value otherwise.

### GetTypeCodeZkaOk

`func (o *Transaction) GetTypeCodeZkaOk() (*string, bool)`

GetTypeCodeZkaOk returns a tuple with the TypeCodeZka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCodeZka

`func (o *Transaction) SetTypeCodeZka(v string)`

SetTypeCodeZka sets TypeCodeZka field to given value.


### SetTypeCodeZkaNil

`func (o *Transaction) SetTypeCodeZkaNil(b bool)`

 SetTypeCodeZkaNil sets the value for TypeCodeZka to be an explicit nil

### UnsetTypeCodeZka
`func (o *Transaction) UnsetTypeCodeZka()`

UnsetTypeCodeZka ensures that no value is present for TypeCodeZka, not even an explicit nil
### GetTypeCodeSwift

`func (o *Transaction) GetTypeCodeSwift() string`

GetTypeCodeSwift returns the TypeCodeSwift field if non-nil, zero value otherwise.

### GetTypeCodeSwiftOk

`func (o *Transaction) GetTypeCodeSwiftOk() (*string, bool)`

GetTypeCodeSwiftOk returns a tuple with the TypeCodeSwift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCodeSwift

`func (o *Transaction) SetTypeCodeSwift(v string)`

SetTypeCodeSwift sets TypeCodeSwift field to given value.


### SetTypeCodeSwiftNil

`func (o *Transaction) SetTypeCodeSwiftNil(b bool)`

 SetTypeCodeSwiftNil sets the value for TypeCodeSwift to be an explicit nil

### UnsetTypeCodeSwift
`func (o *Transaction) UnsetTypeCodeSwift()`

UnsetTypeCodeSwift ensures that no value is present for TypeCodeSwift, not even an explicit nil
### GetSepaPurposeCode

`func (o *Transaction) GetSepaPurposeCode() string`

GetSepaPurposeCode returns the SepaPurposeCode field if non-nil, zero value otherwise.

### GetSepaPurposeCodeOk

`func (o *Transaction) GetSepaPurposeCodeOk() (*string, bool)`

GetSepaPurposeCodeOk returns a tuple with the SepaPurposeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSepaPurposeCode

`func (o *Transaction) SetSepaPurposeCode(v string)`

SetSepaPurposeCode sets SepaPurposeCode field to given value.


### SetSepaPurposeCodeNil

`func (o *Transaction) SetSepaPurposeCodeNil(b bool)`

 SetSepaPurposeCodeNil sets the value for SepaPurposeCode to be an explicit nil

### UnsetSepaPurposeCode
`func (o *Transaction) UnsetSepaPurposeCode()`

UnsetSepaPurposeCode ensures that no value is present for SepaPurposeCode, not even an explicit nil
### GetBankTransactionCode

`func (o *Transaction) GetBankTransactionCode() string`

GetBankTransactionCode returns the BankTransactionCode field if non-nil, zero value otherwise.

### GetBankTransactionCodeOk

`func (o *Transaction) GetBankTransactionCodeOk() (*string, bool)`

GetBankTransactionCodeOk returns a tuple with the BankTransactionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransactionCode

`func (o *Transaction) SetBankTransactionCode(v string)`

SetBankTransactionCode sets BankTransactionCode field to given value.


### SetBankTransactionCodeNil

`func (o *Transaction) SetBankTransactionCodeNil(b bool)`

 SetBankTransactionCodeNil sets the value for BankTransactionCode to be an explicit nil

### UnsetBankTransactionCode
`func (o *Transaction) UnsetBankTransactionCode()`

UnsetBankTransactionCode ensures that no value is present for BankTransactionCode, not even an explicit nil
### GetPrimanota

`func (o *Transaction) GetPrimanota() string`

GetPrimanota returns the Primanota field if non-nil, zero value otherwise.

### GetPrimanotaOk

`func (o *Transaction) GetPrimanotaOk() (*string, bool)`

GetPrimanotaOk returns a tuple with the Primanota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimanota

`func (o *Transaction) SetPrimanota(v string)`

SetPrimanota sets Primanota field to given value.


### SetPrimanotaNil

`func (o *Transaction) SetPrimanotaNil(b bool)`

 SetPrimanotaNil sets the value for Primanota to be an explicit nil

### UnsetPrimanota
`func (o *Transaction) UnsetPrimanota()`

UnsetPrimanota ensures that no value is present for Primanota, not even an explicit nil
### GetCategory

`func (o *Transaction) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Transaction) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Transaction) SetCategory(v Category)`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *Transaction) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *Transaction) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetLabels

`func (o *Transaction) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Transaction) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Transaction) SetLabels(v []Label)`

SetLabels sets Labels field to given value.


### GetIsPotentialDuplicate

`func (o *Transaction) GetIsPotentialDuplicate() bool`

GetIsPotentialDuplicate returns the IsPotentialDuplicate field if non-nil, zero value otherwise.

### GetIsPotentialDuplicateOk

`func (o *Transaction) GetIsPotentialDuplicateOk() (*bool, bool)`

GetIsPotentialDuplicateOk returns a tuple with the IsPotentialDuplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPotentialDuplicate

`func (o *Transaction) SetIsPotentialDuplicate(v bool)`

SetIsPotentialDuplicate sets IsPotentialDuplicate field to given value.


### GetIsAdjustingEntry

`func (o *Transaction) GetIsAdjustingEntry() bool`

GetIsAdjustingEntry returns the IsAdjustingEntry field if non-nil, zero value otherwise.

### GetIsAdjustingEntryOk

`func (o *Transaction) GetIsAdjustingEntryOk() (*bool, bool)`

GetIsAdjustingEntryOk returns a tuple with the IsAdjustingEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAdjustingEntry

`func (o *Transaction) SetIsAdjustingEntry(v bool)`

SetIsAdjustingEntry sets IsAdjustingEntry field to given value.


### GetIsNew

`func (o *Transaction) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *Transaction) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *Transaction) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.


### GetImportDate

`func (o *Transaction) GetImportDate() string`

GetImportDate returns the ImportDate field if non-nil, zero value otherwise.

### GetImportDateOk

`func (o *Transaction) GetImportDateOk() (*string, bool)`

GetImportDateOk returns a tuple with the ImportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportDate

`func (o *Transaction) SetImportDate(v string)`

SetImportDate sets ImportDate field to given value.


### GetChildren

`func (o *Transaction) GetChildren() []int64`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *Transaction) GetChildrenOk() (*[]int64, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *Transaction) SetChildren(v []int64)`

SetChildren sets Children field to given value.


### SetChildrenNil

`func (o *Transaction) SetChildrenNil(b bool)`

 SetChildrenNil sets the value for Children to be an explicit nil

### UnsetChildren
`func (o *Transaction) UnsetChildren()`

UnsetChildren ensures that no value is present for Children, not even an explicit nil
### GetPaypalData

`func (o *Transaction) GetPaypalData() PaypalTransactionData`

GetPaypalData returns the PaypalData field if non-nil, zero value otherwise.

### GetPaypalDataOk

`func (o *Transaction) GetPaypalDataOk() (*PaypalTransactionData, bool)`

GetPaypalDataOk returns a tuple with the PaypalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalData

`func (o *Transaction) SetPaypalData(v PaypalTransactionData)`

SetPaypalData sets PaypalData field to given value.


### SetPaypalDataNil

`func (o *Transaction) SetPaypalDataNil(b bool)`

 SetPaypalDataNil sets the value for PaypalData to be an explicit nil

### UnsetPaypalData
`func (o *Transaction) UnsetPaypalData()`

UnsetPaypalData ensures that no value is present for PaypalData, not even an explicit nil
### GetEndToEndReference

`func (o *Transaction) GetEndToEndReference() string`

GetEndToEndReference returns the EndToEndReference field if non-nil, zero value otherwise.

### GetEndToEndReferenceOk

`func (o *Transaction) GetEndToEndReferenceOk() (*string, bool)`

GetEndToEndReferenceOk returns a tuple with the EndToEndReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndToEndReference

`func (o *Transaction) SetEndToEndReference(v string)`

SetEndToEndReference sets EndToEndReference field to given value.


### SetEndToEndReferenceNil

`func (o *Transaction) SetEndToEndReferenceNil(b bool)`

 SetEndToEndReferenceNil sets the value for EndToEndReference to be an explicit nil

### UnsetEndToEndReference
`func (o *Transaction) UnsetEndToEndReference()`

UnsetEndToEndReference ensures that no value is present for EndToEndReference, not even an explicit nil
### GetCompensationAmount

`func (o *Transaction) GetCompensationAmount() float64`

GetCompensationAmount returns the CompensationAmount field if non-nil, zero value otherwise.

### GetCompensationAmountOk

`func (o *Transaction) GetCompensationAmountOk() (*float64, bool)`

GetCompensationAmountOk returns a tuple with the CompensationAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompensationAmount

`func (o *Transaction) SetCompensationAmount(v float64)`

SetCompensationAmount sets CompensationAmount field to given value.


### SetCompensationAmountNil

`func (o *Transaction) SetCompensationAmountNil(b bool)`

 SetCompensationAmountNil sets the value for CompensationAmount to be an explicit nil

### UnsetCompensationAmount
`func (o *Transaction) UnsetCompensationAmount()`

UnsetCompensationAmount ensures that no value is present for CompensationAmount, not even an explicit nil
### GetOriginalAmount

`func (o *Transaction) GetOriginalAmount() float64`

GetOriginalAmount returns the OriginalAmount field if non-nil, zero value otherwise.

### GetOriginalAmountOk

`func (o *Transaction) GetOriginalAmountOk() (*float64, bool)`

GetOriginalAmountOk returns a tuple with the OriginalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAmount

`func (o *Transaction) SetOriginalAmount(v float64)`

SetOriginalAmount sets OriginalAmount field to given value.


### SetOriginalAmountNil

`func (o *Transaction) SetOriginalAmountNil(b bool)`

 SetOriginalAmountNil sets the value for OriginalAmount to be an explicit nil

### UnsetOriginalAmount
`func (o *Transaction) UnsetOriginalAmount()`

UnsetOriginalAmount ensures that no value is present for OriginalAmount, not even an explicit nil
### GetOriginalCurrency

`func (o *Transaction) GetOriginalCurrency() Currency`

GetOriginalCurrency returns the OriginalCurrency field if non-nil, zero value otherwise.

### GetOriginalCurrencyOk

`func (o *Transaction) GetOriginalCurrencyOk() (*Currency, bool)`

GetOriginalCurrencyOk returns a tuple with the OriginalCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalCurrency

`func (o *Transaction) SetOriginalCurrency(v Currency)`

SetOriginalCurrency sets OriginalCurrency field to given value.


### SetOriginalCurrencyNil

`func (o *Transaction) SetOriginalCurrencyNil(b bool)`

 SetOriginalCurrencyNil sets the value for OriginalCurrency to be an explicit nil

### UnsetOriginalCurrency
`func (o *Transaction) UnsetOriginalCurrency()`

UnsetOriginalCurrency ensures that no value is present for OriginalCurrency, not even an explicit nil
### GetDifferentDebitor

`func (o *Transaction) GetDifferentDebitor() string`

GetDifferentDebitor returns the DifferentDebitor field if non-nil, zero value otherwise.

### GetDifferentDebitorOk

`func (o *Transaction) GetDifferentDebitorOk() (*string, bool)`

GetDifferentDebitorOk returns a tuple with the DifferentDebitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferentDebitor

`func (o *Transaction) SetDifferentDebitor(v string)`

SetDifferentDebitor sets DifferentDebitor field to given value.


### SetDifferentDebitorNil

`func (o *Transaction) SetDifferentDebitorNil(b bool)`

 SetDifferentDebitorNil sets the value for DifferentDebitor to be an explicit nil

### UnsetDifferentDebitor
`func (o *Transaction) UnsetDifferentDebitor()`

UnsetDifferentDebitor ensures that no value is present for DifferentDebitor, not even an explicit nil
### GetDifferentCreditor

`func (o *Transaction) GetDifferentCreditor() string`

GetDifferentCreditor returns the DifferentCreditor field if non-nil, zero value otherwise.

### GetDifferentCreditorOk

`func (o *Transaction) GetDifferentCreditorOk() (*string, bool)`

GetDifferentCreditorOk returns a tuple with the DifferentCreditor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferentCreditor

`func (o *Transaction) SetDifferentCreditor(v string)`

SetDifferentCreditor sets DifferentCreditor field to given value.


### SetDifferentCreditorNil

`func (o *Transaction) SetDifferentCreditorNil(b bool)`

 SetDifferentCreditorNil sets the value for DifferentCreditor to be an explicit nil

### UnsetDifferentCreditor
`func (o *Transaction) UnsetDifferentCreditor()`

UnsetDifferentCreditor ensures that no value is present for DifferentCreditor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


