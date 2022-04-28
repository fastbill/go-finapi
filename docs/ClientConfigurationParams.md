# ClientConfigurationParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserNotificationCallbackUrl** | Pointer to **string** | Callback URL to which finAPI sends the notification messages that are triggered from the automatic batch update of the users&#39; bank connections. This field is only relevant if the automatic batch update is enabled for your client. For details about what the notification messages look like, please see the documentation in the &#39;Notification Rules&#39; section. finAPI will call this URL with HTTP method POST. Note that the response of the call is not processed by finAPI. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha environment, it MUST be a SSL-secured (https) URL on the finAPI live system.&lt;p&gt;The maximum allowed length of the URL is 512. If you have previously set a callback URL and now want to clear it (thus disabling user-related notifications altogether), you can pass an empty string (\&quot;\&quot;). | [optional] 
**UserSynchronizationCallbackUrl** | Pointer to **string** | Callback URL for user synchronization. This field should be set if you - as a finAPI customer - have multiple clients using finAPI. In such case, all of your clients will share the same user base, making it possible for a user to be created in one client, but then deleted in another. To keep the client-side user data consistent in all clients, you should set a callback URL for each client. finAPI will send a notification to the callback URL of each client whenever a user of your user base gets deleted. Note that finAPI will send a deletion notification to ALL clients, including the one that made the user deletion request to finAPI. So when deleting a user in finAPI, a client should rely on the callback to delete the user on its own side. &lt;p&gt;The notification that finAPI sends to the clients&#39; callback URLs will be a POST request, with this body: &lt;pre&gt;{    \&quot;userId\&quot; : string // contains the identifier of the deleted user    \&quot;event\&quot; : string // this will always be \&quot;DELETED\&quot; }&lt;/pre&gt;&lt;br/&gt;Note that finAPI does not process the response of this call. Also note that while the callback URL may be a non-secured (http) URL on the finAPI sandbox or alpha system, it MUST be a SSL-secured (https) URL on the live system.&lt;/p&gt;As long as you have just one client, you can ignore this field and let it be null. However keep in mind that in this case your client will not receive any callback when a user gets deleted - so the deletion of the user on the client-side must not be forgotten. Of course you may still use the callback URL even for just one client, if you want to implement the deletion of the user on the client-side via the callback from finAPI.&lt;p&gt; The maximum allowed length of the URL is 512. If you have previously set a callback URL and now want to clear it (thus disabling user synchronization related notifications for this client), you can pass an empty string (\&quot;\&quot;). | [optional] 
**RefreshTokensValidityPeriod** | Pointer to **int32** | The validity period that newly requested refresh tokens initially have (in seconds). The value must be greater than or equal to 60, or 0. A value of 0 means that the tokens never expire (Unless explicitly invalidated, e.g. by revocation , or when a user gets locked, or when the password is reset for a user). | [optional] 
**UserAccessTokensValidityPeriod** | Pointer to **int32** | The validity period that newly requested access tokens for users initially have (in seconds). The value must be greater than or equal to 60, or 0. A value of 0 means that the tokens never expire. | [optional] 
**ClientAccessTokensValidityPeriod** | Pointer to **int32** | The validity period that newly requested access tokens for clients initially have (in seconds). The value must be greater than or equal to 60, or 0. A value of 0 means that the tokens never expire. | [optional] 
**IsPinStorageAvailableInWebForm** | Pointer to **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/profiles&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;storeSecrets&#39; field instead.&lt;br/&gt;&lt;br/&gt;Whether finAPI&#39;s Web Form should provide a checkbox for the user allowing him to choose whether to store his banking PIN in finAPI. If this field is set to false, then the user won&#39;t have an option to store his PIN. | [optional] 
**StoreSecretsAvailableInWebForm** | Pointer to **bool** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/profiles&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;storeSecrets&#39; field instead.&lt;br/&gt;&lt;br/&gt;Whether finAPI&#39;s Web Form will allow the user to choose whether to store login secrets (like a PIN) in finAPI. If this field is set to false, the option will not be available and the secrets not stored. | [optional] 
**ApplicationName** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;?product&#x3D;web_form_2.0#post-/api/profiles&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;logo&#39; and &#39;favicon&#39; field instead. The Web Form will now be able to render your logo both in the header and as a favicon.&lt;br/&gt;&lt;br/&gt;When an application name is set (e.g. \&quot;My App\&quot;), then finAPI&#39;s Web Form will display a text to the user \&quot;Weiterleitung auf finAPI von ...\&quot; (e.g. \&quot;Weiterleitung auf finAPI von My App\&quot;). If you have previously set a application name and now want to clear it, you can pass an empty string (\&quot;\&quot;). Maximum length: 64 | [optional] 
**FinTSProductRegistrationNumber** | Pointer to **string** | The FinTS product registration number. Please follow &lt;a href&#x3D;&#39;https://www.hbci-zka.de/register/prod_register.htm&#39; target&#x3D;&#39;_blank&#39;&gt;this link&lt;/a&gt; to apply for a registration number. Only customers who have an AISP or PISP license must define their FinTS product registration number. Customers who are relying on the finAPI Web Form will be assigned to finAPI&#39;s FinTS product registration number automatically and do not have to register themselves. During a batch update, finAPI is using the FinTS product registration number of the client, that was used to create the user. If you have previously set a FinTS product registration number and now want to clear it, you can pass an empty string (\&quot;\&quot;). Only hexa decimal characters in capital case with a maximum length of 25 characters are allowed. E.g. &#39;ABCDEF1234567890ABCDEF123&#39; | [optional] 
**SupportSubjectDefault** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/For-best-results!.2477654019.html&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;errorRedirectUrl&#39; and &#39;customerSupportUrl&#39; query parameters instead.&lt;br/&gt;&lt;br/&gt;Default value for the subject element of support emails. Maximum length is 100. Pass an empty string (&#39;&#39;) if you want to clear the current subject default value. | [optional] 
**SupportEmail** | Pointer to **string** | THIS FIELD IS DEPRECATED AND WILL BE REMOVED.&lt;br/&gt;Please refer &lt;a href&#x3D;&#39;https://documentation.finapi.io/webform/For-best-results!.2477654019.html&#39; target&#x3D;&#39;_blank&#39;&gt;here&lt;/a&gt; to the &#39;errorRedirectUrl&#39; and &#39;customerSupportUrl&#39; query parameters instead.&lt;br/&gt;&lt;br/&gt;Email address to sent support requests to from the Web Form. Maximum length is 320. Pass an empty string (&#39;&#39;) if you want to clear the current email address. | [optional] 
**BetaBanksEnabled** | Pointer to **bool** | Whether the set of banks that are available to your client should include “Beta banks”. Beta banks provide pre-release interfaces that are still in a beta phase. Communication to the bank via such interfaces might be unstable, and the correctness and/or quality of data delivery or payment execution cannot be guaranteed.&lt;br/&gt;As the word “BETA” already indicates, Beta banks are subject to changes. Their properties, as well as their behaviour can change based on continuous tests and customer feedback. Also, to keep our bank list clean, we might remove Beta banks at any point in time, including all related user data (bank connections, accounts, transactions etc). We still recommend you to enable beta banks in your application, because it enables us to release a stable interface faster. However, you should point it out to your users when using a beta bank (also see field Bank.isBeta).&lt;br/&gt;&lt;br/&gt;If this field is true, then the GET /banks services will include beta banks in their results, and you can use beta banks in any service where you can pass a bank identifier. If the field is false, then beta banks will not exist for your client. | [optional] 

