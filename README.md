# Go finAPI <!-- omit in toc -->

> This package provides a finAPI SDK for Go. It was created via [swagger-codegen](https://github.com/swagger-api/swagger-codegen) as described in [the finAPI documentation](https://finapi.zendesk.com/hc/en-us/articles/219390177-SDK-creation-for-finAPI-). Addionally some wrapper code was added to improve the usability.

- API version: v1.106.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

- [Usage](#usage)
- [Documentation for API Endpoints](#documentation-for-api-endpoints)
- [Documentation For Models](#documentation-for-models)
- [How to Update This SDK](#how-to-update-this-sdk)

## Usage
```go
import ( 
    "github.com/fastbill/go-finapi/v5"
    finapimodel "github.com/fastbill/go-finapi/v5/model"
)


config := finapi.ClientConfig{
    Endpoint: "",
    ClientID: "",
    ClientSecret: "",
}
client := finapi.NewClient(config)

// Use a client context to perform actions like creating new users.
ctx, err := client.NewClientContext()
// handle error
newUser := finapimodel.UserCreateParams{
    Email: "",
    Password: "",
    IsAutoUpdateEnabled: true,
}
user, res, err := client.UsersApi.CreateUser(ctx, newUser)

// Use a user context for user based actions like connection banks etc.
userCtx, err := client.NewUserContext("put finapi user id here", "put user password here")
// handle error
newBankConn := finapimodel.ImportBankConnectionParams{
    BankId: 1234,
    BankingUserId: "",
    BankingPin: "",
    StorePin: true,
    MaxDaysForDownload: 90,
    AccountTypeIds: []int{1,2,3}
}
bankConnection, res, err := client.BankConnectionsApi.ImportBankConnection(userCtx, newBankConn)
```

## Documentation for API Endpoints

All URIs are relative to *https://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountsApi* | [**DeleteAccount**](docs/AccountsApi.md#deleteaccount) | **Delete** /api/v1/accounts/{id} | Delete an account
*AccountsApi* | [**DeleteAllAccounts**](docs/AccountsApi.md#deleteallaccounts) | **Delete** /api/v1/accounts | Delete all accounts
*AccountsApi* | [**EditAccount**](docs/AccountsApi.md#editaccount) | **Patch** /api/v1/accounts/{id} | Edit an account
*AccountsApi* | [**ExecuteSepaDirectDebit**](docs/AccountsApi.md#executesepadirectdebit) | **Post** /api/v1/accounts/executeSepaDirectDebit | Execute SEPA Direct Debit
*AccountsApi* | [**ExecuteSepaMoneyTransfer**](docs/AccountsApi.md#executesepamoneytransfer) | **Post** /api/v1/accounts/executeSepaMoneyTransfer | Execute SEPA Money Transfer
*AccountsApi* | [**GetAccount**](docs/AccountsApi.md#getaccount) | **Get** /api/v1/accounts/{id} | Get an account
*AccountsApi* | [**GetAndSearchAllAccounts**](docs/AccountsApi.md#getandsearchallaccounts) | **Get** /api/v1/accounts | Get and search all accounts
*AccountsApi* | [**GetDailyBalances**](docs/AccountsApi.md#getdailybalances) | **Get** /api/v1/accounts/dailyBalances | Get daily balances
*AccountsApi* | [**GetMultipleAccounts**](docs/AccountsApi.md#getmultipleaccounts) | **Get** /api/v1/accounts/{ids} | Get multiple accounts
*AccountsApi* | [**RequestSepaDirectDebit**](docs/AccountsApi.md#requestsepadirectdebit) | **Post** /api/v1/accounts/requestSepaDirectDebit | Request SEPA Direct Debit
*AccountsApi* | [**RequestSepaMoneyTransfer**](docs/AccountsApi.md#requestsepamoneytransfer) | **Post** /api/v1/accounts/requestSepaMoneyTransfer | Request SEPA Money Transfer
*AuthorizationApi* | [**GetToken**](docs/AuthorizationApi.md#gettoken) | **Post** /oauth/token | Get tokens
*AuthorizationApi* | [**RevokeToken**](docs/AuthorizationApi.md#revoketoken) | **Post** /oauth/revoke | Revoke a token
*BankConnectionsApi* | [**ConnectInterface**](docs/BankConnectionsApi.md#connectinterface) | **Post** /api/v1/bankConnections/connectInterface | Connect a new interface
*BankConnectionsApi* | [**DeleteAccessData**](docs/BankConnectionsApi.md#deleteaccessdata) | **Delete** /api/v1/bankConnections/{id}/aisConsent | Delete a consent
*BankConnectionsApi* | [**DeleteAllBankConnections**](docs/BankConnectionsApi.md#deleteallbankconnections) | **Delete** /api/v1/bankConnections | Delete all bank connections
*BankConnectionsApi* | [**DeleteBankConnection**](docs/BankConnectionsApi.md#deletebankconnection) | **Delete** /api/v1/bankConnections/{id} | Delete a bank connection
*BankConnectionsApi* | [**EditBankConnection**](docs/BankConnectionsApi.md#editbankconnection) | **Patch** /api/v1/bankConnections/{id} | Edit a bank connection
*BankConnectionsApi* | [**GetAllBankConnections**](docs/BankConnectionsApi.md#getallbankconnections) | **Get** /api/v1/bankConnections | Get all bank connections
*BankConnectionsApi* | [**GetBankConnection**](docs/BankConnectionsApi.md#getbankconnection) | **Get** /api/v1/bankConnections/{id} | Get a bank connection
*BankConnectionsApi* | [**GetMultipleBankConnections**](docs/BankConnectionsApi.md#getmultiplebankconnections) | **Get** /api/v1/bankConnections/{ids} | Get multiple bank connections
*BankConnectionsApi* | [**ImportBankConnection**](docs/BankConnectionsApi.md#importbankconnection) | **Post** /api/v1/bankConnections/import | Import a new bank connection
*BankConnectionsApi* | [**RemoveInterface**](docs/BankConnectionsApi.md#removeinterface) | **Post** /api/v1/bankConnections/removeInterface | Remove an interface
*BankConnectionsApi* | [**UpdateBankConnection**](docs/BankConnectionsApi.md#updatebankconnection) | **Post** /api/v1/bankConnections/update | Update a bank connection
*BanksApi* | [**GetAndSearchAllBanks**](docs/BanksApi.md#getandsearchallbanks) | **Get** /api/v1/banks | Get and search all banks
*BanksApi* | [**GetBank**](docs/BanksApi.md#getbank) | **Get** /api/v1/banks/{id} | Get a bank
*BanksApi* | [**GetMultipleBanks**](docs/BanksApi.md#getmultiplebanks) | **Get** /api/v1/banks/{ids} | Get multiple banks
*CategoriesApi* | [**CreateCategory**](docs/CategoriesApi.md#createcategory) | **Post** /api/v1/categories | Create a new category
*CategoriesApi* | [**DeleteAllCategories**](docs/CategoriesApi.md#deleteallcategories) | **Delete** /api/v1/categories | Delete all categories
*CategoriesApi* | [**DeleteCategory**](docs/CategoriesApi.md#deletecategory) | **Delete** /api/v1/categories/{id} | Delete a category
*CategoriesApi* | [**EditCategory**](docs/CategoriesApi.md#editcategory) | **Patch** /api/v1/categories/{id} | Edit a category
*CategoriesApi* | [**GetAndSearchAllCategories**](docs/CategoriesApi.md#getandsearchallcategories) | **Get** /api/v1/categories | Get and search all categories
*CategoriesApi* | [**GetCashFlows**](docs/CategoriesApi.md#getcashflows) | **Get** /api/v1/categories/cashFlows | Get cash flows
*CategoriesApi* | [**GetCategory**](docs/CategoriesApi.md#getcategory) | **Get** /api/v1/categories/{id} | Get a category
*CategoriesApi* | [**GetMultipleCategories**](docs/CategoriesApi.md#getmultiplecategories) | **Get** /api/v1/categories/{ids} | Get multiple categories
*CategoriesApi* | [**TrainCategorization**](docs/CategoriesApi.md#traincategorization) | **Post** /api/v1/categories/trainCategorization | Train categorization
*ClientConfigurationApi* | [**EditClientConfiguration**](docs/ClientConfigurationApi.md#editclientconfiguration) | **Patch** /api/v1/clientConfiguration | Edit client configuration
*ClientConfigurationApi* | [**GetClientConfiguration**](docs/ClientConfigurationApi.md#getclientconfiguration) | **Get** /api/v1/clientConfiguration | Get client configuration
*LabelsApi* | [**CreateLabel**](docs/LabelsApi.md#createlabel) | **Post** /api/v1/labels | Create a new label
*LabelsApi* | [**DeleteAllLabels**](docs/LabelsApi.md#deletealllabels) | **Delete** /api/v1/labels | Delete all labels
*LabelsApi* | [**DeleteLabel**](docs/LabelsApi.md#deletelabel) | **Delete** /api/v1/labels/{id} | Delete a label
*LabelsApi* | [**EditLabel**](docs/LabelsApi.md#editlabel) | **Patch** /api/v1/labels/{id} | Edit a label
*LabelsApi* | [**GetAndSearchAllLabels**](docs/LabelsApi.md#getandsearchalllabels) | **Get** /api/v1/labels | Get and search all labels
*LabelsApi* | [**GetLabel**](docs/LabelsApi.md#getlabel) | **Get** /api/v1/labels/{id} | Get a label
*LabelsApi* | [**GetMultipleLabels**](docs/LabelsApi.md#getmultiplelabels) | **Get** /api/v1/labels/{ids} | Get multiple labels
*MandatorAdministrationApi* | [**ChangeClientCredentials**](docs/MandatorAdministrationApi.md#changeclientcredentials) | **Post** /api/v1/mandatorAdmin/changeClientCredentials | Change client credentials
*MandatorAdministrationApi* | [**CreateIbanRules**](docs/MandatorAdministrationApi.md#createibanrules) | **Post** /api/v1/mandatorAdmin/ibanRules | Create IBAN rules
*MandatorAdministrationApi* | [**CreateKeywordRules**](docs/MandatorAdministrationApi.md#createkeywordrules) | **Post** /api/v1/mandatorAdmin/keywordRules | Create keyword rules
*MandatorAdministrationApi* | [**DeleteIbanRules**](docs/MandatorAdministrationApi.md#deleteibanrules) | **Post** /api/v1/mandatorAdmin/ibanRules/delete | Delete IBAN rules
*MandatorAdministrationApi* | [**DeleteKeywordRules**](docs/MandatorAdministrationApi.md#deletekeywordrules) | **Post** /api/v1/mandatorAdmin/keywordRules/delete | Delete keyword rules
*MandatorAdministrationApi* | [**DeleteUsers**](docs/MandatorAdministrationApi.md#deleteusers) | **Post** /api/v1/mandatorAdmin/deleteUsers | Delete users
*MandatorAdministrationApi* | [**GetIbanRuleList**](docs/MandatorAdministrationApi.md#getibanrulelist) | **Get** /api/v1/mandatorAdmin/ibanRules | Get IBAN rules
*MandatorAdministrationApi* | [**GetKeywordRuleList**](docs/MandatorAdministrationApi.md#getkeywordrulelist) | **Get** /api/v1/mandatorAdmin/keywordRules | Get keyword rules
*MandatorAdministrationApi* | [**GetUserList**](docs/MandatorAdministrationApi.md#getuserlist) | **Get** /api/v1/mandatorAdmin/getUserList | Get user list
*MocksAndTestsApi* | [**CheckCategorization**](docs/MocksAndTestsApi.md#checkcategorization) | **Post** /api/v1/tests/checkCategorization | Check categorization
*MocksAndTestsApi* | [**MockBatchUpdate**](docs/MocksAndTestsApi.md#mockbatchupdate) | **Post** /api/v1/tests/mockBatchUpdate | Mock batch update
*NotificationRulesApi* | [**CreateNotificationRule**](docs/NotificationRulesApi.md#createnotificationrule) | **Post** /api/v1/notificationRules | Create a new notification rule
*NotificationRulesApi* | [**DeleteAllNotificationRules**](docs/NotificationRulesApi.md#deleteallnotificationrules) | **Delete** /api/v1/notificationRules | Delete all notification rules
*NotificationRulesApi* | [**DeleteNotificationRule**](docs/NotificationRulesApi.md#deletenotificationrule) | **Delete** /api/v1/notificationRules/{id} | Delete a notification rule
*NotificationRulesApi* | [**GetAndSearchAllNotificationRules**](docs/NotificationRulesApi.md#getandsearchallnotificationrules) | **Get** /api/v1/notificationRules | Get and search all notification rules
*NotificationRulesApi* | [**GetNotificationRule**](docs/NotificationRulesApi.md#getnotificationrule) | **Get** /api/v1/notificationRules/{id} | Get a notification rule
*PaymentsApi* | [**CreateDirectDebit**](docs/PaymentsApi.md#createdirectdebit) | **Post** /api/v1/payments/directDebits | Create direct debit
*PaymentsApi* | [**CreateMoneyTransfer**](docs/PaymentsApi.md#createmoneytransfer) | **Post** /api/v1/payments/moneyTransfers | Create money transfer
*PaymentsApi* | [**GetPayments**](docs/PaymentsApi.md#getpayments) | **Get** /api/v1/payments | Get payments
*PaymentsApi* | [**SubmitPayment**](docs/PaymentsApi.md#submitpayment) | **Post** /api/v1/payments/submit | Submit payment
*SecuritiesApi* | [**GetAndSearchAllSecurities**](docs/SecuritiesApi.md#getandsearchallsecurities) | **Get** /api/v1/securities | Get and search all securities
*SecuritiesApi* | [**GetMultipleSecurities**](docs/SecuritiesApi.md#getmultiplesecurities) | **Get** /api/v1/securities/{ids} | Get multiple securities
*SecuritiesApi* | [**GetSecurity**](docs/SecuritiesApi.md#getsecurity) | **Get** /api/v1/securities/{id} | Get a security
*TPPCertificatesApi* | [**CreateNewCertificate**](docs/TPPCertificatesApi.md#createnewcertificate) | **Post** /api/v1/tppCertificates | Create a new certificate
*TPPCertificatesApi* | [**DeleteCertificate**](docs/TPPCertificatesApi.md#deletecertificate) | **Delete** /api/v1/tppCertificates/{id} | Delete a certificate
*TPPCertificatesApi* | [**GetAllCertificates**](docs/TPPCertificatesApi.md#getallcertificates) | **Get** /api/v1/tppCertificates | Get all certificates
*TPPCertificatesApi* | [**GetCertificate**](docs/TPPCertificatesApi.md#getcertificate) | **Get** /api/v1/tppCertificates/{id} | Get a certificate
*TPPCredentialsApi* | [**CreateTppCredential**](docs/TPPCredentialsApi.md#createtppcredential) | **Post** /api/v1/tppCredentials | Create new TPP credentials
*TPPCredentialsApi* | [**DeleteTppCredential**](docs/TPPCredentialsApi.md#deletetppcredential) | **Delete** /api/v1/tppCredentials/{id} | Delete a set of TPP credentials
*TPPCredentialsApi* | [**EditTppCredential**](docs/TPPCredentialsApi.md#edittppcredential) | **Patch** /api/v1/tppCredentials/{id} | Edit a set of TPP credentials
*TPPCredentialsApi* | [**GetAllTppCredentials**](docs/TPPCredentialsApi.md#getalltppcredentials) | **Get** /api/v1/tppCredentials | Get all TPP credentials
*TPPCredentialsApi* | [**GetAndSearchTppAuthenticationGroups**](docs/TPPCredentialsApi.md#getandsearchtppauthenticationgroups) | **Get** /api/v1/tppCredentials/tppAuthenticationGroups | Get all TPP Authentication Groups
*TPPCredentialsApi* | [**GetTppCredential**](docs/TPPCredentialsApi.md#gettppcredential) | **Get** /api/v1/tppCredentials/{id} | Get a set of TPP credentials
*TransactionsApi* | [**DeleteAllTransactions**](docs/TransactionsApi.md#deletealltransactions) | **Delete** /api/v1/transactions | Delete all transactions
*TransactionsApi* | [**DeleteTransaction**](docs/TransactionsApi.md#deletetransaction) | **Delete** /api/v1/transactions/{id} | Delete a transaction
*TransactionsApi* | [**EditMultipleTransactions**](docs/TransactionsApi.md#editmultipletransactions) | **Patch** /api/v1/transactions | Edit multiple transactions
*TransactionsApi* | [**EditMultipleTransactionsDeprecated**](docs/TransactionsApi.md#editmultipletransactionsdeprecated) | **Patch** /api/v1/transactions/{ids} | Edit multiple transactions (DEPRECATED)
*TransactionsApi* | [**EditTransaction**](docs/TransactionsApi.md#edittransaction) | **Patch** /api/v1/transactions/{id} | Edit a transaction
*TransactionsApi* | [**GetAndSearchAllTransactions**](docs/TransactionsApi.md#getandsearchalltransactions) | **Get** /api/v1/transactions | Get and search all transactions
*TransactionsApi* | [**GetMultipleTransactions**](docs/TransactionsApi.md#getmultipletransactions) | **Get** /api/v1/transactions/{ids} | Get multiple transactions
*TransactionsApi* | [**GetTransaction**](docs/TransactionsApi.md#gettransaction) | **Get** /api/v1/transactions/{id} | Get a transaction
*TransactionsApi* | [**RestoreTransaction**](docs/TransactionsApi.md#restoretransaction) | **Post** /api/v1/transactions/{id}/restore | Restore a transaction
*TransactionsApi* | [**SplitTransaction**](docs/TransactionsApi.md#splittransaction) | **Post** /api/v1/transactions/{id}/split | Split a transaction
*TransactionsApi* | [**TriggerCategorization**](docs/TransactionsApi.md#triggercategorization) | **Post** /api/v1/transactions/triggerCategorization | Trigger categorization
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /api/v1/users | Create a new user
*UsersApi* | [**DeleteAuthorizedUser**](docs/UsersApi.md#deleteauthorizeduser) | **Delete** /api/v1/users | Delete the authorized user
*UsersApi* | [**DeleteUnverifiedUser**](docs/UsersApi.md#deleteunverifieduser) | **Delete** /api/v1/users/{userId} | Delete an unverified user
*UsersApi* | [**EditAuthorizedUser**](docs/UsersApi.md#editauthorizeduser) | **Patch** /api/v1/users | Edit the authorized user
*UsersApi* | [**ExecutePasswordChange**](docs/UsersApi.md#executepasswordchange) | **Post** /api/v1/users/executePasswordChange | Execute password change
*UsersApi* | [**GetAuthorizedUser**](docs/UsersApi.md#getauthorizeduser) | **Get** /api/v1/users | Get the authorized user
*UsersApi* | [**GetVerificationStatus**](docs/UsersApi.md#getverificationstatus) | **Get** /api/v1/users/verificationStatus | Get a user&#39;s verification status
*UsersApi* | [**RequestPasswordChange**](docs/UsersApi.md#requestpasswordchange) | **Post** /api/v1/users/requestPasswordChange | Request password change
*UsersApi* | [**VerifyUser**](docs/UsersApi.md#verifyuser) | **Post** /api/v1/users/verify/{userId} | Verify a user
*WebFormsApi* | [**GetWebForm**](docs/WebFormsApi.md#getwebform) | **Get** /api/v1/webForms/{id} | Get a web form


## Documentation For Models

 - [AccessToken](docs/AccessToken.md)
 - [Account](docs/Account.md)
 - [AccountInterface](docs/AccountInterface.md)
 - [AccountList](docs/AccountList.md)
 - [AccountParams](docs/AccountParams.md)
 - [AccountReference](docs/AccountReference.md)
 - [BadCredentialsError](docs/BadCredentialsError.md)
 - [Bank](docs/Bank.md)
 - [BankConnection](docs/BankConnection.md)
 - [BankConnectionInterface](docs/BankConnectionInterface.md)
 - [BankConnectionList](docs/BankConnectionList.md)
 - [BankConnectionOwner](docs/BankConnectionOwner.md)
 - [BankConsent](docs/BankConsent.md)
 - [BankGroup](docs/BankGroup.md)
 - [BankInterface](docs/BankInterface.md)
 - [BankInterfaceLoginField](docs/BankInterfaceLoginField.md)
 - [BankList](docs/BankList.md)
 - [CashFlow](docs/CashFlow.md)
 - [CashFlowList](docs/CashFlowList.md)
 - [CategorizationCheckResult](docs/CategorizationCheckResult.md)
 - [CategorizationCheckResults](docs/CategorizationCheckResults.md)
 - [Category](docs/Category.md)
 - [CategoryList](docs/CategoryList.md)
 - [CategoryParams](docs/CategoryParams.md)
 - [ChangeClientCredentialsParams](docs/ChangeClientCredentialsParams.md)
 - [CheckCategorizationData](docs/CheckCategorizationData.md)
 - [CheckCategorizationTransactionData](docs/CheckCategorizationTransactionData.md)
 - [ClearingAccountData](docs/ClearingAccountData.md)
 - [ClientConfiguration](docs/ClientConfiguration.md)
 - [ClientConfigurationParams](docs/ClientConfigurationParams.md)
 - [ConnectInterfaceParams](docs/ConnectInterfaceParams.md)
 - [CreateDirectDebitParams](docs/CreateDirectDebitParams.md)
 - [CreateMoneyTransferParams](docs/CreateMoneyTransferParams.md)
 - [DailyBalance](docs/DailyBalance.md)
 - [DailyBalanceList](docs/DailyBalanceList.md)
 - [DeleteConsent](docs/DeleteConsent.md)
 - [DirectDebitOrderParams](docs/DirectDebitOrderParams.md)
 - [DirectDebitOrderingResponse](docs/DirectDebitOrderingResponse.md)
 - [EditBankConnectionParams](docs/EditBankConnectionParams.md)
 - [EditCategoryParams](docs/EditCategoryParams.md)
 - [EditTppCredentialParams](docs/EditTppCredentialParams.md)
 - [ErrorDetails](docs/ErrorDetails.md)
 - [ErrorMessage](docs/ErrorMessage.md)
 - [ExecutePasswordChangeParams](docs/ExecutePasswordChangeParams.md)
 - [ExecuteSepaDirectDebitParams](docs/ExecuteSepaDirectDebitParams.md)
 - [ExecuteSepaMoneyTransferParams](docs/ExecuteSepaMoneyTransferParams.md)
 - [IbanRule](docs/IbanRule.md)
 - [IbanRuleIdentifiersParams](docs/IbanRuleIdentifiersParams.md)
 - [IbanRuleList](docs/IbanRuleList.md)
 - [IbanRuleParams](docs/IbanRuleParams.md)
 - [IbanRulesParams](docs/IbanRulesParams.md)
 - [IdentifierList](docs/IdentifierList.md)
 - [ImportBankConnectionParams](docs/ImportBankConnectionParams.md)
 - [KeywordRule](docs/KeywordRule.md)
 - [KeywordRuleIdentifiersParams](docs/KeywordRuleIdentifiersParams.md)
 - [KeywordRuleList](docs/KeywordRuleList.md)
 - [KeywordRuleParams](docs/KeywordRuleParams.md)
 - [KeywordRulesParams](docs/KeywordRulesParams.md)
 - [Label](docs/Label.md)
 - [LabelList](docs/LabelList.md)
 - [LabelParams](docs/LabelParams.md)
 - [LoginCredential](docs/LoginCredential.md)
 - [LoginCredentialResource](docs/LoginCredentialResource.md)
 - [MockAccountData](docs/MockAccountData.md)
 - [MockBankConnectionUpdate](docs/MockBankConnectionUpdate.md)
 - [MockBatchUpdateParams](docs/MockBatchUpdateParams.md)
 - [MoneyTransferOrderParams](docs/MoneyTransferOrderParams.md)
 - [MoneyTransferOrderingResponse](docs/MoneyTransferOrderingResponse.md)
 - [MonthlyUserStats](docs/MonthlyUserStats.md)
 - [MultiStepAuthenticationCallback](docs/MultiStepAuthenticationCallback.md)
 - [MultiStepAuthenticationChallenge](docs/MultiStepAuthenticationChallenge.md)
 - [NewTransaction](docs/NewTransaction.md)
 - [NotificationRule](docs/NotificationRule.md)
 - [NotificationRuleList](docs/NotificationRuleList.md)
 - [NotificationRuleParams](docs/NotificationRuleParams.md)
 - [PageableBankList](docs/PageableBankList.md)
 - [PageableCategoryList](docs/PageableCategoryList.md)
 - [PageableIbanRuleList](docs/PageableIbanRuleList.md)
 - [PageableKeywordRuleList](docs/PageableKeywordRuleList.md)
 - [PageableLabelList](docs/PageableLabelList.md)
 - [PageablePaymentResources](docs/PageablePaymentResources.md)
 - [PageableSecurityList](docs/PageableSecurityList.md)
 - [PageableTppAuthenticationGroupResources](docs/PageableTppAuthenticationGroupResources.md)
 - [PageableTppCertificateList](docs/PageableTppCertificateList.md)
 - [PageableTppCredentialResources](docs/PageableTppCredentialResources.md)
 - [PageableTransactionList](docs/PageableTransactionList.md)
 - [PageableUserInfoList](docs/PageableUserInfoList.md)
 - [Paging](docs/Paging.md)
 - [PasswordChangingResource](docs/PasswordChangingResource.md)
 - [Payment](docs/Payment.md)
 - [PaymentExecutionResponse](docs/PaymentExecutionResponse.md)
 - [PaypalTransactionData](docs/PaypalTransactionData.md)
 - [RemoveInterfaceParams](docs/RemoveInterfaceParams.md)
 - [RequestPasswordChangeParams](docs/RequestPasswordChangeParams.md)
 - [RequestSepaDirectDebitParams](docs/RequestSepaDirectDebitParams.md)
 - [RequestSepaMoneyTransferParams](docs/RequestSepaMoneyTransferParams.md)
 - [Security](docs/Security.md)
 - [SecurityList](docs/SecurityList.md)
 - [SingleDirectDebitData](docs/SingleDirectDebitData.md)
 - [SingleMoneyTransferRecipientData](docs/SingleMoneyTransferRecipientData.md)
 - [SplitTransactionsParams](docs/SplitTransactionsParams.md)
 - [SubTransactionParams](docs/SubTransactionParams.md)
 - [SubmitPaymentParams](docs/SubmitPaymentParams.md)
 - [TppAuthenticationGroup](docs/TppAuthenticationGroup.md)
 - [TppCertificate](docs/TppCertificate.md)
 - [TppCertificateParams](docs/TppCertificateParams.md)
 - [TppCredentials](docs/TppCredentials.md)
 - [TppCredentialsParams](docs/TppCredentialsParams.md)
 - [TrainCategorizationData](docs/TrainCategorizationData.md)
 - [TrainCategorizationTransactionData](docs/TrainCategorizationTransactionData.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionList](docs/TransactionList.md)
 - [TriggerCategorizationParams](docs/TriggerCategorizationParams.md)
 - [TwoStepProcedure](docs/TwoStepProcedure.md)
 - [UpdateBankConnectionParams](docs/UpdateBankConnectionParams.md)
 - [UpdateMultipleTransactionsParams](docs/UpdateMultipleTransactionsParams.md)
 - [UpdateResult](docs/UpdateResult.md)
 - [UpdateTransactionsParams](docs/UpdateTransactionsParams.md)
 - [User](docs/User.md)
 - [UserCreateParams](docs/UserCreateParams.md)
 - [UserIdentifiersList](docs/UserIdentifiersList.md)
 - [UserIdentifiersParams](docs/UserIdentifiersParams.md)
 - [UserInfo](docs/UserInfo.md)
 - [UserUpdateParams](docs/UserUpdateParams.md)
 - [VerificationStatusResource](docs/VerificationStatusResource.md)
 - [WebForm](docs/WebForm.md)

## How to Update This SDK
1. Download the newest version of the official SDK from [docs.finapi.io](https://docs.finapi.io). For that click the download button on the top of the page and select "go" as language.
2. Save the zip file you get there to some place on your PC and extract it. Open the folder `go-client`.
3. Replace the folders `docs` and `api` in the current FastBill finAPI SDK repository with the new ones that you downloaded.
4. Delete everything in the current `model` folder and then copy all files starting with `model_` from the freshly downloaded SDK to the repositories `model` folder.
5. Copy all files besides the ones with `model_`, `README.md` and `git_push.sh` and use them to replace the files on root level in the current repository.
6. Rename all occurances of `package swagger` to `package finapi` in the repository.
7. Add the import statement `. "github.com/fastbill/go-finapi/v5/model"` to all the files in the root folder of the repository besides `configuration.go`, `finapi.go`, `client.go` and `response.go`. When you save the file your IDE should also automatically fix all Go formatting errors in the file. Also some missing import statements for the package `optional` might be added in this process.
8. In the file `client.go` make sure to keep the following change for the serialization of the Swagger error (either revert the change from copying the file or just paste in the old version again):
    ```go
    func (e GenericSwaggerError) Error() string {
        return e.error + ", body: " + string(e.body)
    }
    ```
9. In the file `client.go` make sure to keep the custom HTTP client with the timeout, not the DefaultClient that you get from the fresh SDK download.
    ```go
    func NewAPIClient(cfg *Configuration) *APIClient {
        if cfg.HTTPClient == nil {
            cfg.HTTPClient = &http.Client{
                // The AWS ALB will terminate the connection (e.g. to legacy) after 60s so there is no point
                // in waiting longer than that for finAPI. Even when they would respond later we could not send
                // the response anymore.
                Timeout: 55 * time.Second,
            }
        }
        // ...
    ```
10. From the `README.md` file in the downloaded SDK copy the sections `Documentation for API Endpoints` and `Documentation For Models` and use them to replace those sections in the repository `README.md` file. Also update the API version mentioned at the top of the readme.
11. Replace all `float32` with `float64` and `Float32` with `Float64` in the entire SDK. This is needed because otherwise we get rounding errors for large transaction amounts. Float32 only garantees 6 correct digits, not enought for amounts like `123456.78`.
12. Review the changes in the Git diff before commiting them. Files in `docs` folder can be commited without thourough checking since they only contain text changes.
