# ErrorDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Error message | [optional] [default to null]
**Code** | **string** | Error code. See the documentation of the individual services for details about what values may be returned. | [default to null]
**Type_** | **string** | Error type. BUSINESS errors depict German error messages for the user, e.g. from a bank server. TECHNICAL errors depict internal errors. | [default to null]
**MultiStepAuthentication** | [***MultiStepAuthenticationChallenge**](MultiStepAuthenticationChallenge.md) | This field is set when a multi-step authentication is required, i.e. when you need to repeat the original service call and provide additional data. The field contains information about what additional data is required. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


