# MultiStepAuthenticationChallenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hash** | **string** | Hash for this multi-step authentication flow. Must be passed back to finAPI when continuing the flow. | 
**Status** | [**MsaStatus**](MsaStatus.md) | &lt;strong&gt;Type:&lt;/strong&gt; MsaStatus&lt;br/&gt; Indicates the current status of the multi-step authentication flow:&lt;br/&gt;&lt;br/&gt;TWO_STEP_PROCEDURE_REQUIRED means that the bank has requested an SCA method selection for the user. In this case, the service should be recalled with a chosen TSP-ID set to the &#39;twoStepProcedureId&#39; field.&lt;br/&gt;&lt;br/&gt;CHALLENGE_RESPONSE_REQUIRED means that the bank has requested a challenge code for the previously given TSP (SCA). This status can be completed by setting the &#39;challengeResponse&#39; field.&lt;br/&gt;&lt;br/&gt;REDIRECT_REQUIRED means that the user must be redirected to the bank&#39;s website, where the authentication can be finished.&lt;br/&gt;&lt;br/&gt;DECOUPLED_AUTH_REQUIRED means that the bank has asked for the decoupled authentication. In this case, the &#39;decoupledCallback&#39; field must be set to true to complete the authentication.&lt;br/&gt;&lt;br/&gt;DECOUPLED_AUTH_IN_PROGRESS means that the bank is waiting for the completion of the decoupled authentication by the user. Until this is done, the service should be recalled at most every 5 seconds with the &#39;decoupledCallback&#39; field set to &#39;true&#39;. Once the decoupled authentication is completed by the user, the service returns a successful response. | 
**ChallengeMessage** | **NullableString** | In case of status &#x3D; CHALLENGE_RESPONSE_REQUIRED, this field contains a message from the bank containing instructions for the user on how to proceed with the authorization. | 
**AnswerFieldLabel** | **NullableString** | Suggestion from the bank on how you can label your input field where the user should enter his challenge response. | 
**RedirectUrl** | **NullableString** | In case of status &#x3D; REDIRECT_REQUIRED, this field contains the URL to which you must direct the user. It already includes the redirect URL back to your client that you have passed when initiating the service call. | 
**RedirectContext** | **NullableString** | Set in case of status &#x3D; REDIRECT_REQUIRED. When the bank redirects the user back to your client, the redirect URL will contain this string, which you must process to identify the user context for the callback on your side. | 
**RedirectContextField** | **NullableString** | Set in case of status &#x3D; REDIRECT_REQUIRED. This field is set to the name of the query parameter that contains the &#39;redirectContext&#39; in the redirect URL from the bank back to your client. | 
**TwoStepProcedures** | [**[]TwoStepProcedure**](TwoStepProcedure.md) | &lt;strong&gt;Type:&lt;/strong&gt; TwoStepProcedure&lt;br/&gt; In case of status &#x3D; TWO_STEP_PROCEDURE_REQUIRED, this field contains the available two-step procedures. Note that this set does not necessarily match the set that is stored in the respective bank connection interface. You should always use the set from this field for the multi-step authentication flow. | 
**PhotoTanMimeType** | **NullableString** | In case that the &#39;photoTanData&#39; field is set (i.e. not null), this field contains the MIME type to use for interpreting the photo data (e.g.: &#39;image/png&#39;) | 
**PhotoTanData** | **NullableString** | In case that the bank server has instructed the user to scan a photo (or more generally speaking, any kind of QR-code-like data), then this field will contain the raw data of the photo as a BASE-64 string.  | 
**OpticalData** | **NullableString** | In case that the bank server has instructed the user to scan a flicker code, then this field will contain the raw data for the flicker animation as a BASE-64 string. | 

## Methods

### NewMultiStepAuthenticationChallenge

`func NewMultiStepAuthenticationChallenge(hash string, status MsaStatus, challengeMessage NullableString, answerFieldLabel NullableString, redirectUrl NullableString, redirectContext NullableString, redirectContextField NullableString, twoStepProcedures []TwoStepProcedure, photoTanMimeType NullableString, photoTanData NullableString, opticalData NullableString, ) *MultiStepAuthenticationChallenge`

NewMultiStepAuthenticationChallenge instantiates a new MultiStepAuthenticationChallenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultiStepAuthenticationChallengeWithDefaults

`func NewMultiStepAuthenticationChallengeWithDefaults() *MultiStepAuthenticationChallenge`

NewMultiStepAuthenticationChallengeWithDefaults instantiates a new MultiStepAuthenticationChallenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHash

`func (o *MultiStepAuthenticationChallenge) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *MultiStepAuthenticationChallenge) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *MultiStepAuthenticationChallenge) SetHash(v string)`

