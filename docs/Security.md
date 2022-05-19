# Security

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Identifier. Note: Whenever a security account is being updated, its security positions will be internally re-created, meaning that the identifier of a security position might change over time. | 
**AccountId** | **int64** | Security account identifier | 
**Name** | **NullableString** | Name | 
**Isin** | **NullableString** | ISIN | 
**Wkn** | **NullableString** | WKN | 
**Quote** | **NullableFloat64** | Quote | 
**QuoteCurrency** | **NullableString** | Currency of quote | 
**QuoteType** | [**NullableSecurityPositionQuoteType**](SecurityPositionQuoteType.md) | &lt;strong&gt;Type:&lt;/strong&gt; SecurityPositionQuoteType&lt;br/&gt; Type of quote. &#39;PERC&#39; if quote is a percentage value, &#39;ACTU&#39; if quote is the actual amount | 
**QuoteDate** | **NullableString** | Quote date in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**QuantityNominal** | **NullableFloat64** | Value of quantity or nominal | 
**QuantityNominalType** | [**NullableSecurityPositionQuantityNominalType**](SecurityPositionQuantityNominalType.md) | &lt;strong&gt;Type:&lt;/strong&gt; SecurityPositionQuantityNominalType&lt;br/&gt; Type of quantity or nominal value. &#39;UNIT&#39; if value is a quantity, &#39;FAMT&#39; if value is the nominal amount | 
**MarketValue** | **NullableFloat64** | Market value | 
**MarketValueCurrency** | **NullableString** | Currency of market value | 
**EntryQuote** | **NullableFloat64** | Entry quote | 
**EntryQuoteCurrency** | **NullableString** | Currency of entry quote | 
**ProfitOrLoss** | **NullableFloat64** | Current profit or loss | 

## Methods

### NewSecurity

`func NewSecurity(id int64, accountId int64, name NullableString, isin NullableString, wkn NullableString, quote NullableFloat64, quoteCurrency NullableString, quoteType NullableSecurityPositionQuoteType, quoteDate NullableString, quantityNominal NullableFloat64, quantityNominalType NullableSecurityPositionQuantityNominalType, marketValue NullableFloat64, marketValueCurrency NullableString, entryQuote NullableFloat64, entryQuoteCurrency NullableString, profitOrLoss NullableFloat64, ) *Security`

NewSecurity instantiates a new Security object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityWithDefaults

`func NewSecurityWithDefaults() *Security`

NewSecurityWithDefaults instantiates a new Security object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Security) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Security) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Security) SetId(v int64)`

SetId sets Id field to given value.


### GetAccountId

`func (o *Security) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Security) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Security) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *Security) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Security) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Security) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Security) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Security) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetIsin

`func (o *Security) GetIsin() string`

GetIsin returns the Isin field if non-nil, zero value otherwise.

### GetIsinOk

`func (o *Security) GetIsinOk() (*string, bool)`

GetIsinOk returns a tuple with the Isin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsin

`func (o *Security) SetIsin(v string)`

SetIsin sets Isin field to given value.


### SetIsinNil

`func (o *Security) SetIsinNil(b bool)`

 SetIsinNil sets the value for Isin to be an explicit nil

### UnsetIsin
`func (o *Security) UnsetIsin()`

UnsetIsin ensures that no value is present for Isin, not even an explicit nil
### GetWkn

`func (o *Security) GetWkn() string`

GetWkn returns the Wkn field if non-nil, zero value otherwise.

### GetWknOk

`func (o *Security) GetWknOk() (*string, bool)`

GetWknOk returns a tuple with the Wkn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWkn

`func (o *Security) SetWkn(v string)`

SetWkn sets Wkn field to given value.


### SetWknNil

`func (o *Security) SetWknNil(b bool)`

 SetWknNil sets the value for Wkn to be an explicit nil

### UnsetWkn
`func (o *Security) UnsetWkn()`

UnsetWkn ensures that no value is present for Wkn, not even an explicit nil
### GetQuote

