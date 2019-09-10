# SingleDirectDebitData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DebitorName** | **string** | Name of the debitor. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the debitor name. Even if the debitor name does not depict the actual registered account holder of the specified debitor account, the direct debit request might still be successful. | [default to null]
**DebitorIban** | **string** | IBAN of the debitor&#39;s account | [default to null]
**DebitorBic** | **string** | BIC of the debitor&#39;s account. Note: This field is optional if - and only if - the bank connection of the account that you want to transfer money to supports the IBAN-Only direct debit. You can find this out via GET /bankConnections/&lt;id&gt;. Also note that when a BIC is given, then this BIC will be used for the direct debit request independent of whether it is required or not. | [optional] [default to null]
**Amount** | **float32** | The amount to transfer. Must be a positive decimal number with at most two decimal places (e.g. 99.99) | [default to null]
**Purpose** | **string** | The purpose of the transfer transaction | [optional] [default to null]
**SepaPurposeCode** | **string** | SEPA purpose code, according to ISO 20022, external codes set. | [optional] [default to null]
**MandateId** | **string** | Mandate ID that this direct debit order is based on. | [default to null]
**MandateDate** | **string** | Date of the mandate that this direct debit order is based on, in the format &#39;YYYY-MM-DD&#39; | [default to null]
**CreditorId** | **string** | Creditor ID of the source account&#39;s holder | [optional] [default to null]
**EndToEndId** | **string** | End-To-End ID for the transfer transaction | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


