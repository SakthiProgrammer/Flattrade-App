package client

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

type UpdateClientRec struct {
	ClientID       uint   `json:"client_id" gorm:"column:client_id"`
	FirstName      string `json:"first_name" gorm:"column:first_name"`
	LastName       string `json:"last_name" gorm:"column:last_name"`
	Email          string `json:"email" gorm:"column:email"`
	PhoneNumber    string `json:"phone_number" gorm:"column:phone_number"`
	PanNumber      string `json:"pan_number" gorm:"column:pan_number"`
	NomineeName    string `json:"nominee_name" gorm:"column:nominee_name"`
	BankID         uint   `json:"bank_id" gorm:"column:bank_id"`
	BankAccount    string `json:"bank_account" gorm:"column:bank_account"`
	KycIsCompleted string `json:"kyc_is_completed" gorm:"column:kyc_iScompleted"`
	// CreatedBy      string `json:"created_by" gorm:"column:Created_by"`
	// CreatedAt      string `json:"created_at" gorm:"column:Created_at"`
	UpdatedBy string    `json:"_" gorm:"column:updated_by"`
	UpdatedAt time.Time `json:"_" gorm:"column:updated_at"`
}

type UpdateClientResp struct {
	Client UpdateClientRec `json:"client_details"`
	ErrMsg string          `json:"errMsg"`
	Status string          `json:"status"`
}

func UpdateClient(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("UpdateClient-(+)")

	var lClientResp UpdateClientResp
	var lClient UpdateClientRec

	if r.Method == http.MethodPut {

		lClientResp.Status = common.SuccessCode

		lBody, lErr := io.ReadAll(r.Body)

		lErr = json.Unmarshal(lBody, &lClient)

		if lErr != nil {
			log.Println("CUC-001", lErr.Error())
			lClientResp.Status = common.ErrorCode
			lClientResp.ErrMsg = lErr.Error()

		} else {

		}

		if lErr != nil {
			log.Println("CUC-002", lErr.Error())
			lClientResp.Status = common.ErrorCode
			lClientResp.ErrMsg = lErr.Error()
		} else {
			updateClientInDB(&lClientResp, &lClient)
		}

	} else {

		lClientResp.ErrMsg = "Invalid Method"
		lClientResp.Status = common.ErrorCode

	}

	lData, lErr := json.Marshal(lClientResp)

	if lErr != nil {
		log.Println("BGB-003", lErr.Error())
		lClientResp.Status = common.ErrorCode
		lClientResp.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("UpdateClient-(-)")

}

func updateClientInDB(lClientResp *UpdateClientResp, lClient *UpdateClientRec) {

	log.Println("updateClientInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {

		log.Println("CUCIDB-001", lErr.Error())
		lClientResp.ErrMsg = lErr.Error()
		lClientResp.Status = common.ErrorCode

	} else {

		fmt.Println(lClientResp)
		config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
		AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
		// lClient.CreatedBy = "Admin: " + AdminId
		// lClient.CreatedAt = time.Now()
		lClient.UpdatedBy = "Admin: " + AdminId
		lClient.UpdatedAt = time.Now()

		lResult := lGormDB.Table("st_918_client_table").Where("client_id = ?", lClient.ClientID).UpdateColumns(&lClient)

		if lResult.Error != nil {
			log.Println("BUBINDB-002", lResult.Error)
			lClientResp.Status = common.ErrorCode
			lClientResp.ErrMsg = lResult.Error.Error()

		} else {
			lClientResp.Client = *lClient
			lClientResp.Status = common.SuccessCode
		}
	}

	log.Println("updateClientInDB-(-)")

}
