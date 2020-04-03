/*
 * Adyen Payment API
 *
 * A set of API endpoints that allow you to initiate, settle, and modify payments on the Adyen payments platform. You can use the API to accept card payments (including One-Click and 3D Secure), bank transfers, ewallets, and many other payment methods.  To learn more about the API, visit [Classic integration](https://docs.adyen.com/classic-integration).  ## Authentication To connect to the Payments API, you must use your basic authentication credentials. For this, create your web service user, as described in [How to get the WS user password](https://docs.adyen.com/user-management/how-to-get-the-web-service-ws-user-password). Then use its credentials to authenticate your request, for example:  ``` curl -U \"ws@Company.YourCompany\":\"YourWsPassword\" \\ -H \"Content-Type: application/json\" \\ ... ``` Note that when going live, you need to generate new web service user credentials to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Payments API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://pal-test.adyen.com/pal/servlet/Payment/v52/authorise ```
 *
 * API version: 52
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package payment
import (
	"time"
)
// ForexQuote struct for ForexQuote
type ForexQuote struct {
	// The account name.
	Account string `json:"account,omitempty"`
	// The account type.
	AccountType string `json:"accountType,omitempty"`
	BaseAmount *Amount `json:"baseAmount,omitempty"`
	// The base points.
	BasePoints int32 `json:"basePoints"`
	Buy *Amount `json:"buy,omitempty"`
	Interbank *Amount `json:"interbank,omitempty"`
	// The reference assigned to the forex quote request.
	Reference string `json:"reference,omitempty"`
	Sell *Amount `json:"sell,omitempty"`
	// The signature to validate the integrity.
	Signature string `json:"signature,omitempty"`
	// The source of the forex quote.
	Source string `json:"source,omitempty"`
	// The type of forex.
	Type string `json:"type,omitempty"`
	// The date until which the forex quote is valid.
	ValidTill time.Time `json:"validTill"`
}
