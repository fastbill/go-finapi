# NotificationRule

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Notification rule identifier | [default to null]
**TriggerEvent** | **string** | Trigger event type | [default to null]
**Params** | **map[string]string** | Additional parameters that are specific to the trigger event type. Please refer to the documentation for details. | [optional] [default to null]
**CallbackHandle** | **string** | The string that finAPI includes into the notifications that it sends based on this rule. | [optional] [default to null]
**IncludeDetails** | **bool** | Whether the notification messages that will be sent based on this rule contain encrypted detailed data or not | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