## Methods

### NewClientConfigurationParams

`func NewClientConfigurationParams() *ClientConfigurationParams`

NewClientConfigurationParams instantiates a new ClientConfigurationParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConfigurationParamsWithDefaults

`func NewClientConfigurationParamsWithDefaults() *ClientConfigurationParams`

NewClientConfigurationParamsWithDefaults instantiates a new ClientConfigurationParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserNotificationCallbackUrl

`func (o *ClientConfigurationParams) GetUserNotificationCallbackUrl() string`

GetUserNotificationCallbackUrl returns the UserNotificationCallbackUrl field if non-nil, zero value otherwise.

### GetUserNotificationCallbackUrlOk

`func (o *ClientConfigurationParams) GetUserNotificationCallbackUrlOk() (*string, bool)`

GetUserNotificationCallbackUrlOk returns a tuple with the UserNotificationCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNotificationCallbackUrl

`func (o *ClientConfigurationParams) SetUserNotificationCallbackUrl(v string)`

SetUserNotificationCallbackUrl sets UserNotificationCallbackUrl field to given value.

### HasUserNotificationCallbackUrl

`func (o *ClientConfigurationParams) HasUserNotificationCallbackUrl() bool`

HasUserNotificationCallbackUrl returns a boolean if a field has been set.

### GetUserSynchronizationCallbackUrl

