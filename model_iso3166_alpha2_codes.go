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

// ISO3166Alpha2Codes the model 'ISO3166Alpha2Codes'
type ISO3166Alpha2Codes string

// List of ISO3166Alpha2Codes
const (
	ISO3166ALPHA2CODES_AD ISO3166Alpha2Codes = "AD"
	ISO3166ALPHA2CODES_AE ISO3166Alpha2Codes = "AE"
	ISO3166ALPHA2CODES_AF ISO3166Alpha2Codes = "AF"
	ISO3166ALPHA2CODES_AG ISO3166Alpha2Codes = "AG"
	ISO3166ALPHA2CODES_AI ISO3166Alpha2Codes = "AI"
	ISO3166ALPHA2CODES_AL ISO3166Alpha2Codes = "AL"
	ISO3166ALPHA2CODES_AM ISO3166Alpha2Codes = "AM"
	ISO3166ALPHA2CODES_AO ISO3166Alpha2Codes = "AO"
	ISO3166ALPHA2CODES_AQ ISO3166Alpha2Codes = "AQ"
	ISO3166ALPHA2CODES_AR ISO3166Alpha2Codes = "AR"
	ISO3166ALPHA2CODES_AS ISO3166Alpha2Codes = "AS"
	ISO3166ALPHA2CODES_AT ISO3166Alpha2Codes = "AT"
	ISO3166ALPHA2CODES_AU ISO3166Alpha2Codes = "AU"
	ISO3166ALPHA2CODES_AW ISO3166Alpha2Codes = "AW"
	ISO3166ALPHA2CODES_AX ISO3166Alpha2Codes = "AX"
	ISO3166ALPHA2CODES_AZ ISO3166Alpha2Codes = "AZ"
	ISO3166ALPHA2CODES_BA ISO3166Alpha2Codes = "BA"
	ISO3166ALPHA2CODES_BB ISO3166Alpha2Codes = "BB"
	ISO3166ALPHA2CODES_BD ISO3166Alpha2Codes = "BD"
	ISO3166ALPHA2CODES_BE ISO3166Alpha2Codes = "BE"
	ISO3166ALPHA2CODES_BF ISO3166Alpha2Codes = "BF"
	ISO3166ALPHA2CODES_BG ISO3166Alpha2Codes = "BG"
	ISO3166ALPHA2CODES_BH ISO3166Alpha2Codes = "BH"
	ISO3166ALPHA2CODES_BI ISO3166Alpha2Codes = "BI"
	ISO3166ALPHA2CODES_BJ ISO3166Alpha2Codes = "BJ"
	ISO3166ALPHA2CODES_BL ISO3166Alpha2Codes = "BL"
	ISO3166ALPHA2CODES_BM ISO3166Alpha2Codes = "BM"
	ISO3166ALPHA2CODES_BN ISO3166Alpha2Codes = "BN"
	ISO3166ALPHA2CODES_BO ISO3166Alpha2Codes = "BO"
	ISO3166ALPHA2CODES_BQ ISO3166Alpha2Codes = "BQ"
	ISO3166ALPHA2CODES_BR ISO3166Alpha2Codes = "BR"
	ISO3166ALPHA2CODES_BS ISO3166Alpha2Codes = "BS"
	ISO3166ALPHA2CODES_BT ISO3166Alpha2Codes = "BT"
	ISO3166ALPHA2CODES_BV ISO3166Alpha2Codes = "BV"
	ISO3166ALPHA2CODES_BW ISO3166Alpha2Codes = "BW"
	ISO3166ALPHA2CODES_BY ISO3166Alpha2Codes = "BY"
	ISO3166ALPHA2CODES_BZ ISO3166Alpha2Codes = "BZ"
	ISO3166ALPHA2CODES_CA ISO3166Alpha2Codes = "CA"
	ISO3166ALPHA2CODES_CC ISO3166Alpha2Codes = "CC"
	ISO3166ALPHA2CODES_CD ISO3166Alpha2Codes = "CD"
	ISO3166ALPHA2CODES_CF ISO3166Alpha2Codes = "CF"
	ISO3166ALPHA2CODES_CG ISO3166Alpha2Codes = "CG"
	ISO3166ALPHA2CODES_CH ISO3166Alpha2Codes = "CH"
	ISO3166ALPHA2CODES_CI ISO3166Alpha2Codes = "CI"
	ISO3166ALPHA2CODES_CK ISO3166Alpha2Codes = "CK"
	ISO3166ALPHA2CODES_CL ISO3166Alpha2Codes = "CL"
	ISO3166ALPHA2CODES_CM ISO3166Alpha2Codes = "CM"
	ISO3166ALPHA2CODES_CN ISO3166Alpha2Codes = "CN"
	ISO3166ALPHA2CODES_CO ISO3166Alpha2Codes = "CO"
	ISO3166ALPHA2CODES_CR ISO3166Alpha2Codes = "CR"
	ISO3166ALPHA2CODES_CU ISO3166Alpha2Codes = "CU"
	ISO3166ALPHA2CODES_CV ISO3166Alpha2Codes = "CV"
	ISO3166ALPHA2CODES_CW ISO3166Alpha2Codes = "CW"
	ISO3166ALPHA2CODES_CX ISO3166Alpha2Codes = "CX"
	ISO3166ALPHA2CODES_CY ISO3166Alpha2Codes = "CY"
	ISO3166ALPHA2CODES_CZ ISO3166Alpha2Codes = "CZ"
	ISO3166ALPHA2CODES_DE ISO3166Alpha2Codes = "DE"
	ISO3166ALPHA2CODES_DJ ISO3166Alpha2Codes = "DJ"
	ISO3166ALPHA2CODES_DK ISO3166Alpha2Codes = "DK"
	ISO3166ALPHA2CODES_DM ISO3166Alpha2Codes = "DM"
	ISO3166ALPHA2CODES_DO ISO3166Alpha2Codes = "DO"
	ISO3166ALPHA2CODES_DZ ISO3166Alpha2Codes = "DZ"
	ISO3166ALPHA2CODES_EC ISO3166Alpha2Codes = "EC"
	ISO3166ALPHA2CODES_EE ISO3166Alpha2Codes = "EE"
	ISO3166ALPHA2CODES_EG ISO3166Alpha2Codes = "EG"
	ISO3166ALPHA2CODES_EH ISO3166Alpha2Codes = "EH"
	ISO3166ALPHA2CODES_ER ISO3166Alpha2Codes = "ER"
	ISO3166ALPHA2CODES_ES ISO3166Alpha2Codes = "ES"
	ISO3166ALPHA2CODES_ET ISO3166Alpha2Codes = "ET"
	ISO3166ALPHA2CODES_FI ISO3166Alpha2Codes = "FI"
	ISO3166ALPHA2CODES_FJ ISO3166Alpha2Codes = "FJ"
	ISO3166ALPHA2CODES_FK ISO3166Alpha2Codes = "FK"
	ISO3166ALPHA2CODES_FM ISO3166Alpha2Codes = "FM"
	ISO3166ALPHA2CODES_FO ISO3166Alpha2Codes = "FO"
	ISO3166ALPHA2CODES_FR ISO3166Alpha2Codes = "FR"
	ISO3166ALPHA2CODES_GA ISO3166Alpha2Codes = "GA"
	ISO3166ALPHA2CODES_GB ISO3166Alpha2Codes = "GB"
	ISO3166ALPHA2CODES_GD ISO3166Alpha2Codes = "GD"
	ISO3166ALPHA2CODES_GE ISO3166Alpha2Codes = "GE"
	ISO3166ALPHA2CODES_GF ISO3166Alpha2Codes = "GF"
	ISO3166ALPHA2CODES_GG ISO3166Alpha2Codes = "GG"
	ISO3166ALPHA2CODES_GH ISO3166Alpha2Codes = "GH"
	ISO3166ALPHA2CODES_GI ISO3166Alpha2Codes = "GI"
	ISO3166ALPHA2CODES_GL ISO3166Alpha2Codes = "GL"
	ISO3166ALPHA2CODES_GM ISO3166Alpha2Codes = "GM"
	ISO3166ALPHA2CODES_GN ISO3166Alpha2Codes = "GN"
	ISO3166ALPHA2CODES_GP ISO3166Alpha2Codes = "GP"
	ISO3166ALPHA2CODES_GQ ISO3166Alpha2Codes = "GQ"
	ISO3166ALPHA2CODES_GR ISO3166Alpha2Codes = "GR"
	ISO3166ALPHA2CODES_GS ISO3166Alpha2Codes = "GS"
	ISO3166ALPHA2CODES_GT ISO3166Alpha2Codes = "GT"
	ISO3166ALPHA2CODES_GU ISO3166Alpha2Codes = "GU"
	ISO3166ALPHA2CODES_GW ISO3166Alpha2Codes = "GW"
	ISO3166ALPHA2CODES_GY ISO3166Alpha2Codes = "GY"
	ISO3166ALPHA2CODES_HK ISO3166Alpha2Codes = "HK"
	ISO3166ALPHA2CODES_HM ISO3166Alpha2Codes = "HM"
	ISO3166ALPHA2CODES_HN ISO3166Alpha2Codes = "HN"
	ISO3166ALPHA2CODES_HR ISO3166Alpha2Codes = "HR"
	ISO3166ALPHA2CODES_HT ISO3166Alpha2Codes = "HT"
	ISO3166ALPHA2CODES_HU ISO3166Alpha2Codes = "HU"
	ISO3166ALPHA2CODES_ID ISO3166Alpha2Codes = "ID"
	ISO3166ALPHA2CODES_IE ISO3166Alpha2Codes = "IE"
	ISO3166ALPHA2CODES_IL ISO3166Alpha2Codes = "IL"
	ISO3166ALPHA2CODES_IM ISO3166Alpha2Codes = "IM"
	ISO3166ALPHA2CODES_IN ISO3166Alpha2Codes = "IN"
	ISO3166ALPHA2CODES_IO ISO3166Alpha2Codes = "IO"
	ISO3166ALPHA2CODES_IQ ISO3166Alpha2Codes = "IQ"
	ISO3166ALPHA2CODES_IR ISO3166Alpha2Codes = "IR"
	ISO3166ALPHA2CODES_IS ISO3166Alpha2Codes = "IS"
	ISO3166ALPHA2CODES_IT ISO3166Alpha2Codes = "IT"
	ISO3166ALPHA2CODES_JE ISO3166Alpha2Codes = "JE"
	ISO3166ALPHA2CODES_JM ISO3166Alpha2Codes = "JM"
	ISO3166ALPHA2CODES_JO ISO3166Alpha2Codes = "JO"
	ISO3166ALPHA2CODES_JP ISO3166Alpha2Codes = "JP"
	ISO3166ALPHA2CODES_KE ISO3166Alpha2Codes = "KE"
	ISO3166ALPHA2CODES_KG ISO3166Alpha2Codes = "KG"
	ISO3166ALPHA2CODES_KH ISO3166Alpha2Codes = "KH"
	ISO3166ALPHA2CODES_KI ISO3166Alpha2Codes = "KI"
	ISO3166ALPHA2CODES_KM ISO3166Alpha2Codes = "KM"
	ISO3166ALPHA2CODES_KN ISO3166Alpha2Codes = "KN"
	ISO3166ALPHA2CODES_KP ISO3166Alpha2Codes = "KP"
	ISO3166ALPHA2CODES_KR ISO3166Alpha2Codes = "KR"
	ISO3166ALPHA2CODES_KW ISO3166Alpha2Codes = "KW"
	ISO3166ALPHA2CODES_KY ISO3166Alpha2Codes = "KY"
	ISO3166ALPHA2CODES_KZ ISO3166Alpha2Codes = "KZ"
	ISO3166ALPHA2CODES_LA ISO3166Alpha2Codes = "LA"
	ISO3166ALPHA2CODES_LB ISO3166Alpha2Codes = "LB"
	ISO3166ALPHA2CODES_LC ISO3166Alpha2Codes = "LC"
	ISO3166ALPHA2CODES_LI ISO3166Alpha2Codes = "LI"
	ISO3166ALPHA2CODES_LK ISO3166Alpha2Codes = "LK"
	ISO3166ALPHA2CODES_LR ISO3166Alpha2Codes = "LR"
	ISO3166ALPHA2CODES_LS ISO3166Alpha2Codes = "LS"
	ISO3166ALPHA2CODES_LT ISO3166Alpha2Codes = "LT"
	ISO3166ALPHA2CODES_LU ISO3166Alpha2Codes = "LU"
	ISO3166ALPHA2CODES_LV ISO3166Alpha2Codes = "LV"
	ISO3166ALPHA2CODES_LY ISO3166Alpha2Codes = "LY"
	ISO3166ALPHA2CODES_MA ISO3166Alpha2Codes = "MA"
	ISO3166ALPHA2CODES_MC ISO3166Alpha2Codes = "MC"
	ISO3166ALPHA2CODES_MD ISO3166Alpha2Codes = "MD"
	ISO3166ALPHA2CODES_ME ISO3166Alpha2Codes = "ME"
	ISO3166ALPHA2CODES_MF ISO3166Alpha2Codes = "MF"
	ISO3166ALPHA2CODES_MG ISO3166Alpha2Codes = "MG"
	ISO3166ALPHA2CODES_MH ISO3166Alpha2Codes = "MH"
	ISO3166ALPHA2CODES_MK ISO3166Alpha2Codes = "MK"
	ISO3166ALPHA2CODES_ML ISO3166Alpha2Codes = "ML"
	ISO3166ALPHA2CODES_MM ISO3166Alpha2Codes = "MM"
	ISO3166ALPHA2CODES_MN ISO3166Alpha2Codes = "MN"
	ISO3166ALPHA2CODES_MO ISO3166Alpha2Codes = "MO"
	ISO3166ALPHA2CODES_MP ISO3166Alpha2Codes = "MP"
	ISO3166ALPHA2CODES_MQ ISO3166Alpha2Codes = "MQ"
	ISO3166ALPHA2CODES_MR ISO3166Alpha2Codes = "MR"
	ISO3166ALPHA2CODES_MS ISO3166Alpha2Codes = "MS"
	ISO3166ALPHA2CODES_MT ISO3166Alpha2Codes = "MT"
	ISO3166ALPHA2CODES_MU ISO3166Alpha2Codes = "MU"
	ISO3166ALPHA2CODES_MV ISO3166Alpha2Codes = "MV"
	ISO3166ALPHA2CODES_MW ISO3166Alpha2Codes = "MW"
	ISO3166ALPHA2CODES_MX ISO3166Alpha2Codes = "MX"
	ISO3166ALPHA2CODES_MY ISO3166Alpha2Codes = "MY"
	ISO3166ALPHA2CODES_MZ ISO3166Alpha2Codes = "MZ"
	ISO3166ALPHA2CODES_NA ISO3166Alpha2Codes = "NA"
	ISO3166ALPHA2CODES_NC ISO3166Alpha2Codes = "NC"
	ISO3166ALPHA2CODES_NE ISO3166Alpha2Codes = "NE"
	ISO3166ALPHA2CODES_NF ISO3166Alpha2Codes = "NF"
	ISO3166ALPHA2CODES_NG ISO3166Alpha2Codes = "NG"
	ISO3166ALPHA2CODES_NI ISO3166Alpha2Codes = "NI"
	ISO3166ALPHA2CODES_NL ISO3166Alpha2Codes = "NL"
	ISO3166ALPHA2CODES_NO ISO3166Alpha2Codes = "NO"
	ISO3166ALPHA2CODES_NP ISO3166Alpha2Codes = "NP"
	ISO3166ALPHA2CODES_NR ISO3166Alpha2Codes = "NR"
	ISO3166ALPHA2CODES_NU ISO3166Alpha2Codes = "NU"
	ISO3166ALPHA2CODES_NZ ISO3166Alpha2Codes = "NZ"
	ISO3166ALPHA2CODES_OM ISO3166Alpha2Codes = "OM"
	ISO3166ALPHA2CODES_PA ISO3166Alpha2Codes = "PA"
	ISO3166ALPHA2CODES_PE ISO3166Alpha2Codes = "PE"
	ISO3166ALPHA2CODES_PF ISO3166Alpha2Codes = "PF"
	ISO3166ALPHA2CODES_PG ISO3166Alpha2Codes = "PG"
	ISO3166ALPHA2CODES_PH ISO3166Alpha2Codes = "PH"
	ISO3166ALPHA2CODES_PK ISO3166Alpha2Codes = "PK"
	ISO3166ALPHA2CODES_PL ISO3166Alpha2Codes = "PL"
	ISO3166ALPHA2CODES_PM ISO3166Alpha2Codes = "PM"
	ISO3166ALPHA2CODES_PN ISO3166Alpha2Codes = "PN"
	ISO3166ALPHA2CODES_PR ISO3166Alpha2Codes = "PR"
	ISO3166ALPHA2CODES_PS ISO3166Alpha2Codes = "PS"
	ISO3166ALPHA2CODES_PT ISO3166Alpha2Codes = "PT"
	ISO3166ALPHA2CODES_PW ISO3166Alpha2Codes = "PW"
	ISO3166ALPHA2CODES_PY ISO3166Alpha2Codes = "PY"
	ISO3166ALPHA2CODES_QA ISO3166Alpha2Codes = "QA"
	ISO3166ALPHA2CODES_RE ISO3166Alpha2Codes = "RE"
	ISO3166ALPHA2CODES_RO ISO3166Alpha2Codes = "RO"
	ISO3166ALPHA2CODES_RS ISO3166Alpha2Codes = "RS"
	ISO3166ALPHA2CODES_RU ISO3166Alpha2Codes = "RU"
	ISO3166ALPHA2CODES_RW ISO3166Alpha2Codes = "RW"
	ISO3166ALPHA2CODES_SA ISO3166Alpha2Codes = "SA"
	ISO3166ALPHA2CODES_SB ISO3166Alpha2Codes = "SB"
	ISO3166ALPHA2CODES_SC ISO3166Alpha2Codes = "SC"
	ISO3166ALPHA2CODES_SD ISO3166Alpha2Codes = "SD"
	ISO3166ALPHA2CODES_SE ISO3166Alpha2Codes = "SE"
	ISO3166ALPHA2CODES_SG ISO3166Alpha2Codes = "SG"
	ISO3166ALPHA2CODES_SH ISO3166Alpha2Codes = "SH"
	ISO3166ALPHA2CODES_SI ISO3166Alpha2Codes = "SI"
	ISO3166ALPHA2CODES_SJ ISO3166Alpha2Codes = "SJ"
	ISO3166ALPHA2CODES_SK ISO3166Alpha2Codes = "SK"
	ISO3166ALPHA2CODES_SL ISO3166Alpha2Codes = "SL"
	ISO3166ALPHA2CODES_SM ISO3166Alpha2Codes = "SM"
	ISO3166ALPHA2CODES_SN ISO3166Alpha2Codes = "SN"
	ISO3166ALPHA2CODES_SO ISO3166Alpha2Codes = "SO"
	ISO3166ALPHA2CODES_SR ISO3166Alpha2Codes = "SR"
	ISO3166ALPHA2CODES_SS ISO3166Alpha2Codes = "SS"
	ISO3166ALPHA2CODES_ST ISO3166Alpha2Codes = "ST"
	ISO3166ALPHA2CODES_SV ISO3166Alpha2Codes = "SV"
	ISO3166ALPHA2CODES_SX ISO3166Alpha2Codes = "SX"
	ISO3166ALPHA2CODES_SY ISO3166Alpha2Codes = "SY"
	ISO3166ALPHA2CODES_SZ ISO3166Alpha2Codes = "SZ"
	ISO3166ALPHA2CODES_TC ISO3166Alpha2Codes = "TC"
	ISO3166ALPHA2CODES_TD ISO3166Alpha2Codes = "TD"
	ISO3166ALPHA2CODES_TF ISO3166Alpha2Codes = "TF"
	ISO3166ALPHA2CODES_TG ISO3166Alpha2Codes = "TG"
	ISO3166ALPHA2CODES_TH ISO3166Alpha2Codes = "TH"
	ISO3166ALPHA2CODES_TJ ISO3166Alpha2Codes = "TJ"
	ISO3166ALPHA2CODES_TK ISO3166Alpha2Codes = "TK"
	ISO3166ALPHA2CODES_TL ISO3166Alpha2Codes = "TL"
	ISO3166ALPHA2CODES_TM ISO3166Alpha2Codes = "TM"
	ISO3166ALPHA2CODES_TN ISO3166Alpha2Codes = "TN"
	ISO3166ALPHA2CODES_TO ISO3166Alpha2Codes = "TO"
	ISO3166ALPHA2CODES_TR ISO3166Alpha2Codes = "TR"
	ISO3166ALPHA2CODES_TT ISO3166Alpha2Codes = "TT"
	ISO3166ALPHA2CODES_TV ISO3166Alpha2Codes = "TV"
	ISO3166ALPHA2CODES_TW ISO3166Alpha2Codes = "TW"
	ISO3166ALPHA2CODES_TZ ISO3166Alpha2Codes = "TZ"
	ISO3166ALPHA2CODES_UA ISO3166Alpha2Codes = "UA"
	ISO3166ALPHA2CODES_UG ISO3166Alpha2Codes = "UG"
	ISO3166ALPHA2CODES_UM ISO3166Alpha2Codes = "UM"
	ISO3166ALPHA2CODES_US ISO3166Alpha2Codes = "US"
	ISO3166ALPHA2CODES_UY ISO3166Alpha2Codes = "UY"
	ISO3166ALPHA2CODES_UZ ISO3166Alpha2Codes = "UZ"
	ISO3166ALPHA2CODES_VA ISO3166Alpha2Codes = "VA"
	ISO3166ALPHA2CODES_VC ISO3166Alpha2Codes = "VC"
	ISO3166ALPHA2CODES_VE ISO3166Alpha2Codes = "VE"
	ISO3166ALPHA2CODES_VG ISO3166Alpha2Codes = "VG"
	ISO3166ALPHA2CODES_VI ISO3166Alpha2Codes = "VI"
	ISO3166ALPHA2CODES_VN ISO3166Alpha2Codes = "VN"
	ISO3166ALPHA2CODES_VU ISO3166Alpha2Codes = "VU"
	ISO3166ALPHA2CODES_WF ISO3166Alpha2Codes = "WF"
	ISO3166ALPHA2CODES_WS ISO3166Alpha2Codes = "WS"
	ISO3166ALPHA2CODES_XK ISO3166Alpha2Codes = "XK"
	ISO3166ALPHA2CODES_YE ISO3166Alpha2Codes = "YE"
	ISO3166ALPHA2CODES_YT ISO3166Alpha2Codes = "YT"
	ISO3166ALPHA2CODES_ZA ISO3166Alpha2Codes = "ZA"
	ISO3166ALPHA2CODES_ZM ISO3166Alpha2Codes = "ZM"
	ISO3166ALPHA2CODES_ZW ISO3166Alpha2Codes = "ZW"
)

func (v *ISO3166Alpha2Codes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ISO3166Alpha2Codes(value)
	for _, existing := range []ISO3166Alpha2Codes{ "AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "XK", "YE", "YT", "ZA", "ZM", "ZW",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ISO3166Alpha2Codes", value)
}

// Ptr returns reference to ISO3166Alpha2Codes value
func (v ISO3166Alpha2Codes) Ptr() *ISO3166Alpha2Codes {
	return &v
}

type NullableISO3166Alpha2Codes struct {
	value *ISO3166Alpha2Codes
	isSet bool
}

func (v NullableISO3166Alpha2Codes) Get() *ISO3166Alpha2Codes {
	return v.value
}

func (v *NullableISO3166Alpha2Codes) Set(val *ISO3166Alpha2Codes) {
	v.value = val
	v.isSet = true
}

func (v NullableISO3166Alpha2Codes) IsSet() bool {
	return v.isSet
}

func (v *NullableISO3166Alpha2Codes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableISO3166Alpha2Codes(val *ISO3166Alpha2Codes) *NullableISO3166Alpha2Codes {
	return &NullableISO3166Alpha2Codes{value: val, isSet: true}
}

func (v NullableISO3166Alpha2Codes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableISO3166Alpha2Codes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