`func (o *Security) GetQuote() float64`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *Security) GetQuoteOk() (*float64, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *Security) SetQuote(v float64)`

SetQuote sets Quote field to given value.


### SetQuoteNil

`func (o *Security) SetQuoteNil(b bool)`

 SetQuoteNil sets the value for Quote to be an explicit nil

### UnsetQuote
`func (o *Security) UnsetQuote()`

UnsetQuote ensures that no value is present for Quote, not even an explicit nil
### GetQuoteCurrency

`func (o *Security) GetQuoteCurrency() string`

GetQuoteCurrency returns the QuoteCurrency field if non-nil, zero value otherwise.

### GetQuoteCurrencyOk

`func (o *Security) GetQuoteCurrencyOk() (*string, bool)`

GetQuoteCurrencyOk returns a tuple with the QuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteCurrency

`func (o *Security) SetQuoteCurrency(v string)`

SetQuoteCurrency sets QuoteCurrency field to given value.


### SetQuoteCurrencyNil

`func (o *Security) SetQuoteCurrencyNil(b bool)`

 SetQuoteCurrencyNil sets the value for QuoteCurrency to be an explicit nil

### UnsetQuoteCurrency
`func (o *Security) UnsetQuoteCurrency()`

UnsetQuoteCurrency ensures that no value is present for QuoteCurrency, not even an explicit nil
### GetQuoteType

`func (o *Security) GetQuoteType() SecurityPositionQuoteType`

GetQuoteType returns the QuoteType field if non-nil, zero value otherwise.

### GetQuoteTypeOk

`func (o *Security) GetQuoteTypeOk() (*SecurityPositionQuoteType, bool)`

GetQuoteTypeOk returns a tuple with the QuoteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteType

`func (o *Security) SetQuoteType(v SecurityPositionQuoteType)`

SetQuoteType sets QuoteType field to given value.


### SetQuoteTypeNil

`func (o *Security) SetQuoteTypeNil(b bool)`

 SetQuoteTypeNil sets the value for QuoteType to be an explicit nil

### UnsetQuoteType
`func (o *Security) UnsetQuoteType()`

UnsetQuoteType ensures that no value is present for QuoteType, not even an explicit nil
### GetQuoteDate

`func (o *Security) GetQuoteDate() string`

GetQuoteDate returns the QuoteDate field if non-nil, zero value otherwise.

### GetQuoteDateOk

`func (o *Security) GetQuoteDateOk() (*string, bool)`

GetQuoteDateOk returns a tuple with the QuoteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteDate

`func (o *Security) SetQuoteDate(v string)`

SetQuoteDate sets QuoteDate field to given value.


### SetQuoteDateNil

`func (o *Security) SetQuoteDateNil(b bool)`

 SetQuoteDateNil sets the value for QuoteDate to be an explicit nil

### UnsetQuoteDate
`func (o *Security) UnsetQuoteDate()`

UnsetQuoteDate ensures that no value is present for QuoteDate, not even an explicit nil
### GetQuantityNominal

`func (o *Security) GetQuantityNominal() float64`

GetQuantityNominal returns the QuantityNominal field if non-nil, zero value otherwise.

### GetQuantityNominalOk

`func (o *Security) GetQuantityNominalOk() (*float64, bool)`

GetQuantityNominalOk returns a tuple with the QuantityNominal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityNominal

`func (o *Security) SetQuantityNominal(v float64)`

SetQuantityNominal sets QuantityNominal field to given value.


### SetQuantityNominalNil

`func (o *Security) SetQuantityNominalNil(b bool)`

 SetQuantityNominalNil sets the value for QuantityNominal to be an explicit nil

### UnsetQuantityNominal
`func (o *Security) UnsetQuantityNominal()`

UnsetQuantityNominal ensures that no value is present for QuantityNominal, not even an explicit nil
### GetQuantityNominalType

`func (o *Security) GetQuantityNominalType() SecurityPositionQuantityNominalType`

GetQuantityNominalType returns the QuantityNominalType field if non-nil, zero value otherwise.

### GetQuantityNominalTypeOk

`func (o *Security) GetQuantityNominalTypeOk() (*SecurityPositionQuantityNominalType, bool)`

GetQuantityNominalTypeOk returns a tuple with the QuantityNominalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityNominalType

`func (o *Security) SetQuantityNominalType(v SecurityPositionQuantityNominalType)`

SetQuantityNominalType sets QuantityNominalType field to given value.


### SetQuantityNominalTypeNil

`func (o *Security) SetQuantityNominalTypeNil(b bool)`

 SetQuantityNominalTypeNil sets the value for QuantityNominalType to be an explicit nil

### UnsetQuantityNominalType
`func (o *Security) UnsetQuantityNominalType()`

UnsetQuantityNominalType ensures that no value is present for QuantityNominalType, not even an explicit nil
### GetMarketValue

`func (o *Security) GetMarketValue() float64`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *Security) GetMarketValueOk() (*float64, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *Security) SetMarketValue(v float64)`

