package client

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type GetClientRec struct {
	ClientID       uint   `json:"client_id" gorm:"column:client_id"`
	FirstName      string `json:"first_name" gorm:"column:first_name"`
	LastName       string `json:"last_name" gorm:"column:last_name"`
	Email          string `json:"email" gorm:"column:email"`
	PhoneNumber    string `json:"phone_number" gorm:"column:phone_number"`
	PanNumber      string `json:"pan_number" gorm:"column:pan_number"`
	NomineeName    string `json:"nominee_name" gorm:"column:nominee_name"`
	BankID         uint   `json:"bank_id" gorm:"column:bank_id"`
	UniqueId       string `json:"unique_id" gorm:"column:unique_id"`
	BankAccount    string `json:"bank_account" gorm:"column:bank_account"`
	KycIsCompleted string `json:"kyc_is_completed" gorm:"column:kyc_iScompleted"`
	CreatedBy      string `json:"created_by" gorm:"column:Created_by"`
	CreatedAt      string `json:"created_at" gorm:"column:Created_at"`
	UpdatedBy      string `json:"updated_by" gorm:"column:updated_by"`
	UpdatedAt      string `json:"updated_at" gorm:"column:updated_at"`
}

type GetClientResp struct {
	ClientArr []GetClientRec `json:"client_details"`
	ErrMsg    string         `json:"errMsg"`
	Status    string         `json:"status"`
}

func GetClients(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ID, CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	lId := r.Header.Get("ID")

	log.Println("GetClients-(+)")

	var lClientResp GetClientResp

	// var lStock Stock

	if r.Method == http.MethodGet {
		lClientResp.Status = common.SuccessCode

		lGormDb, lErr := gormdb.GormDBConnection()

		lSql, _ := lGormDb.DB()

		defer lSql.Close()

		if lErr != nil {
			log.Println("CGCS-001", lErr.Error())
			lClientResp.Status = common.ErrorCode
			lClientResp.ErrMsg = lErr.Error()

		} else {

			if lId == "" {

				lResult := lGormDb.Table("st_918_client_table").Where("kyc_iScompleted = ?", common.Pending).Find(&lClientResp.ClientArr)

				if lResult.Error != nil {
					log.Println("CGCS-002", lResult.Error)
					lClientResp.ErrMsg = lResult.Error.Error()
					lClientResp.Status = common.ErrorCode
				} else {
					log.Println("Successfully Get")
					lClientResp.Status = common.SuccessCode
					fmt.Printf("%+v", lClientResp)
				}

			} else if lId != "" {

				lId := strToInt(lId, &lClientResp)

				if lId != -1 {
					lResult := lGormDb.Table("st_918_client_table").Where("client_id=?", lId).Find(&lClientResp.ClientArr)

					if lResult.Error != nil {
						log.Println("CGCS-003", lResult.Error)
						lClientResp.ErrMsg = lResult.Error.Error()
						lClientResp.Status = common.ErrorCode
					} else {
						log.Println("Successfully Get")
						lClientResp.Status = common.SuccessCode
						fmt.Printf("%+v", lClientResp)
					}

				}
			}

		}

	} else {
		log.Println("CGCS-004", "Invalid Method")
		lClientResp.ErrMsg = "Invalid Method"
		lClientResp.Status = common.ErrorCode
	}

	lData, lErr := json.Marshal(lClientResp)

	if lErr != nil {
		log.Println("CGCS-005", lErr.Error())
		lClientResp.ErrMsg = lErr.Error()
		lClientResp.Status = common.SuccessCode
	} else {
		fmt.Fprintf(w, string(lData))
	}
	log.Println("GetClients-(-)")
}

func strToInt(lId string, lClientResp *GetClientResp) int {
	log.Println("strToInt-(+)")
	lIndId, lErr := strconv.Atoi(lId)

	if lErr != nil {
		lClientResp.Status = common.ErrorCode
		lClientResp.ErrMsg = lErr.Error()
		return -1
	}
	log.Println("strToInt-(-)")
	return lIndId
}