SetHash sets Hash field to given value.


### GetStatus

`func (o *MultiStepAuthenticationChallenge) GetStatus() MsaStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MultiStepAuthenticationChallenge) GetStatusOk() (*MsaStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MultiStepAuthenticationChallenge) SetStatus(v MsaStatus)`

SetStatus sets Status field to given value.


### GetChallengeMessage

`func (o *MultiStepAuthenticationChallenge) GetChallengeMessage() string`

GetChallengeMessage returns the ChallengeMessage field if non-nil, zero value otherwise.

### GetChallengeMessageOk

`func (o *MultiStepAuthenticationChallenge) GetChallengeMessageOk() (*string, bool)`

GetChallengeMessageOk returns a tuple with the ChallengeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeMessage

`func (o *MultiStepAuthenticationChallenge) SetChallengeMessage(v string)`

SetChallengeMessage sets ChallengeMessage field to given value.


### SetChallengeMessageNil

`func (o *MultiStepAuthenticationChallenge) SetChallengeMessageNil(b bool)`

 SetChallengeMessageNil sets the value for ChallengeMessage to be an explicit nil

### UnsetChallengeMessage
`func (o *MultiStepAuthenticationChallenge) UnsetChallengeMessage()`

UnsetChallengeMessage ensures that no value is present for ChallengeMessage, not even an explicit nil
### GetAnswerFieldLabel

`func (o *MultiStepAuthenticationChallenge) GetAnswerFieldLabel() string`

GetAnswerFieldLabel returns the AnswerFieldLabel field if non-nil, zero value otherwise.

### GetAnswerFieldLabelOk

`func (o *MultiStepAuthenticationChallenge) GetAnswerFieldLabelOk() (*string, bool)`

GetAnswerFieldLabelOk returns a tuple with the AnswerFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerFieldLabel

`func (o *MultiStepAuthenticationChallenge) SetAnswerFieldLabel(v string)`

SetAnswerFieldLabel sets AnswerFieldLabel field to given value.


### SetAnswerFieldLabelNil

`func (o *MultiStepAuthenticationChallenge) SetAnswerFieldLabelNil(b bool)`

 SetAnswerFieldLabelNil sets the value for AnswerFieldLabel to be an explicit nil

### UnsetAnswerFieldLabel
`func (o *MultiStepAuthenticationChallenge) UnsetAnswerFieldLabel()`

UnsetAnswerFieldLabel ensures that no value is present for AnswerFieldLabel, not even an explicit nil
### GetRedirectUrl

