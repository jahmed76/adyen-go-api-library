/*
 * Adyen for Platforms: Notifications
 *
 * The Notification API sends notifications to the endpoints specified in a given subscription. Subscriptions are managed through the Notification Configuration API. The API specifications listed here detail the format of each notification.  For more information, refer to our [documentation](https://docs.adyen.com/platforms/notifications).
 *
 * API version: 6
 * Contact: support@adyen.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package platformsnotificationevents

import (
	"encoding/json"
)

// AccountHolderPayoutNotificationContent struct for AccountHolderPayoutNotificationContent
type AccountHolderPayoutNotificationContent struct {
	// The code of the account from which the payout was made.
	AccountCode string `json:"accountCode"`
	// The code of the Account Holder to which the payout was made.
	AccountHolderCode string `json:"accountHolderCode"`
	// The payout amounts (per currency).
	Amounts *[]Amount `json:"amounts,omitempty"`
	BankAccountDetail *BankAccountDetail `json:"bankAccountDetail,omitempty"`
	// A description of the payout.
	Description *string `json:"description,omitempty"`
	EstimatedArrivalDate *LocalDate `json:"estimatedArrivalDate,omitempty"`
	// Invalid fields list.
	InvalidFields *[]ErrorFieldType `json:"invalidFields,omitempty"`
	// The merchant reference.
	MerchantReference *string `json:"merchantReference,omitempty"`
	// The PSP reference of the original payout.
	OriginalPspReference *string `json:"originalPspReference,omitempty"`
	// Payout account country.
	PayoutAccountCountry *string `json:"payoutAccountCountry,omitempty"`
	// Payout bank account number.
	PayoutAccountNumber *string `json:"payoutAccountNumber,omitempty"`
	// Payout bank name.
	PayoutBankName *string `json:"payoutBankName,omitempty"`
	// Payout branch code.
	PayoutBranchCode *string `json:"payoutBranchCode,omitempty"`
	// Payout transaction id.
	PayoutReference *int64 `json:"payoutReference,omitempty"`
	// Speed with which payouts for this account are processed. Permitted values: `STANDARD`, `SAME_DAY`.
	PayoutSpeed *string `json:"payoutSpeed,omitempty"`
	Status *OperationStatus `json:"status,omitempty"`
}

// NewAccountHolderPayoutNotificationContent instantiates a new AccountHolderPayoutNotificationContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountHolderPayoutNotificationContent(accountCode string, accountHolderCode string, ) *AccountHolderPayoutNotificationContent {
	this := AccountHolderPayoutNotificationContent{}
	this.AccountCode = accountCode
	this.AccountHolderCode = accountHolderCode
	return &this
}

// NewAccountHolderPayoutNotificationContentWithDefaults instantiates a new AccountHolderPayoutNotificationContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountHolderPayoutNotificationContentWithDefaults() *AccountHolderPayoutNotificationContent {
	this := AccountHolderPayoutNotificationContent{}
	return &this
}

// GetAccountCode returns the AccountCode field value
func (o *AccountHolderPayoutNotificationContent) GetAccountCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountCode
}

// GetAccountCodeOk returns a tuple with the AccountCode field value
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetAccountCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountCode, true
}

// SetAccountCode sets field value
func (o *AccountHolderPayoutNotificationContent) SetAccountCode(v string) {
	o.AccountCode = v
}

// GetAccountHolderCode returns the AccountHolderCode field value
func (o *AccountHolderPayoutNotificationContent) GetAccountHolderCode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.AccountHolderCode
}

// GetAccountHolderCodeOk returns a tuple with the AccountHolderCode field value
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetAccountHolderCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountHolderCode, true
}

// SetAccountHolderCode sets field value
func (o *AccountHolderPayoutNotificationContent) SetAccountHolderCode(v string) {
	o.AccountHolderCode = v
}

// GetAmounts returns the Amounts field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetAmounts() []Amount {
	if o == nil || o.Amounts == nil {
		var ret []Amount
		return ret
	}
	return *o.Amounts
}

// GetAmountsOk returns a tuple with the Amounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetAmountsOk() (*[]Amount, bool) {
	if o == nil || o.Amounts == nil {
		return nil, false
	}
	return o.Amounts, true
}

// HasAmounts returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasAmounts() bool {
	if o != nil && o.Amounts != nil {
		return true
	}

	return false
}

// SetAmounts gets a reference to the given []Amount and assigns it to the Amounts field.
func (o *AccountHolderPayoutNotificationContent) SetAmounts(v []Amount) {
	o.Amounts = &v
}

// GetBankAccountDetail returns the BankAccountDetail field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetBankAccountDetail() BankAccountDetail {
	if o == nil || o.BankAccountDetail == nil {
		var ret BankAccountDetail
		return ret
	}
	return *o.BankAccountDetail
}

// GetBankAccountDetailOk returns a tuple with the BankAccountDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetBankAccountDetailOk() (*BankAccountDetail, bool) {
	if o == nil || o.BankAccountDetail == nil {
		return nil, false
	}
	return o.BankAccountDetail, true
}

// HasBankAccountDetail returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasBankAccountDetail() bool {
	if o != nil && o.BankAccountDetail != nil {
		return true
	}

	return false
}

// SetBankAccountDetail gets a reference to the given BankAccountDetail and assigns it to the BankAccountDetail field.
func (o *AccountHolderPayoutNotificationContent) SetBankAccountDetail(v BankAccountDetail) {
	o.BankAccountDetail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountHolderPayoutNotificationContent) SetDescription(v string) {
	o.Description = &v
}

// GetEstimatedArrivalDate returns the EstimatedArrivalDate field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetEstimatedArrivalDate() LocalDate {
	if o == nil || o.EstimatedArrivalDate == nil {
		var ret LocalDate
		return ret
	}
	return *o.EstimatedArrivalDate
}

// GetEstimatedArrivalDateOk returns a tuple with the EstimatedArrivalDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetEstimatedArrivalDateOk() (*LocalDate, bool) {
	if o == nil || o.EstimatedArrivalDate == nil {
		return nil, false
	}
	return o.EstimatedArrivalDate, true
}

// HasEstimatedArrivalDate returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasEstimatedArrivalDate() bool {
	if o != nil && o.EstimatedArrivalDate != nil {
		return true
	}

	return false
}

// SetEstimatedArrivalDate gets a reference to the given LocalDate and assigns it to the EstimatedArrivalDate field.
func (o *AccountHolderPayoutNotificationContent) SetEstimatedArrivalDate(v LocalDate) {
	o.EstimatedArrivalDate = &v
}

// GetInvalidFields returns the InvalidFields field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetInvalidFields() []ErrorFieldType {
	if o == nil || o.InvalidFields == nil {
		var ret []ErrorFieldType
		return ret
	}
	return *o.InvalidFields
}

// GetInvalidFieldsOk returns a tuple with the InvalidFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetInvalidFieldsOk() (*[]ErrorFieldType, bool) {
	if o == nil || o.InvalidFields == nil {
		return nil, false
	}
	return o.InvalidFields, true
}

// HasInvalidFields returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasInvalidFields() bool {
	if o != nil && o.InvalidFields != nil {
		return true
	}

	return false
}

// SetInvalidFields gets a reference to the given []ErrorFieldType and assigns it to the InvalidFields field.
func (o *AccountHolderPayoutNotificationContent) SetInvalidFields(v []ErrorFieldType) {
	o.InvalidFields = &v
}

// GetMerchantReference returns the MerchantReference field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetMerchantReference() string {
	if o == nil || o.MerchantReference == nil {
		var ret string
		return ret
	}
	return *o.MerchantReference
}

// GetMerchantReferenceOk returns a tuple with the MerchantReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetMerchantReferenceOk() (*string, bool) {
	if o == nil || o.MerchantReference == nil {
		return nil, false
	}
	return o.MerchantReference, true
}

// HasMerchantReference returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasMerchantReference() bool {
	if o != nil && o.MerchantReference != nil {
		return true
	}

	return false
}

// SetMerchantReference gets a reference to the given string and assigns it to the MerchantReference field.
func (o *AccountHolderPayoutNotificationContent) SetMerchantReference(v string) {
	o.MerchantReference = &v
}

// GetOriginalPspReference returns the OriginalPspReference field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetOriginalPspReference() string {
	if o == nil || o.OriginalPspReference == nil {
		var ret string
		return ret
	}
	return *o.OriginalPspReference
}

// GetOriginalPspReferenceOk returns a tuple with the OriginalPspReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetOriginalPspReferenceOk() (*string, bool) {
	if o == nil || o.OriginalPspReference == nil {
		return nil, false
	}
	return o.OriginalPspReference, true
}

// HasOriginalPspReference returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasOriginalPspReference() bool {
	if o != nil && o.OriginalPspReference != nil {
		return true
	}

	return false
}

// SetOriginalPspReference gets a reference to the given string and assigns it to the OriginalPspReference field.
func (o *AccountHolderPayoutNotificationContent) SetOriginalPspReference(v string) {
	o.OriginalPspReference = &v
}

// GetPayoutAccountCountry returns the PayoutAccountCountry field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetPayoutAccountCountry() string {
	if o == nil || o.PayoutAccountCountry == nil {
		var ret string
		return ret
	}
	return *o.PayoutAccountCountry
}

// GetPayoutAccountCountryOk returns a tuple with the PayoutAccountCountry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetPayoutAccountCountryOk() (*string, bool) {
	if o == nil || o.PayoutAccountCountry == nil {
		return nil, false
	}
	return o.PayoutAccountCountry, true
}

// HasPayoutAccountCountry returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasPayoutAccountCountry() bool {
	if o != nil && o.PayoutAccountCountry != nil {
		return true
	}

	return false
}

// SetPayoutAccountCountry gets a reference to the given string and assigns it to the PayoutAccountCountry field.
func (o *AccountHolderPayoutNotificationContent) SetPayoutAccountCountry(v string) {
	o.PayoutAccountCountry = &v
}

// GetPayoutAccountNumber returns the PayoutAccountNumber field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetPayoutAccountNumber() string {
	if o == nil || o.PayoutAccountNumber == nil {
		var ret string
		return ret
	}
	return *o.PayoutAccountNumber
}

// GetPayoutAccountNumberOk returns a tuple with the PayoutAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetPayoutAccountNumberOk() (*string, bool) {
	if o == nil || o.PayoutAccountNumber == nil {
		return nil, false
	}
	return o.PayoutAccountNumber, true
}

// HasPayoutAccountNumber returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasPayoutAccountNumber() bool {
	if o != nil && o.PayoutAccountNumber != nil {
		return true
	}

	return false
}

// SetPayoutAccountNumber gets a reference to the given string and assigns it to the PayoutAccountNumber field.
func (o *AccountHolderPayoutNotificationContent) SetPayoutAccountNumber(v string) {
	o.PayoutAccountNumber = &v
}

// GetPayoutBankName returns the PayoutBankName field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetPayoutBankName() string {
	if o == nil || o.PayoutBankName == nil {
		var ret string
		return ret
	}
	return *o.PayoutBankName
}

// GetPayoutBankNameOk returns a tuple with the PayoutBankName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetPayoutBankNameOk() (*string, bool) {
	if o == nil || o.PayoutBankName == nil {
		return nil, false
	}
	return o.PayoutBankName, true
}

// HasPayoutBankName returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasPayoutBankName() bool {
	if o != nil && o.PayoutBankName != nil {
		return true
	}

	return false
}

// SetPayoutBankName gets a reference to the given string and assigns it to the PayoutBankName field.
func (o *AccountHolderPayoutNotificationContent) SetPayoutBankName(v string) {
	o.PayoutBankName = &v
}

// GetPayoutBranchCode returns the PayoutBranchCode field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetPayoutBranchCode() string {
	if o == nil || o.PayoutBranchCode == nil {
		var ret string
		return ret
	}
	return *o.PayoutBranchCode
}

// GetPayoutBranchCodeOk returns a tuple with the PayoutBranchCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetPayoutBranchCodeOk() (*string, bool) {
	if o == nil || o.PayoutBranchCode == nil {
		return nil, false
	}
	return o.PayoutBranchCode, true
}

// HasPayoutBranchCode returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasPayoutBranchCode() bool {
	if o != nil && o.PayoutBranchCode != nil {
		return true
	}

	return false
}

// SetPayoutBranchCode gets a reference to the given string and assigns it to the PayoutBranchCode field.
func (o *AccountHolderPayoutNotificationContent) SetPayoutBranchCode(v string) {
	o.PayoutBranchCode = &v
}

// GetPayoutReference returns the PayoutReference field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetPayoutReference() int64 {
	if o == nil || o.PayoutReference == nil {
		var ret int64
		return ret
	}
	return *o.PayoutReference
}

// GetPayoutReferenceOk returns a tuple with the PayoutReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetPayoutReferenceOk() (*int64, bool) {
	if o == nil || o.PayoutReference == nil {
		return nil, false
	}
	return o.PayoutReference, true
}

// HasPayoutReference returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasPayoutReference() bool {
	if o != nil && o.PayoutReference != nil {
		return true
	}

	return false
}

// SetPayoutReference gets a reference to the given int64 and assigns it to the PayoutReference field.
func (o *AccountHolderPayoutNotificationContent) SetPayoutReference(v int64) {
	o.PayoutReference = &v
}

// GetPayoutSpeed returns the PayoutSpeed field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetPayoutSpeed() string {
	if o == nil || o.PayoutSpeed == nil {
		var ret string
		return ret
	}
	return *o.PayoutSpeed
}

// GetPayoutSpeedOk returns a tuple with the PayoutSpeed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetPayoutSpeedOk() (*string, bool) {
	if o == nil || o.PayoutSpeed == nil {
		return nil, false
	}
	return o.PayoutSpeed, true
}

// HasPayoutSpeed returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasPayoutSpeed() bool {
	if o != nil && o.PayoutSpeed != nil {
		return true
	}

	return false
}

// SetPayoutSpeed gets a reference to the given string and assigns it to the PayoutSpeed field.
func (o *AccountHolderPayoutNotificationContent) SetPayoutSpeed(v string) {
	o.PayoutSpeed = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AccountHolderPayoutNotificationContent) GetStatus() OperationStatus {
	if o == nil || o.Status == nil {
		var ret OperationStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountHolderPayoutNotificationContent) GetStatusOk() (*OperationStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AccountHolderPayoutNotificationContent) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given OperationStatus and assigns it to the Status field.
func (o *AccountHolderPayoutNotificationContent) SetStatus(v OperationStatus) {
	o.Status = &v
}

func (o AccountHolderPayoutNotificationContent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountCode"] = o.AccountCode
	}
	if true {
		toSerialize["accountHolderCode"] = o.AccountHolderCode
	}
	if o.Amounts != nil {
		toSerialize["amounts"] = o.Amounts
	}
	if o.BankAccountDetail != nil {
		toSerialize["bankAccountDetail"] = o.BankAccountDetail
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.EstimatedArrivalDate != nil {
		toSerialize["estimatedArrivalDate"] = o.EstimatedArrivalDate
	}
	if o.InvalidFields != nil {
		toSerialize["invalidFields"] = o.InvalidFields
	}
	if o.MerchantReference != nil {
		toSerialize["merchantReference"] = o.MerchantReference
	}
	if o.OriginalPspReference != nil {
		toSerialize["originalPspReference"] = o.OriginalPspReference
	}
	if o.PayoutAccountCountry != nil {
		toSerialize["payoutAccountCountry"] = o.PayoutAccountCountry
	}
	if o.PayoutAccountNumber != nil {
		toSerialize["payoutAccountNumber"] = o.PayoutAccountNumber
	}
	if o.PayoutBankName != nil {
		toSerialize["payoutBankName"] = o.PayoutBankName
	}
	if o.PayoutBranchCode != nil {
		toSerialize["payoutBranchCode"] = o.PayoutBranchCode
	}
	if o.PayoutReference != nil {
		toSerialize["payoutReference"] = o.PayoutReference
	}
	if o.PayoutSpeed != nil {
		toSerialize["payoutSpeed"] = o.PayoutSpeed
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableAccountHolderPayoutNotificationContent struct {
	value *AccountHolderPayoutNotificationContent
	isSet bool
}

func (v NullableAccountHolderPayoutNotificationContent) Get() *AccountHolderPayoutNotificationContent {
	return v.value
}

func (v *NullableAccountHolderPayoutNotificationContent) Set(val *AccountHolderPayoutNotificationContent) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountHolderPayoutNotificationContent) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountHolderPayoutNotificationContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountHolderPayoutNotificationContent(val *AccountHolderPayoutNotificationContent) *NullableAccountHolderPayoutNotificationContent {
	return &NullableAccountHolderPayoutNotificationContent{value: val, isSet: true}
}

func (v NullableAccountHolderPayoutNotificationContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountHolderPayoutNotificationContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


