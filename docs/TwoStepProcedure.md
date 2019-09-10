# TwoStepProcedure

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProcedureId** | **string** | Bank-given ID of the procedure | [default to null]
**ProcedureName** | **string** | Bank-given name of the procedure | [default to null]
**ProcedureChallengeType** | **string** | The challenge type of the procedure. Possible values are:&lt;br/&gt;&lt;br/&gt;&amp;bull; &lt;code&gt;TEXT&lt;/code&gt; - the challenge will be a text that contains instructions for the user on how to proceed with the authorization.&lt;br/&gt;&amp;bull; &lt;code&gt;PHOTO&lt;/code&gt; - the challenge will contain a BASE-64 string depicting a photo (or any kind of QR-code-like data) that must be shown to the user.&lt;br/&gt;&amp;bull; &lt;code&gt;FLICKER_CODE&lt;/code&gt; - the challenge will contain a BASE-64 string depicting a flicker code animation that must be shown to the user.&lt;br/&gt;&lt;br/&gt;Note that this challenge type information does not originate from the bank, but is determined by finAPI internally. There is no guarantee that the determined challenge type is correct. Note also that this field may not be set, meaning that finAPI could not determine the challenge type of the procedure. | [optional] [default to null]
**ImplicitExecute** | **bool** | If &#39;true&#39;, then there is no need for your client to pass back anything to finAPI to continue the authorization when using this procedure. The authorization will be dealt with directly between the user, finAPI, and the bank. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


