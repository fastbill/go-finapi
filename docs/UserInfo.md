# UserInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User&#39;s identifier | 
**RegistrationDate** | **string** | User&#39;s registration date, in the format &#39;YYYY-MM-DD&#39; | 
**DeletionDate** | **NullableString** | User&#39;s deletion date, in the format &#39;YYYY-MM-DD&#39;. May be null if the user has not been deleted. | 
**LastActiveDate** | **NullableString** | User&#39;s last active date, in the format &#39;YYYY-MM-DD&#39;. May be null if the user has not yet logged in. | 
**BankConnectionCount** | **int32** | Number of bank connections that currently exist for this user. | 
**LatestBankConnectionImportDate** | **NullableString** | Latest date of when a bank connection was imported for this user, in the format &#39;YYYY-MM-DD&#39;. This field is null when there has never been a bank connection import. | 
**LatestBankConnectionDeletionDate** | **NullableString** | Latest date of when a bank connection was deleted for this user, in the format &#39;YYYY-MM-DD&#39;. This field is null when there has never been a bank connection deletion. | 
**MonthlyStats** | [**[]MonthlyUserStats**](MonthlyUserStats.md) | &lt;strong&gt;Type:&lt;/strong&gt; MonthlyUserStats&lt;br/&gt; Additional information about the user&#39;s data or activities, broken down in months. The list will by default contain an entry for each month starting with the month of when the user was registered, up to the current month. The date range may vary when you have limited it in the request. &lt;br/&gt;&lt;br/&gt;Please note:&lt;br/&gt;&amp;bull; this field is only set when &#39;includeMonthlyStats&#39; &#x3D; true, otherwise it will be null.&lt;br/&gt;&amp;bull; the list is always ordered from the latest month first, to the oldest month last.&lt;br/&gt;&amp;bull; the list will never contain an entry for a month that was prior to the month of when the user was registered, or after the month of when the user was deleted, even when you have explicitly set a respective date range. This means that the list may be empty if you are requesting a date range where the user didn&#39;t exist yet, or didn&#39;t exist any longer. | 
**IsLocked** | **bool** | Whether the user is currently locked (for further information, see the &#39;maxUserLoginAttempts&#39; setting in your client&#39;s configuration). Note that deleted users will always have this field set to &#39;false&#39;. | 

## Methods

### NewUserInfo

`func NewUserInfo(userId string, registrationDate string, deletionDate NullableString, lastActiveDate NullableString, bankConnectionCount int32, latestBankConnectionImportDate NullableString, latestBankConnectionDeletionDate NullableString, monthlyStats []MonthlyUserStats, isLocked bool, ) *UserInfo`

NewUserInfo instantiates a new UserInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserInfoWithDefaults

`func NewUserInfoWithDefaults() *UserInfo`

NewUserInfoWithDefaults instantiates a new UserInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetRegistrationDate

`func (o *UserInfo) GetRegistrationDate() string`

GetRegistrationDate returns the RegistrationDate field if non-nil, zero value otherwise.

### GetRegistrationDateOk

`func (o *UserInfo) GetRegistrationDateOk() (*string, bool)`

GetRegistrationDateOk returns a tuple with the RegistrationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrationDate

`func (o *UserInfo) SetRegistrationDate(v string)`

SetRegistrationDate sets RegistrationDate field to given value.


### GetDeletionDate

`func (o *UserInfo) GetDeletionDate() string`

GetDeletionDate returns the DeletionDate field if non-nil, zero value otherwise.

### GetDeletionDateOk

`func (o *UserInfo) GetDeletionDateOk() (*string, bool)`

GetDeletionDateOk returns a tuple with the DeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionDate

`func (o *UserInfo) SetDeletionDate(v string)`

SetDeletionDate sets DeletionDate field to given value.


### SetDeletionDateNil

`func (o *UserInfo) SetDeletionDateNil(b bool)`

 SetDeletionDateNil sets the value for DeletionDate to be an explicit nil

### UnsetDeletionDate
`func (o *UserInfo) UnsetDeletionDate()`

UnsetDeletionDate ensures that no value is present for DeletionDate, not even an explicit nil
### GetLastActiveDate

`func (o *UserInfo) GetLastActiveDate() string`

GetLastActiveDate returns the LastActiveDate field if non-nil, zero value otherwise.

### GetLastActiveDateOk

`func (o *UserInfo) GetLastActiveDateOk() (*string, bool)`

GetLastActiveDateOk returns a tuple with the LastActiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDate

`func (o *UserInfo) SetLastActiveDate(v string)`

SetLastActiveDate sets LastActiveDate field to given value.


### SetLastActiveDateNil

`func (o *UserInfo) SetLastActiveDateNil(b bool)`

 SetLastActiveDateNil sets the value for LastActiveDate to be an explicit nil

### UnsetLastActiveDate
`func (o *UserInfo) UnsetLastActiveDate()`

UnsetLastActiveDate ensures that no value is present for LastActiveDate, not even an explicit nil
### GetBankConnectionCount

`func (o *UserInfo) GetBankConnectionCount() int32`

GetBankConnectionCount returns the BankConnectionCount field if non-nil, zero value otherwise.

### GetBankConnectionCountOk

`func (o *UserInfo) GetBankConnectionCountOk() (*int32, bool)`

GetBankConnectionCountOk returns a tuple with the BankConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankConnectionCount

`func (o *UserInfo) SetBankConnectionCount(v int32)`

SetBankConnectionCount sets BankConnectionCount field to given value.


### GetLatestBankConnectionImportDate

`func (o *UserInfo) GetLatestBankConnectionImportDate() string`

GetLatestBankConnectionImportDate returns the LatestBankConnectionImportDate field if non-nil, zero value otherwise.

### GetLatestBankConnectionImportDateOk

`func (o *UserInfo) GetLatestBankConnectionImportDateOk() (*string, bool)`

GetLatestBankConnectionImportDateOk returns a tuple with the LatestBankConnectionImportDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestBankConnectionImportDate

`func (o *UserInfo) SetLatestBankConnectionImportDate(v string)`

SetLatestBankConnectionImportDate sets LatestBankConnectionImportDate field to given value.


### SetLatestBankConnectionImportDateNil

`func (o *UserInfo) SetLatestBankConnectionImportDateNil(b bool)`

 SetLatestBankConnectionImportDateNil sets the value for LatestBankConnectionImportDate to be an explicit nil

### UnsetLatestBankConnectionImportDate
`func (o *UserInfo) UnsetLatestBankConnectionImportDate()`

UnsetLatestBankConnectionImportDate ensures that no value is present for LatestBankConnectionImportDate, not even an explicit nil
### GetLatestBankConnectionDeletionDate

`func (o *UserInfo) GetLatestBankConnectionDeletionDate() string`

GetLatestBankConnectionDeletionDate returns the LatestBankConnectionDeletionDate field if non-nil, zero value otherwise.

### GetLatestBankConnectionDeletionDateOk

`func (o *UserInfo) GetLatestBankConnectionDeletionDateOk() (*string, bool)`

GetLatestBankConnectionDeletionDateOk returns a tuple with the LatestBankConnectionDeletionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestBankConnectionDeletionDate

`func (o *UserInfo) SetLatestBankConnectionDeletionDate(v string)`

SetLatestBankConnectionDeletionDate sets LatestBankConnectionDeletionDate field to given value.


### SetLatestBankConnectionDeletionDateNil

`func (o *UserInfo) SetLatestBankConnectionDeletionDateNil(b bool)`

 SetLatestBankConnectionDeletionDateNil sets the value for LatestBankConnectionDeletionDate to be an explicit nil

### UnsetLatestBankConnectionDeletionDate
`func (o *UserInfo) UnsetLatestBankConnectionDeletionDate()`

UnsetLatestBankConnectionDeletionDate ensures that no value is present for LatestBankConnectionDeletionDate, not even an explicit nil
### GetMonthlyStats

`func (o *UserInfo) GetMonthlyStats() []MonthlyUserStats`

GetMonthlyStats returns the MonthlyStats field if non-nil, zero value otherwise.

### GetMonthlyStatsOk

`func (o *UserInfo) GetMonthlyStatsOk() (*[]MonthlyUserStats, bool)`

GetMonthlyStatsOk returns a tuple with the MonthlyStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyStats

`func (o *UserInfo) SetMonthlyStats(v []MonthlyUserStats)`

SetMonthlyStats sets MonthlyStats field to given value.


### SetMonthlyStatsNil

`func (o *UserInfo) SetMonthlyStatsNil(b bool)`

 SetMonthlyStatsNil sets the value for MonthlyStats to be an explicit nil

### UnsetMonthlyStats
`func (o *UserInfo) UnsetMonthlyStats()`

UnsetMonthlyStats ensures that no value is present for MonthlyStats, not even an explicit nil
### GetIsLocked

`func (o *UserInfo) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UserInfo) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UserInfo) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


