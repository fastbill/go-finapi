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

// ClientConfiguration Client configuration parameters
type ClientConfiguration struct {
	// Whether your client is allowed to call PFM services (Personal Finance Management). The set of PFM services is the following:<br/><br/>&bull; all /mandatorAdmin/ibanRules/_* and /mandatorAdmin/keywordRules/_* services<br/>&bull; GET /accounts/dailyBalances<br/>&bull; all /transactions/_* services, except for GET /transactions/[id(s)] and DELETE /transactions/[id]<br/>&bull; all /categories/_* services, except for GET /categories/[id(s)]<br/>&bull; all /labels/_* services<br/>&bull; all /notificationRules/_* services<br/>&bull; all /tests/_* services
	PfmServicesEnabled bool `json:"pfmServicesEnabled"`
	// Whether finAPI performs a regular automatic update of your users' bank connections. To find out how the automatic batch update is configured for your client, i.e. which bank connections get updated, and at which time and interval, please contact your Sys-Admin. Note that even if the automatic batch update is enabled for your client, individual users can still disable the feature for their own bank connections.
	IsAutomaticBatchUpdateEnabled bool `json:"isAutomaticBatchUpdateEnabled"`
	// Whether development mode is enabled. This setting is enabled on mandator level and allows any user to access the 'Mock batch update' service. <br/><br/>NOTE: This flag is meant for testing purposes during development of your application. <br/>This is why this will never be enabled on a production environment.
	IsDevelopmentModeEnabled bool `json:"isDevelopmentModeEnabled"`
	// Whether finAPI will download data (balance and transactions) for bank accounts with a currency other than EUR (affects all users). If this flag is false, then non-EUR accounts will still be returned in the account list, but they will have no balance and no transactions. Note that this currently applies to Checking accounts only.
	IsNonEuroAccountsSupported bool `json:"isNonEuroAccountsSupported"`
	// Whether transactions will be categorized as soon as they are downloaded. <br/>In case this flag is false, the user needs to manually trigger categorization using the 'Trigger categorization' service.
	IsAutoCategorizationEnabled bool `json:"isAutoCategorizationEnabled"`
	// <strong>Type:</strong> MandatorLicense<br/> The license associated with your client. <br/>The licensing model affects the TPP registration data used to connect to the bank (e.g. <b>finTSProductRegistrationNumber</b> for FINTS_SERVER interface). Licenses are administered by finAPI. Please contact the support to change the license that was set up for you.<br/>Possible values are:<br/>UNLICENSED: finAPI will use its own TPP registration to connect to the bank for both account information services (AIS) and payment initiation services (PIS).<br/>AISP: finAPI will use its own TPP registration to connect to the bank for PIS, and your registration for AIS.<br/>PISP: finAPI will use its own TPP registration to connect to the bank for AIS, and your registration for PIS.<br/>FULLY_LICENSED: finAPI will use your TPP registration to connect to the bank for both AIS and PIS.
	MandatorLicense MandatorLicense `json:"mandatorLicense"`
	// <strong>Type:</strong> PreferredConsentType<br/> The preferred consent type that will be used for the XS2A interface.<br/><br/><b>ONETIME</b> - The consent can only be used once to download data associated with the account. The consent won’t be saved by finAPI.<br/><b>RECURRING</b> - The consent is valid for up to 90 days and can be used by finAPI to access and download account data for up to 4 times per day.<br/><br/>NOTE: If the bank does not support the preferred consent type, then finAPI will default to the other type.
	PreferredConsentType PreferredConsentType `json:"preferredConsentType"`
	// Callback URL to which finAPI sends the notification messages that are triggered from the automatic batch update of the users' bank connections. This field is only relevant if the automatic batch update is enabled for your client. For details about what the notification messages look like, please see the documentation in the 'Notification Rules' section. finAPI will call this URL with HTTP method POST. Note that the response of the call is not processed by finAPI. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha environment, it MUST be a SSL-secured (https) URL on the finAPI live system.
	UserNotificationCallbackUrl NullableString `json:"userNotificationCallbackUrl"`
	// Callback URL for user synchronization. This field should be set if you - as a finAPI customer - have multiple clients using finAPI. In such case, all of your clients will share the same user base, making it possible for a user to be created in one client, but then deleted in another. To keep the client-side user data consistent in all clients, you should set a callback URL for each client. finAPI will send a notification to the callback URL of each client whenever a user of your user base gets deleted. Note that finAPI will send a deletion notification to ALL clients, including the one that made the user deletion request to finAPI. So when deleting a user in finAPI, a client should rely on the callback to delete the user on its own side. <p>The notification that finAPI sends to the clients' callback URLs will be a POST request, with this body: <pre>{    \"userId\" : string // contains the identifier of the deleted user    \"event\" : string // this will always be \"DELETED\" }</pre><br/>Note that finAPI does not process the response of this call. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha environment, it MUST be a SSL-secured (https) URL on the finAPI live system.</p>As long as you have just one client, you can ignore this field and let it be null. However keep in mind that in this case your client will not receive any callback when a user gets deleted - so the deletion of the user on the client-side must not be forgotten. Of course you may still use the callback URL even for just one client, if you want to implement the deletion of the user on the client-side via the callback from finAPI.
	UserSynchronizationCallbackUrl NullableString `json:"userSynchronizationCallbackUrl"`
	// The validity period that newly requested refresh tokens initially have (in seconds). A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation, or when a user gets locked, or when the password is reset for a user).
	RefreshTokensValidityPeriod int32 `json:"refreshTokensValidityPeriod"`
	// The validity period that newly requested access tokens for users initially have (in seconds). A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation, or when a user gets locked, or when the password is reset for a user).
	UserAccessTokensValidityPeriod int32 `json:"userAccessTokensValidityPeriod"`
	// The validity period that newly requested access tokens for clients initially have (in seconds). A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation).
	ClientAccessTokensValidityPeriod int32 `json:"clientAccessTokensValidityPeriod"`
	// Number of consecutive failed login attempts of a user into his finAPI account that is allowed before finAPI locks the user's account. When a user's account is locked, finAPI will invalidate all user's tokens and it will deny any service call in the context of this user (i.e. any call to a service using one of the user's authorization tokens, as well as the service for requesting a new token for this user). To unlock a user's account, a new password must be set for the account by the client (see the services /users/requestPasswordChange and /users/executePasswordChange). Once a new password has been set, all services will be available again for this user and the user's failed login attempts counter is reset to 0. The user's failed login attempts counter is also reset whenever a new authorization token has been successfully retrieved, or whenever the user himself changes his password.<br/><br/>Note that when this field has a value of 0, it means that there is no limit for user login attempts, i.e. finAPI will never lock user accounts.
	MaxUserLoginAttempts int32 `json:"maxUserLoginAttempts"`
	// This setting defines the upper limit of how much of an account's transactions history may be downloaded whenever a new account is imported, across all of your users. More technically, it depicts the maximum number of days for which transactions might get downloaded, starting from - and including - the date of the account import. '0' means that there is no limitation.
	TransactionImportLimitation int32 `json:"transactionImportLimitation"`
	// Whether users that are created with this client are automatically verified on creation. If this field is set to 'false', then any user that is created with this client must first be verified with the \"Verify a user\" service before he can be authorized. If the field is 'true', then no verification is required by the client and the user can be authorized immediately after creation.
	IsUserAutoVerificationEnabled bool `json:"isUserAutoVerificationEnabled"`
	// Whether this client is a 'Mandator Admin'. Mandator Admins are special clients that can access the 'Mandator Administration' section of finAPI. If you do not yet have credentials for a Mandator Admin, please contact us at support@finapi.io. For further information, please refer to <a href='https://documentation.finapi.io/access/Application-management.2763423767.html' target='_blank'>this page</a> on our Access Public Documentation.
	IsMandatorAdmin bool `json:"isMandatorAdmin"`
	// Whether finAPI is allowed to use the WEB_SCRAPER interface for data download or payments. <br/><br/>If this field is set to 'true', then finAPI might download data from the online banking websites of banks (either in addition to other interfaces, or as the sole data source for the download). Also, it will be possible to do payments via the WEB_SCRAPER interface.<br/><br/>If this field is set to 'false', then finAPI will not use any web scrapers. Payments via the WEB_SCRAPER interface will not be possible, and finAPI will not allow any data download for banks where no other interface except WEB_SCRAPER is available. <br/><br/>Please contact your Sys-Admin if you want to change this setting.
	IsWebScrapingEnabled bool `json:"isWebScrapingEnabled"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/><br/>Whether this client is allowed to access XS2A services.
	IsXs2aEnabled bool `json:"isXs2aEnabled"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'storeSecretsAvailableInWebForm' field instead.<br/><br/>Whether finAPI's Web Form will provide a checkbox for the user allowing him to choose whether to store his banking PIN in finAPI. If this field is set to false, then the user won't have an option to store his PIN.
	PinStorageAvailableInWebForm bool `json:"pinStorageAvailableInWebForm"`
	// Whether this client is allowed to do PIS.<br/><br/>Note that on the Sandbox environment, it is always possible to execute payments (regardless of what this field says), as long as you are using a test bank (see Bank.isTestBank)
	PaymentsEnabled bool `json:"paymentsEnabled"`
	// Whether this client is allowed to do standalone PIS (doing money transfers and standing orders for accounts that are not imported in finAPI).<br/><br/>Note that on the Sandbox environment, it is always possible to execute payments and standing orders (regardless of what this field says), as long as you are using a test bank (see Bank.isTestBank)
	IsStandalonePaymentsEnabled bool `json:"isStandalonePaymentsEnabled"`
	AvailableBankGroups []string `json:"availableBankGroups"`
	Products []Product `json:"products"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='?product=web_form_2.0#post-/api/profiles' target='_blank'>here</a> to the 'logo' and 'favicon' field instead. The Web Form will now be able to render your logo both in the header and as a favicon.<br/><br/>Application name. When an application name is set (e.g. \"My App\"), then finAPI's Web Form will display a text to the user \"Weiterleitung auf finAPI von ...\" (e.g. \"Weiterleitung auf finAPI von MyApp\").
	ApplicationName NullableString `json:"applicationName"`
	// The FinTS product registration number. If a value is stored, this will always be 'XXXXX'.
	FinTSProductRegistrationNumber NullableString `json:"finTSProductRegistrationNumber"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='?product=web_form_2.0#post-/api/profiles' target='_blank'>here</a> to the 'storeSecrets' field instead.<br/><br/>Whether finAPI's Web Form should provide a checkbox for the user allowing him to choose whether to store his banking PIN in finAPI. If this field is set to false, then the user won't have an option to store his PIN.
	StoreSecretsAvailableInWebForm bool `json:"storeSecretsAvailableInWebForm"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='https://documentation.finapi.io/webform/For-best-results!.2477654019.html' target='_blank'>here</a> to the 'errorRedirectUrl' and 'customerSupportUrl' query parameters instead.<br/><br/>Default value for the subject element of support emails.
	SupportSubjectDefault NullableString `json:"supportSubjectDefault"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer <a href='https://documentation.finapi.io/webform/For-best-results!.2477654019.html' target='_blank'>here</a> to the 'errorRedirectUrl' and 'customerSupportUrl' query parameters instead.<br/><br/>Email address to sent support requests to from the Web Form.
	SupportEmail NullableString `json:"supportEmail"`
	// <strong>Type:</strong> WebFormMode<br/> Indicates whether the client is using the finAPI Web Form for Account Initiation Services.<br/><br/>Possible values: <br/>&bull; <code>DISABLED</code> - No Web Form is triggered<br/>&bull; <code>INTERNAL</code> - THIS VALUE IS DEPRECATED AND WILL BE REMOVED. Hence, we request customers to foresee a migration to Web Form 2.0 (value <code>EXTERNAL</code>).<br/>End users will be directed to the classical Web Form implementation.<br/>&bull; <code>EXTERNAL</code> - End users will be directed to the <a href='https://documentation.finapi.io/webform/Introduction.2038136860.html'  target='_blank'>new Web Form</a> implementation.
	AisWebFormMode WebFormMode `json:"aisWebFormMode"`
	// <strong>Type:</strong> WebFormMode<br/> Indicates whether the client is using the finAPI Web Form for Standard Payment Initiation Services (Payments for accounts that have been imported in finAPI).<br/><br/>Possible values: <br/>&bull; <code>DISABLED</code> - No Web Form is triggered<br/>&bull; <code>INTERNAL</code> - THIS VALUE IS DEPRECATED AND WILL BE REMOVED. Hence, we request customers to foresee a migration to Web Form 2.0 (value <code>EXTERNAL</code>).<br/>End users will be directed to the classical Web Form implementation.<br/>&bull; <code>EXTERNAL</code> - End users will be directed to the <a href='https://documentation.finapi.io/webform/Introduction.2038136860.html'  target='_blank'>new Web Form</a> implementation.
	PisWebFormMode WebFormMode `json:"pisWebFormMode"`
	// <strong>Type:</strong> WebFormMode<br/> Indicates whether the client is using the finAPI Web Form for Standalone Payment Initiation Services (Payments without account import).<br/><br/>Possible values: <br/>&bull; <code>DISABLED</code> - No Web Form is triggered<br/>&bull; <code>INTERNAL</code> - THIS VALUE IS DEPRECATED AND WILL BE REMOVED. Hence, we request customers to foresee a migration to Web Form 2.0 (value <code>EXTERNAL</code>).<br/>End users will be directed to the classical Web Form implementation.<br/>&bull; <code>EXTERNAL</code> - End users will be directed to the <a href='https://documentation.finapi.io/webform/Introduction.2038136860.html'  target='_blank'>new Web Form</a> implementation.
	PisStandaloneWebFormMode WebFormMode `json:"pisStandaloneWebFormMode"`
	// Whether the set of banks that are available to your client contains “Beta banks”. Beta banks provide pre-release interfaces that are still in a beta phase. Communication to the bank via such interfaces might be unstable, and the correctness and/or quality of data delivery or payment execution cannot be guaranteed.<br/>As the word “BETA” already indicates, Beta banks are subject to changes. Their properties, as well as their behaviour can change based on continuous tests and customer feedback. Also, to keep our bank list clean, we might remove Beta banks at any point in time, including all related user data (bank connections, accounts, transactions etc). We still recommend you to enable beta banks in your application, because it enables us to release a stable interface faster. However, you should point it out to your users when using a beta bank (also see field Bank.isBeta).<br/><br/>If this field is true, then the GET /banks services will include beta banks in their results, and you can use beta banks in any service where you can pass a bank identifier. If the field is false, then beta banks will not exist for your client.
	BetaBanksEnabled bool `json:"betaBanksEnabled"`
	// <strong>Type:</strong> Category<br/> Defines the set of transaction categories to which your client is restricted. When retrieving transactions (via the GET /transactions services), you may request only those transactions whose 'category' is one of the listed categories. If this field is null, then there are no restrictions for your client, and you may retrieve the full set of imported transactions.
	CategoryRestrictions []Category `json:"categoryRestrictions"`
	// THIS FIELD IS DEPRECATED AND WILL BE REMOVED.<br/> Please refer <a href='?product=web_form_2.0#post-/api/profiles' target='_blank'>here</a> to the 'skipConfirmationView' field instead.<br/><br/>This flag indicates whether the Web Form should get removed from the parent page automatically once it’s finished. It applies ONLY to the classical embedded Web Form. That means it’s only applied if aisWebFormMode, pisWebFormMode or pisStandaloneWebFormMode are defined as INTERNAL. In case you are using our standalone Web Form by redirecting the user to our Web Form link, this feature has no effect.
	AutoDismountWebForm bool `json:"autoDismountWebForm"`
	// The list of allowed origins for cross-origin requests. The CORS configuration applies to all the API services except for the /oauth services. If this list is empty, then CORS is not enabled for this client. Please contact the support if you want to enable or change the client's CORS configuration.
	CorsAllowedOrigins []string `json:"corsAllowedOrigins"`
}

