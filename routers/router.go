package routers

import (
	"github.com/ak-yudha/loan-service/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(loanController *controllers.LoanController) *gin.Engine {
	router := gin.Default()

	router.POST("/loans", loanController.CreateLoan)
	router.POST("/loans/:loan_id/approve", loanController.ApproveLoan)
	router.POST("/loans/:loan_id/invest", loanController.InvestInLoan)
	router.POST("/loans/:loan_id/disburse", loanController.DisburseLoan)

	return router
}
