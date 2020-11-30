/*
 * Adyen for Platforms: Account API
 *
 * The Account API provides endpoints for managing account-related entities on your platform. These related entities include account holders, accounts, bank accounts, shareholders, and KYC-related documents. The management operations include actions such as creation, retrieval, updating, and deletion of them.  For more information, refer to our [documentation](https://docs.adyen.com/platforms). ## Authentication To connect to the Account API, you must use basic authentication credentials of your web service user. If you don't have one, please contact the [Adyen Support Team](https://support.adyen.com/hc/en-us/requests/new). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@MarketPlace.YourMarketPlace\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning The Account API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://cal-test.adyen.com/cal/services/Account/v6/createAccountHolder ```
 *
 * API version: 6
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package platformsaccount
// UpdateAccountResponse struct for UpdateAccountResponse
type UpdateAccountResponse struct {
	// The code of the account.
	AccountCode string `json:"accountCode"`
	// The bankAccountUUID of the bank account held by the account holder to couple the account with. Scheduled payouts in currencies matching the currency of this bank account will be sent to this bank account. Payouts in different currencies will be sent to a matching bank account of the account holder.
	BankAccountUUID string `json:"bankAccountUUID,omitempty"`
	// The description of the account.
	Description string `json:"description,omitempty"`
	// A list of fields that caused the `/updateAccount` request to fail.
	InvalidFields *[]ErrorFieldType `json:"invalidFields,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	// The payout method code held by the account holder to couple the account with. Scheduled card payouts will be sent using this payout method code.
	PayoutMethodCode string `json:"payoutMethodCode,omitempty"`
	PayoutSchedule *PayoutScheduleResponse `json:"payoutSchedule,omitempty"`
	// Speed with which payouts for this account are processed. Permitted values: `STANDARD`, `SAME_DAY`.
	PayoutSpeed string `json:"payoutSpeed,omitempty"`
	// The reference of a request. Can be used to uniquely identify the request.
	PspReference string `json:"pspReference,omitempty"`
	// The result code.
	ResultCode string `json:"resultCode,omitempty"`
}
