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

// AuthorContextKey is a custom type for context keys to avoid collisions
type AuthorContextKey string

const (
	// AuthorRequestKey is the key to store/retrieve AuthorRequest from context
	AuthorRequestKey AuthorContextKey = "authorRequest"
)

// AuthorValidatorMiddleware validates the request body for author creation/update
func AuthorValidatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var authorRequest request.AuthorRequest
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields() // Return an error if the JSON contains unknown fields

		if err := decoder.Decode(&authorRequest); err != nil {
			helper.Response(w, http.StatusBadRequest, "Invalid request payload: "+err.Error(), nil)
			return
		}

		// Close the body after decoding
		// r.Body.Close() // Do not close the body here, it will be closed by the handler if it needs to be reread

		validate := validator.New()
		if err := validate.Struct(authorRequest); err != nil {
			validationErrors := err.(validator.ValidationErrors)

			// if err := validate.Struct(authorRequest); err != nil {
			// 	validationErrors := make(map[string]string)
			// 	for _, err := range err.(validator.ValidationErrors) {
			// 		validationErrors[err.Field()] = "Validation failed on field '" + err.Field() + "' with tag '" + err.Tag() + "'"
			// 	}

			// 	helper.Response(w, http.StatusBadRequest, fmt.Sprintf("Validation error: %v", validationErrors), nil)
			// 	return
			// }
			// ============

			if len(validationErrors) > 0 {
				// Ambil error pertama (index 0)
				firstErr := validationErrors[0]
				msg := fmt.Sprintf("Validation failed on field '%s' with tag '%s'", firstErr.Field(), firstErr.Tag())
				helper.Response(w, http.StatusBadRequest, msg, nil)
				return
			}

			// Jika tidak ada error, fallback ke pesan umum
			helper.Response(w, http.StatusBadRequest, "Validation error", nil)
			return
		}

		// Store the validated authorRequest in the request context
		ctx := context.WithValue(r.Context(), AuthorRequestKey, authorRequest)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
