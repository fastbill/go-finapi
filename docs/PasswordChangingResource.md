# PasswordChangingResource

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | User identifier | [default to null]
**UserEmail** | **string** | User&#39;s email, encrypted. Decrypt with your data decryption key. If the user has no email set, then this field will be null. | [optional] [default to null]
**PasswordChangeToken** | **string** | Encrypted password change token. Decrypt this token with your data decryption key, and pass the decrypted token to the /users/executePasswordChange service in order to set a new password for the user. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


