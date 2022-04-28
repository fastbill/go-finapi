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

// Currency the model 'Currency'
type Currency string

// List of Currency
const (
	CURRENCY_AED Currency = "AED"
	CURRENCY_AFN Currency = "AFN"
	CURRENCY_ALL Currency = "ALL"
	CURRENCY_AMD Currency = "AMD"
	CURRENCY_ANG Currency = "ANG"
	CURRENCY_AOA Currency = "AOA"
	CURRENCY_ARS Currency = "ARS"
	CURRENCY_AUD Currency = "AUD"
	CURRENCY_AWG Currency = "AWG"
	CURRENCY_AZN Currency = "AZN"
	CURRENCY_BAM Currency = "BAM"
	CURRENCY_BBD Currency = "BBD"
	CURRENCY_BDT Currency = "BDT"
	CURRENCY_BGN Currency = "BGN"
	CURRENCY_BHD Currency = "BHD"
	CURRENCY_BIF Currency = "BIF"
	CURRENCY_BMD Currency = "BMD"
	CURRENCY_BND Currency = "BND"
	CURRENCY_BOB Currency = "BOB"
	CURRENCY_BOV Currency = "BOV"
	CURRENCY_BRL Currency = "BRL"
	CURRENCY_BSD Currency = "BSD"
	CURRENCY_BTN Currency = "BTN"
	CURRENCY_BWP Currency = "BWP"
	CURRENCY_BYN Currency = "BYN"
	CURRENCY_BZD Currency = "BZD"
	CURRENCY_CAD Currency = "CAD"
	CURRENCY_CDF Currency = "CDF"
	CURRENCY_CHE Currency = "CHE"
	CURRENCY_CHF Currency = "CHF"
	CURRENCY_CHN Currency = "CHN"
	CURRENCY_CHW Currency = "CHW"
	CURRENCY_CLF Currency = "CLF"
	CURRENCY_CLP Currency = "CLP"
	CURRENCY_CNY Currency = "CNY"
	CURRENCY_COP Currency = "COP"
	CURRENCY_COU Currency = "COU"
	CURRENCY_CRC Currency = "CRC"
	CURRENCY_CUC Currency = "CUC"
	CURRENCY_CUP Currency = "CUP"
	CURRENCY_CVE Currency = "CVE"
	CURRENCY_CZK Currency = "CZK"
	CURRENCY_DJF Currency = "DJF"
	CURRENCY_DKK Currency = "DKK"
	CURRENCY_DOP Currency = "DOP"
	CURRENCY_DZD Currency = "DZD"
	CURRENCY_EGP Currency = "EGP"
	CURRENCY_ERN Currency = "ERN"
	CURRENCY_ETB Currency = "ETB"
	CURRENCY_EUR Currency = "EUR"
	CURRENCY_FJD Currency = "FJD"
	CURRENCY_FKP Currency = "FKP"
	CURRENCY_GBP Currency = "GBP"
	CURRENCY_GEL Currency = "GEL"
	CURRENCY_GGP Currency = "GGP"
	CURRENCY_GHS Currency = "GHS"
	CURRENCY_GIP Currency = "GIP"
	CURRENCY_GMD Currency = "GMD"
	CURRENCY_GNF Currency = "GNF"
	CURRENCY_GTQ Currency = "GTQ"
	CURRENCY_GYD Currency = "GYD"
	CURRENCY_HKD Currency = "HKD"
	CURRENCY_HNL Currency = "HNL"
	CURRENCY_HRK Currency = "HRK"
	CURRENCY_HTG Currency = "HTG"
	CURRENCY_HUF Currency = "HUF"
	CURRENCY_IDR Currency = "IDR"
	CURRENCY_ILS Currency = "ILS"
	CURRENCY_IMP Currency = "IMP"
	CURRENCY_INR Currency = "INR"
	CURRENCY_IQD Currency = "IQD"
	CURRENCY_IRR Currency = "IRR"
	CURRENCY_ISK Currency = "ISK"
	CURRENCY_JEP Currency = "JEP"
	CURRENCY_JMD Currency = "JMD"
	CURRENCY_JOD Currency = "JOD"
	CURRENCY_JPY Currency = "JPY"
	CURRENCY_KES Currency = "KES"
	CURRENCY_KGS Currency = "KGS"
	CURRENCY_KHR Currency = "KHR"
	CURRENCY_KID Currency = "KID"
	CURRENCY_KMF Currency = "KMF"
	CURRENCY_KPW Currency = "KPW"
	CURRENCY_KRW Currency = "KRW"
	CURRENCY_KWD Currency = "KWD"
	CURRENCY_KYD Currency = "KYD"
	CURRENCY_KZT Currency = "KZT"
	CURRENCY_LAK Currency = "LAK"
	CURRENCY_LBP Currency = "LBP"
	CURRENCY_LKR Currency = "LKR"
	CURRENCY_LRD Currency = "LRD"
	CURRENCY_LSL Currency = "LSL"
	CURRENCY_LYD Currency = "LYD"
	CURRENCY_MAD Currency = "MAD"
	CURRENCY_MDL Currency = "MDL"
	CURRENCY_MGA Currency = "MGA"
	CURRENCY_MKD Currency = "MKD"
	CURRENCY_MMK Currency = "MMK"
	CURRENCY_MNT Currency = "MNT"
	CURRENCY_MOP Currency = "MOP"
	CURRENCY_MRU Currency = "MRU"
	CURRENCY_MUR Currency = "MUR"
	CURRENCY_MVR Currency = "MVR"
	CURRENCY_MWK Currency = "MWK"
	CURRENCY_MXN Currency = "MXN"
	CURRENCY_MXV Currency = "MXV"
	CURRENCY_MYR Currency = "MYR"
	CURRENCY_MZN Currency = "MZN"
	CURRENCY_NAD Currency = "NAD"
	CURRENCY_NGN Currency = "NGN"
	CURRENCY_NIO Currency = "NIO"
	CURRENCY_NIS Currency = "NIS"
	CURRENCY_NOK Currency = "NOK"
	CURRENCY_NPR Currency = "NPR"
	CURRENCY_NTD Currency = "NTD"
	CURRENCY_NZD Currency = "NZD"
	CURRENCY_OMR Currency = "OMR"
	CURRENCY_PAB Currency = "PAB"
	CURRENCY_PEN Currency = "PEN"
	CURRENCY_PGK Currency = "PGK"
	CURRENCY_PHP Currency = "PHP"
	CURRENCY_PKR Currency = "PKR"
	CURRENCY_PLN Currency = "PLN"
	CURRENCY_PRB Currency = "PRB"
	CURRENCY_PYG Currency = "PYG"
	CURRENCY_QAR Currency = "QAR"
	CURRENCY_RMB Currency = "RMB"
	CURRENCY_RON Currency = "RON"
	CURRENCY_RSD Currency = "RSD"
	CURRENCY_RUB Currency = "RUB"
	CURRENCY_RWF Currency = "RWF"
	CURRENCY_SAR Currency = "SAR"
	CURRENCY_SBD Currency = "SBD"
	CURRENCY_SCR Currency = "SCR"
	CURRENCY_SDG Currency = "SDG"
	CURRENCY_SEK Currency = "SEK"
	CURRENCY_SGD Currency = "SGD"
	CURRENCY_SHP Currency = "SHP"
	CURRENCY_SLL Currency = "SLL"
	CURRENCY_SLS Currency = "SLS"
	CURRENCY_SOS Currency = "SOS"
	CURRENCY_SRD Currency = "SRD"
	CURRENCY_SSP Currency = "SSP"
	CURRENCY_STN Currency = "STN"
	CURRENCY_SVC Currency = "SVC"
	CURRENCY_SYP Currency = "SYP"
	CURRENCY_SZL Currency = "SZL"
	CURRENCY_THB Currency = "THB"
	CURRENCY_TJS Currency = "TJS"
	CURRENCY_TMT Currency = "TMT"
	CURRENCY_TND Currency = "TND"
	CURRENCY_TOP Currency = "TOP"
	CURRENCY_TRY Currency = "TRY"
	CURRENCY_TTD Currency = "TTD"
	CURRENCY_TVD Currency = "TVD"
	CURRENCY_TWD Currency = "TWD"
	CURRENCY_TZS Currency = "TZS"
	CURRENCY_UAH Currency = "UAH"
	CURRENCY_UGX Currency = "UGX"
	CURRENCY_USD Currency = "USD"
	CURRENCY_USN Currency = "USN"
	CURRENCY_UYI Currency = "UYI"
	CURRENCY_UYU Currency = "UYU"
	CURRENCY_UYW Currency = "UYW"
	CURRENCY_UZS Currency = "UZS"
	CURRENCY_VEF Currency = "VEF"
	CURRENCY_VES Currency = "VES"
	CURRENCY_VND Currency = "VND"
	CURRENCY_VUV Currency = "VUV"
	CURRENCY_WST Currency = "WST"
	CURRENCY_XAF Currency = "XAF"
	CURRENCY_XAG Currency = "XAG"
	CURRENCY_XAU Currency = "XAU"
	CURRENCY_XBA Currency = "XBA"
	CURRENCY_XBB Currency = "XBB"
	CURRENCY_XBC Currency = "XBC"
	CURRENCY_XBD Currency = "XBD"
	CURRENCY_XCD Currency = "XCD"
	CURRENCY_XDR Currency = "XDR"
	CURRENCY_XOF Currency = "XOF"
	CURRENCY_XPD Currency = "XPD"
	CURRENCY_XPF Currency = "XPF"
	CURRENCY_XPT Currency = "XPT"
	CURRENCY_XSU Currency = "XSU"
	CURRENCY_XTS Currency = "XTS"
	CURRENCY_XUA Currency = "XUA"
	CURRENCY_XXX Currency = "XXX"
	CURRENCY_YER Currency = "YER"
	CURRENCY_ZAR Currency = "ZAR"
	CURRENCY_ZMW Currency = "ZMW"
	CURRENCY_ZWB Currency = "ZWB"
	CURRENCY_ZWL Currency = "ZWL"
)

