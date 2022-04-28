# PageableKeywordRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeywordRules** | [**[]KeywordRule**](KeywordRule.md) | &lt;strong&gt;Type:&lt;/strong&gt; KeywordRule&lt;br/&gt; List of keyword rules | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewPageableKeywordRuleList

`func NewPageableKeywordRuleList(keywordRules []KeywordRule, paging Paging, ) *PageableKeywordRuleList`

NewPageableKeywordRuleList instantiates a new PageableKeywordRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableKeywordRuleListWithDefaults

`func NewPageableKeywordRuleListWithDefaults() *PageableKeywordRuleList`

NewPageableKeywordRuleListWithDefaults instantiates a new PageableKeywordRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeywordRules

`func (o *PageableKeywordRuleList) GetKeywordRules() []KeywordRule`

GetKeywordRules returns the KeywordRules field if non-nil, zero value otherwise.

### GetKeywordRulesOk

`func (o *PageableKeywordRuleList) GetKeywordRulesOk() (*[]KeywordRule, bool)`

GetKeywordRulesOk returns a tuple with the KeywordRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywordRules

`func (o *PageableKeywordRuleList) SetKeywordRules(v []KeywordRule)`

SetKeywordRules sets KeywordRules field to given value.


### GetPaging

`func (o *PageableKeywordRuleList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableKeywordRuleList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableKeywordRuleList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


