# PageableIbanRuleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IbanRules** | [**[]IbanRule**](IbanRule.md) | &lt;strong&gt;Type:&lt;/strong&gt; IbanRule&lt;br/&gt; List of iban rules information | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewPageableIbanRuleList

`func NewPageableIbanRuleList(ibanRules []IbanRule, paging Paging, ) *PageableIbanRuleList`

NewPageableIbanRuleList instantiates a new PageableIbanRuleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableIbanRuleListWithDefaults

`func NewPageableIbanRuleListWithDefaults() *PageableIbanRuleList`

NewPageableIbanRuleListWithDefaults instantiates a new PageableIbanRuleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbanRules

`func (o *PageableIbanRuleList) GetIbanRules() []IbanRule`

GetIbanRules returns the IbanRules field if non-nil, zero value otherwise.

### GetIbanRulesOk

`func (o *PageableIbanRuleList) GetIbanRulesOk() (*[]IbanRule, bool)`

GetIbanRulesOk returns a tuple with the IbanRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanRules

`func (o *PageableIbanRuleList) SetIbanRules(v []IbanRule)`

SetIbanRules sets IbanRules field to given value.


### GetPaging

`func (o *PageableIbanRuleList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *PageableIbanRuleList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *PageableIbanRuleList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


