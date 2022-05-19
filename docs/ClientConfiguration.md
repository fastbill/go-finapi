# ClientConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PfmServicesEnabled** | **bool** | Whether your client is allowed to call PFM services (Personal Finance Management). The set of PFM services is the following:&lt;br/&gt;&lt;br/&gt;&amp;bull; all /mandatorAdmin/ibanRules/_* and /mandatorAdmin/keywordRules/_* services&lt;br/&gt;&amp;bull; GET /accounts/dailyBalances&lt;br/&gt;&amp;bull; all /transactions/_* services, except for GET /transactions/[id(s)] and DELETE /transactions/[id]&lt;br/&gt;&amp;bull; all /categories/_* services, except for GET /categories/[id(s)]&lt;br/&gt;&amp;bull; all /labels/_* services&lt;br/&gt;&amp;bull; all /notificationRules/_* services&lt;br/&gt;&amp;bull; all /tests/_* services | 
**IsAutomaticBatchUpdateEnabled** | **bool** | Whether finAPI performs a regular automatic update of your users&#39; bank connections. To find out how the automatic batch update is configured for your client, i.e. which bank connections get updated, and at which time and interval, please contact your Sys-Admin. Note that even if the automatic batch update is enabled for your client, individual users can still disable the feature for their own bank connections. | 
**IsDevelopmentModeEnabled** | **bool** | Whether development mode is enabled. This setting is enabled on mandator level and allows any user to access the &#39;Mock batch update&#39; service. &lt;br/&gt;&lt;br/&gt;NOTE: This flag is meant for testing purposes during development of your application. &lt;br/&gt;This is why this will never be enabled on a production environment. | 
**IsNonEuroAccountsSupported** | **bool** | Whether finAPI will download data (balance and transactions) for bank accounts with a currency other than EUR (affects all users). If this flag is false, then non-EUR accounts will still be returned in the account list, but they will have no balance and no transactions. Note that this currently applies to Checking accounts only. | 
**IsAutoCategorizationEnabled** | **bool** | Whether transactions will be categorized as soon as they are downloaded. &lt;br/&gt;In case this flag is false, the user needs to manually trigger categorization using the &#39;Trigger categorization&#39; service. | 
**MandatorLicense** | [**MandatorLicense**](MandatorLicense.md) | &lt;strong&gt;Type:&lt;/strong&gt; MandatorLicense&lt;br/&gt; The license associated with your client. &lt;br/&gt;The licensing model affects the TPP registration data used to connect to the bank (e.g. &lt;b&gt;finTSProductRegistrationNumber&lt;/b&gt; for FINTS_SERVER interface). Licenses are administered by finAPI. Please contact the support to change the license that was set up for you.&lt;br/&gt;Possible values are:&lt;br/&gt;UNLICENSED: finAPI will use its own TPP registration to connect to the bank for both account information services (AIS) and payment initiation services (PIS).&lt;br/&gt;AISP: finAPI will use its own TPP registration to connect to the bank for PIS, and your registration for AIS.&lt;br/&gt;PISP: finAPI will use its own TPP registration to connect to the bank for AIS, and your registration for PIS.&lt;br/&gt;FULLY_LICENSED: finAPI will use your TPP registration to connect to the bank for both AIS and PIS. | 
**PreferredConsentType** | [**PreferredConsentType**](PreferredConsentType.md) | &lt;strong&gt;Type:&lt;/strong&gt; PreferredConsentType&lt;br/&gt; The preferred consent type that will be used for the XS2A interface.&lt;br/&gt;&lt;br/&gt;&lt;b&gt;ONETIME&lt;/b&gt; - The consent can only be used once to download data associated with the account. The consent won’t be saved by finAPI.&lt;br/&gt;&lt;b&gt;RECURRING&lt;/b&gt; - The consent is valid for up to 90 days and can be used by finAPI to access and download account data for up to 4 times per day.&lt;br/&gt;&lt;br/&gt;NOTE: If the bank does not support the preferred consent type, then finAPI will default to the other type. | 
**UserNotificationCallbackUrl** | **NullableString** | Callback URL to which finAPI sends the notification messages that are triggered from the automatic batch update of the users&#39; bank connections. This field is only relevant if the automatic batch update is enabled for your client. For details about what the notification messages look like, please see the documentation in the &#39;Notification Rules&#39; section. finAPI will call this URL with HTTP method POST. Note that the response of the call is not processed by finAPI. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha environment, it MUST be a SSL-secured (https) URL on the finAPI live system. | 
**UserSynchronizationCallbackUrl** | **NullableString** | Callback URL for user synchronization. This field should be set if you - as a finAPI customer - have multiple clients using finAPI. In such case, all of your clients will share the same user base, making it possible for a user to be created in one client, but then deleted in another. To keep the client-side user data consistent in all clients, you should set a callback URL for each client. finAPI will send a notification to the callback URL of each client whenever a user of your user base gets deleted. Note that finAPI will send a deletion notification to ALL clients, including the one that made the user deletion request to finAPI. So when deleting a user in finAPI, a client should rely on the callback to delete the user on its own side. &lt;p&gt;The notification that finAPI sends to the clients&#39; callback URLs will be a POST request, with this body: &lt;pre&gt;{    \&quot;userId\&quot; : string // contains the identifier of the deleted user    \&quot;event\&quot; : string // this will always be \&quot;DELETED\&quot; }&lt;/pre&gt;&lt;br/&gt;Note that finAPI does not process the response of this call. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha environment, it MUST be a SSL-secured (https) URL on the finAPI live system.&lt;/p&gt;As long as you have just one client, you can ignore this field and let it be null. However keep in mind that in this case your client will not receive any callback when a user gets deleted - so the deletion of the user on the client-side must not be forgotten. Of course you may still use the callback URL even for just one client, if you want to implement the deletion of the user on the client-side via the callback from finAPI. | 
**RefreshTokensValidityPeriod** | **int32** | The validity period that newly requested refresh tokens initially have (in seconds). A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation, or when a user gets locked, or when the password is reset for a user). | 
**UserAccessTokensValidityPeriod** | **int32** | The validity period that newly requested access tokens for users initially have (in seconds). A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation, or when a user gets locked, or when the password is reset for a user). | 
**ClientAccessTokensValidityPeriod** | **int32** | The validity period that newly requested access tokens for clients initially have (in seconds). A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation). | 
**MaxUserLoginAttempts** | **int32** | Number of consecutive failed login attempts of a user into his finAPI account that is allowed before finAPI locks the user&#39;s account. When a user&#39;s account is locked, finAPI will invalidate all user&#39;s tokens and it will deny any service call in the context of this user (i.e. any call to a service using one of the user&#39;s authorization tokens, as well as the service for requesting a new token for this user). To unlock a user&#39;s account, a new password must be set for the account by the client (see the services /users/requestPasswordChange and /users/executePasswordChange). Once a new password has been set, all services will be available again for this user and the user&#39;s failed login attempts counter is reset to 0. The user&#39;s failed login attempts counter is also reset whenever a new authorization token has been successfully retrieved, or whenever the user himself changes his password.&lt;br/&gt;&lt;br/&gt;Note that when this field has a value of 0, it means that there is no limit for user login attempts, i.e. finAPI will never lock user accounts. | 
**TransactionImportLimitation** | **int32** | This setting defines the upper limit of how much of an account&#39;s transactions history may be downloaded whenever a new account is imported, across all of your users. More technically, it depicts the maximum number of days for which transactions might get downloaded, starting from - and including - the date of the account import. &#39;0&#39; means that there is no limitation. | 
**IsUserAutoVerificationEnabled** | **bool** | Whether users that are created with this client are automatically verified on creation. If this field is set to &#39;false&#39;, then any user that is created with this client must first be verified with the \&quot;Verify a user\&quot; service before he can be authorized. If the field is &#39;true&#39;, then no verification is required by the client and the user can be authorized immediately after creation. | 
**IsMandatorAdmin** | **bool** | Whether this client is a &#39;Mandator Admin&#39;. Mandator Admins are special clients that can access the &#39;Mandator Administration&#39; section of finAPI. If you do not yet have credentials for a Mandator Admin, please contact us at support@finapi.io. For further information, please refer to &lt;a href&#x3D;&#39;https://documentation.finapi.io/access/Application-management.2763423767.html&#39; target&#x3D;&#39;_blank&#39;&gt;this page&lt;/a&gt; on our Access Public Documentation. | 
**IsWebScrapingEnabled** | **bool** | Whether finAPI is allowed to use the WEB_SCRAPER interface for data download or payments. &lt;br/&gt;&lt;br/&gt;If this field is set to &#39;true&#39;, then finAPI might download data from the online banking websites of banks (either in addition to other interfaces, or as the sole data source for the download). Also, it will be possible to do payments via the WEB_SCRAPER interface.&lt;br/&gt;&lt;br/&gt;If this field is set to &#39;false&#39;, then finAPI will not use any web scrapers. Payments via the WEB_SCRAPER interface will not be possible, and finAPI will not allow any data download for banks where no other interface except WEB_SCRAPER is available. &lt;br/&gt;&lt;br/&gt;Please contact your Sys-Admin if you want to change this setting. | 
**IsXs2aEnabled** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;&lt;br/&gt;Whether this client is allowed to access XS2A services. | 
**PinStorageAvailableInWebForm** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer to the &#39;storeSecretsAvailableInWebForm&#39; field instead.&lt;br/&gt;&lt;br/&gt;Whether finAPI&#39;s Web Form will provide a checkbox for the user allowing him to choose whether to store his banking PIN in finAPI. If this field is set to false, then the user won&#39;t have an option to store his PIN. | 
**PaymentsEnabled** | **bool** | Whether this client is allowed to do PIS.&lt;br/&gt;&lt;br/&gt;Note that on the Sandbox environment, it is always possible to execute payments (regardless of what this field says), as long as you are using a test bank (see Bank.isTestBank) | 
**IsStandalonePaymentsEnabled** | **bool** | Whether this client is allowed to do standalone PIS (doing money transfers and standing orders for accounts that are not imported in finAPI).&lt;br/&gt;&lt;br/&gt;Note that on the Sandbox environment, it is always possible to execute payments and standing orders (regardless of what this field says), as long as you are using a test bank (see Bank.isTestBank) | 
**AvailableBankGroups** | **[]string** |  | 
**Products** | [**[]Product**](Product.md) |  | 
**ApplicationName** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/profiles&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;logo&#39; and &#39;favicon&#39; field instead. The Web Form will now be able to render your logo both in the header and as a favicon.&lt;br/&gt;&lt;br/&gt;Application name. When an application name is set (e.g. \&quot;My App\&quot;), then finAPI&#39;s Web Form will display a text to the user \&quot;Weiterleitung auf finAPI von ...\&quot; (e.g. \&quot;Weiterleitung auf finAPI von MyApp\&quot;). | 
**FinTSProductRegistrationNumber** | **NullableString** | The FinTS product registration number. If a value is stored, this will always be &#39;XXXXX&#39;. | 
**StoreSecretsAvailableInWebForm** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/profiles&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;storeSecrets&#39; field instead.&lt;br/&gt;&lt;br/&gt;Whether finAPI&#39;s Web Form should provide a checkbox for the user allowing him to choose whether to store his banking PIN in finAPI. If this field is set to false, then the user won&#39;t have an option to store his PIN. | 
**SupportSubjectDefault** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/For-best-results!.2477654019.html&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;errorRedirectUrl&#39; and &#39;customerSupportUrl&#39; query parameters instead.&lt;br/&gt;&lt;br/&gt;Default value for the subject element of support emails. | 
**SupportEmail** | **NullableString** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/For-best-results!.2477654019.html&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;errorRedirectUrl&#39; and &#39;customerSupportUrl&#39; query parameters instead.&lt;br/&gt;&lt;br/&gt;Email address to sent support requests to from the Web Form. | 
**AisWebFormMode** | [**WebFormMode**](WebFormMode.md) | &lt;strong&gt;Type:&lt;/strong&gt; WebFormMode&lt;br/&gt; Indicates whether the client is using the finAPI Web Form for Account Initiation Services.&lt;br/&gt;&lt;br/&gt;Possible values: &lt;br/&gt;&amp;bull; &lt;code&gt;DISABLED&lt;/code&gt; - No Web Form is triggered&lt;br/&gt;&amp;bull; &lt;code&gt;INTERNAL&lt;/code&gt; - THIS VALUE IS DEPRECATED AND WILL BE REMOVED. Hence, we request customers to foresee a migration to Web Form 2.0 (value &lt;code&gt;EXTERNAL&lt;/code&gt;).&lt;br/&gt;End users will be directed to the classical Web Form implementation.&lt;br/&gt;&amp;bull; &lt;code&gt;EXTERNAL&lt;/code&gt; - End users will be directed to the &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/Introduction.2038136860.html&#39;  target&#x3D;&#39;_blank&#39;&gt;new Web Form&lt;/a&gt; implementation. | 
**PisWebFormMode** | [**WebFormMode**](WebFormMode.md) | &lt;strong&gt;Type:&lt;/strong&gt; WebFormMode&lt;br/&gt; Indicates whether the client is using the finAPI Web Form for Standard Payment Initiation Services (Payments for accounts that have been imported in finAPI).&lt;br/&gt;&lt;br/&gt;Possible values: &lt;br/&gt;&amp;bull; &lt;code&gt;DISABLED&lt;/code&gt; - No Web Form is triggered&lt;br/&gt;&amp;bull; &lt;code&gt;INTERNAL&lt;/code&gt; - THIS VALUE IS DEPRECATED AND WILL BE REMOVED. Hence, we request customers to foresee a migration to Web Form 2.0 (value &lt;code&gt;EXTERNAL&lt;/code&gt;).&lt;br/&gt;End users will be directed to the classical Web Form implementation.&lt;br/&gt;&amp;bull; &lt;code&gt;EXTERNAL&lt;/code&gt; - End users will be directed to the &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/Introduction.2038136860.html&#39;  target&#x3D;&#39;_blank&#39;&gt;new Web Form&lt;/a&gt; implementation. | 
**PisStandaloneWebFormMode** | [**WebFormMode**](WebFormMode.md) | &lt;strong&gt;Type:&lt;/strong&gt; WebFormMode&lt;br/&gt; Indicates whether the client is using the finAPI Web Form for Standalone Payment Initiation Services (Payments without account import).&lt;br/&gt;&lt;br/&gt;Possible values: &lt;br/&gt;&amp;bull; &lt;code&gt;DISABLED&lt;/code&gt; - No Web Form is triggered&lt;br/&gt;&amp;bull; &lt;code&gt;INTERNAL&lt;/code&gt; - THIS VALUE IS DEPRECATED AND WILL BE REMOVED. Hence, we request customers to foresee a migration to Web Form 2.0 (value &lt;code&gt;EXTERNAL&lt;/code&gt;).&lt;br/&gt;End users will be directed to the classical Web Form implementation.&lt;br/&gt;&amp;bull; &lt;code&gt;EXTERNAL&lt;/code&gt; - End users will be directed to the &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/Introduction.2038136860.html&#39;  target&#x3D;&#39;_blank&#39;&gt;new Web Form&lt;/a&gt; implementation. | 
**BetaBanksEnabled** | **bool** | Whether the set of banks that are available to your client contains “Beta banks”. Beta banks provide pre-release interfaces that are still in a beta phase. Communication to the bank via such interfaces might be unstable, and the correctness and/or quality of data delivery or payment execution cannot be guaranteed.&lt;br/&gt;As the word “BETA” already indicates, Beta banks are subject to changes. Their properties, as well as their behaviour can change based on continuous tests and customer feedback. Also, to keep our bank list clean, we might remove Beta banks at any point in time, including all related user data (bank connections, accounts, transactions etc). We still recommend you to enable beta banks in your application, because it enables us to release a stable interface faster. However, you should point it out to your users when using a beta bank (also see field Bank.isBeta).&lt;br/&gt;&lt;br/&gt;If this field is true, then the GET /banks services will include beta banks in their results, and you can use beta banks in any service where you can pass a bank identifier. If the field is false, then beta banks will not exist for your client. | 
**CategoryRestrictions** | [**[]Category**](Category.md) | &lt;strong&gt;Type:&lt;/strong&gt; Category&lt;br/&gt; Defines the set of transaction categories to which your client is restricted. When retrieving transactions (via the GET /transactions services), you may request only those transactions whose &#39;category&#39; is one of the listed categories. If this field is null, then there are no restrictions for your client, and you may retrieve the full set of imported transactions. | 
**AutoDismountWebForm** | **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt; Please refer &lt;a href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/profiles&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;skipConfirmationView&#39; field instead.&lt;br/&gt;&lt;br/&gt;This flag indicates whether the Web Form should get removed from the parent page automatically once it’s finished. It applies ONLY to the classical embedded Web Form. That means it’s only applied if aisWebFormMode, pisWebFormMode or pisStandaloneWebFormMode are defined as INTERNAL. In case you are using our standalone Web Form by redirecting the user to our Web Form link, this feature has no effect. | 
**CorsAllowedOrigins** | **[]string** | The list of allowed origins for cross-origin requests. The CORS configuration applies to all the API services except for the /oauth services. If this list is empty, then CORS is not enabled for this client. Please contact the support if you want to enable or change the client&#39;s CORS configuration. | 

