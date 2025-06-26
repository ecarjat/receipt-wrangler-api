package routers

import (
	"receipt-wrangler/api/internal/handlers"
	"receipt-wrangler/api/internal/middleware"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/go-chi/chi/v5"
)

func BuildCurrencyRouter(tokenValidator *jwtmiddleware.JWTMiddleware) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.MoveJWTCookieToHeader, tokenValidator.CheckJWT)

	router.Get("/", handlers.GetAllCurrencies)
	router.Post("/", handlers.CreateCurrency)
	router.Get("/{currencyId}", handlers.GetCurrencyById)
	router.Put("/{currencyId}", handlers.UpdateCurrency)
	router.Delete("/{currencyId}", handlers.DeleteCurrency)

	return router
}