func (v *Currency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Currency(value)
	for _, existing := range []Currency{ "AED", "AFN", "ALL", "AMD", "ANG", "AOA", "ARS", "AUD", "AWG", "AZN", "BAM", "BBD", "BDT", "BGN", "BHD", "BIF", "BMD", "BND", "BOB", "BOV", "BRL", "BSD", "BTN", "BWP", "BYN", "BZD", "CAD", "CDF", "CHE", "CHF", "CHN", "CHW", "CLF", "CLP", "CNY", "COP", "COU", "CRC", "CUC", "CUP", "CVE", "CZK", "DJF", "DKK", "DOP", "DZD", "EGP", "ERN", "ETB", "EUR", "FJD", "FKP", "GBP", "GEL", "GGP", "GHS", "GIP", "GMD", "GNF", "GTQ", "GYD", "HKD", "HNL", "HRK", "HTG", "HUF", "IDR", "ILS", "IMP", "INR", "IQD", "IRR", "ISK", "JEP", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KID", "KMF", "KPW", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LYD", "MAD", "MDL", "MGA", "MKD", "MMK", "MNT", "MOP", "MRU", "MUR", "MVR", "MWK", "MXN", "MXV", "MYR", "MZN", "NAD", "NGN", "NIO", "NIS", "NOK", "NPR", "NTD", "NZD", "OMR", "PAB", "PEN", "PGK", "PHP", "PKR", "PLN", "PRB", "PYG", "QAR", "RMB", "RON", "RSD", "RUB", "RWF", "SAR", "SBD", "SCR", "SDG", "SEK", "SGD", "SHP", "SLL", "SLS", "SOS", "SRD", "SSP", "STN", "SVC", "SYP", "SZL", "THB", "TJS", "TMT", "TND", "TOP", "TRY", "TTD", "TVD", "TWD", "TZS", "UAH", "UGX", "USD", "USN", "UYI", "UYU", "UYW", "UZS", "VEF", "VES", "VND", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XDR", "XOF", "XPD", "XPF", "XPT", "XSU", "XTS", "XUA", "XXX", "YER", "ZAR", "ZMW", "ZWB", "ZWL",   } {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Currency", value)
}

// Ptr returns reference to Currency value
func (v Currency) Ptr() *Currency {
	return &v
}

type NullableCurrency struct {
	value *Currency
	isSet bool
}

func (v NullableCurrency) Get() *Currency {
	return v.value
}

func (v *NullableCurrency) Set(val *Currency) {
	v.value = val
	v.isSet = true
}

func (v NullableCurrency) IsSet() bool {
	return v.isSet
}

func (v *NullableCurrency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCurrency(val *Currency) *NullableCurrency {
	return &NullableCurrency{value: val, isSet: true}
}

func (v NullableCurrency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCurrency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

