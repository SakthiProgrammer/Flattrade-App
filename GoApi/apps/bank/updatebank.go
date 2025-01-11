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

type UpdateBankRec struct {
	ID         int    `json:"id" gorm:"column:Id"`
	BankName   string `json:"bank_name" gorm:"column:Bank_Name"`
	BranchName string `json:"branch_name" gorm:"column:Branch_name"`
	IFSCCODE   string `json:"ifsc_code" gorm:"column:Ifsc_code"`
	Address    string `json:"address" gorm:"column:Address"`
	// CreatedBy  string    `json:"created_by" gorm:"column:Created_By"`
	// CreatedAt  string    `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy string    `json:"-" gorm:"column:Updated_By"`
	UpdatedAt time.Time `json:"-" gorm:"column:Updated_At"`
}

type UpdateBankResp struct {
	BankDetailsArr UpdateBankRec `json:"bank_details"`
	ErrMsg         string        `json:"errMsg"`
	Status         string        `json:"status"`
}

func UpdateBank(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("UpdateBank-(+)")

	var lBankResp UpdateBankResp
	var lBank UpdateBankRec

	if r.Method == http.MethodPut {

		lBankResp.Status = common.SuccessCode

		lBody, lErr := io.ReadAll(r.Body)

		lErr = json.Unmarshal(lBody, &lBank)

		if lErr != nil {
			log.Println("BUPB-001", lErr.Error())
			lBankResp.Status = common.ErrorCode
			lBankResp.ErrMsg = lErr.Error()

		} else {

		}

		if lErr != nil {
			log.Println("")
			lBankResp.Status = common.ErrorCode
			lBankResp.ErrMsg = lErr.Error()
		} else {
			updateBankInDB(&lBankResp, &lBank)
		}

	} else {

		lBankResp.ErrMsg = "Invalid Method"
		lBankResp.Status = common.ErrorCode

	}

	lData, lErr := json.Marshal(lBankResp)

	if lErr != nil {
		log.Println("BGB-003", lErr.Error())
		lBankResp.Status = common.ErrorCode
		lBankResp.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("UpdateBank-(-)")

}

func updateBankInDB(lBankResp *UpdateBankResp, lBank *UpdateBankRec) {

	log.Println("updateBankInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {

		log.Println("BUBINDB-001", lErr.Error())
		lBankResp.ErrMsg = lErr.Error()
		lBankResp.Status = common.ErrorCode

	} else {

		fmt.Println(lBankResp)
		config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
		AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
		// lBank.CreatedBy = "Admin: " + AdminId
		// lBank.CreatedAt = time.Now()
		lBank.UpdatedBy = "Admin: " + AdminId
		lBank.UpdatedAt = time.Now()

		lResult := lGormDB.Table("st_918_bank_details").Where("id = ?", lBank.ID).UpdateColumns(&lBank)

		if lResult.Error != nil {
			log.Println("BUBINDB-002", lResult.Error)
			lBankResp.Status = common.ErrorCode
			lBankResp.ErrMsg = lResult.Error.Error()

		} else {
			lBankResp.BankDetailsArr = *lBank
			lBankResp.Status = common.SuccessCode
		}
	}

	log.Println("updateBankInDB-(-)")

}
