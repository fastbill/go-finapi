/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: 1.106.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Bank consent information
type BankConsent struct {
	// Status of the consent. If <code>PRESENT</code>, it means that finAPI has a consent stored and can use it to connect to the bank. If <code>NOT_PRESENT</code>, finAPI will need to acquire a consent when connecting to the bank, which may require login credentials (either passed by the client, or stored in finAPI), and/or user involvement. Note that even when a consent is <code>PRESENT</code>, it may no longer be valid and finAPI will still have to acquire a new consent.
	Status string `json:"status"`
	// Expiration time of the consent in the format 'YYYY-MM-DD HH:MM:SS.SSS' (german time).
	ExpiresAt string `json:"expiresAt"`
}