// NewClientConfiguration instantiates a new ClientConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientConfiguration(pfmServicesEnabled bool, isAutomaticBatchUpdateEnabled bool, isDevelopmentModeEnabled bool, isNonEuroAccountsSupported bool, isAutoCategorizationEnabled bool, mandatorLicense MandatorLicense, preferredConsentType PreferredConsentType, userNotificationCallbackUrl NullableString, userSynchronizationCallbackUrl NullableString, refreshTokensValidityPeriod int32, userAccessTokensValidityPeriod int32, clientAccessTokensValidityPeriod int32, maxUserLoginAttempts int32, transactionImportLimitation int32, isUserAutoVerificationEnabled bool, isMandatorAdmin bool, isWebScrapingEnabled bool, isXs2aEnabled bool, pinStorageAvailableInWebForm bool, paymentsEnabled bool, isStandalonePaymentsEnabled bool, availableBankGroups []string, products []Product, applicationName NullableString, finTSProductRegistrationNumber NullableString, storeSecretsAvailableInWebForm bool, supportSubjectDefault NullableString, supportEmail NullableString, aisWebFormMode WebFormMode, pisWebFormMode WebFormMode, pisStandaloneWebFormMode WebFormMode, betaBanksEnabled bool, categoryRestrictions []Category, autoDismountWebForm bool, corsAllowedOrigins []string, ) *ClientConfiguration {
	this := ClientConfiguration{}
	this.PfmServicesEnabled = pfmServicesEnabled
	this.IsAutomaticBatchUpdateEnabled = isAutomaticBatchUpdateEnabled
	this.IsDevelopmentModeEnabled = isDevelopmentModeEnabled
	this.IsNonEuroAccountsSupported = isNonEuroAccountsSupported
	this.IsAutoCategorizationEnabled = isAutoCategorizationEnabled
	this.MandatorLicense = mandatorLicense
	this.PreferredConsentType = preferredConsentType
	this.UserNotificationCallbackUrl = userNotificationCallbackUrl
	this.UserSynchronizationCallbackUrl = userSynchronizationCallbackUrl
	this.RefreshTokensValidityPeriod = refreshTokensValidityPeriod
	this.UserAccessTokensValidityPeriod = userAccessTokensValidityPeriod
	this.ClientAccessTokensValidityPeriod = clientAccessTokensValidityPeriod
	this.MaxUserLoginAttempts = maxUserLoginAttempts
	this.TransactionImportLimitation = transactionImportLimitation
	this.IsUserAutoVerificationEnabled = isUserAutoVerificationEnabled
	this.IsMandatorAdmin = isMandatorAdmin
	this.IsWebScrapingEnabled = isWebScrapingEnabled
	this.IsXs2aEnabled = isXs2aEnabled
	this.PinStorageAvailableInWebForm = pinStorageAvailableInWebForm
	this.PaymentsEnabled = paymentsEnabled
	this.IsStandalonePaymentsEnabled = isStandalonePaymentsEnabled
	this.AvailableBankGroups = availableBankGroups
	this.Products = products
	this.ApplicationName = applicationName
	this.FinTSProductRegistrationNumber = finTSProductRegistrationNumber
	this.StoreSecretsAvailableInWebForm = storeSecretsAvailableInWebForm
	this.SupportSubjectDefault = supportSubjectDefault
	this.SupportEmail = supportEmail
	this.AisWebFormMode = aisWebFormMode
	this.PisWebFormMode = pisWebFormMode
	this.PisStandaloneWebFormMode = pisStandaloneWebFormMode
	this.BetaBanksEnabled = betaBanksEnabled
	this.CategoryRestrictions = categoryRestrictions
	this.AutoDismountWebForm = autoDismountWebForm
	this.CorsAllowedOrigins = corsAllowedOrigins
	return &this
}

