# DirectDebitOrderingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuccessMessage** | **NullableString** | Technical message from the bank server, confirming the success of the request. Typically, you would not want to present this message to the user. Note that this field may not be set. However if it is not set, it does not necessarily mean that there was an error in processing the request. | 
**WarnMessage** | **NullableString** | In some cases, a bank server may accept the requested order, but return a warn message. This message may be of technical nature, but could also be of interest to the user. | 
**PaymentId** | **int64** | Payment identifier. Can be used to retrieve the status of the payment (see &#39;Get payments&#39; service). | 
**ChallengeMessage** | **NullableString** | Message from the bank server containing information or instructions on how to retrieve the TAN that is needed to execute the requested order. This message should be presented to the user. Note that some bank servers may limit the message to just the most crucial information, e.g. the message may contain just a single number that depicts the target TAN number on a user&#39;s TAN list. You may want to parse the challenge message for such cases and extend it with more detailed information before showing it to the user. | 
**AnswerFieldLabel** | **NullableString** | Suggestion from the bank server on how you can label your input field where the user must enter his TAN. A typical value that many bank servers give is &#39;TAN-Nummer&#39;. | 
**TanListNumber** | **NullableString** | In case that the bank server has instructed the user to look up a TAN from a TAN list, this field may contain the identification number of the TAN list. However in most cases, this field is only set (i.e. not null) when the user has multiple active TAN lists. | 
**OpticalData** | **NullableString** | In case that the bank server has instructed the user to scan a flicker code, then this field will contain the raw data for the flicker animation as a BASE-64 string. Otherwise, this field will be not set (i.e. null). For more information on how to process the flicker code data, please refer to &lt;a href&#x3D;&#39;https://documentation.finapi.io/access/Flicker-Code-Template.2807824454.html&#39; target&#x3D;&#39;_blank&#39;&gt;this&lt;/a&gt; article. | 
**PhotoTanMimeType** | **NullableString** | In case that the &#39;photoTanData&#39; field is set (i.e. not null), this field contains the MIME type to use for interpreting the photo data (e.g.: &#39;image/png&#39;) | 
**PhotoTanData** | **NullableString** | In case that the bank server has instructed the user to scan a photo (or more generally speaking, any kind of QR-code-like data), then this field will contain the raw data of the photo as a BASE-64 string. Otherwise, this field will be not set (i.e. null). | 

## Methods

### NewDirectDebitOrderingResponse

`func NewDirectDebitOrderingResponse(successMessage NullableString, warnMessage NullableString, paymentId int64, challengeMessage NullableString, answerFieldLabel NullableString, tanListNumber NullableString, opticalData NullableString, photoTanMimeType NullableString, photoTanData NullableString, ) *DirectDebitOrderingResponse`

NewDirectDebitOrderingResponse instantiates a new DirectDebitOrderingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectDebitOrderingResponseWithDefaults

`func NewDirectDebitOrderingResponseWithDefaults() *DirectDebitOrderingResponse`

NewDirectDebitOrderingResponseWithDefaults instantiates a new DirectDebitOrderingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccessMessage

`func (o *DirectDebitOrderingResponse) GetSuccessMessage() string`

GetSuccessMessage returns the SuccessMessage field if non-nil, zero value otherwise.

### GetSuccessMessageOk

`func (o *DirectDebitOrderingResponse) GetSuccessMessageOk() (*string, bool)`

GetSuccessMessageOk returns a tuple with the SuccessMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessMessage

`func (o *DirectDebitOrderingResponse) SetSuccessMessage(v string)`

SetSuccessMessage sets SuccessMessage field to given value.


### SetSuccessMessageNil

`func (o *DirectDebitOrderingResponse) SetSuccessMessageNil(b bool)`

 SetSuccessMessageNil sets the value for SuccessMessage to be an explicit nil

### UnsetSuccessMessage
`func (o *DirectDebitOrderingResponse) UnsetSuccessMessage()`

UnsetSuccessMessage ensures that no value is present for SuccessMessage, not even an explicit nil
### GetWarnMessage

`func (o *DirectDebitOrderingResponse) GetWarnMessage() string`

GetWarnMessage returns the WarnMessage field if non-nil, zero value otherwise.

### GetWarnMessageOk

`func (o *DirectDebitOrderingResponse) GetWarnMessageOk() (*string, bool)`

GetWarnMessageOk returns a tuple with the WarnMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnMessage

`func (o *DirectDebitOrderingResponse) SetWarnMessage(v string)`

SetWarnMessage sets WarnMessage field to given value.


### SetWarnMessageNil

`func (o *DirectDebitOrderingResponse) SetWarnMessageNil(b bool)`

 SetWarnMessageNil sets the value for WarnMessage to be an explicit nil

### UnsetWarnMessage
`func (o *DirectDebitOrderingResponse) UnsetWarnMessage()`

UnsetWarnMessage ensures that no value is present for WarnMessage, not even an explicit nil
### GetPaymentId

`func (o *DirectDebitOrderingResponse) GetPaymentId() int64`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *DirectDebitOrderingResponse) GetPaymentIdOk() (*int64, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *DirectDebitOrderingResponse) SetPaymentId(v int64)`

SetPaymentId sets PaymentId field to given value.


### GetChallengeMessage

`func (o *DirectDebitOrderingResponse) GetChallengeMessage() string`

GetChallengeMessage returns the ChallengeMessage field if non-nil, zero value otherwise.

### GetChallengeMessageOk

`func (o *DirectDebitOrderingResponse) GetChallengeMessageOk() (*string, bool)`

GetChallengeMessageOk returns a tuple with the ChallengeMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallengeMessage

`func (o *DirectDebitOrderingResponse) SetChallengeMessage(v string)`

SetChallengeMessage sets ChallengeMessage field to given value.


### SetChallengeMessageNil

`func (o *DirectDebitOrderingResponse) SetChallengeMessageNil(b bool)`

 SetChallengeMessageNil sets the value for ChallengeMessage to be an explicit nil

### UnsetChallengeMessage
`func (o *DirectDebitOrderingResponse) UnsetChallengeMessage()`

UnsetChallengeMessage ensures that no value is present for ChallengeMessage, not even an explicit nil
### GetAnswerFieldLabel

`func (o *DirectDebitOrderingResponse) GetAnswerFieldLabel() string`

GetAnswerFieldLabel returns the AnswerFieldLabel field if non-nil, zero value otherwise.

### GetAnswerFieldLabelOk

`func (o *DirectDebitOrderingResponse) GetAnswerFieldLabelOk() (*string, bool)`

GetAnswerFieldLabelOk returns a tuple with the AnswerFieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerFieldLabel

`func (o *DirectDebitOrderingResponse) SetAnswerFieldLabel(v string)`

SetAnswerFieldLabel sets AnswerFieldLabel field to given value.


### SetAnswerFieldLabelNil

`func (o *DirectDebitOrderingResponse) SetAnswerFieldLabelNil(b bool)`

 SetAnswerFieldLabelNil sets the value for AnswerFieldLabel to be an explicit nil

### UnsetAnswerFieldLabel
`func (o *DirectDebitOrderingResponse) UnsetAnswerFieldLabel()`

UnsetAnswerFieldLabel ensures that no value is present for AnswerFieldLabel, not even an explicit nil
### GetTanListNumber

`func (o *DirectDebitOrderingResponse) GetTanListNumber() string`

GetTanListNumber returns the TanListNumber field if non-nil, zero value otherwise.

### GetTanListNumberOk

`func (o *DirectDebitOrderingResponse) GetTanListNumberOk() (*string, bool)`

GetTanListNumberOk returns a tuple with the TanListNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTanListNumber

`func (o *DirectDebitOrderingResponse) SetTanListNumber(v string)`

SetTanListNumber sets TanListNumber field to given value.


### SetTanListNumberNil

`func (o *DirectDebitOrderingResponse) SetTanListNumberNil(b bool)`

 SetTanListNumberNil sets the value for TanListNumber to be an explicit nil

### UnsetTanListNumber
`func (o *DirectDebitOrderingResponse) UnsetTanListNumber()`

UnsetTanListNumber ensures that no value is present for TanListNumber, not even an explicit nil
### GetOpticalData

`func (o *DirectDebitOrderingResponse) GetOpticalData() string`

GetOpticalData returns the OpticalData field if non-nil, zero value otherwise.

### GetOpticalDataOk

`func (o *DirectDebitOrderingResponse) GetOpticalDataOk() (*string, bool)`

GetOpticalDataOk returns a tuple with the OpticalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpticalData

`func (o *DirectDebitOrderingResponse) SetOpticalData(v string)`

SetOpticalData sets OpticalData field to given value.


### SetOpticalDataNil

`func (o *DirectDebitOrderingResponse) SetOpticalDataNil(b bool)`

 SetOpticalDataNil sets the value for OpticalData to be an explicit nil

### UnsetOpticalData
`func (o *DirectDebitOrderingResponse) UnsetOpticalData()`

UnsetOpticalData ensures that no value is present for OpticalData, not even an explicit nil
### GetPhotoTanMimeType

`func (o *DirectDebitOrderingResponse) GetPhotoTanMimeType() string`

GetPhotoTanMimeType returns the PhotoTanMimeType field if non-nil, zero value otherwise.

### GetPhotoTanMimeTypeOk

`func (o *DirectDebitOrderingResponse) GetPhotoTanMimeTypeOk() (*string, bool)`

GetPhotoTanMimeTypeOk returns a tuple with the PhotoTanMimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoTanMimeType

`func (o *DirectDebitOrderingResponse) SetPhotoTanMimeType(v string)`

SetPhotoTanMimeType sets PhotoTanMimeType field to given value.


### SetPhotoTanMimeTypeNil

`func (o *DirectDebitOrderingResponse) SetPhotoTanMimeTypeNil(b bool)`

 SetPhotoTanMimeTypeNil sets the value for PhotoTanMimeType to be an explicit nil

### UnsetPhotoTanMimeType
`func (o *DirectDebitOrderingResponse) UnsetPhotoTanMimeType()`

UnsetPhotoTanMimeType ensures that no value is present for PhotoTanMimeType, not even an explicit nil
### GetPhotoTanData

`func (o *DirectDebitOrderingResponse) GetPhotoTanData() string`

GetPhotoTanData returns the PhotoTanData field if non-nil, zero value otherwise.

### GetPhotoTanDataOk

`func (o *DirectDebitOrderingResponse) GetPhotoTanDataOk() (*string, bool)`

GetPhotoTanDataOk returns a tuple with the PhotoTanData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoTanData

`func (o *DirectDebitOrderingResponse) SetPhotoTanData(v string)`

SetPhotoTanData sets PhotoTanData field to given value.


### SetPhotoTanDataNil

`func (o *DirectDebitOrderingResponse) SetPhotoTanDataNil(b bool)`

 SetPhotoTanDataNil sets the value for PhotoTanData to be an explicit nil

### UnsetPhotoTanData
`func (o *DirectDebitOrderingResponse) UnsetPhotoTanData()`

UnsetPhotoTanData ensures that no value is present for PhotoTanData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


