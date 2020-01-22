/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.93.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// User's verification status
type VerificationStatusResource struct {
	// User's identifier
	UserId string `json:"userId"`
	// User's verification status
	IsUserVerified bool `json:"isUserVerified"`
}
