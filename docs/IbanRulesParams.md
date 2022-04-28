# IbanRulesParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IbanRules** | [**[]IbanRuleParams**](IbanRuleParams.md) | &lt;strong&gt;Type:&lt;/strong&gt; IbanRuleParams&lt;br/&gt; IBAN rule definitions. The minimum number of rule definitions is 1. The maximum number of rule definitions is 100. | 

## Methods

### NewIbanRulesParams

`func NewIbanRulesParams(ibanRules []IbanRuleParams, ) *IbanRulesParams`

NewIbanRulesParams instantiates a new IbanRulesParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIbanRulesParamsWithDefaults

`func NewIbanRulesParamsWithDefaults() *IbanRulesParams`

NewIbanRulesParamsWithDefaults instantiates a new IbanRulesParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIbanRules

`func (o *IbanRulesParams) GetIbanRules() []IbanRuleParams`

GetIbanRules returns the IbanRules field if non-nil, zero value otherwise.

### GetIbanRulesOk

`func (o *IbanRulesParams) GetIbanRulesOk() (*[]IbanRuleParams, bool)`

GetIbanRulesOk returns a tuple with the IbanRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIbanRules

`func (o *IbanRulesParams) SetIbanRules(v []IbanRuleParams)`

SetIbanRules sets IbanRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


