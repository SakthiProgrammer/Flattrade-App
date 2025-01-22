package trade

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"log"
	"net/http"
)

type GetTradeRec struct {
	TradeID             uint    `json:"trade_id"`
	TradeType           string  `json:"trade_type"`
	TotalPrice          float64 `json:"total_price"`
	TradeDate           string  `json:"trade_date"`
	BackOfficerApproval string  `json:"back_officer_approval"`
	BillerApproval      string  `json:"biller_approval"`
	ApproverApproval    string  `json:"approver_approval"`
}

func (GetTradeRec) TradeTable() string {
	return "st_918_trade_table"
}

type GetBankRec struct {
	BankId   uint   `json:"bank_id"`
	BankName string `json:"bank_name"`
	IFSCCode string `json:"ifsc_code"`
	Address  string `json:"address"`
}

func (GetBankRec) BankTable() string {
	return "st_919_bank_table"
}

type GetClientRec struct {
	ClientID       uint          `json:"client_id"`
	FirstName      string        `json:"first_name"`
	LastName       string        `json:"last_name"`
	Email          string        `json:"email"`
	PhoneNumber    string        `json:"phone_number"`
	PanNumber      string        `json:"pan_number"`
	NomineeName    string        `json:"nominee_name"`
	UniqueId       string        `json:"unique_id"`
	BankAccount    string        `json:"bank_account"`
	KycIsCompleted string        `json:"kyc_is_completed"`
	TradesArr      []GetTradeRec `json:"trade_details"`
	BankRec        GetBankRec    `json:"bank_detail"`
}

type GetClientResp struct {
	GetClientRec `json:"client_details"`
	Status       string `json:"status"`
	ErrMsg       string `json:"errMsg"`
}

func GetClientFullDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "USERID, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("GetClientFullDetails-(+)")

	var lClient GetClientRec
	var lClientResp GetClientResp

	lClientResp.Status = common.SuccessCode

	lUserId := r.Header.Get("USERID")

	if r.Method == http.MethodGet {

		if lUserId == "" {

			lClientResp.Status = common.ErrorCode
			lClientResp.ErrMsg = "Provide USERID in Header"

		} else {
			lGormDb, lErr := gormdb.GormDBConnection()

			lSql, _ := lGormDb.DB()

			defer lSql.Close()

			if lErr != nil {
				log.Println("TGCFD-001", lErr.Error())
				lClientResp.Status = common.ErrorCode
				lClientResp.ErrMsg = lErr.Error()

			} else {

				// lResult := lGormDb.Preload("TradesArr").Preload("BankRec").Where("client_id = ? AND trades_arr.back_officer_approval = ?", lUserId, "approved").First(&lClient)
				lResult := lGormDb.Preload("TradesArr").Preload("BankRec").Where("client_id = ?", lUserId).First(&lClient)

				if lResult.Error != nil {
					log.Println("TGCFD-002", lResult.Error)
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
		log.Println("TGCFD-", lErr.Error())
		lClientResp.ErrMsg = lErr.Error()
		lClientResp.Status = common.ErrorCode
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("GetClientFullDetails-(-)")
}
