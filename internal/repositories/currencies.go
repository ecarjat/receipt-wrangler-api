package repositories

import (
	"gorm.io/gorm"
	"receipt-wrangler/api/internal/models"
)

type CurrencyRepository struct {
	BaseRepository
}

func NewCurrencyRepository(tx *gorm.DB) CurrencyRepository {
	return CurrencyRepository{BaseRepository: BaseRepository{
		DB: GetDB(),
		TX: tx,
	}}
}

func (r CurrencyRepository) GetAllCurrencies() ([]models.Currency, error) {
	db := r.GetDB()
	var currencies []models.Currency
	if err := db.Find(&currencies).Error; err != nil {
		return nil, err
	}
	return currencies, nil
}

func (r CurrencyRepository) GetCurrencyById(id uint) (models.Currency, error) {
	db := r.GetDB()
	var currency models.Currency
	if err := db.First(&currency, id).Error; err != nil {
		return models.Currency{}, err
	}
	return currency, nil
}

func (r CurrencyRepository) CreateCurrency(currency models.Currency) (models.Currency, error) {
	db := r.GetDB()
	if err := db.Create(&currency).Error; err != nil {
		return models.Currency{}, err
	}
	return currency, nil
}

func (r CurrencyRepository) UpdateCurrency(currency models.Currency) (models.Currency, error) {
	db := r.GetDB()
	if err := db.Model(&currency).Select("Code", "Name").Where("id = ?", currency.ID).Updates(currency).Error; err != nil {
		return models.Currency{}, err
	}
	return currency, nil
}

func (r CurrencyRepository) DeleteCurrency(id uint) error {
	db := r.GetDB()
	if err := db.Delete(&models.Currency{}, id).Error; err != nil {
		return err
	}
	return nil
}
