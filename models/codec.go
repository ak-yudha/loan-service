package models

type InvestLoanRequest struct {
	InvestorID string  `json:"investor_id"`
	Amount     float64 `json:"amount"`
}

type ApproveLoanRequest struct {
	PictureProof string `json:"picture_proof"`
	EmployeeID   string `json:"employee_id"`
	Date         string `json:"date"`
}

type CreateLoanRequest struct {
	BorrowerID         string  `json:"borrower_id"`
	PrincipalAmount    float64 `json:"principal_amount"`
	Rate               float64 `json:"rate"`
	ROI                float64 `json:"roi"`
	AgreementLetterURL string  `json:"agreement_letter_url"`
}

type DisburseLoanRequest struct {
	EmployeeID         string `json:"employee_id"`
	AgreementLetterURL string `json:"agreement_letter_url"`
	Date               string `json:"date"`
}