SetMarketValue sets MarketValue field to given value.


### SetMarketValueNil

`func (o *Security) SetMarketValueNil(b bool)`

 SetMarketValueNil sets the value for MarketValue to be an explicit nil

### UnsetMarketValue
`func (o *Security) UnsetMarketValue()`

UnsetMarketValue ensures that no value is present for MarketValue, not even an explicit nil
### GetMarketValueCurrency

`func (o *Security) GetMarketValueCurrency() string`

GetMarketValueCurrency returns the MarketValueCurrency field if non-nil, zero value otherwise.

### GetMarketValueCurrencyOk

`func (o *Security) GetMarketValueCurrencyOk() (*string, bool)`

GetMarketValueCurrencyOk returns a tuple with the MarketValueCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValueCurrency

`func (o *Security) SetMarketValueCurrency(v string)`

SetMarketValueCurrency sets MarketValueCurrency field to given value.


### SetMarketValueCurrencyNil

`func (o *Security) SetMarketValueCurrencyNil(b bool)`

 SetMarketValueCurrencyNil sets the value for MarketValueCurrency to be an explicit nil

### UnsetMarketValueCurrency
`func (o *Security) UnsetMarketValueCurrency()`

UnsetMarketValueCurrency ensures that no value is present for MarketValueCurrency, not even an explicit nil
### GetEntryQuote

`func (o *Security) GetEntryQuote() float64`

GetEntryQuote returns the EntryQuote field if non-nil, zero value otherwise.

### GetEntryQuoteOk

`func (o *Security) GetEntryQuoteOk() (*float64, bool)`

GetEntryQuoteOk returns a tuple with the EntryQuote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryQuote

`func (o *Security) SetEntryQuote(v float64)`

SetEntryQuote sets EntryQuote field to given value.


### SetEntryQuoteNil

`func (o *Security) SetEntryQuoteNil(b bool)`

 SetEntryQuoteNil sets the value for EntryQuote to be an explicit nil

### UnsetEntryQuote
`func (o *Security) UnsetEntryQuote()`

UnsetEntryQuote ensures that no value is present for EntryQuote, not even an explicit nil
### GetEntryQuoteCurrency

`func (o *Security) GetEntryQuoteCurrency() string`

GetEntryQuoteCurrency returns the EntryQuoteCurrency field if non-nil, zero value otherwise.

### GetEntryQuoteCurrencyOk

`func (o *Security) GetEntryQuoteCurrencyOk() (*string, bool)`

GetEntryQuoteCurrencyOk returns a tuple with the EntryQuoteCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryQuoteCurrency

`func (o *Security) SetEntryQuoteCurrency(v string)`

SetEntryQuoteCurrency sets EntryQuoteCurrency field to given value.


### SetEntryQuoteCurrencyNil

`func (o *Security) SetEntryQuoteCurrencyNil(b bool)`

 SetEntryQuoteCurrencyNil sets the value for EntryQuoteCurrency to be an explicit nil

### UnsetEntryQuoteCurrency
`func (o *Security) UnsetEntryQuoteCurrency()`

UnsetEntryQuoteCurrency ensures that no value is present for EntryQuoteCurrency, not even an explicit nil
### GetProfitOrLoss

`func (o *Security) GetProfitOrLoss() float64`

GetProfitOrLoss returns the ProfitOrLoss field if non-nil, zero value otherwise.

### GetProfitOrLossOk

`func (o *Security) GetProfitOrLossOk() (*float64, bool)`

GetProfitOrLossOk returns a tuple with the ProfitOrLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfitOrLoss

`func (o *Security) SetProfitOrLoss(v float64)`

SetProfitOrLoss sets ProfitOrLoss field to given value.


### SetProfitOrLossNil

`func (o *Security) SetProfitOrLossNil(b bool)`

 SetProfitOrLossNil sets the value for ProfitOrLoss to be an explicit nil

### UnsetProfitOrLoss
`func (o *Security) UnsetProfitOrLoss()`

UnsetProfitOrLoss ensures that no value is present for ProfitOrLoss, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