## Methods

### NewClientConfiguration

`func NewClientConfiguration(pfmServicesEnabled bool, isAutomaticBatchUpdateEnabled bool, isDevelopmentModeEnabled bool, isNonEuroAccountsSupported bool, isAutoCategorizationEnabled bool, mandatorLicense MandatorLicense, preferredConsentType PreferredConsentType, userNotificationCallbackUrl NullableString, userSynchronizationCallbackUrl NullableString, refreshTokensValidityPeriod int32, userAccessTokensValidityPeriod int32, clientAccessTokensValidityPeriod int32, maxUserLoginAttempts int32, transactionImportLimitation int32, isUserAutoVerificationEnabled bool, isMandatorAdmin bool, isWebScrapingEnabled bool, isXs2aEnabled bool, pinStorageAvailableInWebForm bool, paymentsEnabled bool, isStandalonePaymentsEnabled bool, availableBankGroups []string, products []Product, applicationName NullableString, finTSProductRegistrationNumber NullableString, storeSecretsAvailableInWebForm bool, supportSubjectDefault NullableString, supportEmail NullableString, aisWebFormMode WebFormMode, pisWebFormMode WebFormMode, pisStandaloneWebFormMode WebFormMode, betaBanksEnabled bool, categoryRestrictions []Category, autoDismountWebForm bool, corsAllowedOrigins []string, ) *ClientConfiguration`

