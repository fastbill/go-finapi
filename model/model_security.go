/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.79.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Container for a security position's data
type Security struct {
	// Identifier. Note: Whenever a security account is being updated, its security positions will be internally re-created, meaning that the identifier of a security position might change over time.
	Id int64 `json:"id"`
	// Security account identifier
	AccountId int64 `json:"accountId"`
	// Name
	Name string `json:"name,omitempty"`
	// ISIN
	Isin string `json:"isin,omitempty"`
	// WKN
	Wkn string `json:"wkn,omitempty"`
	// Quote
	Quote float32 `json:"quote,omitempty"`
	// Currency of quote
	QuoteCurrency string `json:"quoteCurrency,omitempty"`
	// Type of quote. 'PERC' if quote is a percentage value, 'ACTU' if quote is the actual amount
	QuoteType string `json:"quoteType,omitempty"`
	// Quote date in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	QuoteDate string `json:"quoteDate,omitempty"`
	// Value of quantity or nominal
	QuantityNominal float32 `json:"quantityNominal,omitempty"`
	// Type of quantity or nominal value. 'UNIT' if value is a quantity, 'FAMT' if value is the nominal amount
	QuantityNominalType string `json:"quantityNominalType,omitempty"`
	// Market value
	MarketValue float32 `json:"marketValue,omitempty"`
	// Currency of market value
	MarketValueCurrency string `json:"marketValueCurrency,omitempty"`
	// Entry quote
	EntryQuote float32 `json:"entryQuote,omitempty"`
	// Currency of entry quote
	EntryQuoteCurrency string `json:"entryQuoteCurrency,omitempty"`
	// Current profit or loss
	ProfitOrLoss float32 `json:"profitOrLoss,omitempty"`
}
