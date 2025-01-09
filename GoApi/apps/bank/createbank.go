package bank

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"flattrade/genpkg"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Bank struct {
	BankName   string    `json:"bank_name" gorm:"column:bank_name"`
	BranchName string    `json:"branch_name" gorm:"column:branch_name"`
	IFSCCODE   string    `json:"ifsc_code" gorm:"column:ifsc_code"`
	Address    string    `json:"address" gorm:"column:address"`
	CreatedBy  string    `json:"created_by"gorm:"column:created_by"`
	CreatedAt  time.Time `json:"created_at"gorm:"column:created_at"`
	UpdatedBy  string    `json:"updated_by"gorm:"column:updated_by"`
	UpdatedAt  time.Time `json:"updated_at"gorm:"column:updated_at"`
}

type BankResponse struct {
	BankDetails Bank   `json:"bank_details"`
	ErrMsg      string `json:"errMsg"`
	Status      string `json:"status"`
}

func CreateBank(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("CreateBank-(+)")

	var lBankResponse BankResponse
	var lBank Bank

	lBankResponse.Status = common.SuccessCode
	if r.Method == http.MethodPost {

		lBody, lErr := io.ReadAll(r.Body)

		if lErr != nil {

			log.Println("BCB-001", lErr.Error())
			lBankResponse.Status = common.ErrorCode
			lBankResponse.ErrMsg = lErr.Error()

		} else {

			lErr = json.Unmarshal(lBody, &lBank)

			if lErr != nil {
				log.Println("BCB-002", lErr.Error())
				lBankResponse.Status = common.ErrorCode
				lBankResponse.ErrMsg = lErr.Error()

			} else {

				createBankInDB(&lBankResponse, &lBank)
			}
		}

	} else {

		lBankResponse.Status = common.ErrorCode
		lBankResponse.ErrMsg = "Invalid Method"
	}

	lData, lErr := json.Marshal(lBankResponse)

	if lErr != nil {
		log.Println("LRC", lErr.Error())
		lBankResponse.Status = common.ErrorCode
		lBankResponse.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}
	log.Println("CreateBank-(-)")

}

func createBankInDB(lBankResponse *BankResponse, lBank *Bank) {

	log.Println("createClient-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	if lErr != nil {

		log.Println("BCBIDB-001", lErr.Error())
		lBankResponse.ErrMsg = lErr.Error()
		lBankResponse.Status = common.ErrorCode

	} else {

		fmt.Println(lBankResponse)
		config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
		AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
		lBank.CreatedBy = "Admin: " + AdminId
		lBank.CreatedAt = time.Now()
		lBank.UpdatedBy = "Admin: " + AdminId
		lBank.UpdatedAt = time.Now()

		lResult := lGormDB.Table("st_918_bank_details").Create(lBank)

		if lResult.Error != nil {
			log.Println("BCBIDB-002", lResult.Error)
			lBankResponse.Status = common.ErrorCode
			lBankResponse.ErrMsg = lResult.Error.Error()

		} else {
			lBankResponse.BankDetails = *lBank
			lBankResponse.Status = common.SuccessCode
		}
	}
	log.Println("createClient-(-)")

}
