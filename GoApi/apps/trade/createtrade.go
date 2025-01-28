package trade

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
)

type CreateTradeRec struct {
	ClientID                  uint      `json:"client_id" gorm:"column:client_id"`
	TradeID                   uint      `json:"trade_id" gorm:"-"`
	StockID                   uint      `json:"stock_id" gorm:"column:stock_id"`
	TradeType                 string    `json:"trade_type" gorm:"column:trade_type"`
	Quantity                  uint      `json:"quantity" gorm:"column:quantity"`
	TradePrice                float64   `json:"trade_price" gorm:"column:trade_price"`
	TradeDate                 time.Time `json:"trade_date" gorm:"column:trade_date"`
	BackOfficerApprovalStatus string    `json:"back_officer_approval_status" gorm:"column:back_officer_approval_status"`
	BillerApprovalStatus      string    `json:"biller_approval_status" gorm:"column:biller_Approvel_status"`
	ApproverApprovalStatus    string    `json:"approver_approval_status" gorm:"column:approver_Approvel_status"`
	CreatedBy                 string    `json:"created_by" gorm:"column:Created_By"`
	CreatedAt                 time.Time `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy                 string    `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt                 time.Time `json:"updated_at" gorm:"column:Updated_At"`
}

type CreateTradeResp struct {
	TradeRec CreateTradeRec `json:"trade_details"`
	ErrMsg   string         `json:"errMsg"`
	Status   string         `json:"status"`
}

func CreateTrade(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("CreateBank-(+)")

	var lTradeResp CreateTradeResp
	var lTrade CreateTradeRec

	lTradeResp.Status = common.SuccessCode

	if r.Method == http.MethodPost {

		lBody, lErr := io.ReadAll(r.Body)

		if lErr != nil {

			log.Println("TCT-001", lErr.Error())
			lTradeResp.Status = common.ErrorCode
			lTradeResp.ErrMsg = lErr.Error()

		} else {

			lErr = json.Unmarshal(lBody, &lTrade)

			if lErr != nil {
				log.Println("TCT-002", lErr.Error())
				lTradeResp.Status = common.ErrorCode
				lTradeResp.ErrMsg = lErr.Error()

			} else {

				createTradeInDB(&lTradeResp, &lTrade)
			}
		}

	} else {

		lTradeResp.Status = common.ErrorCode
		lTradeResp.ErrMsg = "Invalid Method"
	}

	lData, lErr := json.Marshal(lTradeResp)

	if lErr != nil {
		log.Println("TCT-003", lErr.Error())
		lTradeResp.Status = common.ErrorCode
		lTradeResp.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}
	log.Println("CreateBank-(-)")

}

func createTradeInDB(lTradeResp *CreateTradeResp, lTrade *CreateTradeRec) {

	log.Println("createTradeInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {
		log.Println("TCTIDB-001", lErr.Error())
		lTradeResp.ErrMsg = lErr.Error()
		lTradeResp.Status = common.ErrorCode

	} else {

		var lName string
		lResult := lGormDB.Raw("SELECT CONCAT(first_name, ' ', last_name) AS created_by FROM st_918_client_table WHERE client_id = ?", lTrade.ClientID).Scan(&lName)
		if lResult.Error != nil {
			fmt.Println("TCTIDB-002", lResult.Error)
		} else {

			lCurrTime := time.Now()
			lTrade.CreatedBy = lName
			lTrade.CreatedAt = lCurrTime
			lTrade.UpdatedBy = lName
			lTrade.UpdatedAt = lCurrTime
			lPending := common.Pending
			lTrade.ApproverApprovalStatus = lPending
			lTrade.BackOfficerApprovalStatus = lPending
			lTrade.BillerApprovalStatus = lPending

			var lResult *gorm.DB

			if lTrade.TradeType == "BUY" {

				lResult = lGormDB.Table("st_918_trade_table").Create(&lTrade)

			} else {

				var lQuantity uint
				lResult = lGormDB.Table("st_918_trade_table").
					Select("quantity").
					Where("Id = ?", lTrade.TradeID).
					Scan(&lQuantity)

				if lQuantity < lTrade.Quantity {
					lTradeResp.ErrMsg = "You Don't have that quantity to SELL"
					lTradeResp.Status = common.ErrorCode
				} else {

					lResult = lGormDB.Exec(`
	UPDATE st_918_trade_table 
	SET 
		quantity = (SELECT quantity FROM st_918_trade_table WHERE Id = ?) - ?, 
		trade_price = ? 
	WHERE Id = ?`,
						lTrade.TradeID,
						lTrade.Quantity,
						lTrade.TradePrice, lTrade.TradeID)

					lResult = lGormDB.Table("st_918_trade_table").Create(&lTrade)
				}
			}

			if lResult.Error != nil {
				log.Println("TCTIDB-003", lResult.Error)
				lTradeResp.Status = common.ErrorCode
				lTradeResp.ErrMsg = lResult.Error.Error()

			} else {
				lTradeResp.TradeRec = *lTrade
				lTradeResp.Status = common.SuccessCode
			}
		}

	}

	log.Println("createTradeInDB-(-)")

}