// NewClientConfigurationWithDefaults instantiates a new ClientConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientConfigurationWithDefaults() *ClientConfiguration {
	this := ClientConfiguration{}
	return &this
}

// GetPfmServicesEnabled returns the PfmServicesEnabled field value
func (o *ClientConfiguration) GetPfmServicesEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.PfmServicesEnabled
}

// GetPfmServicesEnabledOk returns a tuple with the PfmServicesEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetPfmServicesEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PfmServicesEnabled, true
}

// SetPfmServicesEnabled sets field value
func (o *ClientConfiguration) SetPfmServicesEnabled(v bool) {
	o.PfmServicesEnabled = v
}

// GetIsAutomaticBatchUpdateEnabled returns the IsAutomaticBatchUpdateEnabled field value
func (o *ClientConfiguration) GetIsAutomaticBatchUpdateEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsAutomaticBatchUpdateEnabled
}

// GetIsAutomaticBatchUpdateEnabledOk returns a tuple with the IsAutomaticBatchUpdateEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetIsAutomaticBatchUpdateEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsAutomaticBatchUpdateEnabled, true
}

// SetIsAutomaticBatchUpdateEnabled sets field value
func (o *ClientConfiguration) SetIsAutomaticBatchUpdateEnabled(v bool) {
	o.IsAutomaticBatchUpdateEnabled = v
}

