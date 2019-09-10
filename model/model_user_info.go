/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: v1.79.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Container for user information
type UserInfo struct {
	// User's identifier
	UserId string `json:"userId"`
	// User's registration date, in the format 'YYYY-MM-DD'
	RegistrationDate string `json:"registrationDate"`
	// User's deletion date, in the format 'YYYY-MM-DD'. May be null if the user has not been deleted.
	DeletionDate string `json:"deletionDate,omitempty"`
	// User's last active date, in the format 'YYYY-MM-DD'. May be null if the user has not yet logged in.
	LastActiveDate string `json:"lastActiveDate,omitempty"`
	// Number of bank connections that currently exist for this user.
	BankConnectionCount int32 `json:"bankConnectionCount"`
	// Latest date of when a bank connection was imported for this user, in the format 'YYYY-MM-DD'. This field is null when there has never been a bank connection import.
	LatestBankConnectionImportDate string `json:"latestBankConnectionImportDate,omitempty"`
	// Latest date of when a bank connection was deleted for this user, in the format 'YYYY-MM-DD'. This field is null when there has never been a bank connection deletion.
	LatestBankConnectionDeletionDate string `json:"latestBankConnectionDeletionDate,omitempty"`
	// Additional information about the user's data or activities, broken down in months. The list will by default contain an entry for each month starting with the month of when the user was registered, up to the current month. The date range may vary when you have limited it in the request. <br/><br/>Please note:<br/>&bull; this field is only set when 'includeMonthlyStats' = true, otherwise it will be null.<br/>&bull; the list is always ordered from the latest month first, to the oldest month last.<br/>&bull; the list will never contain an entry for a month that was prior to the month of when the user was registered, or after the month of when the user was deleted, even when you have explicitly set a respective date range. This means that the list may be empty if you are requesting a date range where the user didn't exist yet, or didn't exist any longer.
	MonthlyStats []MonthlyUserStats `json:"monthlyStats,omitempty"`
	// Whether the user is currently locked (for further information, see the 'maxUserLoginAttempts' setting in your client's configuration). Note that deleted users will always have this field set to 'false'.
	IsLocked bool `json:"isLocked,omitempty"`
}
