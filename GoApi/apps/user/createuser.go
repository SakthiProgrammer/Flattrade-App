package user

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"column:Id"`
	UserName  string    `json:"user_name" gorm:"column:User_Name"`
	Role      string    `json:"role" gorm:"column:Role"`
	Status    string    `json:"status" gorm:"column:Status"`
	CreatedBy string    `json:"created_by" gorm:"column:Created_By"`
	CreatedAt time.Time `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy string    `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:Updated_At"`
}

type UserResp struct {
	UserArr []User `json:"user_details"`
	ErrMsg  string `json:"errMsg"`
	Status  string `json:"status"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "PRODUCTS, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("CreateUser - (+)")
	var lUser User
	var lUserResp UserResp
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

func createUserInDB(lUserResp *UserResp, lUser *User) {

	log.Println("createUserInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {
		log.Println("UCU-004", lErr.Error())
		lUserResp.ErrMsg = lErr.Error()
		lUserResp.Status = common.ErrorCode
	} else {

		lResult := lGormDB.Table("st_918_Config_User_Table").Create(&lUser)

		if lResult.Error != nil {
			log.Println("UCU-005", lResult.Error)
			lUserResp.ErrMsg = lErr.Error()
			lUserResp.Status = common.ErrorCode

		} else {
			fmt.Println(lResult.RowsAffected)
			lUserResp.UserArr = append(lUserResp.UserArr, *lUser)
		}
	}

	log.Println("createUserInDB-(-)")

}
