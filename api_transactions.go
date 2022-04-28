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
	"strings"
	"reflect"
)

// Linger please
var (
	_ _context.Context
)

// TransactionsApiService TransactionsApi service
type TransactionsApiService service

type ApiDeleteAllTransactionsRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	maxDeletionDate *string
	minImportDate *string
	accountIds *[]int64
	safeMode *bool
	rememberDeletion *bool
	xHTTPMethodOverride *string
	xRequestId *string
}

func (r ApiDeleteAllTransactionsRequest) MaxDeletionDate(maxDeletionDate string) ApiDeleteAllTransactionsRequest {
	r.maxDeletionDate = &maxDeletionDate
	return r
}
func (r ApiDeleteAllTransactionsRequest) MinImportDate(minImportDate string) ApiDeleteAllTransactionsRequest {
	r.minImportDate = &minImportDate
	return r
}
func (r ApiDeleteAllTransactionsRequest) AccountIds(accountIds []int64) ApiDeleteAllTransactionsRequest {
	r.accountIds = &accountIds
	return r
}
func (r ApiDeleteAllTransactionsRequest) SafeMode(safeMode bool) ApiDeleteAllTransactionsRequest {
	r.safeMode = &safeMode
	return r
}
func (r ApiDeleteAllTransactionsRequest) RememberDeletion(rememberDeletion bool) ApiDeleteAllTransactionsRequest {
	r.rememberDeletion = &rememberDeletion
	return r
}
func (r ApiDeleteAllTransactionsRequest) XHTTPMethodOverride(xHTTPMethodOverride string) ApiDeleteAllTransactionsRequest {
	r.xHTTPMethodOverride = &xHTTPMethodOverride
	return r
}
func (r ApiDeleteAllTransactionsRequest) XRequestId(xRequestId string) ApiDeleteAllTransactionsRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiDeleteAllTransactionsRequest) Execute() (IdentifierList, *_nethttp.Response, error) {
	return r.ApiService.DeleteAllTransactionsExecute(r)
}

/*
 * DeleteAllTransactions Delete all transactions
 * Delete a set, or the entirety, of transactions of the currently authorized user.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiDeleteAllTransactionsRequest
 */
