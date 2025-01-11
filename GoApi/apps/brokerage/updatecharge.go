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
	"time"
)

type UpdateChargeRec struct {
	ChargeID         uint      `json:"charge_id" gorm:"column:charge_id"`
	ChargeType       string    `json:"charge_type" gorm:"column:charge_type"`
	ChargePercentage float64   `json:"charge_percentage" gorm:"column:charge_percentage"`
	EffectiveDate    time.Time `json:"effective_date" gorm:"column:effective_date"`
	// CreatedBy        string  `json:"created_by" gorm:"column:Created_By"`
	// CreatedAt        string  `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy string    `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:Updated_At"`
}

type UpdateChargeResp struct {
	Charge UpdateChargeRec `json:"charge_details"`
	ErrMsg string          `json:"errMsg"`
	Status string          `json:"status"`
}

func UpdateChage(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("UpdateChage-(+)")

	var lChargeResp UpdateChargeResp
	var lCharge UpdateChargeRec

	if r.Method == http.MethodPut {

		lChargeResp.Status = common.SuccessCode

		lBody, lErr := io.ReadAll(r.Body)

		lErr = json.Unmarshal(lBody, &lCharge)

		if lErr != nil {
			log.Println("BUC-001", lErr.Error())
			lChargeResp.Status = common.ErrorCode
			lChargeResp.ErrMsg = lErr.Error()

		} else {

		}

		if lErr != nil {
			log.Println("BUC-002", lErr.Error())
			lChargeResp.Status = common.ErrorCode
			lChargeResp.ErrMsg = lErr.Error()
		} else {
			updateChargeInDB(&lChargeResp, &lCharge)
		}

	} else {

		lChargeResp.ErrMsg = "Invalid Method"
		lChargeResp.Status = common.ErrorCode

	}

	lData, lErr := json.Marshal(lChargeResp)

	if lErr != nil {
		log.Println("BGB-003", lErr.Error())
		lChargeResp.Status = common.ErrorCode
		lChargeResp.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("UpdateChage-(-)")

}

func updateChargeInDB(lChargeResp *UpdateChargeResp, lCharge *UpdateChargeRec) {

	log.Println("updateChargeInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {

		log.Println("BUCIDB-001", lErr.Error())
		lChargeResp.ErrMsg = lErr.Error()
		lChargeResp.Status = common.ErrorCode

	} else {

		fmt.Println(lChargeResp)
		config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
		AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
		// lCharge.CreatedBy = "Admin: " + AdminId
		// lCharge.CreatedAt = time.Now()
		lCharge.UpdatedBy = "Admin: " + AdminId
		lCharge.UpdatedAt = time.Now()

		lResult := lGormDB.Table("st_919_config_charges_table").Where("charge_id = ?", lCharge.ChargeID).UpdateColumns(&lCharge)

		if lResult.Error != nil {
			log.Println("BUBINDB-002", lResult.Error)
			lChargeResp.Status = common.ErrorCode
			lChargeResp.ErrMsg = lResult.Error.Error()

		} else {
			lChargeResp.Charge = *lCharge
			lChargeResp.Status = common.SuccessCode
		}
	}

	log.Println("updateChargeInDB-(-)")

}