NewClientConfiguration instantiates a new ClientConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfigurationWithDefaults

`func NewClientConfigurationWithDefaults() *ClientConfiguration`

NewClientConfigurationWithDefaults instantiates a new ClientConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPfmServicesEnabled

`func (o *ClientConfiguration) GetPfmServicesEnabled() bool`

GetPfmServicesEnabled returns the PfmServicesEnabled field if non-nil, zero value otherwise.

### GetPfmServicesEnabledOk

`func (o *ClientConfiguration) GetPfmServicesEnabledOk() (*bool, bool)`

GetPfmServicesEnabledOk returns a tuple with the PfmServicesEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfmServicesEnabled

`func (o *ClientConfiguration) SetPfmServicesEnabled(v bool)`

SetPfmServicesEnabled sets PfmServicesEnabled field to given value.


### GetIsAutomaticBatchUpdateEnabled

`func (o *ClientConfiguration) GetIsAutomaticBatchUpdateEnabled() bool`

GetIsAutomaticBatchUpdateEnabled returns the IsAutomaticBatchUpdateEnabled field if non-nil, zero value otherwise.

### GetIsAutomaticBatchUpdateEnabledOk

`func (o *ClientConfiguration) GetIsAutomaticBatchUpdateEnabledOk() (*bool, bool)`

