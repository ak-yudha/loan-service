package controllers

import (
	"github.com/ak-yudha/loan-service/models"
	"github.com/ak-yudha/loan-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LoanController struct {
	Service services.LoanService
}

func NewLoanController(service services.LoanService) *LoanController {
	return &LoanController{Service: service}
}

func (c *LoanController) CreateLoan(ctx *gin.Context) {
	var req models.CreateLoanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loan, err := c.Service.CreateLoan(req.BorrowerID, req.PrincipalAmount, req.Rate, req.ROI, req.AgreementLetterURL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) ApproveLoan(ctx *gin.Context) {
	loanID, _ := strconv.Atoi(ctx.Param("loan_id"))
	var req models.ApproveLoanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loan, err := c.Service.ApproveLoan(uint(loanID), req.PictureProof, req.EmployeeID, req.Date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) InvestInLoan(ctx *gin.Context) {
	loanID, _ := strconv.Atoi(ctx.Param("loan_id"))
	var req models.InvestLoanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loan, err := c.Service.InvestInLoan(uint(loanID), req.InvestorID, req.Amount)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loan)
}

func (c *LoanController) DisburseLoan(ctx *gin.Context) {
	loanID, _ := strconv.Atoi(ctx.Param("loan_id"))
	var req models.DisburseLoanRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	loan, err := c.Service.DisburseLoan(uint(loanID), req.EmployeeID, req.AgreementLetterURL, req.Date)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, loan)
}
