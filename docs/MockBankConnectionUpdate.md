# MockBankConnectionUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankConnectionId** | **int64** | Bank connection identifier | [default to null]
**SimulateBankLoginError** | **bool** | Whether to simulate the case that the update fails due to incorrect banking credentials. Note that there is no real communication to any bank server involved, so you won&#39;t lock your accounts when enabling this flag. Default value is &#39;false&#39;. | [optional] [default to null]
**MockAccountsData** | [**[]MockAccountData**](MockAccountData.md) | Mock accounts data. Note that for accounts that exist in a bank connection but that you do not specify in this list, the service will act like those accounts are not received by the bank servers. This means that any accounts that you do not specify here will be marked as deprecated. If you do not specify this list at all, all accounts in the bank connection will be marked as deprecated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