GetIsAutomaticBatchUpdateEnabledOk returns a tuple with the IsAutomaticBatchUpdateEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutomaticBatchUpdateEnabled

`func (o *ClientConfiguration) SetIsAutomaticBatchUpdateEnabled(v bool)`

SetIsAutomaticBatchUpdateEnabled sets IsAutomaticBatchUpdateEnabled field to given value.


### GetIsDevelopmentModeEnabled

`func (o *ClientConfiguration) GetIsDevelopmentModeEnabled() bool`

GetIsDevelopmentModeEnabled returns the IsDevelopmentModeEnabled field if non-nil, zero value otherwise.

### GetIsDevelopmentModeEnabledOk

`func (o *ClientConfiguration) GetIsDevelopmentModeEnabledOk() (*bool, bool)`

GetIsDevelopmentModeEnabledOk returns a tuple with the IsDevelopmentModeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDevelopmentModeEnabled

`func (o *ClientConfiguration) SetIsDevelopmentModeEnabled(v bool)`

SetIsDevelopmentModeEnabled sets IsDevelopmentModeEnabled field to given value.


### GetIsNonEuroAccountsSupported

`func (o *ClientConfiguration) GetIsNonEuroAccountsSupported() bool`

GetIsNonEuroAccountsSupported returns the IsNonEuroAccountsSupported field if non-nil, zero value otherwise.

### GetIsNonEuroAccountsSupportedOk

`func (o *ClientConfiguration) GetIsNonEuroAccountsSupportedOk() (*bool, bool)`

GetIsNonEuroAccountsSupportedOk returns a tuple with the IsNonEuroAccountsSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNonEuroAccountsSupported

`func (o *ClientConfiguration) SetIsNonEuroAccountsSupported(v bool)`

SetIsNonEuroAccountsSupported sets IsNonEuroAccountsSupported field to given value.


### GetIsAutoCategorizationEnabled

`func (o *ClientConfiguration) GetIsAutoCategorizationEnabled() bool`

GetIsAutoCategorizationEnabled returns the IsAutoCategorizationEnabled field if non-nil, zero value otherwise.

### GetIsAutoCategorizationEnabledOk

