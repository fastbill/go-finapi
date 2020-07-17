# DeleteConsent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Local** | **string** | Result of deleting a consent stored in the finAPI database (local):&lt;br/&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;DELETED&lt;/code&gt; - when there was a stored consent and it was deleted.&lt;br/&gt;&amp;bull; &lt;code&gt;NOT_EXIST&lt;/code&gt; - if there is no stored consent.&lt;br/&gt; | [default to null]
**Remote** | **string** | Result of deleting a consent stored on the bank&#39;s side (remote):&lt;br/&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;DELETED&lt;/code&gt; - if the consent was successfully deleted on the bank side.&lt;br/&gt;&amp;bull; &lt;code&gt;NOT_SUPPORTED&lt;/code&gt; - if the bank doesn&#39;t support the feature of deleting consents.&lt;br/&gt;&amp;bull; &lt;code&gt;NOT_EXIST&lt;/code&gt; - if either there is no remote consent, or there is no local consent (and that makes impossible to identify any remote data).&lt;br/&gt; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


