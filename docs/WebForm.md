# WebForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Web Form identifier, as returned in the 451 response of the REST service call that initiated the Web Form flow. | 
**Token** | **string** | Token for the finAPI Web Form page, as contained in the 451 response of the REST service call that initiated the Web Form flow (in the &#39;Location&#39; header). | 
**Status** | [**WebFormStatus**](WebFormStatus.md) | &lt;strong&gt;Type:&lt;/strong&gt; WebFormStatus&lt;br/&gt; Status of a Web Form. Possible values are:&lt;br/&gt;&amp;bull; NOT_YET_OPENED - the Web Form URL was not yet called;&lt;br/&gt;&amp;bull; IN_PROGRESS - the Web Form has been opened but not yet submitted by the user;&lt;br/&gt;&amp;bull; COMPLETED - the user has opened and submitted the Web Form;&lt;br/&gt;&amp;bull; ABORTED - the user has opened but then aborted the Web Form, or the Web Form was aborted by the finAPI system because it has expired (this is the case when a Web Form is opened and then not submitted within 10 minutes) | 
**ServiceResponseCode** | **NullableInt32** | HTTP response code of the REST service call that initiated the Web Form flow. This field can be queried as soon as the status becomes COMPLETED or ABORTED. Note that it is still not guaranteed in this case that the field has a value, i.e. it might be null. | 
**ServiceResponseBody** | **NullableString** | HTTP response body of the REST service call that initiated the Web Form flow. This field can be queried as soon as the status becomes COMPLETED or ABORTED. Note that it is still not guaranteed in this case that the field has a value, i.e. it might be null. | 

## Methods

### NewWebForm

`func NewWebForm(id int64, token string, status WebFormStatus, serviceResponseCode NullableInt32, serviceResponseBody NullableString, ) *WebForm`

NewWebForm instantiates a new WebForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebFormWithDefaults

`func NewWebFormWithDefaults() *WebForm`

NewWebFormWithDefaults instantiates a new WebForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebForm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebForm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebForm) SetId(v int64)`

SetId sets Id field to given value.


### GetToken

`func (o *WebForm) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *WebForm) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *WebForm) SetToken(v string)`

SetToken sets Token field to given value.


### GetStatus

`func (o *WebForm) GetStatus() WebFormStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebForm) GetStatusOk() (*WebFormStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebForm) SetStatus(v WebFormStatus)`

SetStatus sets Status field to given value.


### GetServiceResponseCode

`func (o *WebForm) GetServiceResponseCode() int32`

GetServiceResponseCode returns the ServiceResponseCode field if non-nil, zero value otherwise.

### GetServiceResponseCodeOk

`func (o *WebForm) GetServiceResponseCodeOk() (*int32, bool)`

GetServiceResponseCodeOk returns a tuple with the ServiceResponseCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceResponseCode

`func (o *WebForm) SetServiceResponseCode(v int32)`

SetServiceResponseCode sets ServiceResponseCode field to given value.


### SetServiceResponseCodeNil

`func (o *WebForm) SetServiceResponseCodeNil(b bool)`

 SetServiceResponseCodeNil sets the value for ServiceResponseCode to be an explicit nil

### UnsetServiceResponseCode
`func (o *WebForm) UnsetServiceResponseCode()`

UnsetServiceResponseCode ensures that no value is present for ServiceResponseCode, not even an explicit nil
### GetServiceResponseBody

`func (o *WebForm) GetServiceResponseBody() string`

GetServiceResponseBody returns the ServiceResponseBody field if non-nil, zero value otherwise.

### GetServiceResponseBodyOk

`func (o *WebForm) GetServiceResponseBodyOk() (*string, bool)`

GetServiceResponseBodyOk returns a tuple with the ServiceResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceResponseBody

`func (o *WebForm) SetServiceResponseBody(v string)`

SetServiceResponseBody sets ServiceResponseBody field to given value.


### SetServiceResponseBodyNil

`func (o *WebForm) SetServiceResponseBodyNil(b bool)`

 SetServiceResponseBodyNil sets the value for ServiceResponseBody to be an explicit nil

### UnsetServiceResponseBody
`func (o *WebForm) UnsetServiceResponseBody()`

UnsetServiceResponseBody ensures that no value is present for ServiceResponseBody, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