`func (o *ClientConfiguration) GetIsAutoCategorizationEnabledOk() (*bool, bool)`

GetIsAutoCategorizationEnabledOk returns a tuple with the IsAutoCategorizationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAutoCategorizationEnabled

`func (o *ClientConfiguration) SetIsAutoCategorizationEnabled(v bool)`

SetIsAutoCategorizationEnabled sets IsAutoCategorizationEnabled field to given value.


### GetMandatorLicense

`func (o *ClientConfiguration) GetMandatorLicense() MandatorLicense`

GetMandatorLicense returns the MandatorLicense field if non-nil, zero value otherwise.

### GetMandatorLicenseOk

`func (o *ClientConfiguration) GetMandatorLicenseOk() (*MandatorLicense, bool)`

GetMandatorLicenseOk returns a tuple with the MandatorLicense field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMandatorLicense

`func (o *ClientConfiguration) SetMandatorLicense(v MandatorLicense)`

SetMandatorLicense sets MandatorLicense field to given value.


### GetPreferredConsentType

`func (o *ClientConfiguration) GetPreferredConsentType() PreferredConsentType`

GetPreferredConsentType returns the PreferredConsentType field if non-nil, zero value otherwise.

### GetPreferredConsentTypeOk

`func (o *ClientConfiguration) GetPreferredConsentTypeOk() (*PreferredConsentType, bool)`

GetPreferredConsentTypeOk returns a tuple with the PreferredConsentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredConsentType

`func (o *ClientConfiguration) SetPreferredConsentType(v PreferredConsentType)`

SetPreferredConsentType sets PreferredConsentType field to given value.


### GetUserNotificationCallbackUrl

`func (o *ClientConfiguration) GetUserNotificationCallbackUrl() string`

GetUserNotificationCallbackUrl returns the UserNotificationCallbackUrl field if non-nil, zero value otherwise.

### GetUserNotificationCallbackUrlOk

`func (o *ClientConfiguration) GetUserNotificationCallbackUrlOk() (*string, bool)`

GetUserNotificationCallbackUrlOk returns a tuple with the UserNotificationCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNotificationCallbackUrl

`func (o *ClientConfiguration) SetUserNotificationCallbackUrl(v string)`

SetUserNotificationCallbackUrl sets UserNotificationCallbackUrl field to given value.


### SetUserNotificationCallbackUrlNil

`func (o *ClientConfiguration) SetUserNotificationCallbackUrlNil(b bool)`

 SetUserNotificationCallbackUrlNil sets the value for UserNotificationCallbackUrl to be an explicit nil

### UnsetUserNotificationCallbackUrl
`func (o *ClientConfiguration) UnsetUserNotificationCallbackUrl()`

UnsetUserNotificationCallbackUrl ensures that no value is present for UserNotificationCallbackUrl, not even an explicit nil
### GetUserSynchronizationCallbackUrl

`func (o *ClientConfiguration) GetUserSynchronizationCallbackUrl() string`

GetUserSynchronizationCallbackUrl returns the UserSynchronizationCallbackUrl field if non-nil, zero value otherwise.

### GetUserSynchronizationCallbackUrlOk

`func (o *ClientConfiguration) GetUserSynchronizationCallbackUrlOk() (*string, bool)`

GetUserSynchronizationCallbackUrlOk returns a tuple with the UserSynchronizationCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSynchronizationCallbackUrl

`func (o *ClientConfiguration) SetUserSynchronizationCallbackUrl(v string)`

SetUserSynchronizationCallbackUrl sets UserSynchronizationCallbackUrl field to given value.


### SetUserSynchronizationCallbackUrlNil

`func (o *ClientConfiguration) SetUserSynchronizationCallbackUrlNil(b bool)`

 SetUserSynchronizationCallbackUrlNil sets the value for UserSynchronizationCallbackUrl to be an explicit nil

### UnsetUserSynchronizationCallbackUrl
`func (o *ClientConfiguration) UnsetUserSynchronizationCallbackUrl()`

UnsetUserSynchronizationCallbackUrl ensures that no value is present for UserSynchronizationCallbackUrl, not even an explicit nil
### GetRefreshTokensValidityPeriod

`func (o *ClientConfiguration) GetRefreshTokensValidityPeriod() int32`

GetRefreshTokensValidityPeriod returns the RefreshTokensValidityPeriod field if non-nil, zero value otherwise.

### GetRefreshTokensValidityPeriodOk

`func (o *ClientConfiguration) GetRefreshTokensValidityPeriodOk() (*int32, bool)`

GetRefreshTokensValidityPeriodOk returns a tuple with the RefreshTokensValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokensValidityPeriod

`func (o *ClientConfiguration) SetRefreshTokensValidityPeriod(v int32)`

SetRefreshTokensValidityPeriod sets RefreshTokensValidityPeriod field to given value.


### GetUserAccessTokensValidityPeriod

`func (o *ClientConfiguration) GetUserAccessTokensValidityPeriod() int32`

GetUserAccessTokensValidityPeriod returns the UserAccessTokensValidityPeriod field if non-nil, zero value otherwise.

### GetUserAccessTokensValidityPeriodOk

`func (o *ClientConfiguration) GetUserAccessTokensValidityPeriodOk() (*int32, bool)`

GetUserAccessTokensValidityPeriodOk returns a tuple with the UserAccessTokensValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessTokensValidityPeriod

`func (o *ClientConfiguration) SetUserAccessTokensValidityPeriod(v int32)`

SetUserAccessTokensValidityPeriod sets UserAccessTokensValidityPeriod field to given value.


### GetClientAccessTokensValidityPeriod

