package tables

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	ProductId   int64           `json:"productId" example:"10"`
	ProductCode string          `json:"productCode" example:"P001"`
	ProductName string          `json:"productName" example:"Indomie Goreng"`
	Price       decimal.Decimal `json:"price" example:"3000"`
	Version     int64           `json:"version" example:"0"`
}
