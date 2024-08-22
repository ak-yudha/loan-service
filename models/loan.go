package models

import "time"

type LoanState string

const (
	Proposed  LoanState = "proposed"
	Approved  LoanState = "approved"
	Invested  LoanState = "invested"
	Disbursed LoanState = "disbursed"
)

type Loan struct {
	ID                 uint              `json:"id" gorm:"primary_key"`
	BorrowerID         string            `json:"borrower_id"`
	PrincipalAmount    float64           `json:"principal_amount"`
	Rate               float64           `json:"rate"`
	ROI                float64           `json:"roi"`
	AgreementLetterURL string            `json:"agreement_letter_url"`
	State              LoanState         `json:"state"`
	ApprovalInfo       *ApprovalInfo     `json:"approval_info" gorm:"foreignkey:LoanID"`
	Investments        []Investment      `json:"investments" gorm:"foreignkey:LoanID"`
	DisbursementInfo   *DisbursementInfo `json:"disbursement_info" gorm:"foreignkey:LoanID"`
	CreatedAt          time.Time         `json:"created_at"`
	UpdatedAt          time.Time         `json:"updated_at"`
}

type ApprovalInfo struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	LoanID       uint      `json:"loan_id"`
	PictureProof string    `json:"picture_proof"`
	EmployeeID   string    `json:"employee_id"`
	Date         time.Time `json:"date"`
}

type Investment struct {
	ID         uint    `json:"id" gorm:"primary_key"`
	LoanID     uint    `json:"loan_id"`
	InvestorID string  `json:"investor_id"`
	Amount     float64 `json:"amount"`
}

type DisbursementInfo struct {
	ID                 uint      `json:"id" gorm:"primary_key"`
	LoanID             uint      `json:"loan_id"`
	EmployeeID         string    `json:"employee_id"`
	AgreementLetterURL string    `json:"agreement_letter_url"`
	Date               time.Time `json:"date"`
}
