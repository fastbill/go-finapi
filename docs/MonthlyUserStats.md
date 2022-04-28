# MonthlyUserStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Month** | **string** | The month that the contained information applies to, in the format &#39;YYYY-MM&#39;. | 
**MinBankConnectionCount** | **int32** | Minimum count of bank connections that this user has had at any point during the month. | 
**MaxBankConnectionCount** | **int32** | Maximum count of bank connections that this user has had at any point during the month. | 

## Methods

### NewMonthlyUserStats

`func NewMonthlyUserStats(month string, minBankConnectionCount int32, maxBankConnectionCount int32, ) *MonthlyUserStats`

NewMonthlyUserStats instantiates a new MonthlyUserStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthlyUserStatsWithDefaults

`func NewMonthlyUserStatsWithDefaults() *MonthlyUserStats`

NewMonthlyUserStatsWithDefaults instantiates a new MonthlyUserStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonth

`func (o *MonthlyUserStats) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *MonthlyUserStats) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *MonthlyUserStats) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetMinBankConnectionCount

`func (o *MonthlyUserStats) GetMinBankConnectionCount() int32`

GetMinBankConnectionCount returns the MinBankConnectionCount field if non-nil, zero value otherwise.

### GetMinBankConnectionCountOk

`func (o *MonthlyUserStats) GetMinBankConnectionCountOk() (*int32, bool)`

GetMinBankConnectionCountOk returns a tuple with the MinBankConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinBankConnectionCount

`func (o *MonthlyUserStats) SetMinBankConnectionCount(v int32)`

SetMinBankConnectionCount sets MinBankConnectionCount field to given value.


### GetMaxBankConnectionCount

`func (o *MonthlyUserStats) GetMaxBankConnectionCount() int32`

GetMaxBankConnectionCount returns the MaxBankConnectionCount field if non-nil, zero value otherwise.

### GetMaxBankConnectionCountOk

`func (o *MonthlyUserStats) GetMaxBankConnectionCountOk() (*int32, bool)`

GetMaxBankConnectionCountOk returns a tuple with the MaxBankConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBankConnectionCount

`func (o *MonthlyUserStats) SetMaxBankConnectionCount(v int32)`

SetMaxBankConnectionCount sets MaxBankConnectionCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


