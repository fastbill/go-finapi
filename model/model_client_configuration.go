/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.79.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Client configuration parameters
type ClientConfiguration struct {
	// Whether finAPI performs a regular automatic update of your users' bank connections. To find out how the automatic batch update is configured for your client, i.e. which bank connections get updated, and at which time and interval, please contact your Sys-Admin. Note that even if the automatic batch update is enabled for your client, individual users can still disable the feature for their own bank connections.
	IsAutomaticBatchUpdateEnabled bool `json:"isAutomaticBatchUpdateEnabled"`
	// Callback URL to which finAPI sends the notification messages that are triggered from the automatic batch update of the users' bank connections. This field is only relevant if the automatic batch update is enabled for your client. For details about what the notification messages look like, please see the documentation in the 'Notification Rules' section. finAPI will call this URL with HTTP method POST. Note that the response of the call is not processed by finAPI. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha environment, it MUST be a SSL-secured (https) URL on the finAPI live system.
	UserNotificationCallbackUrl string `json:"userNotificationCallbackUrl,omitempty"`
	// Callback URL for user synchronization. This field should be set if you - as a finAPI customer - have multiple clients using finAPI. In such case, all of your clients will share the same user base, making it possible for a user to be created in one client, but then deleted in another. To keep the client-side user data consistent in all clients, you should set a callback URL for each client. finAPI will send a notification to the callback URL of each client whenever a user of your user base gets deleted. Note that finAPI will send a deletion notification to ALL clients, including the one that made the user deletion request to finAPI. So when deleting a user in finAPI, a client should rely on the callback to delete the user on its own side. <p>The notification that finAPI sends to the clients' callback URLs will be a POST request, with this body: <pre>{    \"userId\" : string // contains the identifier of the deleted user    \"event\" : string // this will always be \"DELETED\" }</pre><br/>Note that finAPI does not process the response of this call. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha environment, it MUST be a SSL-secured (https) URL on the finAPI live system.</p>As long as you have just one client, you can ignore this field and let it be null. However keep in mind that in this case your client will not receive any callback when a user gets deleted - so the deletion of the user on the client-side must not be forgotten. Of course you may still use the callback URL even for just one client, if you want to implement the deletion of the user on the client-side via the callback from finAPI.
	UserSynchronizationCallbackUrl string `json:"userSynchronizationCallbackUrl,omitempty"`
	// The validity period that newly requested refresh tokens initially have (in seconds). A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation, or when a user gets locked, or when the password is reset for a user).
	RefreshTokensValidityPeriod int32 `json:"refreshTokensValidityPeriod,omitempty"`
	// The validity period that newly requested access tokens for users initially have (in seconds). A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation, or when a user gets locked, or when the password is reset for a user).
	UserAccessTokensValidityPeriod int32 `json:"userAccessTokensValidityPeriod,omitempty"`
	// The validity period that newly requested access tokens for clients initially have (in seconds). A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation).
	ClientAccessTokensValidityPeriod int32 `json:"clientAccessTokensValidityPeriod,omitempty"`
	// Number of consecutive failed login attempts of a user into his finAPI account that is allowed before finAPI locks the user's account. When a user's account is locked, finAPI will invalidate all user's tokens and it will deny any service call in the context of this user (i.e. any call to a service using one of the user's authorization tokens, as well as the service for requesting a new token for this user). To unlock a user's account, a new password must be set for the account by the client (see the services /users/requestPasswordChange and /users/executePasswordChange). Once a new password has been set, all services will be available again for this user and the user's failed login attempts counter is reset to 0. The user's failed login attempts counter is also reset whenever a new authorization token has been successfully retrieved, or whenever the user himself changes his password.<br/><br/>Note that when this field has a value of 0, it means that there is no limit for user login attempts, i.e. finAPI will never lock user accounts.
	MaxUserLoginAttempts int32 `json:"maxUserLoginAttempts"`
	// Whether users that are created with this client are automatically verified on creation. If this field is set to 'false', then any user that is created with this client must first be verified with the \"Verify a user\" service before he can be authorized. If the field is 'true', then no verification is required by the client and the user can be authorized immediately after creation.
	IsUserAutoVerificationEnabled bool `json:"isUserAutoVerificationEnabled"`
	// Whether this client is a 'Mandator Admin'. Mandator Admins are special clients that can access the 'Mandator Administration' section of finAPI. If you do not yet have credentials for a Mandator Admin, please contact us at support@finapi.io. For further information, please refer to <a href='https://finapi.zendesk.com/hc/en-us/articles/115003661827-Difference-between-app-clients-and-mandator-admin-client' target='_blank'>this article</a> on our Dev Portal.
	IsMandatorAdmin bool `json:"isMandatorAdmin"`
	// Whether finAPI is allowed to use the WEB_SCRAPER interface for data download. If this field is set to 'true', then finAPI might download data from the online banking websites of banks (either in addition to other interfaces, or as the sole data source for the download). If this field is set to 'false', then finAPI will not use any web scrapers. For banks where no other interface except WEB_SCRAPER is available, finAPI will not allow any data download at all if web scraping is disabled for your client. Please contact your Sys-Admin if you want to change this setting.
	IsWebScrapingEnabled bool `json:"isWebScrapingEnabled"`
	// Whether this client is allowed to access XS2A services.
	IsXs2aEnabled bool `json:"isXs2aEnabled"`
	// List of bank groups that are available to this client. A bank group is a collection of all banks that are located in a certain country, and is defined by the country's ISO 3166 ALPHA-2 code (see also field 'location' of Bank resource). If you want to extend or limit the available bank groups for your client, please contact your Sys-Admin.<br/><br/>Note: There is no bank group for international institutes (i.e. institutes that are not bound to any specific country). Instead, those institutes are always available. If this list is empty, it means that ONLY international institutes are available.
	AvailableBankGroups []string `json:"availableBankGroups"`
	// Application name. When an application name is set (e.g. \"My App\"), then <a href='https://finapi.zendesk.com/hc/en-us/articles/360002596391' target='_blank'>finAPI's web form</a> will display a text to the user \"Weiterleitung auf finAPI von ...\" (e.g. \"Weiterleitung auf finAPI von MyApp\").
	ApplicationName string `json:"applicationName,omitempty"`
	// The FinTS product registration number. If a value is stored, this will always be 'XXXXX'.
	FinTSProductRegistrationNumber string `json:"finTSProductRegistrationNumber,omitempty"`
	// Whether <a href='https://finapi.zendesk.com/hc/en-us/articles/360002596391' target='_blank'>finAPI's web form</a> will provide a checkbox for the user allowing him to choose whether to store login secrets (like a PIN) in finAPI. If this field is set to false, then the user won't have an option to store this data.
	StoreSecretsAvailableInWebForm bool `json:"storeSecretsAvailableInWebForm"`
	// Whether <a href='https://finapi.zendesk.com/hc/en-us/articles/360002596391' target='_blank'>finAPI's web form</a> will provide a checkbox for the user allowing him to choose whether to store his banking PIN in finAPI. If this field is set to false, then the user won't have an option to store his PIN.<br><br>NOTE: This field is deprecated and will be removed at some point. Refer to field 'storeSecretsAvailableInWebForm' instead.
	PinStorageAvailableInWebForm bool `json:"pinStorageAvailableInWebForm"`
	// Whether this client is allowed to do payments
	PaymentsEnabled bool `json:"paymentsEnabled"`
}
