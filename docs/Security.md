# Security

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Identifier. Note: Whenever a security account is being updated, its security positions will be internally re-created, meaning that the identifier of a security position might change over time. | [default to null]
**AccountId** | **int64** | Security account identifier | [default to null]
**Name** | **string** | Name | [optional] [default to null]
**Isin** | **string** | ISIN | [optional] [default to null]
**Wkn** | **string** | WKN | [optional] [default to null]
**Quote** | **float32** | Quote | [optional] [default to null]
**QuoteCurrency** | **string** | Currency of quote | [optional] [default to null]
**QuoteType** | **string** | Type of quote. &#39;PERC&#39; if quote is a percentage value, &#39;ACTU&#39; if quote is the actual amount | [optional] [default to null]
**QuoteDate** | **string** | Quote date in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | [optional] [default to null]
**QuantityNominal** | **float32** | Value of quantity or nominal | [optional] [default to null]
**QuantityNominalType** | **string** | Type of quantity or nominal value. &#39;UNIT&#39; if value is a quantity, &#39;FAMT&#39; if value is the nominal amount | [optional] [default to null]
**MarketValue** | **float32** | Market value | [optional] [default to null]
**MarketValueCurrency** | **string** | Currency of market value | [optional] [default to null]
**EntryQuote** | **float32** | Entry quote | [optional] [default to null]
**EntryQuoteCurrency** | **string** | Currency of entry quote | [optional] [default to null]
**ProfitOrLoss** | **float32** | Current profit or loss | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


