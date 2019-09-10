# MultiStepAuthenticationChallenge

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Hash for this multi-step authentication flow. Must be passed back to finAPI when continuing the flow. | [default to null]
**Status** | **string** | Indicates the current status of the multi-step authentication flow:&lt;br/&gt;&amp;bull; CHALLENGE_RESPONSE_REQUIRED - means that the bank has sent a &#39;challengeMessage&#39; which must be presented to the user. The user&#39;s answer must then be passed back to finAPI as &#39;challengeResponse&#39;;&lt;br/&gt;&amp;bull; TWO_STEP_PROCEDURE_REQUIRED - means that the user has to select one of the &#39;twoStepProcedures&#39;. The selection must then be passed back to finAPI as &#39;twoStepProcedureId&#39;;&lt;br/&gt;&amp;bull; REDIRECT_REQUIRED - means that the bank has sent a &#39;redirectUrl&#39;, to which the user must be forwarded. The bank&#39;s callback must then be passed back to finAPI as &#39;redirectCallback&#39;. | [default to null]
**ChallengeMessage** | **string** | In case of status &#x3D; CHALLENGE_RESPONSE_REQUIRED, this field contains a message from the bank containing instructions for the user on how to proceed with the authorization. | [optional] [default to null]
**AnswerFieldLabel** | **string** | Suggestion from the bank on how you can label your input field where the user should enter his challenge response. | [optional] [default to null]
**RedirectUrl** | **string** | In case of status &#x3D; REDIRECT_REQUIRED, this field contains the URL to which you must direct the user. It already includes the redirect URL back to your client that you have passed when initiating the service call. | [optional] [default to null]
**RedirectContext** | **string** | Set in case of status &#x3D; REDIRECT_REQUIRED. When the bank redirects the user back to your client, the redirect URL will contain this string, which you must process to identify the user context for the callback on your side. | [optional] [default to null]
**RedirectContextField** | **string** | Set in case of status &#x3D; REDIRECT_REQUIRED. This field is set to the name of the query parameter that contains the &#39;redirectContext&#39; in the redirect URL from the bank back to your client. | [optional] [default to null]
**TwoStepProcedures** | [**[]TwoStepProcedure**](TwoStepProcedure.md) | In case of status &#x3D; TWO_STEP_PROCEDURE_REQUIRED, this field contains the available two-step procedures. Note that this set does not necessarily match the set that is stored in the respective bank connection interface. You should always use the set from this field for the multi-step authentication flow. | [optional] [default to null]
**OpticalData** | **string** | In case that the bank server has instructed the user to scan a flicker code, then this field will contain the raw data for the flicker animation as a BASE-64 string. | [optional] [default to null]
**PhotoTanMimeType** | **string** | In case that the &#39;photoTanData&#39; field is set (i.e. not null), this field contains the MIME type to use for interpreting the photo data (e.g.: &#39;image/png&#39;) | [optional] [default to null]
**PhotoTanData** | **string** | In case that the bank server has instructed the user to scan a photo (or more generally speaking, any kind of QR-code-like data), then this field will contain the raw data of the photo as a BASE-64 string.  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


