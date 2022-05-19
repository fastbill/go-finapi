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

// ClientConfigurationParams Client configuration parameters
type ClientConfigurationParams struct {
	// Callback URL to which finAPI sends the notification messages that are triggered from the automatic batch update of the users' bank connections. This field is only relevant if the automatic batch update is enabled for your client. For details about what the notification messages look like, please see the documentation in the 'Notification Rules' section. finAPI will call this URL with HTTP method POST. Note that the response of the call is not processed by finAPI. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha environment, it MUST be a SSL-secured (https) URL on the finAPI live system.<p>The maximum allowed length of the URL is 512. If you have previously set a callback URL and now want to clear it (thus disabling user-related notifications altogether), you can pass an empty string (\"\").
	UserNotificationCallbackUrl *string `json:"userNotificationCallbackUrl,omitempty"`
	// Callback URL for user synchronization. This field should be set if you - as a finAPI customer - have multiple clients using finAPI. In such case, all of your clients will share the same user base, making it possible for a user to be created in one client, but then deleted in another. To keep the client-side user data consistent in all clients, you should set a callback URL for each client. finAPI will send a notification to the callback URL of each client whenever a user of your user base gets deleted. Note that finAPI will send a deletion notification to ALL clients, including the one that made the user deletion request to finAPI. So when deleting a user in finAPI, a client should rely on the callback to delete the user on its own side. <p>The notification that finAPI sends to the clients' callback URLs will be a POST request, with this body: <pre>{    \"userId\" : string // contains the identifier of the deleted user    \"event\" : string // this will always be \"DELETED\" }</pre><br/>Note that finAPI does not process the response of this call. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha system, it MUST be a SSL-secured (https) URL on the live system.</p>As long as you have just one client, you can ignore this field and let it be null. However keep in mind that in this case your client will not receive any callback when a user gets deleted - so the deletion of the user on the client-side must not be forgotten. Of course you may still use the callback URL even for just one client, if you want to implement the deletion of the user on the client-side via the callback from finAPI.<p> The maximum allowed length of the URL is 512. If you have previously set a callback URL and now want to clear it (thus disabling user synchronization related notifications for this client), you can pass an empty string (\"\").
	UserSynchronizationCallbackUrl *string `json:"userSynchronizationCallbackUrl,omitempty"`
	// The validity period that newly requested refresh tokens initially have (in seconds). The value must be greater than or equal to 60, or 0. A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation , or when a user gets locked, or when the password is reset for a user).
	RefreshTokensValidityPeriod *int32 `json:"refreshTokensValidityPeriod,omitempty"`
	// The validity period that newly requested access tokens for users initially have (in seconds). The value must be greater than or equal to 60, or 0. A value of 0 means that the tokens never expire.
	UserAccessTokensValidityPeriod *int32 `json:"userAccessTokensValidityPeriod,omitempty"`
	// The validity period that newly requested access tokens for clients initially have (in seconds). The value must be greater than or equal to 60, or 0. A value of 0 means that the tokens never expire.
	ClientAccessTokensValidityPeriod *int32 `json:"clientAccessTokensValidityPeriod,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='?product=web_form_2.0#post-/api/profiles' target='_blank'>here</a> to the 'storeSecrets' field instead.<br/><br/>Whether finAPI's Web Form should provide a checkbox for the user allowing him to choose whether to store his banking PIN in finAPI. If this field is set to false, then the user won't have an option to store his PIN.
	IsPinStorageAvailableInWebForm *bool `json:"isPinStorageAvailableInWebForm,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='?product=web_form_2.0#post-/api/profiles' target='_blank'>here</a> to the 'storeSecrets' field instead.<br/><br/>Whether finAPI's Web Form will allow the user to choose whether to store login secrets (like a PIN) in finAPI. If this field is set to false, the option will not be available and the secrets not stored.
	StoreSecretsAvailableInWebForm *bool `json:"storeSecretsAvailableInWebForm,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='?product=web_form_2.0#post-/api/profiles' target='_blank'>here</a> to the 'logo' and 'favicon' field instead. The Web Form will now be able to render your logo both in the header and as a favicon.<br/><br/>When an application name is set (e.g. \"My App\"), then finAPI's Web Form will display a text to the user \"Weiterleitung auf finAPI von ...\" (e.g. \"Weiterleitung auf finAPI von My App\"). If you have previously set a application name and now want to clear it, you can pass an empty string (\"\"). Maximum length: 64
	ApplicationName *string `json:"applicationName,omitempty"`
	// The FinTS product registration number. Please follow <a href='https://www.hbci-zka.de/register/prod_register.htm' target='_blank'>this link</a> to apply for a registration number. Only customers who have an AISP or PISP license must define their FinTS product registration number. Customers who are relying on the finAPI Web Form will be assigned to finAPI's FinTS product registration number automatically and do not have to register themselves. During a batch update, finAPI is using the FinTS product registration number of the client, that was used to create the user. If you have previously set a FinTS product registration number and now want to clear it, you can pass an empty string (\"\"). Only hexa decimal characters in capital case with a maximum length of 25 characters are allowed. E.g. 'ABCDEF1234567890ABCDEF123'
	FinTSProductRegistrationNumber *string `json:"finTSProductRegistrationNumber,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='https://documentation.finapi.io/webform/For-best-results!.2477654019.html' target='_blank'>here</a> to the 'errorRedirectUrl' and 'customerSupportUrl' query parameters instead.<br/><br/>Default value for the subject element of support emails. Maximum length is 100. Pass an empty string ('') if you want to clear the current subject default value.
	SupportSubjectDefault *string `json:"supportSubjectDefault,omitempty"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='https://documentation.finapi.io/webform/For-best-results!.2477654019.html' target='_blank'>here</a> to the 'errorRedirectUrl' and 'customerSupportUrl' query parameters instead.<br/><br/>Email address to sent support requests to from the Web Form. Maximum length is 320. Pass an empty string ('') if you want to clear the current email address.
	SupportEmail *string `json:"supportEmail,omitempty"`
	// Whether the set of banks that are available to your client should include “Beta banks”. Beta banks provide pre-release interfaces that are still in a beta phase. Communication to the bank via such interfaces might be unstable, and the correctness and/or quality of data delivery or payment execution cannot be guaranteed.<br/>As the word “BETA” already indicates, Beta banks are subject to changes. Their properties, as well as their behaviour can change based on continuous tests and customer feedback. Also, to keep our bank list clean, we might remove Beta banks at any point in time, including all related user data (bank connections, accounts, transactions etc). We still recommend you to enable beta banks in your application, because it enables us to release a stable interface faster. However, you should point it out to your users when using a beta bank (also see field Bank.isBeta).<br/><br/>If this field is true, then the GET /banks services will include beta banks in their results, and you can use beta banks in any service where you can pass a bank identifier. If the field is false, then beta banks will not exist for your client.
	BetaBanksEnabled *bool `json:"betaBanksEnabled,omitempty"`
}

// NewClientConfigurationParams instantiates a new ClientConfigurationParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientConfigurationParams() *ClientConfigurationParams {
	this := ClientConfigurationParams{}
	return &this
}

// NewClientConfigurationParamsWithDefaults instantiates a new ClientConfigurationParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientConfigurationParamsWithDefaults() *ClientConfigurationParams {
	this := ClientConfigurationParams{}
	return &this
}

// GetUserNotificationCallbackUrl returns the UserNotificationCallbackUrl field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetUserNotificationCallbackUrl() string {
	if o == nil || o.UserNotificationCallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.UserNotificationCallbackUrl
}

// GetUserNotificationCallbackUrlOk returns a tuple with the UserNotificationCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetUserNotificationCallbackUrlOk() (*string, bool) {
	if o == nil || o.UserNotificationCallbackUrl == nil {
		return nil, false
	}
	return o.UserNotificationCallbackUrl, true
}

// HasUserNotificationCallbackUrl returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasUserNotificationCallbackUrl() bool {
	if o != nil && o.UserNotificationCallbackUrl != nil {
		return true
	}

	return false
}

// SetUserNotificationCallbackUrl gets a reference to the given string and assigns it to the UserNotificationCallbackUrl field.
func (o *ClientConfigurationParams) SetUserNotificationCallbackUrl(v string) {
	o.UserNotificationCallbackUrl = &v
}

// GetUserSynchronizationCallbackUrl returns the UserSynchronizationCallbackUrl field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetUserSynchronizationCallbackUrl() string {
	if o == nil || o.UserSynchronizationCallbackUrl == nil {
		var ret string
		return ret
	}
	return *o.UserSynchronizationCallbackUrl
}

// GetUserSynchronizationCallbackUrlOk returns a tuple with the UserSynchronizationCallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetUserSynchronizationCallbackUrlOk() (*string, bool) {
	if o == nil || o.UserSynchronizationCallbackUrl == nil {
		return nil, false
	}
	return o.UserSynchronizationCallbackUrl, true
}

// HasUserSynchronizationCallbackUrl returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasUserSynchronizationCallbackUrl() bool {
	if o != nil && o.UserSynchronizationCallbackUrl != nil {
		return true
	}

	return false
}

// SetUserSynchronizationCallbackUrl gets a reference to the given string and assigns it to the UserSynchronizationCallbackUrl field.
func (o *ClientConfigurationParams) SetUserSynchronizationCallbackUrl(v string) {
	o.UserSynchronizationCallbackUrl = &v
}

// GetRefreshTokensValidityPeriod returns the RefreshTokensValidityPeriod field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetRefreshTokensValidityPeriod() int32 {
	if o == nil || o.RefreshTokensValidityPeriod == nil {
		var ret int32
		return ret
	}
	return *o.RefreshTokensValidityPeriod
}

// GetRefreshTokensValidityPeriodOk returns a tuple with the RefreshTokensValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetRefreshTokensValidityPeriodOk() (*int32, bool) {
	if o == nil || o.RefreshTokensValidityPeriod == nil {
		return nil, false
	}
	return o.RefreshTokensValidityPeriod, true
}

// HasRefreshTokensValidityPeriod returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasRefreshTokensValidityPeriod() bool {
	if o != nil && o.RefreshTokensValidityPeriod != nil {
		return true
	}

	return false
}

// SetRefreshTokensValidityPeriod gets a reference to the given int32 and assigns it to the RefreshTokensValidityPeriod field.
func (o *ClientConfigurationParams) SetRefreshTokensValidityPeriod(v int32) {
	o.RefreshTokensValidityPeriod = &v
}

// GetUserAccessTokensValidityPeriod returns the UserAccessTokensValidityPeriod field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetUserAccessTokensValidityPeriod() int32 {
	if o == nil || o.UserAccessTokensValidityPeriod == nil {
		var ret int32
		return ret
	}
	return *o.UserAccessTokensValidityPeriod
}

// GetUserAccessTokensValidityPeriodOk returns a tuple with the UserAccessTokensValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetUserAccessTokensValidityPeriodOk() (*int32, bool) {
	if o == nil || o.UserAccessTokensValidityPeriod == nil {
		return nil, false
	}
	return o.UserAccessTokensValidityPeriod, true
}

// HasUserAccessTokensValidityPeriod returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasUserAccessTokensValidityPeriod() bool {
	if o != nil && o.UserAccessTokensValidityPeriod != nil {
		return true
	}

	return false
}

// SetUserAccessTokensValidityPeriod gets a reference to the given int32 and assigns it to the UserAccessTokensValidityPeriod field.
func (o *ClientConfigurationParams) SetUserAccessTokensValidityPeriod(v int32) {
	o.UserAccessTokensValidityPeriod = &v
}

// GetClientAccessTokensValidityPeriod returns the ClientAccessTokensValidityPeriod field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetClientAccessTokensValidityPeriod() int32 {
	if o == nil || o.ClientAccessTokensValidityPeriod == nil {
		var ret int32
		return ret
	}
	return *o.ClientAccessTokensValidityPeriod
}

// GetClientAccessTokensValidityPeriodOk returns a tuple with the ClientAccessTokensValidityPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetClientAccessTokensValidityPeriodOk() (*int32, bool) {
	if o == nil || o.ClientAccessTokensValidityPeriod == nil {
		return nil, false
	}
	return o.ClientAccessTokensValidityPeriod, true
}

// HasClientAccessTokensValidityPeriod returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasClientAccessTokensValidityPeriod() bool {
	if o != nil && o.ClientAccessTokensValidityPeriod != nil {
		return true
	}

	return false
}

// SetClientAccessTokensValidityPeriod gets a reference to the given int32 and assigns it to the ClientAccessTokensValidityPeriod field.
func (o *ClientConfigurationParams) SetClientAccessTokensValidityPeriod(v int32) {
	o.ClientAccessTokensValidityPeriod = &v
}

// GetIsPinStorageAvailableInWebForm returns the IsPinStorageAvailableInWebForm field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetIsPinStorageAvailableInWebForm() bool {
	if o == nil || o.IsPinStorageAvailableInWebForm == nil {
		var ret bool
		return ret
	}
	return *o.IsPinStorageAvailableInWebForm
}

// GetIsPinStorageAvailableInWebFormOk returns a tuple with the IsPinStorageAvailableInWebForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetIsPinStorageAvailableInWebFormOk() (*bool, bool) {
	if o == nil || o.IsPinStorageAvailableInWebForm == nil {
		return nil, false
	}
	return o.IsPinStorageAvailableInWebForm, true
}

// HasIsPinStorageAvailableInWebForm returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasIsPinStorageAvailableInWebForm() bool {
	if o != nil && o.IsPinStorageAvailableInWebForm != nil {
		return true
	}

	return false
}

// SetIsPinStorageAvailableInWebForm gets a reference to the given bool and assigns it to the IsPinStorageAvailableInWebForm field.
func (o *ClientConfigurationParams) SetIsPinStorageAvailableInWebForm(v bool) {
	o.IsPinStorageAvailableInWebForm = &v
}

// GetStoreSecretsAvailableInWebForm returns the StoreSecretsAvailableInWebForm field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetStoreSecretsAvailableInWebForm() bool {
	if o == nil || o.StoreSecretsAvailableInWebForm == nil {
		var ret bool
		return ret
	}
	return *o.StoreSecretsAvailableInWebForm
}

// GetStoreSecretsAvailableInWebFormOk returns a tuple with the StoreSecretsAvailableInWebForm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetStoreSecretsAvailableInWebFormOk() (*bool, bool) {
	if o == nil || o.StoreSecretsAvailableInWebForm == nil {
		return nil, false
	}
	return o.StoreSecretsAvailableInWebForm, true
}

// HasStoreSecretsAvailableInWebForm returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasStoreSecretsAvailableInWebForm() bool {
	if o != nil && o.StoreSecretsAvailableInWebForm != nil {
		return true
	}

	return false
}

// SetStoreSecretsAvailableInWebForm gets a reference to the given bool and assigns it to the StoreSecretsAvailableInWebForm field.
func (o *ClientConfigurationParams) SetStoreSecretsAvailableInWebForm(v bool) {
	o.StoreSecretsAvailableInWebForm = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasApplicationName() bool {
	if o != nil && o.ApplicationName != nil {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *ClientConfigurationParams) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetFinTSProductRegistrationNumber returns the FinTSProductRegistrationNumber field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetFinTSProductRegistrationNumber() string {
	if o == nil || o.FinTSProductRegistrationNumber == nil {
		var ret string
		return ret
	}
	return *o.FinTSProductRegistrationNumber
}

// GetFinTSProductRegistrationNumberOk returns a tuple with the FinTSProductRegistrationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetFinTSProductRegistrationNumberOk() (*string, bool) {
	if o == nil || o.FinTSProductRegistrationNumber == nil {
		return nil, false
	}
	return o.FinTSProductRegistrationNumber, true
}

// HasFinTSProductRegistrationNumber returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasFinTSProductRegistrationNumber() bool {
	if o != nil && o.FinTSProductRegistrationNumber != nil {
		return true
	}

	return false
}

// SetFinTSProductRegistrationNumber gets a reference to the given string and assigns it to the FinTSProductRegistrationNumber field.
func (o *ClientConfigurationParams) SetFinTSProductRegistrationNumber(v string) {
	o.FinTSProductRegistrationNumber = &v
}

// GetSupportSubjectDefault returns the SupportSubjectDefault field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetSupportSubjectDefault() string {
	if o == nil || o.SupportSubjectDefault == nil {
		var ret string
		return ret
	}
	return *o.SupportSubjectDefault
}

// GetSupportSubjectDefaultOk returns a tuple with the SupportSubjectDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetSupportSubjectDefaultOk() (*string, bool) {
	if o == nil || o.SupportSubjectDefault == nil {
		return nil, false
	}
	return o.SupportSubjectDefault, true
}

// HasSupportSubjectDefault returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasSupportSubjectDefault() bool {
	if o != nil && o.SupportSubjectDefault != nil {
		return true
	}

	return false
}

// SetSupportSubjectDefault gets a reference to the given string and assigns it to the SupportSubjectDefault field.
func (o *ClientConfigurationParams) SetSupportSubjectDefault(v string) {
	o.SupportSubjectDefault = &v
}

// GetSupportEmail returns the SupportEmail field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetSupportEmail() string {
	if o == nil || o.SupportEmail == nil {
		var ret string
		return ret
	}
	return *o.SupportEmail
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetSupportEmailOk() (*string, bool) {
	if o == nil || o.SupportEmail == nil {
		return nil, false
	}
	return o.SupportEmail, true
}

// HasSupportEmail returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasSupportEmail() bool {
	if o != nil && o.SupportEmail != nil {
		return true
	}

	return false
}

// SetSupportEmail gets a reference to the given string and assigns it to the SupportEmail field.
func (o *ClientConfigurationParams) SetSupportEmail(v string) {
	o.SupportEmail = &v
}

// GetBetaBanksEnabled returns the BetaBanksEnabled field value if set, zero value otherwise.
func (o *ClientConfigurationParams) GetBetaBanksEnabled() bool {
	if o == nil || o.BetaBanksEnabled == nil {
		var ret bool
		return ret
	}
	return *o.BetaBanksEnabled
}

// GetBetaBanksEnabledOk returns a tuple with the BetaBanksEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientConfigurationParams) GetBetaBanksEnabledOk() (*bool, bool) {
	if o == nil || o.BetaBanksEnabled == nil {
		return nil, false
	}
	return o.BetaBanksEnabled, true
}

// HasBetaBanksEnabled returns a boolean if a field has been set.
func (o *ClientConfigurationParams) HasBetaBanksEnabled() bool {
	if o != nil && o.BetaBanksEnabled != nil {
		return true
	}

	return false
}

// SetBetaBanksEnabled gets a reference to the given bool and assigns it to the BetaBanksEnabled field.
func (o *ClientConfigurationParams) SetBetaBanksEnabled(v bool) {
	o.BetaBanksEnabled = &v
}

func (o ClientConfigurationParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UserNotificationCallbackUrl != nil {
		toSerialize["userNotificationCallbackUrl"] = o.UserNotificationCallbackUrl
	}
	if o.UserSynchronizationCallbackUrl != nil {
		toSerialize["userSynchronizationCallbackUrl"] = o.UserSynchronizationCallbackUrl
	}
	if o.RefreshTokensValidityPeriod != nil {
		toSerialize["refreshTokensValidityPeriod"] = o.RefreshTokensValidityPeriod
	}
	if o.UserAccessTokensValidityPeriod != nil {
		toSerialize["userAccessTokensValidityPeriod"] = o.UserAccessTokensValidityPeriod
	}
	if o.ClientAccessTokensValidityPeriod != nil {
		toSerialize["clientAccessTokensValidityPeriod"] = o.ClientAccessTokensValidityPeriod
	}
	if o.IsPinStorageAvailableInWebForm != nil {
		toSerialize["isPinStorageAvailableInWebForm"] = o.IsPinStorageAvailableInWebForm
	}
	if o.StoreSecretsAvailableInWebForm != nil {
		toSerialize["storeSecretsAvailableInWebForm"] = o.StoreSecretsAvailableInWebForm
	}
	if o.ApplicationName != nil {
		toSerialize["applicationName"] = o.ApplicationName
	}
	if o.FinTSProductRegistrationNumber != nil {
		toSerialize["finTSProductRegistrationNumber"] = o.FinTSProductRegistrationNumber
	}
	if o.SupportSubjectDefault != nil {
		toSerialize["supportSubjectDefault"] = o.SupportSubjectDefault
	}
	if o.SupportEmail != nil {
		toSerialize["supportEmail"] = o.SupportEmail
	}
	if o.BetaBanksEnabled != nil {
		toSerialize["betaBanksEnabled"] = o.BetaBanksEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableClientConfigurationParams struct {
	value *ClientConfigurationParams
	isSet bool
}

func (v NullableClientConfigurationParams) Get() *ClientConfigurationParams {
	return v.value
}

func (v *NullableClientConfigurationParams) Set(val *ClientConfigurationParams) {
	v.value = val
	v.isSet = true
}

func (v NullableClientConfigurationParams) IsSet() bool {
	return v.isSet
}

func (v *NullableClientConfigurationParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientConfigurationParams(val *ClientConfigurationParams) *NullableClientConfigurationParams {
	return &NullableClientConfigurationParams{value: val, isSet: true}
}

func (v NullableClientConfigurationParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientConfigurationParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


