package bank

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"log"
	"net/http"
)

type BankRec struct {
	ID         int    `json:"id" gorm:"column:Id"`
	BankName   string `json:"bank_name" gorm:"column:Bank_Name"`
	BranchName string `json:"branch_name" gorm:"column:Branch_Name"`
	IFSCCODE   string `json:"ifsc_code" gorm:"column:Ifsc_code"`
	Address    string `json:"address" gorm:"column:Address"`
	CreatedBy  string `json:"created_by" gorm:"column:Created_By"`
	CreatedAt  string `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy  string `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt  string `json:"updated_at" gorm:"column:Updated_At"`
}

type BankResp struct {
	BankDetailsArr []BankRec `json:"bank_details"`
	ErrMsg         string    `json:"errMsg"`
	Status         string    `json:"status"`
}

func GetBanks(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("GetBank-(+)")

	var lBankResp BankResp
	lBankResp.Status = common.SuccessCode

	if r.Method == http.MethodGet {

		lGormDb, lErr := gormdb.GormDBConnection()

		if lErr != nil {

			log.Println("BCB-001", lErr.Error())
			lBankResp.Status = common.ErrorCode
			lBankResp.ErrMsg = lErr.Error()

		} else {

			lResult := lGormDb.Table("st_918_bank_details").Find(&lBankResp.BankDetailsArr)

			if lResult.Error != nil {

				log.Println("BCB-002", lResult.Error)
				lBankResp.Status = common.ErrorCode
				lBankResp.ErrMsg = lResult.Error.Error()

			} else {
				log.Println("Success")
				lBankResp.Status = common.SuccessCode
				fmt.Printf("%+v", lBankResp)
			}

		}

	} else {

		lBankResp.Status = common.ErrorCode
		lBankResp.ErrMsg = "Invalid Method"
	}

	lData, lErr := json.Marshal(lBankResp)

	if lErr != nil {
		log.Println("BGB-003", lErr.Error())
		lBankResp.Status = common.ErrorCode
		lBankResp.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}
	log.Println("GetBank-(-)")
}