`func (o *ClientConfigurationParams) GetUserSynchronizationCallbackUrl() string`

GetUserSynchronizationCallbackUrl returns the UserSynchronizationCallbackUrl field if non-nil, zero value otherwise.

### GetUserSynchronizationCallbackUrlOk

`func (o *ClientConfigurationParams) GetUserSynchronizationCallbackUrlOk() (*string, bool)`

GetUserSynchronizationCallbackUrlOk returns a tuple with the UserSynchronizationCallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSynchronizationCallbackUrl

`func (o *ClientConfigurationParams) SetUserSynchronizationCallbackUrl(v string)`

SetUserSynchronizationCallbackUrl sets UserSynchronizationCallbackUrl field to given value.

### HasUserSynchronizationCallbackUrl

`func (o *ClientConfigurationParams) HasUserSynchronizationCallbackUrl() bool`

HasUserSynchronizationCallbackUrl returns a boolean if a field has been set.

### GetRefreshTokensValidityPeriod

`func (o *ClientConfigurationParams) GetRefreshTokensValidityPeriod() int32`

GetRefreshTokensValidityPeriod returns the RefreshTokensValidityPeriod field if non-nil, zero value otherwise.

### GetRefreshTokensValidityPeriodOk

`func (o *ClientConfigurationParams) GetRefreshTokensValidityPeriodOk() (*int32, bool)`

GetRefreshTokensValidityPeriodOk returns a tuple with the RefreshTokensValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokensValidityPeriod

`func (o *ClientConfigurationParams) SetRefreshTokensValidityPeriod(v int32)`

SetRefreshTokensValidityPeriod sets RefreshTokensValidityPeriod field to given value.

### HasRefreshTokensValidityPeriod

`func (o *ClientConfigurationParams) HasRefreshTokensValidityPeriod() bool`

HasRefreshTokensValidityPeriod returns a boolean if a field has been set.

### GetUserAccessTokensValidityPeriod

`func (o *ClientConfigurationParams) GetUserAccessTokensValidityPeriod() int32`

GetUserAccessTokensValidityPeriod returns the UserAccessTokensValidityPeriod field if non-nil, zero value otherwise.

### GetUserAccessTokensValidityPeriodOk

`func (o *ClientConfigurationParams) GetUserAccessTokensValidityPeriodOk() (*int32, bool)`

GetUserAccessTokensValidityPeriodOk returns a tuple with the UserAccessTokensValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAccessTokensValidityPeriod

`func (o *ClientConfigurationParams) SetUserAccessTokensValidityPeriod(v int32)`

SetUserAccessTokensValidityPeriod sets UserAccessTokensValidityPeriod field to given value.

### HasUserAccessTokensValidityPeriod

`func (o *ClientConfigurationParams) HasUserAccessTokensValidityPeriod() bool`

HasUserAccessTokensValidityPeriod returns a boolean if a field has been set.

### GetClientAccessTokensValidityPeriod

`func (o *ClientConfigurationParams) GetClientAccessTokensValidityPeriod() int32`

GetClientAccessTokensValidityPeriod returns the ClientAccessTokensValidityPeriod field if non-nil, zero value otherwise.

### GetClientAccessTokensValidityPeriodOk

`func (o *ClientConfigurationParams) GetClientAccessTokensValidityPeriodOk() (*int32, bool)`

GetClientAccessTokensValidityPeriodOk returns a tuple with the ClientAccessTokensValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAccessTokensValidityPeriod

`func (o *ClientConfigurationParams) SetClientAccessTokensValidityPeriod(v int32)`

SetClientAccessTokensValidityPeriod sets ClientAccessTokensValidityPeriod field to given value.

### HasClientAccessTokensValidityPeriod

`func (o *ClientConfigurationParams) HasClientAccessTokensValidityPeriod() bool`

HasClientAccessTokensValidityPeriod returns a boolean if a field has been set.

### GetIsPinStorageAvailableInWebForm

`func (o *ClientConfigurationParams) GetIsPinStorageAvailableInWebForm() bool`

GetIsPinStorageAvailableInWebForm returns the IsPinStorageAvailableInWebForm field if non-nil, zero value otherwise.

### GetIsPinStorageAvailableInWebFormOk

`func (o *ClientConfigurationParams) GetIsPinStorageAvailableInWebFormOk() (*bool, bool)`

GetIsPinStorageAvailableInWebFormOk returns a tuple with the IsPinStorageAvailableInWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPinStorageAvailableInWebForm

