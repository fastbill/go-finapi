# DailyBalanceList

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestCommonBalanceTimestamp** | **string** | The latestCommonBalanceTimestamp is the latest timestamp at which all regarded  accounts have been up to date. Only balances with their date being smaller than the latestCommonBalanceTimestamp are reliable. Example: A user has two accounts: A (last update today, so balance from today) and B (last update yesterday, so balance from yesterday). The service /accounts/dailyBalances will return a balance for yesterday and for today, with the info latestCommonBalanceTimestamp&#x3D;yesterday. Since account B might have received transactions this morning, today&#39;s balance might be wrong. So either make sure that all regarded accounts are up to date before calling this service, or use the results carefully in combination with the latestCommonBalanceTimestamp. The format is &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | [optional] [default to null]
**DailyBalances** | [**[]DailyBalance**](DailyBalance.md) | List of daily balances for the requested period and account(s) | [default to null]
**Paging** | [***Paging**](Paging.md) | Information for pagination | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


