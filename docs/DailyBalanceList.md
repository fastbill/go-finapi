# DailyBalanceList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatestCommonBalanceTimestamp** | **NullableString** | The latestCommonBalanceTimestamp is the latest timestamp at which all regarded  accounts have been up to date. Only balances with their date being smaller than the latestCommonBalanceTimestamp are reliable. Example: A user has two accounts: A (last update today, so balance from today) and B (last update yesterday, so balance from yesterday). The service /accounts/dailyBalances will return a balance for yesterday and for today, with the info latestCommonBalanceTimestamp&#x3D;yesterday. Since account B might have received transactions this morning, today&#39;s balance might be wrong. So either make sure that all regarded accounts are up to date before calling this service, or use the results carefully in combination with the latestCommonBalanceTimestamp. The format is &#39;YYYY-MM-DD HH:MM:SS.SSS&#39; (german time). | 
**DailyBalances** | [**[]DailyBalance**](DailyBalance.md) | &lt;strong&gt;Type:&lt;/strong&gt; DailyBalance&lt;br/&gt; List of daily balances for the requested period and account(s) | 
**Paging** | [**Paging**](Paging.md) | &lt;strong&gt;Type:&lt;/strong&gt; Paging&lt;br/&gt; Information for pagination | 

## Methods

### NewDailyBalanceList

`func NewDailyBalanceList(latestCommonBalanceTimestamp NullableString, dailyBalances []DailyBalance, paging Paging, ) *DailyBalanceList`

NewDailyBalanceList instantiates a new DailyBalanceList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyBalanceListWithDefaults

`func NewDailyBalanceListWithDefaults() *DailyBalanceList`

NewDailyBalanceListWithDefaults instantiates a new DailyBalanceList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatestCommonBalanceTimestamp

`func (o *DailyBalanceList) GetLatestCommonBalanceTimestamp() string`

GetLatestCommonBalanceTimestamp returns the LatestCommonBalanceTimestamp field if non-nil, zero value otherwise.

### GetLatestCommonBalanceTimestampOk

`func (o *DailyBalanceList) GetLatestCommonBalanceTimestampOk() (*string, bool)`

GetLatestCommonBalanceTimestampOk returns a tuple with the LatestCommonBalanceTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestCommonBalanceTimestamp

`func (o *DailyBalanceList) SetLatestCommonBalanceTimestamp(v string)`

SetLatestCommonBalanceTimestamp sets LatestCommonBalanceTimestamp field to given value.


### SetLatestCommonBalanceTimestampNil

`func (o *DailyBalanceList) SetLatestCommonBalanceTimestampNil(b bool)`

 SetLatestCommonBalanceTimestampNil sets the value for LatestCommonBalanceTimestamp to be an explicit nil

### UnsetLatestCommonBalanceTimestamp
`func (o *DailyBalanceList) UnsetLatestCommonBalanceTimestamp()`

UnsetLatestCommonBalanceTimestamp ensures that no value is present for LatestCommonBalanceTimestamp, not even an explicit nil
### GetDailyBalances

`func (o *DailyBalanceList) GetDailyBalances() []DailyBalance`

GetDailyBalances returns the DailyBalances field if non-nil, zero value otherwise.

### GetDailyBalancesOk

`func (o *DailyBalanceList) GetDailyBalancesOk() (*[]DailyBalance, bool)`

GetDailyBalancesOk returns a tuple with the DailyBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyBalances

`func (o *DailyBalanceList) SetDailyBalances(v []DailyBalance)`

SetDailyBalances sets DailyBalances field to given value.


### GetPaging

`func (o *DailyBalanceList) GetPaging() Paging`

GetPaging returns the Paging field if non-nil, zero value otherwise.

### GetPagingOk

`func (o *DailyBalanceList) GetPagingOk() (*Paging, bool)`

GetPagingOk returns a tuple with the Paging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaging

`func (o *DailyBalanceList) SetPaging(v Paging)`

SetPaging sets Paging field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


