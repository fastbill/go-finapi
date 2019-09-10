# ErrorMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ErrorDetails**](ErrorDetails.md) | List of errors | [default to null]
**Date** | **string** | Server date of when the error(s) occurred, in the format YYYY-MM-DD HH:MM:SS.SSS | [default to null]
**RequestId** | **string** | ID of the request that caused this error. This is either what you have passed for the header &#39;X-REQUEST-ID&#39;, or an auto-generated ID in case you didn&#39;t pass any value for that header. | [optional] [default to null]
**Endpoint** | **string** | The service endpoint that was called | [default to null]
**AuthContext** | **string** | Information about the authorization context of the service call | [default to null]
**Bank** | **string** | BLZ and name (in format \&quot;&lt;BLZ&gt; - &lt;name&gt;\&quot;) of a bank that was used for the original request | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


