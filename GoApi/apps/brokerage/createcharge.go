package brokerage

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"flattrade/genpkg"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type CreateChargeRec struct {
	// ChargeID         uint      `json:"charge_id" gorm:"column:charge_iD"`
	ChargeType       string    `json:"charge_type" gorm:"column:charge_type"`
	ChargePercentage float64   `json:"charge_percentage" gorm:"column:charge_percentage"`
	EffectiveDate    time.Time `json:"effective_date" gorm:"column:effective_date"`
	CreatedBy        string    `json:"created_by" gorm:"column:Created_By"`
	CreatedAt        time.Time `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy        string    `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"column:Updated_At"`
}

type CreateChargeResp struct {
	ChargeRec CreateChargeRec `json:"charge_details"`
	ErrMsg    string          `json:"errMsg"`
	Status    string          `json:"status"`
}

func CreateChager(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("CreateChager-(+)")

	var lChargeResp CreateChargeResp
	var lCharge CreateChargeRec

	lChargeResp.Status = common.SuccessCode
	if r.Method == http.MethodPost {

		lBody, lErr := io.ReadAll(r.Body)

		if lErr != nil {

			log.Println("BCC-001", lErr.Error())
			lChargeResp.Status = common.ErrorCode
			lChargeResp.ErrMsg = lErr.Error()

		} else {

			lErr = json.Unmarshal(lBody, &lCharge)

			if lErr != nil {
				log.Println("BCC-002", lErr.Error())
				lChargeResp.Status = common.ErrorCode
				lChargeResp.ErrMsg = lErr.Error()

			} else {

				createChargeInDB(&lChargeResp, &lCharge)
			}
		}

	} else {

		lChargeResp.Status = common.ErrorCode
		lChargeResp.ErrMsg = "Invalid Method"
	}

	lData, lErr := json.Marshal(lChargeResp)

	if lErr != nil {
		log.Println("BCC-003", lErr.Error())
		lChargeResp.Status = common.ErrorCode
		lChargeResp.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}
	log.Println("CreateChager-(-)")

}

func createChargeInDB(lChargeResp *CreateChargeResp, lCharge *CreateChargeRec) {

	log.Println("createChargeInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {

		log.Println("BCCIDB-001", lErr.Error())
		lChargeResp.ErrMsg = lErr.Error()
		lChargeResp.Status = common.ErrorCode

	} else {

		fmt.Println(lChargeResp)
		config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
		AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
		lName := "Admin: " + AdminId

		// Clean up the EffectiveDate string by trimming any unwanted quotes
		cleanedEffectiveDate := strings.Trim(lCharge.EffectiveDate.Format("2006-01-02T15:04:05Z07:00"), "\"")

		// Parse the cleaned date string (without the time and timezone)
		parsedTime, err := time.Parse("2006-01-02T15:04:05Z07:00", cleanedEffectiveDate)
		if err != nil {
			log.Println("Error parsing EffectiveDate:", err)
			lChargeResp.ErrMsg = "Invalid date format"
			lChargeResp.Status = common.ErrorCode
			return
		}

		// Extract the date part only (set the time to midnight)
		parsedDateOnly := parsedTime.Truncate(24 * time.Hour) // Truncate time to 00:00:00

		lCharge.EffectiveDate = parsedDateOnly

		lCurrTime := time.Now()
		lCharge.CreatedBy = lName
		lCharge.CreatedAt = lCurrTime
		lCharge.UpdatedBy = lName
		lCharge.UpdatedAt = lCurrTime

		lResult := lGormDB.Table("st_919_config_charges_table").Create(&lCharge)

		if lResult.Error != nil {
			log.Println("BCCIDB-002", lResult.Error)
			lChargeResp.Status = common.ErrorCode
			lChargeResp.ErrMsg = lResult.Error.Error()

		} else {
			lChargeResp.ChargeRec = *lCharge
			lChargeResp.Status = common.SuccessCode
		}
	}

	log.Println("createChargeInDB-(-)")

}
