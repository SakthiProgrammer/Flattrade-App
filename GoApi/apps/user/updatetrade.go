package user

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type TradeStatusRec struct {
	TradeID        uint   `json:trade_id`
	Role           string `json:"role"`
	UserID         string `json:"user_id` // for updated by
	ApprovalStatus string `json:"status"`
}

type UpdateTradeResp struct {
	// Trade  TradeStatusRec `json:"trade_id"`
	TradeStatus TradeStatusRec `json:"trade_status" `
	Status      string         `json:"status" `
	ErrMsg      string         `json:"errMsg" `
}

func UpdateTradeStatus(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "ROLE, USERID, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("UpdateTradeStatus-(+)")

	var lResp UpdateTradeResp
	if r.Method == http.MethodPut {
		var lTrade TradeStatusRec

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

				lTrade.Role = strings.ToUpper(lTrade.Role)
				if common.ValidateUserRole(lTrade.Role) {

					updateTradeINDB(&lResp, lTrade)

				} else {

					lResp.Status = common.ErrorCode
					lResp.ErrMsg = "Invalid User Role"
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

	lSql, _ := lGormDb.DB()

	defer lSql.Close()

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
			Where("Id = ?", lTrade.TradeID).
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

			lResp.Status = common.SuccessCode
			lResp.TradeStatus = lTrade
			fmt.Println("%+v", lResp)
		}
	}

	log.Println("updateTradeINDB-(-)")

}
