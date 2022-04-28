# NotificationRuleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerEvent** | **string** | Trigger event type | 
**Params** | Pointer to **map[string]string** | Additional parameters that are specific to the chosen trigger event type. Please refer to the documentation for details. | [optional] 
**CallbackHandle** | Pointer to **string** | An arbitrary string that finAPI will include into the notifications that it sends based on this rule and that you can use to identify the notification in your application. For instance, you could include the identifier of the user that you create this rule for. Maximum allowed length of the string is 512 characters.&lt;br/&gt;&lt;br/&gt;Note that for this parameter, you can pass the symbols &#39;/&#39;, &#39;&#x3D;&#39;, &#39;%&#39; and &#39;\&quot;&#39; in addition to the symbols that are generally allowed in finAPI (see &lt;a href&#x3D;&#39;https://documentation.finapi.io/access/Allowed-Characters.2764767279.html&#39; target&#x3D;&#39;_blank&#39;&gt;Allowed Characters&lt;/a&gt;). This was done to enable you to set Base64 encoded strings and JSON structures for the callback handle. | [optional] 
**IncludeDetails** | Pointer to **bool** | Whether the notification messages that will be sent based on this rule should contain encrypted detailed data or not. Default value is &#39;false&#39;. | [optional] [default to false]

## Methods

### NewNotificationRuleParams

`func NewNotificationRuleParams(triggerEvent string, ) *NotificationRuleParams`

NewNotificationRuleParams instantiates a new NotificationRuleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleParamsWithDefaults

`func NewNotificationRuleParamsWithDefaults() *NotificationRuleParams`

NewNotificationRuleParamsWithDefaults instantiates a new NotificationRuleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerEvent

`func (o *NotificationRuleParams) GetTriggerEvent() string`

GetTriggerEvent returns the TriggerEvent field if non-nil, zero value otherwise.

### GetTriggerEventOk

`func (o *NotificationRuleParams) GetTriggerEventOk() (*string, bool)`

GetTriggerEventOk returns a tuple with the TriggerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerEvent

`func (o *NotificationRuleParams) SetTriggerEvent(v string)`

SetTriggerEvent sets TriggerEvent field to given value.


### GetParams

`func (o *NotificationRuleParams) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NotificationRuleParams) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NotificationRuleParams) SetParams(v map[string]string)`

SetParams sets Params field to given value.

### HasParams

`func (o *NotificationRuleParams) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetCallbackHandle

`func (o *NotificationRuleParams) GetCallbackHandle() string`

GetCallbackHandle returns the CallbackHandle field if non-nil, zero value otherwise.

### GetCallbackHandleOk

`func (o *NotificationRuleParams) GetCallbackHandleOk() (*string, bool)`

GetCallbackHandleOk returns a tuple with the CallbackHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackHandle

`func (o *NotificationRuleParams) SetCallbackHandle(v string)`

SetCallbackHandle sets CallbackHandle field to given value.

### HasCallbackHandle

`func (o *NotificationRuleParams) HasCallbackHandle() bool`

HasCallbackHandle returns a boolean if a field has been set.

### GetIncludeDetails

`func (o *NotificationRuleParams) GetIncludeDetails() bool`

GetIncludeDetails returns the IncludeDetails field if non-nil, zero value otherwise.

### GetIncludeDetailsOk

`func (o *NotificationRuleParams) GetIncludeDetailsOk() (*bool, bool)`

GetIncludeDetailsOk returns a tuple with the IncludeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDetails

`func (o *NotificationRuleParams) SetIncludeDetails(v bool)`

SetIncludeDetails sets IncludeDetails field to given value.

### HasIncludeDetails

`func (o *NotificationRuleParams) HasIncludeDetails() bool`

HasIncludeDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


