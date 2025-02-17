/*
 * Adyen Checkout API
 *
 * Adyen Checkout API provides a simple and flexible way to initiate and authorise online payments. You can use the same integration for payments made with cards (including 3D Secure), mobile wallets, and local payment methods (for example, iDEAL and Sofort).  This API reference provides information on available endpoints and how to interact with them. To learn more about the API, visit [Checkout documentation](https://docs.adyen.com/online-payments).  ## Authentication Each request to the Checkout API must be signed with an API key. For this, obtain an API Key from your Customer Area, as described in [How to get the API key](https://docs.adyen.com/development-resources/api-credentials#generate-api-key). Then set this key to the `X-API-Key` header value, for example:  ``` curl -H \"Content-Type: application/json\" \\ -H \"X-API-Key: Your_Checkout_API_key\" \\ ... ``` Note that when going live, you need to generate a new API Key to access the [live endpoints](https://docs.adyen.com/development-resources/live-endpoints).  ## Versioning Checkout API supports versioning of its endpoints through a version suffix in the endpoint URL. This suffix has the following format: \"vXX\", where XX is the version number.  For example: ``` https://checkout-test.adyen.com/v67/payments ```
 *
 * API version: 67
 * Contact: developer-experience@adyen.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package checkout

import (
	"encoding/json"
	"time"
)

// PaymentRequest struct for PaymentRequest
type PaymentRequest struct {
	AccountInfo *AccountInfo `json:"accountInfo,omitempty"`
	// This field contains additional data, which may be required for a particular payment request.  The `additionalData` object consists of entries, each of which includes the key and value.
	AdditionalData  map[string]string `json:"additionalData,omitempty"`
	Amount          Amount            `json:"amount"`
	ApplicationInfo *ApplicationInfo  `json:"applicationInfo,omitempty"`
	BillingAddress  *Address          `json:"billingAddress,omitempty"`
	BrowserInfo     *BrowserInfo      `json:"browserInfo,omitempty"`
	// The delay between the authorisation and scheduled auto-capture, specified in hours.
	CaptureDelayHours int32 `json:"captureDelayHours,omitempty"`
	// The platform where a payment transaction takes place. This field is optional for filtering out payment methods that are only available on specific platforms. If this value is not set, then we will try to infer it from the `sdkVersion` or `token`.  Possible values: * iOS * Android * Web
	Channel string   `json:"channel,omitempty"`
	Company *Company `json:"company,omitempty"`
	// Conversion ID that corresponds to the Id generated for tracking user payment journey.
	ConversionId string `json:"conversionId,omitempty"`
	// The shopper country.  Format: [ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) Example: NL or DE
	CountryCode string `json:"countryCode,omitempty"`
	// The shopper's date of birth.  Format [ISO-8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DD
	DateOfBirth     *time.Time  `json:"dateOfBirth,omitempty"`
	DccQuote        *ForexQuote `json:"dccQuote,omitempty"`
	DeliveryAddress *Address    `json:"deliveryAddress,omitempty"`
	// The date and time the purchased goods should be delivered.  Format [ISO 8601](https://www.w3.org/TR/NOTE-datetime): YYYY-MM-DDThh:mm:ss.sssTZD  Example: 2017-07-17T13:42:40.428+01:00
	DeliveryDate *time.Time `json:"deliveryDate,omitempty"`
	// A string containing the shopper's device fingerprint. For more information, refer to [Device fingerprinting](https://docs.adyen.com/risk-management/device-fingerprinting).
	DeviceFingerprint string `json:"deviceFingerprint,omitempty"`
	// When true and `shopperReference` is provided, the shopper will be asked if the payment details should be stored for future one-click payments.
	EnableOneClick bool `json:"enableOneClick,omitempty"`
	// When true and `shopperReference` is provided, the payment details will be tokenized for payouts.
	EnablePayOut bool `json:"enablePayOut,omitempty"`
	// When true and `shopperReference` is provided, the payment details will be tokenized for recurring payments.
	EnableRecurring bool `json:"enableRecurring,omitempty"`
	// The type of the entity the payment is processed for.
	EntityType string `json:"entityType,omitempty"`
	// An integer value that is added to the normal fraud score. The value can be either positive or negative.
	FraudOffset  int32         `json:"fraudOffset,omitempty"`
	Installments *Installments `json:"installments,omitempty"`
	// Price and product information about the purchased items, to be included on the invoice sent to the shopper. > This field is required for 3x 4x Oney, Affirm, AfterPay, Klarna, RatePay, and Zip.
	LineItems *[]LineItem `json:"lineItems,omitempty"`
    // Mandate for India products
    Mandate *SubscriptionDetails `json:"mandate,omitempty"`
    // The [merchant category code](https://en.wikipedia.org/wiki/Merchant_category_code) (MCC) is a four-digit number, which relates to a particular market segment. This code reflects the predominant activity that is conducted by the merchant.
	Mcc string `json:"mcc,omitempty"`
	// The merchant account identifier, with which you want to process the transaction.
	MerchantAccount string `json:"merchantAccount"`
	// This reference allows linking multiple transactions to each other for reporting purposes (i.e. order auth-rate). The reference should be unique per billing cycle. The same merchant order reference should never be reused after the first authorised attempt. If used, this field should be supplied for all incoming authorisations. > We strongly recommend you send the `merchantOrderReference` value to benefit from linking payment requests when authorisation retries take place. In addition, we recommend you provide `retry.orderAttemptNumber`, `retry.chainAttemptNumber`, and `retry.skipRetry` values in `PaymentRequest.additionalData`.
	MerchantOrderReference string                 `json:"merchantOrderReference,omitempty"`
	MerchantRiskIndicator  *MerchantRiskIndicator `json:"merchantRiskIndicator,omitempty"`
	// Metadata consists of entries, each of which includes a key and a value. Limitations: Maximum 20 key-value pairs per request. When exceeding, the \"177\" error occurs: \"Metadata size exceeds limit\".
	Metadata map[string]string `json:"metadata,omitempty"`
	MpiData  *ThreeDSecureData `json:"mpiData,omitempty"`
	Order    *CheckoutOrder    `json:"order,omitempty"`
	// When you are doing multiple partial (gift card) payments, this is the `pspReference` of the first payment. We use this to link the multiple payments to each other. As your own reference for linking multiple payments, use the `merchantOrderReference`instead.
	OrderReference string `json:"orderReference,omitempty"`
	// Required for the 3D Secure 2 `channel` **Web** integration.  Set this parameter to the origin URL of the page that you are loading the 3D Secure Component from.
	Origin string `json:"origin,omitempty"`
	// The type and required details of a payment method to use.
	PaymentMethod interface{} `json:"paymentMethod"`
	// Date after which no further authorisations shall be performed. Only for 3D Secure 2.
	RecurringExpiry string `json:"recurringExpiry,omitempty"`
	// Minimum number of days between authorisations. Only for 3D Secure 2.
	RecurringFrequency string `json:"recurringFrequency,omitempty"`
	// Defines a recurring payment type. Allowed values: * `Subscription` – A transaction for a fixed or variable amount, which follows a fixed schedule. * `CardOnFile` – With a card-on-file (CoF) transaction, card details are stored to enable one-click or omnichannel journeys, or simply to streamline the checkout process. Any subscription not following a fixed schedule is also considered a card-on-file transaction. * `UnscheduledCardOnFile` – An unscheduled card-on-file (UCoF) transaction is a transaction that occurs on a non-fixed schedule and/or have variable amounts. For example, automatic top-ups when a cardholder's balance drops below a certain amount.
	RecurringProcessingModel string `json:"recurringProcessingModel,omitempty"`
	// Specifies the redirect method (GET or POST) when redirecting back from the issuer.
	RedirectFromIssuerMethod string `json:"redirectFromIssuerMethod,omitempty"`
	// Specifies the redirect method (GET or POST) when redirecting to the issuer.
	RedirectToIssuerMethod string `json:"redirectToIssuerMethod,omitempty"`
	// The reference to uniquely identify a payment. This reference is used in all communication with you about the payment status. We recommend using a unique value per payment; however, it is not a requirement. If you need to provide multiple references for a transaction, separate them with hyphens (\"-\"). Maximum length: 80 characters.
	Reference string `json:"reference"`
	// The URL to return to in case of a redirection. The format depends on the channel. This URL can have a maximum of 1024 characters. * For web, include the protocol `http://` or `https://`. You can also include your own additional query parameters, for example, shopper ID or order reference number. Example: `https://your-company.com/checkout?shopperOrder=12xy` * For iOS, use the custom URL for your app. To know more about setting custom URL schemes, refer to the [Apple Developer documentation](https://developer.apple.com/documentation/uikit/inter-process_communication/allowing_apps_and_websites_to_link_to_your_content/defining_a_custom_url_scheme_for_your_app). Example: `my-app://` * For Android, use a custom URL handled by an Activity on your app. You can configure it with an [intent filter](https://developer.android.com/guide/components/intents-filters). Example: `my-app://your.package.name`
	ReturnUrl string    `json:"returnUrl"`
	RiskData  *RiskData `json:"riskData,omitempty"`
	// The date and time until when the session remains valid, in [ISO 8601](https://www.w3.org/TR/NOTE-datetime) format.  For example: 2020-07-18T15:42:40.428+01:00
	SessionValidity string `json:"sessionValidity,omitempty"`
	// The shopper's email address. We recommend that you provide this data, as it is used in velocity fraud checks. > For 3D Secure 2 transactions, schemes require `shopperEmail` for all browser-based and mobile implementations.
	ShopperEmail string `json:"shopperEmail,omitempty"`
	// The shopper's IP address. In general, we recommend that you provide this data, as it is used in a number of risk checks (for instance, number of payment attempts or location-based checks). > For 3D Secure 2 transactions, schemes require `shopperIP` for all browser-based implementations. This field is also mandatory for some merchants depending on your business model. For more information, [contact Support](https://support.adyen.com/hc/en-us/requests/new).
	ShopperIP string `json:"shopperIP,omitempty"`
	// Specifies the sales channel, through which the shopper gives their card details, and whether the shopper is a returning customer. For the web service API, Adyen assumes Ecommerce shopper interaction by default.  This field has the following possible values: * `Ecommerce` - Online transactions where the cardholder is present (online). For better authorisation rates, we recommend sending the card security code (CSC) along with the request. * `ContAuth` - Card on file and/or subscription transactions, where the cardholder is known to the merchant (returning customer). If the shopper is present (online), you can supply also the CSC to improve authorisation (one-click payment). * `Moto` - Mail-order and telephone-order transactions where the shopper is in contact with the merchant via email or telephone. * `POS` - Point-of-sale transactions where the shopper is physically present to make a payment using a secure payment terminal.
	ShopperInteraction string `json:"shopperInteraction,omitempty"`
	// The combination of a language code and a country code to specify the language to be used in the payment.
	ShopperLocale string `json:"shopperLocale,omitempty"`
	ShopperName   *Name  `json:"shopperName,omitempty"`
	// Your reference to uniquely identify this shopper (for example, user ID or account ID). Minimum length: 3 characters. > This field is required for recurring payments.
	ShopperReference string `json:"shopperReference,omitempty"`
	// The text to be shown on the shopper's bank statement. To enable this field, contact our [Support Team](https://support.adyen.com/hc/en-us/requests/new).  We recommend sending a maximum of 25 characters, otherwise banks might truncate the string.
	ShopperStatement string `json:"shopperStatement,omitempty"`
	// The shopper's social security number.
	SocialSecurityNumber string `json:"socialSecurityNumber,omitempty"`
	// An array of objects specifying how the payment should be split when using [Adyen for Platforms](https://docs.adyen.com/platforms/processing-payments#providing-split-information) or [Issuing](https://docs.adyen.com/issuing/manage-funds#split).
	Splits *[]Split `json:"splits,omitempty"`
	// The physical store, for which this payment is processed.
	Store string `json:"store,omitempty"`
	// When true and `shopperReference` is provided, the payment details will be stored.
	StorePaymentMethod bool `json:"storePaymentMethod,omitempty"`
	// The shopper's telephone number.
	TelephoneNumber     string               `json:"telephoneNumber,omitempty"`
	ThreeDS2RequestData *ThreeDS2RequestData `json:"threeDS2RequestData,omitempty"`
	// If set to true, you will only perform the [3D Secure 2 authentication](https://docs.adyen.com/online-payments/3d-secure/other-3ds-flows/authentication-only), and not the payment authorisation.
	ThreeDSAuthenticationOnly bool `json:"threeDSAuthenticationOnly,omitempty"`
	// Set to true if the payment should be routed to a trusted MID.
	TrustedShopper bool `json:"trustedShopper,omitempty"`
}

// UnmarshalJSON unmarshals a quoted json string to PaymentRequest struct
func (req *PaymentRequest) UnmarshalJSON(b []byte) error {

	type PaymentRequestAlias PaymentRequest

	temp := &struct {
		*PaymentRequestAlias
	}{
		PaymentRequestAlias: (*PaymentRequestAlias)(req),
	}
	if err := json.Unmarshal(b, &temp); err != nil {
		return err
	}
	if temp.PaymentMethod != nil {

		pmtype := temp.PaymentMethod.(map[string]interface{})["type"].(string)

		switch pmtype {
		case "scheme":
			intermediate := &struct {
				PaymentMethod *CardDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod
		case "ach":
			intermediate := &struct {
				PaymentMethod *AchDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "amazonpay":
			intermediate := &struct {
				PaymentMethod *AmazonPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "androidpay":
			intermediate := &struct {
				PaymentMethod *AndroidPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "applepay":
			intermediate := &struct {
				PaymentMethod *ApplePayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "directdebit_GB":
			intermediate := &struct {
				PaymentMethod *BacsDirectDebitDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "billdesk_online":
			intermediate := &struct {
				PaymentMethod *BillDeskOnlineDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "billdesk_wallet":
			intermediate := &struct {
				PaymentMethod *BillDeskWalletDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "blik":
			intermediate := &struct {
				PaymentMethod *BlikDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "cellulant":
			intermediate := &struct {
				PaymentMethod *CellulantDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "doku_mandiri_va", "doku_cimb_va", "doku_danamon_va", "doku_bni_va", "doku_permata_lite_atm", "doku_permata_atm", "doku_bri_va", "doku_bca_va", "doku_alfamart", "doku_indomaret", "doku_sinarmas_va":
			intermediate := &struct {
				PaymentMethod *DokuDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "dotpay":
			intermediate := &struct {
				PaymentMethod *DotpayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "dragonpay_ebanking", "dragonpay_otc_banking", "dragonpay_otc_non_banking", "dragonpay_otc_philippines":
			intermediate := &struct {
				PaymentMethod *DragonpayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "econtext_seveneleven", "econtext_stores":
			intermediate := &struct {
				PaymentMethod *EcontextVoucherDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "entercash":
			intermediate := &struct {
				PaymentMethod *EntercashDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "giropay":
			intermediate := &struct {
				PaymentMethod *GiropayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "paywithgoogle":
			intermediate := &struct {
				PaymentMethod *GooglePayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "ideal":
			intermediate := &struct {
				PaymentMethod *IdealDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "klarna", "klarnapayments", "klarnapayments_account", "klarnapayments_b2b", "klarna_paynow", "klarna_account", "klarna_b2b":
			intermediate := &struct {
				PaymentMethod *KlarnaDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "lianlianpay_ebanking_enterprise", "lianlianpay_ebanking_credit", "lianlianpay_ebanking_debit":
			intermediate := &struct {
				PaymentMethod *LianLianPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "masterpass":
			intermediate := &struct {
				PaymentMethod *MasterpassDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "mbway":
			intermediate := &struct {
				PaymentMethod *MbwayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "molpay_ebanking_fpx_MY", "molpay_ebanking_TH", "molpay_ebanking_VN", "molpay_ebanking_MY", "molpay_ebanking_direct_MY", "molpay_fpx":
			intermediate := &struct {
				PaymentMethod *MolPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "mobilepay":
			intermediate := &struct {
				PaymentMethod *MobilePayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "paypal":
			intermediate := &struct {
				PaymentMethod *PayPalDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "payu_IN_upi":
			intermediate := &struct {
				PaymentMethod *PayUUpiDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "qiwiwallet":
			intermediate := &struct {
				PaymentMethod *QiwiWalletDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "samsungpay":
			intermediate := &struct {
				PaymentMethod *SamsungPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "sepadirectdebit":
			intermediate := &struct {
				PaymentMethod *SepaDirectDebitDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "upi":
			intermediate := &struct {
				PaymentMethod *UpiDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "vipps":
			intermediate := &struct {
				PaymentMethod *VippsDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "visacheckout":
			intermediate := &struct {
				PaymentMethod *VisaCheckoutDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "wechatpay":
			intermediate := &struct {
				PaymentMethod *WeChatPayDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		case "wechatpayMiniProgram":
			intermediate := &struct {
				PaymentMethod *WeChatPayMiniProgramDetails `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod

		default:
			intermediate := &struct {
				PaymentMethod map[string]interface{} `json:"paymentMethod"`
				*PaymentRequestAlias
			}{
				PaymentRequestAlias: (*PaymentRequestAlias)(req),
			}

			if err := json.Unmarshal(b, &intermediate); err != nil {
				return err
			}
			req.PaymentMethod = intermediate.PaymentMethod
		}
	}

	return nil
}
