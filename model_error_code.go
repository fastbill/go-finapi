/*
 * finAPI Access (with deprecation)
 *
 * <strong>RESTful API for Account Information Services (AIS) and Payment Initiation Services (PIS)</strong>  The following pages give you some general information on how to use our APIs.<br/> The actual API services documentation then follows further below. You can use the menu to jump between API sections. <br/> <br/> This page has a built-in HTTP(S) client, so you can test the services directly from within this page, by filling in the request parameters and/or body in the respective services, and then hitting the TRY button. Note that you need to be authorized to make a successful API call. To authorize, refer to the 'Authorization' section of the API, or just use the OAUTH button that can be found near the TRY button. <br/>  <h2 id=\"general-information\">General information</h2>  <h3 id=\"general-error-responses\"><strong>Error Responses</strong></h3> When an API call returns with an error, then in general it has the structure shown in the following example:  <pre> {   \"errors\": [     {       \"message\": \"Interface 'FINTS_SERVER' is not supported for this operation.\",       \"code\": \"BAD_REQUEST\",       \"type\": \"TECHNICAL\"     }   ],   \"date\": \"2020-11-19 16:54:06.854\",   \"requestId\": \"selfgen-312042e7-df55-47e4-bffd-956a68ef37b5\",   \"endpoint\": \"POST /api/v1/bankConnections/import\",   \"authContext\": \"1/21\",   \"bank\": \"DEMO0002 - finAPI Test Redirect Bank\" } </pre>  If an API call requires an additional authentication by the user, HTTP code 510 is returned and the error response contains the additional \"multiStepAuthentication\" object, see the following example:  <pre> {   \"errors\": [     {       \"message\": \"Es ist eine zusätzliche Authentifizierung erforderlich. Bitte geben Sie folgenden Code an: 123456\",       \"code\": \"ADDITIONAL_AUTHENTICATION_REQUIRED\",       \"type\": \"BUSINESS\",       \"multiStepAuthentication\": {         \"hash\": \"678b13f4be9ed7d981a840af8131223a\",         \"status\": \"CHALLENGE_RESPONSE_REQUIRED\",         \"challengeMessage\": \"Es ist eine zusätzliche Authentifizierung erforderlich. Bitte geben Sie folgenden Code an: 123456\",         \"answerFieldLabel\": \"TAN\",         \"redirectUrl\": null,         \"redirectContext\": null,         \"redirectContextField\": null,         \"twoStepProcedures\": null,         \"photoTanMimeType\": null,         \"photoTanData\": null,         \"opticalData\": null       }     }   ],   \"date\": \"2019-11-29 09:51:55.931\",   \"requestId\": \"selfgen-45059c99-1b14-4df7-9bd3-9d5f126df294\",   \"endpoint\": \"POST /api/v1/bankConnections/import\",   \"authContext\": \"1/18\",   \"bank\": \"DEMO0001 - finAPI Test Bank\" } </pre>  An exception to this error format are API authentication errors, where the following structure is returned:  <pre> {   \"error\": \"invalid_token\",   \"error_description\": \"Invalid access token: cccbce46-xxxx-xxxx-xxxx-xxxxxxxxxx\" } </pre>  <h3 id=\"general-paging\"><strong>Paging</strong></h3> API services that may potentially return a lot of data implement paging. They return a limited number of entries within a \"page\". Further entries must be fetched with subsequent calls. <br/><br/> Any API service that implements paging provides the following input parameters:<br/> &bull; \"page\": the number of the page to be retrieved (starting with 1).<br/> &bull; \"perPage\": the number of entries within a page. The default and maximum value is stated in the documentation of the respective services.  A paged response contains an additional \"paging\" object with the following structure:  <pre> {   ...   ,   \"paging\": {     \"page\": 1,     \"perPage\": 20,     \"pageCount\": 234,     \"totalCount\": 4662   } } </pre>  <h3 id=\"general-internationalization\"><strong>Internationalization</strong></h3> The finAPI services support internationalization which means you can define the language you prefer for API service responses. <br/><br/> The following languages are available: German, English, Czech, Slovak. <br/><br/> The preferred language can be defined by providing the official HTTP <strong>Accept-Language</strong> header. <br/><br/> finAPI reacts on the official iso language codes &quot;de&quot;, &quot;en&quot;, &quot;cs&quot; and &quot;sk&quot; for the named languages. Additional subtags supported by the Accept-Language header may be provided, e.g. &quot;en-US&quot;, but are ignored. <br/> If no Accept-Language header is given, German is used as the default language. <br/><br/> Exceptions:<br/> &bull; Bank login hints and login fields are only available in the language of the bank and not being translated.<br/> &bull; Direct messages from the bank systems typically returned as BUSINESS errors will not be translated.<br/> &bull; BUSINESS errors created by finAPI directly are available in German and English.<br/> &bull; TECHNICAL errors messages meant for developers are mostly in English, but also may be translated.  <h3 id=\"general-request-ids\"><strong>Request IDs</strong></h3> With any API call, you can pass a request ID via a header with name \"X-Request-Id\". The request ID can be an arbitrary string with up to 255 characters. Passing a longer string will result in an error. <br/><br/> If you don't pass a request ID for a call, finAPI will generate a random ID internally. <br/><br/> The request ID is always returned back in the response of a service, as a header with name \"X-Request-Id\". <br/><br/> We highly recommend to always pass a (preferably unique) request ID, and include it into your client application logs whenever you make a request or receive a response (especially in the case of an error response). finAPI is also logging request IDs on its end. Having a request ID can help the finAPI support team to work more efficiently and solve tickets faster.  <h3 id=\"general-overriding-http-methods\"><strong>Overriding HTTP methods</strong></h3> Some HTTP clients do not support the HTTP methods PATCH or DELETE. If you are using such a client in your application, you can use a POST request instead with a special HTTP header indicating the originally intended HTTP method. <br/><br/> The header's name is <strong>X-HTTP-Method-Override</strong>. Set its value to either <strong>PATCH</strong> or <strong>DELETE</strong>. POST Requests having this header set will be treated either as PATCH or DELETE by the finAPI servers. <br/><br/> Example: <br/><br/> <strong>X-HTTP-Method-Override: PATCH</strong><br/> POST /api/v1/label/51<br/> {\"name\": \"changed label\"}<br/><br/> will be interpreted by finAPI as:<br/><br/> PATCH /api/v1/label/51<br/> {\"name\": \"changed label\"}<br/>  <h3 id=\"general-user-metadata\"><strong>User metadata</strong></h3> With the migration to PSD2 APIs, a new term called \"User metadata\" (also known as \"PSU metadata\") has been introduced to the API. This user metadata aims to inform the banking API if there was a real end-user behind an HTTP request or if the request was triggered by a system (e.g. by an automatic batch update). In the latter case, the bank may apply some restrictions such as limiting the number of HTTP requests for a single consent. Also, some operations may be forbidden entirely by the banking API. For example, some banks do not allow issuing a new consent without the end-user being involved. Therefore, it is certainly necessary and obligatory for the customer to provide the PSU metadata for such operations. <br/><br/> As finAPI does not have direct interaction with the end-user, it is the client application's responsibility to provide all the necessary information about the end-user. This must be done by sending additional headers with every request triggered on behalf of the end-user. <br/><br/> At the moment, the following headers are supported by the API:<br/> &bull; \"PSU-IP-Address\" - the IP address of the user's device.<br/> &bull; \"PSU-Device-OS\" - the user's device and/or operating system identification.<br/> &bull; \"PSU-User-Agent\" - the user's web browser or other client device identification.  <h3 id=\"general-faq\"><strong>FAQ</strong></h3> <strong>Is there a finAPI SDK?</strong> <br/> Currently we do not offer a native SDK, but there is the option to generate an SDK for almost any target language via OpenAPI. Use the 'Download SDK' button on this page for SDK generation. <br/> <br/> <strong>How can I enable finAPI's automatic batch update?</strong> <br/> Currently there is no way to set up the batch update via the API. Please contact support@finapi.io for this. <br/> <br/> <strong>Why do I need to keep authorizing when calling services on this page?</strong> <br/> This page is a \"one-page-app\". Reloading the page resets the OAuth authorization context. There is generally no need to reload the page, so just don't do it and your authorization will persist. 
 *
 * API version: 1.151.1
 * Contact: kontakt@finapi.io
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package finapi

import (
	"encoding/json"
	"fmt"
)

// ErrorCode the model 'ErrorCode'
type ErrorCode string

// List of ErrorCode
const (
	ERRORCODE_ADDITIONAL_AUTHENTICATION_REQUIRED ErrorCode = "ADDITIONAL_AUTHENTICATION_REQUIRED"
	ERRORCODE_BAD_REQUEST ErrorCode = "BAD_REQUEST"
	ERRORCODE_BANK_SERVER_REJECTION ErrorCode = "BANK_SERVER_REJECTION"
	ERRORCODE_COLLECTIVE_MONEY_TRANSFER_NOT_SUPPORTED ErrorCode = "COLLECTIVE_MONEY_TRANSFER_NOT_SUPPORTED"
	ERRORCODE_ENTITY_EXISTS ErrorCode = "ENTITY_EXISTS"
	ERRORCODE_FORBIDDEN ErrorCode = "FORBIDDEN"
	ERRORCODE_IBAN_ONLY_DIRECT_DEBIT_NOT_SUPPORTED ErrorCode = "IBAN_ONLY_DIRECT_DEBIT_NOT_SUPPORTED"
	ERRORCODE_IBAN_ONLY_MONEY_TRANSFER_NOT_SUPPORTED ErrorCode = "IBAN_ONLY_MONEY_TRANSFER_NOT_SUPPORTED"
	ERRORCODE_ILLEGAL_ENTITY_STATE ErrorCode = "ILLEGAL_ENTITY_STATE"
	ERRORCODE_ILLEGAL_FIELD_VALUE ErrorCode = "ILLEGAL_FIELD_VALUE"
	ERRORCODE_ILLEGAL_PAGE_SIZE ErrorCode = "ILLEGAL_PAGE_SIZE"
	ERRORCODE_INVALID_CONSENT ErrorCode = "INVALID_CONSENT"
	ERRORCODE_INVALID_FILTER_OPTIONS ErrorCode = "INVALID_FILTER_OPTIONS"
	ERRORCODE_INVALID_TOKEN ErrorCode = "INVALID_TOKEN"
	ERRORCODE_INVALID_TWO_STEP_PROCEDURE ErrorCode = "INVALID_TWO_STEP_PROCEDURE"
	ERRORCODE_LOCKED ErrorCode = "LOCKED"
	ERRORCODE_LOGIN_FIELDS_MISMATCH ErrorCode = "LOGIN_FIELDS_MISMATCH"
	ERRORCODE_METHOD_NOT_ALLOWED ErrorCode = "METHOD_NOT_ALLOWED"
	ERRORCODE_MISSING_CREDENTIALS ErrorCode = "MISSING_CREDENTIALS"
	ERRORCODE_MISSING_FIELD ErrorCode = "MISSING_FIELD"
	ERRORCODE_MISSING_TWO_STEP_PROCEDURE ErrorCode = "MISSING_TWO_STEP_PROCEDURE"
	ERRORCODE_NO_ACCOUNTS_FOR_TYPE_LIST ErrorCode = "NO_ACCOUNTS_FOR_TYPE_LIST"
	ERRORCODE_NO_CERTIFICATE ErrorCode = "NO_CERTIFICATE"
	ERRORCODE_NO_EXISTING_CHALLENGE ErrorCode = "NO_EXISTING_CHALLENGE"
	ERRORCODE_NO_TPP_CLIENT_CREDENTIALS ErrorCode = "NO_TPP_CLIENT_CREDENTIALS"
	ERRORCODE_NO_PSU_METADATA ErrorCode = "NO_PSU_METADATA"
	ERRORCODE_TOO_MANY_IDS ErrorCode = "TOO_MANY_IDS"
	ERRORCODE_UNAUTHORIZED_ACCESS ErrorCode = "UNAUTHORIZED_ACCESS"
	ERRORCODE_UNEXPECTED_ERROR ErrorCode = "UNEXPECTED_ERROR"
	ERRORCODE_UNKNOWN_ENTITY ErrorCode = "UNKNOWN_ENTITY"
	ERRORCODE_UNSUPPORTED_BANK ErrorCode = "UNSUPPORTED_BANK"
	ERRORCODE_UNSUPPORTED_FEATURE ErrorCode = "UNSUPPORTED_FEATURE"
	ERRORCODE_UNSUPPORTED_MEDIA_TYPE ErrorCode = "UNSUPPORTED_MEDIA_TYPE"
	ERRORCODE_UNSUPPORTED_ORDER ErrorCode = "UNSUPPORTED_ORDER"
	ERRORCODE_USER_ALREADY_VERIFIED ErrorCode = "USER_ALREADY_VERIFIED"
	ERRORCODE_USER_NOT_VERIFIED ErrorCode = "USER_NOT_VERIFIED"
	ERRORCODE_WEB_FORM_ABORTED ErrorCode = "WEB_FORM_ABORTED"
	ERRORCODE_WEB_FORM_REQUIRED ErrorCode = "WEB_FORM_REQUIRED"
)

func (v *ErrorCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ErrorCode(value)
	for _, existing := range []ErrorCode{ "ADDITIONAL_AUTHENTICATION_REQUIRED", "BAD_REQUEST", "BANK_SERVER_REJECTION", "COLLECTIVE_MONEY_TRANSFER_NOT_SUPPORTED", "ENTITY_EXISTS", "FORBIDDEN", "IBAN_ONLY_DIRECT_DEBIT_NOT_SUPPORTED", "IBAN_ONLY_MONEY_TRANSFER_NOT_SUPPORTED", "ILLEGAL_ENTITY_STATE", "ILLEGAL_FIELD_VALUE", "ILLEGAL_PAGE_SIZE", "INVALID_CONSENT", "INVALID_FILTER_OPTIONS", "INVALID_TOKEN", "INVALID_TWO_STEP_PROCEDURE", "LOCKED", "LOGIN_FIELDS_MISMATCH", "METHOD_NOT_ALLOWED", "MISSING_CREDENTIALS", "MISSING_FIELD", "MISSING_TWO_STEP_PROCEDURE", "NO_ACCOUNTS_FOR_TYPE_LIST", "NO_CERTIFICATE", "NO_EXISTING_CHALLENGE", "NO_TPP_CLIENT_CREDENTIALS", "NO_PSU_METADATA", "TOO_MANY_IDS", "UNAUTHORIZED_ACCESS", "UNEXPECTED_ERROR", "UNKNOWN_ENTITY", "UNSUPPORTED_BANK", "UNSUPPORTED_FEATURE", "UNSUPPORTED_MEDIA_TYPE", "UNSUPPORTED_ORDER", "USER_ALREADY_VERIFIED", "USER_NOT_VERIFIED", "WEB_FORM_ABORTED", "WEB_FORM_REQUIRED",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ErrorCode", value)
}

// Ptr returns reference to ErrorCode value
func (v ErrorCode) Ptr() *ErrorCode {
	return &v
}

type NullableErrorCode struct {
	value *ErrorCode
	isSet bool
}

func (v NullableErrorCode) Get() *ErrorCode {
	return v.value
}

func (v *NullableErrorCode) Set(val *ErrorCode) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCode) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCode(val *ErrorCode) *NullableErrorCode {
	return &NullableErrorCode{value: val, isSet: true}
}

func (v NullableErrorCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

