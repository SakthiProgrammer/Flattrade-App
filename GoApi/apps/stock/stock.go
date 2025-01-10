package stock

import "time"

type Stock struct {
	ID         int       `json:"id" gorm:"column:Id"`
	StockName  string    `json:"stock_name" gorm:"column:Stock_Name"`
	StockPrice float64   `json:"stock_price" gorm:"column:Stock_Price"`
	Segment    string    `json:"segment" gorm:"column:Segment"`
	ISIN       string    `json:"isin" gorm:"column:ISIN"`
	CreatedBy  string    `json:"created_by" gorm:"column:Created_By"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:Created_At"`
	UpdatedBy  string    `json:"updated_by" gorm:"column:Updated_By"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:Updated_At"`
}

type StockResp struct {
	StockDetailsArr []Stock `json:"stock_details"`
	ErrMsg          string  `json:"errMsg"`
	Status          string  `json:"status"`
}
