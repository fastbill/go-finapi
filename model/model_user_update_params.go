/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: 1.106.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// User details
type UserUpdateParams struct {
	// User's new email address. Maximum length is 320. Pass an empty string (\"\") if you want to clear the current email address.
	Email string `json:"email,omitempty"`
	// User's new phone number. Maximum length is 50. Pass an empty string (\"\") if you want to clear the current phone number.
	Phone string `json:"phone,omitempty"`
	// Whether the user's bank connections will get updated in the course of finAPI's automatic batch update. Note that the automatic batch update will only update bank connections where all of the following applies:<br><br> - the PIN is stored in finAPI for the bank connection<br> - the user has accepted the latest Terms and Conditions (this only applies to users whose mandator has PISP license or doesn't have a license at all)<br> - the user has allowed finAPI to use his old stored credentials (this only applies to users which had stored credentials before introducing a web form feature and whose mandator has PISP license or doesn't have a license at all)<br> - the previous update using the stored credentials did not fail due to the credentials being incorrect (or there was no previous update with the stored credentials)<br> - the bank that the bank connection relates to is included in the automatic batch update (please contact your Sys-Admin for details about the batch update configuration)<br><br>Also note that the automatic batch update must generally be enabled for your client in order for this field to have any effect.<br/><br/>WARNING: The automatic update will always download transactions and security positions for any account that it updates! This means that the user will no longer be able to download just the balances for his accounts once the automatic update has run (The 'skipPositionsDownload' flag in the bankConnections/update service can no longer be set to true).
	IsAutoUpdateEnabled bool `json:"isAutoUpdateEnabled,omitempty"`
}
