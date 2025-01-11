package brokerage

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"log"
	"net/http"
)

type GetChargeRec struct {
	ChargeID         uint    `json:"charge_id" gorm:"column:charge_id"`
	ChargeType       string  `json:"charge_type" gorm:"column:charge_type"`
	ChargePercentage float64 `json:"charge_percentage" gorm:"column:charge_percentage"`
	EffectiveDate    string  `json:"effective_date" gorm:"column:effective_date"`
	CreatedBy        string  `json:"created_by" gorm:"column:Created_By"`
	CreatedAt        string  `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy        string  `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt        string  `json:"updated_at" gorm:"column:Updated_At"`
}

type GetChargeResp struct {
	ChargeArr []GetChargeRec `json:"charge_details"`
	ErrMsg    string         `json:"errMsg"`
	Status    string         `json:"status"`
}

func GetCharges(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("GetCharges-(+)")

	var lChargeResp GetChargeResp
	lChargeResp.Status = common.SuccessCode

	if r.Method == http.MethodGet {

		lGormDb, lErr := gormdb.GormDBConnection()

		lSql, _ := lGormDb.DB()

		defer lSql.Close()

		if lErr != nil {

			log.Println("BGCS-001", lErr.Error())
			lChargeResp.Status = common.ErrorCode
			lChargeResp.ErrMsg = lErr.Error()

		} else {

			lResult := lGormDb.Table("st_919_config_charges_table").Find(&lChargeResp.ChargeArr)

			if lResult.Error != nil {

				log.Println("BGCS-002", lResult.Error)
				lChargeResp.Status = common.ErrorCode
				lChargeResp.ErrMsg = lResult.Error.Error()

			} else {
				log.Println("Success")
				lChargeResp.Status = common.SuccessCode
				fmt.Printf("%+v", lChargeResp)
			}

		}

	} else {

		lChargeResp.Status = common.ErrorCode
		lChargeResp.ErrMsg = "Invalid Method"
	}

	lData, lErr := json.Marshal(lChargeResp)

	if lErr != nil {
		log.Println("BGB-003", lErr.Error())
		lChargeResp.Status = common.ErrorCode
		lChargeResp.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}
	log.Println("GetCharges-(-)")
}
