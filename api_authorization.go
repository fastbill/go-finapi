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
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// AuthorizationApiService AuthorizationApi service
type AuthorizationApiService service

type ApiGetTokenRequest struct {
	ctx _context.Context
	ApiService *AuthorizationApiService
	grantType *string
	clientId *string
	clientSecret *string
	xRequestId *string
	refreshToken *string
	username *string
	password *string
}

func (r ApiGetTokenRequest) GrantType(grantType string) ApiGetTokenRequest {
	r.grantType = &grantType
	return r
}
func (r ApiGetTokenRequest) ClientId(clientId string) ApiGetTokenRequest {
	r.clientId = &clientId
	return r
}
func (r ApiGetTokenRequest) ClientSecret(clientSecret string) ApiGetTokenRequest {
	r.clientSecret = &clientSecret
	return r
}
func (r ApiGetTokenRequest) XRequestId(xRequestId string) ApiGetTokenRequest {
	r.xRequestId = &xRequestId
	return r
}
func (r ApiGetTokenRequest) RefreshToken(refreshToken string) ApiGetTokenRequest {
	r.refreshToken = &refreshToken
	return r
}
func (r ApiGetTokenRequest) Username(username string) ApiGetTokenRequest {
	r.username = &username
	return r
}
func (r ApiGetTokenRequest) Password(password string) ApiGetTokenRequest {
	r.password = &password
	return r
}

func (r ApiGetTokenRequest) Execute() (AccessToken, *_nethttp.Response, error) {
	return r.ApiService.GetTokenExecute(r)
}

/*
 * GetToken Get tokens
 * finAPI implements the OAuth 2.0 Standard for authorizing applications and users within applications. OAuth uses the terminology of clients and users. A client represents an application that calls finAPI services. A service call might be in the context of a user of the client (e.g.: getting a user's bank connections), or outside any user context (e.g.: editing your client's configuration, or creating a new user for your client). In any case, every service call must be authorized by an access_token. This service can be used to get such an access_token, for either one of the client's users, or for the client itself. Also, this service can be used to refresh the access_token of a user that has previously requested an access_token.<br/><br/>To get a token, you must always pass a valid client identifier and client secret (=client credentials). You can get free client credentials for the sandbox <a href='http://www.finapi.io/jetzt-testen/' target='_blank'>here</a>. Alternatively, you can also contact us at <a href='mailto:support@finapi.io'>support@finapi.io</a>.<br/><br/>The authorization process is similar for both a user within a client, and for the client itself: <br/>&bull; To authorize a client (i.e. application), use <code>grant_type=client_credentials</code><br/>&bull; To authorize a user, use <code>grant_type=password</code><br/><br/>If the given parameters are valid, the service will respond with the authorization data. <br/>Here is an example of a response when authorizing a user: <br/><pre>{
   "access_token": "yvMbx_TgwdYE0hgOVb8N4ZOvxOukqfjzYOGRZcJiCjQuRGkVIBfjjV3YG4zKTGiY2aPn2cQTGaQOT8uo5uo7_QOXts1s5UBSVuRHc6a8X30RrGBTyqV9h26SUHcZPNbZ",
   "token_type": "bearer",
   "refresh_token": "0b9KjiBVlZLz7a4HshSAIcFuscStiXT1VzT5mgNYwCQ_dWctTDsaIjedAhD1LpsOFJ7x6K8Emf8M3VOQkwNFR9FHijALYSQw2UeRwAC2MvrOKwfF1dHmOq5VEVYEaGf6",
   "expires_in": 3600,
   "scope": "all"
}</pre><br/><p>Use the returned access_token for other service calls by sending it in a 'Authorization' header, with the word 'Bearer' in front of the token. Like this:</p><pre>Authorization: Bearer yvMbx_TgwdYE0hgOVb8N4ZOvxOukqfjzYOGRZcJiCjQuRGkVIBfjjV3YG4zKTGiY2aPn2cQTGaQOT8uo5uo7_QOXts1s5UBSVuRHc6a8X30RrGBTyqV9h26SUHcZPNbZ</pre><p><b>WARNING</b>: Sending the access_token as a request parameter is deprecated and will probably be no longer supported in the next release of finAPI. Please always send the access_token in the request header, as shown above.</p><p>By default, the access tokens have an expiration time of one hour (however, you can change this via the service PATCH /clientConfiguration). If a token has expired, then using the token for a service call will result in a HTTP code 401. To restore access you can simply get a new token (as it is described above) or use <code>grant_type=refresh_token</code> (which works for user-related tokens only). In the latter case you just have to pass the previously received <code>refresh_token</code> for the user.</p><p>If explicit user verification is required (the 'isUserAutoVerificationEnabled' flag in the client configuration is set to false, see Client Configuration) and the user that you want to authorize is not yet verified by the client (see Verify a User), then the service will respond with HTTP code 403. If the user is locked (see 'maxUserLoginAttempts' in the Client Configuration), the service will respond with HTTP code 423.</p><p>If the current role has no privileges to call a certain service (e.g. if a user tries to create a new user, or if a client tries to access user data outside of any user context), then the request will fail with the HTTP code 403.</p><p><b>IMPORTANT NOTES:</b><br/>&bull; Even though finAPI is not logging query parameters, it is still recommended to pass the parameters in the POST body instead of in the URL. Also, please set the Content-Type of your request to 'application/x-www-form-urlencoded' when calling this service.<br/>&bull; You should use this service only when you actually need a new token. As long as a token exists and has not expired, the service will always return the same token for the same credentials. Calling this service repeatedly with the same credentials contradicts the idea behind the tokens in OAuth, and will have a negative impact on the performance of your application. So instead of retrieving the same tokens over and over with this service, you should cache the tokens and re-use them as long as they have not expired - or at least as long as you're using the same tokens repeatedly, e.g. for the time of an active user session in your application.</p>
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetTokenRequest
 */
