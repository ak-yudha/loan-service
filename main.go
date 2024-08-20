package main

import (
	"github.com/ak-yudha/loan-service/controllers"
	"github.com/ak-yudha/loan-service/repositories"
	"github.com/ak-yudha/loan-service/routers"
	"github.com/ak-yudha/loan-service/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/loan_engine?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Loan{}, &models.ApprovalInfo{}, &models.Investment{}, &models.DisbursementInfo{})

	loanRepo := repositories.NewMySQLLoanRepository(db)
	loanService := services.NewLoanService(loanRepo)
	loanController := controllers.NewLoanController(loanService)

	r := routers.SetupRouter(loanController)
	r.Run(":8080")
}