func (a *TransactionsApiService) DeleteAllTransactions(ctx _context.Context) ApiDeleteAllTransactionsRequest {
	return ApiDeleteAllTransactionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return IdentifierList
 */
func (a *TransactionsApiService) DeleteAllTransactionsExecute(r ApiDeleteAllTransactionsRequest) (IdentifierList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IdentifierList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.DeleteAllTransactions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.maxDeletionDate != nil {
		localVarQueryParams.Add("maxDeletionDate", parameterToString(*r.maxDeletionDate, ""))
	}
	if r.minImportDate != nil {
		localVarQueryParams.Add("minImportDate", parameterToString(*r.minImportDate, ""))
	}
	if r.accountIds != nil {
		t := *r.accountIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("accountIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("accountIds", parameterToString(t, "multi"))
		}
	}
	if r.safeMode != nil {
		localVarQueryParams.Add("safeMode", parameterToString(*r.safeMode, ""))
	}
	if r.rememberDeletion != nil {
		localVarQueryParams.Add("rememberDeletion", parameterToString(*r.rememberDeletion, ""))
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
	if r.xHTTPMethodOverride != nil {
		localVarHeaderParams["X-HTTP-Method-Override"] = parameterToString(*r.xHTTPMethodOverride, "")
	}
	if r.xRequestId != nil {
		localVarHeaderParams["X-Request-Id"] = parameterToString(*r.xRequestId, "")
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

type ApiDeleteTransactionRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	id int64
	xHTTPMethodOverride *string
	xRequestId *string
}

func (r ApiDeleteTransactionRequest) XHTTPMethodOverride(xHTTPMethodOverride string) ApiDeleteTransactionRequest {
	r.xHTTPMethodOverride = &xHTTPMethodOverride
	return r
}
func (r ApiDeleteTransactionRequest) XRequestId(xRequestId string) ApiDeleteTransactionRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiDeleteTransactionRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteTransactionExecute(r)
}

/*
 * DeleteTransaction Delete a transaction
 * Delete a single transaction of the user that is authorized by the access_token.<br/><br/> A transaction can only get deleted if at least one of the following holds true:<br/>1. The transaction belongs to an account of a test bank <br/>2. The transaction's 'potentialDuplicate' flag is set to TRUE<br/>3. The transaction is an adjusting entry ('Zwischensaldo' transaction) that was added by finAPI<br/><br/>Note that the 'Delete all transactions' service has additional functionality and allows you to delete transactions that you cannot delete with this service.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Identifier of transaction
 * @return ApiDeleteTransactionRequest
 */
func (a *TransactionsApiService) DeleteTransaction(ctx _context.Context, id int64) ApiDeleteTransactionRequest {
	return ApiDeleteTransactionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 */
func (a *TransactionsApiService) DeleteTransactionExecute(r ApiDeleteTransactionRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.DeleteTransaction")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	if r.xHTTPMethodOverride != nil {
		localVarHeaderParams["X-HTTP-Method-Override"] = parameterToString(*r.xHTTPMethodOverride, "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v BadCredentialsError
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiEditMultipleTransactionsRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	updateMultipleTransactionsParams *UpdateMultipleTransactionsParams
	xHTTPMethodOverride *string
	xRequestId *string
}

func (r ApiEditMultipleTransactionsRequest) UpdateMultipleTransactionsParams(updateMultipleTransactionsParams UpdateMultipleTransactionsParams) ApiEditMultipleTransactionsRequest {
	r.updateMultipleTransactionsParams = &updateMultipleTransactionsParams
	return r
}
func (r ApiEditMultipleTransactionsRequest) XHTTPMethodOverride(xHTTPMethodOverride string) ApiEditMultipleTransactionsRequest {
	r.xHTTPMethodOverride = &xHTTPMethodOverride
	return r
}
func (r ApiEditMultipleTransactionsRequest) XRequestId(xRequestId string) ApiEditMultipleTransactionsRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiEditMultipleTransactionsRequest) Execute() (IdentifierList, *_nethttp.Response, error) {
	return r.ApiService.EditMultipleTransactionsExecute(r)
}

/*
 * EditMultipleTransactions Edit multiple transactions
 * Edit one or multiple transactions. You can edit the following fields: 'isNew=true|false' and/or 'isPotentialDuplicate=false' and/or 'categoryId=&lt;id&gt;' and/or 'labelIds=[&lt;ids&gt;]'. To clear the category of the given transactions (so that they are no longer categorized), pass the value '0' as the categoryId. To clear the labels of the given transactions, pass an empty array of label identifiers: '[]'. The parameters 'categoryId' and 'labelIds' are forbidden if 'ids' is NOT set (i.e. you cannot update the category or labels for ALL transactions). The result is a list of identifiers of only those transactions that have changed as a result of this service call.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiEditMultipleTransactionsRequest
 */
func (a *TransactionsApiService) EditMultipleTransactions(ctx _context.Context) ApiEditMultipleTransactionsRequest {
	return ApiEditMultipleTransactionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return IdentifierList
 */
func (a *TransactionsApiService) EditMultipleTransactionsExecute(r ApiEditMultipleTransactionsRequest) (IdentifierList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IdentifierList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.EditMultipleTransactions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.updateMultipleTransactionsParams == nil {
		return localVarReturnValue, nil, reportError("updateMultipleTransactionsParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if r.xHTTPMethodOverride != nil {
		localVarHeaderParams["X-HTTP-Method-Override"] = parameterToString(*r.xHTTPMethodOverride, "")
	}
	if r.xRequestId != nil {
		localVarHeaderParams["X-Request-Id"] = parameterToString(*r.xRequestId, "")
	}
	// body params
	localVarPostBody = r.updateMultipleTransactionsParams
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiEditMultipleTransactionsDeprecatedRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	ids []int64
	updateTransactionsParams *UpdateTransactionsParams
	xHTTPMethodOverride *string
	xRequestId *string
}

func (r ApiEditMultipleTransactionsDeprecatedRequest) UpdateTransactionsParams(updateTransactionsParams UpdateTransactionsParams) ApiEditMultipleTransactionsDeprecatedRequest {
	r.updateTransactionsParams = &updateTransactionsParams
	return r
}
func (r ApiEditMultipleTransactionsDeprecatedRequest) XHTTPMethodOverride(xHTTPMethodOverride string) ApiEditMultipleTransactionsDeprecatedRequest {
	r.xHTTPMethodOverride = &xHTTPMethodOverride
	return r
}
func (r ApiEditMultipleTransactionsDeprecatedRequest) XRequestId(xRequestId string) ApiEditMultipleTransactionsDeprecatedRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiEditMultipleTransactionsDeprecatedRequest) Execute() (IdentifierList, *_nethttp.Response, error) {
	return r.ApiService.EditMultipleTransactionsDeprecatedExecute(r)
}

/*
 * EditMultipleTransactionsDeprecated Edit multiple transactions (DEPRECATED)
 * THIS SERVICE IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'Edit multiple transactions' service instead if you want to edit multiple transactions, and pass a comma-separated list of identifiers with the 'ids' parameter.<br/><br/>Change the fields of multiple transactions. You can change the following fields: 'isNew=true|false' and/or 'isPotentialDuplicate=false' and/or 'categoryId=&lt;id&gt;' and/or 'labelIds=[<ids>]'. The max number of input identifiers is limited to 100. To clear the category of the given transactions (so that they are no longer categorized), pass the value '0' as the categoryId. To clear the labels of the given transactions, pass an empty array of label identifiers: '[]'. The result is a list of identifiers of only those transactions that have changed as a result of this service call.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ids Comma-separated list of identifiers of updated transactions
 * @return ApiEditMultipleTransactionsDeprecatedRequest
 */
func (a *TransactionsApiService) EditMultipleTransactionsDeprecated(ctx _context.Context, ids []int64) ApiEditMultipleTransactionsDeprecatedRequest {
	return ApiEditMultipleTransactionsDeprecatedRequest{
		ApiService: a,
		ctx: ctx,
		ids: ids,
	}
}

/*
 * Execute executes the request
 * @return IdentifierList
 */
func (a *TransactionsApiService) EditMultipleTransactionsDeprecatedExecute(r ApiEditMultipleTransactionsDeprecatedRequest) (IdentifierList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  IdentifierList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.EditMultipleTransactionsDeprecated")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions/{ids}"
	localVarPath = strings.Replace(localVarPath, "{"+"ids"+"}", _neturl.PathEscape(parameterToString(r.ids, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.updateTransactionsParams == nil {
		return localVarReturnValue, nil, reportError("updateTransactionsParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if r.xHTTPMethodOverride != nil {
		localVarHeaderParams["X-HTTP-Method-Override"] = parameterToString(*r.xHTTPMethodOverride, "")
	}
	if r.xRequestId != nil {
		localVarHeaderParams["X-Request-Id"] = parameterToString(*r.xRequestId, "")
	}
	// body params
	localVarPostBody = r.updateTransactionsParams
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiEditTransactionRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	id int64
	updateTransactionsParams *UpdateTransactionsParams
	xHTTPMethodOverride *string
	xRequestId *string
}

func (r ApiEditTransactionRequest) UpdateTransactionsParams(updateTransactionsParams UpdateTransactionsParams) ApiEditTransactionRequest {
	r.updateTransactionsParams = &updateTransactionsParams
	return r
}
func (r ApiEditTransactionRequest) XHTTPMethodOverride(xHTTPMethodOverride string) ApiEditTransactionRequest {
	r.xHTTPMethodOverride = &xHTTPMethodOverride
	return r
}
func (r ApiEditTransactionRequest) XRequestId(xRequestId string) ApiEditTransactionRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiEditTransactionRequest) Execute() (Transaction, *_nethttp.Response, error) {
	return r.ApiService.EditTransactionExecute(r)
}

/*
 * EditTransaction Edit a transaction
 * Change a transaction's fields. You can change the following fields: 'isNew=true|false' and/or 'isPotentialDuplicate=false' and/or 'categoryId=&lt;id&gt;' and/or 'labelIds=[&lt;ids&gt;]'. To clear a transaction's category (so that it is no longer categorized), pass the value '0' as the categoryId. To clear the labels of the given transaction, pass an empty array of label identifiers: '[]'.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Identifier of transaction
 * @return ApiEditTransactionRequest
 */
func (a *TransactionsApiService) EditTransaction(ctx _context.Context, id int64) ApiEditTransactionRequest {
	return ApiEditTransactionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Transaction
 */
func (a *TransactionsApiService) EditTransactionExecute(r ApiEditTransactionRequest) (Transaction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.EditTransaction")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.updateTransactionsParams == nil {
		return localVarReturnValue, nil, reportError("updateTransactionsParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	if r.xHTTPMethodOverride != nil {
		localVarHeaderParams["X-HTTP-Method-Override"] = parameterToString(*r.xHTTPMethodOverride, "")
	}
	if r.xRequestId != nil {
		localVarHeaderParams["X-Request-Id"] = parameterToString(*r.xRequestId, "")
	}
	// body params
	localVarPostBody = r.updateTransactionsParams
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiGetAndSearchAllTransactionsRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	view *string
	ids *[]int64
	search *string
	counterpart *string
	purpose *string
	accountIds *[]int64
	minBankBookingDate *string
	maxBankBookingDate *string
	minFinapiBookingDate *string
	maxFinapiBookingDate *string
	minAmount *float64
	maxAmount *float64
	direction *string
	labelIds *[]int64
	categoryIds *[]int64
	includeChildCategories *bool
	isNew *bool
	isPotentialDuplicate *bool
	isAdjustingEntry *bool
	minImportDate *string
	maxImportDate *string
	page *int32
	perPage *int32
	order *[]string
	xRequestId *string
}

func (r ApiGetAndSearchAllTransactionsRequest) View(view string) ApiGetAndSearchAllTransactionsRequest {
	r.view = &view
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) Ids(ids []int64) ApiGetAndSearchAllTransactionsRequest {
	r.ids = &ids
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) Search(search string) ApiGetAndSearchAllTransactionsRequest {
	r.search = &search
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) Counterpart(counterpart string) ApiGetAndSearchAllTransactionsRequest {
	r.counterpart = &counterpart
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) Purpose(purpose string) ApiGetAndSearchAllTransactionsRequest {
	r.purpose = &purpose
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) AccountIds(accountIds []int64) ApiGetAndSearchAllTransactionsRequest {
	r.accountIds = &accountIds
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) MinBankBookingDate(minBankBookingDate string) ApiGetAndSearchAllTransactionsRequest {
	r.minBankBookingDate = &minBankBookingDate
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) MaxBankBookingDate(maxBankBookingDate string) ApiGetAndSearchAllTransactionsRequest {
	r.maxBankBookingDate = &maxBankBookingDate
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) MinFinapiBookingDate(minFinapiBookingDate string) ApiGetAndSearchAllTransactionsRequest {
	r.minFinapiBookingDate = &minFinapiBookingDate
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) MaxFinapiBookingDate(maxFinapiBookingDate string) ApiGetAndSearchAllTransactionsRequest {
	r.maxFinapiBookingDate = &maxFinapiBookingDate
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) MinAmount(minAmount float64) ApiGetAndSearchAllTransactionsRequest {
	r.minAmount = &minAmount
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) MaxAmount(maxAmount float64) ApiGetAndSearchAllTransactionsRequest {
	r.maxAmount = &maxAmount
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) Direction(direction string) ApiGetAndSearchAllTransactionsRequest {
	r.direction = &direction
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) LabelIds(labelIds []int64) ApiGetAndSearchAllTransactionsRequest {
	r.labelIds = &labelIds
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) CategoryIds(categoryIds []int64) ApiGetAndSearchAllTransactionsRequest {
	r.categoryIds = &categoryIds
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) IncludeChildCategories(includeChildCategories bool) ApiGetAndSearchAllTransactionsRequest {
	r.includeChildCategories = &includeChildCategories
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) IsNew(isNew bool) ApiGetAndSearchAllTransactionsRequest {
	r.isNew = &isNew
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) IsPotentialDuplicate(isPotentialDuplicate bool) ApiGetAndSearchAllTransactionsRequest {
	r.isPotentialDuplicate = &isPotentialDuplicate
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) IsAdjustingEntry(isAdjustingEntry bool) ApiGetAndSearchAllTransactionsRequest {
	r.isAdjustingEntry = &isAdjustingEntry
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) MinImportDate(minImportDate string) ApiGetAndSearchAllTransactionsRequest {
	r.minImportDate = &minImportDate
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) MaxImportDate(maxImportDate string) ApiGetAndSearchAllTransactionsRequest {
	r.maxImportDate = &maxImportDate
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) Page(page int32) ApiGetAndSearchAllTransactionsRequest {
	r.page = &page
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) PerPage(perPage int32) ApiGetAndSearchAllTransactionsRequest {
	r.perPage = &perPage
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) Order(order []string) ApiGetAndSearchAllTransactionsRequest {
	r.order = &order
	return r
}
func (r ApiGetAndSearchAllTransactionsRequest) XRequestId(xRequestId string) ApiGetAndSearchAllTransactionsRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiGetAndSearchAllTransactionsRequest) Execute() (PageableTransactionList, *_nethttp.Response, error) {
	return r.ApiService.GetAndSearchAllTransactionsExecute(r)
}

/*
 * GetAndSearchAllTransactions Get and search all transactions
 * Get transactions of the user that is authorized by the access_token. Must pass the user's access_token. You can set optional search criteria to get only those transactions that you are interested in. If you do not specify any search criteria, then this service functions as a 'get all' service.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetAndSearchAllTransactionsRequest
 */
func (a *TransactionsApiService) GetAndSearchAllTransactions(ctx _context.Context) ApiGetAndSearchAllTransactionsRequest {
	return ApiGetAndSearchAllTransactionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PageableTransactionList
 */
func (a *TransactionsApiService) GetAndSearchAllTransactionsExecute(r ApiGetAndSearchAllTransactionsRequest) (PageableTransactionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PageableTransactionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.GetAndSearchAllTransactions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.view == nil {
		return localVarReturnValue, nil, reportError("view is required and must be specified")
	}

	if r.ids != nil {
		t := *r.ids
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("ids", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("ids", parameterToString(t, "multi"))
		}
	}
	localVarQueryParams.Add("view", parameterToString(*r.view, ""))
	if r.search != nil {
		localVarQueryParams.Add("search", parameterToString(*r.search, ""))
	}
	if r.counterpart != nil {
		localVarQueryParams.Add("counterpart", parameterToString(*r.counterpart, ""))
	}
	if r.purpose != nil {
		localVarQueryParams.Add("purpose", parameterToString(*r.purpose, ""))
	}
	if r.accountIds != nil {
		t := *r.accountIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("accountIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("accountIds", parameterToString(t, "multi"))
		}
	}
	if r.minBankBookingDate != nil {
		localVarQueryParams.Add("minBankBookingDate", parameterToString(*r.minBankBookingDate, ""))
	}
	if r.maxBankBookingDate != nil {
		localVarQueryParams.Add("maxBankBookingDate", parameterToString(*r.maxBankBookingDate, ""))
	}
	if r.minFinapiBookingDate != nil {
		localVarQueryParams.Add("minFinapiBookingDate", parameterToString(*r.minFinapiBookingDate, ""))
	}
	if r.maxFinapiBookingDate != nil {
		localVarQueryParams.Add("maxFinapiBookingDate", parameterToString(*r.maxFinapiBookingDate, ""))
	}
	if r.minAmount != nil {
		localVarQueryParams.Add("minAmount", parameterToString(*r.minAmount, ""))
	}
	if r.maxAmount != nil {
		localVarQueryParams.Add("maxAmount", parameterToString(*r.maxAmount, ""))
	}
	if r.direction != nil {
		localVarQueryParams.Add("direction", parameterToString(*r.direction, ""))
	}
	if r.labelIds != nil {
		t := *r.labelIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("labelIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("labelIds", parameterToString(t, "multi"))
		}
	}
	if r.categoryIds != nil {
		t := *r.categoryIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("categoryIds", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("categoryIds", parameterToString(t, "multi"))
		}
	}
	if r.includeChildCategories != nil {
		localVarQueryParams.Add("includeChildCategories", parameterToString(*r.includeChildCategories, ""))
	}
	if r.isNew != nil {
		localVarQueryParams.Add("isNew", parameterToString(*r.isNew, ""))
	}
	if r.isPotentialDuplicate != nil {
		localVarQueryParams.Add("isPotentialDuplicate", parameterToString(*r.isPotentialDuplicate, ""))
	}
	if r.isAdjustingEntry != nil {
		localVarQueryParams.Add("isAdjustingEntry", parameterToString(*r.isAdjustingEntry, ""))
	}
	if r.minImportDate != nil {
		localVarQueryParams.Add("minImportDate", parameterToString(*r.minImportDate, ""))
	}
	if r.maxImportDate != nil {
		localVarQueryParams.Add("maxImportDate", parameterToString(*r.maxImportDate, ""))
	}
	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.perPage != nil {
		localVarQueryParams.Add("perPage", parameterToString(*r.perPage, ""))
	}
	if r.order != nil {
		t := *r.order
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				localVarQueryParams.Add("order", parameterToString(s.Index(i), "multi"))
			}
		} else {
			localVarQueryParams.Add("order", parameterToString(t, "multi"))
		}
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

type ApiGetMultipleTransactionsRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	ids []int64
	xRequestId *string
}

func (r ApiGetMultipleTransactionsRequest) XRequestId(xRequestId string) ApiGetMultipleTransactionsRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiGetMultipleTransactionsRequest) Execute() (TransactionList, *_nethttp.Response, error) {
	return r.ApiService.GetMultipleTransactionsExecute(r)
}

/*
 * GetMultipleTransactions Get multiple transactions
 * THIS SERVICE IS DEPRECATED AND WILL BE REMOVED.<br/>Please refer to the 'Get and search all transactions' service instead if you want to get multiple transactions, and pass a comma-separated list of identifiers with the 'ids' parameter.<br/><br/>Get a list of multiple transactions of the user that is authorized by the access_token. Must pass the transactions' identifiers and the user's access_token. Transactions whose identifiers do not exist or do not relate to the authorized user will not be contained in the result (If this applies to all of the given identifiers, then the result will be an empty list).
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param ids Comma-separated list of identifiers of requested transactions
 * @return ApiGetMultipleTransactionsRequest
 */
func (a *TransactionsApiService) GetMultipleTransactions(ctx _context.Context, ids []int64) ApiGetMultipleTransactionsRequest {
	return ApiGetMultipleTransactionsRequest{
		ApiService: a,
		ctx: ctx,
		ids: ids,
	}
}

/*
 * Execute executes the request
 * @return TransactionList
 */
func (a *TransactionsApiService) GetMultipleTransactionsExecute(r ApiGetMultipleTransactionsRequest) (TransactionList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  TransactionList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.GetMultipleTransactions")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions/{ids}"
	localVarPath = strings.Replace(localVarPath, "{"+"ids"+"}", _neturl.PathEscape(parameterToString(r.ids, "csv")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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

type ApiGetTransactionRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	id int64
	xRequestId *string
}

func (r ApiGetTransactionRequest) XRequestId(xRequestId string) ApiGetTransactionRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiGetTransactionRequest) Execute() (Transaction, *_nethttp.Response, error) {
	return r.ApiService.GetTransactionExecute(r)
}

/*
 * GetTransaction Get a transaction
 * Get a single transaction of the user that is authorized by the access_token. Must pass the transaction's identifier and the user's access_token.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Identifier of transaction
 * @return ApiGetTransactionRequest
 */
func (a *TransactionsApiService) GetTransaction(ctx _context.Context, id int64) ApiGetTransactionRequest {
	return ApiGetTransactionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Transaction
 */
func (a *TransactionsApiService) GetTransactionExecute(r ApiGetTransactionRequest) (Transaction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.GetTransaction")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiRestoreTransactionRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	id int64
	xRequestId *string
}

func (r ApiRestoreTransactionRequest) XRequestId(xRequestId string) ApiRestoreTransactionRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiRestoreTransactionRequest) Execute() (Transaction, *_nethttp.Response, error) {
	return r.ApiService.RestoreTransactionExecute(r)
}

/*
 * RestoreTransaction Restore a transaction
 * Restore a previously split transaction. Removes all of its sub-transactions.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Transaction identifier
 * @return ApiRestoreTransactionRequest
 */
func (a *TransactionsApiService) RestoreTransaction(ctx _context.Context, id int64) ApiRestoreTransactionRequest {
	return ApiRestoreTransactionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Transaction
 */
func (a *TransactionsApiService) RestoreTransactionExecute(r ApiRestoreTransactionRequest) (Transaction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.RestoreTransaction")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions/{id}/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiSplitTransactionRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	id int64
	splitTransactionsParams *SplitTransactionsParams
	xRequestId *string
}

func (r ApiSplitTransactionRequest) SplitTransactionsParams(splitTransactionsParams SplitTransactionsParams) ApiSplitTransactionRequest {
	r.splitTransactionsParams = &splitTransactionsParams
	return r
}
func (r ApiSplitTransactionRequest) XRequestId(xRequestId string) ApiSplitTransactionRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiSplitTransactionRequest) Execute() (Transaction, *_nethttp.Response, error) {
	return r.ApiService.SplitTransactionExecute(r)
}

/*
 * SplitTransaction Split a transaction
 * Split a transaction into several logical sub-transactions. If the given transaction is split already, all its current sub-transactions will get deleted before the new sub-transactions will get created.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id Transaction identifier
 * @return ApiSplitTransactionRequest
 */
func (a *TransactionsApiService) SplitTransaction(ctx _context.Context, id int64) ApiSplitTransactionRequest {
	return ApiSplitTransactionRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Transaction
 */
func (a *TransactionsApiService) SplitTransactionExecute(r ApiSplitTransactionRequest) (Transaction, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Transaction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.SplitTransaction")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions/{id}/split"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.splitTransactionsParams == nil {
		return localVarReturnValue, nil, reportError("splitTransactionsParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.splitTransactionsParams
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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

type ApiTriggerCategorizationRequest struct {
	ctx _context.Context
	ApiService *TransactionsApiService
	triggerCategorizationParams *TriggerCategorizationParams
	xRequestId *string
}

func (r ApiTriggerCategorizationRequest) TriggerCategorizationParams(triggerCategorizationParams TriggerCategorizationParams) ApiTriggerCategorizationRequest {
	r.triggerCategorizationParams = &triggerCategorizationParams
	return r
}
func (r ApiTriggerCategorizationRequest) XRequestId(xRequestId string) ApiTriggerCategorizationRequest {
	r.xRequestId = &xRequestId
	return r
}

func (r ApiTriggerCategorizationRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.TriggerCategorizationExecute(r)
}

/*
 * TriggerCategorization Trigger categorization
 * Triggers finAPI's background categorization process for all uncategorized transactions of the given bank connection(s) (or of all of the user's bank connections, if no bank connection identifiers are passed). The service returns as soon as the categorizations are scheduled. At this point, the bank connections will have their 'categorizationStatus' set to 'PENDING'. Use the service "Get a bank connection" or "Get all bank connections" to check when the categorization has finished (this is the case when the categorizationStatus has switched to 'READY').<br/><br/>Note that if at least one of the target bank connections is currently locked at the time when you call this service (i.e. the bank connection is currently being updated, or another categorization is already scheduled for it), then no categorization will be triggered at all and the service will respond with HTTP code 422.<br/><br/>Please also note:<br/>&bull; finAPI's background categorization process is executed automatically whenever you import or update a bank connection (though in case of update, it will categorize only the new transactions, and not re-run categorization for previously imported transactions). This means that in general you do not have to call this service after an import or update. Use this service only when you wish to re-run the categorization of all existing uncategorized transactions.<br/>&bull; if you wish to just manually assign categories to transactions, please use the service "Edit a transaction" or "Edit multiple transactions" instead.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiTriggerCategorizationRequest
 */
func (a *TransactionsApiService) TriggerCategorization(ctx _context.Context) ApiTriggerCategorizationRequest {
	return ApiTriggerCategorizationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *TransactionsApiService) TriggerCategorizationExecute(r ApiTriggerCategorizationRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransactionsApiService.TriggerCategorization")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/transactions/triggerCategorization"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.triggerCategorizationParams == nil {
		return nil, reportError("triggerCategorizationParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.triggerCategorizationParams
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v BadCredentialsError
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ErrorMessage
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
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
