package user

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type TradeStatusRec struct {
	TradeID        uint   `json:trade_id`
	Role           string `json:"role"`
	UserID         string `json:"user_id` // for updated by
	ApprovalStatus string `json:"status"`
}

// type UpdateTradeRec struct {
// 	TradeID             uint    `json:"trade_id" gorm:"column:Id"`
// 	ClientID            uint    `json:"client_id" gorm:"column:client_id"`
// 	TradeType           string  `json:"trade_type" gorm:"column:trade_type"`
// 	Quantity            uint    `json:"quantiry" gorm:"column:quantity"`
// 	TotalPrice          float64 `json:"total_price" gorm:"column:trade_price"`
// 	TradeDate           string  `json:"trade_date" gorm:"column:trade_date"`
// 	BackOfficerApproval string  `json:"back_officer_approval" gorm:"column:back_officer_approval_status"`
// 	BillerApproval      string  `json:"biller_approval" gorm:"column:biller_Approvel_status"`
// 	ApproverApproval    string  `json:"approver_approval" gorm:"column:approver_Approvel_status"`
// }

type UpdateTradeResp struct {
	// Trade  TradeStatusRec `json:"trade_id"`
	TradeStatus TradeStatusRec `json:"trade_status" `
	Status      string         `json:"status" `
	ErrMsg      string         `json:"status" `
}

func UpdateTradeStatus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "ROLE, USERID, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("UpdateTradeStatus-(+)")

	var lResp UpdateTradeResp
	var lTrade TradeStatusRec
	if r.Method != http.MethodPut {

		lData, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			log.Println("UUTS-001", lErr.Error())
			lResp.ErrMsg = lErr.Error()
			lResp.Status = common.ErrorCode

		} else {
			lErr = json.Unmarshal(lData, &lTrade)

			if lErr != nil {

				log.Println("UUTS-002", lErr.Error())
				lResp.ErrMsg = lErr.Error()
				lResp.Status = common.ErrorCode
			} else {

				if lErr != nil {
					log.Println("UUTS-003", lErr.Error())
					lResp.ErrMsg = lErr.Error()
					lResp.Status = common.ErrorCode

				} else {

					updateTradeINDB(&lResp, lTrade)

				}

			}
		}

	} else {
		lResp.Status = common.ErrorCode
		lResp.ErrMsg = "Invalid Method"
	}

	lData, lErr := json.Marshal(lResp)

	if lErr != nil {
		log.Println("UUTS-005", lErr.Error())
		lResp.ErrMsg = lErr.Error()
		lResp.Status = common.ErrorCode
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("UpdateTradeStatus-(-)")

}

func updateTradeINDB(lResp *UpdateTradeResp, lTrade TradeStatusRec) {

	log.Println("updateTradeINDB-(+)")

	lGormDb, lErr := gormdb.GormDBConnection()

	if lErr != nil {
		log.Println("UUTTINDB-001", lErr.Error())
		lResp.ErrMsg = lErr.Error()
		lResp.Status = common.ErrorCode

	} else {

		// back_officer_approval_status
		// biller_Approvel_status
		// approver_Approvel_status
		lColumn := ""

		switch lTrade.Role {
		case "BO":
			lColumn = "back_officer_approval_status"
		case "B":
			lColumn = "biller_Approvel_status"
		case "APPR":
			lColumn = "approver_Approvel_status"
		default:
			log.Println(lTrade.Role)
		}

		lResult := lGormDb.Table("st_918_trade_table").
			Where("trade_id = ?", lTrade.TradeID).
			Updates(map[string]interface{}{
				lColumn:      lTrade.ApprovalStatus,
				"updated_by": lTrade.UserID,
				"updated_at": time.Now(),
			})

		if lResult.Error != nil {
			log.Println("UUTTINDB-002", lResult.Error)
			lResp.Status = common.ErrorCode
			lResp.ErrMsg = lResult.Error.Error()
		} else {
			fmt.Println("%+v", lResp)
		}
	}

	log.Println("updateTradeINDB-(-)")

}
