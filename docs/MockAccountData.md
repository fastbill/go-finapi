# MockAccountData

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **int64** | Account identifier | [default to null]
**AccountBalance** | **float32** | The balance that this account should be set to. Note that when the balance does not add up to the current balance plus the sum of the transactions you pass in the &#39;newTransactions&#39; field, finAPI will fix the balance deviation with the insertion of an adjusting entry (&#39;Zwischensaldo&#39; transaction). | [default to null]
**NewTransactions** | [**[]NewTransaction**](NewTransaction.md) | New transactions that should be imported into the account (maximum 1000 transactions at once). Please make sure that the value you pass in the &#39;accountBalance&#39; field equals the current account balance plus the sum of the new transactions that you pass here, otherwise finAPI will detect a deviation in the balance and fix it with the insertion of an adjusting entry (&#39;Zwischensaldo&#39; transaction). Please also note that it is not guaranteed that all transactions that you pass here will actually get imported. More specifically, finAPI will ignore any transactions whose booking date is prior to the booking date of the latest currently existing transactions minus 10 days (which is the &#39;update window&#39; that finAPI uses when importing new transactions). Also, finAPI will ignore any transactions that are considered duplicates of already existing transactions within the update window. This is the case for instance when you try to import a new transaction with the same booking date and same amount as an already existing transaction. In such cases, you might get an adjusting entry too (&#39;Zwischensaldo&#39; transaction), as your given balance might not add up to the transactions that will exist in the account after the update. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


