# KeywordRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Rule identifier | 
**Category** | [**Category**](Category.md) | &lt;strong&gt;Type:&lt;/strong&gt; Category&lt;br/&gt; The category that this rule assigns to the transactions that it matches | 
**Direction** | [**TransactionDirection**](TransactionDirection.md) | &lt;strong&gt;Type:&lt;/strong&gt; TransactionDirection&lt;br/&gt; Direction for the rule. &#39;Income&#39; means that the rule applies to transactions with a positive amount only, &#39;Spending&#39; means it applies to transactions with a negative amount only. | 
**CreationDate** | **string** | Timestamp of when the rule was created, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time) | 
**Keywords** | **[]string** |  | 
**AllKeywordsMustMatch** | **bool** | This field is only relevant if the rule contains multiple keywords. If set to &#39;true&#39; it means that all keywords have to be found in a transaction to apply the given category. If set to &#39;false&#39;, then even a single matching keyword in a transaction can trigger this rule. | 

## Methods

### NewKeywordRule

`func NewKeywordRule(id int64, category Category, direction TransactionDirection, creationDate string, keywords []string, allKeywordsMustMatch bool, ) *KeywordRule`

NewKeywordRule instantiates a new KeywordRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeywordRuleWithDefaults

`func NewKeywordRuleWithDefaults() *KeywordRule`

NewKeywordRuleWithDefaults instantiates a new KeywordRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeywordRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeywordRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeywordRule) SetId(v int64)`

SetId sets Id field to given value.


### GetCategory

`func (o *KeywordRule) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *KeywordRule) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *KeywordRule) SetCategory(v Category)`

SetCategory sets Category field to given value.


### GetDirection

`func (o *KeywordRule) GetDirection() TransactionDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *KeywordRule) GetDirectionOk() (*TransactionDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *KeywordRule) SetDirection(v TransactionDirection)`

SetDirection sets Direction field to given value.


### GetCreationDate

`func (o *KeywordRule) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *KeywordRule) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *KeywordRule) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.


### GetKeywords

`func (o *KeywordRule) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *KeywordRule) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *KeywordRule) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.


### GetAllKeywordsMustMatch

`func (o *KeywordRule) GetAllKeywordsMustMatch() bool`

GetAllKeywordsMustMatch returns the AllKeywordsMustMatch field if non-nil, zero value otherwise.

### GetAllKeywordsMustMatchOk

`func (o *KeywordRule) GetAllKeywordsMustMatchOk() (*bool, bool)`

GetAllKeywordsMustMatchOk returns a tuple with the AllKeywordsMustMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllKeywordsMustMatch

`func (o *KeywordRule) SetAllKeywordsMustMatch(v bool)`

SetAllKeywordsMustMatch sets AllKeywordsMustMatch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