// GetIsDevelopmentModeEnabled returns the IsDevelopmentModeEnabled field value
func (o *ClientConfiguration) GetIsDevelopmentModeEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsDevelopmentModeEnabled
}

// GetIsDevelopmentModeEnabledOk returns a tuple with the IsDevelopmentModeEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetIsDevelopmentModeEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsDevelopmentModeEnabled, true
}

// SetIsDevelopmentModeEnabled sets field value
func (o *ClientConfiguration) SetIsDevelopmentModeEnabled(v bool) {
	o.IsDevelopmentModeEnabled = v
}

// GetIsNonEuroAccountsSupported returns the IsNonEuroAccountsSupported field value
func (o *ClientConfiguration) GetIsNonEuroAccountsSupported() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsNonEuroAccountsSupported
}

// GetIsNonEuroAccountsSupportedOk returns a tuple with the IsNonEuroAccountsSupported field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetIsNonEuroAccountsSupportedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsNonEuroAccountsSupported, true
}

// SetIsNonEuroAccountsSupported sets field value
func (o *ClientConfiguration) SetIsNonEuroAccountsSupported(v bool) {
	o.IsNonEuroAccountsSupported = v
}

// GetIsAutoCategorizationEnabled returns the IsAutoCategorizationEnabled field value
func (o *ClientConfiguration) GetIsAutoCategorizationEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsAutoCategorizationEnabled
}

