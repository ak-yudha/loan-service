package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/ak-yudha/loan-service/models"
	"github.com/ak-yudha/loan-service/repositories"
)

var layout = "2006-01-02 15:04:05"

// LoanService interface defines the methods related to loan operations.
type LoanService interface {
	CreateLoan(borrowerID string, principalAmount, rate, roi float64, agreementLetterURL string) (*models.Loan, error)
	ApproveLoan(loanID uint, pictureProof, employeeID string, dateLoan string) (*models.Loan, error)
	InvestInLoan(loanID uint, investorID string, amount float64) (*models.Loan, error)
	DisburseLoan(loanID uint, employeeID, agreementLetterURL, disbursementDate string) (*models.Loan, error)
}

// loanServiceImpl is the concrete implementation of LoanService interface.
type loanServiceImpl struct {
	repo repositories.LoanRepository
}

// NewLoanService creates a new instance of LoanService.
func NewLoanService(repo repositories.LoanRepository) LoanService {
	return &loanServiceImpl{repo: repo}
}

func (s *loanServiceImpl) CreateLoan(borrowerID string, principalAmount, rate, roi float64, agreementLetterURL string) (*models.Loan, error) {
	loan := &models.Loan{
		BorrowerID:         borrowerID,
		PrincipalAmount:    principalAmount,
		Rate:               rate,
		ROI:                roi,
		AgreementLetterURL: agreementLetterURL,
		State:              models.Proposed,
	}
	err := s.repo.CreateLoan(loan)
	return loan, err
}

func (s *loanServiceImpl) ApproveLoan(loanID uint, pictureProof, employeeID string, dateLoan string) (*models.Loan, error) {
	loan, err := s.repo.GetLoanByID(loanID)
	if err != nil {
		return nil, err
	}
	if loan.State != models.Proposed {
		return nil, errors.New("loan can only be approved from proposed state")
	}

	date, err := time.Parse(layout, dateLoan)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, err
	}

	loan.ApprovalInfo = &models.ApprovalInfo{
		PictureProof: pictureProof,
		EmployeeID:   employeeID,
		Date:         date,
	}
	loan.State = models.Approved

	err = s.repo.UpdateLoan(loan)
	return loan, err
}

func (s *loanServiceImpl) InvestInLoan(loanID uint, investorID string, amount float64) (*models.Loan, error) {
	loan, err := s.repo.GetLoanByID(loanID)
	if err != nil {
		return nil, err
	}
	if loan.State != models.Approved {
		return nil, errors.New("loan can only be invested in when in approved state")
	}

	totalInvested := amount
	for _, inv := range loan.Investments {
		totalInvested += inv.Amount
	}
	if totalInvested > loan.PrincipalAmount {
		return nil, errors.New("total invested amount cannot exceed principal amount")
	}

	investment := models.Investment{
		LoanID:     loanID,
		InvestorID: investorID,
		Amount:     amount,
	}
	loan.Investments = append(loan.Investments, investment)

	if totalInvested == loan.PrincipalAmount {
		loan.State = models.Invested
		// TODO: Send email to investors
	}

	err = s.repo.UpdateLoan(loan)
	return loan, err
}

func (s *loanServiceImpl) DisburseLoan(loanID uint, employeeID, agreementLetterURL, disbursementDate string) (*models.Loan, error) {
	loan, err := s.repo.GetLoanByID(loanID)
	if err != nil {
		return nil, err
	}
	if loan.State != models.Invested {
		return nil, errors.New("loan can only be disbursed when in invested state")
	}

	date, err := time.Parse(layout, disbursementDate)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil, err
	}

	loan.DisbursementInfo = &models.DisbursementInfo{
		EmployeeID:         employeeID,
		AgreementLetterURL: agreementLetterURL,
		Date:               date,
	}
	loan.State = models.Disbursed

	err = s.repo.UpdateLoan(loan)
	return loan, err
}
