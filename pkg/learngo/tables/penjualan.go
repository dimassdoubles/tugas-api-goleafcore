package tables

import (
	"git.solusiteknologi.co.id/goleaf/goleafcore/glentity"
	"github.com/shopspring/decimal"
)

type Penjualan struct {
	PenjualanId     int64           `json:"penjualanId" example="10"`
	TotalPenjualan  decimal.Decimal `json:"totalPenjualan" example="120000"`
	TotalPembayaran decimal.Decimal `json:"totalPembayaran" example="150000"`
	TotalKembalian  decimal.Decimal `json:"totalKembalian" example="30000"`

	glentity.BaseEntity
}