// GetIsAutoCategorizationEnabledOk returns a tuple with the IsAutoCategorizationEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetIsAutoCategorizationEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsAutoCategorizationEnabled, true
}

// SetIsAutoCategorizationEnabled sets field value
func (o *ClientConfiguration) SetIsAutoCategorizationEnabled(v bool) {
	o.IsAutoCategorizationEnabled = v
}

// GetMandatorLicense returns the MandatorLicense field value
func (o *ClientConfiguration) GetMandatorLicense() MandatorLicense {
	if o == nil  {
		var ret MandatorLicense
		return ret
	}

	return o.MandatorLicense
}

// GetMandatorLicenseOk returns a tuple with the MandatorLicense field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetMandatorLicenseOk() (*MandatorLicense, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MandatorLicense, true
}

// SetMandatorLicense sets field value
func (o *ClientConfiguration) SetMandatorLicense(v MandatorLicense) {
	o.MandatorLicense = v
}

// GetPreferredConsentType returns the PreferredConsentType field value
func (o *ClientConfiguration) GetPreferredConsentType() PreferredConsentType {
	if o == nil  {
		var ret PreferredConsentType
		return ret
	}

	return o.PreferredConsentType
}

// GetPreferredConsentTypeOk returns a tuple with the PreferredConsentType field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetPreferredConsentTypeOk() (*PreferredConsentType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PreferredConsentType, true
}

// SetPreferredConsentType sets field value
func (o *ClientConfiguration) SetPreferredConsentType(v PreferredConsentType) {
	o.PreferredConsentType = v
}

// GetUserNotificationCallbackUrl returns the UserNotificationCallbackUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClientConfiguration) GetUserNotificationCallbackUrl() string {
	if o == nil || o.UserNotificationCallbackUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserNotificationCallbackUrl.Get()
}

// GetUserNotificationCallbackUrlOk returns a tuple with the UserNotificationCallbackUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientConfiguration) GetUserNotificationCallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserNotificationCallbackUrl.Get(), o.UserNotificationCallbackUrl.IsSet()
}

// SetUserNotificationCallbackUrl sets field value
func (o *ClientConfiguration) SetUserNotificationCallbackUrl(v string) {
	o.UserNotificationCallbackUrl.Set(&v)
}

// GetUserSynchronizationCallbackUrl returns the UserSynchronizationCallbackUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClientConfiguration) GetUserSynchronizationCallbackUrl() string {
	if o == nil || o.UserSynchronizationCallbackUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.UserSynchronizationCallbackUrl.Get()
}

// GetUserSynchronizationCallbackUrlOk returns a tuple with the UserSynchronizationCallbackUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientConfiguration) GetUserSynchronizationCallbackUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.UserSynchronizationCallbackUrl.Get(), o.UserSynchronizationCallbackUrl.IsSet()
}

// SetUserSynchronizationCallbackUrl sets field value
func (o *ClientConfiguration) SetUserSynchronizationCallbackUrl(v string) {
	o.UserSynchronizationCallbackUrl.Set(&v)
}

// GetRefreshTokensValidityPeriod returns the RefreshTokensValidityPeriod field value
func (o *ClientConfiguration) GetRefreshTokensValidityPeriod() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.RefreshTokensValidityPeriod
}

// GetRefreshTokensValidityPeriodOk returns a tuple with the RefreshTokensValidityPeriod field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetRefreshTokensValidityPeriodOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RefreshTokensValidityPeriod, true
}

// SetRefreshTokensValidityPeriod sets field value
func (o *ClientConfiguration) SetRefreshTokensValidityPeriod(v int32) {
	o.RefreshTokensValidityPeriod = v
}

// GetUserAccessTokensValidityPeriod returns the UserAccessTokensValidityPeriod field value
func (o *ClientConfiguration) GetUserAccessTokensValidityPeriod() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.UserAccessTokensValidityPeriod
}

// GetUserAccessTokensValidityPeriodOk returns a tuple with the UserAccessTokensValidityPeriod field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetUserAccessTokensValidityPeriodOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserAccessTokensValidityPeriod, true
}

// SetUserAccessTokensValidityPeriod sets field value
func (o *ClientConfiguration) SetUserAccessTokensValidityPeriod(v int32) {
	o.UserAccessTokensValidityPeriod = v
}

// GetClientAccessTokensValidityPeriod returns the ClientAccessTokensValidityPeriod field value
func (o *ClientConfiguration) GetClientAccessTokensValidityPeriod() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.ClientAccessTokensValidityPeriod
}

// GetClientAccessTokensValidityPeriodOk returns a tuple with the ClientAccessTokensValidityPeriod field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetClientAccessTokensValidityPeriodOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientAccessTokensValidityPeriod, true
}

// SetClientAccessTokensValidityPeriod sets field value
func (o *ClientConfiguration) SetClientAccessTokensValidityPeriod(v int32) {
	o.ClientAccessTokensValidityPeriod = v
}

// GetMaxUserLoginAttempts returns the MaxUserLoginAttempts field value
func (o *ClientConfiguration) GetMaxUserLoginAttempts() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.MaxUserLoginAttempts
}

// GetMaxUserLoginAttemptsOk returns a tuple with the MaxUserLoginAttempts field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetMaxUserLoginAttemptsOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MaxUserLoginAttempts, true
}

// SetMaxUserLoginAttempts sets field value
func (o *ClientConfiguration) SetMaxUserLoginAttempts(v int32) {
	o.MaxUserLoginAttempts = v
}

