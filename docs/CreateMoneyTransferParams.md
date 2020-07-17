# CreateMoneyTransferParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | Identifier of the account that should be used for the payment. | [default to null]
**MoneyTransfers** | [**[]MoneyTransferOrderParams**](MoneyTransferOrderParams.md) | List of money transfer orders (may contain at most 15000 items). Please note that collective money transfer may not always be supported. | [default to null]
**SingleBooking** | **bool** | This field is only relevant when you pass multiple orders. It determines whether the orders should be processed by the bank as one collective booking (in case of &#39;false&#39;), or as single bookings (in case of &#39;true&#39;). Note that it is subject to the bank whether it will regard the field. Default value is &#39;false&#39;. | [optional] [default to null]
**ExecutionDate** | **string** | Execution date for the money transfer(s), in the format &#39;YYYY-MM-DD&#39;. May not be in the past. If not specified, then the current date will be used. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