`func (o *ClientConfiguration) GetClientAccessTokensValidityPeriod() int32`

GetClientAccessTokensValidityPeriod returns the ClientAccessTokensValidityPeriod field if non-nil, zero value otherwise.

### GetClientAccessTokensValidityPeriodOk

`func (o *ClientConfiguration) GetClientAccessTokensValidityPeriodOk() (*int32, bool)`

GetClientAccessTokensValidityPeriodOk returns a tuple with the ClientAccessTokensValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAccessTokensValidityPeriod

`func (o *ClientConfiguration) SetClientAccessTokensValidityPeriod(v int32)`

SetClientAccessTokensValidityPeriod sets ClientAccessTokensValidityPeriod field to given value.


### GetMaxUserLoginAttempts

`func (o *ClientConfiguration) GetMaxUserLoginAttempts() int32`

GetMaxUserLoginAttempts returns the MaxUserLoginAttempts field if non-nil, zero value otherwise.

### GetMaxUserLoginAttemptsOk

`func (o *ClientConfiguration) GetMaxUserLoginAttemptsOk() (*int32, bool)`

GetMaxUserLoginAttemptsOk returns a tuple with the MaxUserLoginAttempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUserLoginAttempts

`func (o *ClientConfiguration) SetMaxUserLoginAttempts(v int32)`

SetMaxUserLoginAttempts sets MaxUserLoginAttempts field to given value.


### GetTransactionImportLimitation

`func (o *ClientConfiguration) GetTransactionImportLimitation() int32`

GetTransactionImportLimitation returns the TransactionImportLimitation field if non-nil, zero value otherwise.

### GetTransactionImportLimitationOk

`func (o *ClientConfiguration) GetTransactionImportLimitationOk() (*int32, bool)`

GetTransactionImportLimitationOk returns a tuple with the TransactionImportLimitation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionImportLimitation

`func (o *ClientConfiguration) SetTransactionImportLimitation(v int32)`

SetTransactionImportLimitation sets TransactionImportLimitation field to given value.


### GetIsUserAutoVerificationEnabled

`func (o *ClientConfiguration) GetIsUserAutoVerificationEnabled() bool`

GetIsUserAutoVerificationEnabled returns the IsUserAutoVerificationEnabled field if non-nil, zero value otherwise.

### GetIsUserAutoVerificationEnabledOk

`func (o *ClientConfiguration) GetIsUserAutoVerificationEnabledOk() (*bool, bool)`

GetIsUserAutoVerificationEnabledOk returns a tuple with the IsUserAutoVerificationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUserAutoVerificationEnabled

`func (o *ClientConfiguration) SetIsUserAutoVerificationEnabled(v bool)`

SetIsUserAutoVerificationEnabled sets IsUserAutoVerificationEnabled field to given value.


### GetIsMandatorAdmin

`func (o *ClientConfiguration) GetIsMandatorAdmin() bool`

GetIsMandatorAdmin returns the IsMandatorAdmin field if non-nil, zero value otherwise.

### GetIsMandatorAdminOk

`func (o *ClientConfiguration) GetIsMandatorAdminOk() (*bool, bool)`

GetIsMandatorAdminOk returns a tuple with the IsMandatorAdmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMandatorAdmin

`func (o *ClientConfiguration) SetIsMandatorAdmin(v bool)`

SetIsMandatorAdmin sets IsMandatorAdmin field to given value.


### GetIsWebScrapingEnabled

`func (o *ClientConfiguration) GetIsWebScrapingEnabled() bool`

GetIsWebScrapingEnabled returns the IsWebScrapingEnabled field if non-nil, zero value otherwise.

### GetIsWebScrapingEnabledOk

`func (o *ClientConfiguration) GetIsWebScrapingEnabledOk() (*bool, bool)`

GetIsWebScrapingEnabledOk returns a tuple with the IsWebScrapingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsWebScrapingEnabled

`func (o *ClientConfiguration) SetIsWebScrapingEnabled(v bool)`

SetIsWebScrapingEnabled sets IsWebScrapingEnabled field to given value.


### GetIsXs2aEnabled

`func (o *ClientConfiguration) GetIsXs2aEnabled() bool`

GetIsXs2aEnabled returns the IsXs2aEnabled field if non-nil, zero value otherwise.

### GetIsXs2aEnabledOk

`func (o *ClientConfiguration) GetIsXs2aEnabledOk() (*bool, bool)`

GetIsXs2aEnabledOk returns a tuple with the IsXs2aEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsXs2aEnabled

`func (o *ClientConfiguration) SetIsXs2aEnabled(v bool)`

SetIsXs2aEnabled sets IsXs2aEnabled field to given value.


### GetPinStorageAvailableInWebForm

`func (o *ClientConfiguration) GetPinStorageAvailableInWebForm() bool`

GetPinStorageAvailableInWebForm returns the PinStorageAvailableInWebForm field if non-nil, zero value otherwise.

### GetPinStorageAvailableInWebFormOk

`func (o *ClientConfiguration) GetPinStorageAvailableInWebFormOk() (*bool, bool)`

GetPinStorageAvailableInWebFormOk returns a tuple with the PinStorageAvailableInWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinStorageAvailableInWebForm

`func (o *ClientConfiguration) SetPinStorageAvailableInWebForm(v bool)`

SetPinStorageAvailableInWebForm sets PinStorageAvailableInWebForm field to given value.


### GetPaymentsEnabled

`func (o *ClientConfiguration) GetPaymentsEnabled() bool`

GetPaymentsEnabled returns the PaymentsEnabled field if non-nil, zero value otherwise.

### GetPaymentsEnabledOk

