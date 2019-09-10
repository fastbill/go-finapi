# TppCertificateParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type_** | **string** | A type of certificate submitted | [default to null]
**PublicKey** | **string** | A certificate (public key) | [default to null]
**PrivateKey** | **string** | A private key in PKCS #8 format. PKCS #8 private keys are typically exchanged in the PEM base64-encoded format (https://support.quovadisglobal.com/kb/a37/what-is-pem-format.aspx)&lt;/br&gt;&lt;/br&gt;NOTE: The certificate should have one of the following headers:&lt;/br&gt;- &#39;-----BEGIN PRIVATE KEY-----&#39;&lt;/br&gt;- &#39;-----BEGIN ENCRYPTED PRIVATE KEY-----&#39;&lt;br&gt;Any other header denotes that the private key is NOT in PKCS #8 format! | [default to null]
**Passphrase** | **string** | Optional passphrase for the private key | [optional] [default to null]
**Label** | **string** | Optional label to certificate to identify in the list of certificates | [default to null]
**ValidFromDate** | **string** | Start day of the certificate&#39;s validity, in the format &#39;YYYY-MM-DD&#39;. Default is the passed certificate validFrom date | [optional] [default to null]
**ValidUntilDate** | **string** | Expiration day of the certificate&#39;s validity, in the format &#39;YYYY-MM-DD&#39;. Default is the passed certificate validUntil date | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


