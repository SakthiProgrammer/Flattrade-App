package user

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"log"
	"net/http"
)

type GetUserRec struct {
	ID        int    `json:"id" gorm:"column:user_id"`
	UserName  string `json:"user_name" gorm:"column:username"`
	Role      string `json:"role" gorm:"column:role"`
	Status    string `json:"status" gorm:"column:status"`
	Password  string `json:"password" gorm:"column:password"`
	CreatedBy string `json:"created_by" gorm:"column:Created_By"`
	CreatedAt string `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy string `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt string `json:"updated_at" gorm:"column:Updated_At"`
}

type GetUserResp struct {
	UserArr []GetUserRec `json:"user_details"`
	ErrMsg  string       `json:"errMsg"`
	Status  string       `json:"status"`
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("GetUsers-(+)")

	var lUserResp GetUserResp

	if r.Method == http.MethodGet {
		lUserResp.Status = common.SuccessCode

		lGormDb, lErr := gormdb.GormDBConnection()

		lSql, _ := lGormDb.DB()

		defer lSql.Close()

		if lErr != nil {
			log.Println("SGS-001", lErr.Error())
			lUserResp.Status = common.ErrorCode
			lUserResp.ErrMsg = lErr.Error()

		} else {

			lResult := lGormDb.Table("st_918_config_users_table").Find(&lUserResp.UserArr)

			if lResult.Error != nil {
				log.Println("SGS-002", lResult.Error)
				lUserResp.ErrMsg = lResult.Error.Error()
				lUserResp.Status = common.ErrorCode
			} else {
				log.Println("Successfully Get")
				lUserResp.Status = common.SuccessCode
				fmt.Printf("%+v", lUserResp)
			}
		}

	} else {
		log.Println("SGS-", "Invalid Method")
		lUserResp.ErrMsg = "Invalid Method"
		lUserResp.Status = common.ErrorCode
	}

	lData, lErr := json.Marshal(lUserResp)

	if lErr != nil {
		log.Println("SGS-", lErr.Error())
		lUserResp.ErrMsg = lErr.Error()
		lUserResp.Status = common.SuccessCode
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("GetUsers-(-)")

}
