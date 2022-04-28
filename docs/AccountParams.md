# AccountParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountName** | Pointer to **string** | Account name. Maximum length is 512. | [optional] 
**AccountTypeId** | Pointer to **int64** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;accountType&#39; field instead.&lt;br/&gt;&lt;br/&gt;Identifier of account type.&lt;br/&gt;&lt;br/&gt;1 &#x3D; Checking,&lt;br/&gt;2 &#x3D; Savings,&lt;br/&gt;3 &#x3D; CreditCard,&lt;br/&gt;4 &#x3D; Security,&lt;br/&gt;5 &#x3D; Loan,&lt;br/&gt;6 &#x3D; Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),&lt;br/&gt;7 &#x3D; Membership,&lt;br/&gt;8 &#x3D; Bausparen&lt;br/&gt; | [optional] 
**AccountType** | Pointer to [**AccountType**](AccountType.md) | &lt;strong&gt;Type:&lt;/strong&gt; AccountType&lt;br/&gt; An account type.&lt;br/&gt;&lt;br/&gt;Checking,&lt;br/&gt;Savings,&lt;br/&gt;CreditCard,&lt;br/&gt;Security,&lt;br/&gt;Loan,&lt;br/&gt;Pocket (DEPRECATED; will not be returned for any account unless this type has explicitly been set via PATCH),&lt;br/&gt;Membership,&lt;br/&gt;Bausparen&lt;br/&gt;&lt;br/&gt; | [optional] 
**IsNew** | Pointer to **bool** | Whether this account should be flagged as &#39;new&#39; or not. Any newly imported account will have this flag initially set to true, and remain so until you set it to false (see PATCH /accounts/&lt;id&gt;). How you use this field is up to your interpretation, however it is recommended to set the flag to false for all accounts right after the initial import of the bank connection. This way, you will be able recognize accounts that get newly imported during a later update of the bank connection, by checking for any accounts with the flag set to true after every update of the bank connection. | [optional] 

## Methods

### NewAccountParams

`func NewAccountParams() *AccountParams`

NewAccountParams instantiates a new AccountParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountParamsWithDefaults

`func NewAccountParamsWithDefaults() *AccountParams`

NewAccountParamsWithDefaults instantiates a new AccountParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountName

`func (o *AccountParams) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *AccountParams) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *AccountParams) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *AccountParams) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetAccountTypeId

`func (o *AccountParams) GetAccountTypeId() int64`

GetAccountTypeId returns the AccountTypeId field if non-nil, zero value otherwise.

### GetAccountTypeIdOk

`func (o *AccountParams) GetAccountTypeIdOk() (*int64, bool)`

GetAccountTypeIdOk returns a tuple with the AccountTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypeId

`func (o *AccountParams) SetAccountTypeId(v int64)`

SetAccountTypeId sets AccountTypeId field to given value.

### HasAccountTypeId

`func (o *AccountParams) HasAccountTypeId() bool`

HasAccountTypeId returns a boolean if a field has been set.

### GetAccountType

`func (o *AccountParams) GetAccountType() AccountType`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *AccountParams) GetAccountTypeOk() (*AccountType, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *AccountParams) SetAccountType(v AccountType)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *AccountParams) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetIsNew

`func (o *AccountParams) GetIsNew() bool`

GetIsNew returns the IsNew field if non-nil, zero value otherwise.

### GetIsNewOk

`func (o *AccountParams) GetIsNewOk() (*bool, bool)`

GetIsNewOk returns a tuple with the IsNew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNew

`func (o *AccountParams) SetIsNew(v bool)`

SetIsNew sets IsNew field to given value.

### HasIsNew

`func (o *AccountParams) HasIsNew() bool`

HasIsNew returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


