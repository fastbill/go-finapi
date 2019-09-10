# WebForm

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | Web form identifier, as returned in the 451 response of the REST service call that initiated the web form flow. | [default to null]
**Token** | **string** | Token for the finAPI web form page, as contained in the 451 response of the REST service call that initiated the web form flow (in the &#39;Location&#39; header). | [default to null]
**Status** | **string** | Status of a web form. Possible values are:&lt;br/&gt;&amp;bull; NOT_YET_OPENED - the web form URL was not yet called;&lt;br/&gt;&amp;bull; IN_PROGRESS - the web form has been opened but not yet submitted by the user;&lt;br/&gt;&amp;bull; COMPLETED - the user has opened and submitted the web form;&lt;br/&gt;&amp;bull; ABORTED - the user has opened but then aborted the web form, or the web form was aborted by the finAPI system because it has expired (this is the case when a web form is opened and then not submitted within 20 minutes) | [default to null]
**ServiceResponseCode** | **int32** | HTTP response code of the REST service call that initiated the web form flow. This field can be queried as soon as the status becomes COMPLETED or ABORTED. Note that it is still not guaranteed in this case that the field has a value, i.e. it might be null. | [optional] [default to null]
**ServiceResponseBody** | **string** | HTTP response body of the REST service call that initiated the web form flow. This field can be queried as soon as the status becomes COMPLETED or ABORTED. Note that it is still not guaranteed in this case that the field has a value, i.e. it might be null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


