package stock

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

type UpdateStockRec struct {
	ID         int     `json:"id" gorm:"column:stock_id"`
	StockName  string  `json:"stock_name" gorm:"column:stock_name"`
	StockPrice float64 `json:"stock_price" gorm:"column:stock_price"`
	Segment    string  `json:"segment" gorm:"column:segment"`
	ISIN       string  `json:"isin" gorm:"column:isin"`
	// CreatedBy  string    `json:"created_by" gorm:"column:created_by"`
	// CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedBy string    `json:"-" gorm:"column:updated_by"`
	UpdatedAt time.Time `json:"-" gorm:"column:updated_at"`
}

type UpdateStockResp struct {
	Stock  UpdateStockRec `json:"stock_details"`
	ErrMsg string         `json:"errMsg"`
	Status string         `json:"status"`
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "PUT, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "Stock, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("UpdateStock- (+)")

	var lStock UpdateStockRec
	var lStockResp UpdateStockResp

	lStockResp.Status = common.SuccessCode
	if r.Method == http.MethodPut {

		lData, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			log.Println("SUPS-001", lErr.Error())
			lStockResp.ErrMsg = lErr.Error()
			lStockResp.Status = common.ErrorCode

		} else {
			lErr = json.Unmarshal(lData, &lStock)

			if lErr != nil {

				log.Println("SUPS-002", lErr.Error())
				lStockResp.ErrMsg = lErr.Error()
				lStockResp.Status = common.ErrorCode
			} else {

				if lErr != nil {
					log.Println("SUPS-003", lErr.Error())
					lStockResp.ErrMsg = lErr.Error()
					lStockResp.Status = common.ErrorCode

				} else {

					updateClientInDB(&lStockResp, &lStock)

				}

			}

		}

	} else {
		// w.Header().Set(http.S)
		lStockResp.ErrMsg = "Method Not Allowed!"
		lStockResp.Status = common.ErrorCode
	}

	lData, lErr := json.Marshal(lStockResp)

	if lErr != nil {
		log.Println("SUPS-006", lErr.Error())
		lStockResp.ErrMsg = lErr.Error()
		lStockResp.Status = common.ErrorCode

	} else {
		fmt.Fprint(w, string(lData))
	}

	log.Println("UpdateStock- (-)")
}

func updateClientInDB(lStockResp *UpdateStockResp, lStock *UpdateStockRec) {

	log.Println("updateClientInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {

		log.Println("SUPSINDB-001", lErr.Error())
		lStockResp.ErrMsg = lErr.Error()
		lStockResp.Status = common.ErrorCode

	} else {

		fmt.Println(lStockResp)
		config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
		AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
		// lStock.CreatedBy = "Admin: " + AdminId
		// lStock.CreatedAt = time.Now()
		lStock.UpdatedBy = "Admin: " + AdminId
		lStock.UpdatedAt = time.Now()

		lResult := lGormDB.Table("st_918_stock_table").Where("stock_id = ?", lStock.ID).UpdateColumns(&lStock)

		if lResult.Error != nil {
			log.Println("SUPSINDB-002", lResult.Error)
			lStockResp.Status = common.ErrorCode
			lStockResp.ErrMsg = lResult.Error.Error()

		} else {
			lStockResp.Stock = *lStock
			lStockResp.Status = common.SuccessCode
		}
	}

	log.Println("updateClientInDB-(-)")

}
