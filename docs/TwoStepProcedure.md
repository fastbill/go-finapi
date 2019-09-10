# TwoStepProcedure

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcedureId** | **string** | Bank-given ID of the procedure | [default to null]
**ProcedureName** | **string** | Bank-given name of the procedure | [default to null]
**ProcedureChallengeType** | **string** | The challenge type of the procedure. Possible values are:&lt;br/&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;TEXT&lt;/code&gt; - the challenge will be a text that contains instructions for the user on how to retrieve the TAN.&lt;br/&gt;&amp;bull; &lt;code&gt;PHOTO&lt;/code&gt; - the challenge will contain a BASE-64 string depicting a photo (or any kind of QR-code-like data) that must be shown to the user.&lt;br/&gt;&amp;bull; &lt;code&gt;FLICKER_CODE&lt;/code&gt; - the challenge will contain a BASE-64 string depicting a flicker code animation that must be shown to the user.&lt;br/&gt;&lt;br/&gt;Note that this challenge type information does not originate from the bank server, but is determined by finAPI internally. There is no guarantee that the determined challenge type is correct. Note also that this field may not be set, meaning that finAPI could not determine the challenge type of the procedure.&lt;br/&gt;&lt;br/&gt;For further information on how to deal with the challenges, please see &lt;a href&#x3D;&#39;https://finapi.zendesk.com/hc/en-us/articles/219117247-SEPA-Money-Transfer&#39;&gt;this article&lt;/a&gt; on our Dev Portal. | [optional] [default to null]
**ImplicitExecute** | **bool** | If &#39;true&#39;, then requesting a SEPA order with this procedure will implicitly trigger the execution of the order. For example, if you do a money transfer with this procedure, then calling the /requestSepaMoneyTransfer service will immediately execute the order (a call to /executeSepaMoneyTransfer will not be necessary). On the other hand, if this flag is &#39;false&#39;, then doing a money transfer with this procedure will require you to first call /requestSepaMoneyTransfer, and then call /executeSepaMoneyTransfer. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