// GetTransactionImportLimitation returns the TransactionImportLimitation field value
func (o *ClientConfiguration) GetTransactionImportLimitation() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.TransactionImportLimitation
}

// GetTransactionImportLimitationOk returns a tuple with the TransactionImportLimitation field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetTransactionImportLimitationOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransactionImportLimitation, true
}

// SetTransactionImportLimitation sets field value
func (o *ClientConfiguration) SetTransactionImportLimitation(v int32) {
	o.TransactionImportLimitation = v
}

// GetIsUserAutoVerificationEnabled returns the IsUserAutoVerificationEnabled field value
func (o *ClientConfiguration) GetIsUserAutoVerificationEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsUserAutoVerificationEnabled
}

// GetIsUserAutoVerificationEnabledOk returns a tuple with the IsUserAutoVerificationEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetIsUserAutoVerificationEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsUserAutoVerificationEnabled, true
}

// SetIsUserAutoVerificationEnabled sets field value
func (o *ClientConfiguration) SetIsUserAutoVerificationEnabled(v bool) {
	o.IsUserAutoVerificationEnabled = v
}

// GetIsMandatorAdmin returns the IsMandatorAdmin field value
func (o *ClientConfiguration) GetIsMandatorAdmin() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsMandatorAdmin
}

// GetIsMandatorAdminOk returns a tuple with the IsMandatorAdmin field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetIsMandatorAdminOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsMandatorAdmin, true
}

// SetIsMandatorAdmin sets field value
func (o *ClientConfiguration) SetIsMandatorAdmin(v bool) {
	o.IsMandatorAdmin = v
}

// GetIsWebScrapingEnabled returns the IsWebScrapingEnabled field value
func (o *ClientConfiguration) GetIsWebScrapingEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsWebScrapingEnabled
}

// GetIsWebScrapingEnabledOk returns a tuple with the IsWebScrapingEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetIsWebScrapingEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsWebScrapingEnabled, true
}

// SetIsWebScrapingEnabled sets field value
func (o *ClientConfiguration) SetIsWebScrapingEnabled(v bool) {
	o.IsWebScrapingEnabled = v
}

// GetIsXs2aEnabled returns the IsXs2aEnabled field value
func (o *ClientConfiguration) GetIsXs2aEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsXs2aEnabled
}

// GetIsXs2aEnabledOk returns a tuple with the IsXs2aEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetIsXs2aEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsXs2aEnabled, true
}

// SetIsXs2aEnabled sets field value
func (o *ClientConfiguration) SetIsXs2aEnabled(v bool) {
	o.IsXs2aEnabled = v
}

// GetPinStorageAvailableInWebForm returns the PinStorageAvailableInWebForm field value
func (o *ClientConfiguration) GetPinStorageAvailableInWebForm() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.PinStorageAvailableInWebForm
}

// GetPinStorageAvailableInWebFormOk returns a tuple with the PinStorageAvailableInWebForm field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetPinStorageAvailableInWebFormOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PinStorageAvailableInWebForm, true
}

// SetPinStorageAvailableInWebForm sets field value
func (o *ClientConfiguration) SetPinStorageAvailableInWebForm(v bool) {
	o.PinStorageAvailableInWebForm = v
}

// GetPaymentsEnabled returns the PaymentsEnabled field value
func (o *ClientConfiguration) GetPaymentsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.PaymentsEnabled
}

// GetPaymentsEnabledOk returns a tuple with the PaymentsEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetPaymentsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PaymentsEnabled, true
}

// SetPaymentsEnabled sets field value
func (o *ClientConfiguration) SetPaymentsEnabled(v bool) {
	o.PaymentsEnabled = v
}

// GetIsStandalonePaymentsEnabled returns the IsStandalonePaymentsEnabled field value
func (o *ClientConfiguration) GetIsStandalonePaymentsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsStandalonePaymentsEnabled
}

// GetIsStandalonePaymentsEnabledOk returns a tuple with the IsStandalonePaymentsEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetIsStandalonePaymentsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsStandalonePaymentsEnabled, true
}

// SetIsStandalonePaymentsEnabled sets field value
func (o *ClientConfiguration) SetIsStandalonePaymentsEnabled(v bool) {
	o.IsStandalonePaymentsEnabled = v
}

// GetAvailableBankGroups returns the AvailableBankGroups field value
func (o *ClientConfiguration) GetAvailableBankGroups() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.AvailableBankGroups
}

// GetAvailableBankGroupsOk returns a tuple with the AvailableBankGroups field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetAvailableBankGroupsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AvailableBankGroups, true
}

// SetAvailableBankGroups sets field value
func (o *ClientConfiguration) SetAvailableBankGroups(v []string) {
	o.AvailableBankGroups = v
}

// GetProducts returns the Products field value
func (o *ClientConfiguration) GetProducts() []Product {
	if o == nil  {
		var ret []Product
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetProductsOk() (*[]Product, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Products, true
}

// SetProducts sets field value
func (o *ClientConfiguration) SetProducts(v []Product) {
	o.Products = v
}

// GetApplicationName returns the ApplicationName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClientConfiguration) GetApplicationName() string {
	if o == nil || o.ApplicationName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ApplicationName.Get()
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientConfiguration) GetApplicationNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ApplicationName.Get(), o.ApplicationName.IsSet()
}

// SetApplicationName sets field value
func (o *ClientConfiguration) SetApplicationName(v string) {
	o.ApplicationName.Set(&v)
}

// GetFinTSProductRegistrationNumber returns the FinTSProductRegistrationNumber field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClientConfiguration) GetFinTSProductRegistrationNumber() string {
	if o == nil || o.FinTSProductRegistrationNumber.Get() == nil {
		var ret string
		return ret
	}

	return *o.FinTSProductRegistrationNumber.Get()
}

