# PaymentExecutionResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessMessage** | **string** | Technical message from the bank server, confirming the success of the request. Typically, you would not want to present this message to the user. Note that this field may not be set. However if it is not set, it does not necessarily mean that there was an error in processing the request. | [optional] [default to null]
**WarnMessage** | **string** | In some cases, a bank server may accept the requested order, but return a warn message. This message may be of technical nature, but could also be of interest to the user. | [optional] [default to null]
**PaymentId** | **int64** | Payment identifier. Can be used to retrieve the status of the payment (see &#39;Get payments&#39; service). | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


