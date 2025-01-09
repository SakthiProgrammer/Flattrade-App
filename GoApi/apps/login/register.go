package login

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	common "flattrade/common"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	// ClientID       int       `json:"client_id" gorm:"column:client_id"`
	PhoneNumber    string `json:"phone_number"gorm:"column:phone_number"`
	FirstName      string `json:"first_name" gorm:"column:first_name"`
	LastName       string `json:"last_name" gorm:"column:last_name"`
	PanNumber      string `json:"pan_number" gorm:"column:pan_number"`
	NomineeName    string `json:"nominee_name" gorm:"column:nominee_name"`
	KycIsCompleted bool   `json:"kyc_is_completed" gorm:"column:kyc_iscompleted"`
	BankAccount    string `json:"bank_account" gorm:"column:bank_account"`
	Email          string `json:"email" gorm:"column:email"`
	// BankID         int       `json:"bank_id" gorm:"column:bank_id"`
	Password  string    `json:"password"gorm:"column:password"`
	CreatedBy string    `json:"created_by"gorm:"column:created_by"`
	CreatedAt time.Time `json:"created_at"gorm:"column:created_at"`
	UpdatedBy string    `json:"updated_by"gorm:"column:updated_by"`
	UpdatedAt time.Time `json:"updated_at"gorm:"column:updated_at"`
}

type ClinetResponse struct {
	ClientInfo Client `json:"client_info"`
	ErrMsg     string `json:"err_msg"`
	Status     string `json:"status"`
}

func Register(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("LoginUser-(+)")

	log.Println("RegisterUser-(+)")

	var lClientResponse ClinetResponse
	var lClient Client

	lClientResponse.Status = common.SuccessCode

	if r.Method == http.MethodPost {

		lBody, lErr := io.ReadAll(r.Body)

		if lErr != nil {

			log.Println("LRU-001", lErr.Error())
			lClientResponse.ErrMsg = lErr.Error()
			lClientResponse.Status = common.ErrorCode

		} else {

			lErr = json.Unmarshal(lBody, &lClient)

			if lErr != nil {

				log.Println("LRU-002", lErr.Error())
				lClientResponse.ErrMsg = lErr.Error()
				lClientResponse.Status = common.ErrorCode

			} else {

				createClient(&lClientResponse, &lClient)

			}
		}

	} else {
		lClientResponse.Status = common.ErrorCode
		lClientResponse.ErrMsg = "Invalid Method"
	}

	lData, lErr := json.Marshal(lClientResponse)

	if lErr != nil {
		log.Println("")
		lClientResponse.Status = common.ErrorCode
		lClientResponse.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}
	log.Println("RegisterUser-(-)")
}

func createClient(lClientResponse *ClinetResponse, lClient *Client) {

	log.Println("createClient-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	if lErr != nil {

		log.Println("LCU-001", lErr.Error())
		lClientResponse.ErrMsg = lErr.Error()
		lClientResponse.Status = common.ErrorCode

	} else {

		fmt.Println(lClientResponse)

		lClient.CreatedBy = lClient.FirstName + " " + lClient.LastName
		lClient.CreatedAt = time.Now()
		lClient.UpdatedBy = lClient.FirstName + " " + lClient.LastName
		lClient.UpdatedAt = time.Now()

		lResult := lGormDB.Table("st_918_client_table").Create(lClient)

		if lResult.Error != nil {
			log.Println("LCU-002", lResult.Error.Error())
			lClientResponse.Status = common.ErrorCode
			lClientResponse.ErrMsg = lResult.Error.Error()

		} else {

			lClientResponse.ClientInfo = *lClient
			lClientResponse.Status = common.SuccessCode
		}
	}
	log.Println("createClient-(-)")

}
