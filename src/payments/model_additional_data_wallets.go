/*
 * Adyen Payment API
 *
 * A set of API endpoints that allow you to initiate, settle, and modify payments on the Adyen payments platform. You can use the API to accept card payments (including One-Click and 3D Secure), bank transfers, ewallets, and many other payment methods.  To learn more about the API, visit [Classic integration](https://docs.adyen.com/classic-integration).  ## Authentication To connect to the Payments API, you must use your basic authentication credentials. For this, create your web service user, as described in [How to get the WS user password](https://docs.adyen.com/user-management/how-to-get-the-web-service-ws-user-password). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@Company.YourCompany\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Payments API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://pal-test.adyen.com/pal/servlet/Payment/v64/authorise ```
 *
 * API version: 64
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payments
// AdditionalDataWallets struct for AdditionalDataWallets
type AdditionalDataWallets struct {
	// The Android Pay token retrieved from the SDK.
	AndroidpayToken string `json:"androidpay.token,omitempty"`
	// The Mastercard Masterpass Transaction ID retrieved from the SDK.
	MasterpassTransactionId string `json:"masterpass.transactionId,omitempty"`
	// The Apple Pay token retrieved from the SDK.
	PaymentToken string `json:"payment.token,omitempty"`
	// The Google Pay token retrieved from the SDK.
	PaywithgoogleToken string `json:"paywithgoogle.token,omitempty"`
	// The Samsung Pay token retrieved from the SDK.
	SamsungpayToken string `json:"samsungpay.token,omitempty"`
	// The Visa Checkout Call ID retrieved from the SDK.
	VisacheckoutCallId string `json:"visacheckout.callId,omitempty"`
}
