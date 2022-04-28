# EditBankConnectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New name for the bank connection. Maximum length is 64. If you do not want to change the current name let this field remain unset. If you want to clear the current name, set the field&#39;s value to an empty string (\&quot;\&quot;).&lt;br/&gt;&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If you are a Web Form 2.0 customer, and would like to update the name of your bank connection, please use the API parameter. | [optional] 
**BankingUserId** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; and &#39;interface&#39; fields instead. If any of those two fields is used, then the value of this field will be ignored.&lt;br/&gt;&lt;br/&gt;New online banking user ID. If you do not want to change the current user ID let this field remain unset. In case you need to use finAPI&#39;s Web Form to let the user update the field, just set the field to any value, so that the service recognizes that you wish to use the Web Form flow. Note that you cannot clear the current user ID, i.e. a bank connection must always have a user ID (except for when it is a &#39;demo connection&#39;). Max length: 170. | [optional] 
**BankingCustomerId** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; and &#39;interface&#39; fields instead. If any of those two fields is used, then the value of this field will be ignored.&lt;br/&gt;&lt;br/&gt;New online banking customer ID. If you do not want to change the current customer ID let this field remain unset. In case you need to use finAPI&#39;s Web Form to let the user update the field, just set the field to non-empty value, so that the service recognizes that you wish to use the Web Form flow. If you want to clear the current customer ID, set the field&#39;s value to an empty string (\&quot;\&quot;). Max length: 170. | [optional] 
**BankingPin** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;loginCredentials&#39; and &#39;interface&#39; fields instead. If any of those two fields is used, then the value of this field will be ignored.&lt;br/&gt;&lt;br/&gt;New online banking PIN. If you do not want to change the current PIN let this field remain unset. In case you need to use finAPI&#39;s Web Form to let the user update the field, just set the field to non-empty value, so that the service recognizes that you wish to use the Web Form flow. If you want to clear the current PIN, set the field&#39;s value to an empty string (\&quot;\&quot;).&lt;br/&gt;&lt;br/&gt;Any symbols are allowed. Max length: 170. | [optional] 
**Interface** | Pointer to [**BankingInterface**](BankingInterface.md) | &lt;strong&gt;Type:&lt;/strong&gt; BankingInterface&lt;br/&gt; The interface for which you want to edit data. Must be given when you pass &#39;loginCredentials&#39; and/or a &#39;defaultTwoStepProcedureId&#39;. | [optional] 
**LoginCredentials** | Pointer to [**[]LoginCredential**](LoginCredential.md) | &lt;strong&gt;Type:&lt;/strong&gt; LoginCredential&lt;br/&gt; Set of login credentials that you want to edit. Must be passed in combination with the &#39;interface&#39; field. The labels that you pass must match with the login credential labels that the respective interface defines. If you want to clear the stored value for a credential, you can pass an empty string (\&quot;\&quot;) as value.In case you need to use finAPI&#39;s Web Form to let the user update the login credentials, send all fields the user wishes to update with a non-empty value.In case all fields contain an empty string (\&quot;\&quot;), no Web Form will be generated. Note that any change in the credentials will automatically remove the saved consent data associated with those credentials.&lt;br/&gt;&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If you are a  Web Form 2.0 customer, and would like to allow your end-users to change the credentials they have stored in our system, then please navigate &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/tasks/backgroundUpdate&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to implement the same functionality. | [optional] 
**DefaultTwoStepProcedureId** | Pointer to **string** | NOTE: In the future, this field will work only in combination with the &#39;interface&#39; field.&lt;br/&gt;&lt;br/&gt;New default two-step-procedure. Must match the &#39;procedureId&#39; of one of the procedures that are listed in the bank connection. If you do not want to change this field let it remain unset. If you want to clear the current default two-step-procedure, set the field&#39;s value to an empty string (\&quot;\&quot;).&lt;br/&gt;&lt;br/&gt;&lt;strong&gt;NOTE:&lt;/strong&gt; If you are a Web Form 2.0 customer and would like to allow your end users to update their preferred TAN procedure that is stored in our system, then please navigate &lt;a target&#x3D;\&quot;_blank\&quot; href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/tasks/backgroundUpdate&#39;&gt;here&lt;/a&gt; to implement the same functionality. | [optional] 

## Methods

### NewEditBankConnectionParams

`func NewEditBankConnectionParams() *EditBankConnectionParams`

