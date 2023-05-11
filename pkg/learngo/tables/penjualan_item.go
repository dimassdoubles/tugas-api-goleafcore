package tables

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glentity"
	"github.com/shopspring/decimal"
)

type PenjualanItem struct {
	PenjualanItemId int64           `json:"penjualanItem" example:"10"`
	PenjualanId     int64           `json:"penjualanId" example:"05"`
	ProductId       int64           `json:"productId" example:"15"`
	Qty             decimal.Decimal `json:"qty" example:"20"`
	Price           decimal.Decimal `json:"price" example:"40000"`

	glentity.BaseEntity
}
