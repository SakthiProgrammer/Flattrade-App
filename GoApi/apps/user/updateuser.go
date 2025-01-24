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

type UpdateUserRec struct {
	ID       int    `json:"id" gorm:"column:user_id"`
	UserName string `json:"user_name" gorm:"column:username"`
	Role     string `json:"role" gorm:"column:role"`
	Status   string `json:"status" gorm:"column:status"`
	Password string `json:"password" gorm:"column:password"`
	// CreatedBy string `json:"created_by" gorm:"column:Created_By"`
	// CreatedAt string `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy string    `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:Updated_At"`
}

type UpdateUserResp struct {
	User   UpdateUserRec `json:"user_details"`
	ErrMsg string        `json:"errMsg"`
	Status string        `json:"status"`
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("UpdateUser-(+)")

	var lUserResp UpdateUserResp
	var lUser UpdateUserRec

	if r.Method == http.MethodPut {

		lUserResp.Status = common.SuccessCode

		lBody, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			log.Println("UUPU-001", lErr.Error())
			lUserResp.Status = common.ErrorCode
			lUserResp.ErrMsg = lErr.Error()

		} else {

			lErr = json.Unmarshal(lBody, &lUser)

			if lErr != nil {
				log.Println("UUPU-002", lErr.Error())
				lUserResp.Status = common.ErrorCode
				lUserResp.ErrMsg = lErr.Error()
			} else {
				updateUserInDB(&lUserResp, &lUser)
			}
		}

	} else {

		lUserResp.ErrMsg = "Invalid Method"
		lUserResp.Status = common.ErrorCode

	}

	lData, lErr := json.Marshal(lUserResp)

	if lErr != nil {
		log.Println("BGB-003", lErr.Error())
		lUserResp.Status = common.ErrorCode
		lUserResp.ErrMsg = "Server Error"
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("UpdateUser-(-)")

}

func updateUserInDB(lUserResp *UpdateUserResp, lUser *UpdateUserRec) {

	log.Println("updateUserInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {

		log.Println("UUPUIDB-001", lErr.Error())
		lUserResp.ErrMsg = lErr.Error()
		lUserResp.Status = common.ErrorCode

	} else {

		fmt.Println(lUserResp)
		config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
		AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
		// lUser.CreatedBy = "Admin: " + AdminId
		// lUser.CreatedAt = time.Now()
		lUser.UpdatedBy = "Admin: " + AdminId
		lUser.UpdatedAt = time.Now()

		lResult := lGormDB.Table("st_918_config_users_table").Where("user_id = ?", lUser.ID).UpdateColumns(&lUser)

		if lResult.Error != nil {
			log.Println("UUPUIDB-002", lResult.Error)
			lUserResp.Status = common.ErrorCode
			lUserResp.ErrMsg = lResult.Error.Error()

		} else {

			lUserResp.User = *lUser
			lUserResp.Status = common.SuccessCode
		}
	}

	log.Println("updateUserInDB-(-)")

}