NewEditBankConnectionParams instantiates a new EditBankConnectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditBankConnectionParamsWithDefaults

`func NewEditBankConnectionParamsWithDefaults() *EditBankConnectionParams`

NewEditBankConnectionParamsWithDefaults instantiates a new EditBankConnectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditBankConnectionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditBankConnectionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditBankConnectionParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditBankConnectionParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBankingUserId

`func (o *EditBankConnectionParams) GetBankingUserId() string`

GetBankingUserId returns the BankingUserId field if non-nil, zero value otherwise.

### GetBankingUserIdOk

`func (o *EditBankConnectionParams) GetBankingUserIdOk() (*string, bool)`

GetBankingUserIdOk returns a tuple with the BankingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingUserId

`func (o *EditBankConnectionParams) SetBankingUserId(v string)`

SetBankingUserId sets BankingUserId field to given value.

### HasBankingUserId

`func (o *EditBankConnectionParams) HasBankingUserId() bool`

HasBankingUserId returns a boolean if a field has been set.

### GetBankingCustomerId

`func (o *EditBankConnectionParams) GetBankingCustomerId() string`

GetBankingCustomerId returns the BankingCustomerId field if non-nil, zero value otherwise.

### GetBankingCustomerIdOk

`func (o *EditBankConnectionParams) GetBankingCustomerIdOk() (*string, bool)`

GetBankingCustomerIdOk returns a tuple with the BankingCustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingCustomerId

`func (o *EditBankConnectionParams) SetBankingCustomerId(v string)`

SetBankingCustomerId sets BankingCustomerId field to given value.

### HasBankingCustomerId

`func (o *EditBankConnectionParams) HasBankingCustomerId() bool`

HasBankingCustomerId returns a boolean if a field has been set.

### GetBankingPin

`func (o *EditBankConnectionParams) GetBankingPin() string`

GetBankingPin returns the BankingPin field if non-nil, zero value otherwise.

### GetBankingPinOk

`func (o *EditBankConnectionParams) GetBankingPinOk() (*string, bool)`

GetBankingPinOk returns a tuple with the BankingPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankingPin

`func (o *EditBankConnectionParams) SetBankingPin(v string)`

SetBankingPin sets BankingPin field to given value.

### HasBankingPin

`func (o *EditBankConnectionParams) HasBankingPin() bool`

HasBankingPin returns a boolean if a field has been set.

### GetInterface

`func (o *EditBankConnectionParams) GetInterface() BankingInterface`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *EditBankConnectionParams) GetInterfaceOk() (*BankingInterface, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *EditBankConnectionParams) SetInterface(v BankingInterface)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *EditBankConnectionParams) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetLoginCredentials

`func (o *EditBankConnectionParams) GetLoginCredentials() []LoginCredential`

GetLoginCredentials returns the LoginCredentials field if non-nil, zero value otherwise.

### GetLoginCredentialsOk

`func (o *EditBankConnectionParams) GetLoginCredentialsOk() (*[]LoginCredential, bool)`

GetLoginCredentialsOk returns a tuple with the LoginCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginCredentials

`func (o *EditBankConnectionParams) SetLoginCredentials(v []LoginCredential)`

SetLoginCredentials sets LoginCredentials field to given value.

### HasLoginCredentials

`func (o *EditBankConnectionParams) HasLoginCredentials() bool`

HasLoginCredentials returns a boolean if a field has been set.

### GetDefaultTwoStepProcedureId

`func (o *EditBankConnectionParams) GetDefaultTwoStepProcedureId() string`

GetDefaultTwoStepProcedureId returns the DefaultTwoStepProcedureId field if non-nil, zero value otherwise.

### GetDefaultTwoStepProcedureIdOk

`func (o *EditBankConnectionParams) GetDefaultTwoStepProcedureIdOk() (*string, bool)`

GetDefaultTwoStepProcedureIdOk returns a tuple with the DefaultTwoStepProcedureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTwoStepProcedureId

`func (o *EditBankConnectionParams) SetDefaultTwoStepProcedureId(v string)`

SetDefaultTwoStepProcedureId sets DefaultTwoStepProcedureId field to given value.

### HasDefaultTwoStepProcedureId

`func (o *EditBankConnectionParams) HasDefaultTwoStepProcedureId() bool`

HasDefaultTwoStepProcedureId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


