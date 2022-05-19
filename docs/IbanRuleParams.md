# IbanRuleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Iban** | **string** | IBAN (case-insensitive) | 
**CategoryId** | **int64** | ID of the category that this rule should assign to the matching transactions | 
**Direction** | [**CategorizationRuleDirection**](CategorizationRuleDirection.md) | &lt;strong&gt;Type:&lt;/strong&gt; CategorizationRuleDirection&lt;br/&gt; Direction for the rule. &#39;Income&#39; means that the rule applies to transactions with a positive amount only, &#39;Spending&#39; means it applies to transactions with a negative amount only. &#39;Both&#39; means that it applies to both kind of transactions. Note that in case of &#39;Both&#39;, finAPI will create two individual rules (one with direction &#39;Income&#39; and one with direction &#39;Spending&#39;). | 

## Methods

### NewIbanRuleParams

`func NewIbanRuleParams(iban string, categoryId int64, direction CategorizationRuleDirection, ) *IbanRuleParams`

NewIbanRuleParams instantiates a new IbanRuleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbanRuleParamsWithDefaults

`func NewIbanRuleParamsWithDefaults() *IbanRuleParams`

NewIbanRuleParamsWithDefaults instantiates a new IbanRuleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIban

`func (o *IbanRuleParams) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *IbanRuleParams) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *IbanRuleParams) SetIban(v string)`

SetIban sets Iban field to given value.


### GetCategoryId

`func (o *IbanRuleParams) GetCategoryId() int64`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *IbanRuleParams) GetCategoryIdOk() (*int64, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *IbanRuleParams) SetCategoryId(v int64)`

SetCategoryId sets CategoryId field to given value.


### GetDirection

`func (o *IbanRuleParams) GetDirection() CategorizationRuleDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *IbanRuleParams) GetDirectionOk() (*CategorizationRuleDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *IbanRuleParams) SetDirection(v CategorizationRuleDirection)`

SetDirection sets Direction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


