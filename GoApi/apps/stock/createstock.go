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

type CreateStockRec struct {
	ID         int       `json:"id" gorm:"column:stock_id"`
	StockName  string    `json:"stock_name" gorm:"column:stock_name"`
	StockPrice float64   `json:"stock_price" gorm:"column:stock_price"`
	Segment    string    `json:"segment" gorm:"column:segment"`
	ISIN       string    `json:"isin" gorm:"column:isin"`
	CreatedBy  string    `json:"created_by" gorm:"column:created_by"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedBy  string    `json:"updated_by" gorm:"column:updated_by"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at"`
}

type CreateStockResp struct {
	StockDetailsArr CreateStockRec `json:"stock_details"`
	ErrMsg          string         `json:"errMsg"`
	Status          string         `json:"status"`
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "PRODUCTS, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json")

	log.Println("CreateStock - (+)")
	var lStock CreateStockRec
	var lStockResp CreateStockResp
	lStockResp.Status = common.SuccessCode

	if r.Method == http.MethodPost {

		lData, lErr := io.ReadAll(r.Body)

		if lErr != nil {
			log.Println("SCS-001", lErr.Error())
			lStockResp.ErrMsg = lErr.Error()
			lStockResp.Status = common.ErrorCode
		} else {
			lErr = json.Unmarshal(lData, &lStock)

			if lErr != nil {
				log.Println("SCS-002", lErr.Error())
				lStockResp.ErrMsg = lErr.Error()
				lStockResp.Status = common.ErrorCode
			} else {

				if lErr != nil {
					log.Println("SCS-003", lErr.Error())
					lStockResp.ErrMsg = lErr.Error()
					lStockResp.Status = common.ErrorCode

				} else {
					createStocksInDB(&lStockResp, &lStock)
				}

			}

		}

	} else {
		lStockResp.ErrMsg = "Invalid Method"
		lStockResp.Status = common.ErrorCode
	}

	lData, lErr := json.Marshal(lStockResp)

	if lErr != nil {
		log.Println("SCS-006", lErr.Error())
		lStockResp.ErrMsg = lErr.Error()
		lStockResp.Status = common.ErrorCode

	} else {
		fmt.Fprint(w, string(lData))
	}

	log.Println("CreateStock - (-)")

}

func createStocksInDB(lStockResp *CreateStockResp, lStock *CreateStockRec) {

	log.Println("CreateStocksInDB-(+)")

	lGormDB, lErr := gormdb.GormDBConnection()

	lSql, _ := lGormDB.DB()

	defer lSql.Close()

	if lErr != nil {
		log.Println("SCS-004", lErr.Error())
		lStockResp.ErrMsg = lErr.Error()
		lStockResp.Status = common.ErrorCode
	} else {

		config := genpkg.ReadTomlConfig("./toml/dbconfig.toml")
		AdminId := fmt.Sprintf("%v", config.(map[string]interface{})["AdminId"])
		lName := "Admin: " + AdminId
		lCurrTime := time.Now()
		lStock.CreatedBy = lName
		lStock.CreatedAt = lCurrTime
		lStock.UpdatedBy = lName
		lStock.UpdatedAt = lCurrTime

		lResult := lGormDB.Table("st_918_stock_table").Create(&lStock)

		if lResult.Error != nil {
			log.Println("SCS-005", lResult.Error)
			lStockResp.ErrMsg = lErr.Error()
			lStockResp.Status = common.ErrorCode

		} else {
			fmt.Println(lResult.RowsAffected)
			lStockResp.StockDetailsArr = *lStock
		}
	}

	log.Println("CreateStocksInDB-(-)")
}
