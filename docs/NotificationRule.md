# NotificationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Notification rule identifier | 
**TriggerEvent** | **string** | Trigger event type | 
**Params** | **map[string]string** | Additional parameters that are specific to the trigger event type. Please refer to the documentation for details. | 
**CallbackHandle** | **NullableString** | The string that finAPI includes into the notifications that it sends based on this rule. | 
**IncludeDetails** | **bool** | Whether the notification messages that will be sent based on this rule contain encrypted detailed data or not. | 

## Methods

### NewNotificationRule

`func NewNotificationRule(id int64, triggerEvent string, params map[string]string, callbackHandle NullableString, includeDetails bool, ) *NotificationRule`

NewNotificationRule instantiates a new NotificationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationRuleWithDefaults

`func NewNotificationRuleWithDefaults() *NotificationRule`

NewNotificationRuleWithDefaults instantiates a new NotificationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationRule) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationRule) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationRule) SetId(v int64)`

SetId sets Id field to given value.


### GetTriggerEvent

`func (o *NotificationRule) GetTriggerEvent() string`

GetTriggerEvent returns the TriggerEvent field if non-nil, zero value otherwise.

### GetTriggerEventOk

`func (o *NotificationRule) GetTriggerEventOk() (*string, bool)`

GetTriggerEventOk returns a tuple with the TriggerEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerEvent

`func (o *NotificationRule) SetTriggerEvent(v string)`

SetTriggerEvent sets TriggerEvent field to given value.


### GetParams

`func (o *NotificationRule) GetParams() map[string]string`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *NotificationRule) GetParamsOk() (*map[string]string, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *NotificationRule) SetParams(v map[string]string)`

SetParams sets Params field to given value.


### GetCallbackHandle

`func (o *NotificationRule) GetCallbackHandle() string`

GetCallbackHandle returns the CallbackHandle field if non-nil, zero value otherwise.

### GetCallbackHandleOk

`func (o *NotificationRule) GetCallbackHandleOk() (*string, bool)`

GetCallbackHandleOk returns a tuple with the CallbackHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackHandle

`func (o *NotificationRule) SetCallbackHandle(v string)`

SetCallbackHandle sets CallbackHandle field to given value.


### SetCallbackHandleNil

`func (o *NotificationRule) SetCallbackHandleNil(b bool)`

 SetCallbackHandleNil sets the value for CallbackHandle to be an explicit nil

### UnsetCallbackHandle
`func (o *NotificationRule) UnsetCallbackHandle()`

UnsetCallbackHandle ensures that no value is present for CallbackHandle, not even an explicit nil
### GetIncludeDetails

`func (o *NotificationRule) GetIncludeDetails() bool`

GetIncludeDetails returns the IncludeDetails field if non-nil, zero value otherwise.

### GetIncludeDetailsOk

`func (o *NotificationRule) GetIncludeDetailsOk() (*bool, bool)`

GetIncludeDetailsOk returns a tuple with the IncludeDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeDetails

`func (o *NotificationRule) SetIncludeDetails(v bool)`

SetIncludeDetails sets IncludeDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


