package repositories

import (
	"github.com/ak-yudha/loan-service/models"
	"gorm.io/gorm"
)

type LoanRepository interface {
	CreateLoan(loan *models.Loan) error
	GetLoanByID(id uint) (*models.Loan, error)
	UpdateLoan(loan *models.Loan) error
}

type MySQLLoanRepository struct {
	DB *gorm.DB
}

func NewMySQLLoanRepository(db *gorm.DB) LoanRepository {
	return &MySQLLoanRepository{DB: db}
}

func (r *MySQLLoanRepository) CreateLoan(loan *models.Loan) error {
	return r.DB.Create(loan).Error
}

func (r *MySQLLoanRepository) GetLoanByID(id uint) (*models.Loan, error) {
	var loan models.Loan
	err := r.DB.Preload("ApprovalInfo").Preload("Investments").Preload("DisbursementInfo").First(&loan, id).Error
	return &loan, err
}

func (r *MySQLLoanRepository) UpdateLoan(loan *models.Loan) error {
	return r.DB.Save(loan).Error
}