// GetFinTSProductRegistrationNumberOk returns a tuple with the FinTSProductRegistrationNumber field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientConfiguration) GetFinTSProductRegistrationNumberOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FinTSProductRegistrationNumber.Get(), o.FinTSProductRegistrationNumber.IsSet()
}

// SetFinTSProductRegistrationNumber sets field value
func (o *ClientConfiguration) SetFinTSProductRegistrationNumber(v string) {
	o.FinTSProductRegistrationNumber.Set(&v)
}

// GetStoreSecretsAvailableInWebForm returns the StoreSecretsAvailableInWebForm field value
func (o *ClientConfiguration) GetStoreSecretsAvailableInWebForm() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.StoreSecretsAvailableInWebForm
}

// GetStoreSecretsAvailableInWebFormOk returns a tuple with the StoreSecretsAvailableInWebForm field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetStoreSecretsAvailableInWebFormOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StoreSecretsAvailableInWebForm, true
}

// SetStoreSecretsAvailableInWebForm sets field value
func (o *ClientConfiguration) SetStoreSecretsAvailableInWebForm(v bool) {
	o.StoreSecretsAvailableInWebForm = v
}

// GetSupportSubjectDefault returns the SupportSubjectDefault field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClientConfiguration) GetSupportSubjectDefault() string {
	if o == nil || o.SupportSubjectDefault.Get() == nil {
		var ret string
		return ret
	}

	return *o.SupportSubjectDefault.Get()
}

// GetSupportSubjectDefaultOk returns a tuple with the SupportSubjectDefault field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientConfiguration) GetSupportSubjectDefaultOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SupportSubjectDefault.Get(), o.SupportSubjectDefault.IsSet()
}

// SetSupportSubjectDefault sets field value
func (o *ClientConfiguration) SetSupportSubjectDefault(v string) {
	o.SupportSubjectDefault.Set(&v)
}

// GetSupportEmail returns the SupportEmail field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClientConfiguration) GetSupportEmail() string {
	if o == nil || o.SupportEmail.Get() == nil {
		var ret string
		return ret
	}

	return *o.SupportEmail.Get()
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientConfiguration) GetSupportEmailOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.SupportEmail.Get(), o.SupportEmail.IsSet()
}

// SetSupportEmail sets field value
func (o *ClientConfiguration) SetSupportEmail(v string) {
	o.SupportEmail.Set(&v)
}

// GetAisWebFormMode returns the AisWebFormMode field value
func (o *ClientConfiguration) GetAisWebFormMode() WebFormMode {
	if o == nil  {
		var ret WebFormMode
		return ret
	}

	return o.AisWebFormMode
}

// GetAisWebFormModeOk returns a tuple with the AisWebFormMode field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetAisWebFormModeOk() (*WebFormMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AisWebFormMode, true
}

// SetAisWebFormMode sets field value
func (o *ClientConfiguration) SetAisWebFormMode(v WebFormMode) {
	o.AisWebFormMode = v
}

// GetPisWebFormMode returns the PisWebFormMode field value
func (o *ClientConfiguration) GetPisWebFormMode() WebFormMode {
	if o == nil  {
		var ret WebFormMode
		return ret
	}

	return o.PisWebFormMode
}

// GetPisWebFormModeOk returns a tuple with the PisWebFormMode field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetPisWebFormModeOk() (*WebFormMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PisWebFormMode, true
}

// SetPisWebFormMode sets field value
func (o *ClientConfiguration) SetPisWebFormMode(v WebFormMode) {
	o.PisWebFormMode = v
}

// GetPisStandaloneWebFormMode returns the PisStandaloneWebFormMode field value
func (o *ClientConfiguration) GetPisStandaloneWebFormMode() WebFormMode {
	if o == nil  {
		var ret WebFormMode
		return ret
	}

	return o.PisStandaloneWebFormMode
}

// GetPisStandaloneWebFormModeOk returns a tuple with the PisStandaloneWebFormMode field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetPisStandaloneWebFormModeOk() (*WebFormMode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PisStandaloneWebFormMode, true
}

// SetPisStandaloneWebFormMode sets field value
func (o *ClientConfiguration) SetPisStandaloneWebFormMode(v WebFormMode) {
	o.PisStandaloneWebFormMode = v
}

// GetBetaBanksEnabled returns the BetaBanksEnabled field value
func (o *ClientConfiguration) GetBetaBanksEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.BetaBanksEnabled
}

// GetBetaBanksEnabledOk returns a tuple with the BetaBanksEnabled field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetBetaBanksEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BetaBanksEnabled, true
}

// SetBetaBanksEnabled sets field value
func (o *ClientConfiguration) SetBetaBanksEnabled(v bool) {
	o.BetaBanksEnabled = v
}

// GetCategoryRestrictions returns the CategoryRestrictions field value
// If the value is explicit nil, the zero value for []Category will be returned
func (o *ClientConfiguration) GetCategoryRestrictions() []Category {
	if o == nil  {
		var ret []Category
		return ret
	}

	return o.CategoryRestrictions
}

// GetCategoryRestrictionsOk returns a tuple with the CategoryRestrictions field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientConfiguration) GetCategoryRestrictionsOk() (*[]Category, bool) {
	if o == nil || o.CategoryRestrictions == nil {
		return nil, false
	}
	return &o.CategoryRestrictions, true
}

// SetCategoryRestrictions sets field value
func (o *ClientConfiguration) SetCategoryRestrictions(v []Category) {
	o.CategoryRestrictions = v
}

