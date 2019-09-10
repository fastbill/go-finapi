# KeywordRuleParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int64** | ID of the category that this rule should assign to the matching transactions | [default to null]
**Direction** | **string** | Direction for the rule. &#39;Income&#39; means that the rule applies to transactions with a positive amount only, &#39;Spending&#39; means it applies to transactions with a negative amount only. &#39;Both&#39; means that it applies to both kind of transactions. Note that in case of &#39;Both&#39;, finAPI will create two individual rules (one with direction &#39;Income&#39; and one with direction &#39;Spending&#39;). | [default to null]
**Keywords** | **[]string** | Set of keywords for the rule (Keywords are regarded case-insensitive). The minimum number of keywords is 1. The maximum number of keywords is 100. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


