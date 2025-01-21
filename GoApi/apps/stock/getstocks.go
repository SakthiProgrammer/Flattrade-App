package stock

import (
	"encoding/json"
	"flattrade/apps/DBConnection/gormdb"
	"flattrade/common"
	"fmt"
	"log"
	"net/http"
)

type GetStockRec struct {
	ID         int     `json:"id" gorm:"column:stock_id"`
	StockName  string  `json:"stock_name" gorm:"column:stock_name"`
	StockPrice float64 `json:"stock_price" gorm:"column:stock_price"`
	Segment    string  `json:"segment" gorm:"column:segment"`
	ISIN       string  `json:"isin" gorm:"column:isin"`
	CreatedBy  string  `json:"created_by" gorm:"column:created_by"`
	CreatedAt  string  `json:"created_at" gorm:"column:created_at"`
	UpdatedBy  string  `json:"updated_by" gorm:"column:updated_by"`
	UpdatedAt  string  `json:"updated_at" gorm:"column:updated_at"`
}

type GetStockResp struct {
	StockDetailsArr []GetStockRec `json:"stock_details"`
	ErrMsg          string        `json:"errMsg"`
	Status          string        `json:"status"`
}

func GetStocks(w http.ResponseWriter, r *http.Request) {

	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "ADMIN, CLIENT, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(w).Header().Set("Content-Type", "application/json")

	log.Println("GetStocks-(+)")

	var lStockResp GetStockResp

	// var lStock Stock

	if r.Method == http.MethodGet {
		lStockResp.Status = common.SuccessCode

		lGormDb, lErr := gormdb.GormDBConnection()

		lSql, _ := lGormDb.DB()

		defer lSql.Close()

		if lErr != nil {
			log.Println("SGS-001", lErr.Error())
			lStockResp.Status = common.ErrorCode
			lStockResp.ErrMsg = lErr.Error()

		} else {

			lResult := lGormDb.Table("st_918_stock_table").Find(&lStockResp.StockDetailsArr)

			if lResult.Error != nil {
				log.Println("SGS-002", lResult.Error)
				lStockResp.ErrMsg = lResult.Error.Error()
				lStockResp.Status = common.ErrorCode
			} else {
				log.Println("Successfully Get")
				lStockResp.Status = common.SuccessCode
				fmt.Printf("%+v", lStockResp)
			}
		}

	} else {
		log.Println("SGS-", "Invalid Method")
		lStockResp.ErrMsg = "Invalid Method"
		lStockResp.Status = common.ErrorCode
	}

	lData, lErr := json.Marshal(lStockResp)

	if lErr != nil {
		log.Println("SGS-", lErr.Error())
		lStockResp.ErrMsg = lErr.Error()
		lStockResp.Status = common.SuccessCode
	} else {
		fmt.Fprintf(w, string(lData))
	}

	log.Println("GetStocks-(-)")

}
