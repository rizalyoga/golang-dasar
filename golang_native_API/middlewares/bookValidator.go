package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"golang_native_api/helper"
	"golang_native_api/request"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type BookContextKey string

const BookRequestKey BookContextKey = "bookRequestKey"

func BookValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bookRequest request.BookRequest
		decoder := json.NewDecoder(r.Body)

		// fungsi agar tidak menerima field yang tidak dikenal dalam json
		decoder.DisallowUnknownFields()

		if err := decoder.Decode(&bookRequest); err != nil {
			helper.Response(w, http.StatusBadRequest, "Invalid request payload: "+err.Error(), nil)
			return
		}

		validate := validator.New()
		if err := validate.Struct(bookRequest); err != nil {
			validationErrors := err.(validator.ValidationErrors)

			if len(validationErrors) > 0 {
				firstErr := validationErrors[0]
				msg := fmt.Sprintf("Validation failed on field '%s' with tag '%s'", firstErr.Field(), firstErr.Tag())
				helper.Response(w, http.StatusBadRequest, msg, nil)
				return
			}

			helper.Response(w, http.StatusBadRequest, "Validation error", nil)
			return
		}

		ctx := context.WithValue(r.Context(), BookRequestKey, bookRequest)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
