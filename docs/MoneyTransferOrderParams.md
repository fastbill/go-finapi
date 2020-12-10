# MoneyTransferOrderParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CounterpartName** | **string** | Name of the counterpart. Note: Neither finAPI nor the involved bank servers are guaranteed to validate the counterpart name. Even if the name does not depict the actual registered account holder of the target account, the order might still be successful. | [default to null]
**CounterpartIban** | **string** | IBAN of the counterpart&#39;s account. | [default to null]
**CounterpartBic** | **string** | BIC of the counterpart&#39;s account. Only required when there is no &#39;IBAN_ONLY&#39;-capability in the respective account/interface combination that is to be used when submitting the payment. | [optional] [default to null]
**Amount** | **float64** | The amount of the payment. Must be a positive decimal number with at most two decimal places. | [default to null]
**Purpose** | **string** | The purpose of the transfer transaction | [optional] [default to null]
**SepaPurposeCode** | **string** | SEPA purpose code, according to ISO 20022, external codes set.&lt;br/&gt;Please note that the SEPA purpose code may be ignored by some banks. | [optional] [default to null]
**EndToEndId** | **string** | End-To-End ID for the transfer transaction | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