`func (o *ClientConfiguration) GetPaymentsEnabledOk() (*bool, bool)`

GetPaymentsEnabledOk returns a tuple with the PaymentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentsEnabled

`func (o *ClientConfiguration) SetPaymentsEnabled(v bool)`

SetPaymentsEnabled sets PaymentsEnabled field to given value.


### GetIsStandalonePaymentsEnabled

`func (o *ClientConfiguration) GetIsStandalonePaymentsEnabled() bool`

GetIsStandalonePaymentsEnabled returns the IsStandalonePaymentsEnabled field if non-nil, zero value otherwise.

### GetIsStandalonePaymentsEnabledOk

`func (o *ClientConfiguration) GetIsStandalonePaymentsEnabledOk() (*bool, bool)`

GetIsStandalonePaymentsEnabledOk returns a tuple with the IsStandalonePaymentsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStandalonePaymentsEnabled

`func (o *ClientConfiguration) SetIsStandalonePaymentsEnabled(v bool)`

SetIsStandalonePaymentsEnabled sets IsStandalonePaymentsEnabled field to given value.


### GetAvailableBankGroups

`func (o *ClientConfiguration) GetAvailableBankGroups() []string`

GetAvailableBankGroups returns the AvailableBankGroups field if non-nil, zero value otherwise.

### GetAvailableBankGroupsOk

`func (o *ClientConfiguration) GetAvailableBankGroupsOk() (*[]string, bool)`

GetAvailableBankGroupsOk returns a tuple with the AvailableBankGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBankGroups

`func (o *ClientConfiguration) SetAvailableBankGroups(v []string)`

SetAvailableBankGroups sets AvailableBankGroups field to given value.


### GetProducts

`func (o *ClientConfiguration) GetProducts() []Product`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *ClientConfiguration) GetProductsOk() (*[]Product, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *ClientConfiguration) SetProducts(v []Product)`

SetProducts sets Products field to given value.


### GetApplicationName

`func (o *ClientConfiguration) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ClientConfiguration) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ClientConfiguration) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.


### SetApplicationNameNil

`func (o *ClientConfiguration) SetApplicationNameNil(b bool)`

 SetApplicationNameNil sets the value for ApplicationName to be an explicit nil

### UnsetApplicationName
`func (o *ClientConfiguration) UnsetApplicationName()`

UnsetApplicationName ensures that no value is present for ApplicationName, not even an explicit nil
### GetFinTSProductRegistrationNumber

`func (o *ClientConfiguration) GetFinTSProductRegistrationNumber() string`

GetFinTSProductRegistrationNumber returns the FinTSProductRegistrationNumber field if non-nil, zero value otherwise.

### GetFinTSProductRegistrationNumberOk

`func (o *ClientConfiguration) GetFinTSProductRegistrationNumberOk() (*string, bool)`

GetFinTSProductRegistrationNumberOk returns a tuple with the FinTSProductRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinTSProductRegistrationNumber

`func (o *ClientConfiguration) SetFinTSProductRegistrationNumber(v string)`

SetFinTSProductRegistrationNumber sets FinTSProductRegistrationNumber field to given value.


### SetFinTSProductRegistrationNumberNil

`func (o *ClientConfiguration) SetFinTSProductRegistrationNumberNil(b bool)`

 SetFinTSProductRegistrationNumberNil sets the value for FinTSProductRegistrationNumber to be an explicit nil

### UnsetFinTSProductRegistrationNumber
`func (o *ClientConfiguration) UnsetFinTSProductRegistrationNumber()`

UnsetFinTSProductRegistrationNumber ensures that no value is present for FinTSProductRegistrationNumber, not even an explicit nil
### GetStoreSecretsAvailableInWebForm

`func (o *ClientConfiguration) GetStoreSecretsAvailableInWebForm() bool`

GetStoreSecretsAvailableInWebForm returns the StoreSecretsAvailableInWebForm field if non-nil, zero value otherwise.

### GetStoreSecretsAvailableInWebFormOk

`func (o *ClientConfiguration) GetStoreSecretsAvailableInWebFormOk() (*bool, bool)`

GetStoreSecretsAvailableInWebFormOk returns a tuple with the StoreSecretsAvailableInWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSecretsAvailableInWebForm

`func (o *ClientConfiguration) SetStoreSecretsAvailableInWebForm(v bool)`

SetStoreSecretsAvailableInWebForm sets StoreSecretsAvailableInWebForm field to given value.


### GetSupportSubjectDefault

`func (o *ClientConfiguration) GetSupportSubjectDefault() string`

GetSupportSubjectDefault returns the SupportSubjectDefault field if non-nil, zero value otherwise.

### GetSupportSubjectDefaultOk

`func (o *ClientConfiguration) GetSupportSubjectDefaultOk() (*string, bool)`

GetSupportSubjectDefaultOk returns a tuple with the SupportSubjectDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSubjectDefault

`func (o *ClientConfiguration) SetSupportSubjectDefault(v string)`

SetSupportSubjectDefault sets SupportSubjectDefault field to given value.


### SetSupportSubjectDefaultNil

`func (o *ClientConfiguration) SetSupportSubjectDefaultNil(b bool)`

 SetSupportSubjectDefaultNil sets the value for SupportSubjectDefault to be an explicit nil

### UnsetSupportSubjectDefault
`func (o *ClientConfiguration) UnsetSupportSubjectDefault()`

UnsetSupportSubjectDefault ensures that no value is present for SupportSubjectDefault, not even an explicit nil
### GetSupportEmail