`func (o *MultiStepAuthenticationChallenge) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *MultiStepAuthenticationChallenge) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *MultiStepAuthenticationChallenge) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### SetRedirectUrlNil

`func (o *MultiStepAuthenticationChallenge) SetRedirectUrlNil(b bool)`

 SetRedirectUrlNil sets the value for RedirectUrl to be an explicit nil

### UnsetRedirectUrl
`func (o *MultiStepAuthenticationChallenge) UnsetRedirectUrl()`

UnsetRedirectUrl ensures that no value is present for RedirectUrl, not even an explicit nil
### GetRedirectContext

`func (o *MultiStepAuthenticationChallenge) GetRedirectContext() string`

GetRedirectContext returns the RedirectContext field if non-nil, zero value otherwise.

### GetRedirectContextOk

`func (o *MultiStepAuthenticationChallenge) GetRedirectContextOk() (*string, bool)`

GetRedirectContextOk returns a tuple with the RedirectContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectContext

`func (o *MultiStepAuthenticationChallenge) SetRedirectContext(v string)`

SetRedirectContext sets RedirectContext field to given value.


### SetRedirectContextNil

`func (o *MultiStepAuthenticationChallenge) SetRedirectContextNil(b bool)`

 SetRedirectContextNil sets the value for RedirectContext to be an explicit nil

### UnsetRedirectContext
`func (o *MultiStepAuthenticationChallenge) UnsetRedirectContext()`

UnsetRedirectContext ensures that no value is present for RedirectContext, not even an explicit nil
### GetRedirectContextField

`func (o *MultiStepAuthenticationChallenge) GetRedirectContextField() string`

GetRedirectContextField returns the RedirectContextField field if non-nil, zero value otherwise.

### GetRedirectContextFieldOk

`func (o *MultiStepAuthenticationChallenge) GetRedirectContextFieldOk() (*string, bool)`

GetRedirectContextFieldOk returns a tuple with the RedirectContextField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectContextField

`func (o *MultiStepAuthenticationChallenge) SetRedirectContextField(v string)`

SetRedirectContextField sets RedirectContextField field to given value.


### SetRedirectContextFieldNil

`func (o *MultiStepAuthenticationChallenge) SetRedirectContextFieldNil(b bool)`

 SetRedirectContextFieldNil sets the value for RedirectContextField to be an explicit nil

### UnsetRedirectContextField
`func (o *MultiStepAuthenticationChallenge) UnsetRedirectContextField()`

UnsetRedirectContextField ensures that no value is present for RedirectContextField, not even an explicit nil
### GetTwoStepProcedures

`func (o *MultiStepAuthenticationChallenge) GetTwoStepProcedures() []TwoStepProcedure`

GetTwoStepProcedures returns the TwoStepProcedures field if non-nil, zero value otherwise.

### GetTwoStepProceduresOk

`func (o *MultiStepAuthenticationChallenge) GetTwoStepProceduresOk() (*[]TwoStepProcedure, bool)`

GetTwoStepProceduresOk returns a tuple with the TwoStepProcedures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoStepProcedures

`func (o *MultiStepAuthenticationChallenge) SetTwoStepProcedures(v []TwoStepProcedure)`

SetTwoStepProcedures sets TwoStepProcedures field to given value.


### SetTwoStepProceduresNil

`func (o *MultiStepAuthenticationChallenge) SetTwoStepProceduresNil(b bool)`

 SetTwoStepProceduresNil sets the value for TwoStepProcedures to be an explicit nil

### UnsetTwoStepProcedures
`func (o *MultiStepAuthenticationChallenge) UnsetTwoStepProcedures()`

UnsetTwoStepProcedures ensures that no value is present for TwoStepProcedures, not even an explicit nil
### GetPhotoTanMimeType

`func (o *MultiStepAuthenticationChallenge) GetPhotoTanMimeType() string`

GetPhotoTanMimeType returns the PhotoTanMimeType field if non-nil, zero value otherwise.

### GetPhotoTanMimeTypeOk

`func (o *MultiStepAuthenticationChallenge) GetPhotoTanMimeTypeOk() (*string, bool)`

GetPhotoTanMimeTypeOk returns a tuple with the PhotoTanMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoTanMimeType

`func (o *MultiStepAuthenticationChallenge) SetPhotoTanMimeType(v string)`

SetPhotoTanMimeType sets PhotoTanMimeType field to given value.


### SetPhotoTanMimeTypeNil

`func (o *MultiStepAuthenticationChallenge) SetPhotoTanMimeTypeNil(b bool)`

 SetPhotoTanMimeTypeNil sets the value for PhotoTanMimeType to be an explicit nil

### UnsetPhotoTanMimeType
`func (o *MultiStepAuthenticationChallenge) UnsetPhotoTanMimeType()`

UnsetPhotoTanMimeType ensures that no value is present for PhotoTanMimeType, not even an explicit nil
### GetPhotoTanData

`func (o *MultiStepAuthenticationChallenge) GetPhotoTanData() string`

GetPhotoTanData returns the PhotoTanData field if non-nil, zero value otherwise.

### GetPhotoTanDataOk

`func (o *MultiStepAuthenticationChallenge) GetPhotoTanDataOk() (*string, bool)`

GetPhotoTanDataOk returns a tuple with the PhotoTanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoTanData

`func (o *MultiStepAuthenticationChallenge) SetPhotoTanData(v string)`

SetPhotoTanData sets PhotoTanData field to given value.


### SetPhotoTanDataNil

`func (o *MultiStepAuthenticationChallenge) SetPhotoTanDataNil(b bool)`

 SetPhotoTanDataNil sets the value for PhotoTanData to be an explicit nil

### UnsetPhotoTanData
`func (o *MultiStepAuthenticationChallenge) UnsetPhotoTanData()`

UnsetPhotoTanData ensures that no value is present for PhotoTanData, not even an explicit nil
### GetOpticalData

`func (o *MultiStepAuthenticationChallenge) GetOpticalData() string`

GetOpticalData returns the OpticalData field if non-nil, zero value otherwise.

### GetOpticalDataOk

`func (o *MultiStepAuthenticationChallenge) GetOpticalDataOk() (*string, bool)`

GetOpticalDataOk returns a tuple with the OpticalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpticalData

`func (o *MultiStepAuthenticationChallenge) SetOpticalData(v string)`

SetOpticalData sets OpticalData field to given value.


### SetOpticalDataNil

`func (o *MultiStepAuthenticationChallenge) SetOpticalDataNil(b bool)`

 SetOpticalDataNil sets the value for OpticalData to be an explicit nil

### UnsetOpticalData
`func (o *MultiStepAuthenticationChallenge) UnsetOpticalData()`

UnsetOpticalData ensures that no value is present for OpticalData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


