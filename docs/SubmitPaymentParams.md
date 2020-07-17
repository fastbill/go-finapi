# SubmitPaymentParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **int64** | Payment identifier | [default to null]
**Interface_** | **string** | Bank interface. Possible values:&lt;br&gt;&lt;br&gt;&amp;bull; &lt;code&gt;FINTS_SERVER&lt;/code&gt; - means that finAPI will execute the payment via the bank&#39;s FinTS interface.&lt;br&gt;&amp;bull; &lt;code&gt;WEB_SCRAPER&lt;/code&gt; - means that finAPI will parse data from the bank&#39;s online banking website.&lt;br&gt;&amp;bull; &lt;code&gt;XS2A&lt;/code&gt; - means that finAPI will execute the payment via the bank&#39;s XS2A interface.Please note that XS2A doesn&#39;t support direct debits yet. &lt;br/&gt;To determine what interface(s) you can choose to submit a payment, please refer to the field AccountInterface.capabilities of the account that is related to the payment, or if this is a standalone payment without a related account imported in finAPI, refer to the field BankInterface.isMoneyTransferSupported.&lt;br/&gt;For standalone money transfers in particular, we suggest to always use XS2A if supported, and only use FINTS_SERVER or WEB_SCRAPER as a fallback, because non-XS2A interfaces might require not just a single, but multiple authentications when submitting the payment.&lt;br/&gt; | [default to null]
**LoginCredentials** | [**[]LoginCredential**](LoginCredential.md) | Login credentials. May not be required when the credentials are stored in finAPI, or when the bank interface has no login credentials. | [optional] [default to null]
**RedirectUrl** | **string** | Must only be passed when the used interface has the property REDIRECT_APPROACH and no web form flow is used. The user will be redirected to the given URL from the bank&#39;s website after having entered his credentials. | [optional] [default to null]
**MultiStepAuthentication** | [***MultiStepAuthenticationCallback**](MultiStepAuthenticationCallback.md) | Container for multi-step authentication data. Required when a previous service call initiated a multi-step authentication. | [optional] [default to null]
**HideTransactionDetailsInWebForm** | **bool** | Whether the finAPI web form should hide transaction details when prompting the caller for the second factor. Default value is false. | [optional] [default to null]
**ForceWebForm** | **bool** | If the user has stored his credentials and a default two step procedure, the finAPI web form will only be shown in case a TAN has to be provided or an error has occurred. In case the user wants to change his stored data, you can set this field to true. It will force the webform flow for the user and allow changes to the data he has stored. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


