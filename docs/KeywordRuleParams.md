# KeywordRuleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategoryId** | **int64** | ID of the category that this rule should assign to the matching transactions | 
**Direction** | [**CategorizationRuleDirection**](CategorizationRuleDirection.md) | &lt;strong&gt;Type:&lt;/strong&gt; CategorizationRuleDirection&lt;br/&gt; Direction for the rule. &#39;Income&#39; means that the rule applies to transactions with a positive amount only, &#39;Spending&#39; means it applies to transactions with a negative amount only. &#39;Both&#39; means that it applies to both kind of transactions. Note that in case of &#39;Both&#39;, finAPI will create two individual rules (one with direction &#39;Income&#39; and one with direction &#39;Spending&#39;). | 
**Keywords** | **[]string** |  | 
**AllKeywordsMustMatch** | Pointer to **bool** | This field is only relevant if you pass multiple keywords. If set to &#39;true&#39;, it means that all keywords have to be found in a transaction to apply the given category. If set to &#39;false&#39;, then even a single matching keyword in a transaction can trigger this rule. Default value is &#39;false&#39;. | [optional] [default to false]

## Methods

### NewKeywordRuleParams

`func NewKeywordRuleParams(categoryId int64, direction CategorizationRuleDirection, keywords []string, ) *KeywordRuleParams`

NewKeywordRuleParams instantiates a new KeywordRuleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeywordRuleParamsWithDefaults

`func NewKeywordRuleParamsWithDefaults() *KeywordRuleParams`

NewKeywordRuleParamsWithDefaults instantiates a new KeywordRuleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategoryId

`func (o *KeywordRuleParams) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *KeywordRuleParams) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *KeywordRuleParams) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.


### GetDirection

`func (o *KeywordRuleParams) GetDirection() CategorizationRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *KeywordRuleParams) GetDirectionOk() (*CategorizationRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *KeywordRuleParams) SetDirection(v CategorizationRuleDirection)`

SetDirection sets Direction field to given value.


### GetKeywords

`func (o *KeywordRuleParams) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *KeywordRuleParams) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *KeywordRuleParams) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetAllKeywordsMustMatch

`func (o *KeywordRuleParams) GetAllKeywordsMustMatch() bool`

GetAllKeywordsMustMatch returns the AllKeywordsMustMatch field if non-nil, zero value otherwise.

### GetAllKeywordsMustMatchOk

`func (o *KeywordRuleParams) GetAllKeywordsMustMatchOk() (*bool, bool)`

GetAllKeywordsMustMatchOk returns a tuple with the AllKeywordsMustMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllKeywordsMustMatch

`func (o *KeywordRuleParams) SetAllKeywordsMustMatch(v bool)`

SetAllKeywordsMustMatch sets AllKeywordsMustMatch field to given value.

### HasAllKeywordsMustMatch

`func (o *KeywordRuleParams) HasAllKeywordsMustMatch() bool`

HasAllKeywordsMustMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


