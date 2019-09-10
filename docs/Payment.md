# Payment

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Payment identifier | [default to null]
**AccountId** | **int64** | Identifier of the account to which this payment relates | [default to null]
**RequestDate** | **string** | Time of when this payment was requested, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time) | [default to null]
**ExecutionDate** | **string** | Time of when this payment was executed, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time) | [optional] [default to null]
**Type_** | **string** | Payment type | [default to null]
**Status** | **string** | Current payment status:&lt;br/&gt; &amp;bull; PENDING: means that this payment has been requested, but not yet executed.&lt;br/&gt; &amp;bull; SUCCESSFUL: means that this payment has been successfully executed.&lt;br/&gt; &amp;bull; NOT_SUCCESSFUL: means that this payment could not be executed successfully.&lt;br/&gt; &amp;bull; DISCARDED: means that this payment was discarded because another payment was requested for the same account before this payment was executed. | [default to null]
**BankMessage** | **string** | Contains the bank&#39;s response to the execution of this payment. This field is not set until the payment gets executed. Note that even after the execution of the payment, the field might still not be set, if the bank did not send any response message. | [optional] [default to null]
**Amount** | **float32** | Total money amount of the payment order(s), as absolute value | [default to null]
**OrderCount** | **int32** | Total count of orders included in this payment | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


