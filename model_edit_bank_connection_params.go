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
)

// EditBankConnectionParams Container for bank connection edit params
type EditBankConnectionParams struct {
	// New name for the bank connection. Maximum length is 64. If you do not want to change the current name let this field remain unset. If you want to clear the current name, set the field's value to an empty string (\"\").<br/><br/><strong>NOTE:</strong> If you are a Web Form 2.0 customer, and would like to update the name of your bank connection, please use the API parameter.
	Name *string `json:"name,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' and 'interface' fields instead. If any of those two fields is used, then the value of this field will be ignored.<br/><br/>New online banking user ID. If you do not want to change the current user ID let this field remain unset. In case you need to use finAPI's Web Form to let the user update the field, just set the field to any value, so that the service recognizes that you wish to use the Web Form flow. Note that you cannot clear the current user ID, i.e. a bank connection must always have a user ID (except for when it is a 'demo connection'). Max length: 170.
	BankingUserId *string `json:"bankingUserId,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' and 'interface' fields instead. If any of those two fields is used, then the value of this field will be ignored.<br/><br/>New online banking customer ID. If you do not want to change the current customer ID let this field remain unset. In case you need to use finAPI's Web Form to let the user update the field, just set the field to non-empty value, so that the service recognizes that you wish to use the Web Form flow. If you want to clear the current customer ID, set the field's value to an empty string (\"\"). Max length: 170.
	BankingCustomerId *string `json:"bankingCustomerId,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'loginCredentials' and 'interface' fields instead. If any of those two fields is used, then the value of this field will be ignored.<br/><br/>New online banking PIN. If you do not want to change the current PIN let this field remain unset. In case you need to use finAPI's Web Form to let the user update the field, just set the field to non-empty value, so that the service recognizes that you wish to use the Web Form flow. If you want to clear the current PIN, set the field's value to an empty string (\"\").<br/><br/>Any symbols are allowed. Max length: 170.
	BankingPin *string `json:"bankingPin,omitempty"`
	// <strong>Type:</strong> BankingInterface<br/> The interface for which you want to edit data. Must be given when you pass 'loginCredentials' and/or a 'defaultTwoStepProcedureId'.
	Interface *BankingInterface `json:"interface,omitempty"`
	// <strong>Type:</strong> LoginCredential<br/> Set of login credentials that you want to edit. Must be passed in combination with the 'interface' field. The labels that you pass must match with the login credential labels that the respective interface defines. If you want to clear the stored value for a credential, you can pass an empty string (\"\") as value.In case you need to use finAPI's Web Form to let the user update the login credentials, send all fields the user wishes to update with a non-empty value.In case all fields contain an empty string (\"\"), no Web Form will be generated. Note that any change in the credentials will automatically remove the saved consent data associated with those credentials.<br/><br/><strong>NOTE:</strong> If you are a  Web Form 2.0 customer, and would like to allow your end-users to change the credentials they have stored in our system, then please navigate <a target=\"_blank\" href='?product=web_form_2.0#post-/api/tasks/backgroundUpdate' target='_blank'>here</a> to implement the same functionality.
	LoginCredentials *[]LoginCredential `json:"loginCredentials,omitempty"`
	// NOTE: In the future, this field will work only in combination with the 'interface' field.<br/><br/>New default two-step-procedure. Must match the 'procedureId' of one of the procedures that are listed in the bank connection. If you do not want to change this field let it remain unset. If you want to clear the current default two-step-procedure, set the field's value to an empty string (\"\").<br/><br/><strong>NOTE:</strong> If you are a Web Form 2.0 customer and would like to allow your end users to update their preferred TAN procedure that is stored in our system, then please navigate <a target=\"_blank\" href='?product=web_form_2.0#post-/api/tasks/backgroundUpdate'>here</a> to implement the same functionality.
	DefaultTwoStepProcedureId *string `json:"defaultTwoStepProcedureId,omitempty"`
}

// NewEditBankConnectionParams instantiates a new EditBankConnectionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEditBankConnectionParams() *EditBankConnectionParams {
	this := EditBankConnectionParams{}
	return &this
}

// NewEditBankConnectionParamsWithDefaults instantiates a new EditBankConnectionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEditBankConnectionParamsWithDefaults() *EditBankConnectionParams {
	this := EditBankConnectionParams{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EditBankConnectionParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditBankConnectionParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EditBankConnectionParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EditBankConnectionParams) SetName(v string) {
	o.Name = &v
}

// GetBankingUserId returns the BankingUserId field value if set, zero value otherwise.
func (o *EditBankConnectionParams) GetBankingUserId() string {
	if o == nil || o.BankingUserId == nil {
		var ret string
		return ret
	}
	return *o.BankingUserId
}

// GetBankingUserIdOk returns a tuple with the BankingUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditBankConnectionParams) GetBankingUserIdOk() (*string, bool) {
	if o == nil || o.BankingUserId == nil {
		return nil, false
	}
	return o.BankingUserId, true
}

// HasBankingUserId returns a boolean if a field has been set.
func (o *EditBankConnectionParams) HasBankingUserId() bool {
	if o != nil && o.BankingUserId != nil {
		return true
	}

	return false
}

// SetBankingUserId gets a reference to the given string and assigns it to the BankingUserId field.
func (o *EditBankConnectionParams) SetBankingUserId(v string) {
	o.BankingUserId = &v
}

// GetBankingCustomerId returns the BankingCustomerId field value if set, zero value otherwise.
func (o *EditBankConnectionParams) GetBankingCustomerId() string {
	if o == nil || o.BankingCustomerId == nil {
		var ret string
		return ret
	}
	return *o.BankingCustomerId
}

// GetBankingCustomerIdOk returns a tuple with the BankingCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditBankConnectionParams) GetBankingCustomerIdOk() (*string, bool) {
	if o == nil || o.BankingCustomerId == nil {
		return nil, false
	}
	return o.BankingCustomerId, true
}

// HasBankingCustomerId returns a boolean if a field has been set.
func (o *EditBankConnectionParams) HasBankingCustomerId() bool {
	if o != nil && o.BankingCustomerId != nil {
		return true
	}

	return false
}

// SetBankingCustomerId gets a reference to the given string and assigns it to the BankingCustomerId field.
func (o *EditBankConnectionParams) SetBankingCustomerId(v string) {
	o.BankingCustomerId = &v
}

// GetBankingPin returns the BankingPin field value if set, zero value otherwise.
func (o *EditBankConnectionParams) GetBankingPin() string {
	if o == nil || o.BankingPin == nil {
		var ret string
		return ret
	}
	return *o.BankingPin
}

// GetBankingPinOk returns a tuple with the BankingPin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditBankConnectionParams) GetBankingPinOk() (*string, bool) {
	if o == nil || o.BankingPin == nil {
		return nil, false
	}
	return o.BankingPin, true
}

// HasBankingPin returns a boolean if a field has been set.
func (o *EditBankConnectionParams) HasBankingPin() bool {
	if o != nil && o.BankingPin != nil {
		return true
	}

	return false
}

// SetBankingPin gets a reference to the given string and assigns it to the BankingPin field.
func (o *EditBankConnectionParams) SetBankingPin(v string) {
	o.BankingPin = &v
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *EditBankConnectionParams) GetInterface() BankingInterface {
	if o == nil || o.Interface == nil {
		var ret BankingInterface
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditBankConnectionParams) GetInterfaceOk() (*BankingInterface, bool) {
	if o == nil || o.Interface == nil {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *EditBankConnectionParams) HasInterface() bool {
	if o != nil && o.Interface != nil {
		return true
	}

	return false
}

// SetInterface gets a reference to the given BankingInterface and assigns it to the Interface field.
func (o *EditBankConnectionParams) SetInterface(v BankingInterface) {
	o.Interface = &v
}

// GetLoginCredentials returns the LoginCredentials field value if set, zero value otherwise.
func (o *EditBankConnectionParams) GetLoginCredentials() []LoginCredential {
	if o == nil || o.LoginCredentials == nil {
		var ret []LoginCredential
		return ret
	}
	return *o.LoginCredentials
}

// GetLoginCredentialsOk returns a tuple with the LoginCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditBankConnectionParams) GetLoginCredentialsOk() (*[]LoginCredential, bool) {
	if o == nil || o.LoginCredentials == nil {
		return nil, false
	}
	return o.LoginCredentials, true
}

// HasLoginCredentials returns a boolean if a field has been set.
func (o *EditBankConnectionParams) HasLoginCredentials() bool {
	if o != nil && o.LoginCredentials != nil {
		return true
	}

	return false
}

// SetLoginCredentials gets a reference to the given []LoginCredential and assigns it to the LoginCredentials field.
func (o *EditBankConnectionParams) SetLoginCredentials(v []LoginCredential) {
	o.LoginCredentials = &v
}

// GetDefaultTwoStepProcedureId returns the DefaultTwoStepProcedureId field value if set, zero value otherwise.
func (o *EditBankConnectionParams) GetDefaultTwoStepProcedureId() string {
	if o == nil || o.DefaultTwoStepProcedureId == nil {
		var ret string
		return ret
	}
	return *o.DefaultTwoStepProcedureId
}

// GetDefaultTwoStepProcedureIdOk returns a tuple with the DefaultTwoStepProcedureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EditBankConnectionParams) GetDefaultTwoStepProcedureIdOk() (*string, bool) {
	if o == nil || o.DefaultTwoStepProcedureId == nil {
		return nil, false
	}
	return o.DefaultTwoStepProcedureId, true
}

// HasDefaultTwoStepProcedureId returns a boolean if a field has been set.
func (o *EditBankConnectionParams) HasDefaultTwoStepProcedureId() bool {
	if o != nil && o.DefaultTwoStepProcedureId != nil {
		return true
	}

	return false
}

// SetDefaultTwoStepProcedureId gets a reference to the given string and assigns it to the DefaultTwoStepProcedureId field.
func (o *EditBankConnectionParams) SetDefaultTwoStepProcedureId(v string) {
	o.DefaultTwoStepProcedureId = &v
}

func (o EditBankConnectionParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.BankingUserId != nil {
		toSerialize["bankingUserId"] = o.BankingUserId
	}
	if o.BankingCustomerId != nil {
		toSerialize["bankingCustomerId"] = o.BankingCustomerId
	}
	if o.BankingPin != nil {
		toSerialize["bankingPin"] = o.BankingPin
	}
	if o.Interface != nil {
		toSerialize["interface"] = o.Interface
	}
	if o.LoginCredentials != nil {
		toSerialize["loginCredentials"] = o.LoginCredentials
	}
	if o.DefaultTwoStepProcedureId != nil {
		toSerialize["defaultTwoStepProcedureId"] = o.DefaultTwoStepProcedureId
	}
	return json.Marshal(toSerialize)
}

type NullableEditBankConnectionParams struct {
	value *EditBankConnectionParams
	isSet bool
}

func (v NullableEditBankConnectionParams) Get() *EditBankConnectionParams {
	return v.value
}

func (v *NullableEditBankConnectionParams) Set(val *EditBankConnectionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEditBankConnectionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEditBankConnectionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEditBankConnectionParams(val *EditBankConnectionParams) *NullableEditBankConnectionParams {
	return &NullableEditBankConnectionParams{value: val, isSet: true}
}

func (v NullableEditBankConnectionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEditBankConnectionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


