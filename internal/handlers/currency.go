package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"receipt-wrangler/api/internal/constants"
	"receipt-wrangler/api/internal/models"
	"receipt-wrangler/api/internal/repositories"
	"receipt-wrangler/api/internal/structs"
	"receipt-wrangler/api/internal/utils"
)

func GetAllCurrencies(w http.ResponseWriter, r *http.Request) {
	handler := structs.Handler{
		ErrorMessage: "Error retrieving currencies",
		Writer:       w,
		Request:      r,
		UserRole:     models.USER,
		ResponseType: constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			repo := repositories.NewCurrencyRepository(nil)
			currencies, err := repo.GetAllCurrencies()
			if err != nil {
				return http.StatusInternalServerError, err
			}
			bytes, err := utils.MarshalResponseData(&currencies)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
			return 0, nil
		},
	}
	HandleRequest(handler)
}

func CreateCurrency(w http.ResponseWriter, r *http.Request) {
	handler := structs.Handler{
		ErrorMessage: "Error creating currency",
		Writer:       w,
		Request:      r,
		UserRole:     models.ADMIN,
		ResponseType: constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			currency := models.Currency{}
			if err := currency.LoadDataFromRequest(w, r); err != nil {
				return http.StatusInternalServerError, err
			}
			repo := repositories.NewCurrencyRepository(nil)
			created, err := repo.CreateCurrency(currency)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			data, err := json.Marshal(created)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return 0, nil
		},
	}
	HandleRequest(handler)
}

func GetCurrencyById(w http.ResponseWriter, r *http.Request) {
	handler := structs.Handler{
		ErrorMessage: "Error retrieving currency",
		Writer:       w,
		Request:      r,
		UserRole:     models.USER,
		ResponseType: constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			idParam := chi.URLParam(r, "currencyId")
			id, err := utils.StringToUint(idParam)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			repo := repositories.NewCurrencyRepository(nil)
			currency, err := repo.GetCurrencyById(id)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			data, err := json.Marshal(currency)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return 0, nil
		},
	}
	HandleRequest(handler)
}

func UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	handler := structs.Handler{
		ErrorMessage: "Error updating currency",
		Writer:       w,
		Request:      r,
		UserRole:     models.ADMIN,
		ResponseType: constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			idParam := chi.URLParam(r, "currencyId")
			id, err := utils.StringToUint(idParam)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			currency := models.Currency{}
			if err := currency.LoadDataFromRequest(w, r); err != nil {
				return http.StatusInternalServerError, err
			}
			currency.ID = id
			repo := repositories.NewCurrencyRepository(nil)
			updated, err := repo.UpdateCurrency(currency)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			data, err := json.Marshal(updated)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			w.WriteHeader(http.StatusOK)
			w.Write(data)
			return 0, nil
		},
	}
	HandleRequest(handler)
}

func DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	handler := structs.Handler{
		ErrorMessage: "Error deleting currency",
		Writer:       w,
		Request:      r,
		UserRole:     models.ADMIN,
		ResponseType: constants.ApplicationJson,
		HandlerFunction: func(w http.ResponseWriter, r *http.Request) (int, error) {
			idParam := chi.URLParam(r, "currencyId")
			id, err := utils.StringToUint(idParam)
			if err != nil {
				return http.StatusInternalServerError, err
			}
			repo := repositories.NewCurrencyRepository(nil)
			if err := repo.DeleteCurrency(id); err != nil {
				return http.StatusInternalServerError, err
			}
			w.WriteHeader(http.StatusOK)
			return 0, nil
		},
	}
	HandleRequest(handler)
}
