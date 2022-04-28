# IbanRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Rule identifier | 
**Category** | [**Category**](Category.md) | &lt;strong&gt;Type:&lt;/strong&gt; Category&lt;br/&gt; The category that this rule assigns to the transactions that it matches | 
**Direction** | [**TransactionDirection**](TransactionDirection.md) | &lt;strong&gt;Type:&lt;/strong&gt; TransactionDirection&lt;br/&gt; Direction for the rule. &#39;Income&#39; means that the rule applies to transactions with a positive amount only, &#39;Spending&#39; means it applies to transactions with a negative amount only. | 
**CreationDate** | **string** | Timestamp of when the rule was created, in the format &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time) | 
**Iban** | **string** | The IBAN for this rule | 

## Methods

### NewIbanRule

`func NewIbanRule(id int64, category Category, direction TransactionDirection, creationDate string, iban string, ) *IbanRule`

NewIbanRule instantiates a new IbanRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbanRuleWithDefaults

`func NewIbanRuleWithDefaults() *IbanRule`

NewIbanRuleWithDefaults instantiates a new IbanRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IbanRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IbanRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IbanRule) SetId(v int64)`

SetId sets Id field to given value.


### GetCategory

`func (o *IbanRule) GetCategory() Category`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IbanRule) GetCategoryOk() (*Category, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IbanRule) SetCategory(v Category)`

SetCategory sets Category field to given value.


### GetDirection

`func (o *IbanRule) GetDirection() TransactionDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *IbanRule) GetDirectionOk() (*TransactionDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *IbanRule) SetDirection(v TransactionDirection)`

SetDirection sets Direction field to given value.


### GetCreationDate

`func (o *IbanRule) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *IbanRule) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *IbanRule) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.


### GetIban

`func (o *IbanRule) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *IbanRule) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *IbanRule) SetIban(v string)`

SetIban sets Iban field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


