package user

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// client_id
// stock_id
// trade_type
// quantity
// trade_price
// trade_date
// back_officer_approval_status
// biller_Approvel_status
// approver_Approvel_status
// Created_By
// Created_At
// Updated_By
// Updated_At

type GetStockRec struct {
	StockID    uint    `json:"stock_id" gorm:"column:stock_id"`
	StockName  string  `json:"stock_name" gorm:"column:stock_name"`
	StockPrice float64 `json:"stock_price" gorm:"column:stock_price"`
	Segment    string  `json:"segment" gorm:"column:segment"`
	ISIN       string  `json:"isin" gorm:"column:isin"`
	// CreatedBy  string  `json:"created_by" gorm:"column:created_by"`
	// CreatedAt  string  `json:"created_at" gorm:"column:created_at"`
	// UpdatedBy  string  `json:"updated_by" gorm:"column:updated_by"`
	// UpdatedAt  string  `json:"updated_at" gorm:"column:updated_at"`
}

func (GetStockRec) TableName() string {
	return "st_918_stock_table"
}

/* ------ */

type GetTradeRec struct {
	TradeID             uint    `json:"trade_id" gorm:"column:Id"`
	ClientID            uint    `json:"client_id" gorm:"column:client_id"`
	TradeType           string  `json:"trade_type" gorm:"column:trade_type"`
	Quantity            uint    `json:"quantiry" gorm:"column:quantity"`
	TotalPrice          float64 `json:"total_price" gorm:"column:trade_price"`
	TradeDate           string  `json:"trade_date" gorm:"column:trade_date"`
	BackOfficerApproval string  `json:"back_officer_approval" gorm:"column:back_officer_approval_status"`
	BillerApproval      string  `json:"biller_approval" gorm:"column:biller_Approvel_status"`
	ApproverApproval    string  `json:"approver_approval" gorm:"column:approver_Approvel_status"`
}

func (GetTradeRec) TableName() string {
	return "st_918_trade_table"
}

// Id
// Bank_Name
// Branch_name
// Ifsc_code
// Address
// Created_By
// Created_At
// Updated_By
// // Updated_At
// type GetBankRec struct {
// 	BankID     uint   `json:"-" gorm:"column:Id"`
// 	BankName   string `json:"bank_name" gorm:"column:Bank_Name"`
// 	BranchName string `json:"branch_name" gorm:"column:Branch_name"`
// 	IFSCCode   string `json:"ifsc_code" gorm:"column:Ifsc_code"`
// 	Address    string `json:"address" gorm:"column:Address"`
// }

// func (GetBankRec) TableName() string {
// 	return "st_918_bank_details"
// }

// client_id
// first_name
// last_name
// email
// phone_number
// pan_number
// nominee_name
// bank_id
// bank_account
// kyc_iScompleted
// Created_by
// Created_at
// updated_by
// updated_at
// password
// unique_id
type GetClientRec struct {
	ClientID    uint   `json:"client_id" gorm:"column:client_id"`
	FirstName   string `json:"first_name" gorm:"column:first_name"`
	LastName    string `json:"last_name" gorm:"column:last_name"`
	Email       string `json:"email" gorm:"column:email"`
	BankID      uint   `json:"-" gorm:"column:bank_id"`
	PhoneNumber string `json:"phone_number" gorm:"column:phone_number"`
	PanNumber   string `json:"pan_number" gorm:"column:pan_number"`
	NomineeName string `json:"nominee_name" gorm:"column:nominee_name"`
	// BankDetail     GetBankRec    `json:"bank_detail" gorm:"foreignKey:BankID;references:BankID"`
	UniqueId       string        `json:"unique_id" gorm:"column:unique_id"`
	BankAccount    string        `json:"bank_account" gorm:"column:bank_account"`
	KycIsCompleted string        `json:"kyc_is_completed" gorm:"column:kyc_iScompleted"`
	TradesArr      []GetTradeRec `json:"trade_details" gorm:"foreignKey:ClientID;references:ClientID"`
}

func (GetClientRec) TableName() string {
	return "st_918_client_table"
}

