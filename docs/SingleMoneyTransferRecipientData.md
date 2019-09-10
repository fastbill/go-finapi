# SingleMoneyTransferRecipientData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecipientName** | **string** | Name of the recipient. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the recipient name. Even if the recipient name does not depict the actual registered account holder of the specified recipient account, the money transfer request might still be successful. This field is optional only when you pass a clearing account as the recipient. Otherwise, this field is required. | [optional] [default to null]
**RecipientIban** | **string** | IBAN of the recipient&#39;s account. This field is optional only when you pass a clearing account as the recipient. Otherwise, this field is required. | [optional] [default to null]
**RecipientBic** | **string** | BIC of the recipient&#39;s account. Note: This field is optional when you pass a clearing account as the recipient or if the bank connection of the account that you want to transfer money from supports the IBAN-Only money transfer. You can find this out via GET /bankConnections/&lt;id&gt;. Also note that when a BIC is given, then this BIC will be used for the money transfer request independent of whether it is required or not (unless you pass a clearing account, in which case this field will always be ignored). | [optional] [default to null]
**ClearingAccountId** | **string** | Identifier of a clearing account. If this field is set, then the fields &#39;recipientName&#39;, &#39;recipientIban&#39; and &#39;recipientBic&#39; will be ignored and the recipient account will be the specified clearing account. | [optional] [default to null]
**Amount** | **float32** | The amount to transfer. Must be a positive decimal number with at most two decimal places (e.g. 99.99) | [default to null]
**Purpose** | **string** | The purpose of the transfer transaction | [optional] [default to null]
**SepaPurposeCode** | **string** | SEPA purpose code, according to ISO 20022, external codes set. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


