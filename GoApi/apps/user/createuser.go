package user

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

type CreateUserRec struct {
	ID        int       `json:"id" gorm:"column:user_id"`
	UserName  string    `json:"user_name" gorm:"column:username"`
	Role      string    `json:"role" gorm:"column:role"`
	Status    string    `json:"status" gorm:"column:status"`
	Password  string    `json:"password" gorm:"column:password"`
	CreatedBy string    `json:"created_by" gorm:"column:Created_By"`
	CreatedAt time.Time `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy string    `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:Updated_At"`
}

type CreateUserResp struct {
	User   CreateUserRec `json:"user_details"`
	ErrMsg string        `json:"errMsg"`
	Status string        `json:"status"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "PRODUCTS, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("CreateUser - (+)")
	var lUser CreateUserRec
	var lUserResp CreateUserResp
	lUserResp.Status = common.SuccessCode

	if r.Method == http.MethodPost {

		lData, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			log.Println("UCU-001", lErr.Error())
			lUserResp.ErrMsg = lErr.Error()
			lUserResp.Status = common.ErrorCode
		} else {
			lErr = json.Unmarshal(lData, &lUser)

			if lErr != nil {
				log.Println("UCU-002", lErr.Error())
				lUserResp.ErrMsg = lErr.Error()
				lUserResp.Status = common.ErrorCode
			} else {

				if lErr != nil {
					log.Println("UCU-003", lErr.Error())
					lUserResp.ErrMsg = lErr.Error()
					lUserResp.Status = common.ErrorCode

				} else {
					createUserInDB(&lUserResp, &lUser)
				}

			}

		}

	} else {
		lUserResp.ErrMsg = "Invalid Method"
		lUserResp.Status = common.ErrorCode
	}

	lData, lErr := json.Marshal(lUserResp)

	if lErr != nil {
		log.Println("UCU-006", lErr.Error())
		lUserResp.ErrMsg = lErr.Error()
		lUserResp.Status = common.ErrorCode

	} else {
		fmt.Fprint(w, string(lData))
	}

	log.Println("CreateUser - (-)")

}

func createUserInDB(lUserResp *CreateUserResp, lUser *CreateUserRec) {

	log.Println("createUserInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {
		log.Println("UCU-004", lErr.Error())
		lUserResp.ErrMsg = lErr.Error()
		lUserResp.Status = common.ErrorCode
	} else {

		config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
		AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
		lName := "Admin: " + AdminId
		lCurrTime := time.Now()
		lUser.CreatedBy = lName
		lUser.CreatedAt = lCurrTime
		lUser.UpdatedBy = lName
		lUser.UpdatedAt = lCurrTime

		lResult := lGormDB.Table("st_918_config_users_table").Create(&lUser)

		if lResult.Error != nil {
			log.Println("UCU-005", lResult.Error)
			lUserResp.ErrMsg = lErr.Error()
			lUserResp.Status = common.ErrorCode

		} else {
			fmt.Println(lResult.RowsAffected)
			lUserResp.User = *lUser
		}
	}

	log.Println("createUserInDB-(-)")

}