`func (o *ClientConfiguration) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *ClientConfiguration) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *ClientConfiguration) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.


### SetSupportEmailNil

`func (o *ClientConfiguration) SetSupportEmailNil(b bool)`

 SetSupportEmailNil sets the value for SupportEmail to be an explicit nil

### UnsetSupportEmail
`func (o *ClientConfiguration) UnsetSupportEmail()`

UnsetSupportEmail ensures that no value is present for SupportEmail, not even an explicit nil
### GetAisWebFormMode

`func (o *ClientConfiguration) GetAisWebFormMode() WebFormMode`

GetAisWebFormMode returns the AisWebFormMode field if non-nil, zero value otherwise.

### GetAisWebFormModeOk

`func (o *ClientConfiguration) GetAisWebFormModeOk() (*WebFormMode, bool)`

GetAisWebFormModeOk returns a tuple with the AisWebFormMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAisWebFormMode

`func (o *ClientConfiguration) SetAisWebFormMode(v WebFormMode)`

SetAisWebFormMode sets AisWebFormMode field to given value.


### GetPisWebFormMode

`func (o *ClientConfiguration) GetPisWebFormMode() WebFormMode`

GetPisWebFormMode returns the PisWebFormMode field if non-nil, zero value otherwise.

### GetPisWebFormModeOk

`func (o *ClientConfiguration) GetPisWebFormModeOk() (*WebFormMode, bool)`

GetPisWebFormModeOk returns a tuple with the PisWebFormMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPisWebFormMode

`func (o *ClientConfiguration) SetPisWebFormMode(v WebFormMode)`

SetPisWebFormMode sets PisWebFormMode field to given value.


### GetPisStandaloneWebFormMode

`func (o *ClientConfiguration) GetPisStandaloneWebFormMode() WebFormMode`

GetPisStandaloneWebFormMode returns the PisStandaloneWebFormMode field if non-nil, zero value otherwise.

### GetPisStandaloneWebFormModeOk

`func (o *ClientConfiguration) GetPisStandaloneWebFormModeOk() (*WebFormMode, bool)`

GetPisStandaloneWebFormModeOk returns a tuple with the PisStandaloneWebFormMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPisStandaloneWebFormMode

`func (o *ClientConfiguration) SetPisStandaloneWebFormMode(v WebFormMode)`

SetPisStandaloneWebFormMode sets PisStandaloneWebFormMode field to given value.


### GetBetaBanksEnabled

`func (o *ClientConfiguration) GetBetaBanksEnabled() bool`

GetBetaBanksEnabled returns the BetaBanksEnabled field if non-nil, zero value otherwise.

### GetBetaBanksEnabledOk

`func (o *ClientConfiguration) GetBetaBanksEnabledOk() (*bool, bool)`

GetBetaBanksEnabledOk returns a tuple with the BetaBanksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaBanksEnabled

`func (o *ClientConfiguration) SetBetaBanksEnabled(v bool)`

SetBetaBanksEnabled sets BetaBanksEnabled field to given value.


### GetCategoryRestrictions

`func (o *ClientConfiguration) GetCategoryRestrictions() []Category`

GetCategoryRestrictions returns the CategoryRestrictions field if non-nil, zero value otherwise.

### GetCategoryRestrictionsOk

`func (o *ClientConfiguration) GetCategoryRestrictionsOk() (*[]Category, bool)`

GetCategoryRestrictionsOk returns a tuple with the CategoryRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryRestrictions

`func (o *ClientConfiguration) SetCategoryRestrictions(v []Category)`

SetCategoryRestrictions sets CategoryRestrictions field to given value.


### SetCategoryRestrictionsNil

`func (o *ClientConfiguration) SetCategoryRestrictionsNil(b bool)`

 SetCategoryRestrictionsNil sets the value for CategoryRestrictions to be an explicit nil

### UnsetCategoryRestrictions
`func (o *ClientConfiguration) UnsetCategoryRestrictions()`

UnsetCategoryRestrictions ensures that no value is present for CategoryRestrictions, not even an explicit nil
### GetAutoDismountWebForm

`func (o *ClientConfiguration) GetAutoDismountWebForm() bool`

GetAutoDismountWebForm returns the AutoDismountWebForm field if non-nil, zero value otherwise.

### GetAutoDismountWebFormOk

`func (o *ClientConfiguration) GetAutoDismountWebFormOk() (*bool, bool)`

GetAutoDismountWebFormOk returns a tuple with the AutoDismountWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDismountWebForm

`func (o *ClientConfiguration) SetAutoDismountWebForm(v bool)`

SetAutoDismountWebForm sets AutoDismountWebForm field to given value.


### GetCorsAllowedOrigins

`func (o *ClientConfiguration) GetCorsAllowedOrigins() []string`

GetCorsAllowedOrigins returns the CorsAllowedOrigins field if non-nil, zero value otherwise.

### GetCorsAllowedOriginsOk

`func (o *ClientConfiguration) GetCorsAllowedOriginsOk() (*[]string, bool)`

GetCorsAllowedOriginsOk returns a tuple with the CorsAllowedOrigins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsAllowedOrigins

`func (o *ClientConfiguration) SetCorsAllowedOrigins(v []string)`

SetCorsAllowedOrigins sets CorsAllowedOrigins field to given value.


### SetCorsAllowedOriginsNil

`func (o *ClientConfiguration) SetCorsAllowedOriginsNil(b bool)`

 SetCorsAllowedOriginsNil sets the value for CorsAllowedOrigins to be an explicit nil

### UnsetCorsAllowedOrigins
`func (o *ClientConfiguration) UnsetCorsAllowedOrigins()`

UnsetCorsAllowedOrigins ensures that no value is present for CorsAllowedOrigins, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


