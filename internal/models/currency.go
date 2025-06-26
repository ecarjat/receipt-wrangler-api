package models

import (
	"encoding/json"
	"net/http"
	"receipt-wrangler/api/internal/utils"
)

type Currency struct {
	BaseModel
	Code string `gorm:"not null;uniqueIndex" json:"code"`
	Name string `gorm:"not null" json:"name"`
}

func (currency *Currency) LoadDataFromRequest(w http.ResponseWriter, r *http.Request) error {
	bytes, err := utils.GetBodyData(w, r)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, currency); err != nil {
		return err
	}
	return nil
}
