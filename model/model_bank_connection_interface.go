/*
 * finAPI RESTful Services
 *
 * finAPI RESTful Services
 *
 * API version: 1.106.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package finapi

// Resource representing a bank connection interface
type BankConnectionInterface struct {
	// Bank interface. Possible values:<br><br>&bull; <code>FINTS_SERVER</code> - means that finAPI will download data via the bank's FinTS interface.<br>&bull; <code>WEB_SCRAPER</code> - means that finAPI will parse data from the bank's online banking website.<br>&bull; <code>XS2A</code> - means that finAPI will download data via the bank's XS2A interface.<br>
	Interface_ string `json:"interface"`
	// Login fields for this interface, in the order that we suggest to show them to the user.
	LoginCredentials []LoginCredentialResource `json:"loginCredentials,omitempty"`
	// The default two-step-procedure for this interface. Must match one of the available 'procedureId's from the 'twoStepProcedures' list. When this field is set, then finAPI will automatically try to select the procedure wherever applicable. Note that the list of available procedures of a bank connection may change as a result of an update of the connection, and if this field references a procedure that is no longer available after an update, finAPI will automatically clear the default procedure (set it to null).
	DefaultTwoStepProcedureId string `json:"defaultTwoStepProcedureId,omitempty"`
	// Available two-step-procedures in this interface, used for submitting a money transfer or direct debit request (see /accounts/requestSepaMoneyTransfer or /requestSepaDirectDebit),or for multi-step-authentication during bank connection import or update. The available two-step-procedures mya be re-evaluated each time this bank connection is updated (/bankConnections/update). This means that this list may change as a result of an update.
	TwoStepProcedures []TwoStepProcedure `json:"twoStepProcedures,omitempty"`
	// If this field is set, it means that this interface is handing out a consent to finAPI in exchange for the login credentials. finAPI needs to use this consent to get access to the account list and account data (i.e. Account Information Services, AIS). If this field is not set, it means that this interface does not use such consents.
	AisConsent *BankConsent `json:"aisConsent,omitempty"`
	// Result of the last manual update of the associated bank connection using this interface. If no manual update has ever been done so far with this interface, then this field will not be set.
	LastManualUpdate *UpdateResult `json:"lastManualUpdate,omitempty"`
	// Result of the last auto update of the associated bank connection using this interface (ran by finAPI's automatic batch update process). If no auto update has ever been done so far with this interface, then this field will not be set.
	LastAutoUpdate *UpdateResult `json:"lastAutoUpdate,omitempty"`
}
