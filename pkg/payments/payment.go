package payments

type Payment struct {
	PaymentType    string            `json:"type"`
	ID             string            `json:"id"`
	Version        uint              `json:"version"`
	OrganisationID string            `json:"organisation_id"`
	Attributes     PaymentAttributes `json:"attributes"`
}

type PaymentAttributes struct {
	Amount               string             `json:"amount"`
	BeneficiaryParty     PaymentParty       `json:"beneficiary_party"`
	ChargesInformation   ChargesInformation `json:"charges_information"`
	Currency             string             `json:"currency"`
	DebtorParty          PaymentParty       `json:"debtor_party"`
	EndToEndReference    string             `json:"end_to_end_reference"`
	FX                   FX                 `json:"fx"`
	NumericReference     string             `json:"numeric_reference"`
	PaymentID            string             `json:"payment_id"`
	PaymentPurpose       string             `json:"payment_purpose"`
	PaymentScheme        string             `json:"payment_scheme"`
	PaymentType          string             `json:"payment_type"`
	ProcessingDate       string             `json:"processing_date"`
	Reference            string             `json:"reference"`
	SchemePaymentSubType string             `json:"scheme_payment_sub_type"`
	SchemePaymentType    string             `json:"scheme_payment_type"`
	SponsorParty         PaymentParty       `json:"sponsor_party"`
}

type PaymentParty struct {
	AccountName       string `json:"account_name"`
	AccountNumber     string `json:"account_number"`
	AccountNumberCode string `json:"account_number_code"`
	AccountType       int    `json:"account_type"`
	Address           string `json:"address"`
	BankID            string `json:"bank_id"`
	BankIDCode        string `json:"bank_id_code"`
	Name              string `json:"name"`
}

type ChargesInformation struct {
	BearerCode              string    `json:"bearer_code"`
	SenderCharges           []Charges `json:"sender_charges"`
	ReceiverChargesAmount   string    `json:"receiver_charges_amount"`
	ReceiverChargesCurrency string    `json:"receiver_charges_currency"`
}

type FX struct {
	ContractReference string `json:"contract_reference"`
	ExchangeRate      string `json:"exchange_rate"`
	OriginalAmount    string `json:"original_amount"`
	OriginalCurrency  string `json:"original_currency"`
}

type Charges struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