type GetClientResp struct {
	GetClientRec `json:"client_details"`
	Status       string `json:"status"`
	ErrMsg       string `json:"errMsg"`
}

func ClientTradeFullDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "ROLE, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("ClientTradeFullDetails-(+)")

	// var lClient GetClientRec
	var lClientResp GetClientResp

	lClientResp.Status = common.SuccessCode

	lUserRole := r.Header.Get("ROLE")
	// lUserId := r.Header.Get("USERID")

	if r.Method == http.MethodGet {

		lUserRole = strings.ToUpper(lUserRole)
		fmt.Println(lUserRole)

		if common.ValidateUserRole(lUserRole) {

			fmt.Println("helo")
			lClientResp.Status = common.ErrorCode
			lClientResp.ErrMsg = "Provide Role in Header"

		} else {

			fmt.Println("Hi")

			lGormDb, lErr := gormdb.GormDBConnection()

			lSql, _ := lGormDb.DB()

			defer lSql.Close()

			if lErr != nil {
				log.Println("UCTFD-001", lErr.Error())
				lClientResp.Status = common.ErrorCode
				lClientResp.ErrMsg = lErr.Error()

			} else {

				//  back_officer_approval_status
				// biller_Approvel_status
				// approver_Approvel_status

				/* ==== Roles ====

				Back Officer - BO [Find back_officer_approval_status = pending]
				Biller - B [Find back_officer_approval_status = Approved]
				Approver - APPR [Find biller_Approvel_status = Approved]

				*/

				var lColumnName string
				var lStatus string
				lColumnName = "back_officer_approval_status"
				if lUserRole == "BO" {
					lStatus = common.Pending
				} else if lUserRole == "B" {
					lStatus = common.Approved
				} else if lUserRole == "APPR" {
					lColumnName = "biller_Approvel_status"
					lStatus = common.Approved
				}

				/* lResult := lGormDb.Table("st_918_client_table AS c").
				Joins("JOIN st_918_trade_table AS t ON t.client_id = c.client_id").
				Joins("JOIN st_918_stock_table AS st ON t.stock_id = st.stock_id").
				Select(`
					c.client_id, c.first_name, c.last_name, c.email, c.phone_number, c.pan_number,
					c.nominee_name, c.unique_id, c.bank_account, c.kyc_iscompleted,
					t.trade_id, t.trade_type, t.quantity, t.trade_price AS total_price,
					t.trade_date, t.back_officer_approval AS back_officer_approval,
					t.biller_approvel_status AS biller_approval, t.approver_approvel_status AS approver_approval,
					st.stock_id, st.stock_name, st.stock_price, st.segment, st.isin
				`).
				Where(lColumnName+" = ?", lStatus).
				Find(&lClientResp.GetClientRec)
				*/

				lResult := lGormDb.Table("st_918_client_table c").Joins("JOIN st_918_trade_table  AS t ON t.client_id = c.client_id").Joins("JOIN st_918_stock_table AS st ON t.stock_id = st.stock_id").Preload("TradesArr").
					Where("t."+lColumnName+" = ?", lStatus).Find(&lClientResp.GetClientRec)
				// lResult := lGormDb.Preload("TradesArr").Preload("BankDetail").Where("client_id = ?", lUserId).First(&lClientResp.GetClientRec)
				// lResult := lGormDb.Preload("TradesArr").Where("client_id = ?", lUserId).First(&lClientResp.GetClientRec)

				if lResult.Error != nil {
					log.Println("UCTFD-002", lResult.Error)
					lClientResp.ErrMsg = lResult.Error.Error()
					lClientResp.Status = common.ErrorCode
				} else {
					log.Println("Successfully Get")
					fmt.Println("%+v", lClientResp)
				}
			}
		}
	} else {
		lClientResp.Status = common.ErrorCode
		lClientResp.ErrMsg = "Invalid Method"
	}

	lData, lErr := json.Marshal(lClientResp)

	if lErr != nil {
		log.Println("UCTFD-003", lErr.Error())
		lClientResp.ErrMsg = lErr.Error()
		lClientResp.Status = common.ErrorCode
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("ClientTradeFullDetails-(-)")
}