`func (o *ClientConfigurationParams) SetIsPinStorageAvailableInWebForm(v bool)`

SetIsPinStorageAvailableInWebForm sets IsPinStorageAvailableInWebForm field to given value.

### HasIsPinStorageAvailableInWebForm

`func (o *ClientConfigurationParams) HasIsPinStorageAvailableInWebForm() bool`

HasIsPinStorageAvailableInWebForm returns a boolean if a field has been set.

### GetStoreSecretsAvailableInWebForm

`func (o *ClientConfigurationParams) GetStoreSecretsAvailableInWebForm() bool`

GetStoreSecretsAvailableInWebForm returns the StoreSecretsAvailableInWebForm field if non-nil, zero value otherwise.

### GetStoreSecretsAvailableInWebFormOk

`func (o *ClientConfigurationParams) GetStoreSecretsAvailableInWebFormOk() (*bool, bool)`

GetStoreSecretsAvailableInWebFormOk returns a tuple with the StoreSecretsAvailableInWebForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreSecretsAvailableInWebForm

`func (o *ClientConfigurationParams) SetStoreSecretsAvailableInWebForm(v bool)`

SetStoreSecretsAvailableInWebForm sets StoreSecretsAvailableInWebForm field to given value.

### HasStoreSecretsAvailableInWebForm

`func (o *ClientConfigurationParams) HasStoreSecretsAvailableInWebForm() bool`

HasStoreSecretsAvailableInWebForm returns a boolean if a field has been set.

### GetApplicationName

`func (o *ClientConfigurationParams) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ClientConfigurationParams) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ClientConfigurationParams) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *ClientConfigurationParams) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetFinTSProductRegistrationNumber

`func (o *ClientConfigurationParams) GetFinTSProductRegistrationNumber() string`

GetFinTSProductRegistrationNumber returns the FinTSProductRegistrationNumber field if non-nil, zero value otherwise.

### GetFinTSProductRegistrationNumberOk

`func (o *ClientConfigurationParams) GetFinTSProductRegistrationNumberOk() (*string, bool)`

GetFinTSProductRegistrationNumberOk returns a tuple with the FinTSProductRegistrationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinTSProductRegistrationNumber

`func (o *ClientConfigurationParams) SetFinTSProductRegistrationNumber(v string)`

SetFinTSProductRegistrationNumber sets FinTSProductRegistrationNumber field to given value.

### HasFinTSProductRegistrationNumber

`func (o *ClientConfigurationParams) HasFinTSProductRegistrationNumber() bool`

HasFinTSProductRegistrationNumber returns a boolean if a field has been set.

### GetSupportSubjectDefault

`func (o *ClientConfigurationParams) GetSupportSubjectDefault() string`

GetSupportSubjectDefault returns the SupportSubjectDefault field if non-nil, zero value otherwise.

### GetSupportSubjectDefaultOk

`func (o *ClientConfigurationParams) GetSupportSubjectDefaultOk() (*string, bool)`

GetSupportSubjectDefaultOk returns a tuple with the SupportSubjectDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSubjectDefault

`func (o *ClientConfigurationParams) SetSupportSubjectDefault(v string)`

SetSupportSubjectDefault sets SupportSubjectDefault field to given value.

### HasSupportSubjectDefault

`func (o *ClientConfigurationParams) HasSupportSubjectDefault() bool`

HasSupportSubjectDefault returns a boolean if a field has been set.

### GetSupportEmail

`func (o *ClientConfigurationParams) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *ClientConfigurationParams) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *ClientConfigurationParams) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.

### HasSupportEmail

`func (o *ClientConfigurationParams) HasSupportEmail() bool`

HasSupportEmail returns a boolean if a field has been set.

### GetBetaBanksEnabled

`func (o *ClientConfigurationParams) GetBetaBanksEnabled() bool`

GetBetaBanksEnabled returns the BetaBanksEnabled field if non-nil, zero value otherwise.

### GetBetaBanksEnabledOk

`func (o *ClientConfigurationParams) GetBetaBanksEnabledOk() (*bool, bool)`

GetBetaBanksEnabledOk returns a tuple with the BetaBanksEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBetaBanksEnabled

`func (o *ClientConfigurationParams) SetBetaBanksEnabled(v bool)`

SetBetaBanksEnabled sets BetaBanksEnabled field to given value.

### HasBetaBanksEnabled

`func (o *ClientConfigurationParams) HasBetaBanksEnabled() bool`

HasBetaBanksEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