func (a *AuthorizationApiService) GetToken(ctx _context.Context) ApiGetTokenRequest {
	return ApiGetTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return AccessToken
 */
func (a *AuthorizationApiService) GetTokenExecute(r ApiGetTokenRequest) (AccessToken, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  AccessToken
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationApiService.GetToken")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.grantType == nil {
		return localVarReturnValue, nil, reportError("grantType is required and must be specified")
	}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.clientSecret == nil {
		return localVarReturnValue, nil, reportError("clientSecret is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xRequestId != nil {
		localVarHeaderParams["X-Request-Id"] = parameterToString(*r.xRequestId, "")
	}
	localVarFormParams.Add("grant_type", parameterToString(*r.grantType, ""))
	localVarFormParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarFormParams.Add("client_secret", parameterToString(*r.clientSecret, ""))
	if r.refreshToken != nil {
		localVarFormParams.Add("refresh_token", parameterToString(*r.refreshToken, ""))
	}
	if r.username != nil {
		localVarFormParams.Add("username", parameterToString(*r.username, ""))
	}
	if r.password != nil {
		localVarFormParams.Add("password", parameterToString(*r.password, ""))
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v BadCredentialsError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 423 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRevokeTokenRequest struct {
	ctx _context.Context
	ApiService *AuthorizationApiService
	token *string
	tokenTypeHint *string
	xRequestId *string
}

func (r ApiRevokeTokenRequest) Token(token string) ApiRevokeTokenRequest {
	r.token = &token
	return r
}
func (r ApiRevokeTokenRequest) TokenTypeHint(tokenTypeHint string) ApiRevokeTokenRequest {
	r.tokenTypeHint = &tokenTypeHint
	return r
}
func (r ApiRevokeTokenRequest) XRequestId(xRequestId string) ApiRevokeTokenRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiRevokeTokenRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.RevokeTokenExecute(r)
}

/*
 * RevokeToken Revoke a token
 * An additional endpoint for the OAuth 2.0 Standard, which allows clients to notify finAPI that a previously obtained refresh_token or access_token is no longer required. A successful request will invalidate the given token. The revocation of a particular token may also cause the revocation of related tokens and the underlying authorization grant. For token_type_hint=access_token finAPI will invalidate only the given access_token. For token_type_hint=refresh_token, finAPI will invalidate the refresh token and all access tokens based on the same authorization grant. If the token_type_hint is not defined, finAPI will revoke all access and refresh tokens (if applicable) that are based on the same authorization grant.<br/><br/>Note that the service responds with HTTP status code 200 both if the token has been revoked successfully, and if the client submitted an invalid token.<br/><br/>Note also that the client's access_token is required to authenticate the revocation.<br/><br/>Here is an example of how to revoke a user's refresh_token (and therefore also his access tokens):<pre>Authorization: Bearer {client_access_token}
POST /oauth/revoke?token={refresh_token}&token_type_hint=refresh_token</pre>
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiRevokeTokenRequest
 */
func (a *AuthorizationApiService) RevokeToken(ctx _context.Context) ApiRevokeTokenRequest {
	return ApiRevokeTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *AuthorizationApiService) RevokeTokenExecute(r ApiRevokeTokenRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthorizationApiService.RevokeToken")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth/revoke"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.token == nil {
		return nil, reportError("token is required and must be specified")
	}

	localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	if r.tokenTypeHint != nil {
		localVarQueryParams.Add("token_type_hint", parameterToString(*r.tokenTypeHint, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xRequestId != nil {
		localVarHeaderParams["X-Request-Id"] = parameterToString(*r.xRequestId, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
