/*
 * Adyen BinLookup API
 *
 * The BIN Lookup API provides endpoints for retrieving information, such as cost estimates, and 3D Secure supported version based on a given BIN.
 *
 * API version: 50
 * Contact: support@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package binlookup
// CostEstimateRequest struct for CostEstimateRequest
type CostEstimateRequest struct {
	Amount Amount `json:"amount"`
	Assumptions *CostEstimateAssumptions `json:"assumptions,omitempty"`
	// The card number (4-19 characters) for PCI compliant use cases. Do not use any separators.  > Either the `cardNumber` or `encryptedCardNumber` field must be provided in a payment request.
	CardNumber string `json:"cardNumber,omitempty"`
	// Encrypted data that stores card information for non PCI-compliant use cases. The encrypted data must be created with the Checkout Card Component or Secured Fields Component, and must contain the `encryptedCardNumber` field.  > Either the `cardNumber` or `encryptedCardNumber` field must be provided in a payment request.
	EncryptedCardNumber string `json:"encryptedCardNumber,omitempty"`
	// The merchant account identifier you want to process the (transaction) request with.
	MerchantAccount string `json:"merchantAccount"`
	MerchantDetails *MerchantDetails `json:"merchantDetails,omitempty"`
	Recurring *Recurring `json:"recurring,omitempty"`
	// The `recurringDetailReference` you want to use for this cost estimate. The value `LATEST` can be used to select the most recently stored recurring detail.
	SelectedRecurringDetailReference string `json:"selectedRecurringDetailReference,omitempty"`
	// Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * `Ecommerce` - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * `ContAuth` - Card on file and/or subscription transactions, where the card holder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * `Moto` - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * `POS` - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal.
	ShopperInteraction string `json:"shopperInteraction,omitempty"`
	// Your reference to uniquely identify this shopper (for example, user ID or account ID). Minimum length: 3 characters. > This field is required for recurring payments.
	ShopperReference string `json:"shopperReference,omitempty"`
}