// GetAutoDismountWebForm returns the AutoDismountWebForm field value
func (o *ClientConfiguration) GetAutoDismountWebForm() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.AutoDismountWebForm
}

// GetAutoDismountWebFormOk returns a tuple with the AutoDismountWebForm field value
// and a boolean to check if the value has been set.
func (o *ClientConfiguration) GetAutoDismountWebFormOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AutoDismountWebForm, true
}

// SetAutoDismountWebForm sets field value
func (o *ClientConfiguration) SetAutoDismountWebForm(v bool) {
	o.AutoDismountWebForm = v
}

// GetCorsAllowedOrigins returns the CorsAllowedOrigins field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ClientConfiguration) GetCorsAllowedOrigins() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.CorsAllowedOrigins
}

// GetCorsAllowedOriginsOk returns a tuple with the CorsAllowedOrigins field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientConfiguration) GetCorsAllowedOriginsOk() (*[]string, bool) {
	if o == nil || o.CorsAllowedOrigins == nil {
		return nil, false
	}
	return &o.CorsAllowedOrigins, true
}

// SetCorsAllowedOrigins sets field value
func (o *ClientConfiguration) SetCorsAllowedOrigins(v []string) {
	o.CorsAllowedOrigins = v
}

func (o ClientConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pfmServicesEnabled"] = o.PfmServicesEnabled
	}
	if true {
		toSerialize["isAutomaticBatchUpdateEnabled"] = o.IsAutomaticBatchUpdateEnabled
	}
	if true {
		toSerialize["isDevelopmentModeEnabled"] = o.IsDevelopmentModeEnabled
	}
	if true {
		toSerialize["isNonEuroAccountsSupported"] = o.IsNonEuroAccountsSupported
	}
	if true {
		toSerialize["isAutoCategorizationEnabled"] = o.IsAutoCategorizationEnabled
	}
	if true {
		toSerialize["mandatorLicense"] = o.MandatorLicense
	}
	if true {
		toSerialize["preferredConsentType"] = o.PreferredConsentType
	}
	if true {
		toSerialize["userNotificationCallbackUrl"] = o.UserNotificationCallbackUrl.Get()
	}
	if true {
		toSerialize["userSynchronizationCallbackUrl"] = o.UserSynchronizationCallbackUrl.Get()
	}
	if true {
		toSerialize["refreshTokensValidityPeriod"] = o.RefreshTokensValidityPeriod
	}
	if true {
		toSerialize["userAccessTokensValidityPeriod"] = o.UserAccessTokensValidityPeriod
	}
	if true {
		toSerialize["clientAccessTokensValidityPeriod"] = o.ClientAccessTokensValidityPeriod
	}
	if true {
		toSerialize["maxUserLoginAttempts"] = o.MaxUserLoginAttempts
	}
	if true {
		toSerialize["transactionImportLimitation"] = o.TransactionImportLimitation
	}
	if true {
		toSerialize["isUserAutoVerificationEnabled"] = o.IsUserAutoVerificationEnabled
	}
	if true {
		toSerialize["isMandatorAdmin"] = o.IsMandatorAdmin
	}
	if true {
		toSerialize["isWebScrapingEnabled"] = o.IsWebScrapingEnabled
	}
	if true {
		toSerialize["isXs2aEnabled"] = o.IsXs2aEnabled
	}
	if true {
		toSerialize["pinStorageAvailableInWebForm"] = o.PinStorageAvailableInWebForm
	}
	if true {
		toSerialize["paymentsEnabled"] = o.PaymentsEnabled
	}
	if true {
		toSerialize["isStandalonePaymentsEnabled"] = o.IsStandalonePaymentsEnabled
	}
	if true {
		toSerialize["availableBankGroups"] = o.AvailableBankGroups
	}
	if true {
		toSerialize["products"] = o.Products
	}
	if true {
		toSerialize["applicationName"] = o.ApplicationName.Get()
	}
	if true {
		toSerialize["finTSProductRegistrationNumber"] = o.FinTSProductRegistrationNumber.Get()
	}
	if true {
		toSerialize["storeSecretsAvailableInWebForm"] = o.StoreSecretsAvailableInWebForm
	}
	if true {
		toSerialize["supportSubjectDefault"] = o.SupportSubjectDefault.Get()
	}
	if true {
		toSerialize["supportEmail"] = o.SupportEmail.Get()
	}
	if true {
		toSerialize["aisWebFormMode"] = o.AisWebFormMode
	}
	if true {
		toSerialize["pisWebFormMode"] = o.PisWebFormMode
	}
	if true {
		toSerialize["pisStandaloneWebFormMode"] = o.PisStandaloneWebFormMode
	}
	if true {
		toSerialize["betaBanksEnabled"] = o.BetaBanksEnabled
	}
	if o.CategoryRestrictions != nil {
		toSerialize["categoryRestrictions"] = o.CategoryRestrictions
	}
	if true {
		toSerialize["autoDismountWebForm"] = o.AutoDismountWebForm
	}
	if o.CorsAllowedOrigins != nil {
		toSerialize["corsAllowedOrigins"] = o.CorsAllowedOrigins
	}
	return json.Marshal(toSerialize)
}

type NullableClientConfiguration struct {
	value *ClientConfiguration
	isSet bool
}

func (v NullableClientConfiguration) Get() *ClientConfiguration {
	return v.value
}

func (v *NullableClientConfiguration) Set(val *ClientConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableClientConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableClientConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientConfiguration(val *ClientConfiguration) *NullableClientConfiguration {
	return &NullableClientConfiguration{value: val, isSet: true}
}

func (v NullableClientConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


