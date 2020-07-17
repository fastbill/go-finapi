/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: 1.106.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Password changing details
type PasswordChangingResource struct {
	// User identifier
	UserId string `json:"userId"`
	// User's email, encrypted. Decrypt with your data decryption key. If the user has no email set, then this field will be null.
	UserEmail string `json:"userEmail,omitempty"`
	// Encrypted password change token. Decrypt this token with your data decryption key, and pass the decrypted token to the /users/executePasswordChange service in order to set a new password for the user.
	PasswordChangeToken string `json:"passwordChangeToken"`
}
