package bank

type BaseBank struct {
	// ID         int    `json:"id" gorm:"column:Id"`
	BankName   string `json:"bank_name" gorm:"column:Bank_Name"`
	BranchName string `json:"branch_name" gorm:"column:Branch_name"`
	IFSCCODE   string `json:"ifsc_code" gorm:"column:Ifsc_code"`
	Address    string `json:"address" gorm:"column:Address"`
	// CreatedBy  string `json:"created_by" gorm:"column:Created_By"`
	// CreatedAt  string `json:"created_at" gorm:"column:Created_At"`
	// UpdatedBy  string `json:"updated_by" gorm:"column:Updated_By"`
	// UpdatedAt  string `json:"updated_at" gorm:"column:Updated_At"`
}
